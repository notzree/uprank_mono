package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank-backend/main-backend/ent/freelancer"
	"github.com/notzree/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank-backend/main-backend/ent/user"
	"github.com/notzree/uprank-backend/main-backend/types"
)

func (s *Server) CreateFreelancers(w http.ResponseWriter, r *http.Request) error {

	claims, _ := clerk.SessionClaimsFromContext(r.Context())
	user_id := claims.Subject
	job_id := chi.URLParam(r, "job_id")

	//check if job exists and belongs to user
	_, getJobErr := s.ent.Job.Query().
		Where(
			job.IDEQ(job_id),
			job.HasUserWith(user.IDEQ(user_id)),
		).
		Only(context.Background())
	if getJobErr != nil && ent.IsNotFound(getJobErr) {
		return ResourceMisMatch()
	}

	var req []types.CreateFreelancersRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode request", "err", err)
		return InvalidJSON()
	}
	defer r.Body.Close()

	for _, freelancer := range req {
		if errors := freelancer.Validate(); len(errors) > 0 {
			return InvalidRequestData(errors)
		}
	}
	attachements := make([][]*ent.AttachmentRef, len(req))

	freelancers, createFreelancerErr := s.ent.Freelancer.MapCreateBulk(req, func(c *ent.FreelancerCreate, i int) {
		freelancer := req[i]
		//guaranteed to be valid by freelancer.Validate()
		parsed_fixed_charge_amount, _ := strconv.ParseFloat(freelancer.Fixed_charge_amount, 64)
		parse_hourly_charge_amount, _ := strconv.ParseFloat(freelancer.Hourly_charge_amount, 64)
		c.SetID(freelancer.Url).
			SetName(freelancer.Name).
			SetTitle(freelancer.Title).
			SetDescription(freelancer.Description).
			SetCity(freelancer.Location.City).
			SetCountry(freelancer.Location.Country).
			SetTimezone(freelancer.Location.Timezone).
			SetCv(freelancer.Cv).
			SetAiReccomended(freelancer.Ai_reccomended).
			SetFixedChargeAmount(parsed_fixed_charge_amount).
			SetFixedChargeCurrency(freelancer.Fixed_charge_currency).
			SetHourlyChargeAmount(parse_hourly_charge_amount).
			SetHourlyChargeCurrency(freelancer.Hourly_charge_currency).
			SetInvited(freelancer.Invited).
			SetPhotoURL(freelancer.Photo_url).
			SetRecentHours(freelancer.Recent_hours).
			SetTotalHours(freelancer.Total_hours).
			SetTotalPortfolioItems(freelancer.Total_portfolio_items).
			SetTotalPortfolioV2Items(freelancer.Total_portfolio_v2_items).
			SetUpworkTotalFeedback(freelancer.Total_feedback).
			SetUpworkRecentFeedback(freelancer.Recent_feedback).
			SetUpworkTopRatedStatus(freelancer.Top_rated_status).
			SetUpworkTopRatedPlusStatus(freelancer.Top_rated_plus_status).
			SetUpworkSponsored(freelancer.Sponsored).
			SetUpworkJobSuccessScore(freelancer.Job_success_score).
			SetUpworkReccomended(freelancer.Reccomended).
			SetSkills(freelancer.Skills).
			SetAverageRecentEarnings(freelancer.Earnings_info.Average_recent_earnings).
			SetCombinedAverageRecentEarnings(freelancer.Earnings_info.Combined_average_recent_earnings).
			SetCombinedRecentEarnings(freelancer.Earnings_info.Combined_recent_earnings).
			SetCombinedTotalEarnings(freelancer.Earnings_info.Combined_total_earnings).
			SetCombinedTotalRevenue(freelancer.Earnings_info.Combined_total_revenue).
			SetRecentEarnings(freelancer.Earnings_info.Recent_earnings).
			SetTotalRevenue(freelancer.Earnings_info.Total_revenue).
			SetJobID(job_id)
	}).Save(r.Context())
	if createFreelancerErr != nil {
		return createFreelancerErr
	}

	//TODO: CREATE JOB HISTORIES
	//need to fix type so the json parses correctly, then handle the creation of job histories
	for i, freelancer := range req {
		if len(freelancer.Attachements) == 0 {
			attachements[i] = nil
			continue
		}
		attch, createAttachementErr := s.ent.AttachmentRef.MapCreateBulk(freelancer.Attachements, func(c *ent.AttachmentRefCreate, j int) {
			c.SetName(freelancer.Attachements[j].Name).
				SetLink(freelancer.Attachements[j].Link).
				SetFreelancerID(freelancers[i].ID)
		}).Save(r.Context())

		if createAttachementErr != nil {
			return createAttachementErr
		}

		attachements[i] = attch
	}

	return writeJSON(w, http.StatusCreated, freelancers)
}

// this shit doesn't work lol. I prob need to use CreateBulk with upserts again gg
func (s *Server) UpdateFreelancers(w http.ResponseWriter, r *http.Request) error {
	claims, _ := clerk.SessionClaimsFromContext(r.Context())
	user_id := claims.Subject
	job_id := chi.URLParam(r, "job_id")

	//check if job exists and belongs to user
	job, getJobErr := s.ent.Job.Query().
		Where(
			job.IDEQ(job_id),
			job.HasUserWith(user.IDEQ(user_id)),
		).QueryFreelancers().All(r.Context())
	if getJobErr != nil && ent.IsNotFound(getJobErr) {
		return ResourceMisMatch()
	}

	var req []types.CreateFreelancersRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode request", "err", err)
		return InvalidJSON()
	}
	defer r.Body.Close()

	for _, freelancer := range req {
		if errors := freelancer.Validate(); len(errors) > 0 {
			return InvalidRequestData(errors)
		}
	}

	//We need to compare with the freelancers currently in the db.
	//We need to split them into 3 categories.
	// 1. Freelancers that are already in the db and need to be updated (don't need to update attachements or job histories)
	// 2. Freelancers that are not in the db but in the request and need to be created (with attachements + job histories created)
	// 3. Freelancers that are in the db but not in the request and need to be deleted (with attachements + job histories deleted)

	//Plan: find the intersections of the two sets, then for each array, remove the intersection. Then we are left with the 3 categories

	createFreelancerErr := s.ent.Freelancer.MapCreateBulk(req, func(c *ent.FreelancerCreate, i int) {
		freelancer := req[i]
		//guaranteed to be valid by freelancer.Validate()
		parsed_fixed_charge_amount, _ := strconv.ParseFloat(freelancer.Fixed_charge_amount, 64)
		parse_hourly_charge_amount, _ := strconv.ParseFloat(freelancer.Hourly_charge_amount, 64)
		c.SetID(freelancer.Url).
			SetName(freelancer.Name).
			SetTitle(freelancer.Title).
			SetDescription(freelancer.Description).
			SetCity(freelancer.Location.City).
			SetCountry(freelancer.Location.Country).
			SetTimezone(freelancer.Location.Timezone).
			SetCv(freelancer.Cv).
			SetAiReccomended(freelancer.Ai_reccomended).
			SetFixedChargeAmount(parsed_fixed_charge_amount).
			SetFixedChargeCurrency(freelancer.Fixed_charge_currency).
			SetHourlyChargeAmount(parse_hourly_charge_amount).
			SetHourlyChargeCurrency(freelancer.Hourly_charge_currency).
			SetInvited(freelancer.Invited).
			SetPhotoURL(freelancer.Photo_url).
			SetRecentHours(freelancer.Recent_hours).
			SetTotalHours(freelancer.Total_hours).
			SetTotalPortfolioItems(freelancer.Total_portfolio_items).
			SetTotalPortfolioV2Items(freelancer.Total_portfolio_v2_items).
			SetUpworkTotalFeedback(freelancer.Total_feedback).
			SetUpworkRecentFeedback(freelancer.Recent_feedback).
			SetUpworkTopRatedStatus(freelancer.Top_rated_status).
			SetUpworkTopRatedPlusStatus(freelancer.Top_rated_plus_status).
			SetUpworkSponsored(freelancer.Sponsored).
			SetUpworkJobSuccessScore(freelancer.Job_success_score).
			SetUpworkReccomended(freelancer.Reccomended).
			SetSkills(freelancer.Skills).
			SetAverageRecentEarnings(freelancer.Earnings_info.Average_recent_earnings).
			SetCombinedAverageRecentEarnings(freelancer.Earnings_info.Combined_average_recent_earnings).
			SetCombinedRecentEarnings(freelancer.Earnings_info.Combined_recent_earnings).
			SetCombinedTotalEarnings(freelancer.Earnings_info.Combined_total_earnings).
			SetCombinedTotalRevenue(freelancer.Earnings_info.Combined_total_revenue).
			SetRecentEarnings(freelancer.Earnings_info.Recent_earnings).
			SetTotalRevenue(freelancer.Earnings_info.Total_revenue).
			SetJobID(job_id)
	}).OnConflict().UpdateNewValues().Exec(r.Context())
	if createFreelancerErr != nil {
		return createFreelancerErr
	}

	return writeJSON(w, http.StatusCreated, nil)
}

// queueScraperJob queues the scraping job for the given job_id into our aws sqs queue for the scraper server to pick up
func queueScraperJob(s *Server, scrape_obj types.QueueScrapeFreelancersReqest) error {
	messageGroupId := "uprank-scraper-requests"
	freelancers_json, err := json.Marshal(scrape_obj)
	if err != nil {
		return err
	}
	message_body := string(freelancers_json)
	send_message_input := &sqs.SendMessageInput{
		MessageBody:    &message_body,
		MessageGroupId: &messageGroupId,
		QueueUrl:       &s.scraper_queue_url,
	}
	result, err := s.scraper_queue_client.SendMessage(context.TODO(), send_message_input)
	if err != nil {
		return err
	}
	slog.Info("Sent message to queue", "message_id", *result.MessageId)
	return nil
}

func (s *Server) TestQueue(w http.ResponseWriter, r *http.Request) error {
	var freelancers []types.ScrapeFreelancerData

	job_id := "1788207506953621504"
	queryErr := s.ent.Freelancer.Query().Where(freelancer.HasJobWith(job.IDEQ(job_id))).Select(freelancer.FieldID, freelancer.FieldID).Scan(context.Background(), &freelancers)
	if queryErr != nil {
		return queryErr
	}

	scrape_obj := types.QueueScrapeFreelancersReqest{
		Job_id:      job_id,
		Freelancers: freelancers,
	}

	queueErr := queueScraperJob(s, scrape_obj)
	if queueErr != nil {
		return queueErr
	}
	return writeJSON(w, http.StatusOK, "success")
}
