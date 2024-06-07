package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/notzree/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank-backend/main-backend/ent/upworkfreelancer"
	"github.com/notzree/uprank-backend/main-backend/ent/user"
	"github.com/notzree/uprank-backend/main-backend/types"
)

func (s *Server) CreateFreelancers(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromContext(r.Context())
	if user_id_err != nil {
		return user_id_err
	}
	job_id := chi.URLParam(r, "job_id")
	var req []types.CreateUpworkFreelancerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode request", "err", err)
		return InvalidJSON()
	}
	defer r.Body.Close()

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

	for _, freelancer := range req {
		if errors := freelancer.Validate(); len(errors) > 0 {
			return InvalidRequestData(errors)
		}
	}

	freelancers, createFreelancerErr := s.ent.UpworkFreelancer.MapCreateBulk(req, func(c *ent.UpworkFreelancerCreate, i int) {
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
			AddJobIDs(job_id)
	}).Save(r.Context())
	if createFreelancerErr != nil {
		return createFreelancerErr
	}

	for _, freelancer := range req {

		if len(freelancer.Attachements) == 0 {
			continue
		}
		_, createAttachementErr := s.ent.AttachmentRef.MapCreateBulk(freelancer.Attachements, func(c *ent.AttachmentRefCreate, j int) {
			c.SetName(freelancer.Attachements[j].Name).
				SetLink(freelancer.Attachements[j].Link).
				SetFreelancerID(freelancer.Url)
		}).Save(r.Context())
		if createAttachementErr != nil {
			return createAttachementErr
		}

		if len(freelancer.Work_history) == 0 {
			continue
		}
		_, createWorkHistoryErr := s.ent.WorkHistory.MapCreateBulk(freelancer.Work_history, func(c *ent.WorkHistoryCreate, j int) {
			work_history := freelancer.Work_history[j]
			work_history_start_date, _ := time.Parse(time.RFC3339, work_history.Start_Date)
			work_history_end_date, _ := time.Parse(time.RFC3339, work_history.End_Date)
			c.SetTitle(work_history.Title).
				SetStartDate(work_history_start_date).
				SetEndDate(work_history_end_date).
				SetDescription(work_history.Description).
				SetClientFeedback(work_history.Client_Feedback).
				SetOverallRating(work_history.Client_Rating).
				SetClientTotalSpend(work_history.Client_Total_Spend).
				SetClientTotalHires(work_history.Client_Total_Hires).
				SetClientActiveHires(work_history.Client_Active_Hires).
				SetBudget(work_history.Budget).
				SetFreelancerEarnings(work_history.Total_Earned).
				SetClientCountry(work_history.Client_Location).
				SetFreelancerID(freelancer.Url)
		}).Save(r.Context())
		if createWorkHistoryErr != nil {
			return createWorkHistoryErr
		}

	}

	return writeJSON(w, http.StatusCreated, freelancers)
}

func (s *Server) UpdateFreelancers(w http.ResponseWriter, r *http.Request) error {
	user_id, user_id_err := s.authenticator.GetIdFromContext(r.Context())
	if user_id_err != nil {
		return user_id_err
	}
	job_id := chi.URLParam(r, "job_id")

	var (
		req                   []types.CreateUpworkFreelancerRequest
		freelancers_to_create []types.CreateUpworkFreelancerRequest
		freelancers_to_update []types.CreateUpworkFreelancerRequest
		freelancers_to_delete []string
	)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode request", "err", err)
		return InvalidJSON()
	}
	defer r.Body.Close()

	//check if job exists and belongs to user
	current_freelancers, getJobErr := s.ent.Job.Query().
		Where(
			job.IDEQ(job_id),
			job.HasUserWith(user.IDEQ(user_id)),
		).QueryFreelancers().All(r.Context())
	if getJobErr != nil && ent.IsNotFound(getJobErr) {
		return ResourceMisMatch()
	}

	for _, freelancer := range req {
		if errors := freelancer.Validate(); len(errors) > 0 {
			return InvalidRequestData(errors)
		}
	}

	incoming_freelancer_dict := make(map[string]types.CreateUpworkFreelancerRequest)
	for _, freelancer := range req {
		incoming_freelancer_dict[freelancer.Url] = freelancer
	}
	current_freelancer_dict := make(map[string]ent.UpworkFreelancer)
	for _, freelancer := range current_freelancers {
		current_freelancer_dict[freelancer.ID] = *freelancer
	}

	for _, freelancer := range current_freelancers {
		if _, ok := incoming_freelancer_dict[freelancer.ID]; ok {
			//the freelancer is in the db and in the request
			freelancers_to_update = append(freelancers_to_update, incoming_freelancer_dict[freelancer.ID])
		} else {
			//the freelancer is in the db but not in the request
			freelancers_to_delete = append(freelancers_to_delete, freelancer.ID)
		}
	}
	for _, freelancer := range req {
		if _, ok := current_freelancer_dict[freelancer.Url]; !ok {
			//the freelancer is in the request but not in the db
			freelancers_to_create = append(freelancers_to_create, freelancer)
		}
	}

	//create new freelancers
	_, createFreelancerErr := s.ent.UpworkFreelancer.MapCreateBulk(freelancers_to_create, func(c *ent.UpworkFreelancerCreate, i int) {
		freelancer := freelancers_to_create[i]
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
			AddJobIDs(job_id)

	}).Save(r.Context())

	if createFreelancerErr != nil {
		return createFreelancerErr
	}

	for _, freelancer := range freelancers_to_create {
		if len(freelancer.Attachements) != 0 {
			_, createAttachementErr := s.ent.AttachmentRef.MapCreateBulk(freelancer.Attachements, func(c *ent.AttachmentRefCreate, j int) {
				c.SetName(freelancer.Attachements[j].Name).
					SetLink(freelancer.Attachements[j].Link).
					SetFreelancerID(freelancer.Url)
			}).Save(r.Context())

			if createAttachementErr != nil {
				return createAttachementErr
			}
		}
		if len(freelancer.Work_history) != 0 {
			_, createWorkHistoryErr := s.ent.WorkHistory.MapCreateBulk(freelancer.Work_history, func(c *ent.WorkHistoryCreate, j int) {
				work_history := freelancer.Work_history[j]
				work_history_start_date, _ := time.Parse(time.RFC3339, work_history.Start_Date)
				work_history_end_date, _ := time.Parse(time.RFC3339, work_history.End_Date)
				c.SetTitle(work_history.Title).
					SetStartDate(work_history_start_date).
					SetEndDate(work_history_end_date).
					SetDescription(work_history.Description).
					SetClientFeedback(work_history.Client_Feedback).
					SetOverallRating(work_history.Client_Rating).
					SetClientTotalSpend(work_history.Client_Total_Spend).
					SetClientTotalHires(work_history.Client_Total_Hires).
					SetClientActiveHires(work_history.Client_Active_Hires).
					SetBudget(work_history.Budget).
					SetFreelancerEarnings(work_history.Total_Earned).
					SetClientCountry(work_history.Client_Location).
					SetFreelancerID(freelancer.Url)
			}).Save(r.Context())
			if createWorkHistoryErr != nil {
				return createWorkHistoryErr
			}
		}
	}

	//update freelancers
	updateFreelancersErr := s.ent.UpworkFreelancer.MapCreateBulk(freelancers_to_update, func(c *ent.UpworkFreelancerCreate, i int) {
		freelancer := freelancers_to_update[i]
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
			SetTotalRevenue(freelancer.Earnings_info.Total_revenue).AddJobIDs(job_id)
	}).OnConflictColumns(upworkfreelancer.FieldID).UpdateNewValues().Exec(r.Context())
	if updateFreelancersErr != nil {
		return updateFreelancersErr
	}
	for _, freelancer := range freelancers_to_update {
		if len(freelancer.Attachements) != 0 {
			_, createAttachementErr := s.ent.AttachmentRef.MapCreateBulk(freelancer.Attachements, func(c *ent.AttachmentRefCreate, j int) {
				c.SetName(freelancer.Attachements[j].Name).
					SetLink(freelancer.Attachements[j].Link).
					SetFreelancerID(freelancer.Url)
			}).Save(r.Context())

			if createAttachementErr != nil {
				return createAttachementErr
			}
		}
		if len(freelancer.Work_history) != 0 {
			_, createWorkHistoryErr := s.ent.WorkHistory.MapCreateBulk(freelancer.Work_history, func(c *ent.WorkHistoryCreate, j int) {
				work_history := freelancer.Work_history[j]
				work_history_start_date, _ := time.Parse(time.RFC3339, work_history.Start_Date)
				work_history_end_date, _ := time.Parse(time.RFC3339, work_history.End_Date)
				c.SetTitle(work_history.Title).
					SetStartDate(work_history_start_date).
					SetEndDate(work_history_end_date).
					SetDescription(work_history.Description).
					SetClientFeedback(work_history.Client_Feedback).
					SetOverallRating(work_history.Client_Rating).
					SetClientTotalSpend(work_history.Client_Total_Spend).
					SetClientTotalHires(work_history.Client_Total_Hires).
					SetClientActiveHires(work_history.Client_Active_Hires).
					SetBudget(work_history.Budget).
					SetFreelancerEarnings(work_history.Total_Earned).
					SetClientCountry(work_history.Client_Location).
					SetFreelancerID(freelancer.Url)
			}).Save(r.Context())
			if createWorkHistoryErr != nil {
				return createWorkHistoryErr
			}
		}
	}

	//delete freelancers
	for _, freelancer_id := range freelancers_to_delete {
		deleteFreelancerErr := s.ent.UpworkFreelancer.DeleteOneID(freelancer_id).Exec(r.Context())
		if deleteFreelancerErr != nil {
			return deleteFreelancerErr
		}
	}

	return writeJSON(w, http.StatusCreated, nil)
}
