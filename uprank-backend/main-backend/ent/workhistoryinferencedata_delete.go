// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/predicate"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/workhistoryinferencedata"
)

// WorkhistoryInferenceDataDelete is the builder for deleting a WorkhistoryInferenceData entity.
type WorkhistoryInferenceDataDelete struct {
	config
	hooks    []Hook
	mutation *WorkhistoryInferenceDataMutation
}

// Where appends a list predicates to the WorkhistoryInferenceDataDelete builder.
func (widd *WorkhistoryInferenceDataDelete) Where(ps ...predicate.WorkhistoryInferenceData) *WorkhistoryInferenceDataDelete {
	widd.mutation.Where(ps...)
	return widd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (widd *WorkhistoryInferenceDataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, widd.sqlExec, widd.mutation, widd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (widd *WorkhistoryInferenceDataDelete) ExecX(ctx context.Context) int {
	n, err := widd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (widd *WorkhistoryInferenceDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(workhistoryinferencedata.Table, sqlgraph.NewFieldSpec(workhistoryinferencedata.FieldID, field.TypeInt))
	if ps := widd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, widd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	widd.mutation.done = true
	return affected, err
}

// WorkhistoryInferenceDataDeleteOne is the builder for deleting a single WorkhistoryInferenceData entity.
type WorkhistoryInferenceDataDeleteOne struct {
	widd *WorkhistoryInferenceDataDelete
}

// Where appends a list predicates to the WorkhistoryInferenceDataDelete builder.
func (widdo *WorkhistoryInferenceDataDeleteOne) Where(ps ...predicate.WorkhistoryInferenceData) *WorkhistoryInferenceDataDeleteOne {
	widdo.widd.mutation.Where(ps...)
	return widdo
}

// Exec executes the deletion query.
func (widdo *WorkhistoryInferenceDataDeleteOne) Exec(ctx context.Context) error {
	n, err := widdo.widd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workhistoryinferencedata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (widdo *WorkhistoryInferenceDataDeleteOne) ExecX(ctx context.Context) {
	if err := widdo.Exec(ctx); err != nil {
		panic(err)
	}
}
