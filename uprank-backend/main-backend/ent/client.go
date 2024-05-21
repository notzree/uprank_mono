// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/main-backend/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/notzree/uprank-backend/main-backend/ent/attachmentref"
	"github.com/notzree/uprank-backend/main-backend/ent/freelancer"
	"github.com/notzree/uprank-backend/main-backend/ent/job"
	"github.com/notzree/uprank-backend/main-backend/ent/user"
	"github.com/notzree/uprank-backend/main-backend/ent/workhistory"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AttachmentRef is the client for interacting with the AttachmentRef builders.
	AttachmentRef *AttachmentRefClient
	// Freelancer is the client for interacting with the Freelancer builders.
	Freelancer *FreelancerClient
	// Job is the client for interacting with the Job builders.
	Job *JobClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// WorkHistory is the client for interacting with the WorkHistory builders.
	WorkHistory *WorkHistoryClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AttachmentRef = NewAttachmentRefClient(c.config)
	c.Freelancer = NewFreelancerClient(c.config)
	c.Job = NewJobClient(c.config)
	c.User = NewUserClient(c.config)
	c.WorkHistory = NewWorkHistoryClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		AttachmentRef: NewAttachmentRefClient(cfg),
		Freelancer:    NewFreelancerClient(cfg),
		Job:           NewJobClient(cfg),
		User:          NewUserClient(cfg),
		WorkHistory:   NewWorkHistoryClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		AttachmentRef: NewAttachmentRefClient(cfg),
		Freelancer:    NewFreelancerClient(cfg),
		Job:           NewJobClient(cfg),
		User:          NewUserClient(cfg),
		WorkHistory:   NewWorkHistoryClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AttachmentRef.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AttachmentRef.Use(hooks...)
	c.Freelancer.Use(hooks...)
	c.Job.Use(hooks...)
	c.User.Use(hooks...)
	c.WorkHistory.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.AttachmentRef.Intercept(interceptors...)
	c.Freelancer.Intercept(interceptors...)
	c.Job.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
	c.WorkHistory.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AttachmentRefMutation:
		return c.AttachmentRef.mutate(ctx, m)
	case *FreelancerMutation:
		return c.Freelancer.mutate(ctx, m)
	case *JobMutation:
		return c.Job.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *WorkHistoryMutation:
		return c.WorkHistory.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AttachmentRefClient is a client for the AttachmentRef schema.
type AttachmentRefClient struct {
	config
}

// NewAttachmentRefClient returns a client for the AttachmentRef from the given config.
func NewAttachmentRefClient(c config) *AttachmentRefClient {
	return &AttachmentRefClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `attachmentref.Hooks(f(g(h())))`.
func (c *AttachmentRefClient) Use(hooks ...Hook) {
	c.hooks.AttachmentRef = append(c.hooks.AttachmentRef, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `attachmentref.Intercept(f(g(h())))`.
func (c *AttachmentRefClient) Intercept(interceptors ...Interceptor) {
	c.inters.AttachmentRef = append(c.inters.AttachmentRef, interceptors...)
}

// Create returns a builder for creating a AttachmentRef entity.
func (c *AttachmentRefClient) Create() *AttachmentRefCreate {
	mutation := newAttachmentRefMutation(c.config, OpCreate)
	return &AttachmentRefCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AttachmentRef entities.
func (c *AttachmentRefClient) CreateBulk(builders ...*AttachmentRefCreate) *AttachmentRefCreateBulk {
	return &AttachmentRefCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AttachmentRefClient) MapCreateBulk(slice any, setFunc func(*AttachmentRefCreate, int)) *AttachmentRefCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AttachmentRefCreateBulk{err: fmt.Errorf("calling to AttachmentRefClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AttachmentRefCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AttachmentRefCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AttachmentRef.
func (c *AttachmentRefClient) Update() *AttachmentRefUpdate {
	mutation := newAttachmentRefMutation(c.config, OpUpdate)
	return &AttachmentRefUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AttachmentRefClient) UpdateOne(ar *AttachmentRef) *AttachmentRefUpdateOne {
	mutation := newAttachmentRefMutation(c.config, OpUpdateOne, withAttachmentRef(ar))
	return &AttachmentRefUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AttachmentRefClient) UpdateOneID(id int) *AttachmentRefUpdateOne {
	mutation := newAttachmentRefMutation(c.config, OpUpdateOne, withAttachmentRefID(id))
	return &AttachmentRefUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AttachmentRef.
func (c *AttachmentRefClient) Delete() *AttachmentRefDelete {
	mutation := newAttachmentRefMutation(c.config, OpDelete)
	return &AttachmentRefDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AttachmentRefClient) DeleteOne(ar *AttachmentRef) *AttachmentRefDeleteOne {
	return c.DeleteOneID(ar.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AttachmentRefClient) DeleteOneID(id int) *AttachmentRefDeleteOne {
	builder := c.Delete().Where(attachmentref.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AttachmentRefDeleteOne{builder}
}

// Query returns a query builder for AttachmentRef.
func (c *AttachmentRefClient) Query() *AttachmentRefQuery {
	return &AttachmentRefQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAttachmentRef},
		inters: c.Interceptors(),
	}
}

// Get returns a AttachmentRef entity by its id.
func (c *AttachmentRefClient) Get(ctx context.Context, id int) (*AttachmentRef, error) {
	return c.Query().Where(attachmentref.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AttachmentRefClient) GetX(ctx context.Context, id int) *AttachmentRef {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFreelancer queries the freelancer edge of a AttachmentRef.
func (c *AttachmentRefClient) QueryFreelancer(ar *AttachmentRef) *FreelancerQuery {
	query := (&FreelancerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ar.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(attachmentref.Table, attachmentref.FieldID, id),
			sqlgraph.To(freelancer.Table, freelancer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachmentref.FreelancerTable, attachmentref.FreelancerColumn),
		)
		fromV = sqlgraph.Neighbors(ar.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AttachmentRefClient) Hooks() []Hook {
	return c.hooks.AttachmentRef
}

// Interceptors returns the client interceptors.
func (c *AttachmentRefClient) Interceptors() []Interceptor {
	return c.inters.AttachmentRef
}

func (c *AttachmentRefClient) mutate(ctx context.Context, m *AttachmentRefMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AttachmentRefCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AttachmentRefUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AttachmentRefUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AttachmentRefDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown AttachmentRef mutation op: %q", m.Op())
	}
}

// FreelancerClient is a client for the Freelancer schema.
type FreelancerClient struct {
	config
}

// NewFreelancerClient returns a client for the Freelancer from the given config.
func NewFreelancerClient(c config) *FreelancerClient {
	return &FreelancerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `freelancer.Hooks(f(g(h())))`.
func (c *FreelancerClient) Use(hooks ...Hook) {
	c.hooks.Freelancer = append(c.hooks.Freelancer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `freelancer.Intercept(f(g(h())))`.
func (c *FreelancerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Freelancer = append(c.inters.Freelancer, interceptors...)
}

// Create returns a builder for creating a Freelancer entity.
func (c *FreelancerClient) Create() *FreelancerCreate {
	mutation := newFreelancerMutation(c.config, OpCreate)
	return &FreelancerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Freelancer entities.
func (c *FreelancerClient) CreateBulk(builders ...*FreelancerCreate) *FreelancerCreateBulk {
	return &FreelancerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *FreelancerClient) MapCreateBulk(slice any, setFunc func(*FreelancerCreate, int)) *FreelancerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &FreelancerCreateBulk{err: fmt.Errorf("calling to FreelancerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*FreelancerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &FreelancerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Freelancer.
func (c *FreelancerClient) Update() *FreelancerUpdate {
	mutation := newFreelancerMutation(c.config, OpUpdate)
	return &FreelancerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FreelancerClient) UpdateOne(f *Freelancer) *FreelancerUpdateOne {
	mutation := newFreelancerMutation(c.config, OpUpdateOne, withFreelancer(f))
	return &FreelancerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FreelancerClient) UpdateOneID(id uuid.UUID) *FreelancerUpdateOne {
	mutation := newFreelancerMutation(c.config, OpUpdateOne, withFreelancerID(id))
	return &FreelancerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Freelancer.
func (c *FreelancerClient) Delete() *FreelancerDelete {
	mutation := newFreelancerMutation(c.config, OpDelete)
	return &FreelancerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FreelancerClient) DeleteOne(f *Freelancer) *FreelancerDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FreelancerClient) DeleteOneID(id uuid.UUID) *FreelancerDeleteOne {
	builder := c.Delete().Where(freelancer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FreelancerDeleteOne{builder}
}

// Query returns a query builder for Freelancer.
func (c *FreelancerClient) Query() *FreelancerQuery {
	return &FreelancerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFreelancer},
		inters: c.Interceptors(),
	}
}

// Get returns a Freelancer entity by its id.
func (c *FreelancerClient) Get(ctx context.Context, id uuid.UUID) (*Freelancer, error) {
	return c.Query().Where(freelancer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FreelancerClient) GetX(ctx context.Context, id uuid.UUID) *Freelancer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryJob queries the job edge of a Freelancer.
func (c *FreelancerClient) QueryJob(f *Freelancer) *JobQuery {
	query := (&JobClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(freelancer.Table, freelancer.FieldID, id),
			sqlgraph.To(job.Table, job.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, freelancer.JobTable, freelancer.JobColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAttachments queries the attachments edge of a Freelancer.
func (c *FreelancerClient) QueryAttachments(f *Freelancer) *AttachmentRefQuery {
	query := (&AttachmentRefClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(freelancer.Table, freelancer.FieldID, id),
			sqlgraph.To(attachmentref.Table, attachmentref.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, freelancer.AttachmentsTable, freelancer.AttachmentsColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWorkHistories queries the work_histories edge of a Freelancer.
func (c *FreelancerClient) QueryWorkHistories(f *Freelancer) *WorkHistoryQuery {
	query := (&WorkHistoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(freelancer.Table, freelancer.FieldID, id),
			sqlgraph.To(workhistory.Table, workhistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, freelancer.WorkHistoriesTable, freelancer.WorkHistoriesColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FreelancerClient) Hooks() []Hook {
	return c.hooks.Freelancer
}

// Interceptors returns the client interceptors.
func (c *FreelancerClient) Interceptors() []Interceptor {
	return c.inters.Freelancer
}

func (c *FreelancerClient) mutate(ctx context.Context, m *FreelancerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FreelancerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FreelancerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FreelancerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FreelancerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Freelancer mutation op: %q", m.Op())
	}
}

// JobClient is a client for the Job schema.
type JobClient struct {
	config
}

// NewJobClient returns a client for the Job from the given config.
func NewJobClient(c config) *JobClient {
	return &JobClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `job.Hooks(f(g(h())))`.
func (c *JobClient) Use(hooks ...Hook) {
	c.hooks.Job = append(c.hooks.Job, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `job.Intercept(f(g(h())))`.
func (c *JobClient) Intercept(interceptors ...Interceptor) {
	c.inters.Job = append(c.inters.Job, interceptors...)
}

// Create returns a builder for creating a Job entity.
func (c *JobClient) Create() *JobCreate {
	mutation := newJobMutation(c.config, OpCreate)
	return &JobCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Job entities.
func (c *JobClient) CreateBulk(builders ...*JobCreate) *JobCreateBulk {
	return &JobCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *JobClient) MapCreateBulk(slice any, setFunc func(*JobCreate, int)) *JobCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &JobCreateBulk{err: fmt.Errorf("calling to JobClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*JobCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &JobCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Job.
func (c *JobClient) Update() *JobUpdate {
	mutation := newJobMutation(c.config, OpUpdate)
	return &JobUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *JobClient) UpdateOne(j *Job) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJob(j))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *JobClient) UpdateOneID(id string) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJobID(id))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Job.
func (c *JobClient) Delete() *JobDelete {
	mutation := newJobMutation(c.config, OpDelete)
	return &JobDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *JobClient) DeleteOne(j *Job) *JobDeleteOne {
	return c.DeleteOneID(j.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *JobClient) DeleteOneID(id string) *JobDeleteOne {
	builder := c.Delete().Where(job.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &JobDeleteOne{builder}
}

// Query returns a query builder for Job.
func (c *JobClient) Query() *JobQuery {
	return &JobQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeJob},
		inters: c.Interceptors(),
	}
}

// Get returns a Job entity by its id.
func (c *JobClient) Get(ctx context.Context, id string) (*Job, error) {
	return c.Query().Where(job.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *JobClient) GetX(ctx context.Context, id string) *Job {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Job.
func (c *JobClient) QueryUser(j *Job) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := j.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, job.UserTable, job.UserColumn),
		)
		fromV = sqlgraph.Neighbors(j.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFreelancers queries the freelancers edge of a Job.
func (c *JobClient) QueryFreelancers(j *Job) *FreelancerQuery {
	query := (&FreelancerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := j.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, id),
			sqlgraph.To(freelancer.Table, freelancer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, job.FreelancersTable, job.FreelancersColumn),
		)
		fromV = sqlgraph.Neighbors(j.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *JobClient) Hooks() []Hook {
	return c.hooks.Job
}

// Interceptors returns the client interceptors.
func (c *JobClient) Interceptors() []Interceptor {
	return c.inters.Job
}

func (c *JobClient) mutate(ctx context.Context, m *JobMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&JobCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&JobUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&JobDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Job mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryJobs queries the jobs edge of a User.
func (c *UserClient) QueryJobs(u *User) *JobQuery {
	query := (&JobClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(job.Table, job.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.JobsTable, user.JobsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// WorkHistoryClient is a client for the WorkHistory schema.
type WorkHistoryClient struct {
	config
}

// NewWorkHistoryClient returns a client for the WorkHistory from the given config.
func NewWorkHistoryClient(c config) *WorkHistoryClient {
	return &WorkHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `workhistory.Hooks(f(g(h())))`.
func (c *WorkHistoryClient) Use(hooks ...Hook) {
	c.hooks.WorkHistory = append(c.hooks.WorkHistory, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `workhistory.Intercept(f(g(h())))`.
func (c *WorkHistoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.WorkHistory = append(c.inters.WorkHistory, interceptors...)
}

// Create returns a builder for creating a WorkHistory entity.
func (c *WorkHistoryClient) Create() *WorkHistoryCreate {
	mutation := newWorkHistoryMutation(c.config, OpCreate)
	return &WorkHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WorkHistory entities.
func (c *WorkHistoryClient) CreateBulk(builders ...*WorkHistoryCreate) *WorkHistoryCreateBulk {
	return &WorkHistoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *WorkHistoryClient) MapCreateBulk(slice any, setFunc func(*WorkHistoryCreate, int)) *WorkHistoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &WorkHistoryCreateBulk{err: fmt.Errorf("calling to WorkHistoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*WorkHistoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &WorkHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WorkHistory.
func (c *WorkHistoryClient) Update() *WorkHistoryUpdate {
	mutation := newWorkHistoryMutation(c.config, OpUpdate)
	return &WorkHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkHistoryClient) UpdateOne(wh *WorkHistory) *WorkHistoryUpdateOne {
	mutation := newWorkHistoryMutation(c.config, OpUpdateOne, withWorkHistory(wh))
	return &WorkHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkHistoryClient) UpdateOneID(id int) *WorkHistoryUpdateOne {
	mutation := newWorkHistoryMutation(c.config, OpUpdateOne, withWorkHistoryID(id))
	return &WorkHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WorkHistory.
func (c *WorkHistoryClient) Delete() *WorkHistoryDelete {
	mutation := newWorkHistoryMutation(c.config, OpDelete)
	return &WorkHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WorkHistoryClient) DeleteOne(wh *WorkHistory) *WorkHistoryDeleteOne {
	return c.DeleteOneID(wh.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *WorkHistoryClient) DeleteOneID(id int) *WorkHistoryDeleteOne {
	builder := c.Delete().Where(workhistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkHistoryDeleteOne{builder}
}

// Query returns a query builder for WorkHistory.
func (c *WorkHistoryClient) Query() *WorkHistoryQuery {
	return &WorkHistoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeWorkHistory},
		inters: c.Interceptors(),
	}
}

// Get returns a WorkHistory entity by its id.
func (c *WorkHistoryClient) Get(ctx context.Context, id int) (*WorkHistory, error) {
	return c.Query().Where(workhistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkHistoryClient) GetX(ctx context.Context, id int) *WorkHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUpworkFreelancerProposal queries the upwork_Freelancer_Proposal edge of a WorkHistory.
func (c *WorkHistoryClient) QueryUpworkFreelancerProposal(wh *WorkHistory) *FreelancerQuery {
	query := (&FreelancerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := wh.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workhistory.Table, workhistory.FieldID, id),
			sqlgraph.To(freelancer.Table, freelancer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workhistory.UpworkFreelancerProposalTable, workhistory.UpworkFreelancerProposalColumn),
		)
		fromV = sqlgraph.Neighbors(wh.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WorkHistoryClient) Hooks() []Hook {
	return c.hooks.WorkHistory
}

// Interceptors returns the client interceptors.
func (c *WorkHistoryClient) Interceptors() []Interceptor {
	return c.inters.WorkHistory
}

func (c *WorkHistoryClient) mutate(ctx context.Context, m *WorkHistoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&WorkHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&WorkHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&WorkHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&WorkHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown WorkHistory mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		AttachmentRef, Freelancer, Job, User, WorkHistory []ent.Hook
	}
	inters struct {
		AttachmentRef, Freelancer, Job, User, WorkHistory []ent.Interceptor
	}
)