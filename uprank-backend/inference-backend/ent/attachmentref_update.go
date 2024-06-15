// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notzree/uprank-backend/inference-backend/ent/attachmentref"
	"github.com/notzree/uprank-backend/inference-backend/ent/predicate"
	"github.com/notzree/uprank-backend/inference-backend/ent/upworkfreelancer"
)

// AttachmentRefUpdate is the builder for updating AttachmentRef entities.
type AttachmentRefUpdate struct {
	config
	hooks    []Hook
	mutation *AttachmentRefMutation
}

// Where appends a list predicates to the AttachmentRefUpdate builder.
func (aru *AttachmentRefUpdate) Where(ps ...predicate.AttachmentRef) *AttachmentRefUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetName sets the "name" field.
func (aru *AttachmentRefUpdate) SetName(s string) *AttachmentRefUpdate {
	aru.mutation.SetName(s)
	return aru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (aru *AttachmentRefUpdate) SetNillableName(s *string) *AttachmentRefUpdate {
	if s != nil {
		aru.SetName(*s)
	}
	return aru
}

// SetLink sets the "link" field.
func (aru *AttachmentRefUpdate) SetLink(s string) *AttachmentRefUpdate {
	aru.mutation.SetLink(s)
	return aru
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (aru *AttachmentRefUpdate) SetNillableLink(s *string) *AttachmentRefUpdate {
	if s != nil {
		aru.SetLink(*s)
	}
	return aru
}

// SetFreelancerID sets the "freelancer" edge to the UpworkFreelancer entity by ID.
func (aru *AttachmentRefUpdate) SetFreelancerID(id string) *AttachmentRefUpdate {
	aru.mutation.SetFreelancerID(id)
	return aru
}

// SetFreelancer sets the "freelancer" edge to the UpworkFreelancer entity.
func (aru *AttachmentRefUpdate) SetFreelancer(u *UpworkFreelancer) *AttachmentRefUpdate {
	return aru.SetFreelancerID(u.ID)
}

// Mutation returns the AttachmentRefMutation object of the builder.
func (aru *AttachmentRefUpdate) Mutation() *AttachmentRefMutation {
	return aru.mutation
}

// ClearFreelancer clears the "freelancer" edge to the UpworkFreelancer entity.
func (aru *AttachmentRefUpdate) ClearFreelancer() *AttachmentRefUpdate {
	aru.mutation.ClearFreelancer()
	return aru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AttachmentRefUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aru.sqlSave, aru.mutation, aru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AttachmentRefUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AttachmentRefUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AttachmentRefUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aru *AttachmentRefUpdate) check() error {
	if _, ok := aru.mutation.FreelancerID(); aru.mutation.FreelancerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AttachmentRef.freelancer"`)
	}
	return nil
}

func (aru *AttachmentRefUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(attachmentref.Table, attachmentref.Columns, sqlgraph.NewFieldSpec(attachmentref.FieldID, field.TypeInt))
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.Name(); ok {
		_spec.SetField(attachmentref.FieldName, field.TypeString, value)
	}
	if value, ok := aru.mutation.Link(); ok {
		_spec.SetField(attachmentref.FieldLink, field.TypeString, value)
	}
	if aru.mutation.FreelancerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachmentref.FreelancerTable,
			Columns: []string{attachmentref.FreelancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(upworkfreelancer.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.FreelancerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachmentref.FreelancerTable,
			Columns: []string{attachmentref.FreelancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(upworkfreelancer.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachmentref.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aru.mutation.done = true
	return n, nil
}

// AttachmentRefUpdateOne is the builder for updating a single AttachmentRef entity.
type AttachmentRefUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttachmentRefMutation
}

// SetName sets the "name" field.
func (aruo *AttachmentRefUpdateOne) SetName(s string) *AttachmentRefUpdateOne {
	aruo.mutation.SetName(s)
	return aruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (aruo *AttachmentRefUpdateOne) SetNillableName(s *string) *AttachmentRefUpdateOne {
	if s != nil {
		aruo.SetName(*s)
	}
	return aruo
}

// SetLink sets the "link" field.
func (aruo *AttachmentRefUpdateOne) SetLink(s string) *AttachmentRefUpdateOne {
	aruo.mutation.SetLink(s)
	return aruo
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (aruo *AttachmentRefUpdateOne) SetNillableLink(s *string) *AttachmentRefUpdateOne {
	if s != nil {
		aruo.SetLink(*s)
	}
	return aruo
}

// SetFreelancerID sets the "freelancer" edge to the UpworkFreelancer entity by ID.
func (aruo *AttachmentRefUpdateOne) SetFreelancerID(id string) *AttachmentRefUpdateOne {
	aruo.mutation.SetFreelancerID(id)
	return aruo
}

// SetFreelancer sets the "freelancer" edge to the UpworkFreelancer entity.
func (aruo *AttachmentRefUpdateOne) SetFreelancer(u *UpworkFreelancer) *AttachmentRefUpdateOne {
	return aruo.SetFreelancerID(u.ID)
}

// Mutation returns the AttachmentRefMutation object of the builder.
func (aruo *AttachmentRefUpdateOne) Mutation() *AttachmentRefMutation {
	return aruo.mutation
}

// ClearFreelancer clears the "freelancer" edge to the UpworkFreelancer entity.
func (aruo *AttachmentRefUpdateOne) ClearFreelancer() *AttachmentRefUpdateOne {
	aruo.mutation.ClearFreelancer()
	return aruo
}

// Where appends a list predicates to the AttachmentRefUpdate builder.
func (aruo *AttachmentRefUpdateOne) Where(ps ...predicate.AttachmentRef) *AttachmentRefUpdateOne {
	aruo.mutation.Where(ps...)
	return aruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *AttachmentRefUpdateOne) Select(field string, fields ...string) *AttachmentRefUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated AttachmentRef entity.
func (aruo *AttachmentRefUpdateOne) Save(ctx context.Context) (*AttachmentRef, error) {
	return withHooks(ctx, aruo.sqlSave, aruo.mutation, aruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AttachmentRefUpdateOne) SaveX(ctx context.Context) *AttachmentRef {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AttachmentRefUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AttachmentRefUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aruo *AttachmentRefUpdateOne) check() error {
	if _, ok := aruo.mutation.FreelancerID(); aruo.mutation.FreelancerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AttachmentRef.freelancer"`)
	}
	return nil
}

func (aruo *AttachmentRefUpdateOne) sqlSave(ctx context.Context) (_node *AttachmentRef, err error) {
	if err := aruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(attachmentref.Table, attachmentref.Columns, sqlgraph.NewFieldSpec(attachmentref.FieldID, field.TypeInt))
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AttachmentRef.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attachmentref.FieldID)
		for _, f := range fields {
			if !attachmentref.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attachmentref.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.Name(); ok {
		_spec.SetField(attachmentref.FieldName, field.TypeString, value)
	}
	if value, ok := aruo.mutation.Link(); ok {
		_spec.SetField(attachmentref.FieldLink, field.TypeString, value)
	}
	if aruo.mutation.FreelancerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachmentref.FreelancerTable,
			Columns: []string{attachmentref.FreelancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(upworkfreelancer.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.FreelancerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachmentref.FreelancerTable,
			Columns: []string{attachmentref.FreelancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(upworkfreelancer.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AttachmentRef{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachmentref.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aruo.mutation.done = true
	return _node, nil
}
