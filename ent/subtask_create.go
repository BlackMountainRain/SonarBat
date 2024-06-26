// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sonar-bat/ent/subtask"
	"sonar-bat/ent/task"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SubtaskCreate is the builder for creating a Subtask entity.
type SubtaskCreate struct {
	config
	mutation *SubtaskMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SubtaskCreate) SetCreatedAt(t time.Time) *SubtaskCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SubtaskCreate) SetNillableCreatedAt(t *time.Time) *SubtaskCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SubtaskCreate) SetUpdatedAt(t time.Time) *SubtaskCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SubtaskCreate) SetNillableUpdatedAt(t *time.Time) *SubtaskCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *SubtaskCreate) SetUpdatedBy(u uuid.UUID) *SubtaskCreate {
	sc.mutation.SetUpdatedBy(u)
	return sc
}

// SetCreatedBy sets the "created_by" field.
func (sc *SubtaskCreate) SetCreatedBy(u uuid.UUID) *SubtaskCreate {
	sc.mutation.SetCreatedBy(u)
	return sc
}

// SetTaskID sets the "task_id" field.
func (sc *SubtaskCreate) SetTaskID(u uuid.UUID) *SubtaskCreate {
	sc.mutation.SetTaskID(u)
	return sc
}

// SetStatus sets the "status" field.
func (sc *SubtaskCreate) SetStatus(b bool) *SubtaskCreate {
	sc.mutation.SetStatus(b)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *SubtaskCreate) SetNillableStatus(b *bool) *SubtaskCreate {
	if b != nil {
		sc.SetStatus(*b)
	}
	return sc
}

// SetDetectionType sets the "detection_type" field.
func (sc *SubtaskCreate) SetDetectionType(i int16) *SubtaskCreate {
	sc.mutation.SetDetectionType(i)
	return sc
}

// SetParams sets the "params" field.
func (sc *SubtaskCreate) SetParams(m map[string]interface{}) *SubtaskCreate {
	sc.mutation.SetParams(m)
	return sc
}

// SetSrcEpFilterStrategy sets the "src_ep_filter_strategy" field.
func (sc *SubtaskCreate) SetSrcEpFilterStrategy(m map[string]interface{}) *SubtaskCreate {
	sc.mutation.SetSrcEpFilterStrategy(m)
	return sc
}

// SetSrcEpSelStrategy sets the "src_ep_sel_strategy" field.
func (sc *SubtaskCreate) SetSrcEpSelStrategy(m map[string]interface{}) *SubtaskCreate {
	sc.mutation.SetSrcEpSelStrategy(m)
	return sc
}

// SetSrcEpSelNum sets the "src_ep_sel_num" field.
func (sc *SubtaskCreate) SetSrcEpSelNum(i int) *SubtaskCreate {
	sc.mutation.SetSrcEpSelNum(i)
	return sc
}

// SetDstEpFilterStrategy sets the "dst_ep_filter_strategy" field.
func (sc *SubtaskCreate) SetDstEpFilterStrategy(m map[string]interface{}) *SubtaskCreate {
	sc.mutation.SetDstEpFilterStrategy(m)
	return sc
}

// SetDstEpSelStrategy sets the "dst_ep_sel_strategy" field.
func (sc *SubtaskCreate) SetDstEpSelStrategy(m map[string]interface{}) *SubtaskCreate {
	sc.mutation.SetDstEpSelStrategy(m)
	return sc
}

// SetDstEpSelNum sets the "dst_ep_sel_num" field.
func (sc *SubtaskCreate) SetDstEpSelNum(i int) *SubtaskCreate {
	sc.mutation.SetDstEpSelNum(i)
	return sc
}

// SetID sets the "id" field.
func (sc *SubtaskCreate) SetID(u uuid.UUID) *SubtaskCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SubtaskCreate) SetNillableID(u *uuid.UUID) *SubtaskCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetTask sets the "task" edge to the Task entity.
func (sc *SubtaskCreate) SetTask(t *Task) *SubtaskCreate {
	return sc.SetTaskID(t.ID)
}

// Mutation returns the SubtaskMutation object of the builder.
func (sc *SubtaskCreate) Mutation() *SubtaskMutation {
	return sc.mutation
}

// Save creates the Subtask in the database.
func (sc *SubtaskCreate) Save(ctx context.Context) (*Subtask, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubtaskCreate) SaveX(ctx context.Context) *Subtask {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubtaskCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubtaskCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubtaskCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := subtask.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := subtask.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := subtask.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := subtask.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubtaskCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Subtask.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Subtask.updated_at"`)}
	}
	if _, ok := sc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Subtask.updated_by"`)}
	}
	if _, ok := sc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Subtask.created_by"`)}
	}
	if _, ok := sc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "Subtask.task_id"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Subtask.status"`)}
	}
	if _, ok := sc.mutation.DetectionType(); !ok {
		return &ValidationError{Name: "detection_type", err: errors.New(`ent: missing required field "Subtask.detection_type"`)}
	}
	if _, ok := sc.mutation.Params(); !ok {
		return &ValidationError{Name: "params", err: errors.New(`ent: missing required field "Subtask.params"`)}
	}
	if _, ok := sc.mutation.SrcEpFilterStrategy(); !ok {
		return &ValidationError{Name: "src_ep_filter_strategy", err: errors.New(`ent: missing required field "Subtask.src_ep_filter_strategy"`)}
	}
	if _, ok := sc.mutation.SrcEpSelStrategy(); !ok {
		return &ValidationError{Name: "src_ep_sel_strategy", err: errors.New(`ent: missing required field "Subtask.src_ep_sel_strategy"`)}
	}
	if _, ok := sc.mutation.SrcEpSelNum(); !ok {
		return &ValidationError{Name: "src_ep_sel_num", err: errors.New(`ent: missing required field "Subtask.src_ep_sel_num"`)}
	}
	if _, ok := sc.mutation.DstEpFilterStrategy(); !ok {
		return &ValidationError{Name: "dst_ep_filter_strategy", err: errors.New(`ent: missing required field "Subtask.dst_ep_filter_strategy"`)}
	}
	if _, ok := sc.mutation.DstEpSelStrategy(); !ok {
		return &ValidationError{Name: "dst_ep_sel_strategy", err: errors.New(`ent: missing required field "Subtask.dst_ep_sel_strategy"`)}
	}
	if _, ok := sc.mutation.DstEpSelNum(); !ok {
		return &ValidationError{Name: "dst_ep_sel_num", err: errors.New(`ent: missing required field "Subtask.dst_ep_sel_num"`)}
	}
	if _, ok := sc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "Subtask.task"`)}
	}
	return nil
}

func (sc *SubtaskCreate) sqlSave(ctx context.Context) (*Subtask, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubtaskCreate) createSpec() (*Subtask, *sqlgraph.CreateSpec) {
	var (
		_node = &Subtask{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subtask.Table, sqlgraph.NewFieldSpec(subtask.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(subtask.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(subtask.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.SetField(subtask.FieldUpdatedBy, field.TypeUUID, value)
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.CreatedBy(); ok {
		_spec.SetField(subtask.FieldCreatedBy, field.TypeUUID, value)
		_node.CreatedBy = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(subtask.FieldStatus, field.TypeBool, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.DetectionType(); ok {
		_spec.SetField(subtask.FieldDetectionType, field.TypeInt16, value)
		_node.DetectionType = value
	}
	if value, ok := sc.mutation.Params(); ok {
		_spec.SetField(subtask.FieldParams, field.TypeJSON, value)
		_node.Params = value
	}
	if value, ok := sc.mutation.SrcEpFilterStrategy(); ok {
		_spec.SetField(subtask.FieldSrcEpFilterStrategy, field.TypeJSON, value)
		_node.SrcEpFilterStrategy = value
	}
	if value, ok := sc.mutation.SrcEpSelStrategy(); ok {
		_spec.SetField(subtask.FieldSrcEpSelStrategy, field.TypeJSON, value)
		_node.SrcEpSelStrategy = value
	}
	if value, ok := sc.mutation.SrcEpSelNum(); ok {
		_spec.SetField(subtask.FieldSrcEpSelNum, field.TypeInt, value)
		_node.SrcEpSelNum = value
	}
	if value, ok := sc.mutation.DstEpFilterStrategy(); ok {
		_spec.SetField(subtask.FieldDstEpFilterStrategy, field.TypeJSON, value)
		_node.DstEpFilterStrategy = value
	}
	if value, ok := sc.mutation.DstEpSelStrategy(); ok {
		_spec.SetField(subtask.FieldDstEpSelStrategy, field.TypeJSON, value)
		_node.DstEpSelStrategy = value
	}
	if value, ok := sc.mutation.DstEpSelNum(); ok {
		_spec.SetField(subtask.FieldDstEpSelNum, field.TypeInt, value)
		_node.DstEpSelNum = value
	}
	if nodes := sc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subtask.TaskTable,
			Columns: []string{subtask.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TaskID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubtaskCreateBulk is the builder for creating many Subtask entities in bulk.
type SubtaskCreateBulk struct {
	config
	err      error
	builders []*SubtaskCreate
}

// Save creates the Subtask entities in the database.
func (scb *SubtaskCreateBulk) Save(ctx context.Context) ([]*Subtask, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subtask, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubtaskMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubtaskCreateBulk) SaveX(ctx context.Context) []*Subtask {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubtaskCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubtaskCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
