package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WorkHistory holds the schema definition for the WorkHistory entity.
type WorkHistory struct {
	ent.Schema
}

// Fields of the WorkHistory.
func (WorkHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Time("embedded_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("title"),
		field.String("client_feedback").Optional(),
		field.Float("overall_rating").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}).Optional(),
		field.Float("freelancer_earnings").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}).Optional(),
		field.Time("start_date").Optional(),
		field.Time("end_date").Optional(),
		field.String("description").Optional(),
		field.Float("budget").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}).Optional(),
		field.Float("client_rating").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}).Optional(),
		field.Int("client_review_count").Optional(),
		field.String("client_country").Optional(),
		field.Int("client_total_jobs_posted").Optional(),
		field.Float("client_total_spend").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}).Optional(),
		field.Int("client_total_hires").Optional(),
		field.Int("client_active_hires").Optional(),
		field.Int("client_total_paid_hours").Optional(),
		field.Float("client_average_hourly_rate_paid").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}).Optional(),
		field.String("client_company_category").Optional(),
		field.String("client_company_size").Optional(),
		field.Int("total_proposals").Optional(),
		field.Int("number_of_interviews").Optional(),
		field.Strings("skills").Optional(),
	}
}

// Edges of the WorkHistory.
func (WorkHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("freelancer", UpworkFreelancer.Type).
			Ref("work_histories").
			Unique().Required(),
		edge.To("work_history_inference_data", WorkhistoryInferenceData.Type),
	}
}
