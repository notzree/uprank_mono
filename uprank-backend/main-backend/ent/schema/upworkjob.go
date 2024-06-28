package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UpworkJob struct {
	ent.Schema
}

func (UpworkJob) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().
			Unique().
			Immutable(),
		field.String("title").NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("embedded_at").Optional(),
		field.Time("ranked_at").Optional(),
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

// Edges of the UpworkJob.
func (UpworkJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("upworkfreelancer", UpworkFreelancer.Type),
		edge.From("job", Job.Type).Ref("upworkjob").Unique().Required(),
		edge.From("user", User.Type).Ref("upworkjob"),
	}
}
