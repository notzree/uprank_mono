// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/attachmentref"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/predicate"
	"github.com/notzree/uprank_mono/uprank-backend/main-backend/ent/upworkfreelancer"
)

// AttachmentRefQuery is the builder for querying AttachmentRef entities.
type AttachmentRefQuery struct {
	config
	ctx            *QueryContext
	order          []attachmentref.OrderOption
	inters         []Interceptor
	predicates     []predicate.AttachmentRef
	withFreelancer *UpworkFreelancerQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttachmentRefQuery builder.
func (arq *AttachmentRefQuery) Where(ps ...predicate.AttachmentRef) *AttachmentRefQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit the number of records to be returned by this query.
func (arq *AttachmentRefQuery) Limit(limit int) *AttachmentRefQuery {
	arq.ctx.Limit = &limit
	return arq
}

// Offset to start from.
func (arq *AttachmentRefQuery) Offset(offset int) *AttachmentRefQuery {
	arq.ctx.Offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AttachmentRefQuery) Unique(unique bool) *AttachmentRefQuery {
	arq.ctx.Unique = &unique
	return arq
}

// Order specifies how the records should be ordered.
func (arq *AttachmentRefQuery) Order(o ...attachmentref.OrderOption) *AttachmentRefQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// QueryFreelancer chains the current query on the "freelancer" edge.
func (arq *AttachmentRefQuery) QueryFreelancer() *UpworkFreelancerQuery {
	query := (&UpworkFreelancerClient{config: arq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := arq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attachmentref.Table, attachmentref.FieldID, selector),
			sqlgraph.To(upworkfreelancer.Table, upworkfreelancer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachmentref.FreelancerTable, attachmentref.FreelancerColumn),
		)
		fromU = sqlgraph.SetNeighbors(arq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AttachmentRef entity from the query.
// Returns a *NotFoundError when no AttachmentRef was found.
func (arq *AttachmentRefQuery) First(ctx context.Context) (*AttachmentRef, error) {
	nodes, err := arq.Limit(1).All(setContextOp(ctx, arq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attachmentref.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AttachmentRefQuery) FirstX(ctx context.Context) *AttachmentRef {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AttachmentRef ID from the query.
// Returns a *NotFoundError when no AttachmentRef ID was found.
func (arq *AttachmentRefQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(1).IDs(setContextOp(ctx, arq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attachmentref.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AttachmentRefQuery) FirstIDX(ctx context.Context) int {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AttachmentRef entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AttachmentRef entity is found.
// Returns a *NotFoundError when no AttachmentRef entities are found.
func (arq *AttachmentRefQuery) Only(ctx context.Context) (*AttachmentRef, error) {
	nodes, err := arq.Limit(2).All(setContextOp(ctx, arq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attachmentref.Label}
	default:
		return nil, &NotSingularError{attachmentref.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AttachmentRefQuery) OnlyX(ctx context.Context) *AttachmentRef {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AttachmentRef ID in the query.
// Returns a *NotSingularError when more than one AttachmentRef ID is found.
// Returns a *NotFoundError when no entities are found.
func (arq *AttachmentRefQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(2).IDs(setContextOp(ctx, arq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attachmentref.Label}
	default:
		err = &NotSingularError{attachmentref.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AttachmentRefQuery) OnlyIDX(ctx context.Context) int {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AttachmentRefs.
func (arq *AttachmentRefQuery) All(ctx context.Context) ([]*AttachmentRef, error) {
	ctx = setContextOp(ctx, arq.ctx, "All")
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AttachmentRef, *AttachmentRefQuery]()
	return withInterceptors[[]*AttachmentRef](ctx, arq, qr, arq.inters)
}

// AllX is like All, but panics if an error occurs.
func (arq *AttachmentRefQuery) AllX(ctx context.Context) []*AttachmentRef {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AttachmentRef IDs.
func (arq *AttachmentRefQuery) IDs(ctx context.Context) (ids []int, err error) {
	if arq.ctx.Unique == nil && arq.path != nil {
		arq.Unique(true)
	}
	ctx = setContextOp(ctx, arq.ctx, "IDs")
	if err = arq.Select(attachmentref.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AttachmentRefQuery) IDsX(ctx context.Context) []int {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AttachmentRefQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, arq.ctx, "Count")
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, arq, querierCount[*AttachmentRefQuery](), arq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AttachmentRefQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AttachmentRefQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, arq.ctx, "Exist")
	switch _, err := arq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AttachmentRefQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttachmentRefQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AttachmentRefQuery) Clone() *AttachmentRefQuery {
	if arq == nil {
		return nil
	}
	return &AttachmentRefQuery{
		config:         arq.config,
		ctx:            arq.ctx.Clone(),
		order:          append([]attachmentref.OrderOption{}, arq.order...),
		inters:         append([]Interceptor{}, arq.inters...),
		predicates:     append([]predicate.AttachmentRef{}, arq.predicates...),
		withFreelancer: arq.withFreelancer.Clone(),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

// WithFreelancer tells the query-builder to eager-load the nodes that are connected to
// the "freelancer" edge. The optional arguments are used to configure the query builder of the edge.
func (arq *AttachmentRefQuery) WithFreelancer(opts ...func(*UpworkFreelancerQuery)) *AttachmentRefQuery {
	query := (&UpworkFreelancerClient{config: arq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	arq.withFreelancer = query
	return arq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AttachmentRef.Query().
//		GroupBy(attachmentref.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (arq *AttachmentRefQuery) GroupBy(field string, fields ...string) *AttachmentRefGroupBy {
	arq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttachmentRefGroupBy{build: arq}
	grbuild.flds = &arq.ctx.Fields
	grbuild.label = attachmentref.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.AttachmentRef.Query().
//		Select(attachmentref.FieldName).
//		Scan(ctx, &v)
func (arq *AttachmentRefQuery) Select(fields ...string) *AttachmentRefSelect {
	arq.ctx.Fields = append(arq.ctx.Fields, fields...)
	sbuild := &AttachmentRefSelect{AttachmentRefQuery: arq}
	sbuild.label = attachmentref.Label
	sbuild.flds, sbuild.scan = &arq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttachmentRefSelect configured with the given aggregations.
func (arq *AttachmentRefQuery) Aggregate(fns ...AggregateFunc) *AttachmentRefSelect {
	return arq.Select().Aggregate(fns...)
}

func (arq *AttachmentRefQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range arq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, arq); err != nil {
				return err
			}
		}
	}
	for _, f := range arq.ctx.Fields {
		if !attachmentref.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AttachmentRefQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AttachmentRef, error) {
	var (
		nodes       = []*AttachmentRef{}
		withFKs     = arq.withFKs
		_spec       = arq.querySpec()
		loadedTypes = [1]bool{
			arq.withFreelancer != nil,
		}
	)
	if arq.withFreelancer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, attachmentref.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AttachmentRef).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AttachmentRef{config: arq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := arq.withFreelancer; query != nil {
		if err := arq.loadFreelancer(ctx, query, nodes, nil,
			func(n *AttachmentRef, e *UpworkFreelancer) { n.Edges.Freelancer = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (arq *AttachmentRefQuery) loadFreelancer(ctx context.Context, query *UpworkFreelancerQuery, nodes []*AttachmentRef, init func(*AttachmentRef), assign func(*AttachmentRef, *UpworkFreelancer)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*AttachmentRef)
	for i := range nodes {
		if nodes[i].upwork_freelancer_attachments == nil {
			continue
		}
		fk := *nodes[i].upwork_freelancer_attachments
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(upworkfreelancer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "upwork_freelancer_attachments" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (arq *AttachmentRefQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	_spec.Node.Columns = arq.ctx.Fields
	if len(arq.ctx.Fields) > 0 {
		_spec.Unique = arq.ctx.Unique != nil && *arq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AttachmentRefQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attachmentref.Table, attachmentref.Columns, sqlgraph.NewFieldSpec(attachmentref.FieldID, field.TypeInt))
	_spec.From = arq.sql
	if unique := arq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if arq.path != nil {
		_spec.Unique = true
	}
	if fields := arq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attachmentref.FieldID)
		for i := range fields {
			if fields[i] != attachmentref.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AttachmentRefQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(attachmentref.Table)
	columns := arq.ctx.Fields
	if len(columns) == 0 {
		columns = attachmentref.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if arq.ctx.Unique != nil && *arq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttachmentRefGroupBy is the group-by builder for AttachmentRef entities.
type AttachmentRefGroupBy struct {
	selector
	build *AttachmentRefQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AttachmentRefGroupBy) Aggregate(fns ...AggregateFunc) *AttachmentRefGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the selector query and scans the result into the given value.
func (argb *AttachmentRefGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, argb.build.ctx, "GroupBy")
	if err := argb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttachmentRefQuery, *AttachmentRefGroupBy](ctx, argb.build, argb, argb.build.inters, v)
}

func (argb *AttachmentRefGroupBy) sqlScan(ctx context.Context, root *AttachmentRefQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*argb.flds)+len(argb.fns))
		for _, f := range *argb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*argb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttachmentRefSelect is the builder for selecting fields of AttachmentRef entities.
type AttachmentRefSelect struct {
	*AttachmentRefQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ars *AttachmentRefSelect) Aggregate(fns ...AggregateFunc) *AttachmentRefSelect {
	ars.fns = append(ars.fns, fns...)
	return ars
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AttachmentRefSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ars.ctx, "Select")
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttachmentRefQuery, *AttachmentRefSelect](ctx, ars.AttachmentRefQuery, ars, ars.inters, v)
}

func (ars *AttachmentRefSelect) sqlScan(ctx context.Context, root *AttachmentRefQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ars.fns))
	for _, fn := range ars.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ars.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
