// Code generated by ent, DO NOT EDIT.

package host

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the host type in the database.
	Label = "host"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLiveAt holds the string denoting the live_at field in the database.
	FieldLiveAt = "live_at"
	// FieldIsAgentInstalled holds the string denoting the is_agent_installed field in the database.
	FieldIsAgentInstalled = "is_agent_installed"
	// FieldAgentVersion holds the string denoting the agent_version field in the database.
	FieldAgentVersion = "agent_version"
	// FieldIps holds the string denoting the ips field in the database.
	FieldIps = "ips"
	// FieldNetType holds the string denoting the net_type field in the database.
	FieldNetType = "net_type"
	// FieldAdditions holds the string denoting the additions field in the database.
	FieldAdditions = "additions"
	// EdgeHostBlacklist holds the string denoting the host_blacklist edge name in mutations.
	EdgeHostBlacklist = "host_blacklist"
	// Table holds the table name of the host in the database.
	Table = "hosts"
	// HostBlacklistTable is the table that holds the host_blacklist relation/edge.
	HostBlacklistTable = "host_blacklists"
	// HostBlacklistInverseTable is the table name for the HostBlacklist entity.
	// It exists in this package in order to avoid circular dependency with the "hostblacklist" package.
	HostBlacklistInverseTable = "host_blacklists"
	// HostBlacklistColumn is the table column denoting the host_blacklist relation/edge.
	HostBlacklistColumn = "host_id"
)

// Columns holds all SQL columns for host fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldCreatedBy,
	FieldStatus,
	FieldName,
	FieldLiveAt,
	FieldIsAgentInstalled,
	FieldAgentVersion,
	FieldIps,
	FieldNetType,
	FieldAdditions,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultIsAgentInstalled holds the default value on creation for the "is_agent_installed" field.
	DefaultIsAgentInstalled bool
	// AgentVersionValidator is a validator for the "agent_version" field. It is called by the builders before save.
	AgentVersionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Host queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLiveAt orders the results by the live_at field.
func ByLiveAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLiveAt, opts...).ToFunc()
}

// ByIsAgentInstalled orders the results by the is_agent_installed field.
func ByIsAgentInstalled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAgentInstalled, opts...).ToFunc()
}

// ByAgentVersion orders the results by the agent_version field.
func ByAgentVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAgentVersion, opts...).ToFunc()
}

// ByNetType orders the results by the net_type field.
func ByNetType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNetType, opts...).ToFunc()
}

// ByHostBlacklistCount orders the results by host_blacklist count.
func ByHostBlacklistCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHostBlacklistStep(), opts...)
	}
}

// ByHostBlacklist orders the results by host_blacklist terms.
func ByHostBlacklist(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHostBlacklistStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newHostBlacklistStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HostBlacklistInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HostBlacklistTable, HostBlacklistColumn),
	)
}
