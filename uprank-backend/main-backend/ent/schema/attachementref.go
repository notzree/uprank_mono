package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AttachmentRef holds the schema definition for the AttachmentRef entity.
type AttachmentRef struct {
	ent.Schema
}

// Fields of the AttachmentRef.
func (AttachmentRef) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("link"),
	}
}

// Edges of the AttachmentRef.
func (AttachmentRef) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("freelancer", UpworkFreelancer.Type).
			Ref("attachments").
			Unique().Required(),
	}
}
