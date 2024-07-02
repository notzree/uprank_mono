// Code generated by ent, DO NOT EDIT.

package upworkfreelancer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the upworkfreelancer type in the database.
	Label = "upwork_freelancer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldTimezone holds the string denoting the timezone field in the database.
	FieldTimezone = "timezone"
	// FieldCv holds the string denoting the cv field in the database.
	FieldCv = "cv"
	// FieldAiReccomended holds the string denoting the ai_reccomended field in the database.
	FieldAiReccomended = "ai_reccomended"
	// FieldFixedChargeAmount holds the string denoting the fixed_charge_amount field in the database.
	FieldFixedChargeAmount = "fixed_charge_amount"
	// FieldFixedChargeCurrency holds the string denoting the fixed_charge_currency field in the database.
	FieldFixedChargeCurrency = "fixed_charge_currency"
	// FieldHourlyChargeAmount holds the string denoting the hourly_charge_amount field in the database.
	FieldHourlyChargeAmount = "hourly_charge_amount"
	// FieldHourlyChargeCurrency holds the string denoting the hourly_charge_currency field in the database.
	FieldHourlyChargeCurrency = "hourly_charge_currency"
	// FieldInvited holds the string denoting the invited field in the database.
	FieldInvited = "invited"
	// FieldPhotoURL holds the string denoting the photo_url field in the database.
	FieldPhotoURL = "photo_url"
	// FieldRecentHours holds the string denoting the recent_hours field in the database.
	FieldRecentHours = "recent_hours"
	// FieldTotalHours holds the string denoting the total_hours field in the database.
	FieldTotalHours = "total_hours"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEmbeddedAt holds the string denoting the embedded_at field in the database.
	FieldEmbeddedAt = "embedded_at"
	// FieldTotalPortfolioItems holds the string denoting the total_portfolio_items field in the database.
	FieldTotalPortfolioItems = "total_portfolio_items"
	// FieldTotalPortfolioV2Items holds the string denoting the total_portfolio_v2_items field in the database.
	FieldTotalPortfolioV2Items = "total_portfolio_v2_items"
	// FieldUpworkTotalFeedback holds the string denoting the upwork_total_feedback field in the database.
	FieldUpworkTotalFeedback = "upwork_total_feedback"
	// FieldUpworkRecentFeedback holds the string denoting the upwork_recent_feedback field in the database.
	FieldUpworkRecentFeedback = "upwork_recent_feedback"
	// FieldUpworkTopRatedStatus holds the string denoting the upwork_top_rated_status field in the database.
	FieldUpworkTopRatedStatus = "upwork_top_rated_status"
	// FieldUpworkTopRatedPlusStatus holds the string denoting the upwork_top_rated_plus_status field in the database.
	FieldUpworkTopRatedPlusStatus = "upwork_top_rated_plus_status"
	// FieldUpworkSponsored holds the string denoting the upwork_sponsored field in the database.
	FieldUpworkSponsored = "upwork_sponsored"
	// FieldUpworkJobSuccessScore holds the string denoting the upwork_job_success_score field in the database.
	FieldUpworkJobSuccessScore = "upwork_job_success_score"
	// FieldUpworkReccomended holds the string denoting the upwork_reccomended field in the database.
	FieldUpworkReccomended = "upwork_reccomended"
	// FieldSkills holds the string denoting the skills field in the database.
	FieldSkills = "skills"
	// FieldAverageRecentEarnings holds the string denoting the average_recent_earnings field in the database.
	FieldAverageRecentEarnings = "average_recent_earnings"
	// FieldCombinedAverageRecentEarnings holds the string denoting the combined_average_recent_earnings field in the database.
	FieldCombinedAverageRecentEarnings = "combined_average_recent_earnings"
	// FieldCombinedRecentEarnings holds the string denoting the combined_recent_earnings field in the database.
	FieldCombinedRecentEarnings = "combined_recent_earnings"
	// FieldCombinedTotalEarnings holds the string denoting the combined_total_earnings field in the database.
	FieldCombinedTotalEarnings = "combined_total_earnings"
	// FieldCombinedTotalRevenue holds the string denoting the combined_total_revenue field in the database.
	FieldCombinedTotalRevenue = "combined_total_revenue"
	// FieldRecentEarnings holds the string denoting the recent_earnings field in the database.
	FieldRecentEarnings = "recent_earnings"
	// FieldTotalRevenue holds the string denoting the total_revenue field in the database.
	FieldTotalRevenue = "total_revenue"
	// FieldUprankSpecializationScore holds the string denoting the uprank_specialization_score field in the database.
	FieldUprankSpecializationScore = "uprank_specialization_score"
	// FieldUprankEstimatedCompletionTime holds the string denoting the uprank_estimated_completion_time field in the database.
	FieldUprankEstimatedCompletionTime = "uprank_estimated_completion_time"
	// FieldUprankReccomended holds the string denoting the uprank_reccomended field in the database.
	FieldUprankReccomended = "uprank_reccomended"
	// FieldUprankReccomendedReasons holds the string denoting the uprank_reccomended_reasons field in the database.
	FieldUprankReccomendedReasons = "uprank_reccomended_reasons"
	// FieldUprankNotEnoughData holds the string denoting the uprank_not_enough_data field in the database.
	FieldUprankNotEnoughData = "uprank_not_enough_data"
	// EdgeUpworkJob holds the string denoting the upwork_job edge name in mutations.
	EdgeUpworkJob = "upwork_job"
	// EdgeAttachments holds the string denoting the attachments edge name in mutations.
	EdgeAttachments = "attachments"
	// EdgeWorkHistories holds the string denoting the work_histories edge name in mutations.
	EdgeWorkHistories = "work_histories"
	// Table holds the table name of the upworkfreelancer in the database.
	Table = "upwork_freelancers"
	// UpworkJobTable is the table that holds the upwork_job relation/edge. The primary key declared below.
	UpworkJobTable = "upwork_job_upworkfreelancer"
	// UpworkJobInverseTable is the table name for the UpworkJob entity.
	// It exists in this package in order to avoid circular dependency with the "upworkjob" package.
	UpworkJobInverseTable = "upwork_jobs"
	// AttachmentsTable is the table that holds the attachments relation/edge.
	AttachmentsTable = "attachment_refs"
	// AttachmentsInverseTable is the table name for the AttachmentRef entity.
	// It exists in this package in order to avoid circular dependency with the "attachmentref" package.
	AttachmentsInverseTable = "attachment_refs"
	// AttachmentsColumn is the table column denoting the attachments relation/edge.
	AttachmentsColumn = "upwork_freelancer_attachments"
	// WorkHistoriesTable is the table that holds the work_histories relation/edge.
	WorkHistoriesTable = "work_histories"
	// WorkHistoriesInverseTable is the table name for the WorkHistory entity.
	// It exists in this package in order to avoid circular dependency with the "workhistory" package.
	WorkHistoriesInverseTable = "work_histories"
	// WorkHistoriesColumn is the table column denoting the work_histories relation/edge.
	WorkHistoriesColumn = "upwork_freelancer_work_histories"
)

// Columns holds all SQL columns for upworkfreelancer fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTitle,
	FieldDescription,
	FieldCity,
	FieldCountry,
	FieldTimezone,
	FieldCv,
	FieldAiReccomended,
	FieldFixedChargeAmount,
	FieldFixedChargeCurrency,
	FieldHourlyChargeAmount,
	FieldHourlyChargeCurrency,
	FieldInvited,
	FieldPhotoURL,
	FieldRecentHours,
	FieldTotalHours,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEmbeddedAt,
	FieldTotalPortfolioItems,
	FieldTotalPortfolioV2Items,
	FieldUpworkTotalFeedback,
	FieldUpworkRecentFeedback,
	FieldUpworkTopRatedStatus,
	FieldUpworkTopRatedPlusStatus,
	FieldUpworkSponsored,
	FieldUpworkJobSuccessScore,
	FieldUpworkReccomended,
	FieldSkills,
	FieldAverageRecentEarnings,
	FieldCombinedAverageRecentEarnings,
	FieldCombinedRecentEarnings,
	FieldCombinedTotalEarnings,
	FieldCombinedTotalRevenue,
	FieldRecentEarnings,
	FieldTotalRevenue,
	FieldUprankSpecializationScore,
	FieldUprankEstimatedCompletionTime,
	FieldUprankReccomended,
	FieldUprankReccomendedReasons,
	FieldUprankNotEnoughData,
}

var (
	// UpworkJobPrimaryKey and UpworkJobColumn2 are the table columns denoting the
	// primary key for the upwork_job relation (M2M).
	UpworkJobPrimaryKey = []string{"upwork_job_id", "upwork_freelancer_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultUprankSpecializationScore holds the default value on creation for the "uprank_specialization_score" field.
	DefaultUprankSpecializationScore float64
	// DefaultUprankReccomended holds the default value on creation for the "uprank_reccomended" field.
	DefaultUprankReccomended bool
	// DefaultUprankNotEnoughData holds the default value on creation for the "uprank_not_enough_data" field.
	DefaultUprankNotEnoughData bool
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the UpworkFreelancer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByTimezone orders the results by the timezone field.
func ByTimezone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimezone, opts...).ToFunc()
}

// ByCv orders the results by the cv field.
func ByCv(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCv, opts...).ToFunc()
}

// ByAiReccomended orders the results by the ai_reccomended field.
func ByAiReccomended(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAiReccomended, opts...).ToFunc()
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

// ByInvited orders the results by the invited field.
func ByInvited(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvited, opts...).ToFunc()
}

// ByPhotoURL orders the results by the photo_url field.
func ByPhotoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL, opts...).ToFunc()
}

// ByRecentHours orders the results by the recent_hours field.
func ByRecentHours(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecentHours, opts...).ToFunc()
}

// ByTotalHours orders the results by the total_hours field.
func ByTotalHours(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalHours, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByEmbeddedAt orders the results by the embedded_at field.
func ByEmbeddedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmbeddedAt, opts...).ToFunc()
}

// ByTotalPortfolioItems orders the results by the total_portfolio_items field.
func ByTotalPortfolioItems(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalPortfolioItems, opts...).ToFunc()
}

// ByTotalPortfolioV2Items orders the results by the total_portfolio_v2_items field.
func ByTotalPortfolioV2Items(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalPortfolioV2Items, opts...).ToFunc()
}

// ByUpworkTotalFeedback orders the results by the upwork_total_feedback field.
func ByUpworkTotalFeedback(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpworkTotalFeedback, opts...).ToFunc()
}

// ByUpworkRecentFeedback orders the results by the upwork_recent_feedback field.
func ByUpworkRecentFeedback(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpworkRecentFeedback, opts...).ToFunc()
}

// ByUpworkTopRatedStatus orders the results by the upwork_top_rated_status field.
func ByUpworkTopRatedStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpworkTopRatedStatus, opts...).ToFunc()
}

// ByUpworkTopRatedPlusStatus orders the results by the upwork_top_rated_plus_status field.
func ByUpworkTopRatedPlusStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpworkTopRatedPlusStatus, opts...).ToFunc()
}

// ByUpworkSponsored orders the results by the upwork_sponsored field.
func ByUpworkSponsored(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpworkSponsored, opts...).ToFunc()
}

// ByUpworkJobSuccessScore orders the results by the upwork_job_success_score field.
func ByUpworkJobSuccessScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpworkJobSuccessScore, opts...).ToFunc()
}

// ByUpworkReccomended orders the results by the upwork_reccomended field.
func ByUpworkReccomended(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpworkReccomended, opts...).ToFunc()
}

// ByAverageRecentEarnings orders the results by the average_recent_earnings field.
func ByAverageRecentEarnings(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAverageRecentEarnings, opts...).ToFunc()
}

// ByCombinedAverageRecentEarnings orders the results by the combined_average_recent_earnings field.
func ByCombinedAverageRecentEarnings(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCombinedAverageRecentEarnings, opts...).ToFunc()
}

// ByCombinedRecentEarnings orders the results by the combined_recent_earnings field.
func ByCombinedRecentEarnings(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCombinedRecentEarnings, opts...).ToFunc()
}

// ByCombinedTotalEarnings orders the results by the combined_total_earnings field.
func ByCombinedTotalEarnings(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCombinedTotalEarnings, opts...).ToFunc()
}

// ByCombinedTotalRevenue orders the results by the combined_total_revenue field.
func ByCombinedTotalRevenue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCombinedTotalRevenue, opts...).ToFunc()
}

// ByRecentEarnings orders the results by the recent_earnings field.
func ByRecentEarnings(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecentEarnings, opts...).ToFunc()
}

// ByTotalRevenue orders the results by the total_revenue field.
func ByTotalRevenue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalRevenue, opts...).ToFunc()
}

// ByUprankSpecializationScore orders the results by the uprank_specialization_score field.
func ByUprankSpecializationScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUprankSpecializationScore, opts...).ToFunc()
}

// ByUprankEstimatedCompletionTime orders the results by the uprank_estimated_completion_time field.
func ByUprankEstimatedCompletionTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUprankEstimatedCompletionTime, opts...).ToFunc()
}

// ByUprankReccomended orders the results by the uprank_reccomended field.
func ByUprankReccomended(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUprankReccomended, opts...).ToFunc()
}

// ByUprankReccomendedReasons orders the results by the uprank_reccomended_reasons field.
func ByUprankReccomendedReasons(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUprankReccomendedReasons, opts...).ToFunc()
}

// ByUprankNotEnoughData orders the results by the uprank_not_enough_data field.
func ByUprankNotEnoughData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUprankNotEnoughData, opts...).ToFunc()
}

// ByUpworkJobCount orders the results by upwork_job count.
func ByUpworkJobCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUpworkJobStep(), opts...)
	}
}

// ByUpworkJob orders the results by upwork_job terms.
func ByUpworkJob(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUpworkJobStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAttachmentsCount orders the results by attachments count.
func ByAttachmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttachmentsStep(), opts...)
	}
}

// ByAttachments orders the results by attachments terms.
func ByAttachments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttachmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkHistoriesCount orders the results by work_histories count.
func ByWorkHistoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkHistoriesStep(), opts...)
	}
}

// ByWorkHistories orders the results by work_histories terms.
func ByWorkHistories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkHistoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUpworkJobStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UpworkJobInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UpworkJobTable, UpworkJobPrimaryKey...),
	)
}
func newAttachmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttachmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttachmentsTable, AttachmentsColumn),
	)
}
func newWorkHistoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkHistoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkHistoriesTable, WorkHistoriesColumn),
	)
}
