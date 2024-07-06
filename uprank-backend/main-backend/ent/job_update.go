// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notzree/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank-backend/main-backend/ent/predicate"
	"github.com/notzree/uprank-backend/main-backend/ent/upworkjob"
	"github.com/notzree/uprank-backend/main-backend/ent/user"
)

// JobUpdate is the builder for updating Job entities.
type JobUpdate struct {
	config
	hooks    []Hook
	mutation *JobMutation
}

// Where appends a list predicates to the JobUpdate builder.
func (ju *JobUpdate) Where(ps ...predicate.Job) *JobUpdate {
	ju.mutation.Where(ps...)
	return ju
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ju *JobUpdate) SetUserID(id string) *JobUpdate {
	ju.mutation.SetUserID(id)
	return ju
}

// SetUser sets the "user" edge to the User entity.
func (ju *JobUpdate) SetUser(u *User) *JobUpdate {
	return ju.SetUserID(u.ID)
}

// SetUpworkjobID sets the "upworkjob" edge to the UpworkJob entity by ID.
func (ju *JobUpdate) SetUpworkjobID(id string) *JobUpdate {
	ju.mutation.SetUpworkjobID(id)
	return ju
}

// SetNillableUpworkjobID sets the "upworkjob" edge to the UpworkJob entity by ID if the given value is not nil.
func (ju *JobUpdate) SetNillableUpworkjobID(id *string) *JobUpdate {
	if id != nil {
		ju = ju.SetUpworkjobID(*id)
	}
	return ju
}

// SetUpworkjob sets the "upworkjob" edge to the UpworkJob entity.
func (ju *JobUpdate) SetUpworkjob(u *UpworkJob) *JobUpdate {
	return ju.SetUpworkjobID(u.ID)
}

// Mutation returns the JobMutation object of the builder.
func (ju *JobUpdate) Mutation() *JobMutation {
	return ju.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ju *JobUpdate) ClearUser() *JobUpdate {
	ju.mutation.ClearUser()
	return ju
}

// ClearUpworkjob clears the "upworkjob" edge to the UpworkJob entity.
func (ju *JobUpdate) ClearUpworkjob() *JobUpdate {
	ju.mutation.ClearUpworkjob()
	return ju
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ju *JobUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ju.sqlSave, ju.mutation, ju.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ju *JobUpdate) SaveX(ctx context.Context) int {
	affected, err := ju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ju *JobUpdate) Exec(ctx context.Context) error {
	_, err := ju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ju *JobUpdate) ExecX(ctx context.Context) {
	if err := ju.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ju *JobUpdate) check() error {
	if _, ok := ju.mutation.UserID(); ju.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Job.user"`)
	}
	return nil
}

func (ju *JobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ju.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(job.Table, job.Columns, sqlgraph.NewFieldSpec(job.FieldID, field.TypeUUID))
	if ps := ju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ju.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ju.mutation.UpworkjobCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.UpworkjobIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ju.mutation.done = true
	return n, nil
}

// JobUpdateOne is the builder for updating a single Job entity.
type JobUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobMutation
}

// SetUserID sets the "user" edge to the User entity by ID.
func (juo *JobUpdateOne) SetUserID(id string) *JobUpdateOne {
	juo.mutation.SetUserID(id)
	return juo
}

// SetUser sets the "user" edge to the User entity.
func (juo *JobUpdateOne) SetUser(u *User) *JobUpdateOne {
	return juo.SetUserID(u.ID)
}

// SetUpworkjobID sets the "upworkjob" edge to the UpworkJob entity by ID.
func (juo *JobUpdateOne) SetUpworkjobID(id string) *JobUpdateOne {
	juo.mutation.SetUpworkjobID(id)
	return juo
}

// SetNillableUpworkjobID sets the "upworkjob" edge to the UpworkJob entity by ID if the given value is not nil.
func (juo *JobUpdateOne) SetNillableUpworkjobID(id *string) *JobUpdateOne {
	if id != nil {
		juo = juo.SetUpworkjobID(*id)
	}
	return juo
}

// SetUpworkjob sets the "upworkjob" edge to the UpworkJob entity.
func (juo *JobUpdateOne) SetUpworkjob(u *UpworkJob) *JobUpdateOne {
	return juo.SetUpworkjobID(u.ID)
}

// Mutation returns the JobMutation object of the builder.
func (juo *JobUpdateOne) Mutation() *JobMutation {
	return juo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (juo *JobUpdateOne) ClearUser() *JobUpdateOne {
	juo.mutation.ClearUser()
	return juo
}

// ClearUpworkjob clears the "upworkjob" edge to the UpworkJob entity.
func (juo *JobUpdateOne) ClearUpworkjob() *JobUpdateOne {
	juo.mutation.ClearUpworkjob()
	return juo
}

// Where appends a list predicates to the JobUpdate builder.
func (juo *JobUpdateOne) Where(ps ...predicate.Job) *JobUpdateOne {
	juo.mutation.Where(ps...)
	return juo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (juo *JobUpdateOne) Select(field string, fields ...string) *JobUpdateOne {
	juo.fields = append([]string{field}, fields...)
	return juo
}

// Save executes the query and returns the updated Job entity.
func (juo *JobUpdateOne) Save(ctx context.Context) (*Job, error) {
	return withHooks(ctx, juo.sqlSave, juo.mutation, juo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (juo *JobUpdateOne) SaveX(ctx context.Context) *Job {
	node, err := juo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (juo *JobUpdateOne) Exec(ctx context.Context) error {
	_, err := juo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (juo *JobUpdateOne) ExecX(ctx context.Context) {
	if err := juo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (juo *JobUpdateOne) check() error {
	if _, ok := juo.mutation.UserID(); juo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Job.user"`)
	}
	return nil
}

func (juo *JobUpdateOne) sqlSave(ctx context.Context) (_node *Job, err error) {
	if err := juo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(job.Table, job.Columns, sqlgraph.NewFieldSpec(job.FieldID, field.TypeUUID))
	id, ok := juo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Job.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := juo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, job.FieldID)
		for _, f := range fields {
			if !job.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != job.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := juo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if juo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if juo.mutation.UpworkjobCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.UpworkjobIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Job{config: juo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, juo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	juo.mutation.done = true
	return _node, nil
}
