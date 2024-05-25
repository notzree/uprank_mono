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

		c.SetURL(freelancer.Url).
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
	err := queueScraperJob(s, job_id)
	if err != nil {
		return err
	}
	//itrating over freelancers to write attachements + edges into the db
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

// queueScraperJob queues the scraping job for the given job_id into our aws sqs queue for the scraper server to pick up
func queueScraperJob(s *Server, job_id string) error {
	messageGroupId := "my-group-id"
	message_body := job_id
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
	err := queueScraperJob(s, "test")
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, "success")
}
