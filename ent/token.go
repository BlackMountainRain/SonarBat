// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sonar-bat/ent/token"
	"sonar-bat/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Token is the model entity for the Token schema.
type Token struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy uuid.UUID `json:"updated_by,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TokenQuery when eager-loading is set.
	Edges        TokenEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TokenEdges holds the relations/edges for other nodes in the graph.
type TokenEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TokenEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Token) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case token.FieldStatus:
			values[i] = new(sql.NullBool)
		case token.FieldID, token.FieldUserID:
			values[i] = new(sql.NullInt64)
		case token.FieldName, token.FieldRemark, token.FieldToken:
			values[i] = new(sql.NullString)
		case token.FieldCreatedAt, token.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case token.FieldUpdatedBy, token.FieldCreatedBy:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Token fields.
func (t *Token) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case token.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int64(value.Int64)
		case token.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case token.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case token.FieldUpdatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value != nil {
				t.UpdatedBy = *value
			}
		case token.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				t.CreatedBy = *value
			}
		case token.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				t.UserID = value.Int64
			}
		case token.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = value.Bool
			}
		case token.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case token.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				t.Remark = value.String
			}
		case token.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				t.Token = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Token.
// This includes values selected through modifiers, order, etc.
func (t *Token) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Token entity.
func (t *Token) QueryUser() *UserQuery {
	return NewTokenClient(t.config).QueryUser(t)
}

// Update returns a builder for updating this Token.
// Note that you need to call Token.Unwrap() before calling this method if this Token
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Token) Update() *TokenUpdateOne {
	return NewTokenClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Token entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Token) Unwrap() *Token {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Token is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Token) String() string {
	var builder strings.Builder
	builder.WriteString("Token(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", t.UserID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(t.Remark)
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(t.Token)
	builder.WriteByte(')')
	return builder.String()
}

// Tokens is a parsable slice of Token.
type Tokens []*Token
