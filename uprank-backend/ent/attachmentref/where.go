// Code generated by ent, DO NOT EDIT.

package attachmentref

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/notzree/uprank-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldEQ(FieldName, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldEQ(FieldURL, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldContainsFold(FieldName, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.FieldContainsFold(FieldURL, v))
}

// HasFreelancer applies the HasEdge predicate on the "freelancer" edge.
func HasFreelancer() predicate.AttachmentRef {
	return predicate.AttachmentRef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FreelancerTable, FreelancerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFreelancerWith applies the HasEdge predicate on the "freelancer" edge with a given conditions (other predicates).
func HasFreelancerWith(preds ...predicate.Freelancer) predicate.AttachmentRef {
	return predicate.AttachmentRef(func(s *sql.Selector) {
		step := newFreelancerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AttachmentRef) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AttachmentRef) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AttachmentRef) predicate.AttachmentRef {
	return predicate.AttachmentRef(sql.NotPredicates(p))
}
