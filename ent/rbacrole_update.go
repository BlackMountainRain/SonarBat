// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sonar-bat/ent/predicate"
	"sonar-bat/ent/rbacrole"
	"sonar-bat/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RbacRoleUpdate is the builder for updating RbacRole entities.
type RbacRoleUpdate struct {
	config
	hooks    []Hook
	mutation *RbacRoleMutation
}

// Where appends a list predicates to the RbacRoleUpdate builder.
func (rru *RbacRoleUpdate) Where(ps ...predicate.RbacRole) *RbacRoleUpdate {
	rru.mutation.Where(ps...)
	return rru
}

// SetUpdatedAt sets the "updated_at" field.
func (rru *RbacRoleUpdate) SetUpdatedAt(t time.Time) *RbacRoleUpdate {
	rru.mutation.SetUpdatedAt(t)
	return rru
}

// SetUpdatedBy sets the "updated_by" field.
func (rru *RbacRoleUpdate) SetUpdatedBy(u uuid.UUID) *RbacRoleUpdate {
	rru.mutation.SetUpdatedBy(u)
	return rru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rru *RbacRoleUpdate) SetNillableUpdatedBy(u *uuid.UUID) *RbacRoleUpdate {
	if u != nil {
		rru.SetUpdatedBy(*u)
	}
	return rru
}

// SetCreatedBy sets the "created_by" field.
func (rru *RbacRoleUpdate) SetCreatedBy(u uuid.UUID) *RbacRoleUpdate {
	rru.mutation.SetCreatedBy(u)
	return rru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rru *RbacRoleUpdate) SetNillableCreatedBy(u *uuid.UUID) *RbacRoleUpdate {
	if u != nil {
		rru.SetCreatedBy(*u)
	}
	return rru
}

// SetStatus sets the "status" field.
func (rru *RbacRoleUpdate) SetStatus(b bool) *RbacRoleUpdate {
	rru.mutation.SetStatus(b)
	return rru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rru *RbacRoleUpdate) SetNillableStatus(b *bool) *RbacRoleUpdate {
	if b != nil {
		rru.SetStatus(*b)
	}
	return rru
}

// SetName sets the "name" field.
func (rru *RbacRoleUpdate) SetName(s string) *RbacRoleUpdate {
	rru.mutation.SetName(s)
	return rru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rru *RbacRoleUpdate) SetNillableName(s *string) *RbacRoleUpdate {
	if s != nil {
		rru.SetName(*s)
	}
	return rru
}

// SetDescription sets the "description" field.
func (rru *RbacRoleUpdate) SetDescription(s string) *RbacRoleUpdate {
	rru.mutation.SetDescription(s)
	return rru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rru *RbacRoleUpdate) SetNillableDescription(s *string) *RbacRoleUpdate {
	if s != nil {
		rru.SetDescription(*s)
	}
	return rru
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (rru *RbacRoleUpdate) AddUserIDs(ids ...uuid.UUID) *RbacRoleUpdate {
	rru.mutation.AddUserIDs(ids...)
	return rru
}

// AddUsers adds the "users" edges to the User entity.
func (rru *RbacRoleUpdate) AddUsers(u ...*User) *RbacRoleUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rru.AddUserIDs(ids...)
}

// Mutation returns the RbacRoleMutation object of the builder.
func (rru *RbacRoleUpdate) Mutation() *RbacRoleMutation {
	return rru.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (rru *RbacRoleUpdate) ClearUsers() *RbacRoleUpdate {
	rru.mutation.ClearUsers()
	return rru
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (rru *RbacRoleUpdate) RemoveUserIDs(ids ...uuid.UUID) *RbacRoleUpdate {
	rru.mutation.RemoveUserIDs(ids...)
	return rru
}

// RemoveUsers removes "users" edges to User entities.
func (rru *RbacRoleUpdate) RemoveUsers(u ...*User) *RbacRoleUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rru *RbacRoleUpdate) Save(ctx context.Context) (int, error) {
	rru.defaults()
	return withHooks(ctx, rru.sqlSave, rru.mutation, rru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rru *RbacRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := rru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rru *RbacRoleUpdate) Exec(ctx context.Context) error {
	_, err := rru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rru *RbacRoleUpdate) ExecX(ctx context.Context) {
	if err := rru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rru *RbacRoleUpdate) defaults() {
	if _, ok := rru.mutation.UpdatedAt(); !ok {
		v := rbacrole.UpdateDefaultUpdatedAt()
		rru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rru *RbacRoleUpdate) check() error {
	if v, ok := rru.mutation.Name(); ok {
		if err := rbacrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "RbacRole.name": %w`, err)}
		}
	}
	if v, ok := rru.mutation.Description(); ok {
		if err := rbacrole.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "RbacRole.description": %w`, err)}
		}
	}
	return nil
}

func (rru *RbacRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(rbacrole.Table, rbacrole.Columns, sqlgraph.NewFieldSpec(rbacrole.FieldID, field.TypeUUID))
	if ps := rru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rru.mutation.UpdatedAt(); ok {
		_spec.SetField(rbacrole.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rru.mutation.UpdatedBy(); ok {
		_spec.SetField(rbacrole.FieldUpdatedBy, field.TypeUUID, value)
	}
	if value, ok := rru.mutation.CreatedBy(); ok {
		_spec.SetField(rbacrole.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := rru.mutation.Status(); ok {
		_spec.SetField(rbacrole.FieldStatus, field.TypeBool, value)
	}
	if value, ok := rru.mutation.Name(); ok {
		_spec.SetField(rbacrole.FieldName, field.TypeString, value)
	}
	if value, ok := rru.mutation.Description(); ok {
		_spec.SetField(rbacrole.FieldDescription, field.TypeString, value)
	}
	if rru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rbacrole.UsersTable,
			Columns: rbacrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !rru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rbacrole.UsersTable,
			Columns: rbacrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rbacrole.UsersTable,
			Columns: rbacrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rbacrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rru.mutation.done = true
	return n, nil
}

// RbacRoleUpdateOne is the builder for updating a single RbacRole entity.
type RbacRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RbacRoleMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (rruo *RbacRoleUpdateOne) SetUpdatedAt(t time.Time) *RbacRoleUpdateOne {
	rruo.mutation.SetUpdatedAt(t)
	return rruo
}

// SetUpdatedBy sets the "updated_by" field.
func (rruo *RbacRoleUpdateOne) SetUpdatedBy(u uuid.UUID) *RbacRoleUpdateOne {
	rruo.mutation.SetUpdatedBy(u)
	return rruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rruo *RbacRoleUpdateOne) SetNillableUpdatedBy(u *uuid.UUID) *RbacRoleUpdateOne {
	if u != nil {
		rruo.SetUpdatedBy(*u)
	}
	return rruo
}

// SetCreatedBy sets the "created_by" field.
func (rruo *RbacRoleUpdateOne) SetCreatedBy(u uuid.UUID) *RbacRoleUpdateOne {
	rruo.mutation.SetCreatedBy(u)
	return rruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rruo *RbacRoleUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *RbacRoleUpdateOne {
	if u != nil {
		rruo.SetCreatedBy(*u)
	}
	return rruo
}

// SetStatus sets the "status" field.
func (rruo *RbacRoleUpdateOne) SetStatus(b bool) *RbacRoleUpdateOne {
	rruo.mutation.SetStatus(b)
	return rruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rruo *RbacRoleUpdateOne) SetNillableStatus(b *bool) *RbacRoleUpdateOne {
	if b != nil {
		rruo.SetStatus(*b)
	}
	return rruo
}

// SetName sets the "name" field.
func (rruo *RbacRoleUpdateOne) SetName(s string) *RbacRoleUpdateOne {
	rruo.mutation.SetName(s)
	return rruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rruo *RbacRoleUpdateOne) SetNillableName(s *string) *RbacRoleUpdateOne {
	if s != nil {
		rruo.SetName(*s)
	}
	return rruo
}

// SetDescription sets the "description" field.
func (rruo *RbacRoleUpdateOne) SetDescription(s string) *RbacRoleUpdateOne {
	rruo.mutation.SetDescription(s)
	return rruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rruo *RbacRoleUpdateOne) SetNillableDescription(s *string) *RbacRoleUpdateOne {
	if s != nil {
		rruo.SetDescription(*s)
	}
	return rruo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (rruo *RbacRoleUpdateOne) AddUserIDs(ids ...uuid.UUID) *RbacRoleUpdateOne {
	rruo.mutation.AddUserIDs(ids...)
	return rruo
}

// AddUsers adds the "users" edges to the User entity.
func (rruo *RbacRoleUpdateOne) AddUsers(u ...*User) *RbacRoleUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rruo.AddUserIDs(ids...)
}

// Mutation returns the RbacRoleMutation object of the builder.
func (rruo *RbacRoleUpdateOne) Mutation() *RbacRoleMutation {
	return rruo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (rruo *RbacRoleUpdateOne) ClearUsers() *RbacRoleUpdateOne {
	rruo.mutation.ClearUsers()
	return rruo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (rruo *RbacRoleUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *RbacRoleUpdateOne {
	rruo.mutation.RemoveUserIDs(ids...)
	return rruo
}

// RemoveUsers removes "users" edges to User entities.
func (rruo *RbacRoleUpdateOne) RemoveUsers(u ...*User) *RbacRoleUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rruo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the RbacRoleUpdate builder.
func (rruo *RbacRoleUpdateOne) Where(ps ...predicate.RbacRole) *RbacRoleUpdateOne {
	rruo.mutation.Where(ps...)
	return rruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rruo *RbacRoleUpdateOne) Select(field string, fields ...string) *RbacRoleUpdateOne {
	rruo.fields = append([]string{field}, fields...)
	return rruo
}

// Save executes the query and returns the updated RbacRole entity.
func (rruo *RbacRoleUpdateOne) Save(ctx context.Context) (*RbacRole, error) {
	rruo.defaults()
	return withHooks(ctx, rruo.sqlSave, rruo.mutation, rruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rruo *RbacRoleUpdateOne) SaveX(ctx context.Context) *RbacRole {
	node, err := rruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rruo *RbacRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := rruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rruo *RbacRoleUpdateOne) ExecX(ctx context.Context) {
	if err := rruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rruo *RbacRoleUpdateOne) defaults() {
	if _, ok := rruo.mutation.UpdatedAt(); !ok {
		v := rbacrole.UpdateDefaultUpdatedAt()
		rruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rruo *RbacRoleUpdateOne) check() error {
	if v, ok := rruo.mutation.Name(); ok {
		if err := rbacrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "RbacRole.name": %w`, err)}
		}
	}
	if v, ok := rruo.mutation.Description(); ok {
		if err := rbacrole.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "RbacRole.description": %w`, err)}
		}
	}
	return nil
}

func (rruo *RbacRoleUpdateOne) sqlSave(ctx context.Context) (_node *RbacRole, err error) {
	if err := rruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(rbacrole.Table, rbacrole.Columns, sqlgraph.NewFieldSpec(rbacrole.FieldID, field.TypeUUID))
	id, ok := rruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RbacRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rbacrole.FieldID)
		for _, f := range fields {
			if !rbacrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rbacrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rruo.mutation.UpdatedAt(); ok {
		_spec.SetField(rbacrole.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rruo.mutation.UpdatedBy(); ok {
		_spec.SetField(rbacrole.FieldUpdatedBy, field.TypeUUID, value)
	}
	if value, ok := rruo.mutation.CreatedBy(); ok {
		_spec.SetField(rbacrole.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := rruo.mutation.Status(); ok {
		_spec.SetField(rbacrole.FieldStatus, field.TypeBool, value)
	}
	if value, ok := rruo.mutation.Name(); ok {
		_spec.SetField(rbacrole.FieldName, field.TypeString, value)
	}
	if value, ok := rruo.mutation.Description(); ok {
		_spec.SetField(rbacrole.FieldDescription, field.TypeString, value)
	}
	if rruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rbacrole.UsersTable,
			Columns: rbacrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !rruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rbacrole.UsersTable,
			Columns: rbacrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   rbacrole.UsersTable,
			Columns: rbacrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RbacRole{config: rruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rbacrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rruo.mutation.done = true
	return _node, nil
}
