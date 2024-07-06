// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank-backend/main-backend/ent/predicate"
	"github.com/notzree/uprank-backend/main-backend/ent/upworkfreelancer"
	"github.com/notzree/uprank-backend/main-backend/ent/upworkjob"
	"github.com/notzree/uprank-backend/main-backend/ent/user"
)

// UpworkJobQuery is the builder for querying UpworkJob entities.
type UpworkJobQuery struct {
	config
	ctx                  *QueryContext
	order                []upworkjob.OrderOption
	inters               []Interceptor
	predicates           []predicate.UpworkJob
	withUpworkfreelancer *UpworkFreelancerQuery
	withJob              *JobQuery
	withUser             *UserQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UpworkJobQuery builder.
func (ujq *UpworkJobQuery) Where(ps ...predicate.UpworkJob) *UpworkJobQuery {
	ujq.predicates = append(ujq.predicates, ps...)
	return ujq
}

// Limit the number of records to be returned by this query.
func (ujq *UpworkJobQuery) Limit(limit int) *UpworkJobQuery {
	ujq.ctx.Limit = &limit
	return ujq
}

// Offset to start from.
func (ujq *UpworkJobQuery) Offset(offset int) *UpworkJobQuery {
	ujq.ctx.Offset = &offset
	return ujq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ujq *UpworkJobQuery) Unique(unique bool) *UpworkJobQuery {
	ujq.ctx.Unique = &unique
	return ujq
}

// Order specifies how the records should be ordered.
func (ujq *UpworkJobQuery) Order(o ...upworkjob.OrderOption) *UpworkJobQuery {
	ujq.order = append(ujq.order, o...)
	return ujq
}

// QueryUpworkfreelancer chains the current query on the "upworkfreelancer" edge.
func (ujq *UpworkJobQuery) QueryUpworkfreelancer() *UpworkFreelancerQuery {
	query := (&UpworkFreelancerClient{config: ujq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ujq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ujq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upworkjob.Table, upworkjob.FieldID, selector),
			sqlgraph.To(upworkfreelancer.Table, upworkfreelancer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, upworkjob.UpworkfreelancerTable, upworkjob.UpworkfreelancerPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ujq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryJob chains the current query on the "job" edge.
func (ujq *UpworkJobQuery) QueryJob() *JobQuery {
	query := (&JobClient{config: ujq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ujq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ujq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upworkjob.Table, upworkjob.FieldID, selector),
			sqlgraph.To(job.Table, job.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, upworkjob.JobTable, upworkjob.JobColumn),
		)
		fromU = sqlgraph.SetNeighbors(ujq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (ujq *UpworkJobQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ujq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ujq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ujq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upworkjob.Table, upworkjob.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, upworkjob.UserTable, upworkjob.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ujq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UpworkJob entity from the query.
// Returns a *NotFoundError when no UpworkJob was found.
func (ujq *UpworkJobQuery) First(ctx context.Context) (*UpworkJob, error) {
	nodes, err := ujq.Limit(1).All(setContextOp(ctx, ujq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{upworkjob.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ujq *UpworkJobQuery) FirstX(ctx context.Context) *UpworkJob {
	node, err := ujq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UpworkJob ID from the query.
// Returns a *NotFoundError when no UpworkJob ID was found.
func (ujq *UpworkJobQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ujq.Limit(1).IDs(setContextOp(ctx, ujq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{upworkjob.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ujq *UpworkJobQuery) FirstIDX(ctx context.Context) string {
	id, err := ujq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UpworkJob entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UpworkJob entity is found.
// Returns a *NotFoundError when no UpworkJob entities are found.
func (ujq *UpworkJobQuery) Only(ctx context.Context) (*UpworkJob, error) {
	nodes, err := ujq.Limit(2).All(setContextOp(ctx, ujq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{upworkjob.Label}
	default:
		return nil, &NotSingularError{upworkjob.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ujq *UpworkJobQuery) OnlyX(ctx context.Context) *UpworkJob {
	node, err := ujq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UpworkJob ID in the query.
// Returns a *NotSingularError when more than one UpworkJob ID is found.
// Returns a *NotFoundError when no entities are found.
func (ujq *UpworkJobQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ujq.Limit(2).IDs(setContextOp(ctx, ujq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{upworkjob.Label}
	default:
		err = &NotSingularError{upworkjob.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ujq *UpworkJobQuery) OnlyIDX(ctx context.Context) string {
	id, err := ujq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UpworkJobs.
func (ujq *UpworkJobQuery) All(ctx context.Context) ([]*UpworkJob, error) {
	ctx = setContextOp(ctx, ujq.ctx, "All")
	if err := ujq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UpworkJob, *UpworkJobQuery]()
	return withInterceptors[[]*UpworkJob](ctx, ujq, qr, ujq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ujq *UpworkJobQuery) AllX(ctx context.Context) []*UpworkJob {
	nodes, err := ujq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UpworkJob IDs.
func (ujq *UpworkJobQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ujq.ctx.Unique == nil && ujq.path != nil {
		ujq.Unique(true)
	}
	ctx = setContextOp(ctx, ujq.ctx, "IDs")
	if err = ujq.Select(upworkjob.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ujq *UpworkJobQuery) IDsX(ctx context.Context) []string {
	ids, err := ujq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ujq *UpworkJobQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ujq.ctx, "Count")
	if err := ujq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ujq, querierCount[*UpworkJobQuery](), ujq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ujq *UpworkJobQuery) CountX(ctx context.Context) int {
	count, err := ujq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ujq *UpworkJobQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ujq.ctx, "Exist")
	switch _, err := ujq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ujq *UpworkJobQuery) ExistX(ctx context.Context) bool {
	exist, err := ujq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UpworkJobQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ujq *UpworkJobQuery) Clone() *UpworkJobQuery {
	if ujq == nil {
		return nil
	}
	return &UpworkJobQuery{
		config:               ujq.config,
		ctx:                  ujq.ctx.Clone(),
		order:                append([]upworkjob.OrderOption{}, ujq.order...),
		inters:               append([]Interceptor{}, ujq.inters...),
		predicates:           append([]predicate.UpworkJob{}, ujq.predicates...),
		withUpworkfreelancer: ujq.withUpworkfreelancer.Clone(),
		withJob:              ujq.withJob.Clone(),
		withUser:             ujq.withUser.Clone(),
		// clone intermediate query.
		sql:  ujq.sql.Clone(),
		path: ujq.path,
	}
}

// WithUpworkfreelancer tells the query-builder to eager-load the nodes that are connected to
// the "upworkfreelancer" edge. The optional arguments are used to configure the query builder of the edge.
func (ujq *UpworkJobQuery) WithUpworkfreelancer(opts ...func(*UpworkFreelancerQuery)) *UpworkJobQuery {
	query := (&UpworkFreelancerClient{config: ujq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ujq.withUpworkfreelancer = query
	return ujq
}

// WithJob tells the query-builder to eager-load the nodes that are connected to
// the "job" edge. The optional arguments are used to configure the query builder of the edge.
func (ujq *UpworkJobQuery) WithJob(opts ...func(*JobQuery)) *UpworkJobQuery {
	query := (&JobClient{config: ujq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ujq.withJob = query
	return ujq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ujq *UpworkJobQuery) WithUser(opts ...func(*UserQuery)) *UpworkJobQuery {
	query := (&UserClient{config: ujq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ujq.withUser = query
	return ujq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UpworkJob.Query().
//		GroupBy(upworkjob.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ujq *UpworkJobQuery) GroupBy(field string, fields ...string) *UpworkJobGroupBy {
	ujq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UpworkJobGroupBy{build: ujq}
	grbuild.flds = &ujq.ctx.Fields
	grbuild.label = upworkjob.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.UpworkJob.Query().
//		Select(upworkjob.FieldTitle).
//		Scan(ctx, &v)
func (ujq *UpworkJobQuery) Select(fields ...string) *UpworkJobSelect {
	ujq.ctx.Fields = append(ujq.ctx.Fields, fields...)
	sbuild := &UpworkJobSelect{UpworkJobQuery: ujq}
	sbuild.label = upworkjob.Label
	sbuild.flds, sbuild.scan = &ujq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UpworkJobSelect configured with the given aggregations.
func (ujq *UpworkJobQuery) Aggregate(fns ...AggregateFunc) *UpworkJobSelect {
	return ujq.Select().Aggregate(fns...)
}

func (ujq *UpworkJobQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ujq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ujq); err != nil {
				return err
			}
		}
	}
	for _, f := range ujq.ctx.Fields {
		if !upworkjob.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ujq.path != nil {
		prev, err := ujq.path(ctx)
		if err != nil {
			return err
		}
		ujq.sql = prev
	}
	return nil
}

func (ujq *UpworkJobQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UpworkJob, error) {
	var (
		nodes       = []*UpworkJob{}
		withFKs     = ujq.withFKs
		_spec       = ujq.querySpec()
		loadedTypes = [3]bool{
			ujq.withUpworkfreelancer != nil,
			ujq.withJob != nil,
			ujq.withUser != nil,
		}
	)
	if ujq.withJob != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, upworkjob.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UpworkJob).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UpworkJob{config: ujq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ujq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ujq.withUpworkfreelancer; query != nil {
		if err := ujq.loadUpworkfreelancer(ctx, query, nodes,
			func(n *UpworkJob) { n.Edges.Upworkfreelancer = []*UpworkFreelancer{} },
			func(n *UpworkJob, e *UpworkFreelancer) {
				n.Edges.Upworkfreelancer = append(n.Edges.Upworkfreelancer, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := ujq.withJob; query != nil {
		if err := ujq.loadJob(ctx, query, nodes, nil,
			func(n *UpworkJob, e *Job) { n.Edges.Job = e }); err != nil {
			return nil, err
		}
	}
	if query := ujq.withUser; query != nil {
		if err := ujq.loadUser(ctx, query, nodes,
			func(n *UpworkJob) { n.Edges.User = []*User{} },
			func(n *UpworkJob, e *User) { n.Edges.User = append(n.Edges.User, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ujq *UpworkJobQuery) loadUpworkfreelancer(ctx context.Context, query *UpworkFreelancerQuery, nodes []*UpworkJob, init func(*UpworkJob), assign func(*UpworkJob, *UpworkFreelancer)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*UpworkJob)
	nids := make(map[string]map[*UpworkJob]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(upworkjob.UpworkfreelancerTable)
		s.Join(joinT).On(s.C(upworkfreelancer.FieldID), joinT.C(upworkjob.UpworkfreelancerPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(upworkjob.UpworkfreelancerPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(upworkjob.UpworkfreelancerPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*UpworkJob]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*UpworkFreelancer](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "upworkfreelancer" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (ujq *UpworkJobQuery) loadJob(ctx context.Context, query *JobQuery, nodes []*UpworkJob, init func(*UpworkJob), assign func(*UpworkJob, *Job)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UpworkJob)
	for i := range nodes {
		if nodes[i].job_upworkjob == nil {
			continue
		}
		fk := *nodes[i].job_upworkjob
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(job.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "job_upworkjob" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ujq *UpworkJobQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UpworkJob, init func(*UpworkJob), assign func(*UpworkJob, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*UpworkJob)
	nids := make(map[string]map[*UpworkJob]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(upworkjob.UserTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(upworkjob.UserPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(upworkjob.UserPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(upworkjob.UserPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*UpworkJob]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (ujq *UpworkJobQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ujq.querySpec()
	_spec.Node.Columns = ujq.ctx.Fields
	if len(ujq.ctx.Fields) > 0 {
		_spec.Unique = ujq.ctx.Unique != nil && *ujq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ujq.driver, _spec)
}

func (ujq *UpworkJobQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(upworkjob.Table, upworkjob.Columns, sqlgraph.NewFieldSpec(upworkjob.FieldID, field.TypeString))
	_spec.From = ujq.sql
	if unique := ujq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ujq.path != nil {
		_spec.Unique = true
	}
	if fields := ujq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upworkjob.FieldID)
		for i := range fields {
			if fields[i] != upworkjob.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ujq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ujq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ujq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ujq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ujq *UpworkJobQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ujq.driver.Dialect())
	t1 := builder.Table(upworkjob.Table)
	columns := ujq.ctx.Fields
	if len(columns) == 0 {
		columns = upworkjob.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ujq.sql != nil {
		selector = ujq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ujq.ctx.Unique != nil && *ujq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ujq.predicates {
		p(selector)
	}
	for _, p := range ujq.order {
		p(selector)
	}
	if offset := ujq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ujq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UpworkJobGroupBy is the group-by builder for UpworkJob entities.
type UpworkJobGroupBy struct {
	selector
	build *UpworkJobQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ujgb *UpworkJobGroupBy) Aggregate(fns ...AggregateFunc) *UpworkJobGroupBy {
	ujgb.fns = append(ujgb.fns, fns...)
	return ujgb
}

// Scan applies the selector query and scans the result into the given value.
func (ujgb *UpworkJobGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ujgb.build.ctx, "GroupBy")
	if err := ujgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpworkJobQuery, *UpworkJobGroupBy](ctx, ujgb.build, ujgb, ujgb.build.inters, v)
}

func (ujgb *UpworkJobGroupBy) sqlScan(ctx context.Context, root *UpworkJobQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ujgb.fns))
	for _, fn := range ujgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ujgb.flds)+len(ujgb.fns))
		for _, f := range *ujgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ujgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ujgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UpworkJobSelect is the builder for selecting fields of UpworkJob entities.
type UpworkJobSelect struct {
	*UpworkJobQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ujs *UpworkJobSelect) Aggregate(fns ...AggregateFunc) *UpworkJobSelect {
	ujs.fns = append(ujs.fns, fns...)
	return ujs
}

// Scan applies the selector query and scans the result into the given value.
func (ujs *UpworkJobSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ujs.ctx, "Select")
	if err := ujs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpworkJobQuery, *UpworkJobSelect](ctx, ujs.UpworkJobQuery, ujs, ujs.inters, v)
}

func (ujs *UpworkJobSelect) sqlScan(ctx context.Context, root *UpworkJobQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ujs.fns))
	for _, fn := range ujs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ujs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ujs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
