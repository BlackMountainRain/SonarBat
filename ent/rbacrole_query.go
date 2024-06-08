// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"sonar-bat/ent/predicate"
	"sonar-bat/ent/rbacrole"
	"sonar-bat/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RbacRoleQuery is the builder for querying RbacRole entities.
type RbacRoleQuery struct {
	config
	ctx        *QueryContext
	order      []rbacrole.OrderOption
	inters     []Interceptor
	predicates []predicate.RbacRole
	withUsers  *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RbacRoleQuery builder.
func (rrq *RbacRoleQuery) Where(ps ...predicate.RbacRole) *RbacRoleQuery {
	rrq.predicates = append(rrq.predicates, ps...)
	return rrq
}

// Limit the number of records to be returned by this query.
func (rrq *RbacRoleQuery) Limit(limit int) *RbacRoleQuery {
	rrq.ctx.Limit = &limit
	return rrq
}

// Offset to start from.
func (rrq *RbacRoleQuery) Offset(offset int) *RbacRoleQuery {
	rrq.ctx.Offset = &offset
	return rrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rrq *RbacRoleQuery) Unique(unique bool) *RbacRoleQuery {
	rrq.ctx.Unique = &unique
	return rrq
}

// Order specifies how the records should be ordered.
func (rrq *RbacRoleQuery) Order(o ...rbacrole.OrderOption) *RbacRoleQuery {
	rrq.order = append(rrq.order, o...)
	return rrq
}

// QueryUsers chains the current query on the "users" edge.
func (rrq *RbacRoleQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: rrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rbacrole.Table, rbacrole.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, rbacrole.UsersTable, rbacrole.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RbacRole entity from the query.
// Returns a *NotFoundError when no RbacRole was found.
func (rrq *RbacRoleQuery) First(ctx context.Context) (*RbacRole, error) {
	nodes, err := rrq.Limit(1).All(setContextOp(ctx, rrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rbacrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rrq *RbacRoleQuery) FirstX(ctx context.Context) *RbacRole {
	node, err := rrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RbacRole ID from the query.
// Returns a *NotFoundError when no RbacRole ID was found.
func (rrq *RbacRoleQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rrq.Limit(1).IDs(setContextOp(ctx, rrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rbacrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rrq *RbacRoleQuery) FirstIDX(ctx context.Context) int64 {
	id, err := rrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RbacRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RbacRole entity is found.
// Returns a *NotFoundError when no RbacRole entities are found.
func (rrq *RbacRoleQuery) Only(ctx context.Context) (*RbacRole, error) {
	nodes, err := rrq.Limit(2).All(setContextOp(ctx, rrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rbacrole.Label}
	default:
		return nil, &NotSingularError{rbacrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rrq *RbacRoleQuery) OnlyX(ctx context.Context) *RbacRole {
	node, err := rrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RbacRole ID in the query.
// Returns a *NotSingularError when more than one RbacRole ID is found.
// Returns a *NotFoundError when no entities are found.
func (rrq *RbacRoleQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rrq.Limit(2).IDs(setContextOp(ctx, rrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rbacrole.Label}
	default:
		err = &NotSingularError{rbacrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rrq *RbacRoleQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := rrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RbacRoles.
func (rrq *RbacRoleQuery) All(ctx context.Context) ([]*RbacRole, error) {
	ctx = setContextOp(ctx, rrq.ctx, "All")
	if err := rrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RbacRole, *RbacRoleQuery]()
	return withInterceptors[[]*RbacRole](ctx, rrq, qr, rrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rrq *RbacRoleQuery) AllX(ctx context.Context) []*RbacRole {
	nodes, err := rrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RbacRole IDs.
func (rrq *RbacRoleQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if rrq.ctx.Unique == nil && rrq.path != nil {
		rrq.Unique(true)
	}
	ctx = setContextOp(ctx, rrq.ctx, "IDs")
	if err = rrq.Select(rbacrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rrq *RbacRoleQuery) IDsX(ctx context.Context) []int64 {
	ids, err := rrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rrq *RbacRoleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rrq.ctx, "Count")
	if err := rrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rrq, querierCount[*RbacRoleQuery](), rrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rrq *RbacRoleQuery) CountX(ctx context.Context) int {
	count, err := rrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rrq *RbacRoleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rrq.ctx, "Exist")
	switch _, err := rrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rrq *RbacRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := rrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RbacRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rrq *RbacRoleQuery) Clone() *RbacRoleQuery {
	if rrq == nil {
		return nil
	}
	return &RbacRoleQuery{
		config:     rrq.config,
		ctx:        rrq.ctx.Clone(),
		order:      append([]rbacrole.OrderOption{}, rrq.order...),
		inters:     append([]Interceptor{}, rrq.inters...),
		predicates: append([]predicate.RbacRole{}, rrq.predicates...),
		withUsers:  rrq.withUsers.Clone(),
		// clone intermediate query.
		sql:  rrq.sql.Clone(),
		path: rrq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (rrq *RbacRoleQuery) WithUsers(opts ...func(*UserQuery)) *RbacRoleQuery {
	query := (&UserClient{config: rrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rrq.withUsers = query
	return rrq
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
//	client.RbacRole.Query().
//		GroupBy(rbacrole.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rrq *RbacRoleQuery) GroupBy(field string, fields ...string) *RbacRoleGroupBy {
	rrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RbacRoleGroupBy{build: rrq}
	grbuild.flds = &rrq.ctx.Fields
	grbuild.label = rbacrole.Label
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
//	client.RbacRole.Query().
//		Select(rbacrole.FieldCreatedAt).
//		Scan(ctx, &v)
func (rrq *RbacRoleQuery) Select(fields ...string) *RbacRoleSelect {
	rrq.ctx.Fields = append(rrq.ctx.Fields, fields...)
	sbuild := &RbacRoleSelect{RbacRoleQuery: rrq}
	sbuild.label = rbacrole.Label
	sbuild.flds, sbuild.scan = &rrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RbacRoleSelect configured with the given aggregations.
func (rrq *RbacRoleQuery) Aggregate(fns ...AggregateFunc) *RbacRoleSelect {
	return rrq.Select().Aggregate(fns...)
}

func (rrq *RbacRoleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rrq); err != nil {
				return err
			}
		}
	}
	for _, f := range rrq.ctx.Fields {
		if !rbacrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rrq.path != nil {
		prev, err := rrq.path(ctx)
		if err != nil {
			return err
		}
		rrq.sql = prev
	}
	return nil
}

func (rrq *RbacRoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RbacRole, error) {
	var (
		nodes       = []*RbacRole{}
		_spec       = rrq.querySpec()
		loadedTypes = [1]bool{
			rrq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RbacRole).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RbacRole{config: rrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rrq.withUsers; query != nil {
		if err := rrq.loadUsers(ctx, query, nodes,
			func(n *RbacRole) { n.Edges.Users = []*User{} },
			func(n *RbacRole, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rrq *RbacRoleQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*RbacRole, init func(*RbacRole), assign func(*RbacRole, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*RbacRole)
	nids := make(map[int64]map[*RbacRole]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(rbacrole.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(rbacrole.UsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(rbacrole.UsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(rbacrole.UsersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*RbacRole]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (rrq *RbacRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rrq.querySpec()
	_spec.Node.Columns = rrq.ctx.Fields
	if len(rrq.ctx.Fields) > 0 {
		_spec.Unique = rrq.ctx.Unique != nil && *rrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rrq.driver, _spec)
}

func (rrq *RbacRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rbacrole.Table, rbacrole.Columns, sqlgraph.NewFieldSpec(rbacrole.FieldID, field.TypeInt64))
	_spec.From = rrq.sql
	if unique := rrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rrq.path != nil {
		_spec.Unique = true
	}
	if fields := rrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rbacrole.FieldID)
		for i := range fields {
			if fields[i] != rbacrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rrq *RbacRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rrq.driver.Dialect())
	t1 := builder.Table(rbacrole.Table)
	columns := rrq.ctx.Fields
	if len(columns) == 0 {
		columns = rbacrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rrq.sql != nil {
		selector = rrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rrq.ctx.Unique != nil && *rrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rrq.predicates {
		p(selector)
	}
	for _, p := range rrq.order {
		p(selector)
	}
	if offset := rrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RbacRoleGroupBy is the group-by builder for RbacRole entities.
type RbacRoleGroupBy struct {
	selector
	build *RbacRoleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rrgb *RbacRoleGroupBy) Aggregate(fns ...AggregateFunc) *RbacRoleGroupBy {
	rrgb.fns = append(rrgb.fns, fns...)
	return rrgb
}

// Scan applies the selector query and scans the result into the given value.
func (rrgb *RbacRoleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rrgb.build.ctx, "GroupBy")
	if err := rrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RbacRoleQuery, *RbacRoleGroupBy](ctx, rrgb.build, rrgb, rrgb.build.inters, v)
}

func (rrgb *RbacRoleGroupBy) sqlScan(ctx context.Context, root *RbacRoleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rrgb.fns))
	for _, fn := range rrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rrgb.flds)+len(rrgb.fns))
		for _, f := range *rrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RbacRoleSelect is the builder for selecting fields of RbacRole entities.
type RbacRoleSelect struct {
	*RbacRoleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rrs *RbacRoleSelect) Aggregate(fns ...AggregateFunc) *RbacRoleSelect {
	rrs.fns = append(rrs.fns, fns...)
	return rrs
}

// Scan applies the selector query and scans the result into the given value.
func (rrs *RbacRoleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rrs.ctx, "Select")
	if err := rrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RbacRoleQuery, *RbacRoleSelect](ctx, rrs.RbacRoleQuery, rrs, rrs.inters, v)
}

func (rrs *RbacRoleSelect) sqlScan(ctx context.Context, root *RbacRoleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rrs.fns))
	for _, fn := range rrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
