// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/notzree/uprank-backend/main-backend/ent/upworkfreelancer"
)

// UpworkFreelancer is the model entity for the UpworkFreelancer schema.
type UpworkFreelancer struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Timezone holds the value of the "timezone" field.
	Timezone string `json:"timezone,omitempty"`
	// Cv holds the value of the "cv" field.
	Cv string `json:"cv,omitempty"`
	// AiReccomended holds the value of the "ai_reccomended" field.
	AiReccomended bool `json:"ai_reccomended,omitempty"`
	// FixedChargeAmount holds the value of the "fixed_charge_amount" field.
	FixedChargeAmount float64 `json:"fixed_charge_amount,omitempty"`
	// FixedChargeCurrency holds the value of the "fixed_charge_currency" field.
	FixedChargeCurrency string `json:"fixed_charge_currency,omitempty"`
	// HourlyChargeAmount holds the value of the "hourly_charge_amount" field.
	HourlyChargeAmount float64 `json:"hourly_charge_amount,omitempty"`
	// HourlyChargeCurrency holds the value of the "hourly_charge_currency" field.
	HourlyChargeCurrency string `json:"hourly_charge_currency,omitempty"`
	// Invited holds the value of the "invited" field.
	Invited bool `json:"invited,omitempty"`
	// PhotoURL holds the value of the "photo_url" field.
	PhotoURL string `json:"photo_url,omitempty"`
	// RecentHours holds the value of the "recent_hours" field.
	RecentHours float64 `json:"recent_hours,omitempty"`
	// TotalHours holds the value of the "total_hours" field.
	TotalHours float64 `json:"total_hours,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// EmbeddedAt holds the value of the "embedded_at" field.
	EmbeddedAt time.Time `json:"embedded_at,omitempty"`
	// TotalPortfolioItems holds the value of the "total_portfolio_items" field.
	TotalPortfolioItems int `json:"total_portfolio_items,omitempty"`
	// TotalPortfolioV2Items holds the value of the "total_portfolio_v2_items" field.
	TotalPortfolioV2Items int `json:"total_portfolio_v2_items,omitempty"`
	// UpworkTotalFeedback holds the value of the "upwork_total_feedback" field.
	UpworkTotalFeedback float64 `json:"upwork_total_feedback,omitempty"`
	// UpworkRecentFeedback holds the value of the "upwork_recent_feedback" field.
	UpworkRecentFeedback float64 `json:"upwork_recent_feedback,omitempty"`
	// UpworkTopRatedStatus holds the value of the "upwork_top_rated_status" field.
	UpworkTopRatedStatus bool `json:"upwork_top_rated_status,omitempty"`
	// UpworkTopRatedPlusStatus holds the value of the "upwork_top_rated_plus_status" field.
	UpworkTopRatedPlusStatus bool `json:"upwork_top_rated_plus_status,omitempty"`
	// UpworkSponsored holds the value of the "upwork_sponsored" field.
	UpworkSponsored bool `json:"upwork_sponsored,omitempty"`
	// UpworkJobSuccessScore holds the value of the "upwork_job_success_score" field.
	UpworkJobSuccessScore float64 `json:"upwork_job_success_score,omitempty"`
	// UpworkReccomended holds the value of the "upwork_reccomended" field.
	UpworkReccomended bool `json:"upwork_reccomended,omitempty"`
	// Skills holds the value of the "skills" field.
	Skills []string `json:"skills,omitempty"`
	// AverageRecentEarnings holds the value of the "average_recent_earnings" field.
	AverageRecentEarnings float64 `json:"average_recent_earnings,omitempty"`
	// CombinedAverageRecentEarnings holds the value of the "combined_average_recent_earnings" field.
	CombinedAverageRecentEarnings float64 `json:"combined_average_recent_earnings,omitempty"`
	// CombinedRecentEarnings holds the value of the "combined_recent_earnings" field.
	CombinedRecentEarnings float64 `json:"combined_recent_earnings,omitempty"`
	// CombinedTotalEarnings holds the value of the "combined_total_earnings" field.
	CombinedTotalEarnings float64 `json:"combined_total_earnings,omitempty"`
	// CombinedTotalRevenue holds the value of the "combined_total_revenue" field.
	CombinedTotalRevenue float64 `json:"combined_total_revenue,omitempty"`
	// RecentEarnings holds the value of the "recent_earnings" field.
	RecentEarnings float64 `json:"recent_earnings,omitempty"`
	// TotalRevenue holds the value of the "total_revenue" field.
	TotalRevenue float64 `json:"total_revenue,omitempty"`
	// MissingFields holds the value of the "missing_fields" field.
	MissingFields bool `json:"missing_fields,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UpworkFreelancerQuery when eager-loading is set.
	Edges        UpworkFreelancerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UpworkFreelancerEdges holds the relations/edges for other nodes in the graph.
type UpworkFreelancerEdges struct {
	// UpworkJob holds the value of the upwork_job edge.
	UpworkJob []*UpworkJob `json:"upwork_job,omitempty"`
	// Attachments holds the value of the attachments edge.
	Attachments []*AttachmentRef `json:"attachments,omitempty"`
	// WorkHistories holds the value of the work_histories edge.
	WorkHistories []*WorkHistory `json:"work_histories,omitempty"`
	// FreelancerInferenceData holds the value of the freelancer_inference_data edge.
	FreelancerInferenceData []*FreelancerInferenceData `json:"freelancer_inference_data,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UpworkJobOrErr returns the UpworkJob value or an error if the edge
// was not loaded in eager-loading.
func (e UpworkFreelancerEdges) UpworkJobOrErr() ([]*UpworkJob, error) {
	if e.loadedTypes[0] {
		return e.UpworkJob, nil
	}
	return nil, &NotLoadedError{edge: "upwork_job"}
}

// AttachmentsOrErr returns the Attachments value or an error if the edge
// was not loaded in eager-loading.
func (e UpworkFreelancerEdges) AttachmentsOrErr() ([]*AttachmentRef, error) {
	if e.loadedTypes[1] {
		return e.Attachments, nil
	}
	return nil, &NotLoadedError{edge: "attachments"}
}

// WorkHistoriesOrErr returns the WorkHistories value or an error if the edge
// was not loaded in eager-loading.
func (e UpworkFreelancerEdges) WorkHistoriesOrErr() ([]*WorkHistory, error) {
	if e.loadedTypes[2] {
		return e.WorkHistories, nil
	}
	return nil, &NotLoadedError{edge: "work_histories"}
}

// FreelancerInferenceDataOrErr returns the FreelancerInferenceData value or an error if the edge
// was not loaded in eager-loading.
func (e UpworkFreelancerEdges) FreelancerInferenceDataOrErr() ([]*FreelancerInferenceData, error) {
	if e.loadedTypes[3] {
		return e.FreelancerInferenceData, nil
	}
	return nil, &NotLoadedError{edge: "freelancer_inference_data"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UpworkFreelancer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case upworkfreelancer.FieldSkills:
			values[i] = new([]byte)
		case upworkfreelancer.FieldAiReccomended, upworkfreelancer.FieldInvited, upworkfreelancer.FieldUpworkTopRatedStatus, upworkfreelancer.FieldUpworkTopRatedPlusStatus, upworkfreelancer.FieldUpworkSponsored, upworkfreelancer.FieldUpworkReccomended, upworkfreelancer.FieldMissingFields:
			values[i] = new(sql.NullBool)
		case upworkfreelancer.FieldFixedChargeAmount, upworkfreelancer.FieldHourlyChargeAmount, upworkfreelancer.FieldRecentHours, upworkfreelancer.FieldTotalHours, upworkfreelancer.FieldUpworkTotalFeedback, upworkfreelancer.FieldUpworkRecentFeedback, upworkfreelancer.FieldUpworkJobSuccessScore, upworkfreelancer.FieldAverageRecentEarnings, upworkfreelancer.FieldCombinedAverageRecentEarnings, upworkfreelancer.FieldCombinedRecentEarnings, upworkfreelancer.FieldCombinedTotalEarnings, upworkfreelancer.FieldCombinedTotalRevenue, upworkfreelancer.FieldRecentEarnings, upworkfreelancer.FieldTotalRevenue:
			values[i] = new(sql.NullFloat64)
		case upworkfreelancer.FieldTotalPortfolioItems, upworkfreelancer.FieldTotalPortfolioV2Items:
			values[i] = new(sql.NullInt64)
		case upworkfreelancer.FieldID, upworkfreelancer.FieldName, upworkfreelancer.FieldTitle, upworkfreelancer.FieldDescription, upworkfreelancer.FieldCity, upworkfreelancer.FieldCountry, upworkfreelancer.FieldTimezone, upworkfreelancer.FieldCv, upworkfreelancer.FieldFixedChargeCurrency, upworkfreelancer.FieldHourlyChargeCurrency, upworkfreelancer.FieldPhotoURL:
			values[i] = new(sql.NullString)
		case upworkfreelancer.FieldCreatedAt, upworkfreelancer.FieldUpdatedAt, upworkfreelancer.FieldEmbeddedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UpworkFreelancer fields.
func (uf *UpworkFreelancer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case upworkfreelancer.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				uf.ID = value.String
			}
		case upworkfreelancer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				uf.Name = value.String
			}
		case upworkfreelancer.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				uf.Title = value.String
			}
		case upworkfreelancer.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				uf.Description = value.String
			}
		case upworkfreelancer.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				uf.City = value.String
			}
		case upworkfreelancer.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				uf.Country = value.String
			}
		case upworkfreelancer.FieldTimezone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timezone", values[i])
			} else if value.Valid {
				uf.Timezone = value.String
			}
		case upworkfreelancer.FieldCv:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cv", values[i])
			} else if value.Valid {
				uf.Cv = value.String
			}
		case upworkfreelancer.FieldAiReccomended:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field ai_reccomended", values[i])
			} else if value.Valid {
				uf.AiReccomended = value.Bool
			}
		case upworkfreelancer.FieldFixedChargeAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fixed_charge_amount", values[i])
			} else if value.Valid {
				uf.FixedChargeAmount = value.Float64
			}
		case upworkfreelancer.FieldFixedChargeCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fixed_charge_currency", values[i])
			} else if value.Valid {
				uf.FixedChargeCurrency = value.String
			}
		case upworkfreelancer.FieldHourlyChargeAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field hourly_charge_amount", values[i])
			} else if value.Valid {
				uf.HourlyChargeAmount = value.Float64
			}
		case upworkfreelancer.FieldHourlyChargeCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hourly_charge_currency", values[i])
			} else if value.Valid {
				uf.HourlyChargeCurrency = value.String
			}
		case upworkfreelancer.FieldInvited:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field invited", values[i])
			} else if value.Valid {
				uf.Invited = value.Bool
			}
		case upworkfreelancer.FieldPhotoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url", values[i])
			} else if value.Valid {
				uf.PhotoURL = value.String
			}
		case upworkfreelancer.FieldRecentHours:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field recent_hours", values[i])
			} else if value.Valid {
				uf.RecentHours = value.Float64
			}
		case upworkfreelancer.FieldTotalHours:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_hours", values[i])
			} else if value.Valid {
				uf.TotalHours = value.Float64
			}
		case upworkfreelancer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				uf.CreatedAt = value.Time
			}
		case upworkfreelancer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				uf.UpdatedAt = value.Time
			}
		case upworkfreelancer.FieldEmbeddedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field embedded_at", values[i])
			} else if value.Valid {
				uf.EmbeddedAt = value.Time
			}
		case upworkfreelancer.FieldTotalPortfolioItems:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_portfolio_items", values[i])
			} else if value.Valid {
				uf.TotalPortfolioItems = int(value.Int64)
			}
		case upworkfreelancer.FieldTotalPortfolioV2Items:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_portfolio_v2_items", values[i])
			} else if value.Valid {
				uf.TotalPortfolioV2Items = int(value.Int64)
			}
		case upworkfreelancer.FieldUpworkTotalFeedback:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field upwork_total_feedback", values[i])
			} else if value.Valid {
				uf.UpworkTotalFeedback = value.Float64
			}
		case upworkfreelancer.FieldUpworkRecentFeedback:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field upwork_recent_feedback", values[i])
			} else if value.Valid {
				uf.UpworkRecentFeedback = value.Float64
			}
		case upworkfreelancer.FieldUpworkTopRatedStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field upwork_top_rated_status", values[i])
			} else if value.Valid {
				uf.UpworkTopRatedStatus = value.Bool
			}
		case upworkfreelancer.FieldUpworkTopRatedPlusStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field upwork_top_rated_plus_status", values[i])
			} else if value.Valid {
				uf.UpworkTopRatedPlusStatus = value.Bool
			}
		case upworkfreelancer.FieldUpworkSponsored:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field upwork_sponsored", values[i])
			} else if value.Valid {
				uf.UpworkSponsored = value.Bool
			}
		case upworkfreelancer.FieldUpworkJobSuccessScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field upwork_job_success_score", values[i])
			} else if value.Valid {
				uf.UpworkJobSuccessScore = value.Float64
			}
		case upworkfreelancer.FieldUpworkReccomended:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field upwork_reccomended", values[i])
			} else if value.Valid {
				uf.UpworkReccomended = value.Bool
			}
		case upworkfreelancer.FieldSkills:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field skills", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &uf.Skills); err != nil {
					return fmt.Errorf("unmarshal field skills: %w", err)
				}
			}
		case upworkfreelancer.FieldAverageRecentEarnings:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field average_recent_earnings", values[i])
			} else if value.Valid {
				uf.AverageRecentEarnings = value.Float64
			}
		case upworkfreelancer.FieldCombinedAverageRecentEarnings:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field combined_average_recent_earnings", values[i])
			} else if value.Valid {
				uf.CombinedAverageRecentEarnings = value.Float64
			}
		case upworkfreelancer.FieldCombinedRecentEarnings:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field combined_recent_earnings", values[i])
			} else if value.Valid {
				uf.CombinedRecentEarnings = value.Float64
			}
		case upworkfreelancer.FieldCombinedTotalEarnings:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field combined_total_earnings", values[i])
			} else if value.Valid {
				uf.CombinedTotalEarnings = value.Float64
			}
		case upworkfreelancer.FieldCombinedTotalRevenue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field combined_total_revenue", values[i])
			} else if value.Valid {
				uf.CombinedTotalRevenue = value.Float64
			}
		case upworkfreelancer.FieldRecentEarnings:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field recent_earnings", values[i])
			} else if value.Valid {
				uf.RecentEarnings = value.Float64
			}
		case upworkfreelancer.FieldTotalRevenue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_revenue", values[i])
			} else if value.Valid {
				uf.TotalRevenue = value.Float64
			}
		case upworkfreelancer.FieldMissingFields:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field missing_fields", values[i])
			} else if value.Valid {
				uf.MissingFields = value.Bool
			}
		default:
			uf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UpworkFreelancer.
// This includes values selected through modifiers, order, etc.
func (uf *UpworkFreelancer) Value(name string) (ent.Value, error) {
	return uf.selectValues.Get(name)
}

// QueryUpworkJob queries the "upwork_job" edge of the UpworkFreelancer entity.
func (uf *UpworkFreelancer) QueryUpworkJob() *UpworkJobQuery {
	return NewUpworkFreelancerClient(uf.config).QueryUpworkJob(uf)
}

// QueryAttachments queries the "attachments" edge of the UpworkFreelancer entity.
func (uf *UpworkFreelancer) QueryAttachments() *AttachmentRefQuery {
	return NewUpworkFreelancerClient(uf.config).QueryAttachments(uf)
}

// QueryWorkHistories queries the "work_histories" edge of the UpworkFreelancer entity.
func (uf *UpworkFreelancer) QueryWorkHistories() *WorkHistoryQuery {
	return NewUpworkFreelancerClient(uf.config).QueryWorkHistories(uf)
}

// QueryFreelancerInferenceData queries the "freelancer_inference_data" edge of the UpworkFreelancer entity.
func (uf *UpworkFreelancer) QueryFreelancerInferenceData() *FreelancerInferenceDataQuery {
	return NewUpworkFreelancerClient(uf.config).QueryFreelancerInferenceData(uf)
}

// Update returns a builder for updating this UpworkFreelancer.
// Note that you need to call UpworkFreelancer.Unwrap() before calling this method if this UpworkFreelancer
// was returned from a transaction, and the transaction was committed or rolled back.
func (uf *UpworkFreelancer) Update() *UpworkFreelancerUpdateOne {
	return NewUpworkFreelancerClient(uf.config).UpdateOne(uf)
}

// Unwrap unwraps the UpworkFreelancer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uf *UpworkFreelancer) Unwrap() *UpworkFreelancer {
	_tx, ok := uf.config.driver.(*txDriver)
	if !ok {
		panic("ent: UpworkFreelancer is not a transactional entity")
	}
	uf.config.driver = _tx.drv
	return uf
}

// String implements the fmt.Stringer.
func (uf *UpworkFreelancer) String() string {
	var builder strings.Builder
	builder.WriteString("UpworkFreelancer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uf.ID))
	builder.WriteString("name=")
	builder.WriteString(uf.Name)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(uf.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(uf.Description)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(uf.City)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(uf.Country)
	builder.WriteString(", ")
	builder.WriteString("timezone=")
	builder.WriteString(uf.Timezone)
	builder.WriteString(", ")
	builder.WriteString("cv=")
	builder.WriteString(uf.Cv)
	builder.WriteString(", ")
	builder.WriteString("ai_reccomended=")
	builder.WriteString(fmt.Sprintf("%v", uf.AiReccomended))
	builder.WriteString(", ")
	builder.WriteString("fixed_charge_amount=")
	builder.WriteString(fmt.Sprintf("%v", uf.FixedChargeAmount))
	builder.WriteString(", ")
	builder.WriteString("fixed_charge_currency=")
	builder.WriteString(uf.FixedChargeCurrency)
	builder.WriteString(", ")
	builder.WriteString("hourly_charge_amount=")
	builder.WriteString(fmt.Sprintf("%v", uf.HourlyChargeAmount))
	builder.WriteString(", ")
	builder.WriteString("hourly_charge_currency=")
	builder.WriteString(uf.HourlyChargeCurrency)
	builder.WriteString(", ")
	builder.WriteString("invited=")
	builder.WriteString(fmt.Sprintf("%v", uf.Invited))
	builder.WriteString(", ")
	builder.WriteString("photo_url=")
	builder.WriteString(uf.PhotoURL)
	builder.WriteString(", ")
	builder.WriteString("recent_hours=")
	builder.WriteString(fmt.Sprintf("%v", uf.RecentHours))
	builder.WriteString(", ")
	builder.WriteString("total_hours=")
	builder.WriteString(fmt.Sprintf("%v", uf.TotalHours))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(uf.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(uf.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("embedded_at=")
	builder.WriteString(uf.EmbeddedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("total_portfolio_items=")
	builder.WriteString(fmt.Sprintf("%v", uf.TotalPortfolioItems))
	builder.WriteString(", ")
	builder.WriteString("total_portfolio_v2_items=")
	builder.WriteString(fmt.Sprintf("%v", uf.TotalPortfolioV2Items))
	builder.WriteString(", ")
	builder.WriteString("upwork_total_feedback=")
	builder.WriteString(fmt.Sprintf("%v", uf.UpworkTotalFeedback))
	builder.WriteString(", ")
	builder.WriteString("upwork_recent_feedback=")
	builder.WriteString(fmt.Sprintf("%v", uf.UpworkRecentFeedback))
	builder.WriteString(", ")
	builder.WriteString("upwork_top_rated_status=")
	builder.WriteString(fmt.Sprintf("%v", uf.UpworkTopRatedStatus))
	builder.WriteString(", ")
	builder.WriteString("upwork_top_rated_plus_status=")
	builder.WriteString(fmt.Sprintf("%v", uf.UpworkTopRatedPlusStatus))
	builder.WriteString(", ")
	builder.WriteString("upwork_sponsored=")
	builder.WriteString(fmt.Sprintf("%v", uf.UpworkSponsored))
	builder.WriteString(", ")
	builder.WriteString("upwork_job_success_score=")
	builder.WriteString(fmt.Sprintf("%v", uf.UpworkJobSuccessScore))
	builder.WriteString(", ")
	builder.WriteString("upwork_reccomended=")
	builder.WriteString(fmt.Sprintf("%v", uf.UpworkReccomended))
	builder.WriteString(", ")
	builder.WriteString("skills=")
	builder.WriteString(fmt.Sprintf("%v", uf.Skills))
	builder.WriteString(", ")
	builder.WriteString("average_recent_earnings=")
	builder.WriteString(fmt.Sprintf("%v", uf.AverageRecentEarnings))
	builder.WriteString(", ")
	builder.WriteString("combined_average_recent_earnings=")
	builder.WriteString(fmt.Sprintf("%v", uf.CombinedAverageRecentEarnings))
	builder.WriteString(", ")
	builder.WriteString("combined_recent_earnings=")
	builder.WriteString(fmt.Sprintf("%v", uf.CombinedRecentEarnings))
	builder.WriteString(", ")
	builder.WriteString("combined_total_earnings=")
	builder.WriteString(fmt.Sprintf("%v", uf.CombinedTotalEarnings))
	builder.WriteString(", ")
	builder.WriteString("combined_total_revenue=")
	builder.WriteString(fmt.Sprintf("%v", uf.CombinedTotalRevenue))
	builder.WriteString(", ")
	builder.WriteString("recent_earnings=")
	builder.WriteString(fmt.Sprintf("%v", uf.RecentEarnings))
	builder.WriteString(", ")
	builder.WriteString("total_revenue=")
	builder.WriteString(fmt.Sprintf("%v", uf.TotalRevenue))
	builder.WriteString(", ")
	builder.WriteString("missing_fields=")
	builder.WriteString(fmt.Sprintf("%v", uf.MissingFields))
	builder.WriteByte(')')
	return builder.String()
}

// UpworkFreelancers is a parsable slice of UpworkFreelancer.
type UpworkFreelancers []*UpworkFreelancer
