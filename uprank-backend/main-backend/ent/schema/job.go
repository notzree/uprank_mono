package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Job struct {
	ent.Schema
}

func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().
			Unique().
			Immutable(),
		field.String("title").NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.String("location").Optional(),
		field.String("description").NotEmpty(),
		field.JSON("skills", []string{}).Optional(),
		field.String("experience_level").Optional(),
		field.Bool("hourly"),
		field.Bool("fixed"),
		field.JSON("hourly_rate", []float32{}).Optional(), // hourly rate can be an array of floats with length 1 or 2 (for a range of hourly rates)
		field.Float("fixed_rate").
			SchemaType(map[string]string{
				dialect.Postgres: "numeric",
			}).Optional(),
		field.Float("average_uprank_score").Optional(),
		field.Float("max_uprank_score").Optional(),
		field.Float("min_uprank_score").Optional(),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("jobs").
			Unique().Required(),
		edge.To("freelancers", Freelancer.Type),
	}
}
