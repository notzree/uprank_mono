// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/predicate"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/workhistory"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/workhistoryinferencedata"
)

// WorkhistoryInferenceDataUpdate is the builder for updating WorkhistoryInferenceData entities.
type WorkhistoryInferenceDataUpdate struct {
	config
	hooks    []Hook
	mutation *WorkhistoryInferenceDataMutation
}

// Where appends a list predicates to the WorkhistoryInferenceDataUpdate builder.
func (widu *WorkhistoryInferenceDataUpdate) Where(ps ...predicate.WorkhistoryInferenceData) *WorkhistoryInferenceDataUpdate {
	widu.mutation.Where(ps...)
	return widu
}

// SetFinalizedJobRatingScore sets the "finalized_job_rating_score" field.
func (widu *WorkhistoryInferenceDataUpdate) SetFinalizedJobRatingScore(f float64) *WorkhistoryInferenceDataUpdate {
	widu.mutation.ResetFinalizedJobRatingScore()
	widu.mutation.SetFinalizedJobRatingScore(f)
	return widu
}

// SetNillableFinalizedJobRatingScore sets the "finalized_job_rating_score" field if the given value is not nil.
func (widu *WorkhistoryInferenceDataUpdate) SetNillableFinalizedJobRatingScore(f *float64) *WorkhistoryInferenceDataUpdate {
	if f != nil {
		widu.SetFinalizedJobRatingScore(*f)
	}
	return widu
}

// AddFinalizedJobRatingScore adds f to the "finalized_job_rating_score" field.
func (widu *WorkhistoryInferenceDataUpdate) AddFinalizedJobRatingScore(f float64) *WorkhistoryInferenceDataUpdate {
	widu.mutation.AddFinalizedJobRatingScore(f)
	return widu
}

// SetIsWithinBudget sets the "is_within_budget" field.
func (widu *WorkhistoryInferenceDataUpdate) SetIsWithinBudget(b bool) *WorkhistoryInferenceDataUpdate {
	widu.mutation.SetIsWithinBudget(b)
	return widu
}

// SetNillableIsWithinBudget sets the "is_within_budget" field if the given value is not nil.
func (widu *WorkhistoryInferenceDataUpdate) SetNillableIsWithinBudget(b *bool) *WorkhistoryInferenceDataUpdate {
	if b != nil {
		widu.SetIsWithinBudget(*b)
	}
	return widu
}

// SetWorkHistoriesID sets the "work_histories" edge to the WorkHistory entity by ID.
func (widu *WorkhistoryInferenceDataUpdate) SetWorkHistoriesID(id int) *WorkhistoryInferenceDataUpdate {
	widu.mutation.SetWorkHistoriesID(id)
	return widu
}

// SetWorkHistories sets the "work_histories" edge to the WorkHistory entity.
func (widu *WorkhistoryInferenceDataUpdate) SetWorkHistories(w *WorkHistory) *WorkhistoryInferenceDataUpdate {
	return widu.SetWorkHistoriesID(w.ID)
}

// Mutation returns the WorkhistoryInferenceDataMutation object of the builder.
func (widu *WorkhistoryInferenceDataUpdate) Mutation() *WorkhistoryInferenceDataMutation {
	return widu.mutation
}

// ClearWorkHistories clears the "work_histories" edge to the WorkHistory entity.
func (widu *WorkhistoryInferenceDataUpdate) ClearWorkHistories() *WorkhistoryInferenceDataUpdate {
	widu.mutation.ClearWorkHistories()
	return widu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (widu *WorkhistoryInferenceDataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, widu.sqlSave, widu.mutation, widu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (widu *WorkhistoryInferenceDataUpdate) SaveX(ctx context.Context) int {
	affected, err := widu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (widu *WorkhistoryInferenceDataUpdate) Exec(ctx context.Context) error {
	_, err := widu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (widu *WorkhistoryInferenceDataUpdate) ExecX(ctx context.Context) {
	if err := widu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (widu *WorkhistoryInferenceDataUpdate) check() error {
	if _, ok := widu.mutation.WorkHistoriesID(); widu.mutation.WorkHistoriesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkhistoryInferenceData.work_histories"`)
	}
	return nil
}

func (widu *WorkhistoryInferenceDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := widu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workhistoryinferencedata.Table, workhistoryinferencedata.Columns, sqlgraph.NewFieldSpec(workhistoryinferencedata.FieldID, field.TypeInt))
	if ps := widu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := widu.mutation.FinalizedJobRatingScore(); ok {
		_spec.SetField(workhistoryinferencedata.FieldFinalizedJobRatingScore, field.TypeFloat64, value)
	}
	if value, ok := widu.mutation.AddedFinalizedJobRatingScore(); ok {
		_spec.AddField(workhistoryinferencedata.FieldFinalizedJobRatingScore, field.TypeFloat64, value)
	}
	if value, ok := widu.mutation.IsWithinBudget(); ok {
		_spec.SetField(workhistoryinferencedata.FieldIsWithinBudget, field.TypeBool, value)
	}
	if widu.mutation.WorkHistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workhistoryinferencedata.WorkHistoriesTable,
			Columns: []string{workhistoryinferencedata.WorkHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workhistory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := widu.mutation.WorkHistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workhistoryinferencedata.WorkHistoriesTable,
			Columns: []string{workhistoryinferencedata.WorkHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workhistory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, widu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workhistoryinferencedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	widu.mutation.done = true
	return n, nil
}

// WorkhistoryInferenceDataUpdateOne is the builder for updating a single WorkhistoryInferenceData entity.
type WorkhistoryInferenceDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkhistoryInferenceDataMutation
}

// SetFinalizedJobRatingScore sets the "finalized_job_rating_score" field.
func (widuo *WorkhistoryInferenceDataUpdateOne) SetFinalizedJobRatingScore(f float64) *WorkhistoryInferenceDataUpdateOne {
	widuo.mutation.ResetFinalizedJobRatingScore()
	widuo.mutation.SetFinalizedJobRatingScore(f)
	return widuo
}

// SetNillableFinalizedJobRatingScore sets the "finalized_job_rating_score" field if the given value is not nil.
func (widuo *WorkhistoryInferenceDataUpdateOne) SetNillableFinalizedJobRatingScore(f *float64) *WorkhistoryInferenceDataUpdateOne {
	if f != nil {
		widuo.SetFinalizedJobRatingScore(*f)
	}
	return widuo
}

// AddFinalizedJobRatingScore adds f to the "finalized_job_rating_score" field.
func (widuo *WorkhistoryInferenceDataUpdateOne) AddFinalizedJobRatingScore(f float64) *WorkhistoryInferenceDataUpdateOne {
	widuo.mutation.AddFinalizedJobRatingScore(f)
	return widuo
}

// SetIsWithinBudget sets the "is_within_budget" field.
func (widuo *WorkhistoryInferenceDataUpdateOne) SetIsWithinBudget(b bool) *WorkhistoryInferenceDataUpdateOne {
	widuo.mutation.SetIsWithinBudget(b)
	return widuo
}

// SetNillableIsWithinBudget sets the "is_within_budget" field if the given value is not nil.
func (widuo *WorkhistoryInferenceDataUpdateOne) SetNillableIsWithinBudget(b *bool) *WorkhistoryInferenceDataUpdateOne {
	if b != nil {
		widuo.SetIsWithinBudget(*b)
	}
	return widuo
}

// SetWorkHistoriesID sets the "work_histories" edge to the WorkHistory entity by ID.
func (widuo *WorkhistoryInferenceDataUpdateOne) SetWorkHistoriesID(id int) *WorkhistoryInferenceDataUpdateOne {
	widuo.mutation.SetWorkHistoriesID(id)
	return widuo
}

// SetWorkHistories sets the "work_histories" edge to the WorkHistory entity.
func (widuo *WorkhistoryInferenceDataUpdateOne) SetWorkHistories(w *WorkHistory) *WorkhistoryInferenceDataUpdateOne {
	return widuo.SetWorkHistoriesID(w.ID)
}

// Mutation returns the WorkhistoryInferenceDataMutation object of the builder.
func (widuo *WorkhistoryInferenceDataUpdateOne) Mutation() *WorkhistoryInferenceDataMutation {
	return widuo.mutation
}

// ClearWorkHistories clears the "work_histories" edge to the WorkHistory entity.
func (widuo *WorkhistoryInferenceDataUpdateOne) ClearWorkHistories() *WorkhistoryInferenceDataUpdateOne {
	widuo.mutation.ClearWorkHistories()
	return widuo
}

// Where appends a list predicates to the WorkhistoryInferenceDataUpdate builder.
func (widuo *WorkhistoryInferenceDataUpdateOne) Where(ps ...predicate.WorkhistoryInferenceData) *WorkhistoryInferenceDataUpdateOne {
	widuo.mutation.Where(ps...)
	return widuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (widuo *WorkhistoryInferenceDataUpdateOne) Select(field string, fields ...string) *WorkhistoryInferenceDataUpdateOne {
	widuo.fields = append([]string{field}, fields...)
	return widuo
}

// Save executes the query and returns the updated WorkhistoryInferenceData entity.
func (widuo *WorkhistoryInferenceDataUpdateOne) Save(ctx context.Context) (*WorkhistoryInferenceData, error) {
	return withHooks(ctx, widuo.sqlSave, widuo.mutation, widuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (widuo *WorkhistoryInferenceDataUpdateOne) SaveX(ctx context.Context) *WorkhistoryInferenceData {
	node, err := widuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (widuo *WorkhistoryInferenceDataUpdateOne) Exec(ctx context.Context) error {
	_, err := widuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (widuo *WorkhistoryInferenceDataUpdateOne) ExecX(ctx context.Context) {
	if err := widuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (widuo *WorkhistoryInferenceDataUpdateOne) check() error {
	if _, ok := widuo.mutation.WorkHistoriesID(); widuo.mutation.WorkHistoriesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkhistoryInferenceData.work_histories"`)
	}
	return nil
}

func (widuo *WorkhistoryInferenceDataUpdateOne) sqlSave(ctx context.Context) (_node *WorkhistoryInferenceData, err error) {
	if err := widuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workhistoryinferencedata.Table, workhistoryinferencedata.Columns, sqlgraph.NewFieldSpec(workhistoryinferencedata.FieldID, field.TypeInt))
	id, ok := widuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkhistoryInferenceData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := widuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workhistoryinferencedata.FieldID)
		for _, f := range fields {
			if !workhistoryinferencedata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workhistoryinferencedata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := widuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := widuo.mutation.FinalizedJobRatingScore(); ok {
		_spec.SetField(workhistoryinferencedata.FieldFinalizedJobRatingScore, field.TypeFloat64, value)
	}
	if value, ok := widuo.mutation.AddedFinalizedJobRatingScore(); ok {
		_spec.AddField(workhistoryinferencedata.FieldFinalizedJobRatingScore, field.TypeFloat64, value)
	}
	if value, ok := widuo.mutation.IsWithinBudget(); ok {
		_spec.SetField(workhistoryinferencedata.FieldIsWithinBudget, field.TypeBool, value)
	}
	if widuo.mutation.WorkHistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workhistoryinferencedata.WorkHistoriesTable,
			Columns: []string{workhistoryinferencedata.WorkHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workhistory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := widuo.mutation.WorkHistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workhistoryinferencedata.WorkHistoriesTable,
			Columns: []string{workhistoryinferencedata.WorkHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workhistory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkhistoryInferenceData{config: widuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, widuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workhistoryinferencedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	widuo.mutation.done = true
	return _node, nil
}
