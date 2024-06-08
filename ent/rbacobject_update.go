// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sonar-bat/ent/predicate"
	"sonar-bat/ent/rbacobject"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RbacObjectUpdate is the builder for updating RbacObject entities.
type RbacObjectUpdate struct {
	config
	hooks    []Hook
	mutation *RbacObjectMutation
}

// Where appends a list predicates to the RbacObjectUpdate builder.
func (rou *RbacObjectUpdate) Where(ps ...predicate.RbacObject) *RbacObjectUpdate {
	rou.mutation.Where(ps...)
	return rou
}

// SetUpdatedAt sets the "updated_at" field.
func (rou *RbacObjectUpdate) SetUpdatedAt(t time.Time) *RbacObjectUpdate {
	rou.mutation.SetUpdatedAt(t)
	return rou
}

// SetUpdatedBy sets the "updated_by" field.
func (rou *RbacObjectUpdate) SetUpdatedBy(u uuid.UUID) *RbacObjectUpdate {
	rou.mutation.SetUpdatedBy(u)
	return rou
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rou *RbacObjectUpdate) SetNillableUpdatedBy(u *uuid.UUID) *RbacObjectUpdate {
	if u != nil {
		rou.SetUpdatedBy(*u)
	}
	return rou
}

// SetCreatedBy sets the "created_by" field.
func (rou *RbacObjectUpdate) SetCreatedBy(u uuid.UUID) *RbacObjectUpdate {
	rou.mutation.SetCreatedBy(u)
	return rou
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rou *RbacObjectUpdate) SetNillableCreatedBy(u *uuid.UUID) *RbacObjectUpdate {
	if u != nil {
		rou.SetCreatedBy(*u)
	}
	return rou
}

// SetStatus sets the "status" field.
func (rou *RbacObjectUpdate) SetStatus(b bool) *RbacObjectUpdate {
	rou.mutation.SetStatus(b)
	return rou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rou *RbacObjectUpdate) SetNillableStatus(b *bool) *RbacObjectUpdate {
	if b != nil {
		rou.SetStatus(*b)
	}
	return rou
}

// SetValue sets the "value" field.
func (rou *RbacObjectUpdate) SetValue(s string) *RbacObjectUpdate {
	rou.mutation.SetValue(s)
	return rou
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (rou *RbacObjectUpdate) SetNillableValue(s *string) *RbacObjectUpdate {
	if s != nil {
		rou.SetValue(*s)
	}
	return rou
}

// Mutation returns the RbacObjectMutation object of the builder.
func (rou *RbacObjectUpdate) Mutation() *RbacObjectMutation {
	return rou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rou *RbacObjectUpdate) Save(ctx context.Context) (int, error) {
	rou.defaults()
	return withHooks(ctx, rou.sqlSave, rou.mutation, rou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rou *RbacObjectUpdate) SaveX(ctx context.Context) int {
	affected, err := rou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rou *RbacObjectUpdate) Exec(ctx context.Context) error {
	_, err := rou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rou *RbacObjectUpdate) ExecX(ctx context.Context) {
	if err := rou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rou *RbacObjectUpdate) defaults() {
	if _, ok := rou.mutation.UpdatedAt(); !ok {
		v := rbacobject.UpdateDefaultUpdatedAt()
		rou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rou *RbacObjectUpdate) check() error {
	if v, ok := rou.mutation.Value(); ok {
		if err := rbacobject.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "RbacObject.value": %w`, err)}
		}
	}
	return nil
}

func (rou *RbacObjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(rbacobject.Table, rbacobject.Columns, sqlgraph.NewFieldSpec(rbacobject.FieldID, field.TypeInt64))
	if ps := rou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rou.mutation.UpdatedAt(); ok {
		_spec.SetField(rbacobject.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rou.mutation.UpdatedBy(); ok {
		_spec.SetField(rbacobject.FieldUpdatedBy, field.TypeUUID, value)
	}
	if value, ok := rou.mutation.CreatedBy(); ok {
		_spec.SetField(rbacobject.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := rou.mutation.Status(); ok {
		_spec.SetField(rbacobject.FieldStatus, field.TypeBool, value)
	}
	if value, ok := rou.mutation.Value(); ok {
		_spec.SetField(rbacobject.FieldValue, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rbacobject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rou.mutation.done = true
	return n, nil
}

// RbacObjectUpdateOne is the builder for updating a single RbacObject entity.
type RbacObjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RbacObjectMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (rouo *RbacObjectUpdateOne) SetUpdatedAt(t time.Time) *RbacObjectUpdateOne {
	rouo.mutation.SetUpdatedAt(t)
	return rouo
}

// SetUpdatedBy sets the "updated_by" field.
func (rouo *RbacObjectUpdateOne) SetUpdatedBy(u uuid.UUID) *RbacObjectUpdateOne {
	rouo.mutation.SetUpdatedBy(u)
	return rouo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rouo *RbacObjectUpdateOne) SetNillableUpdatedBy(u *uuid.UUID) *RbacObjectUpdateOne {
	if u != nil {
		rouo.SetUpdatedBy(*u)
	}
	return rouo
}

// SetCreatedBy sets the "created_by" field.
func (rouo *RbacObjectUpdateOne) SetCreatedBy(u uuid.UUID) *RbacObjectUpdateOne {
	rouo.mutation.SetCreatedBy(u)
	return rouo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rouo *RbacObjectUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *RbacObjectUpdateOne {
	if u != nil {
		rouo.SetCreatedBy(*u)
	}
	return rouo
}

// SetStatus sets the "status" field.
func (rouo *RbacObjectUpdateOne) SetStatus(b bool) *RbacObjectUpdateOne {
	rouo.mutation.SetStatus(b)
	return rouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rouo *RbacObjectUpdateOne) SetNillableStatus(b *bool) *RbacObjectUpdateOne {
	if b != nil {
		rouo.SetStatus(*b)
	}
	return rouo
}

// SetValue sets the "value" field.
func (rouo *RbacObjectUpdateOne) SetValue(s string) *RbacObjectUpdateOne {
	rouo.mutation.SetValue(s)
	return rouo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (rouo *RbacObjectUpdateOne) SetNillableValue(s *string) *RbacObjectUpdateOne {
	if s != nil {
		rouo.SetValue(*s)
	}
	return rouo
}

// Mutation returns the RbacObjectMutation object of the builder.
func (rouo *RbacObjectUpdateOne) Mutation() *RbacObjectMutation {
	return rouo.mutation
}

// Where appends a list predicates to the RbacObjectUpdate builder.
func (rouo *RbacObjectUpdateOne) Where(ps ...predicate.RbacObject) *RbacObjectUpdateOne {
	rouo.mutation.Where(ps...)
	return rouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rouo *RbacObjectUpdateOne) Select(field string, fields ...string) *RbacObjectUpdateOne {
	rouo.fields = append([]string{field}, fields...)
	return rouo
}

// Save executes the query and returns the updated RbacObject entity.
func (rouo *RbacObjectUpdateOne) Save(ctx context.Context) (*RbacObject, error) {
	rouo.defaults()
	return withHooks(ctx, rouo.sqlSave, rouo.mutation, rouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rouo *RbacObjectUpdateOne) SaveX(ctx context.Context) *RbacObject {
	node, err := rouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rouo *RbacObjectUpdateOne) Exec(ctx context.Context) error {
	_, err := rouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rouo *RbacObjectUpdateOne) ExecX(ctx context.Context) {
	if err := rouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rouo *RbacObjectUpdateOne) defaults() {
	if _, ok := rouo.mutation.UpdatedAt(); !ok {
		v := rbacobject.UpdateDefaultUpdatedAt()
		rouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rouo *RbacObjectUpdateOne) check() error {
	if v, ok := rouo.mutation.Value(); ok {
		if err := rbacobject.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "RbacObject.value": %w`, err)}
		}
	}
	return nil
}

func (rouo *RbacObjectUpdateOne) sqlSave(ctx context.Context) (_node *RbacObject, err error) {
	if err := rouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(rbacobject.Table, rbacobject.Columns, sqlgraph.NewFieldSpec(rbacobject.FieldID, field.TypeInt64))
	id, ok := rouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RbacObject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rbacobject.FieldID)
		for _, f := range fields {
			if !rbacobject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rbacobject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rouo.mutation.UpdatedAt(); ok {
		_spec.SetField(rbacobject.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rouo.mutation.UpdatedBy(); ok {
		_spec.SetField(rbacobject.FieldUpdatedBy, field.TypeUUID, value)
	}
	if value, ok := rouo.mutation.CreatedBy(); ok {
		_spec.SetField(rbacobject.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := rouo.mutation.Status(); ok {
		_spec.SetField(rbacobject.FieldStatus, field.TypeBool, value)
	}
	if value, ok := rouo.mutation.Value(); ok {
		_spec.SetField(rbacobject.FieldValue, field.TypeString, value)
	}
	_node = &RbacObject{config: rouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rbacobject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rouo.mutation.done = true
	return _node, nil
}
