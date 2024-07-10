// Code generated by ent, DO NOT EDIT.

package freelancerinferencedata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/jackc/pgtype"
	"github.com/notzree/uprank-backend/main-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLTE(FieldID, id))
}

// UprankReccomended applies equality check predicate on the "uprank_reccomended" field. It's identical to UprankReccomendedEQ.
func UprankReccomended(v bool) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldUprankReccomended, v))
}

// UprankReccomendedReasons applies equality check predicate on the "uprank_reccomended_reasons" field. It's identical to UprankReccomendedReasonsEQ.
func UprankReccomendedReasons(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldUprankReccomendedReasons, v))
}

// UprankNotEnoughData applies equality check predicate on the "uprank_not_enough_data" field. It's identical to UprankNotEnoughDataEQ.
func UprankNotEnoughData(v bool) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldUprankNotEnoughData, v))
}

// FinalizedRatingScore applies equality check predicate on the "finalized_rating_score" field. It's identical to FinalizedRatingScoreEQ.
func FinalizedRatingScore(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldFinalizedRatingScore, v))
}

// RawRatingScore applies equality check predicate on the "raw_rating_score" field. It's identical to RawRatingScoreEQ.
func RawRatingScore(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldRawRatingScore, v))
}

// AiEstimatedDuration applies equality check predicate on the "ai_estimated_duration" field. It's identical to AiEstimatedDurationEQ.
func AiEstimatedDuration(v *pgtype.Interval) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldAiEstimatedDuration, v))
}

// BudgetAdherencePercentage applies equality check predicate on the "budget_adherence_percentage" field. It's identical to BudgetAdherencePercentageEQ.
func BudgetAdherencePercentage(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldBudgetAdherencePercentage, v))
}

// BudgetOverrunPercentage applies equality check predicate on the "budget_overrun_percentage" field. It's identical to BudgetOverrunPercentageEQ.
func BudgetOverrunPercentage(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldBudgetOverrunPercentage, v))
}

// UprankReccomendedEQ applies the EQ predicate on the "uprank_reccomended" field.
func UprankReccomendedEQ(v bool) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldUprankReccomended, v))
}

// UprankReccomendedNEQ applies the NEQ predicate on the "uprank_reccomended" field.
func UprankReccomendedNEQ(v bool) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNEQ(FieldUprankReccomended, v))
}

// UprankReccomendedIsNil applies the IsNil predicate on the "uprank_reccomended" field.
func UprankReccomendedIsNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIsNull(FieldUprankReccomended))
}

// UprankReccomendedNotNil applies the NotNil predicate on the "uprank_reccomended" field.
func UprankReccomendedNotNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotNull(FieldUprankReccomended))
}

// UprankReccomendedReasonsEQ applies the EQ predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsEQ(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsNEQ applies the NEQ predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsNEQ(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNEQ(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsIn applies the In predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsIn(vs ...string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIn(FieldUprankReccomendedReasons, vs...))
}

// UprankReccomendedReasonsNotIn applies the NotIn predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsNotIn(vs ...string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotIn(FieldUprankReccomendedReasons, vs...))
}

// UprankReccomendedReasonsGT applies the GT predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsGT(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGT(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsGTE applies the GTE predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsGTE(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGTE(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsLT applies the LT predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsLT(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLT(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsLTE applies the LTE predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsLTE(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLTE(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsContains applies the Contains predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsContains(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldContains(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsHasPrefix applies the HasPrefix predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsHasPrefix(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldHasPrefix(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsHasSuffix applies the HasSuffix predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsHasSuffix(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldHasSuffix(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsIsNil applies the IsNil predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsIsNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIsNull(FieldUprankReccomendedReasons))
}

// UprankReccomendedReasonsNotNil applies the NotNil predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsNotNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotNull(FieldUprankReccomendedReasons))
}

// UprankReccomendedReasonsEqualFold applies the EqualFold predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsEqualFold(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEqualFold(FieldUprankReccomendedReasons, v))
}

// UprankReccomendedReasonsContainsFold applies the ContainsFold predicate on the "uprank_reccomended_reasons" field.
func UprankReccomendedReasonsContainsFold(v string) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldContainsFold(FieldUprankReccomendedReasons, v))
}

// UprankNotEnoughDataEQ applies the EQ predicate on the "uprank_not_enough_data" field.
func UprankNotEnoughDataEQ(v bool) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldUprankNotEnoughData, v))
}

// UprankNotEnoughDataNEQ applies the NEQ predicate on the "uprank_not_enough_data" field.
func UprankNotEnoughDataNEQ(v bool) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNEQ(FieldUprankNotEnoughData, v))
}

// UprankNotEnoughDataIsNil applies the IsNil predicate on the "uprank_not_enough_data" field.
func UprankNotEnoughDataIsNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIsNull(FieldUprankNotEnoughData))
}

// UprankNotEnoughDataNotNil applies the NotNil predicate on the "uprank_not_enough_data" field.
func UprankNotEnoughDataNotNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotNull(FieldUprankNotEnoughData))
}

// FinalizedRatingScoreEQ applies the EQ predicate on the "finalized_rating_score" field.
func FinalizedRatingScoreEQ(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldFinalizedRatingScore, v))
}

// FinalizedRatingScoreNEQ applies the NEQ predicate on the "finalized_rating_score" field.
func FinalizedRatingScoreNEQ(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNEQ(FieldFinalizedRatingScore, v))
}

// FinalizedRatingScoreIn applies the In predicate on the "finalized_rating_score" field.
func FinalizedRatingScoreIn(vs ...float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIn(FieldFinalizedRatingScore, vs...))
}

// FinalizedRatingScoreNotIn applies the NotIn predicate on the "finalized_rating_score" field.
func FinalizedRatingScoreNotIn(vs ...float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotIn(FieldFinalizedRatingScore, vs...))
}

// FinalizedRatingScoreGT applies the GT predicate on the "finalized_rating_score" field.
func FinalizedRatingScoreGT(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGT(FieldFinalizedRatingScore, v))
}

// FinalizedRatingScoreGTE applies the GTE predicate on the "finalized_rating_score" field.
func FinalizedRatingScoreGTE(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGTE(FieldFinalizedRatingScore, v))
}

// FinalizedRatingScoreLT applies the LT predicate on the "finalized_rating_score" field.
func FinalizedRatingScoreLT(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLT(FieldFinalizedRatingScore, v))
}

// FinalizedRatingScoreLTE applies the LTE predicate on the "finalized_rating_score" field.
func FinalizedRatingScoreLTE(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLTE(FieldFinalizedRatingScore, v))
}

// RawRatingScoreEQ applies the EQ predicate on the "raw_rating_score" field.
func RawRatingScoreEQ(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldRawRatingScore, v))
}

// RawRatingScoreNEQ applies the NEQ predicate on the "raw_rating_score" field.
func RawRatingScoreNEQ(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNEQ(FieldRawRatingScore, v))
}

// RawRatingScoreIn applies the In predicate on the "raw_rating_score" field.
func RawRatingScoreIn(vs ...float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIn(FieldRawRatingScore, vs...))
}

// RawRatingScoreNotIn applies the NotIn predicate on the "raw_rating_score" field.
func RawRatingScoreNotIn(vs ...float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotIn(FieldRawRatingScore, vs...))
}

// RawRatingScoreGT applies the GT predicate on the "raw_rating_score" field.
func RawRatingScoreGT(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGT(FieldRawRatingScore, v))
}

// RawRatingScoreGTE applies the GTE predicate on the "raw_rating_score" field.
func RawRatingScoreGTE(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGTE(FieldRawRatingScore, v))
}

// RawRatingScoreLT applies the LT predicate on the "raw_rating_score" field.
func RawRatingScoreLT(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLT(FieldRawRatingScore, v))
}

// RawRatingScoreLTE applies the LTE predicate on the "raw_rating_score" field.
func RawRatingScoreLTE(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLTE(FieldRawRatingScore, v))
}

// RawRatingScoreIsNil applies the IsNil predicate on the "raw_rating_score" field.
func RawRatingScoreIsNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIsNull(FieldRawRatingScore))
}

// RawRatingScoreNotNil applies the NotNil predicate on the "raw_rating_score" field.
func RawRatingScoreNotNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotNull(FieldRawRatingScore))
}

// AiEstimatedDurationEQ applies the EQ predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationEQ(v *pgtype.Interval) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldAiEstimatedDuration, v))
}

// AiEstimatedDurationNEQ applies the NEQ predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationNEQ(v *pgtype.Interval) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNEQ(FieldAiEstimatedDuration, v))
}

// AiEstimatedDurationIn applies the In predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationIn(vs ...*pgtype.Interval) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIn(FieldAiEstimatedDuration, vs...))
}

// AiEstimatedDurationNotIn applies the NotIn predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationNotIn(vs ...*pgtype.Interval) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotIn(FieldAiEstimatedDuration, vs...))
}

// AiEstimatedDurationGT applies the GT predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationGT(v *pgtype.Interval) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGT(FieldAiEstimatedDuration, v))
}

// AiEstimatedDurationGTE applies the GTE predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationGTE(v *pgtype.Interval) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGTE(FieldAiEstimatedDuration, v))
}

// AiEstimatedDurationLT applies the LT predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationLT(v *pgtype.Interval) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLT(FieldAiEstimatedDuration, v))
}

// AiEstimatedDurationLTE applies the LTE predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationLTE(v *pgtype.Interval) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLTE(FieldAiEstimatedDuration, v))
}

// AiEstimatedDurationIsNil applies the IsNil predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationIsNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIsNull(FieldAiEstimatedDuration))
}

// AiEstimatedDurationNotNil applies the NotNil predicate on the "ai_estimated_duration" field.
func AiEstimatedDurationNotNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotNull(FieldAiEstimatedDuration))
}

// BudgetAdherencePercentageEQ applies the EQ predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageEQ(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldBudgetAdherencePercentage, v))
}

// BudgetAdherencePercentageNEQ applies the NEQ predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageNEQ(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNEQ(FieldBudgetAdherencePercentage, v))
}

// BudgetAdherencePercentageIn applies the In predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageIn(vs ...float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIn(FieldBudgetAdherencePercentage, vs...))
}

// BudgetAdherencePercentageNotIn applies the NotIn predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageNotIn(vs ...float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotIn(FieldBudgetAdherencePercentage, vs...))
}

// BudgetAdherencePercentageGT applies the GT predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageGT(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGT(FieldBudgetAdherencePercentage, v))
}

// BudgetAdherencePercentageGTE applies the GTE predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageGTE(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGTE(FieldBudgetAdherencePercentage, v))
}

// BudgetAdherencePercentageLT applies the LT predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageLT(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLT(FieldBudgetAdherencePercentage, v))
}

// BudgetAdherencePercentageLTE applies the LTE predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageLTE(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLTE(FieldBudgetAdherencePercentage, v))
}

// BudgetAdherencePercentageIsNil applies the IsNil predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageIsNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIsNull(FieldBudgetAdherencePercentage))
}

// BudgetAdherencePercentageNotNil applies the NotNil predicate on the "budget_adherence_percentage" field.
func BudgetAdherencePercentageNotNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotNull(FieldBudgetAdherencePercentage))
}

// BudgetOverrunPercentageEQ applies the EQ predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageEQ(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldEQ(FieldBudgetOverrunPercentage, v))
}

// BudgetOverrunPercentageNEQ applies the NEQ predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageNEQ(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNEQ(FieldBudgetOverrunPercentage, v))
}

// BudgetOverrunPercentageIn applies the In predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageIn(vs ...float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIn(FieldBudgetOverrunPercentage, vs...))
}

// BudgetOverrunPercentageNotIn applies the NotIn predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageNotIn(vs ...float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotIn(FieldBudgetOverrunPercentage, vs...))
}

// BudgetOverrunPercentageGT applies the GT predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageGT(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGT(FieldBudgetOverrunPercentage, v))
}

// BudgetOverrunPercentageGTE applies the GTE predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageGTE(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldGTE(FieldBudgetOverrunPercentage, v))
}

// BudgetOverrunPercentageLT applies the LT predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageLT(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLT(FieldBudgetOverrunPercentage, v))
}

// BudgetOverrunPercentageLTE applies the LTE predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageLTE(v float64) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldLTE(FieldBudgetOverrunPercentage, v))
}

// BudgetOverrunPercentageIsNil applies the IsNil predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageIsNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldIsNull(FieldBudgetOverrunPercentage))
}

// BudgetOverrunPercentageNotNil applies the NotNil predicate on the "budget_overrun_percentage" field.
func BudgetOverrunPercentageNotNil() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.FieldNotNull(FieldBudgetOverrunPercentage))
}

// HasUpworkfreelancer applies the HasEdge predicate on the "upworkfreelancer" edge.
func HasUpworkfreelancer() predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UpworkfreelancerTable, UpworkfreelancerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpworkfreelancerWith applies the HasEdge predicate on the "upworkfreelancer" edge with a given conditions (other predicates).
func HasUpworkfreelancerWith(preds ...predicate.UpworkFreelancer) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(func(s *sql.Selector) {
		step := newUpworkfreelancerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FreelancerInferenceData) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FreelancerInferenceData) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FreelancerInferenceData) predicate.FreelancerInferenceData {
	return predicate.FreelancerInferenceData(sql.NotPredicates(p))
}
