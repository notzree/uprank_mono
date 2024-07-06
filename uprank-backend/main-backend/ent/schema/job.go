package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Job struct {
	ent.Schema
}

func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.Enum("origin_platform").GoType(Platform("")).Immutable(), //Enum for the starting platform type the job was created. Can be created either through Uprank or through extension (fiverr, upwork, etc)
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("job").
			Unique().Required(),
		edge.To("upworkjob", UpworkJob.Type).Unique(),
	}
}

// Enum for platform type
type Platform string

const (
	Upwork Platform = "upwork"
	Uprank Platform = "uprank"
)

func (Platform) Values() (kinds []string) {
	for _, s := range []Platform{Upwork, Uprank} {
		kinds = append(kinds, string(s))
	}
	return
}
