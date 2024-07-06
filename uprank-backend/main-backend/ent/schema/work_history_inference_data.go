package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type WorkhistoryInferenceData struct {
	ent.Schema
}

func (WorkhistoryInferenceData) Fields() []ent.Field {
	return []ent.Field{
		field.Float("finalized_job_rating_score"),
		field.Bool("is_within_budget"),
	}
}

func (WorkhistoryInferenceData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("work_histories", WorkHistory.Type).
			Ref("work_history_inference_data").Unique().Required(),
	}
}
