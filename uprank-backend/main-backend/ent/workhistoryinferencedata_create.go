// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/workhistory"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/workhistoryinferencedata"
)

// WorkhistoryInferenceDataCreate is the builder for creating a WorkhistoryInferenceData entity.
type WorkhistoryInferenceDataCreate struct {
	config
	mutation *WorkhistoryInferenceDataMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetFinalizedJobRatingScore sets the "finalized_job_rating_score" field.
func (widc *WorkhistoryInferenceDataCreate) SetFinalizedJobRatingScore(f float64) *WorkhistoryInferenceDataCreate {
	widc.mutation.SetFinalizedJobRatingScore(f)
	return widc
}

// SetIsWithinBudget sets the "is_within_budget" field.
func (widc *WorkhistoryInferenceDataCreate) SetIsWithinBudget(b bool) *WorkhistoryInferenceDataCreate {
	widc.mutation.SetIsWithinBudget(b)
	return widc
}

// SetWorkHistoriesID sets the "work_histories" edge to the WorkHistory entity by ID.
func (widc *WorkhistoryInferenceDataCreate) SetWorkHistoriesID(id int) *WorkhistoryInferenceDataCreate {
	widc.mutation.SetWorkHistoriesID(id)
	return widc
}

// SetWorkHistories sets the "work_histories" edge to the WorkHistory entity.
func (widc *WorkhistoryInferenceDataCreate) SetWorkHistories(w *WorkHistory) *WorkhistoryInferenceDataCreate {
	return widc.SetWorkHistoriesID(w.ID)
}

// Mutation returns the WorkhistoryInferenceDataMutation object of the builder.
func (widc *WorkhistoryInferenceDataCreate) Mutation() *WorkhistoryInferenceDataMutation {
	return widc.mutation
}

// Save creates the WorkhistoryInferenceData in the database.
func (widc *WorkhistoryInferenceDataCreate) Save(ctx context.Context) (*WorkhistoryInferenceData, error) {
	return withHooks(ctx, widc.sqlSave, widc.mutation, widc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (widc *WorkhistoryInferenceDataCreate) SaveX(ctx context.Context) *WorkhistoryInferenceData {
	v, err := widc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (widc *WorkhistoryInferenceDataCreate) Exec(ctx context.Context) error {
	_, err := widc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (widc *WorkhistoryInferenceDataCreate) ExecX(ctx context.Context) {
	if err := widc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (widc *WorkhistoryInferenceDataCreate) check() error {
	if _, ok := widc.mutation.FinalizedJobRatingScore(); !ok {
		return &ValidationError{Name: "finalized_job_rating_score", err: errors.New(`ent: missing required field "WorkhistoryInferenceData.finalized_job_rating_score"`)}
	}
	if _, ok := widc.mutation.IsWithinBudget(); !ok {
		return &ValidationError{Name: "is_within_budget", err: errors.New(`ent: missing required field "WorkhistoryInferenceData.is_within_budget"`)}
	}
	if _, ok := widc.mutation.WorkHistoriesID(); !ok {
		return &ValidationError{Name: "work_histories", err: errors.New(`ent: missing required edge "WorkhistoryInferenceData.work_histories"`)}
	}
	return nil
}

func (widc *WorkhistoryInferenceDataCreate) sqlSave(ctx context.Context) (*WorkhistoryInferenceData, error) {
	if err := widc.check(); err != nil {
		return nil, err
	}
	_node, _spec := widc.createSpec()
	if err := sqlgraph.CreateNode(ctx, widc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	widc.mutation.id = &_node.ID
	widc.mutation.done = true
	return _node, nil
}

func (widc *WorkhistoryInferenceDataCreate) createSpec() (*WorkhistoryInferenceData, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkhistoryInferenceData{config: widc.config}
		_spec = sqlgraph.NewCreateSpec(workhistoryinferencedata.Table, sqlgraph.NewFieldSpec(workhistoryinferencedata.FieldID, field.TypeInt))
	)
	_spec.OnConflict = widc.conflict
	if value, ok := widc.mutation.FinalizedJobRatingScore(); ok {
		_spec.SetField(workhistoryinferencedata.FieldFinalizedJobRatingScore, field.TypeFloat64, value)
		_node.FinalizedJobRatingScore = value
	}
	if value, ok := widc.mutation.IsWithinBudget(); ok {
		_spec.SetField(workhistoryinferencedata.FieldIsWithinBudget, field.TypeBool, value)
		_node.IsWithinBudget = value
	}
	if nodes := widc.mutation.WorkHistoriesIDs(); len(nodes) > 0 {
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
		_node.work_history_work_history_inference_data = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.WorkhistoryInferenceData.Create().
//		SetFinalizedJobRatingScore(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WorkhistoryInferenceDataUpsert) {
//			SetFinalizedJobRatingScore(v+v).
//		}).
//		Exec(ctx)
func (widc *WorkhistoryInferenceDataCreate) OnConflict(opts ...sql.ConflictOption) *WorkhistoryInferenceDataUpsertOne {
	widc.conflict = opts
	return &WorkhistoryInferenceDataUpsertOne{
		create: widc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.WorkhistoryInferenceData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (widc *WorkhistoryInferenceDataCreate) OnConflictColumns(columns ...string) *WorkhistoryInferenceDataUpsertOne {
	widc.conflict = append(widc.conflict, sql.ConflictColumns(columns...))
	return &WorkhistoryInferenceDataUpsertOne{
		create: widc,
	}
}

type (
	// WorkhistoryInferenceDataUpsertOne is the builder for "upsert"-ing
	//  one WorkhistoryInferenceData node.
	WorkhistoryInferenceDataUpsertOne struct {
		create *WorkhistoryInferenceDataCreate
	}

	// WorkhistoryInferenceDataUpsert is the "OnConflict" setter.
	WorkhistoryInferenceDataUpsert struct {
		*sql.UpdateSet
	}
)

// SetFinalizedJobRatingScore sets the "finalized_job_rating_score" field.
func (u *WorkhistoryInferenceDataUpsert) SetFinalizedJobRatingScore(v float64) *WorkhistoryInferenceDataUpsert {
	u.Set(workhistoryinferencedata.FieldFinalizedJobRatingScore, v)
	return u
}

// UpdateFinalizedJobRatingScore sets the "finalized_job_rating_score" field to the value that was provided on create.
func (u *WorkhistoryInferenceDataUpsert) UpdateFinalizedJobRatingScore() *WorkhistoryInferenceDataUpsert {
	u.SetExcluded(workhistoryinferencedata.FieldFinalizedJobRatingScore)
	return u
}

// AddFinalizedJobRatingScore adds v to the "finalized_job_rating_score" field.
func (u *WorkhistoryInferenceDataUpsert) AddFinalizedJobRatingScore(v float64) *WorkhistoryInferenceDataUpsert {
	u.Add(workhistoryinferencedata.FieldFinalizedJobRatingScore, v)
	return u
}

// SetIsWithinBudget sets the "is_within_budget" field.
func (u *WorkhistoryInferenceDataUpsert) SetIsWithinBudget(v bool) *WorkhistoryInferenceDataUpsert {
	u.Set(workhistoryinferencedata.FieldIsWithinBudget, v)
	return u
}

// UpdateIsWithinBudget sets the "is_within_budget" field to the value that was provided on create.
func (u *WorkhistoryInferenceDataUpsert) UpdateIsWithinBudget() *WorkhistoryInferenceDataUpsert {
	u.SetExcluded(workhistoryinferencedata.FieldIsWithinBudget)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.WorkhistoryInferenceData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *WorkhistoryInferenceDataUpsertOne) UpdateNewValues() *WorkhistoryInferenceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.WorkhistoryInferenceData.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *WorkhistoryInferenceDataUpsertOne) Ignore() *WorkhistoryInferenceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WorkhistoryInferenceDataUpsertOne) DoNothing() *WorkhistoryInferenceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WorkhistoryInferenceDataCreate.OnConflict
// documentation for more info.
func (u *WorkhistoryInferenceDataUpsertOne) Update(set func(*WorkhistoryInferenceDataUpsert)) *WorkhistoryInferenceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WorkhistoryInferenceDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetFinalizedJobRatingScore sets the "finalized_job_rating_score" field.
func (u *WorkhistoryInferenceDataUpsertOne) SetFinalizedJobRatingScore(v float64) *WorkhistoryInferenceDataUpsertOne {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.SetFinalizedJobRatingScore(v)
	})
}

// AddFinalizedJobRatingScore adds v to the "finalized_job_rating_score" field.
func (u *WorkhistoryInferenceDataUpsertOne) AddFinalizedJobRatingScore(v float64) *WorkhistoryInferenceDataUpsertOne {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.AddFinalizedJobRatingScore(v)
	})
}

// UpdateFinalizedJobRatingScore sets the "finalized_job_rating_score" field to the value that was provided on create.
func (u *WorkhistoryInferenceDataUpsertOne) UpdateFinalizedJobRatingScore() *WorkhistoryInferenceDataUpsertOne {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.UpdateFinalizedJobRatingScore()
	})
}

// SetIsWithinBudget sets the "is_within_budget" field.
func (u *WorkhistoryInferenceDataUpsertOne) SetIsWithinBudget(v bool) *WorkhistoryInferenceDataUpsertOne {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.SetIsWithinBudget(v)
	})
}

// UpdateIsWithinBudget sets the "is_within_budget" field to the value that was provided on create.
func (u *WorkhistoryInferenceDataUpsertOne) UpdateIsWithinBudget() *WorkhistoryInferenceDataUpsertOne {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.UpdateIsWithinBudget()
	})
}

// Exec executes the query.
func (u *WorkhistoryInferenceDataUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WorkhistoryInferenceDataCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WorkhistoryInferenceDataUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *WorkhistoryInferenceDataUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *WorkhistoryInferenceDataUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// WorkhistoryInferenceDataCreateBulk is the builder for creating many WorkhistoryInferenceData entities in bulk.
type WorkhistoryInferenceDataCreateBulk struct {
	config
	err      error
	builders []*WorkhistoryInferenceDataCreate
	conflict []sql.ConflictOption
}

// Save creates the WorkhistoryInferenceData entities in the database.
func (widcb *WorkhistoryInferenceDataCreateBulk) Save(ctx context.Context) ([]*WorkhistoryInferenceData, error) {
	if widcb.err != nil {
		return nil, widcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(widcb.builders))
	nodes := make([]*WorkhistoryInferenceData, len(widcb.builders))
	mutators := make([]Mutator, len(widcb.builders))
	for i := range widcb.builders {
		func(i int, root context.Context) {
			builder := widcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkhistoryInferenceDataMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, widcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = widcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, widcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, widcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (widcb *WorkhistoryInferenceDataCreateBulk) SaveX(ctx context.Context) []*WorkhistoryInferenceData {
	v, err := widcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (widcb *WorkhistoryInferenceDataCreateBulk) Exec(ctx context.Context) error {
	_, err := widcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (widcb *WorkhistoryInferenceDataCreateBulk) ExecX(ctx context.Context) {
	if err := widcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.WorkhistoryInferenceData.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WorkhistoryInferenceDataUpsert) {
//			SetFinalizedJobRatingScore(v+v).
//		}).
//		Exec(ctx)
func (widcb *WorkhistoryInferenceDataCreateBulk) OnConflict(opts ...sql.ConflictOption) *WorkhistoryInferenceDataUpsertBulk {
	widcb.conflict = opts
	return &WorkhistoryInferenceDataUpsertBulk{
		create: widcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.WorkhistoryInferenceData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (widcb *WorkhistoryInferenceDataCreateBulk) OnConflictColumns(columns ...string) *WorkhistoryInferenceDataUpsertBulk {
	widcb.conflict = append(widcb.conflict, sql.ConflictColumns(columns...))
	return &WorkhistoryInferenceDataUpsertBulk{
		create: widcb,
	}
}

// WorkhistoryInferenceDataUpsertBulk is the builder for "upsert"-ing
// a bulk of WorkhistoryInferenceData nodes.
type WorkhistoryInferenceDataUpsertBulk struct {
	create *WorkhistoryInferenceDataCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.WorkhistoryInferenceData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *WorkhistoryInferenceDataUpsertBulk) UpdateNewValues() *WorkhistoryInferenceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.WorkhistoryInferenceData.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *WorkhistoryInferenceDataUpsertBulk) Ignore() *WorkhistoryInferenceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WorkhistoryInferenceDataUpsertBulk) DoNothing() *WorkhistoryInferenceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WorkhistoryInferenceDataCreateBulk.OnConflict
// documentation for more info.
func (u *WorkhistoryInferenceDataUpsertBulk) Update(set func(*WorkhistoryInferenceDataUpsert)) *WorkhistoryInferenceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WorkhistoryInferenceDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetFinalizedJobRatingScore sets the "finalized_job_rating_score" field.
func (u *WorkhistoryInferenceDataUpsertBulk) SetFinalizedJobRatingScore(v float64) *WorkhistoryInferenceDataUpsertBulk {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.SetFinalizedJobRatingScore(v)
	})
}

// AddFinalizedJobRatingScore adds v to the "finalized_job_rating_score" field.
func (u *WorkhistoryInferenceDataUpsertBulk) AddFinalizedJobRatingScore(v float64) *WorkhistoryInferenceDataUpsertBulk {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.AddFinalizedJobRatingScore(v)
	})
}

// UpdateFinalizedJobRatingScore sets the "finalized_job_rating_score" field to the value that was provided on create.
func (u *WorkhistoryInferenceDataUpsertBulk) UpdateFinalizedJobRatingScore() *WorkhistoryInferenceDataUpsertBulk {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.UpdateFinalizedJobRatingScore()
	})
}

// SetIsWithinBudget sets the "is_within_budget" field.
func (u *WorkhistoryInferenceDataUpsertBulk) SetIsWithinBudget(v bool) *WorkhistoryInferenceDataUpsertBulk {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.SetIsWithinBudget(v)
	})
}

// UpdateIsWithinBudget sets the "is_within_budget" field to the value that was provided on create.
func (u *WorkhistoryInferenceDataUpsertBulk) UpdateIsWithinBudget() *WorkhistoryInferenceDataUpsertBulk {
	return u.Update(func(s *WorkhistoryInferenceDataUpsert) {
		s.UpdateIsWithinBudget()
	})
}

// Exec executes the query.
func (u *WorkhistoryInferenceDataUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the WorkhistoryInferenceDataCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WorkhistoryInferenceDataCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WorkhistoryInferenceDataUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
