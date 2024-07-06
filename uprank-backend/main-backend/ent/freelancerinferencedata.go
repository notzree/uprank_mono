// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/jackc/pgtype"
	"github.com/notzree/uprank-backend/main-backend/ent/freelancerinferencedata"
	"github.com/notzree/uprank-backend/main-backend/ent/upworkfreelancer"
)

// FreelancerInferenceData is the model entity for the FreelancerInferenceData schema.
type FreelancerInferenceData struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UprankReccomended holds the value of the "uprank_reccomended" field.
	UprankReccomended bool `json:"uprank_reccomended,omitempty"`
	// UprankReccomendedReasons holds the value of the "uprank_reccomended_reasons" field.
	UprankReccomendedReasons string `json:"uprank_reccomended_reasons,omitempty"`
	// UprankNotEnoughData holds the value of the "uprank_not_enough_data" field.
	UprankNotEnoughData bool `json:"uprank_not_enough_data,omitempty"`
	// FinalizedRatingScore holds the value of the "finalized_rating_score" field.
	FinalizedRatingScore float64 `json:"finalized_rating_score,omitempty"`
	// AiEstimatedDuration holds the value of the "ai_estimated_duration" field.
	AiEstimatedDuration *pgtype.Interval `json:"ai_estimated_duration,omitempty"`
	// BudgetAdherencePercentage holds the value of the "budget_adherence_percentage" field.
	BudgetAdherencePercentage float64 `json:"budget_adherence_percentage,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FreelancerInferenceDataQuery when eager-loading is set.
	Edges                                       FreelancerInferenceDataEdges `json:"edges"`
	upwork_freelancer_freelancer_inference_data *string
	selectValues                                sql.SelectValues
}

// FreelancerInferenceDataEdges holds the relations/edges for other nodes in the graph.
type FreelancerInferenceDataEdges struct {
	// Upworkfreelancer holds the value of the upworkfreelancer edge.
	Upworkfreelancer *UpworkFreelancer `json:"upworkfreelancer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UpworkfreelancerOrErr returns the Upworkfreelancer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FreelancerInferenceDataEdges) UpworkfreelancerOrErr() (*UpworkFreelancer, error) {
	if e.Upworkfreelancer != nil {
		return e.Upworkfreelancer, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: upworkfreelancer.Label}
	}
	return nil, &NotLoadedError{edge: "upworkfreelancer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FreelancerInferenceData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case freelancerinferencedata.FieldAiEstimatedDuration:
			values[i] = new(pgtype.Interval)
		case freelancerinferencedata.FieldUprankReccomended, freelancerinferencedata.FieldUprankNotEnoughData:
			values[i] = new(sql.NullBool)
		case freelancerinferencedata.FieldFinalizedRatingScore, freelancerinferencedata.FieldBudgetAdherencePercentage:
			values[i] = new(sql.NullFloat64)
		case freelancerinferencedata.FieldID:
			values[i] = new(sql.NullInt64)
		case freelancerinferencedata.FieldUprankReccomendedReasons:
			values[i] = new(sql.NullString)
		case freelancerinferencedata.ForeignKeys[0]: // upwork_freelancer_freelancer_inference_data
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FreelancerInferenceData fields.
func (fid *FreelancerInferenceData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case freelancerinferencedata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fid.ID = int(value.Int64)
		case freelancerinferencedata.FieldUprankReccomended:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field uprank_reccomended", values[i])
			} else if value.Valid {
				fid.UprankReccomended = value.Bool
			}
		case freelancerinferencedata.FieldUprankReccomendedReasons:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uprank_reccomended_reasons", values[i])
			} else if value.Valid {
				fid.UprankReccomendedReasons = value.String
			}
		case freelancerinferencedata.FieldUprankNotEnoughData:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field uprank_not_enough_data", values[i])
			} else if value.Valid {
				fid.UprankNotEnoughData = value.Bool
			}
		case freelancerinferencedata.FieldFinalizedRatingScore:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field finalized_rating_score", values[i])
			} else if value.Valid {
				fid.FinalizedRatingScore = value.Float64
			}
		case freelancerinferencedata.FieldAiEstimatedDuration:
			if value, ok := values[i].(*pgtype.Interval); !ok {
				return fmt.Errorf("unexpected type %T for field ai_estimated_duration", values[i])
			} else if value != nil {
				fid.AiEstimatedDuration = value
			}
		case freelancerinferencedata.FieldBudgetAdherencePercentage:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field budget_adherence_percentage", values[i])
			} else if value.Valid {
				fid.BudgetAdherencePercentage = value.Float64
			}
		case freelancerinferencedata.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field upwork_freelancer_freelancer_inference_data", values[i])
			} else if value.Valid {
				fid.upwork_freelancer_freelancer_inference_data = new(string)
				*fid.upwork_freelancer_freelancer_inference_data = value.String
			}
		default:
			fid.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FreelancerInferenceData.
// This includes values selected through modifiers, order, etc.
func (fid *FreelancerInferenceData) Value(name string) (ent.Value, error) {
	return fid.selectValues.Get(name)
}

// QueryUpworkfreelancer queries the "upworkfreelancer" edge of the FreelancerInferenceData entity.
func (fid *FreelancerInferenceData) QueryUpworkfreelancer() *UpworkFreelancerQuery {
	return NewFreelancerInferenceDataClient(fid.config).QueryUpworkfreelancer(fid)
}

// Update returns a builder for updating this FreelancerInferenceData.
// Note that you need to call FreelancerInferenceData.Unwrap() before calling this method if this FreelancerInferenceData
// was returned from a transaction, and the transaction was committed or rolled back.
func (fid *FreelancerInferenceData) Update() *FreelancerInferenceDataUpdateOne {
	return NewFreelancerInferenceDataClient(fid.config).UpdateOne(fid)
}

// Unwrap unwraps the FreelancerInferenceData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fid *FreelancerInferenceData) Unwrap() *FreelancerInferenceData {
	_tx, ok := fid.config.driver.(*txDriver)
	if !ok {
		panic("ent: FreelancerInferenceData is not a transactional entity")
	}
	fid.config.driver = _tx.drv
	return fid
}

// String implements the fmt.Stringer.
func (fid *FreelancerInferenceData) String() string {
	var builder strings.Builder
	builder.WriteString("FreelancerInferenceData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fid.ID))
	builder.WriteString("uprank_reccomended=")
	builder.WriteString(fmt.Sprintf("%v", fid.UprankReccomended))
	builder.WriteString(", ")
	builder.WriteString("uprank_reccomended_reasons=")
	builder.WriteString(fid.UprankReccomendedReasons)
	builder.WriteString(", ")
	builder.WriteString("uprank_not_enough_data=")
	builder.WriteString(fmt.Sprintf("%v", fid.UprankNotEnoughData))
	builder.WriteString(", ")
	builder.WriteString("finalized_rating_score=")
	builder.WriteString(fmt.Sprintf("%v", fid.FinalizedRatingScore))
	builder.WriteString(", ")
	builder.WriteString("ai_estimated_duration=")
	builder.WriteString(fmt.Sprintf("%v", fid.AiEstimatedDuration))
	builder.WriteString(", ")
	builder.WriteString("budget_adherence_percentage=")
	builder.WriteString(fmt.Sprintf("%v", fid.BudgetAdherencePercentage))
	builder.WriteByte(')')
	return builder.String()
}

// FreelancerInferenceDataSlice is a parsable slice of FreelancerInferenceData.
type FreelancerInferenceDataSlice []*FreelancerInferenceData
