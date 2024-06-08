// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"sonar-bat/ent/predicate"
	"sonar-bat/ent/rbacpolicy"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RbacPolicyQuery is the builder for querying RbacPolicy entities.
type RbacPolicyQuery struct {
	config
	ctx        *QueryContext
	order      []rbacpolicy.OrderOption
	inters     []Interceptor
	predicates []predicate.RbacPolicy
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RbacPolicyQuery builder.
func (rpq *RbacPolicyQuery) Where(ps ...predicate.RbacPolicy) *RbacPolicyQuery {
	rpq.predicates = append(rpq.predicates, ps...)
	return rpq
}

// Limit the number of records to be returned by this query.
func (rpq *RbacPolicyQuery) Limit(limit int) *RbacPolicyQuery {
	rpq.ctx.Limit = &limit
	return rpq
}

// Offset to start from.
func (rpq *RbacPolicyQuery) Offset(offset int) *RbacPolicyQuery {
	rpq.ctx.Offset = &offset
	return rpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rpq *RbacPolicyQuery) Unique(unique bool) *RbacPolicyQuery {
	rpq.ctx.Unique = &unique
	return rpq
}

// Order specifies how the records should be ordered.
func (rpq *RbacPolicyQuery) Order(o ...rbacpolicy.OrderOption) *RbacPolicyQuery {
	rpq.order = append(rpq.order, o...)
	return rpq
}

// First returns the first RbacPolicy entity from the query.
// Returns a *NotFoundError when no RbacPolicy was found.
func (rpq *RbacPolicyQuery) First(ctx context.Context) (*RbacPolicy, error) {
	nodes, err := rpq.Limit(1).All(setContextOp(ctx, rpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rbacpolicy.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rpq *RbacPolicyQuery) FirstX(ctx context.Context) *RbacPolicy {
	node, err := rpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RbacPolicy ID from the query.
// Returns a *NotFoundError when no RbacPolicy ID was found.
func (rpq *RbacPolicyQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rpq.Limit(1).IDs(setContextOp(ctx, rpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rbacpolicy.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rpq *RbacPolicyQuery) FirstIDX(ctx context.Context) int64 {
	id, err := rpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RbacPolicy entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RbacPolicy entity is found.
// Returns a *NotFoundError when no RbacPolicy entities are found.
func (rpq *RbacPolicyQuery) Only(ctx context.Context) (*RbacPolicy, error) {
	nodes, err := rpq.Limit(2).All(setContextOp(ctx, rpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rbacpolicy.Label}
	default:
		return nil, &NotSingularError{rbacpolicy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rpq *RbacPolicyQuery) OnlyX(ctx context.Context) *RbacPolicy {
	node, err := rpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RbacPolicy ID in the query.
// Returns a *NotSingularError when more than one RbacPolicy ID is found.
// Returns a *NotFoundError when no entities are found.
func (rpq *RbacPolicyQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rpq.Limit(2).IDs(setContextOp(ctx, rpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rbacpolicy.Label}
	default:
		err = &NotSingularError{rbacpolicy.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rpq *RbacPolicyQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := rpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RbacPolicies.
func (rpq *RbacPolicyQuery) All(ctx context.Context) ([]*RbacPolicy, error) {
	ctx = setContextOp(ctx, rpq.ctx, "All")
	if err := rpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RbacPolicy, *RbacPolicyQuery]()
	return withInterceptors[[]*RbacPolicy](ctx, rpq, qr, rpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rpq *RbacPolicyQuery) AllX(ctx context.Context) []*RbacPolicy {
	nodes, err := rpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RbacPolicy IDs.
func (rpq *RbacPolicyQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if rpq.ctx.Unique == nil && rpq.path != nil {
		rpq.Unique(true)
	}
	ctx = setContextOp(ctx, rpq.ctx, "IDs")
	if err = rpq.Select(rbacpolicy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rpq *RbacPolicyQuery) IDsX(ctx context.Context) []int64 {
	ids, err := rpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rpq *RbacPolicyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Count")
	if err := rpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rpq, querierCount[*RbacPolicyQuery](), rpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rpq *RbacPolicyQuery) CountX(ctx context.Context) int {
	count, err := rpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rpq *RbacPolicyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Exist")
	switch _, err := rpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rpq *RbacPolicyQuery) ExistX(ctx context.Context) bool {
	exist, err := rpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RbacPolicyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rpq *RbacPolicyQuery) Clone() *RbacPolicyQuery {
	if rpq == nil {
		return nil
	}
	return &RbacPolicyQuery{
		config:     rpq.config,
		ctx:        rpq.ctx.Clone(),
		order:      append([]rbacpolicy.OrderOption{}, rpq.order...),
		inters:     append([]Interceptor{}, rpq.inters...),
		predicates: append([]predicate.RbacPolicy{}, rpq.predicates...),
		// clone intermediate query.
		sql:  rpq.sql.Clone(),
		path: rpq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RbacPolicy.Query().
//		GroupBy(rbacpolicy.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rpq *RbacPolicyQuery) GroupBy(field string, fields ...string) *RbacPolicyGroupBy {
	rpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RbacPolicyGroupBy{build: rpq}
	grbuild.flds = &rpq.ctx.Fields
	grbuild.label = rbacpolicy.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.RbacPolicy.Query().
//		Select(rbacpolicy.FieldCreatedAt).
//		Scan(ctx, &v)
func (rpq *RbacPolicyQuery) Select(fields ...string) *RbacPolicySelect {
	rpq.ctx.Fields = append(rpq.ctx.Fields, fields...)
	sbuild := &RbacPolicySelect{RbacPolicyQuery: rpq}
	sbuild.label = rbacpolicy.Label
	sbuild.flds, sbuild.scan = &rpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RbacPolicySelect configured with the given aggregations.
func (rpq *RbacPolicyQuery) Aggregate(fns ...AggregateFunc) *RbacPolicySelect {
	return rpq.Select().Aggregate(fns...)
}

func (rpq *RbacPolicyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rpq); err != nil {
				return err
			}
		}
	}
	for _, f := range rpq.ctx.Fields {
		if !rbacpolicy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rpq.path != nil {
		prev, err := rpq.path(ctx)
		if err != nil {
			return err
		}
		rpq.sql = prev
	}
	return nil
}

func (rpq *RbacPolicyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RbacPolicy, error) {
	var (
		nodes = []*RbacPolicy{}
		_spec = rpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RbacPolicy).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RbacPolicy{config: rpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rpq *RbacPolicyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rpq.querySpec()
	_spec.Node.Columns = rpq.ctx.Fields
	if len(rpq.ctx.Fields) > 0 {
		_spec.Unique = rpq.ctx.Unique != nil && *rpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rpq.driver, _spec)
}

func (rpq *RbacPolicyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rbacpolicy.Table, rbacpolicy.Columns, sqlgraph.NewFieldSpec(rbacpolicy.FieldID, field.TypeInt64))
	_spec.From = rpq.sql
	if unique := rpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rpq.path != nil {
		_spec.Unique = true
	}
	if fields := rpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rbacpolicy.FieldID)
		for i := range fields {
			if fields[i] != rbacpolicy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rpq *RbacPolicyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rpq.driver.Dialect())
	t1 := builder.Table(rbacpolicy.Table)
	columns := rpq.ctx.Fields
	if len(columns) == 0 {
		columns = rbacpolicy.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rpq.sql != nil {
		selector = rpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rpq.ctx.Unique != nil && *rpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rpq.predicates {
		p(selector)
	}
	for _, p := range rpq.order {
		p(selector)
	}
	if offset := rpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RbacPolicyGroupBy is the group-by builder for RbacPolicy entities.
type RbacPolicyGroupBy struct {
	selector
	build *RbacPolicyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rpgb *RbacPolicyGroupBy) Aggregate(fns ...AggregateFunc) *RbacPolicyGroupBy {
	rpgb.fns = append(rpgb.fns, fns...)
	return rpgb
}

// Scan applies the selector query and scans the result into the given value.
func (rpgb *RbacPolicyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rpgb.build.ctx, "GroupBy")
	if err := rpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RbacPolicyQuery, *RbacPolicyGroupBy](ctx, rpgb.build, rpgb, rpgb.build.inters, v)
}

func (rpgb *RbacPolicyGroupBy) sqlScan(ctx context.Context, root *RbacPolicyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rpgb.fns))
	for _, fn := range rpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rpgb.flds)+len(rpgb.fns))
		for _, f := range *rpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RbacPolicySelect is the builder for selecting fields of RbacPolicy entities.
type RbacPolicySelect struct {
	*RbacPolicyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rps *RbacPolicySelect) Aggregate(fns ...AggregateFunc) *RbacPolicySelect {
	rps.fns = append(rps.fns, fns...)
	return rps
}

// Scan applies the selector query and scans the result into the given value.
func (rps *RbacPolicySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rps.ctx, "Select")
	if err := rps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RbacPolicyQuery, *RbacPolicySelect](ctx, rps.RbacPolicyQuery, rps, rps.inters, v)
}

func (rps *RbacPolicySelect) sqlScan(ctx context.Context, root *RbacPolicyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rps.fns))
	for _, fn := range rps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
