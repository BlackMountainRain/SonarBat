// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sonar-bat/ent/rbacpolicy"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// RbacPolicy is the model entity for the RbacPolicy schema.
type RbacPolicy struct {
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
	// role or user
	Role string `json:"role,omitempty"`
	// module or api
	Obj string `json:"obj,omitempty"`
	// action or method
	Act string `json:"act,omitempty"`
	// uri or path
	URI          string `json:"uri,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RbacPolicy) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rbacpolicy.FieldRole, rbacpolicy.FieldObj, rbacpolicy.FieldAct, rbacpolicy.FieldURI:
			values[i] = new(sql.NullString)
		case rbacpolicy.FieldCreatedAt, rbacpolicy.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case rbacpolicy.FieldID, rbacpolicy.FieldUpdatedBy, rbacpolicy.FieldCreatedBy:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RbacPolicy fields.
func (rp *RbacPolicy) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rbacpolicy.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rp.ID = *value
			}
		case rbacpolicy.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rp.CreatedAt = value.Time
			}
		case rbacpolicy.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rp.UpdatedAt = value.Time
			}
		case rbacpolicy.FieldUpdatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value != nil {
				rp.UpdatedBy = *value
			}
		case rbacpolicy.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				rp.CreatedBy = *value
			}
		case rbacpolicy.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				rp.Role = value.String
			}
		case rbacpolicy.FieldObj:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field obj", values[i])
			} else if value.Valid {
				rp.Obj = value.String
			}
		case rbacpolicy.FieldAct:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field act", values[i])
			} else if value.Valid {
				rp.Act = value.String
			}
		case rbacpolicy.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				rp.URI = value.String
			}
		default:
			rp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RbacPolicy.
// This includes values selected through modifiers, order, etc.
func (rp *RbacPolicy) Value(name string) (ent.Value, error) {
	return rp.selectValues.Get(name)
}

// Update returns a builder for updating this RbacPolicy.
// Note that you need to call RbacPolicy.Unwrap() before calling this method if this RbacPolicy
// was returned from a transaction, and the transaction was committed or rolled back.
func (rp *RbacPolicy) Update() *RbacPolicyUpdateOne {
	return NewRbacPolicyClient(rp.config).UpdateOne(rp)
}

// Unwrap unwraps the RbacPolicy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rp *RbacPolicy) Unwrap() *RbacPolicy {
	_tx, ok := rp.config.driver.(*txDriver)
	if !ok {
		panic("ent: RbacPolicy is not a transactional entity")
	}
	rp.config.driver = _tx.drv
	return rp
}

// String implements the fmt.Stringer.
func (rp *RbacPolicy) String() string {
	var builder strings.Builder
	builder.WriteString("RbacPolicy(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(rp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", rp.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", rp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(rp.Role)
	builder.WriteString(", ")
	builder.WriteString("obj=")
	builder.WriteString(rp.Obj)
	builder.WriteString(", ")
	builder.WriteString("act=")
	builder.WriteString(rp.Act)
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(rp.URI)
	builder.WriteByte(')')
	return builder.String()
}

// RbacPolicies is a parsable slice of RbacPolicy.
type RbacPolicies []*RbacPolicy
