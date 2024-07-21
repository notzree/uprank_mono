package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/upworkfreelancer"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/upworkjob"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/user"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/types"
)

type UpworkFreelancerService struct {
	ent *ent.Client
}

type NewUpworkFreelancerServiceParams struct {
	Ent *ent.Client
}

func NewUpworkFreelancerService(params NewUpworkFreelancerServiceParams) UpworkFreelancerService {
	return UpworkFreelancerService{
		ent: params.Ent,
	}
}

func (s *UpworkFreelancerService) CreateUpworkFreelancer(ctx context.Context, params types.CreateUpworkFreelancerRequest) ([]*ent.UpworkFreelancer, error) {
	is_found, getJobErr := s.ent.UpworkJob.Query().
		Where(upworkjob.IDEQ(params.UpworkJobId)).
		QueryJob().QueryUser().Where(user.IDEQ(params.UserId)).Exist(ctx)
	if getJobErr != nil {
		return nil, getJobErr
	}
	if !is_found {
		return nil, errors.New("resource_mismatch")
	}
	freelancers, createFreelancersErr := s.ent.UpworkFreelancer.MapCreateBulk(params.Data, func(c *ent.UpworkFreelancerCreate, i int) {
		freelancer := params.Data[i]
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
			SetTotalRevenue(freelancer.Earnings_info.Total_revenue).AddUpworkJobIDs(params.UpworkJobId)
	}).Save(ctx)
	if createFreelancersErr != nil {
		return nil, createFreelancersErr
	}
	for _, freelancer := range params.Data {

		if len(freelancer.Attachements) != 0 {
			_, createAttachementErr := s.ent.AttachmentRef.MapCreateBulk(freelancer.Attachements, func(c *ent.AttachmentRefCreate, j int) {
				c.SetName(freelancer.Attachements[j].Name).
					SetLink(freelancer.Attachements[j].Link).
					SetFreelancerID(freelancer.Url)
			}).Save(ctx)
			if createAttachementErr != nil {
				return nil, createAttachementErr
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
			}).Save(ctx)
			if createWorkHistoryErr != nil {
				return nil, createWorkHistoryErr
			}
		}
	}
	return freelancers, nil
}

func (s *UpworkFreelancerService) UpsertUpworkFreelancer(ctx context.Context, params types.CreateUpworkFreelancerRequest) (resp *types.CreateUpworkFreelancerResponse, err error) {
	var (
		freelancers_to_create []types.CreateUpworkFreelancerData
		freelancers_to_update []types.CreateUpworkFreelancerData
		freelancers_to_delete []string
	)
	current_freelancers, getJobErr := s.ent.UpworkJob.Query().
		Where(
			upworkjob.IDEQ(params.UpworkJobId),
			upworkjob.HasJobWith(job.HasUserWith(user.IDEQ(params.UserId))),
		).QueryUpworkfreelancer().All(ctx)

	if getJobErr != nil {
		return nil, getJobErr
	}

	incoming_freelancer_dict := make(map[string]types.CreateUpworkFreelancerData)
	for _, freelancer := range params.Data {
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
	for _, freelancer := range params.Data {
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
			AddUpworkJobIDs(params.UpworkJobId)

	}).Save(ctx)

	if createFreelancerErr != nil {
		return nil, createFreelancerErr
	}

	for _, freelancer := range freelancers_to_create {
		if len(freelancer.Attachements) != 0 {
			_, createAttachementErr := s.ent.AttachmentRef.MapCreateBulk(freelancer.Attachements, func(c *ent.AttachmentRefCreate, j int) {
				c.SetName(freelancer.Attachements[j].Name).
					SetLink(freelancer.Attachements[j].Link).
					SetFreelancerID(freelancer.Url)
			}).Save(ctx)

			if createAttachementErr != nil {
				return nil, createAttachementErr
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
			}).Save(ctx)
			if createWorkHistoryErr != nil {
				return nil, createWorkHistoryErr
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
			SetTotalRevenue(freelancer.Earnings_info.Total_revenue).AddUpworkJobIDs(params.UpworkJobId)
	}).OnConflictColumns(upworkfreelancer.FieldID).UpdateNewValues().Exec(ctx)
	if updateFreelancersErr != nil {
		return nil, updateFreelancersErr
	}
	for _, freelancer := range freelancers_to_update {
		if len(freelancer.Attachements) != 0 {
			_, createAttachementErr := s.ent.AttachmentRef.MapCreateBulk(freelancer.Attachements, func(c *ent.AttachmentRefCreate, j int) {
				c.SetName(freelancer.Attachements[j].Name).
					SetLink(freelancer.Attachements[j].Link).
					SetFreelancerID(freelancer.Url)
			}).Save(ctx)

			if createAttachementErr != nil {
				return nil, createAttachementErr
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
			}).Save(ctx)
			if createWorkHistoryErr != nil {
				return nil, createWorkHistoryErr
			}
		}
	}
	//delete freelancers
	for _, freelancer_id := range freelancers_to_delete {
		deleteFreelancerErr := s.ent.UpworkFreelancer.DeleteOneID(freelancer_id).Exec(ctx)
		if deleteFreelancerErr != nil {
			return nil, deleteFreelancerErr
		}
	}
	return &types.CreateUpworkFreelancerResponse{
		CreatedFreelancerCount: len(freelancers_to_create),
		UpdatedFreelancerCount: len(freelancers_to_update),
		DeletedFreelancerCount: len(freelancers_to_delete),
	}, nil
}

func (s *UpworkFreelancerService) GetAllFreelancersFromUpworkJob(ctx context.Context, params types.QueryUpworkJobRequest) ([]*ent.UpworkFreelancer, error) {
	freelancers, err := s.ent.UpworkJob.Query().
		Where(upworkjob.IDEQ(params.Upwork_job_id)).
		Where(upworkjob.HasJobWith(job.HasUserWith(user.IDEQ(params.User_id)))).
		QueryUpworkfreelancer().All(ctx)
	return freelancers, err
}

// data []types.UpdateUpworkFreelancerRequest, user_id, upwork_job_id string,
func (s *UpworkFreelancerService) UpdateUpworkFreelancer(ctx context.Context, params types.UpdateUpworkFreelancerRequest) ([]string, error) {
	updated_ids := make([]string, 0)
	for _, freelancer := range params.Data {
		updated_freelancer := s.ent.UpworkFreelancer.UpdateOneID(freelancer.Url)
		if freelancer.Name != nil {
			updated_freelancer.SetName(*freelancer.Name)
		}
		if freelancer.Title != nil {
			updated_freelancer.SetTitle(*freelancer.Title)
		}
		if freelancer.Description != nil {
			updated_freelancer.SetDescription(*freelancer.Description)
		}
		if freelancer.Location != nil {
			updated_freelancer.SetCity(freelancer.Location.City)
			updated_freelancer.SetCountry(freelancer.Location.Country)
			updated_freelancer.SetTimezone(freelancer.Location.Timezone)
		}
		if freelancer.Cv != nil {
			updated_freelancer.SetCv(*freelancer.Cv)
		}
		if freelancer.Ai_reccomended != nil {
			updated_freelancer.SetAiReccomended(*freelancer.Ai_reccomended)
		}
		if freelancer.Fixed_charge_amount != nil {
			parsed_fixed_charge_amount, _ := strconv.ParseFloat(*freelancer.Fixed_charge_amount, 64)
			updated_freelancer.SetFixedChargeAmount(parsed_fixed_charge_amount)
		}
		if freelancer.Fixed_charge_currency != nil {
			updated_freelancer.SetFixedChargeCurrency(*freelancer.Fixed_charge_currency)
		}
		if freelancer.Hourly_charge_amount != nil {
			parse_hourly_charge_amount, _ := strconv.ParseFloat(*freelancer.Hourly_charge_amount, 64)
			updated_freelancer.SetHourlyChargeAmount(parse_hourly_charge_amount)
		}
		if freelancer.Hourly_charge_currency != nil {
			updated_freelancer.SetHourlyChargeCurrency(*freelancer.Hourly_charge_currency)
		}
		if freelancer.Invited != nil {
			updated_freelancer.SetInvited(*freelancer.Invited)
		}
		if freelancer.Photo_url != nil {
			updated_freelancer.SetPhotoURL(*freelancer.Photo_url)
		}
		if freelancer.Recent_hours != nil {
			updated_freelancer.SetRecentHours(*freelancer.Recent_hours)
		}
		if freelancer.Total_hours != nil {
			updated_freelancer.SetTotalHours(*freelancer.Total_hours)
		}
		if freelancer.Total_portfolio_items != nil {
			updated_freelancer.SetTotalPortfolioItems(*freelancer.Total_portfolio_items)
		}
		if freelancer.Total_portfolio_v2_items != nil {
			updated_freelancer.SetTotalPortfolioV2Items(*freelancer.Total_portfolio_v2_items)
		}
		if freelancer.Total_feedback != nil {
			updated_freelancer.SetUpworkTotalFeedback(*freelancer.Total_feedback)
		}
		if freelancer.Recent_feedback != nil {
			updated_freelancer.SetUpworkRecentFeedback(*freelancer.Recent_feedback)
		}
		if freelancer.Top_rated_status != nil {
			updated_freelancer.SetUpworkTopRatedStatus(*freelancer.Top_rated_status)
		}
		if freelancer.Top_rated_plus_status != nil {
			updated_freelancer.SetUpworkTopRatedPlusStatus(*freelancer.Top_rated_plus_status)
		}
		if freelancer.Sponsored != nil {
			updated_freelancer.SetUpworkSponsored(*freelancer.Sponsored)
		}
		if freelancer.Job_success_score != nil {
			updated_freelancer.SetUpworkJobSuccessScore(*freelancer.Job_success_score)
		}
		if freelancer.Reccomended != nil {
			updated_freelancer.SetUpworkReccomended(*freelancer.Reccomended)
		}
		if freelancer.Skills != nil {
			updated_freelancer.SetSkills(*freelancer.Skills)
		}
		if freelancer.Earnings_info != nil {
			updated_freelancer.SetAverageRecentEarnings(freelancer.Earnings_info.Average_recent_earnings)

			updated_freelancer.SetCombinedAverageRecentEarnings(freelancer.Earnings_info.Combined_average_recent_earnings)

			updated_freelancer.SetCombinedRecentEarnings(freelancer.Earnings_info.Combined_recent_earnings)

			updated_freelancer.SetCombinedTotalEarnings(freelancer.Earnings_info.Combined_total_earnings)

			updated_freelancer.SetCombinedTotalRevenue(freelancer.Earnings_info.Combined_total_revenue)

			updated_freelancer.SetRecentEarnings(freelancer.Earnings_info.Recent_earnings)

			updated_freelancer.SetTotalRevenue(freelancer.Earnings_info.Total_revenue)

		}
		if freelancer.Embedded_at != nil {
			updated_freelancer.SetEmbeddedAt(*freelancer.Embedded_at)
		}
		//todo impplement updated attachements
		_, err := updated_freelancer.Save(ctx)
		if err != nil {
			return nil, err
		}
		updated_ids = append(updated_ids, freelancer.Url)
	}
	return updated_ids, nil
}
