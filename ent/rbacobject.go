// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sonar-bat/ent/rbacobject"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// RbacObject is the model entity for the RbacObject schema.
type RbacObject struct {
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
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Value holds the value of the "value" field.
	Value        string `json:"value,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RbacObject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rbacobject.FieldStatus:
			values[i] = new(sql.NullBool)
		case rbacobject.FieldValue:
			values[i] = new(sql.NullString)
		case rbacobject.FieldCreatedAt, rbacobject.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case rbacobject.FieldID, rbacobject.FieldUpdatedBy, rbacobject.FieldCreatedBy:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RbacObject fields.
func (ro *RbacObject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rbacobject.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ro.ID = *value
			}
		case rbacobject.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ro.CreatedAt = value.Time
			}
		case rbacobject.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ro.UpdatedAt = value.Time
			}
		case rbacobject.FieldUpdatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value != nil {
				ro.UpdatedBy = *value
			}
		case rbacobject.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				ro.CreatedBy = *value
			}
		case rbacobject.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ro.Status = value.Bool
			}
		case rbacobject.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				ro.Value = value.String
			}
		default:
			ro.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the RbacObject.
// This includes values selected through modifiers, order, etc.
func (ro *RbacObject) GetValue(name string) (ent.Value, error) {
	return ro.selectValues.Get(name)
}

// Update returns a builder for updating this RbacObject.
// Note that you need to call RbacObject.Unwrap() before calling this method if this RbacObject
// was returned from a transaction, and the transaction was committed or rolled back.
func (ro *RbacObject) Update() *RbacObjectUpdateOne {
	return NewRbacObjectClient(ro.config).UpdateOne(ro)
}

// Unwrap unwraps the RbacObject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ro *RbacObject) Unwrap() *RbacObject {
	_tx, ok := ro.config.driver.(*txDriver)
	if !ok {
		panic("ent: RbacObject is not a transactional entity")
	}
	ro.config.driver = _tx.drv
	return ro
}

// String implements the fmt.Stringer.
func (ro *RbacObject) String() string {
	var builder strings.Builder
	builder.WriteString("RbacObject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ro.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ro.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ro.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ro.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ro.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ro.Status))
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(ro.Value)
	builder.WriteByte(')')
	return builder.String()
}

// RbacObjects is a parsable slice of RbacObject.
type RbacObjects []*RbacObject
