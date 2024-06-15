// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sonar-bat/ent/predicate"
	"sonar-bat/ent/rbacrole"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RbacRoleDelete is the builder for deleting a RbacRole entity.
type RbacRoleDelete struct {
	config
	hooks    []Hook
	mutation *RbacRoleMutation
}

// Where appends a list predicates to the RbacRoleDelete builder.
func (rrd *RbacRoleDelete) Where(ps ...predicate.RbacRole) *RbacRoleDelete {
	rrd.mutation.Where(ps...)
	return rrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rrd *RbacRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rrd.sqlExec, rrd.mutation, rrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rrd *RbacRoleDelete) ExecX(ctx context.Context) int {
	n, err := rrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rrd *RbacRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rbacrole.Table, sqlgraph.NewFieldSpec(rbacrole.FieldID, field.TypeUUID))
	if ps := rrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rrd.mutation.done = true
	return affected, err
}

// RbacRoleDeleteOne is the builder for deleting a single RbacRole entity.
type RbacRoleDeleteOne struct {
	rrd *RbacRoleDelete
}

// Where appends a list predicates to the RbacRoleDelete builder.
func (rrdo *RbacRoleDeleteOne) Where(ps ...predicate.RbacRole) *RbacRoleDeleteOne {
	rrdo.rrd.mutation.Where(ps...)
	return rrdo
}

// Exec executes the deletion query.
func (rrdo *RbacRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := rrdo.rrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rbacrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rrdo *RbacRoleDeleteOne) ExecX(ctx context.Context) {
	if err := rrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
