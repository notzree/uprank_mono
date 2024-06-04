package schema

import (
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
		field.String("title"),
		field.String("client_feedback"),
		field.Float("overall_rating").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Int("fixed_charge_amount").Optional(),
		field.String("fixed_charge_currency").Optional(),
		field.Int("hourly_charge_amount").Optional(),
		field.String("hourly_charge_currency").Optional(),
		field.Time("start_date"),
		field.Time("end_date").Optional(),
		field.String("description"),
		field.Int("total_proposals"),
		field.Int("number_of_interviews"),
		field.Strings("skills"),
		field.Float("client_rating").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Int("client_review_count"),
		field.String("client_country"),
		field.Int("client_total_jobs_posted"),
		field.Float("client_total_spend").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Int("client_total_hires").Optional(),
		field.Int("client_total_paid_hours").Optional(),
		field.Float("client_average_hourly_rate_paid").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}).Optional(),
		field.String("client_company_category").Optional(),
		field.String("client_company_size").Optional(),
	}
}

// Edges of the WorkHistory.
func (WorkHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("upwork_Freelancer_Proposal", Freelancer.Type).
			Ref("work_histories").
			Unique().Required(),
	}
}
