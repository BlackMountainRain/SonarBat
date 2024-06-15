// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sonar-bat/ent/host"
	"sonar-bat/ent/hostblacklist"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// HostBlacklist is the model entity for the HostBlacklist schema.
type HostBlacklist struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy uuid.UUID `json:"updated_by,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// HostID holds the value of the "host_id" field.
	HostID uuid.UUID `json:"host_id,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HostBlacklistQuery when eager-loading is set.
	Edges        HostBlacklistEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HostBlacklistEdges holds the relations/edges for other nodes in the graph.
type HostBlacklistEdges struct {
	// Host holds the value of the host edge.
	Host *Host `json:"host,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HostOrErr returns the Host value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostBlacklistEdges) HostOrErr() (*Host, error) {
	if e.Host != nil {
		return e.Host, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: host.Label}
	}
	return nil, &NotLoadedError{edge: "host"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HostBlacklist) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hostblacklist.FieldReason:
			values[i] = new(sql.NullString)
		case hostblacklist.FieldCreatedAt, hostblacklist.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case hostblacklist.FieldID, hostblacklist.FieldUpdatedBy, hostblacklist.FieldCreatedBy, hostblacklist.FieldHostID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HostBlacklist fields.
func (hb *HostBlacklist) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hostblacklist.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				hb.ID = *value
			}
		case hostblacklist.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				hb.CreatedAt = value.Time
			}
		case hostblacklist.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				hb.UpdatedAt = value.Time
			}
		case hostblacklist.FieldUpdatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value != nil {
				hb.UpdatedBy = *value
			}
		case hostblacklist.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				hb.CreatedBy = *value
			}
		case hostblacklist.FieldHostID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field host_id", values[i])
			} else if value != nil {
				hb.HostID = *value
			}
		case hostblacklist.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				hb.Reason = value.String
			}
		default:
			hb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HostBlacklist.
// This includes values selected through modifiers, order, etc.
func (hb *HostBlacklist) Value(name string) (ent.Value, error) {
	return hb.selectValues.Get(name)
}

// QueryHost queries the "host" edge of the HostBlacklist entity.
func (hb *HostBlacklist) QueryHost() *HostQuery {
	return NewHostBlacklistClient(hb.config).QueryHost(hb)
}

// Update returns a builder for updating this HostBlacklist.
// Note that you need to call HostBlacklist.Unwrap() before calling this method if this HostBlacklist
// was returned from a transaction, and the transaction was committed or rolled back.
func (hb *HostBlacklist) Update() *HostBlacklistUpdateOne {
	return NewHostBlacklistClient(hb.config).UpdateOne(hb)
}

// Unwrap unwraps the HostBlacklist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hb *HostBlacklist) Unwrap() *HostBlacklist {
	_tx, ok := hb.config.driver.(*txDriver)
	if !ok {
		panic("ent: HostBlacklist is not a transactional entity")
	}
	hb.config.driver = _tx.drv
	return hb
}

// String implements the fmt.Stringer.
func (hb *HostBlacklist) String() string {
	var builder strings.Builder
	builder.WriteString("HostBlacklist(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hb.ID))
	builder.WriteString("created_at=")
	builder.WriteString(hb.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(hb.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", hb.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", hb.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("host_id=")
	builder.WriteString(fmt.Sprintf("%v", hb.HostID))
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(hb.Reason)
	builder.WriteByte(')')
	return builder.String()
}

// HostBlacklists is a parsable slice of HostBlacklist.
type HostBlacklists []*HostBlacklist