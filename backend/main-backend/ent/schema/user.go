package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"). //from clerk: user.data.id
					NotEmpty().
					Unique().
					Immutable(),
		field.String("first_name").Default("unknown").NotEmpty(),
		field.String("company_name").Default("unknown").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("last_login").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("job", Job.Type),
		edge.To("upworkjob", UpworkJob.Type),
	}
}
