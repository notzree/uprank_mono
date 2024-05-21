// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttachmentRefsColumns holds the columns for the "attachment_refs" table.
	AttachmentRefsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "freelancer_attachments", Type: field.TypeUUID},
	}
	// AttachmentRefsTable holds the schema information for the "attachment_refs" table.
	AttachmentRefsTable = &schema.Table{
		Name:       "attachment_refs",
		Columns:    AttachmentRefsColumns,
		PrimaryKey: []*schema.Column{AttachmentRefsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attachment_refs_freelancers_attachments",
				Columns:    []*schema.Column{AttachmentRefsColumns[3]},
				RefColumns: []*schema.Column{FreelancersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FreelancersColumns holds the columns for the "freelancers" table.
	FreelancersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "country", Type: field.TypeString},
		{Name: "timezone", Type: field.TypeString},
		{Name: "cv", Type: field.TypeString, SchemaType: map[string]string{"postgres": "TEXT"}},
		{Name: "ai_reccomended", Type: field.TypeBool},
		{Name: "fixed_charge_amount", Type: field.TypeInt, Nullable: true},
		{Name: "fixed_charge_currency", Type: field.TypeString},
		{Name: "hourly_charge_amount", Type: field.TypeInt, Nullable: true},
		{Name: "hourly_charge_currency", Type: field.TypeString},
		{Name: "invited", Type: field.TypeBool},
		{Name: "photo_url", Type: field.TypeString},
		{Name: "recent_hours", Type: field.TypeInt},
		{Name: "total_hours", Type: field.TypeInt},
		{Name: "total_portfolio_items", Type: field.TypeInt},
		{Name: "total_portfolio_v2_items", Type: field.TypeInt},
		{Name: "upwork_total_feedback", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "upwork_recent_feedback", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "upwork_top_rated_status", Type: field.TypeBool},
		{Name: "upwork_top_rated_plus_status", Type: field.TypeBool},
		{Name: "upwork_sponsored", Type: field.TypeBool},
		{Name: "upwork_job_success_score", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "upwork_reccomended", Type: field.TypeBool},
		{Name: "skills", Type: field.TypeJSON},
		{Name: "average_recent_earnings", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "combined_average_recent_earnings", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "combined_recent_earnings", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "combined_total_earnings", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "combined_total_revenue", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "recent_earnings", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "total_revenue", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "uprank_score", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "uprank_updated_at", Type: field.TypeTime},
		{Name: "uprank_reccomended", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "uprank_reccomended_reasons", Type: field.TypeString, Nullable: true},
		{Name: "uprank_not_enough_data", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "job_freelancers", Type: field.TypeString},
	}
	// FreelancersTable holds the schema information for the "freelancers" table.
	FreelancersTable = &schema.Table{
		Name:       "freelancers",
		Columns:    FreelancersColumns,
		PrimaryKey: []*schema.Column{FreelancersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "freelancers_jobs_freelancers",
				Columns:    []*schema.Column{FreelancersColumns[40]},
				RefColumns: []*schema.Column{JobsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// JobsColumns holds the columns for the "jobs" table.
	JobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "location", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString},
		{Name: "skills", Type: field.TypeJSON, Nullable: true},
		{Name: "experience_level", Type: field.TypeString, Nullable: true},
		{Name: "hourly", Type: field.TypeBool},
		{Name: "fixed", Type: field.TypeBool},
		{Name: "hourly_rate", Type: field.TypeJSON, Nullable: true},
		{Name: "fixed_rate", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "average_uprank_score", Type: field.TypeFloat64, Nullable: true},
		{Name: "max_uprank_score", Type: field.TypeFloat64, Nullable: true},
		{Name: "min_uprank_score", Type: field.TypeFloat64, Nullable: true},
		{Name: "user_jobs", Type: field.TypeString},
	}
	// JobsTable holds the schema information for the "jobs" table.
	JobsTable = &schema.Table{
		Name:       "jobs",
		Columns:    JobsColumns,
		PrimaryKey: []*schema.Column{JobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "jobs_users_jobs",
				Columns:    []*schema.Column{JobsColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "first_name", Type: field.TypeString, Default: "unknown"},
		{Name: "company_name", Type: field.TypeString, Default: "unknown"},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_login", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WorkHistoriesColumns holds the columns for the "work_histories" table.
	WorkHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "client_feedback", Type: field.TypeString},
		{Name: "overall_rating", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "fixed_charge_amount", Type: field.TypeInt, Nullable: true},
		{Name: "fixed_charge_currency", Type: field.TypeString, Nullable: true},
		{Name: "hourly_charge_amount", Type: field.TypeInt, Nullable: true},
		{Name: "hourly_charge_currency", Type: field.TypeString, Nullable: true},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "job_description", Type: field.TypeString},
		{Name: "total_proposals", Type: field.TypeInt},
		{Name: "number_of_interviews", Type: field.TypeInt},
		{Name: "skills", Type: field.TypeJSON},
		{Name: "client_rating", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "client_review_count", Type: field.TypeInt},
		{Name: "client_country", Type: field.TypeString},
		{Name: "client_total_jobs_posted", Type: field.TypeInt},
		{Name: "client_total_spend", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "client_total_hires", Type: field.TypeInt, Nullable: true},
		{Name: "client_total_paid_hours", Type: field.TypeInt, Nullable: true},
		{Name: "client_average_hourly_rate_paid", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "client_company_category", Type: field.TypeString, Nullable: true},
		{Name: "client_company_size", Type: field.TypeString, Nullable: true},
		{Name: "freelancer_work_histories", Type: field.TypeUUID},
	}
	// WorkHistoriesTable holds the schema information for the "work_histories" table.
	WorkHistoriesTable = &schema.Table{
		Name:       "work_histories",
		Columns:    WorkHistoriesColumns,
		PrimaryKey: []*schema.Column{WorkHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "work_histories_freelancers_work_histories",
				Columns:    []*schema.Column{WorkHistoriesColumns[24]},
				RefColumns: []*schema.Column{FreelancersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttachmentRefsTable,
		FreelancersTable,
		JobsTable,
		UsersTable,
		WorkHistoriesTable,
	}
)

func init() {
	AttachmentRefsTable.ForeignKeys[0].RefTable = FreelancersTable
	FreelancersTable.ForeignKeys[0].RefTable = JobsTable
	JobsTable.ForeignKeys[0].RefTable = UsersTable
	WorkHistoriesTable.ForeignKeys[0].RefTable = FreelancersTable
}