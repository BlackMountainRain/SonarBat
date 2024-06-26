// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sonar-bat/ent/hostblacklist"
	"sonar-bat/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HostBlacklistDelete is the builder for deleting a HostBlacklist entity.
type HostBlacklistDelete struct {
	config
	hooks    []Hook
	mutation *HostBlacklistMutation
}

// Where appends a list predicates to the HostBlacklistDelete builder.
func (hbd *HostBlacklistDelete) Where(ps ...predicate.HostBlacklist) *HostBlacklistDelete {
	hbd.mutation.Where(ps...)
	return hbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hbd *HostBlacklistDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hbd.sqlExec, hbd.mutation, hbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hbd *HostBlacklistDelete) ExecX(ctx context.Context) int {
	n, err := hbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hbd *HostBlacklistDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hostblacklist.Table, sqlgraph.NewFieldSpec(hostblacklist.FieldID, field.TypeUUID))
	if ps := hbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hbd.mutation.done = true
	return affected, err
}

// HostBlacklistDeleteOne is the builder for deleting a single HostBlacklist entity.
type HostBlacklistDeleteOne struct {
	hbd *HostBlacklistDelete
}

// Where appends a list predicates to the HostBlacklistDelete builder.
func (hbdo *HostBlacklistDeleteOne) Where(ps ...predicate.HostBlacklist) *HostBlacklistDeleteOne {
	hbdo.hbd.mutation.Where(ps...)
	return hbdo
}

// Exec executes the deletion query.
func (hbdo *HostBlacklistDeleteOne) Exec(ctx context.Context) error {
	n, err := hbdo.hbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hostblacklist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hbdo *HostBlacklistDeleteOne) ExecX(ctx context.Context) {
	if err := hbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
