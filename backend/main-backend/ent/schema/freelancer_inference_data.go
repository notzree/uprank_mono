package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
)

type FreelancerInferenceData struct {
	ent.Schema
}

func (FreelancerInferenceData) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("uprank_reccomended").Default(false).Optional(),
		field.String("uprank_reccomended_reasons").Optional(),
		field.Bool("uprank_not_enough_data").Default(false).Optional(),
		field.Float("finalized_rating_score"),
		field.Float("raw_rating_score").Optional(),
		field.Other("ai_estimated_duration", &pgtype.Interval{}).SchemaType(map[string]string{
			dialect.Postgres: "INTERVAL",
		}).Optional(),
		field.Float("budget_adherence_percentage").Optional(),
		field.Float("budget_overrun_percentage").Optional(),
	}
}

func (FreelancerInferenceData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("upworkfreelancer", UpworkFreelancer.Type).
			Ref("freelancer_inference_data").Unique().Required(),
	}
}
