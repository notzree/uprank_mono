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
		{Name: "link", Type: field.TypeString},
		{Name: "upwork_freelancer_attachments", Type: field.TypeString},
	}
	// AttachmentRefsTable holds the schema information for the "attachment_refs" table.
	AttachmentRefsTable = &schema.Table{
		Name:       "attachment_refs",
		Columns:    AttachmentRefsColumns,
		PrimaryKey: []*schema.Column{AttachmentRefsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attachment_refs_upwork_freelancers_attachments",
				Columns:    []*schema.Column{AttachmentRefsColumns[3]},
				RefColumns: []*schema.Column{UpworkFreelancersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FreelancerInferenceDataColumns holds the columns for the "freelancer_inference_data" table.
	FreelancerInferenceDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uprank_reccomended", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "uprank_reccomended_reasons", Type: field.TypeString, Nullable: true},
		{Name: "uprank_not_enough_data", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "finalized_rating_score", Type: field.TypeFloat64},
		{Name: "raw_rating_score", Type: field.TypeFloat64, Nullable: true},
		{Name: "ai_estimated_duration", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"postgres": "INTERVAL"}},
		{Name: "budget_adherence_percentage", Type: field.TypeFloat64, Nullable: true},
		{Name: "upwork_freelancer_freelancer_inference_data", Type: field.TypeString, Unique: true},
	}
	// FreelancerInferenceDataTable holds the schema information for the "freelancer_inference_data" table.
	FreelancerInferenceDataTable = &schema.Table{
		Name:       "freelancer_inference_data",
		Columns:    FreelancerInferenceDataColumns,
		PrimaryKey: []*schema.Column{FreelancerInferenceDataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "freelancer_inference_data_upwork_freelancers_freelancer_inference_data",
				Columns:    []*schema.Column{FreelancerInferenceDataColumns[8]},
				RefColumns: []*schema.Column{UpworkFreelancersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// JobsColumns holds the columns for the "jobs" table.
	JobsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "origin_platform", Type: field.TypeEnum, Enums: []string{"upwork", "uprank"}},
		{Name: "user_job", Type: field.TypeString},
	}
	// JobsTable holds the schema information for the "jobs" table.
	JobsTable = &schema.Table{
		Name:       "jobs",
		Columns:    JobsColumns,
		PrimaryKey: []*schema.Column{JobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "jobs_users_job",
				Columns:    []*schema.Column{JobsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UpworkFreelancersColumns holds the columns for the "upwork_freelancers" table.
	UpworkFreelancersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "country", Type: field.TypeString},
		{Name: "timezone", Type: field.TypeString},
		{Name: "cv", Type: field.TypeString, SchemaType: map[string]string{"postgres": "TEXT"}},
		{Name: "ai_reccomended", Type: field.TypeBool},
		{Name: "fixed_charge_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "fixed_charge_currency", Type: field.TypeString},
		{Name: "hourly_charge_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "hourly_charge_currency", Type: field.TypeString},
		{Name: "invited", Type: field.TypeBool},
		{Name: "photo_url", Type: field.TypeString},
		{Name: "recent_hours", Type: field.TypeFloat64},
		{Name: "total_hours", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "embedded_at", Type: field.TypeTime, Nullable: true},
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
		{Name: "missing_fields", Type: field.TypeBool, Default: false},
	}
	// UpworkFreelancersTable holds the schema information for the "upwork_freelancers" table.
	UpworkFreelancersTable = &schema.Table{
		Name:       "upwork_freelancers",
		Columns:    UpworkFreelancersColumns,
		PrimaryKey: []*schema.Column{UpworkFreelancersColumns[0]},
	}
	// UpworkJobsColumns holds the columns for the "upwork_jobs" table.
	UpworkJobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "embedded_at", Type: field.TypeTime, Nullable: true},
		{Name: "ranked_at", Type: field.TypeTime, Nullable: true},
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
		{Name: "job_upworkjob", Type: field.TypeUUID, Unique: true},
	}
	// UpworkJobsTable holds the schema information for the "upwork_jobs" table.
	UpworkJobsTable = &schema.Table{
		Name:       "upwork_jobs",
		Columns:    UpworkJobsColumns,
		PrimaryKey: []*schema.Column{UpworkJobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "upwork_jobs_jobs_upworkjob",
				Columns:    []*schema.Column{UpworkJobsColumns[17]},
				RefColumns: []*schema.Column{JobsColumns[0]},
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
		{Name: "embedded_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "client_feedback", Type: field.TypeString, Nullable: true},
		{Name: "overall_rating", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "freelancer_earnings", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "start_date", Type: field.TypeTime, Nullable: true},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "budget", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "client_rating", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "client_review_count", Type: field.TypeInt, Nullable: true},
		{Name: "client_country", Type: field.TypeString, Nullable: true},
		{Name: "client_total_jobs_posted", Type: field.TypeInt, Nullable: true},
		{Name: "client_total_spend", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "client_total_hires", Type: field.TypeInt, Nullable: true},
		{Name: "client_active_hires", Type: field.TypeInt, Nullable: true},
		{Name: "client_total_paid_hours", Type: field.TypeInt, Nullable: true},
		{Name: "client_average_hourly_rate_paid", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "DECIMAL"}},
		{Name: "client_company_category", Type: field.TypeString, Nullable: true},
		{Name: "client_company_size", Type: field.TypeString, Nullable: true},
		{Name: "total_proposals", Type: field.TypeInt, Nullable: true},
		{Name: "number_of_interviews", Type: field.TypeInt, Nullable: true},
		{Name: "skills", Type: field.TypeJSON, Nullable: true},
		{Name: "upwork_freelancer_work_histories", Type: field.TypeString},
	}
	// WorkHistoriesTable holds the schema information for the "work_histories" table.
	WorkHistoriesTable = &schema.Table{
		Name:       "work_histories",
		Columns:    WorkHistoriesColumns,
		PrimaryKey: []*schema.Column{WorkHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "work_histories_upwork_freelancers_work_histories",
				Columns:    []*schema.Column{WorkHistoriesColumns[26]},
				RefColumns: []*schema.Column{UpworkFreelancersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// WorkhistoryInferenceDataColumns holds the columns for the "workhistory_inference_data" table.
	WorkhistoryInferenceDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "finalized_job_rating_score", Type: field.TypeFloat64},
		{Name: "is_within_budget", Type: field.TypeBool},
		{Name: "work_history_work_history_inference_data", Type: field.TypeInt},
	}
	// WorkhistoryInferenceDataTable holds the schema information for the "workhistory_inference_data" table.
	WorkhistoryInferenceDataTable = &schema.Table{
		Name:       "workhistory_inference_data",
		Columns:    WorkhistoryInferenceDataColumns,
		PrimaryKey: []*schema.Column{WorkhistoryInferenceDataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workhistory_inference_data_work_histories_work_history_inference_data",
				Columns:    []*schema.Column{WorkhistoryInferenceDataColumns[3]},
				RefColumns: []*schema.Column{WorkHistoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UpworkJobUpworkfreelancerColumns holds the columns for the "upwork_job_upworkfreelancer" table.
	UpworkJobUpworkfreelancerColumns = []*schema.Column{
		{Name: "upwork_job_id", Type: field.TypeString},
		{Name: "upwork_freelancer_id", Type: field.TypeString},
	}
	// UpworkJobUpworkfreelancerTable holds the schema information for the "upwork_job_upworkfreelancer" table.
	UpworkJobUpworkfreelancerTable = &schema.Table{
		Name:       "upwork_job_upworkfreelancer",
		Columns:    UpworkJobUpworkfreelancerColumns,
		PrimaryKey: []*schema.Column{UpworkJobUpworkfreelancerColumns[0], UpworkJobUpworkfreelancerColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "upwork_job_upworkfreelancer_upwork_job_id",
				Columns:    []*schema.Column{UpworkJobUpworkfreelancerColumns[0]},
				RefColumns: []*schema.Column{UpworkJobsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "upwork_job_upworkfreelancer_upwork_freelancer_id",
				Columns:    []*schema.Column{UpworkJobUpworkfreelancerColumns[1]},
				RefColumns: []*schema.Column{UpworkFreelancersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserUpworkjobColumns holds the columns for the "user_upworkjob" table.
	UserUpworkjobColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeString},
		{Name: "upwork_job_id", Type: field.TypeString},
	}
	// UserUpworkjobTable holds the schema information for the "user_upworkjob" table.
	UserUpworkjobTable = &schema.Table{
		Name:       "user_upworkjob",
		Columns:    UserUpworkjobColumns,
		PrimaryKey: []*schema.Column{UserUpworkjobColumns[0], UserUpworkjobColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_upworkjob_user_id",
				Columns:    []*schema.Column{UserUpworkjobColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_upworkjob_upwork_job_id",
				Columns:    []*schema.Column{UserUpworkjobColumns[1]},
				RefColumns: []*schema.Column{UpworkJobsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttachmentRefsTable,
		FreelancerInferenceDataTable,
		JobsTable,
		UpworkFreelancersTable,
		UpworkJobsTable,
		UsersTable,
		WorkHistoriesTable,
		WorkhistoryInferenceDataTable,
		UpworkJobUpworkfreelancerTable,
		UserUpworkjobTable,
	}
)

func init() {
	AttachmentRefsTable.ForeignKeys[0].RefTable = UpworkFreelancersTable
	FreelancerInferenceDataTable.ForeignKeys[0].RefTable = UpworkFreelancersTable
	JobsTable.ForeignKeys[0].RefTable = UsersTable
	UpworkJobsTable.ForeignKeys[0].RefTable = JobsTable
	WorkHistoriesTable.ForeignKeys[0].RefTable = UpworkFreelancersTable
	WorkhistoryInferenceDataTable.ForeignKeys[0].RefTable = WorkHistoriesTable
	UpworkJobUpworkfreelancerTable.ForeignKeys[0].RefTable = UpworkJobsTable
	UpworkJobUpworkfreelancerTable.ForeignKeys[1].RefTable = UpworkFreelancersTable
	UserUpworkjobTable.ForeignKeys[0].RefTable = UsersTable
	UserUpworkjobTable.ForeignKeys[1].RefTable = UpworkJobsTable
}
