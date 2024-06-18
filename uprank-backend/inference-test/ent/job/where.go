// Code generated by ent, DO NOT EDIT.

package job

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/inference-backend/ent/predicate"
	"github.com/notzree/uprank-backend/inference-backend/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldID, id))
}

// OriginPlatformEQ applies the EQ predicate on the "origin_platform" field.
func OriginPlatformEQ(v schema.Platform) predicate.Job {
	vc := v
	return predicate.Job(sql.FieldEQ(FieldOriginPlatform, vc))
}

// OriginPlatformNEQ applies the NEQ predicate on the "origin_platform" field.
func OriginPlatformNEQ(v schema.Platform) predicate.Job {
	vc := v
	return predicate.Job(sql.FieldNEQ(FieldOriginPlatform, vc))
}

// OriginPlatformIn applies the In predicate on the "origin_platform" field.
func OriginPlatformIn(vs ...schema.Platform) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(sql.FieldIn(FieldOriginPlatform, v...))
}

// OriginPlatformNotIn applies the NotIn predicate on the "origin_platform" field.
func OriginPlatformNotIn(vs ...schema.Platform) predicate.Job {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(sql.FieldNotIn(FieldOriginPlatform, v...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpworkjob applies the HasEdge predicate on the "upworkjob" edge.
func HasUpworkjob() predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UpworkjobTable, UpworkjobColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpworkjobWith applies the HasEdge predicate on the "upworkjob" edge with a given conditions (other predicates).
func HasUpworkjobWith(preds ...predicate.UpworkJob) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := newUpworkjobStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Job) predicate.Job {
	return predicate.Job(sql.NotPredicates(p))
}