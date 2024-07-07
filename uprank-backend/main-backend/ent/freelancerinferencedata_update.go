// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
	"github.com/notzree/uprank-backend/main-backend/ent/freelancerinferencedata"
	"github.com/notzree/uprank-backend/main-backend/ent/predicate"
	"github.com/notzree/uprank-backend/main-backend/ent/upworkfreelancer"
)

// FreelancerInferenceDataUpdate is the builder for updating FreelancerInferenceData entities.
type FreelancerInferenceDataUpdate struct {
	config
	hooks    []Hook
	mutation *FreelancerInferenceDataMutation
}

// Where appends a list predicates to the FreelancerInferenceDataUpdate builder.
func (fidu *FreelancerInferenceDataUpdate) Where(ps ...predicate.FreelancerInferenceData) *FreelancerInferenceDataUpdate {
	fidu.mutation.Where(ps...)
	return fidu
}

// SetUprankReccomended sets the "uprank_reccomended" field.
func (fidu *FreelancerInferenceDataUpdate) SetUprankReccomended(b bool) *FreelancerInferenceDataUpdate {
	fidu.mutation.SetUprankReccomended(b)
	return fidu
}

// SetNillableUprankReccomended sets the "uprank_reccomended" field if the given value is not nil.
func (fidu *FreelancerInferenceDataUpdate) SetNillableUprankReccomended(b *bool) *FreelancerInferenceDataUpdate {
	if b != nil {
		fidu.SetUprankReccomended(*b)
	}
	return fidu
}

// ClearUprankReccomended clears the value of the "uprank_reccomended" field.
func (fidu *FreelancerInferenceDataUpdate) ClearUprankReccomended() *FreelancerInferenceDataUpdate {
	fidu.mutation.ClearUprankReccomended()
	return fidu
}

// SetUprankReccomendedReasons sets the "uprank_reccomended_reasons" field.
func (fidu *FreelancerInferenceDataUpdate) SetUprankReccomendedReasons(s string) *FreelancerInferenceDataUpdate {
	fidu.mutation.SetUprankReccomendedReasons(s)
	return fidu
}

// SetNillableUprankReccomendedReasons sets the "uprank_reccomended_reasons" field if the given value is not nil.
func (fidu *FreelancerInferenceDataUpdate) SetNillableUprankReccomendedReasons(s *string) *FreelancerInferenceDataUpdate {
	if s != nil {
		fidu.SetUprankReccomendedReasons(*s)
	}
	return fidu
}

// ClearUprankReccomendedReasons clears the value of the "uprank_reccomended_reasons" field.
func (fidu *FreelancerInferenceDataUpdate) ClearUprankReccomendedReasons() *FreelancerInferenceDataUpdate {
	fidu.mutation.ClearUprankReccomendedReasons()
	return fidu
}

// SetUprankNotEnoughData sets the "uprank_not_enough_data" field.
func (fidu *FreelancerInferenceDataUpdate) SetUprankNotEnoughData(b bool) *FreelancerInferenceDataUpdate {
	fidu.mutation.SetUprankNotEnoughData(b)
	return fidu
}

// SetNillableUprankNotEnoughData sets the "uprank_not_enough_data" field if the given value is not nil.
func (fidu *FreelancerInferenceDataUpdate) SetNillableUprankNotEnoughData(b *bool) *FreelancerInferenceDataUpdate {
	if b != nil {
		fidu.SetUprankNotEnoughData(*b)
	}
	return fidu
}

// ClearUprankNotEnoughData clears the value of the "uprank_not_enough_data" field.
func (fidu *FreelancerInferenceDataUpdate) ClearUprankNotEnoughData() *FreelancerInferenceDataUpdate {
	fidu.mutation.ClearUprankNotEnoughData()
	return fidu
}

// SetFinalizedRatingScore sets the "finalized_rating_score" field.
func (fidu *FreelancerInferenceDataUpdate) SetFinalizedRatingScore(f float64) *FreelancerInferenceDataUpdate {
	fidu.mutation.ResetFinalizedRatingScore()
	fidu.mutation.SetFinalizedRatingScore(f)
	return fidu
}

// SetNillableFinalizedRatingScore sets the "finalized_rating_score" field if the given value is not nil.
func (fidu *FreelancerInferenceDataUpdate) SetNillableFinalizedRatingScore(f *float64) *FreelancerInferenceDataUpdate {
	if f != nil {
		fidu.SetFinalizedRatingScore(*f)
	}
	return fidu
}

// AddFinalizedRatingScore adds f to the "finalized_rating_score" field.
func (fidu *FreelancerInferenceDataUpdate) AddFinalizedRatingScore(f float64) *FreelancerInferenceDataUpdate {
	fidu.mutation.AddFinalizedRatingScore(f)
	return fidu
}

// SetAiEstimatedDuration sets the "ai_estimated_duration" field.
func (fidu *FreelancerInferenceDataUpdate) SetAiEstimatedDuration(pg *pgtype.Interval) *FreelancerInferenceDataUpdate {
	fidu.mutation.SetAiEstimatedDuration(pg)
	return fidu
}

// ClearAiEstimatedDuration clears the value of the "ai_estimated_duration" field.
func (fidu *FreelancerInferenceDataUpdate) ClearAiEstimatedDuration() *FreelancerInferenceDataUpdate {
	fidu.mutation.ClearAiEstimatedDuration()
	return fidu
}

// SetBudgetAdherencePercentage sets the "budget_adherence_percentage" field.
func (fidu *FreelancerInferenceDataUpdate) SetBudgetAdherencePercentage(f float64) *FreelancerInferenceDataUpdate {
	fidu.mutation.ResetBudgetAdherencePercentage()
	fidu.mutation.SetBudgetAdherencePercentage(f)
	return fidu
}

// SetNillableBudgetAdherencePercentage sets the "budget_adherence_percentage" field if the given value is not nil.
func (fidu *FreelancerInferenceDataUpdate) SetNillableBudgetAdherencePercentage(f *float64) *FreelancerInferenceDataUpdate {
	if f != nil {
		fidu.SetBudgetAdherencePercentage(*f)
	}
	return fidu
}

// AddBudgetAdherencePercentage adds f to the "budget_adherence_percentage" field.
func (fidu *FreelancerInferenceDataUpdate) AddBudgetAdherencePercentage(f float64) *FreelancerInferenceDataUpdate {
	fidu.mutation.AddBudgetAdherencePercentage(f)
	return fidu
}

// ClearBudgetAdherencePercentage clears the value of the "budget_adherence_percentage" field.
func (fidu *FreelancerInferenceDataUpdate) ClearBudgetAdherencePercentage() *FreelancerInferenceDataUpdate {
	fidu.mutation.ClearBudgetAdherencePercentage()
	return fidu
}

// SetUpworkfreelancerID sets the "upworkfreelancer" edge to the UpworkFreelancer entity by ID.
func (fidu *FreelancerInferenceDataUpdate) SetUpworkfreelancerID(id string) *FreelancerInferenceDataUpdate {
	fidu.mutation.SetUpworkfreelancerID(id)
	return fidu
}

// SetUpworkfreelancer sets the "upworkfreelancer" edge to the UpworkFreelancer entity.
func (fidu *FreelancerInferenceDataUpdate) SetUpworkfreelancer(u *UpworkFreelancer) *FreelancerInferenceDataUpdate {
	return fidu.SetUpworkfreelancerID(u.ID)
}

// Mutation returns the FreelancerInferenceDataMutation object of the builder.
func (fidu *FreelancerInferenceDataUpdate) Mutation() *FreelancerInferenceDataMutation {
	return fidu.mutation
}

// ClearUpworkfreelancer clears the "upworkfreelancer" edge to the UpworkFreelancer entity.
func (fidu *FreelancerInferenceDataUpdate) ClearUpworkfreelancer() *FreelancerInferenceDataUpdate {
	fidu.mutation.ClearUpworkfreelancer()
	return fidu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fidu *FreelancerInferenceDataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fidu.sqlSave, fidu.mutation, fidu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fidu *FreelancerInferenceDataUpdate) SaveX(ctx context.Context) int {
	affected, err := fidu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fidu *FreelancerInferenceDataUpdate) Exec(ctx context.Context) error {
	_, err := fidu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fidu *FreelancerInferenceDataUpdate) ExecX(ctx context.Context) {
	if err := fidu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fidu *FreelancerInferenceDataUpdate) check() error {
	if _, ok := fidu.mutation.UpworkfreelancerID(); fidu.mutation.UpworkfreelancerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FreelancerInferenceData.upworkfreelancer"`)
	}
	return nil
}

func (fidu *FreelancerInferenceDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fidu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(freelancerinferencedata.Table, freelancerinferencedata.Columns, sqlgraph.NewFieldSpec(freelancerinferencedata.FieldID, field.TypeInt))
	if ps := fidu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fidu.mutation.UprankReccomended(); ok {
		_spec.SetField(freelancerinferencedata.FieldUprankReccomended, field.TypeBool, value)
	}
	if fidu.mutation.UprankReccomendedCleared() {
		_spec.ClearField(freelancerinferencedata.FieldUprankReccomended, field.TypeBool)
	}
	if value, ok := fidu.mutation.UprankReccomendedReasons(); ok {
		_spec.SetField(freelancerinferencedata.FieldUprankReccomendedReasons, field.TypeString, value)
	}
	if fidu.mutation.UprankReccomendedReasonsCleared() {
		_spec.ClearField(freelancerinferencedata.FieldUprankReccomendedReasons, field.TypeString)
	}
	if value, ok := fidu.mutation.UprankNotEnoughData(); ok {
		_spec.SetField(freelancerinferencedata.FieldUprankNotEnoughData, field.TypeBool, value)
	}
	if fidu.mutation.UprankNotEnoughDataCleared() {
		_spec.ClearField(freelancerinferencedata.FieldUprankNotEnoughData, field.TypeBool)
	}
	if value, ok := fidu.mutation.FinalizedRatingScore(); ok {
		_spec.SetField(freelancerinferencedata.FieldFinalizedRatingScore, field.TypeFloat64, value)
	}
	if value, ok := fidu.mutation.AddedFinalizedRatingScore(); ok {
		_spec.AddField(freelancerinferencedata.FieldFinalizedRatingScore, field.TypeFloat64, value)
	}
	if value, ok := fidu.mutation.AiEstimatedDuration(); ok {
		_spec.SetField(freelancerinferencedata.FieldAiEstimatedDuration, field.TypeOther, value)
	}
	if fidu.mutation.AiEstimatedDurationCleared() {
		_spec.ClearField(freelancerinferencedata.FieldAiEstimatedDuration, field.TypeOther)
	}
	if value, ok := fidu.mutation.BudgetAdherencePercentage(); ok {
		_spec.SetField(freelancerinferencedata.FieldBudgetAdherencePercentage, field.TypeFloat64, value)
	}
	if value, ok := fidu.mutation.AddedBudgetAdherencePercentage(); ok {
		_spec.AddField(freelancerinferencedata.FieldBudgetAdherencePercentage, field.TypeFloat64, value)
	}
	if fidu.mutation.BudgetAdherencePercentageCleared() {
		_spec.ClearField(freelancerinferencedata.FieldBudgetAdherencePercentage, field.TypeFloat64)
	}
	if fidu.mutation.UpworkfreelancerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   freelancerinferencedata.UpworkfreelancerTable,
			Columns: []string{freelancerinferencedata.UpworkfreelancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(upworkfreelancer.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fidu.mutation.UpworkfreelancerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   freelancerinferencedata.UpworkfreelancerTable,
			Columns: []string{freelancerinferencedata.UpworkfreelancerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fidu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{freelancerinferencedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fidu.mutation.done = true
	return n, nil
}

// FreelancerInferenceDataUpdateOne is the builder for updating a single FreelancerInferenceData entity.
type FreelancerInferenceDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FreelancerInferenceDataMutation
}

// SetUprankReccomended sets the "uprank_reccomended" field.
func (fiduo *FreelancerInferenceDataUpdateOne) SetUprankReccomended(b bool) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.SetUprankReccomended(b)
	return fiduo
}

// SetNillableUprankReccomended sets the "uprank_reccomended" field if the given value is not nil.
func (fiduo *FreelancerInferenceDataUpdateOne) SetNillableUprankReccomended(b *bool) *FreelancerInferenceDataUpdateOne {
	if b != nil {
		fiduo.SetUprankReccomended(*b)
	}
	return fiduo
}

// ClearUprankReccomended clears the value of the "uprank_reccomended" field.
func (fiduo *FreelancerInferenceDataUpdateOne) ClearUprankReccomended() *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.ClearUprankReccomended()
	return fiduo
}

// SetUprankReccomendedReasons sets the "uprank_reccomended_reasons" field.
func (fiduo *FreelancerInferenceDataUpdateOne) SetUprankReccomendedReasons(s string) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.SetUprankReccomendedReasons(s)
	return fiduo
}

// SetNillableUprankReccomendedReasons sets the "uprank_reccomended_reasons" field if the given value is not nil.
func (fiduo *FreelancerInferenceDataUpdateOne) SetNillableUprankReccomendedReasons(s *string) *FreelancerInferenceDataUpdateOne {
	if s != nil {
		fiduo.SetUprankReccomendedReasons(*s)
	}
	return fiduo
}

// ClearUprankReccomendedReasons clears the value of the "uprank_reccomended_reasons" field.
func (fiduo *FreelancerInferenceDataUpdateOne) ClearUprankReccomendedReasons() *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.ClearUprankReccomendedReasons()
	return fiduo
}

// SetUprankNotEnoughData sets the "uprank_not_enough_data" field.
func (fiduo *FreelancerInferenceDataUpdateOne) SetUprankNotEnoughData(b bool) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.SetUprankNotEnoughData(b)
	return fiduo
}

// SetNillableUprankNotEnoughData sets the "uprank_not_enough_data" field if the given value is not nil.
func (fiduo *FreelancerInferenceDataUpdateOne) SetNillableUprankNotEnoughData(b *bool) *FreelancerInferenceDataUpdateOne {
	if b != nil {
		fiduo.SetUprankNotEnoughData(*b)
	}
	return fiduo
}

// ClearUprankNotEnoughData clears the value of the "uprank_not_enough_data" field.
func (fiduo *FreelancerInferenceDataUpdateOne) ClearUprankNotEnoughData() *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.ClearUprankNotEnoughData()
	return fiduo
}

// SetFinalizedRatingScore sets the "finalized_rating_score" field.
func (fiduo *FreelancerInferenceDataUpdateOne) SetFinalizedRatingScore(f float64) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.ResetFinalizedRatingScore()
	fiduo.mutation.SetFinalizedRatingScore(f)
	return fiduo
}

// SetNillableFinalizedRatingScore sets the "finalized_rating_score" field if the given value is not nil.
func (fiduo *FreelancerInferenceDataUpdateOne) SetNillableFinalizedRatingScore(f *float64) *FreelancerInferenceDataUpdateOne {
	if f != nil {
		fiduo.SetFinalizedRatingScore(*f)
	}
	return fiduo
}

// AddFinalizedRatingScore adds f to the "finalized_rating_score" field.
func (fiduo *FreelancerInferenceDataUpdateOne) AddFinalizedRatingScore(f float64) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.AddFinalizedRatingScore(f)
	return fiduo
}

// SetAiEstimatedDuration sets the "ai_estimated_duration" field.
func (fiduo *FreelancerInferenceDataUpdateOne) SetAiEstimatedDuration(pg *pgtype.Interval) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.SetAiEstimatedDuration(pg)
	return fiduo
}

// ClearAiEstimatedDuration clears the value of the "ai_estimated_duration" field.
func (fiduo *FreelancerInferenceDataUpdateOne) ClearAiEstimatedDuration() *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.ClearAiEstimatedDuration()
	return fiduo
}

// SetBudgetAdherencePercentage sets the "budget_adherence_percentage" field.
func (fiduo *FreelancerInferenceDataUpdateOne) SetBudgetAdherencePercentage(f float64) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.ResetBudgetAdherencePercentage()
	fiduo.mutation.SetBudgetAdherencePercentage(f)
	return fiduo
}

// SetNillableBudgetAdherencePercentage sets the "budget_adherence_percentage" field if the given value is not nil.
func (fiduo *FreelancerInferenceDataUpdateOne) SetNillableBudgetAdherencePercentage(f *float64) *FreelancerInferenceDataUpdateOne {
	if f != nil {
		fiduo.SetBudgetAdherencePercentage(*f)
	}
	return fiduo
}

// AddBudgetAdherencePercentage adds f to the "budget_adherence_percentage" field.
func (fiduo *FreelancerInferenceDataUpdateOne) AddBudgetAdherencePercentage(f float64) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.AddBudgetAdherencePercentage(f)
	return fiduo
}

// ClearBudgetAdherencePercentage clears the value of the "budget_adherence_percentage" field.
func (fiduo *FreelancerInferenceDataUpdateOne) ClearBudgetAdherencePercentage() *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.ClearBudgetAdherencePercentage()
	return fiduo
}

// SetUpworkfreelancerID sets the "upworkfreelancer" edge to the UpworkFreelancer entity by ID.
func (fiduo *FreelancerInferenceDataUpdateOne) SetUpworkfreelancerID(id string) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.SetUpworkfreelancerID(id)
	return fiduo
}

// SetUpworkfreelancer sets the "upworkfreelancer" edge to the UpworkFreelancer entity.
func (fiduo *FreelancerInferenceDataUpdateOne) SetUpworkfreelancer(u *UpworkFreelancer) *FreelancerInferenceDataUpdateOne {
	return fiduo.SetUpworkfreelancerID(u.ID)
}

// Mutation returns the FreelancerInferenceDataMutation object of the builder.
func (fiduo *FreelancerInferenceDataUpdateOne) Mutation() *FreelancerInferenceDataMutation {
	return fiduo.mutation
}

// ClearUpworkfreelancer clears the "upworkfreelancer" edge to the UpworkFreelancer entity.
func (fiduo *FreelancerInferenceDataUpdateOne) ClearUpworkfreelancer() *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.ClearUpworkfreelancer()
	return fiduo
}

// Where appends a list predicates to the FreelancerInferenceDataUpdate builder.
func (fiduo *FreelancerInferenceDataUpdateOne) Where(ps ...predicate.FreelancerInferenceData) *FreelancerInferenceDataUpdateOne {
	fiduo.mutation.Where(ps...)
	return fiduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fiduo *FreelancerInferenceDataUpdateOne) Select(field string, fields ...string) *FreelancerInferenceDataUpdateOne {
	fiduo.fields = append([]string{field}, fields...)
	return fiduo
}

// Save executes the query and returns the updated FreelancerInferenceData entity.
func (fiduo *FreelancerInferenceDataUpdateOne) Save(ctx context.Context) (*FreelancerInferenceData, error) {
	return withHooks(ctx, fiduo.sqlSave, fiduo.mutation, fiduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fiduo *FreelancerInferenceDataUpdateOne) SaveX(ctx context.Context) *FreelancerInferenceData {
	node, err := fiduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fiduo *FreelancerInferenceDataUpdateOne) Exec(ctx context.Context) error {
	_, err := fiduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fiduo *FreelancerInferenceDataUpdateOne) ExecX(ctx context.Context) {
	if err := fiduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fiduo *FreelancerInferenceDataUpdateOne) check() error {
	if _, ok := fiduo.mutation.UpworkfreelancerID(); fiduo.mutation.UpworkfreelancerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FreelancerInferenceData.upworkfreelancer"`)
	}
	return nil
}

func (fiduo *FreelancerInferenceDataUpdateOne) sqlSave(ctx context.Context) (_node *FreelancerInferenceData, err error) {
	if err := fiduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(freelancerinferencedata.Table, freelancerinferencedata.Columns, sqlgraph.NewFieldSpec(freelancerinferencedata.FieldID, field.TypeInt))
	id, ok := fiduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FreelancerInferenceData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fiduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, freelancerinferencedata.FieldID)
		for _, f := range fields {
			if !freelancerinferencedata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != freelancerinferencedata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fiduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fiduo.mutation.UprankReccomended(); ok {
		_spec.SetField(freelancerinferencedata.FieldUprankReccomended, field.TypeBool, value)
	}
	if fiduo.mutation.UprankReccomendedCleared() {
		_spec.ClearField(freelancerinferencedata.FieldUprankReccomended, field.TypeBool)
	}
	if value, ok := fiduo.mutation.UprankReccomendedReasons(); ok {
		_spec.SetField(freelancerinferencedata.FieldUprankReccomendedReasons, field.TypeString, value)
	}
	if fiduo.mutation.UprankReccomendedReasonsCleared() {
		_spec.ClearField(freelancerinferencedata.FieldUprankReccomendedReasons, field.TypeString)
	}
	if value, ok := fiduo.mutation.UprankNotEnoughData(); ok {
		_spec.SetField(freelancerinferencedata.FieldUprankNotEnoughData, field.TypeBool, value)
	}
	if fiduo.mutation.UprankNotEnoughDataCleared() {
		_spec.ClearField(freelancerinferencedata.FieldUprankNotEnoughData, field.TypeBool)
	}
	if value, ok := fiduo.mutation.FinalizedRatingScore(); ok {
		_spec.SetField(freelancerinferencedata.FieldFinalizedRatingScore, field.TypeFloat64, value)
	}
	if value, ok := fiduo.mutation.AddedFinalizedRatingScore(); ok {
		_spec.AddField(freelancerinferencedata.FieldFinalizedRatingScore, field.TypeFloat64, value)
	}
	if value, ok := fiduo.mutation.AiEstimatedDuration(); ok {
		_spec.SetField(freelancerinferencedata.FieldAiEstimatedDuration, field.TypeOther, value)
	}
	if fiduo.mutation.AiEstimatedDurationCleared() {
		_spec.ClearField(freelancerinferencedata.FieldAiEstimatedDuration, field.TypeOther)
	}
	if value, ok := fiduo.mutation.BudgetAdherencePercentage(); ok {
		_spec.SetField(freelancerinferencedata.FieldBudgetAdherencePercentage, field.TypeFloat64, value)
	}
	if value, ok := fiduo.mutation.AddedBudgetAdherencePercentage(); ok {
		_spec.AddField(freelancerinferencedata.FieldBudgetAdherencePercentage, field.TypeFloat64, value)
	}
	if fiduo.mutation.BudgetAdherencePercentageCleared() {
		_spec.ClearField(freelancerinferencedata.FieldBudgetAdherencePercentage, field.TypeFloat64)
	}
	if fiduo.mutation.UpworkfreelancerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   freelancerinferencedata.UpworkfreelancerTable,
			Columns: []string{freelancerinferencedata.UpworkfreelancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(upworkfreelancer.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiduo.mutation.UpworkfreelancerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   freelancerinferencedata.UpworkfreelancerTable,
			Columns: []string{freelancerinferencedata.UpworkfreelancerColumn},
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
	_node = &FreelancerInferenceData{config: fiduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fiduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{freelancerinferencedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fiduo.mutation.done = true
	return _node, nil
}
