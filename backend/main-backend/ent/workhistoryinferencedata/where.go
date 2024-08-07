// Code generated by ent, DO NOT EDIT.

package workhistoryinferencedata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldLTE(FieldID, id))
}

// FinalizedJobRatingScore applies equality check predicate on the "finalized_job_rating_score" field. It's identical to FinalizedJobRatingScoreEQ.
func FinalizedJobRatingScore(v float64) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldEQ(FieldFinalizedJobRatingScore, v))
}

// IsWithinBudget applies equality check predicate on the "is_within_budget" field. It's identical to IsWithinBudgetEQ.
func IsWithinBudget(v bool) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldEQ(FieldIsWithinBudget, v))
}

// FinalizedJobRatingScoreEQ applies the EQ predicate on the "finalized_job_rating_score" field.
func FinalizedJobRatingScoreEQ(v float64) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldEQ(FieldFinalizedJobRatingScore, v))
}

// FinalizedJobRatingScoreNEQ applies the NEQ predicate on the "finalized_job_rating_score" field.
func FinalizedJobRatingScoreNEQ(v float64) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldNEQ(FieldFinalizedJobRatingScore, v))
}

// FinalizedJobRatingScoreIn applies the In predicate on the "finalized_job_rating_score" field.
func FinalizedJobRatingScoreIn(vs ...float64) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldIn(FieldFinalizedJobRatingScore, vs...))
}

// FinalizedJobRatingScoreNotIn applies the NotIn predicate on the "finalized_job_rating_score" field.
func FinalizedJobRatingScoreNotIn(vs ...float64) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldNotIn(FieldFinalizedJobRatingScore, vs...))
}

// FinalizedJobRatingScoreGT applies the GT predicate on the "finalized_job_rating_score" field.
func FinalizedJobRatingScoreGT(v float64) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldGT(FieldFinalizedJobRatingScore, v))
}

// FinalizedJobRatingScoreGTE applies the GTE predicate on the "finalized_job_rating_score" field.
func FinalizedJobRatingScoreGTE(v float64) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldGTE(FieldFinalizedJobRatingScore, v))
}

// FinalizedJobRatingScoreLT applies the LT predicate on the "finalized_job_rating_score" field.
func FinalizedJobRatingScoreLT(v float64) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldLT(FieldFinalizedJobRatingScore, v))
}

// FinalizedJobRatingScoreLTE applies the LTE predicate on the "finalized_job_rating_score" field.
func FinalizedJobRatingScoreLTE(v float64) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldLTE(FieldFinalizedJobRatingScore, v))
}

// IsWithinBudgetEQ applies the EQ predicate on the "is_within_budget" field.
func IsWithinBudgetEQ(v bool) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldEQ(FieldIsWithinBudget, v))
}

// IsWithinBudgetNEQ applies the NEQ predicate on the "is_within_budget" field.
func IsWithinBudgetNEQ(v bool) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.FieldNEQ(FieldIsWithinBudget, v))
}

// HasWorkHistories applies the HasEdge predicate on the "work_histories" edge.
func HasWorkHistories() predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkHistoriesTable, WorkHistoriesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkHistoriesWith applies the HasEdge predicate on the "work_histories" edge with a given conditions (other predicates).
func HasWorkHistoriesWith(preds ...predicate.WorkHistory) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(func(s *sql.Selector) {
		step := newWorkHistoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkhistoryInferenceData) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkhistoryInferenceData) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WorkhistoryInferenceData) predicate.WorkhistoryInferenceData {
	return predicate.WorkhistoryInferenceData(sql.NotPredicates(p))
}
