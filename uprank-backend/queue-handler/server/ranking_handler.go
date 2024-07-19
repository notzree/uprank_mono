package server

import (
	"context"

	"github.com/notzree/uprank_mono/uprank-backend/queue-handler/types"
)

func (s *Server) HandleRankingRequest(ctx context.Context, req types.UpworkRankingMessage) error {
	fetched_job_data, freelancer_ranking_data, err := s.svc.FetchJobData(ctx, req)
	if err != nil {
		return NewServiceError(err, "s.svc.FetchJobData")
	}

	upsert_resp, upsert_err := s.svc.UpsertVectors(ctx, *fetched_job_data, req.User_id)
	if upsert_err != nil {
		return NewServiceError(upsert_err, "s.svc.UpsertVectors")
	}
	raw_specialization_scores, compute_raw_specialization_err := s.svc.ComputeRawSpecializationScore(ctx, types.ComputeRawSpecializationScoreRequest{
		Job_data:               *fetched_job_data,
		Job_description_vector: upsert_resp.Job_description_vector,
	})
	if compute_raw_specialization_err != nil {
		return NewServiceError(compute_raw_specialization_err, "s.svc.ComputeRawSpecializationScore")
	}
	save_raw_score_err := s.svc.SaveRawSpecializationScoreWeights(ctx, raw_specialization_scores, &freelancer_ranking_data)
	if save_raw_score_err != nil {
		return NewServiceError(save_raw_score_err, "s.svc.SaveRawSpecializationScoreWeights")
	}
	final_specialization_scores, apply_specialization_weights_err := s.svc.ApplySpecializationScoreWeights(types.ApplySpecializationScoreWeightsRequest{
		Description_scores: *raw_specialization_scores.Job_description_specialization_scores,
		Job_data:           *fetched_job_data,
	}, context.TODO())
	if apply_specialization_weights_err != nil {
		return NewServiceError(apply_specialization_weights_err, "s.svc.ApplySpecializationScoreWeights")
	}
	save_weighted_score_err := s.svc.SaveWeightedSpecializationScoreWeights(ctx, final_specialization_scores, &freelancer_ranking_data)
	if save_weighted_score_err != nil {
		return NewServiceError(save_weighted_score_err, "s.svc.SaveWeightedSpecializationScoreWeights")
	}
	budget_err := s.svc.ApplyBudgetScores(ctx, *fetched_job_data, &freelancer_ranking_data)
	if budget_err != nil {
		return NewServiceError(budget_err, "s.svc.ApplyBudgetScores")
	}

	svc_err := s.svc.PostJobRankingData(types.PostJobRankingDataRequest{
		Freelancer_ranking_data: freelancer_ranking_data,
		Job_id:                  req.Job_id.String(),
		Platform:                req.Platform,
		Platform_id:             req.Platform_id,
		User_id:                 req.User_id,
	}, ctx)
	if svc_err != nil {
		return NewServiceError(svc_err, "s.svc.PostJobRankingData")
	}

	err = s.queue.DeleteMessage(context.TODO(), req.Receipt_handle)
	if err != nil {
		return NewQError(err)
	}
	return nil
}
