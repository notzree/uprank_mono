// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/notzree/uprank-backend/inference-backend/ent/upworkfreelancer"
	"github.com/notzree/uprank-backend/inference-backend/ent/workhistory"
)

// WorkHistory is the model entity for the WorkHistory schema.
type WorkHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// ClientFeedback holds the value of the "client_feedback" field.
	ClientFeedback string `json:"client_feedback,omitempty"`
	// OverallRating holds the value of the "overall_rating" field.
	OverallRating float64 `json:"overall_rating,omitempty"`
	// FreelancerEarnings holds the value of the "freelancer_earnings" field.
	FreelancerEarnings float64 `json:"freelancer_earnings,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate time.Time `json:"end_date,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Budget holds the value of the "budget" field.
	Budget float64 `json:"budget,omitempty"`
	// ClientRating holds the value of the "client_rating" field.
	ClientRating float64 `json:"client_rating,omitempty"`
	// ClientReviewCount holds the value of the "client_review_count" field.
	ClientReviewCount int `json:"client_review_count,omitempty"`
	// ClientCountry holds the value of the "client_country" field.
	ClientCountry string `json:"client_country,omitempty"`
	// ClientTotalJobsPosted holds the value of the "client_total_jobs_posted" field.
	ClientTotalJobsPosted int `json:"client_total_jobs_posted,omitempty"`
	// ClientTotalSpend holds the value of the "client_total_spend" field.
	ClientTotalSpend float64 `json:"client_total_spend,omitempty"`
	// ClientTotalHires holds the value of the "client_total_hires" field.
	ClientTotalHires int `json:"client_total_hires,omitempty"`
	// ClientActiveHires holds the value of the "client_active_hires" field.
	ClientActiveHires int `json:"client_active_hires,omitempty"`
	// ClientTotalPaidHours holds the value of the "client_total_paid_hours" field.
	ClientTotalPaidHours int `json:"client_total_paid_hours,omitempty"`
	// ClientAverageHourlyRatePaid holds the value of the "client_average_hourly_rate_paid" field.
	ClientAverageHourlyRatePaid float64 `json:"client_average_hourly_rate_paid,omitempty"`
	// ClientCompanyCategory holds the value of the "client_company_category" field.
	ClientCompanyCategory string `json:"client_company_category,omitempty"`
	// ClientCompanySize holds the value of the "client_company_size" field.
	ClientCompanySize string `json:"client_company_size,omitempty"`
	// TotalProposals holds the value of the "total_proposals" field.
	TotalProposals int `json:"total_proposals,omitempty"`
	// NumberOfInterviews holds the value of the "number_of_interviews" field.
	NumberOfInterviews int `json:"number_of_interviews,omitempty"`
	// Skills holds the value of the "skills" field.
	Skills []string `json:"skills,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkHistoryQuery when eager-loading is set.
	Edges                            WorkHistoryEdges `json:"edges"`
	upwork_freelancer_work_histories *string
	selectValues                     sql.SelectValues
}

// WorkHistoryEdges holds the relations/edges for other nodes in the graph.
type WorkHistoryEdges struct {
	// Freelancer holds the value of the freelancer edge.
	Freelancer *UpworkFreelancer `json:"freelancer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FreelancerOrErr returns the Freelancer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkHistoryEdges) FreelancerOrErr() (*UpworkFreelancer, error) {
	if e.Freelancer != nil {
		return e.Freelancer, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: upworkfreelancer.Label}
	}
	return nil, &NotLoadedError{edge: "freelancer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workhistory.FieldSkills:
			values[i] = new([]byte)
		case workhistory.FieldOverallRating, workhistory.FieldFreelancerEarnings, workhistory.FieldBudget, workhistory.FieldClientRating, workhistory.FieldClientTotalSpend, workhistory.FieldClientAverageHourlyRatePaid:
			values[i] = new(sql.NullFloat64)
		case workhistory.FieldID, workhistory.FieldClientReviewCount, workhistory.FieldClientTotalJobsPosted, workhistory.FieldClientTotalHires, workhistory.FieldClientActiveHires, workhistory.FieldClientTotalPaidHours, workhistory.FieldTotalProposals, workhistory.FieldNumberOfInterviews:
			values[i] = new(sql.NullInt64)
		case workhistory.FieldTitle, workhistory.FieldClientFeedback, workhistory.FieldDescription, workhistory.FieldClientCountry, workhistory.FieldClientCompanyCategory, workhistory.FieldClientCompanySize:
			values[i] = new(sql.NullString)
		case workhistory.FieldStartDate, workhistory.FieldEndDate:
			values[i] = new(sql.NullTime)
		case workhistory.ForeignKeys[0]: // upwork_freelancer_work_histories
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkHistory fields.
func (wh *WorkHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wh.ID = int(value.Int64)
		case workhistory.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				wh.Title = value.String
			}
		case workhistory.FieldClientFeedback:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_feedback", values[i])
			} else if value.Valid {
				wh.ClientFeedback = value.String
			}
		case workhistory.FieldOverallRating:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field overall_rating", values[i])
			} else if value.Valid {
				wh.OverallRating = value.Float64
			}
		case workhistory.FieldFreelancerEarnings:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field freelancer_earnings", values[i])
			} else if value.Valid {
				wh.FreelancerEarnings = value.Float64
			}
		case workhistory.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				wh.StartDate = value.Time
			}
		case workhistory.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				wh.EndDate = value.Time
			}
		case workhistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				wh.Description = value.String
			}
		case workhistory.FieldBudget:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field budget", values[i])
			} else if value.Valid {
				wh.Budget = value.Float64
			}
		case workhistory.FieldClientRating:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field client_rating", values[i])
			} else if value.Valid {
				wh.ClientRating = value.Float64
			}
		case workhistory.FieldClientReviewCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field client_review_count", values[i])
			} else if value.Valid {
				wh.ClientReviewCount = int(value.Int64)
			}
		case workhistory.FieldClientCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_country", values[i])
			} else if value.Valid {
				wh.ClientCountry = value.String
			}
		case workhistory.FieldClientTotalJobsPosted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field client_total_jobs_posted", values[i])
			} else if value.Valid {
				wh.ClientTotalJobsPosted = int(value.Int64)
			}
		case workhistory.FieldClientTotalSpend:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field client_total_spend", values[i])
			} else if value.Valid {
				wh.ClientTotalSpend = value.Float64
			}
		case workhistory.FieldClientTotalHires:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field client_total_hires", values[i])
			} else if value.Valid {
				wh.ClientTotalHires = int(value.Int64)
			}
		case workhistory.FieldClientActiveHires:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field client_active_hires", values[i])
			} else if value.Valid {
				wh.ClientActiveHires = int(value.Int64)
			}
		case workhistory.FieldClientTotalPaidHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field client_total_paid_hours", values[i])
			} else if value.Valid {
				wh.ClientTotalPaidHours = int(value.Int64)
			}
		case workhistory.FieldClientAverageHourlyRatePaid:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field client_average_hourly_rate_paid", values[i])
			} else if value.Valid {
				wh.ClientAverageHourlyRatePaid = value.Float64
			}
		case workhistory.FieldClientCompanyCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_company_category", values[i])
			} else if value.Valid {
				wh.ClientCompanyCategory = value.String
			}
		case workhistory.FieldClientCompanySize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_company_size", values[i])
			} else if value.Valid {
				wh.ClientCompanySize = value.String
			}
		case workhistory.FieldTotalProposals:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_proposals", values[i])
			} else if value.Valid {
				wh.TotalProposals = int(value.Int64)
			}
		case workhistory.FieldNumberOfInterviews:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number_of_interviews", values[i])
			} else if value.Valid {
				wh.NumberOfInterviews = int(value.Int64)
			}
		case workhistory.FieldSkills:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field skills", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wh.Skills); err != nil {
					return fmt.Errorf("unmarshal field skills: %w", err)
				}
			}
		case workhistory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field upwork_freelancer_work_histories", values[i])
			} else if value.Valid {
				wh.upwork_freelancer_work_histories = new(string)
				*wh.upwork_freelancer_work_histories = value.String
			}
		default:
			wh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WorkHistory.
// This includes values selected through modifiers, order, etc.
func (wh *WorkHistory) Value(name string) (ent.Value, error) {
	return wh.selectValues.Get(name)
}

// QueryFreelancer queries the "freelancer" edge of the WorkHistory entity.
func (wh *WorkHistory) QueryFreelancer() *UpworkFreelancerQuery {
	return NewWorkHistoryClient(wh.config).QueryFreelancer(wh)
}

// Update returns a builder for updating this WorkHistory.
// Note that you need to call WorkHistory.Unwrap() before calling this method if this WorkHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (wh *WorkHistory) Update() *WorkHistoryUpdateOne {
	return NewWorkHistoryClient(wh.config).UpdateOne(wh)
}

// Unwrap unwraps the WorkHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wh *WorkHistory) Unwrap() *WorkHistory {
	_tx, ok := wh.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkHistory is not a transactional entity")
	}
	wh.config.driver = _tx.drv
	return wh
}

// String implements the fmt.Stringer.
func (wh *WorkHistory) String() string {
	var builder strings.Builder
	builder.WriteString("WorkHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wh.ID))
	builder.WriteString("title=")
	builder.WriteString(wh.Title)
	builder.WriteString(", ")
	builder.WriteString("client_feedback=")
	builder.WriteString(wh.ClientFeedback)
	builder.WriteString(", ")
	builder.WriteString("overall_rating=")
	builder.WriteString(fmt.Sprintf("%v", wh.OverallRating))
	builder.WriteString(", ")
	builder.WriteString("freelancer_earnings=")
	builder.WriteString(fmt.Sprintf("%v", wh.FreelancerEarnings))
	builder.WriteString(", ")
	builder.WriteString("start_date=")
	builder.WriteString(wh.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_date=")
	builder.WriteString(wh.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(wh.Description)
	builder.WriteString(", ")
	builder.WriteString("budget=")
	builder.WriteString(fmt.Sprintf("%v", wh.Budget))
	builder.WriteString(", ")
	builder.WriteString("client_rating=")
	builder.WriteString(fmt.Sprintf("%v", wh.ClientRating))
	builder.WriteString(", ")
	builder.WriteString("client_review_count=")
	builder.WriteString(fmt.Sprintf("%v", wh.ClientReviewCount))
	builder.WriteString(", ")
	builder.WriteString("client_country=")
	builder.WriteString(wh.ClientCountry)
	builder.WriteString(", ")
	builder.WriteString("client_total_jobs_posted=")
	builder.WriteString(fmt.Sprintf("%v", wh.ClientTotalJobsPosted))
	builder.WriteString(", ")
	builder.WriteString("client_total_spend=")
	builder.WriteString(fmt.Sprintf("%v", wh.ClientTotalSpend))
	builder.WriteString(", ")
	builder.WriteString("client_total_hires=")
	builder.WriteString(fmt.Sprintf("%v", wh.ClientTotalHires))
	builder.WriteString(", ")
	builder.WriteString("client_active_hires=")
	builder.WriteString(fmt.Sprintf("%v", wh.ClientActiveHires))
	builder.WriteString(", ")
	builder.WriteString("client_total_paid_hours=")
	builder.WriteString(fmt.Sprintf("%v", wh.ClientTotalPaidHours))
	builder.WriteString(", ")
	builder.WriteString("client_average_hourly_rate_paid=")
	builder.WriteString(fmt.Sprintf("%v", wh.ClientAverageHourlyRatePaid))
	builder.WriteString(", ")
	builder.WriteString("client_company_category=")
	builder.WriteString(wh.ClientCompanyCategory)
	builder.WriteString(", ")
	builder.WriteString("client_company_size=")
	builder.WriteString(wh.ClientCompanySize)
	builder.WriteString(", ")
	builder.WriteString("total_proposals=")
	builder.WriteString(fmt.Sprintf("%v", wh.TotalProposals))
	builder.WriteString(", ")
	builder.WriteString("number_of_interviews=")
	builder.WriteString(fmt.Sprintf("%v", wh.NumberOfInterviews))
	builder.WriteString(", ")
	builder.WriteString("skills=")
	builder.WriteString(fmt.Sprintf("%v", wh.Skills))
	builder.WriteByte(')')
	return builder.String()
}

// WorkHistories is a parsable slice of WorkHistory.
type WorkHistories []*WorkHistory