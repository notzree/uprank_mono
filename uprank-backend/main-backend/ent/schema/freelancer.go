package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Freelancer represents a freelancers application to a job
type Freelancer struct {
	ent.Schema
}

func (Freelancer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique(),
		field.String("url").Unique(),
		field.String("name"),
		field.String("title"),
		field.String("description"),
		field.String("city"),
		field.String("country"),
		field.String("timezone"),
		field.String("cv").SchemaType(map[string]string{dialect.Postgres: "TEXT"}),
		field.Bool("ai_reccomended"),
		field.Int("fixed_charge_amount").Optional(),
		field.String("fixed_charge_currency"),
		field.Int("hourly_charge_amount").Optional(),
		field.String("hourly_charge_currency"),
		field.Bool("invited"),
		field.String("photo_url"),
		field.Int("recent_hours"),
		field.Int("total_hours"),
		field.Int("total_portfolio_items"),
		field.Int("total_portfolio_v2_items"),
		field.Float("upwork_total_feedback").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Float("upwork_recent_feedback").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Bool("upwork_top_rated_status"),
		field.Bool("upwork_top_rated_plus_status"),
		field.Bool("upwork_sponsored"),
		field.Float("upwork_job_success_score").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Bool("upwork_reccomended"),
		field.Strings("skills"),
		field.Float("average_recent_earnings").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Float("combined_average_recent_earnings").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Float("combined_recent_earnings").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Float("combined_total_earnings").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Float("combined_total_revenue").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Float("recent_earnings").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Float("total_revenue").SchemaType(map[string]string{dialect.Postgres: "DECIMAL"}),
		field.Int("uprank_score").Default(0).Optional(),
		field.Time("uprank_updated_at").Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("uprank_reccomended").Default(false).Optional(),
		field.String("uprank_reccomended_reasons").Optional(),
		field.Bool("uprank_not_enough_data").Default(false).Optional(),
	}
}

// Edges of the Freelancer.
func (Freelancer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("job", Job.Type).
			Ref("freelancers").
			Unique().Required(),
		edge.To("attachments", AttachmentRef.Type),
		edge.To("work_histories", WorkHistory.Type),
	}
}
