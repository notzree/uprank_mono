// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notzree/uprank-backend/ent/attachmentref"
	"github.com/notzree/uprank-backend/ent/predicate"
)

// AttachmentRefDelete is the builder for deleting a AttachmentRef entity.
type AttachmentRefDelete struct {
	config
	hooks    []Hook
	mutation *AttachmentRefMutation
}

// Where appends a list predicates to the AttachmentRefDelete builder.
func (ard *AttachmentRefDelete) Where(ps ...predicate.AttachmentRef) *AttachmentRefDelete {
	ard.mutation.Where(ps...)
	return ard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ard *AttachmentRefDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ard.sqlExec, ard.mutation, ard.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ard *AttachmentRefDelete) ExecX(ctx context.Context) int {
	n, err := ard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ard *AttachmentRefDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(attachmentref.Table, sqlgraph.NewFieldSpec(attachmentref.FieldID, field.TypeInt))
	if ps := ard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ard.mutation.done = true
	return affected, err
}

// AttachmentRefDeleteOne is the builder for deleting a single AttachmentRef entity.
type AttachmentRefDeleteOne struct {
	ard *AttachmentRefDelete
}

// Where appends a list predicates to the AttachmentRefDelete builder.
func (ardo *AttachmentRefDeleteOne) Where(ps ...predicate.AttachmentRef) *AttachmentRefDeleteOne {
	ardo.ard.mutation.Where(ps...)
	return ardo
}

// Exec executes the deletion query.
func (ardo *AttachmentRefDeleteOne) Exec(ctx context.Context) error {
	n, err := ardo.ard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{attachmentref.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ardo *AttachmentRefDeleteOne) ExecX(ctx context.Context) {
	if err := ardo.Exec(ctx); err != nil {
		panic(err)
	}
}
