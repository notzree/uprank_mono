package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqs_types "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/main-backend/ent"
	"github.com/notzree/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank-backend/main-backend/ent/schema"
	"github.com/notzree/uprank-backend/main-backend/ent/upworkfreelancer"
	"github.com/notzree/uprank-backend/main-backend/ent/upworkjob"
	"github.com/notzree/uprank-backend/main-backend/ent/user"
	"github.com/notzree/uprank-backend/main-backend/types"
)

type V1Servicer struct {
	ent               *ent.Client
	sqs_client        *sqs.Client
	ranking_queue_url string
}

func NewV1Servicer(ent *ent.Client, sqs_client *sqs.Client, ranking_queue_url string) *V1Servicer {
	return &V1Servicer{
		ent:               ent,
		sqs_client:        sqs_client,
		ranking_queue_url: ranking_queue_url,
	}
}

func (s *V1Servicer) CreateUser(data types.CreateUserRequest, ctx context.Context) (*ent.User, error) {
	new_user, err := s.ent.User.Create().
		SetID(data.User.ID).
		SetFirstName(data.User.FirstName).
		SetCompanyName(data.User.CompanyName).
		SetEmail(data.User.Email).
		Save(ctx)
	return new_user, err
}

func (s *V1Servicer) UpdateUser(data types.UpdateUserRequest, ctx context.Context) (*ent.User, error) {
	updated_user, err := s.ent.User.
		UpdateOneID(data.ClerkUserData.ID).
		SetFirstName(data.ClerkUserData.FirstName).
		SetLastLogin(time.Unix(data.ClerkUserData.LastSignInAt, 0)). //Convert timestamp into time.Time
		SetUpdatedAt(time.Unix(data.ClerkUserData.UpdatedAt, 0)).
		Save(ctx)
	return updated_user, err
}

// Creates a job and for each non-null platform job request, creates a platform job and attaches it to the job.
// TODO: ADD VALIDATION OT THE CREATE JOB REQUEST
// THERE SHOULD ONLY BE ONE PLATFORM JOB REQUEST PER JOB CREATION REQUEST
func (s *V1Servicer) CreateJob(data types.CreateJobRequest, user_id string, ctx context.Context) (*ent.Job, error) {
	new_job, create_job_err := s.ent.Job.Create().SetUserID(user_id).SetOriginPlatform(schema.Platform("uprank")).Save(ctx) //TODO: Make the origin platform dynamic
	if data.UpworkJobRequest != nil {
		if create_job_err != nil {
			return nil, create_job_err
		}
		_, create_upwork_job_err := s.AttachUpworkJob(*data.UpworkJobRequest, user_id, &new_job.ID, ctx)
		if create_upwork_job_err != nil {
			return nil, create_upwork_job_err
		}
		return new_job, nil

	}
	//technically should never happen because the type validation should catch this
	return nil, errors.New("missing platform job request")
}

// Creates an upwork job and attaches it to a job.
// Returns the created upwork job and an err
// If a job_id is provided, this is being called from the job creation api and the data (from the client) does not contain a job id.
func (s *V1Servicer) AttachUpworkJob(data types.AttachUpworkJobRequest, user_id string, job_id *uuid.UUID, ctx context.Context) (*ent.UpworkJob, error) {
	var used_job_id uuid.UUID
	if job_id == nil {
		used_job_id = data.Job_Id
	} else {
		used_job_id = *job_id
	}
	new_upwork_job, create_upwork_job_err := s.ent.UpworkJob.Create().
		SetID(data.Id).
		SetTitle(data.Title).
		SetLocation(data.Location).
		SetDescription(data.Description).
		SetSkills(data.Skills).
		SetExperienceLevel(data.Experience_level).
		SetHourly(data.Hourly).
		SetFixed(data.Fixed).
		SetHourlyRate(data.Hourly_rate).
		SetFixedRate(data.Fixed_rate).
		SetJobID(used_job_id).
		AddUserIDs(user_id).
		Save(ctx)
	if create_upwork_job_err != nil {
		return nil, create_upwork_job_err
	}
	return new_upwork_job, create_upwork_job_err
}
func (s *V1Servicer) GetJob(ctx context.Context) (*ent.Job, error) {
	//todo: implement this
	job, err := s.ent.Job.Query().First(ctx)
	return job, err
}

func (s *V1Servicer) GetUpworkJob(upwork_job_id string, user_id string, ctx context.Context) (*ent.UpworkJob, error) {
	job, err := s.ent.UpworkJob.Query().
		Where(upworkjob.IDEQ(upwork_job_id)).
		Where(
			upworkjob.HasJobWith(
				job.HasUserWith(user.IDEQ(user_id)),
			),
		).
		WithUpworkfreelancer().
		Only(ctx)
	return job, err
}

func (s *V1Servicer) GetFreelancersFromUpworkJob(upwork_job_id string, user_id string, ctx context.Context) ([]*ent.UpworkFreelancer, error) {
	freelancers, err := s.ent.UpworkJob.Query().
		Where(upworkjob.IDEQ(upwork_job_id)).
		Where(upworkjob.HasJobWith(job.HasUserWith(user.IDEQ(user_id)))).
		QueryUpworkfreelancer().All(ctx)
	return freelancers, err
}

func (s *V1Servicer) CreateUpworkFreelancer(data []types.CreateUpworkFreelancerRequest, user_id string, upwork_job_id string, ctx context.Context) ([]*ent.UpworkFreelancer, error) {
	is_found, getJobErr := s.ent.UpworkJob.Query().
		Where(upworkjob.IDEQ(upwork_job_id)).
		QueryJob().QueryUser().Where(user.IDEQ(user_id)).Exist(ctx)
	if getJobErr != nil {
		return nil, getJobErr
	}
	if !is_found {
		return nil, errors.New("resource_mismatch")
	}
	freelancers, createFreelancersErr := s.ent.UpworkFreelancer.MapCreateBulk(data, func(c *ent.UpworkFreelancerCreate, i int) {
		freelancer := data[i]
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
			SetTotalRevenue(freelancer.Earnings_info.Total_revenue).AddUpworkJobIDs(upwork_job_id)
	}).Save(ctx)
	if createFreelancersErr != nil {
		return nil, createFreelancersErr
	}
	for _, freelancer := range data {

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

func (s *V1Servicer) UpdateUpworkFreelancer(data []types.CreateUpworkFreelancerRequest, user_id string, upwork_job_id string, ctx context.Context) (created_freelancer_count int, updated_freelancer_count int, deleted_freelancer_count int, err error) {
	var (
		freelancers_to_create []types.CreateUpworkFreelancerRequest
		freelancers_to_update []types.CreateUpworkFreelancerRequest
		freelancers_to_delete []string
	)
	current_freelancers, getJobErr := s.ent.UpworkJob.Query().
		Where(
			upworkjob.IDEQ(upwork_job_id),
			upworkjob.HasJobWith(job.HasUserWith(user.IDEQ(user_id))),
		).QueryUpworkfreelancer().All(ctx)

	if getJobErr != nil {
		return 0, 0, 0, getJobErr
	}

	incoming_freelancer_dict := make(map[string]types.CreateUpworkFreelancerRequest)
	for _, freelancer := range data {
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
	for _, freelancer := range data {
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
			AddUpworkJobIDs(upwork_job_id)

	}).Save(ctx)

	if createFreelancerErr != nil {
		return 0, 0, 0, createFreelancerErr
	}

	for _, freelancer := range freelancers_to_create {
		if len(freelancer.Attachements) != 0 {
			_, createAttachementErr := s.ent.AttachmentRef.MapCreateBulk(freelancer.Attachements, func(c *ent.AttachmentRefCreate, j int) {
				c.SetName(freelancer.Attachements[j].Name).
					SetLink(freelancer.Attachements[j].Link).
					SetFreelancerID(freelancer.Url)
			}).Save(ctx)

			if createAttachementErr != nil {
				return 0, 0, 0, createAttachementErr
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
				return 0, 0, 0, createWorkHistoryErr
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
			SetTotalRevenue(freelancer.Earnings_info.Total_revenue).AddUpworkJobIDs(upwork_job_id)
	}).OnConflictColumns(upworkfreelancer.FieldID).UpdateNewValues().Exec(ctx)
	if updateFreelancersErr != nil {
		return 0, 0, 0, updateFreelancersErr
	}
	for _, freelancer := range freelancers_to_update {
		if len(freelancer.Attachements) != 0 {
			_, createAttachementErr := s.ent.AttachmentRef.MapCreateBulk(freelancer.Attachements, func(c *ent.AttachmentRefCreate, j int) {
				c.SetName(freelancer.Attachements[j].Name).
					SetLink(freelancer.Attachements[j].Link).
					SetFreelancerID(freelancer.Url)
			}).Save(ctx)

			if createAttachementErr != nil {
				return 0, 0, 0, createAttachementErr
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
				return 0, 0, 0, createWorkHistoryErr
			}
		}
	}

	//delete freelancers
	for _, freelancer_id := range freelancers_to_delete {
		deleteFreelancerErr := s.ent.UpworkFreelancer.DeleteOneID(freelancer_id).Exec(ctx)
		if deleteFreelancerErr != nil {
			return 0, 0, 0, deleteFreelancerErr
		}
	}
	return len(freelancers_to_create), len(freelancers_to_update), len(freelancers_to_delete), nil
}

// So Basically, the fifo queue uses the messagebody as the deduplication id so if we send a message with the same body within a span of 5 mins it will be discarded.
func (s *V1Servicer) SendRankingrequest(data types.RankJobRequest, ctx context.Context) error {
	_, err := s.sqs_client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageAttributes: map[string]sqs_types.MessageAttributeValue{
			"job_id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.Job_id.String()),
			},
			"user_id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.User_id),
			},
			"short_lived_token": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.Short_lived_token),
			},
			"platform": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.Platform),
			},
			"platform_id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.Platform_id),
			},
		},
		QueueUrl:       &s.ranking_queue_url,
		MessageBody:    aws.String(fmt.Sprint("Ranking request for job ", data.Job_id, " by user ", data.User_id, time.Now())),
		MessageGroupId: aws.String("RankingRequest"),
	})
	if err != nil {
		return err
	}

	return nil
}
