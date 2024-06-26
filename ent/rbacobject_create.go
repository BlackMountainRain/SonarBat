// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sonar-bat/ent/rbacobject"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RbacObjectCreate is the builder for creating a RbacObject entity.
type RbacObjectCreate struct {
	config
	mutation *RbacObjectMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (roc *RbacObjectCreate) SetCreatedAt(t time.Time) *RbacObjectCreate {
	roc.mutation.SetCreatedAt(t)
	return roc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (roc *RbacObjectCreate) SetNillableCreatedAt(t *time.Time) *RbacObjectCreate {
	if t != nil {
		roc.SetCreatedAt(*t)
	}
	return roc
}

// SetUpdatedAt sets the "updated_at" field.
func (roc *RbacObjectCreate) SetUpdatedAt(t time.Time) *RbacObjectCreate {
	roc.mutation.SetUpdatedAt(t)
	return roc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (roc *RbacObjectCreate) SetNillableUpdatedAt(t *time.Time) *RbacObjectCreate {
	if t != nil {
		roc.SetUpdatedAt(*t)
	}
	return roc
}

// SetUpdatedBy sets the "updated_by" field.
func (roc *RbacObjectCreate) SetUpdatedBy(u uuid.UUID) *RbacObjectCreate {
	roc.mutation.SetUpdatedBy(u)
	return roc
}

// SetCreatedBy sets the "created_by" field.
func (roc *RbacObjectCreate) SetCreatedBy(u uuid.UUID) *RbacObjectCreate {
	roc.mutation.SetCreatedBy(u)
	return roc
}

// SetStatus sets the "status" field.
func (roc *RbacObjectCreate) SetStatus(b bool) *RbacObjectCreate {
	roc.mutation.SetStatus(b)
	return roc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (roc *RbacObjectCreate) SetNillableStatus(b *bool) *RbacObjectCreate {
	if b != nil {
		roc.SetStatus(*b)
	}
	return roc
}

// SetValue sets the "value" field.
func (roc *RbacObjectCreate) SetValue(s string) *RbacObjectCreate {
	roc.mutation.SetValue(s)
	return roc
}

// SetID sets the "id" field.
func (roc *RbacObjectCreate) SetID(u uuid.UUID) *RbacObjectCreate {
	roc.mutation.SetID(u)
	return roc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (roc *RbacObjectCreate) SetNillableID(u *uuid.UUID) *RbacObjectCreate {
	if u != nil {
		roc.SetID(*u)
	}
	return roc
}

// Mutation returns the RbacObjectMutation object of the builder.
func (roc *RbacObjectCreate) Mutation() *RbacObjectMutation {
	return roc.mutation
}

// Save creates the RbacObject in the database.
func (roc *RbacObjectCreate) Save(ctx context.Context) (*RbacObject, error) {
	roc.defaults()
	return withHooks(ctx, roc.sqlSave, roc.mutation, roc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (roc *RbacObjectCreate) SaveX(ctx context.Context) *RbacObject {
	v, err := roc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (roc *RbacObjectCreate) Exec(ctx context.Context) error {
	_, err := roc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (roc *RbacObjectCreate) ExecX(ctx context.Context) {
	if err := roc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (roc *RbacObjectCreate) defaults() {
	if _, ok := roc.mutation.CreatedAt(); !ok {
		v := rbacobject.DefaultCreatedAt()
		roc.mutation.SetCreatedAt(v)
	}
	if _, ok := roc.mutation.UpdatedAt(); !ok {
		v := rbacobject.DefaultUpdatedAt()
		roc.mutation.SetUpdatedAt(v)
	}
	if _, ok := roc.mutation.Status(); !ok {
		v := rbacobject.DefaultStatus
		roc.mutation.SetStatus(v)
	}
	if _, ok := roc.mutation.ID(); !ok {
		v := rbacobject.DefaultID()
		roc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (roc *RbacObjectCreate) check() error {
	if _, ok := roc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RbacObject.created_at"`)}
	}
	if _, ok := roc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RbacObject.updated_at"`)}
	}
	if _, ok := roc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "RbacObject.updated_by"`)}
	}
	if _, ok := roc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "RbacObject.created_by"`)}
	}
	if _, ok := roc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "RbacObject.status"`)}
	}
	if _, ok := roc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "RbacObject.value"`)}
	}
	if v, ok := roc.mutation.Value(); ok {
		if err := rbacobject.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "RbacObject.value": %w`, err)}
		}
	}
	return nil
}

func (roc *RbacObjectCreate) sqlSave(ctx context.Context) (*RbacObject, error) {
	if err := roc.check(); err != nil {
		return nil, err
	}
	_node, _spec := roc.createSpec()
	if err := sqlgraph.CreateNode(ctx, roc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	roc.mutation.id = &_node.ID
	roc.mutation.done = true
	return _node, nil
}

func (roc *RbacObjectCreate) createSpec() (*RbacObject, *sqlgraph.CreateSpec) {
	var (
		_node = &RbacObject{config: roc.config}
		_spec = sqlgraph.NewCreateSpec(rbacobject.Table, sqlgraph.NewFieldSpec(rbacobject.FieldID, field.TypeUUID))
	)
	if id, ok := roc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := roc.mutation.CreatedAt(); ok {
		_spec.SetField(rbacobject.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := roc.mutation.UpdatedAt(); ok {
		_spec.SetField(rbacobject.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := roc.mutation.UpdatedBy(); ok {
		_spec.SetField(rbacobject.FieldUpdatedBy, field.TypeUUID, value)
		_node.UpdatedBy = value
	}
	if value, ok := roc.mutation.CreatedBy(); ok {
		_spec.SetField(rbacobject.FieldCreatedBy, field.TypeUUID, value)
		_node.CreatedBy = value
	}
	if value, ok := roc.mutation.Status(); ok {
		_spec.SetField(rbacobject.FieldStatus, field.TypeBool, value)
		_node.Status = value
	}
	if value, ok := roc.mutation.Value(); ok {
		_spec.SetField(rbacobject.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	return _node, _spec
}

// RbacObjectCreateBulk is the builder for creating many RbacObject entities in bulk.
type RbacObjectCreateBulk struct {
	config
	err      error
	builders []*RbacObjectCreate
}

// Save creates the RbacObject entities in the database.
func (rocb *RbacObjectCreateBulk) Save(ctx context.Context) ([]*RbacObject, error) {
	if rocb.err != nil {
		return nil, rocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rocb.builders))
	nodes := make([]*RbacObject, len(rocb.builders))
	mutators := make([]Mutator, len(rocb.builders))
	for i := range rocb.builders {
		func(i int, root context.Context) {
			builder := rocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RbacObjectMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rocb *RbacObjectCreateBulk) SaveX(ctx context.Context) []*RbacObject {
	v, err := rocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rocb *RbacObjectCreateBulk) Exec(ctx context.Context) error {
	_, err := rocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rocb *RbacObjectCreateBulk) ExecX(ctx context.Context) {
	if err := rocb.Exec(ctx); err != nil {
		panic(err)
	}
}
