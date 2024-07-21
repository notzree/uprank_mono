// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/predicate"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/workhistory"
)

// WorkHistoryDelete is the builder for deleting a WorkHistory entity.
type WorkHistoryDelete struct {
	config
	hooks    []Hook
	mutation *WorkHistoryMutation
}

// Where appends a list predicates to the WorkHistoryDelete builder.
func (whd *WorkHistoryDelete) Where(ps ...predicate.WorkHistory) *WorkHistoryDelete {
	whd.mutation.Where(ps...)
	return whd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (whd *WorkHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, whd.sqlExec, whd.mutation, whd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (whd *WorkHistoryDelete) ExecX(ctx context.Context) int {
	n, err := whd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (whd *WorkHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(workhistory.Table, sqlgraph.NewFieldSpec(workhistory.FieldID, field.TypeInt))
	if ps := whd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, whd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	whd.mutation.done = true
	return affected, err
}

// WorkHistoryDeleteOne is the builder for deleting a single WorkHistory entity.
type WorkHistoryDeleteOne struct {
	whd *WorkHistoryDelete
}

// Where appends a list predicates to the WorkHistoryDelete builder.
func (whdo *WorkHistoryDeleteOne) Where(ps ...predicate.WorkHistory) *WorkHistoryDeleteOne {
	whdo.whd.mutation.Where(ps...)
	return whdo
}

// Exec executes the deletion query.
func (whdo *WorkHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := whdo.whd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (whdo *WorkHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := whdo.Exec(ctx); err != nil {
		panic(err)
	}
}