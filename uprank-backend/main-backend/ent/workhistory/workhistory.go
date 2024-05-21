// Code generated by ent, DO NOT EDIT.

package workhistory

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the workhistory type in the database.
	Label = "work_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldClientFeedback holds the string denoting the client_feedback field in the database.
	FieldClientFeedback = "client_feedback"
	// FieldOverallRating holds the string denoting the overall_rating field in the database.
	FieldOverallRating = "overall_rating"
	// FieldFixedChargeAmount holds the string denoting the fixed_charge_amount field in the database.
	FieldFixedChargeAmount = "fixed_charge_amount"
	// FieldFixedChargeCurrency holds the string denoting the fixed_charge_currency field in the database.
	FieldFixedChargeCurrency = "fixed_charge_currency"
	// FieldHourlyChargeAmount holds the string denoting the hourly_charge_amount field in the database.
	FieldHourlyChargeAmount = "hourly_charge_amount"
	// FieldHourlyChargeCurrency holds the string denoting the hourly_charge_currency field in the database.
	FieldHourlyChargeCurrency = "hourly_charge_currency"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldJobDescription holds the string denoting the job_description field in the database.
	FieldJobDescription = "job_description"
	// FieldTotalProposals holds the string denoting the total_proposals field in the database.
	FieldTotalProposals = "total_proposals"
	// FieldNumberOfInterviews holds the string denoting the number_of_interviews field in the database.
	FieldNumberOfInterviews = "number_of_interviews"
	// FieldSkills holds the string denoting the skills field in the database.
	FieldSkills = "skills"
	// FieldClientRating holds the string denoting the client_rating field in the database.
	FieldClientRating = "client_rating"
	// FieldClientReviewCount holds the string denoting the client_review_count field in the database.
	FieldClientReviewCount = "client_review_count"
	// FieldClientCountry holds the string denoting the client_country field in the database.
	FieldClientCountry = "client_country"
	// FieldClientTotalJobsPosted holds the string denoting the client_total_jobs_posted field in the database.
	FieldClientTotalJobsPosted = "client_total_jobs_posted"
	// FieldClientTotalSpend holds the string denoting the client_total_spend field in the database.
	FieldClientTotalSpend = "client_total_spend"
	// FieldClientTotalHires holds the string denoting the client_total_hires field in the database.
	FieldClientTotalHires = "client_total_hires"
	// FieldClientTotalPaidHours holds the string denoting the client_total_paid_hours field in the database.
	FieldClientTotalPaidHours = "client_total_paid_hours"
	// FieldClientAverageHourlyRatePaid holds the string denoting the client_average_hourly_rate_paid field in the database.
	FieldClientAverageHourlyRatePaid = "client_average_hourly_rate_paid"
	// FieldClientCompanyCategory holds the string denoting the client_company_category field in the database.
	FieldClientCompanyCategory = "client_company_category"
	// FieldClientCompanySize holds the string denoting the client_company_size field in the database.
	FieldClientCompanySize = "client_company_size"
	// EdgeUpworkFreelancerProposal holds the string denoting the upwork_freelancer_proposal edge name in mutations.
	EdgeUpworkFreelancerProposal = "upwork_Freelancer_Proposal"
	// Table holds the table name of the workhistory in the database.
	Table = "work_histories"
	// UpworkFreelancerProposalTable is the table that holds the upwork_Freelancer_Proposal relation/edge.
	UpworkFreelancerProposalTable = "work_histories"
	// UpworkFreelancerProposalInverseTable is the table name for the Freelancer entity.
	// It exists in this package in order to avoid circular dependency with the "freelancer" package.
	UpworkFreelancerProposalInverseTable = "freelancers"
	// UpworkFreelancerProposalColumn is the table column denoting the upwork_Freelancer_Proposal relation/edge.
	UpworkFreelancerProposalColumn = "freelancer_work_histories"
)

// Columns holds all SQL columns for workhistory fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldClientFeedback,
	FieldOverallRating,
	FieldFixedChargeAmount,
	FieldFixedChargeCurrency,
	FieldHourlyChargeAmount,
	FieldHourlyChargeCurrency,
	FieldStartDate,
	FieldEndDate,
	FieldJobDescription,
	FieldTotalProposals,
	FieldNumberOfInterviews,
	FieldSkills,
	FieldClientRating,
	FieldClientReviewCount,
	FieldClientCountry,
	FieldClientTotalJobsPosted,
	FieldClientTotalSpend,
	FieldClientTotalHires,
	FieldClientTotalPaidHours,
	FieldClientAverageHourlyRatePaid,
	FieldClientCompanyCategory,
	FieldClientCompanySize,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "work_histories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"freelancer_work_histories",
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

// OrderOption defines the ordering options for the WorkHistory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByClientFeedback orders the results by the client_feedback field.
func ByClientFeedback(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientFeedback, opts...).ToFunc()
}

// ByOverallRating orders the results by the overall_rating field.
func ByOverallRating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOverallRating, opts...).ToFunc()
}

// ByFixedChargeAmount orders the results by the fixed_charge_amount field.
func ByFixedChargeAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFixedChargeAmount, opts...).ToFunc()
}

// ByFixedChargeCurrency orders the results by the fixed_charge_currency field.
func ByFixedChargeCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFixedChargeCurrency, opts...).ToFunc()
}

// ByHourlyChargeAmount orders the results by the hourly_charge_amount field.
func ByHourlyChargeAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHourlyChargeAmount, opts...).ToFunc()
}

// ByHourlyChargeCurrency orders the results by the hourly_charge_currency field.
func ByHourlyChargeCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHourlyChargeCurrency, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByJobDescription orders the results by the job_description field.
func ByJobDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJobDescription, opts...).ToFunc()
}

// ByTotalProposals orders the results by the total_proposals field.
func ByTotalProposals(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalProposals, opts...).ToFunc()
}

// ByNumberOfInterviews orders the results by the number_of_interviews field.
func ByNumberOfInterviews(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumberOfInterviews, opts...).ToFunc()
}

// ByClientRating orders the results by the client_rating field.
func ByClientRating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientRating, opts...).ToFunc()
}

// ByClientReviewCount orders the results by the client_review_count field.
func ByClientReviewCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientReviewCount, opts...).ToFunc()
}

// ByClientCountry orders the results by the client_country field.
func ByClientCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientCountry, opts...).ToFunc()
}

// ByClientTotalJobsPosted orders the results by the client_total_jobs_posted field.
func ByClientTotalJobsPosted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientTotalJobsPosted, opts...).ToFunc()
}

// ByClientTotalSpend orders the results by the client_total_spend field.
func ByClientTotalSpend(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientTotalSpend, opts...).ToFunc()
}

// ByClientTotalHires orders the results by the client_total_hires field.
func ByClientTotalHires(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientTotalHires, opts...).ToFunc()
}

// ByClientTotalPaidHours orders the results by the client_total_paid_hours field.
func ByClientTotalPaidHours(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientTotalPaidHours, opts...).ToFunc()
}

// ByClientAverageHourlyRatePaid orders the results by the client_average_hourly_rate_paid field.
func ByClientAverageHourlyRatePaid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientAverageHourlyRatePaid, opts...).ToFunc()
}

// ByClientCompanyCategory orders the results by the client_company_category field.
func ByClientCompanyCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientCompanyCategory, opts...).ToFunc()
}

// ByClientCompanySize orders the results by the client_company_size field.
func ByClientCompanySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientCompanySize, opts...).ToFunc()
}

// ByUpworkFreelancerProposalField orders the results by upwork_Freelancer_Proposal field.
func ByUpworkFreelancerProposalField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUpworkFreelancerProposalStep(), sql.OrderByField(field, opts...))
	}
}
func newUpworkFreelancerProposalStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UpworkFreelancerProposalInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UpworkFreelancerProposalTable, UpworkFreelancerProposalColumn),
	)
}