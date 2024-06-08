// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sonar-bat/ent/rbacpolicy"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RbacPolicyCreate is the builder for creating a RbacPolicy entity.
type RbacPolicyCreate struct {
	config
	mutation *RbacPolicyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rpc *RbacPolicyCreate) SetCreatedAt(t time.Time) *RbacPolicyCreate {
	rpc.mutation.SetCreatedAt(t)
	return rpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rpc *RbacPolicyCreate) SetNillableCreatedAt(t *time.Time) *RbacPolicyCreate {
	if t != nil {
		rpc.SetCreatedAt(*t)
	}
	return rpc
}

// SetUpdatedAt sets the "updated_at" field.
func (rpc *RbacPolicyCreate) SetUpdatedAt(t time.Time) *RbacPolicyCreate {
	rpc.mutation.SetUpdatedAt(t)
	return rpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rpc *RbacPolicyCreate) SetNillableUpdatedAt(t *time.Time) *RbacPolicyCreate {
	if t != nil {
		rpc.SetUpdatedAt(*t)
	}
	return rpc
}

// SetUpdatedBy sets the "updated_by" field.
func (rpc *RbacPolicyCreate) SetUpdatedBy(u uuid.UUID) *RbacPolicyCreate {
	rpc.mutation.SetUpdatedBy(u)
	return rpc
}

// SetCreatedBy sets the "created_by" field.
func (rpc *RbacPolicyCreate) SetCreatedBy(u uuid.UUID) *RbacPolicyCreate {
	rpc.mutation.SetCreatedBy(u)
	return rpc
}

// SetRole sets the "role" field.
func (rpc *RbacPolicyCreate) SetRole(s string) *RbacPolicyCreate {
	rpc.mutation.SetRole(s)
	return rpc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (rpc *RbacPolicyCreate) SetNillableRole(s *string) *RbacPolicyCreate {
	if s != nil {
		rpc.SetRole(*s)
	}
	return rpc
}

// SetObj sets the "obj" field.
func (rpc *RbacPolicyCreate) SetObj(s string) *RbacPolicyCreate {
	rpc.mutation.SetObj(s)
	return rpc
}

// SetNillableObj sets the "obj" field if the given value is not nil.
func (rpc *RbacPolicyCreate) SetNillableObj(s *string) *RbacPolicyCreate {
	if s != nil {
		rpc.SetObj(*s)
	}
	return rpc
}

// SetAct sets the "act" field.
func (rpc *RbacPolicyCreate) SetAct(s string) *RbacPolicyCreate {
	rpc.mutation.SetAct(s)
	return rpc
}

// SetNillableAct sets the "act" field if the given value is not nil.
func (rpc *RbacPolicyCreate) SetNillableAct(s *string) *RbacPolicyCreate {
	if s != nil {
		rpc.SetAct(*s)
	}
	return rpc
}

// SetURI sets the "uri" field.
func (rpc *RbacPolicyCreate) SetURI(s string) *RbacPolicyCreate {
	rpc.mutation.SetURI(s)
	return rpc
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (rpc *RbacPolicyCreate) SetNillableURI(s *string) *RbacPolicyCreate {
	if s != nil {
		rpc.SetURI(*s)
	}
	return rpc
}

// SetID sets the "id" field.
func (rpc *RbacPolicyCreate) SetID(i int64) *RbacPolicyCreate {
	rpc.mutation.SetID(i)
	return rpc
}

// Mutation returns the RbacPolicyMutation object of the builder.
func (rpc *RbacPolicyCreate) Mutation() *RbacPolicyMutation {
	return rpc.mutation
}

// Save creates the RbacPolicy in the database.
func (rpc *RbacPolicyCreate) Save(ctx context.Context) (*RbacPolicy, error) {
	rpc.defaults()
	return withHooks(ctx, rpc.sqlSave, rpc.mutation, rpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rpc *RbacPolicyCreate) SaveX(ctx context.Context) *RbacPolicy {
	v, err := rpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpc *RbacPolicyCreate) Exec(ctx context.Context) error {
	_, err := rpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpc *RbacPolicyCreate) ExecX(ctx context.Context) {
	if err := rpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rpc *RbacPolicyCreate) defaults() {
	if _, ok := rpc.mutation.CreatedAt(); !ok {
		v := rbacpolicy.DefaultCreatedAt()
		rpc.mutation.SetCreatedAt(v)
	}
	if _, ok := rpc.mutation.UpdatedAt(); !ok {
		v := rbacpolicy.DefaultUpdatedAt()
		rpc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpc *RbacPolicyCreate) check() error {
	if _, ok := rpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RbacPolicy.created_at"`)}
	}
	if _, ok := rpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RbacPolicy.updated_at"`)}
	}
	if _, ok := rpc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "RbacPolicy.updated_by"`)}
	}
	if _, ok := rpc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "RbacPolicy.created_by"`)}
	}
	if v, ok := rpc.mutation.Role(); ok {
		if err := rbacpolicy.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "RbacPolicy.role": %w`, err)}
		}
	}
	if v, ok := rpc.mutation.Obj(); ok {
		if err := rbacpolicy.ObjValidator(v); err != nil {
			return &ValidationError{Name: "obj", err: fmt.Errorf(`ent: validator failed for field "RbacPolicy.obj": %w`, err)}
		}
	}
	if v, ok := rpc.mutation.Act(); ok {
		if err := rbacpolicy.ActValidator(v); err != nil {
			return &ValidationError{Name: "act", err: fmt.Errorf(`ent: validator failed for field "RbacPolicy.act": %w`, err)}
		}
	}
	if v, ok := rpc.mutation.URI(); ok {
		if err := rbacpolicy.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "RbacPolicy.uri": %w`, err)}
		}
	}
	return nil
}

func (rpc *RbacPolicyCreate) sqlSave(ctx context.Context) (*RbacPolicy, error) {
	if err := rpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rpc.mutation.id = &_node.ID
	rpc.mutation.done = true
	return _node, nil
}

func (rpc *RbacPolicyCreate) createSpec() (*RbacPolicy, *sqlgraph.CreateSpec) {
	var (
		_node = &RbacPolicy{config: rpc.config}
		_spec = sqlgraph.NewCreateSpec(rbacpolicy.Table, sqlgraph.NewFieldSpec(rbacpolicy.FieldID, field.TypeInt64))
	)
	if id, ok := rpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rpc.mutation.CreatedAt(); ok {
		_spec.SetField(rbacpolicy.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rpc.mutation.UpdatedAt(); ok {
		_spec.SetField(rbacpolicy.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rpc.mutation.UpdatedBy(); ok {
		_spec.SetField(rbacpolicy.FieldUpdatedBy, field.TypeUUID, value)
		_node.UpdatedBy = value
	}
	if value, ok := rpc.mutation.CreatedBy(); ok {
		_spec.SetField(rbacpolicy.FieldCreatedBy, field.TypeUUID, value)
		_node.CreatedBy = value
	}
	if value, ok := rpc.mutation.Role(); ok {
		_spec.SetField(rbacpolicy.FieldRole, field.TypeString, value)
		_node.Role = value
	}
	if value, ok := rpc.mutation.Obj(); ok {
		_spec.SetField(rbacpolicy.FieldObj, field.TypeString, value)
		_node.Obj = value
	}
	if value, ok := rpc.mutation.Act(); ok {
		_spec.SetField(rbacpolicy.FieldAct, field.TypeString, value)
		_node.Act = value
	}
	if value, ok := rpc.mutation.URI(); ok {
		_spec.SetField(rbacpolicy.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	return _node, _spec
}

// RbacPolicyCreateBulk is the builder for creating many RbacPolicy entities in bulk.
type RbacPolicyCreateBulk struct {
	config
	err      error
	builders []*RbacPolicyCreate
}

// Save creates the RbacPolicy entities in the database.
func (rpcb *RbacPolicyCreateBulk) Save(ctx context.Context) ([]*RbacPolicy, error) {
	if rpcb.err != nil {
		return nil, rpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rpcb.builders))
	nodes := make([]*RbacPolicy, len(rpcb.builders))
	mutators := make([]Mutator, len(rpcb.builders))
	for i := range rpcb.builders {
		func(i int, root context.Context) {
			builder := rpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RbacPolicyMutation)
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
					_, err = mutators[i+1].Mutate(root, rpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, rpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rpcb *RbacPolicyCreateBulk) SaveX(ctx context.Context) []*RbacPolicy {
	v, err := rpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpcb *RbacPolicyCreateBulk) Exec(ctx context.Context) error {
	_, err := rpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpcb *RbacPolicyCreateBulk) ExecX(ctx context.Context) {
	if err := rpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
