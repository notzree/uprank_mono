// Code generated by ent, DO NOT EDIT.

package job

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/notzree/uprank-backend/main-backend/ent/schema"
)

const (
	// Label holds the string label denoting the job type in the database.
	Label = "job"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldOriginPlatform holds the string denoting the origin_platform field in the database.
	FieldOriginPlatform = "origin_platform"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeUpworkjob holds the string denoting the upworkjob edge name in mutations.
	EdgeUpworkjob = "upworkjob"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// UpworkJobFieldID holds the string denoting the ID field of the UpworkJob.
	UpworkJobFieldID = "id"
	// Table holds the table name of the job in the database.
	Table = "jobs"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "jobs"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_job"
	// UpworkjobTable is the table that holds the upworkjob relation/edge.
	UpworkjobTable = "upwork_jobs"
	// UpworkjobInverseTable is the table name for the UpworkJob entity.
	// It exists in this package in order to avoid circular dependency with the "upworkjob" package.
	UpworkjobInverseTable = "upwork_jobs"
	// UpworkjobColumn is the table column denoting the upworkjob relation/edge.
	UpworkjobColumn = "job_upworkjob"
)

// Columns holds all SQL columns for job fields.
var Columns = []string{
	FieldID,
	FieldOriginPlatform,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "jobs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_job",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OriginPlatformValidator is a validator for the "origin_platform" field enum values. It is called by the builders before save.
func OriginPlatformValidator(op schema.Platform) error {
	switch op {
	case "upwork", "uprank":
		return nil
	default:
		return fmt.Errorf("job: invalid enum value for origin_platform field: %q", op)
	}
}

// OrderOption defines the ordering options for the Job queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOriginPlatform orders the results by the origin_platform field.
func ByOriginPlatform(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOriginPlatform, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByUpworkjobCount orders the results by upworkjob count.
func ByUpworkjobCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUpworkjobStep(), opts...)
	}
}

// ByUpworkjob orders the results by upworkjob terms.
func ByUpworkjob(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUpworkjobStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newUpworkjobStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UpworkjobInverseTable, UpworkJobFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UpworkjobTable, UpworkjobColumn),
	)
}
