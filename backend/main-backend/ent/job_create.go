// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/schema"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/upworkjob"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/user"
)

// JobCreate is the builder for creating a Job entity.
type JobCreate struct {
	config
	mutation *JobMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOriginPlatform sets the "origin_platform" field.
func (jc *JobCreate) SetOriginPlatform(s schema.Platform) *JobCreate {
	jc.mutation.SetOriginPlatform(s)
	return jc
}

// SetID sets the "id" field.
func (jc *JobCreate) SetID(u uuid.UUID) *JobCreate {
	jc.mutation.SetID(u)
	return jc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (jc *JobCreate) SetNillableID(u *uuid.UUID) *JobCreate {
	if u != nil {
		jc.SetID(*u)
	}
	return jc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (jc *JobCreate) SetUserID(id string) *JobCreate {
	jc.mutation.SetUserID(id)
	return jc
}

// SetUser sets the "user" edge to the User entity.
func (jc *JobCreate) SetUser(u *User) *JobCreate {
	return jc.SetUserID(u.ID)
}

// SetUpworkjobID sets the "upworkjob" edge to the UpworkJob entity by ID.
func (jc *JobCreate) SetUpworkjobID(id string) *JobCreate {
	jc.mutation.SetUpworkjobID(id)
	return jc
}

// SetNillableUpworkjobID sets the "upworkjob" edge to the UpworkJob entity by ID if the given value is not nil.
func (jc *JobCreate) SetNillableUpworkjobID(id *string) *JobCreate {
	if id != nil {
		jc = jc.SetUpworkjobID(*id)
	}
	return jc
}

// SetUpworkjob sets the "upworkjob" edge to the UpworkJob entity.
func (jc *JobCreate) SetUpworkjob(u *UpworkJob) *JobCreate {
	return jc.SetUpworkjobID(u.ID)
}

// Mutation returns the JobMutation object of the builder.
func (jc *JobCreate) Mutation() *JobMutation {
	return jc.mutation
}

// Save creates the Job in the database.
func (jc *JobCreate) Save(ctx context.Context) (*Job, error) {
	jc.defaults()
	return withHooks(ctx, jc.sqlSave, jc.mutation, jc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JobCreate) SaveX(ctx context.Context) *Job {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jc *JobCreate) Exec(ctx context.Context) error {
	_, err := jc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jc *JobCreate) ExecX(ctx context.Context) {
	if err := jc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jc *JobCreate) defaults() {
	if _, ok := jc.mutation.ID(); !ok {
		v := job.DefaultID()
		jc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jc *JobCreate) check() error {
	if _, ok := jc.mutation.OriginPlatform(); !ok {
		return &ValidationError{Name: "origin_platform", err: errors.New(`ent: missing required field "Job.origin_platform"`)}
	}
	if v, ok := jc.mutation.OriginPlatform(); ok {
		if err := job.OriginPlatformValidator(v); err != nil {
			return &ValidationError{Name: "origin_platform", err: fmt.Errorf(`ent: validator failed for field "Job.origin_platform": %w`, err)}
		}
	}
	if _, ok := jc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Job.user"`)}
	}
	return nil
}

func (jc *JobCreate) sqlSave(ctx context.Context) (*Job, error) {
	if err := jc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	jc.mutation.id = &_node.ID
	jc.mutation.done = true
	return _node, nil
}

func (jc *JobCreate) createSpec() (*Job, *sqlgraph.CreateSpec) {
	var (
		_node = &Job{config: jc.config}
		_spec = sqlgraph.NewCreateSpec(job.Table, sqlgraph.NewFieldSpec(job.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = jc.conflict
	if id, ok := jc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := jc.mutation.OriginPlatform(); ok {
		_spec.SetField(job.FieldOriginPlatform, field.TypeEnum, value)
		_node.OriginPlatform = value
	}
	if nodes := jc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   job.UserTable,
			Columns: []string{job.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_job = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jc.mutation.UpworkjobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   job.UpworkjobTable,
			Columns: []string{job.UpworkjobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(upworkjob.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Job.Create().
//		SetOriginPlatform(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.JobUpsert) {
//			SetOriginPlatform(v+v).
//		}).
//		Exec(ctx)
func (jc *JobCreate) OnConflict(opts ...sql.ConflictOption) *JobUpsertOne {
	jc.conflict = opts
	return &JobUpsertOne{
		create: jc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Job.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (jc *JobCreate) OnConflictColumns(columns ...string) *JobUpsertOne {
	jc.conflict = append(jc.conflict, sql.ConflictColumns(columns...))
	return &JobUpsertOne{
		create: jc,
	}
}

type (
	// JobUpsertOne is the builder for "upsert"-ing
	//  one Job node.
	JobUpsertOne struct {
		create *JobCreate
	}

	// JobUpsert is the "OnConflict" setter.
	JobUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Job.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(job.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *JobUpsertOne) UpdateNewValues() *JobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(job.FieldID)
		}
		if _, exists := u.create.mutation.OriginPlatform(); exists {
			s.SetIgnore(job.FieldOriginPlatform)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Job.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *JobUpsertOne) Ignore() *JobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *JobUpsertOne) DoNothing() *JobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the JobCreate.OnConflict
// documentation for more info.
func (u *JobUpsertOne) Update(set func(*JobUpsert)) *JobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&JobUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *JobUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for JobCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *JobUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *JobUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: JobUpsertOne.ID is not supported by MySQL driver. Use JobUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *JobUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// JobCreateBulk is the builder for creating many Job entities in bulk.
type JobCreateBulk struct {
	config
	err      error
	builders []*JobCreate
	conflict []sql.ConflictOption
}

// Save creates the Job entities in the database.
func (jcb *JobCreateBulk) Save(ctx context.Context) ([]*Job, error) {
	if jcb.err != nil {
		return nil, jcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jcb.builders))
	nodes := make([]*Job, len(jcb.builders))
	mutators := make([]Mutator, len(jcb.builders))
	for i := range jcb.builders {
		func(i int, root context.Context) {
			builder := jcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobMutation)
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
					_, err = mutators[i+1].Mutate(root, jcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = jcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, jcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jcb *JobCreateBulk) SaveX(ctx context.Context) []*Job {
	v, err := jcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jcb *JobCreateBulk) Exec(ctx context.Context) error {
	_, err := jcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcb *JobCreateBulk) ExecX(ctx context.Context) {
	if err := jcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Job.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.JobUpsert) {
//			SetOriginPlatform(v+v).
//		}).
//		Exec(ctx)
func (jcb *JobCreateBulk) OnConflict(opts ...sql.ConflictOption) *JobUpsertBulk {
	jcb.conflict = opts
	return &JobUpsertBulk{
		create: jcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Job.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (jcb *JobCreateBulk) OnConflictColumns(columns ...string) *JobUpsertBulk {
	jcb.conflict = append(jcb.conflict, sql.ConflictColumns(columns...))
	return &JobUpsertBulk{
		create: jcb,
	}
}

// JobUpsertBulk is the builder for "upsert"-ing
// a bulk of Job nodes.
type JobUpsertBulk struct {
	create *JobCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Job.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(job.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *JobUpsertBulk) UpdateNewValues() *JobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(job.FieldID)
			}
			if _, exists := b.mutation.OriginPlatform(); exists {
				s.SetIgnore(job.FieldOriginPlatform)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Job.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *JobUpsertBulk) Ignore() *JobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *JobUpsertBulk) DoNothing() *JobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the JobCreateBulk.OnConflict
// documentation for more info.
func (u *JobUpsertBulk) Update(set func(*JobUpsert)) *JobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&JobUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *JobUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the JobCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for JobCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *JobUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
