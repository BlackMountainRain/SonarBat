// Code generated by ent, DO NOT EDIT.

package host

import (
	"sonar-bat/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldCreatedBy, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldStatus, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldName, v))
}

// LiveAt applies equality check predicate on the "live_at" field. It's identical to LiveAtEQ.
func LiveAt(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldLiveAt, v))
}

// IsAgentInstalled applies equality check predicate on the "is_agent_installed" field. It's identical to IsAgentInstalledEQ.
func IsAgentInstalled(v bool) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldIsAgentInstalled, v))
}

// AgentVersion applies equality check predicate on the "agent_version" field. It's identical to AgentVersionEQ.
func AgentVersion(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldAgentVersion, v))
}

// NetType applies equality check predicate on the "net_type" field. It's identical to NetTypeEQ.
func NetType(v int16) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldNetType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uuid.UUID) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldCreatedBy, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldStatus, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Host {
	return predicate.Host(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Host {
	return predicate.Host(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Host {
	return predicate.Host(sql.FieldContainsFold(FieldName, v))
}

// LiveAtEQ applies the EQ predicate on the "live_at" field.
func LiveAtEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldLiveAt, v))
}

// LiveAtNEQ applies the NEQ predicate on the "live_at" field.
func LiveAtNEQ(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldLiveAt, v))
}

// LiveAtIn applies the In predicate on the "live_at" field.
func LiveAtIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldLiveAt, vs...))
}

// LiveAtNotIn applies the NotIn predicate on the "live_at" field.
func LiveAtNotIn(vs ...time.Time) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldLiveAt, vs...))
}

// LiveAtGT applies the GT predicate on the "live_at" field.
func LiveAtGT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldLiveAt, v))
}

// LiveAtGTE applies the GTE predicate on the "live_at" field.
func LiveAtGTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldLiveAt, v))
}

// LiveAtLT applies the LT predicate on the "live_at" field.
func LiveAtLT(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldLiveAt, v))
}

// LiveAtLTE applies the LTE predicate on the "live_at" field.
func LiveAtLTE(v time.Time) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldLiveAt, v))
}

// LiveAtIsNil applies the IsNil predicate on the "live_at" field.
func LiveAtIsNil() predicate.Host {
	return predicate.Host(sql.FieldIsNull(FieldLiveAt))
}

// LiveAtNotNil applies the NotNil predicate on the "live_at" field.
func LiveAtNotNil() predicate.Host {
	return predicate.Host(sql.FieldNotNull(FieldLiveAt))
}

// IsAgentInstalledEQ applies the EQ predicate on the "is_agent_installed" field.
func IsAgentInstalledEQ(v bool) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldIsAgentInstalled, v))
}

// IsAgentInstalledNEQ applies the NEQ predicate on the "is_agent_installed" field.
func IsAgentInstalledNEQ(v bool) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldIsAgentInstalled, v))
}

// AgentVersionEQ applies the EQ predicate on the "agent_version" field.
func AgentVersionEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldAgentVersion, v))
}

// AgentVersionNEQ applies the NEQ predicate on the "agent_version" field.
func AgentVersionNEQ(v string) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldAgentVersion, v))
}

// AgentVersionIn applies the In predicate on the "agent_version" field.
func AgentVersionIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldAgentVersion, vs...))
}

// AgentVersionNotIn applies the NotIn predicate on the "agent_version" field.
func AgentVersionNotIn(vs ...string) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldAgentVersion, vs...))
}

// AgentVersionGT applies the GT predicate on the "agent_version" field.
func AgentVersionGT(v string) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldAgentVersion, v))
}

// AgentVersionGTE applies the GTE predicate on the "agent_version" field.
func AgentVersionGTE(v string) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldAgentVersion, v))
}

// AgentVersionLT applies the LT predicate on the "agent_version" field.
func AgentVersionLT(v string) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldAgentVersion, v))
}

// AgentVersionLTE applies the LTE predicate on the "agent_version" field.
func AgentVersionLTE(v string) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldAgentVersion, v))
}

// AgentVersionContains applies the Contains predicate on the "agent_version" field.
func AgentVersionContains(v string) predicate.Host {
	return predicate.Host(sql.FieldContains(FieldAgentVersion, v))
}

// AgentVersionHasPrefix applies the HasPrefix predicate on the "agent_version" field.
func AgentVersionHasPrefix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasPrefix(FieldAgentVersion, v))
}

// AgentVersionHasSuffix applies the HasSuffix predicate on the "agent_version" field.
func AgentVersionHasSuffix(v string) predicate.Host {
	return predicate.Host(sql.FieldHasSuffix(FieldAgentVersion, v))
}

// AgentVersionIsNil applies the IsNil predicate on the "agent_version" field.
func AgentVersionIsNil() predicate.Host {
	return predicate.Host(sql.FieldIsNull(FieldAgentVersion))
}

// AgentVersionNotNil applies the NotNil predicate on the "agent_version" field.
func AgentVersionNotNil() predicate.Host {
	return predicate.Host(sql.FieldNotNull(FieldAgentVersion))
}

// AgentVersionEqualFold applies the EqualFold predicate on the "agent_version" field.
func AgentVersionEqualFold(v string) predicate.Host {
	return predicate.Host(sql.FieldEqualFold(FieldAgentVersion, v))
}

// AgentVersionContainsFold applies the ContainsFold predicate on the "agent_version" field.
func AgentVersionContainsFold(v string) predicate.Host {
	return predicate.Host(sql.FieldContainsFold(FieldAgentVersion, v))
}

// IpsIsNil applies the IsNil predicate on the "ips" field.
func IpsIsNil() predicate.Host {
	return predicate.Host(sql.FieldIsNull(FieldIps))
}

// IpsNotNil applies the NotNil predicate on the "ips" field.
func IpsNotNil() predicate.Host {
	return predicate.Host(sql.FieldNotNull(FieldIps))
}

// NetTypeEQ applies the EQ predicate on the "net_type" field.
func NetTypeEQ(v int16) predicate.Host {
	return predicate.Host(sql.FieldEQ(FieldNetType, v))
}

// NetTypeNEQ applies the NEQ predicate on the "net_type" field.
func NetTypeNEQ(v int16) predicate.Host {
	return predicate.Host(sql.FieldNEQ(FieldNetType, v))
}

// NetTypeIn applies the In predicate on the "net_type" field.
func NetTypeIn(vs ...int16) predicate.Host {
	return predicate.Host(sql.FieldIn(FieldNetType, vs...))
}

// NetTypeNotIn applies the NotIn predicate on the "net_type" field.
func NetTypeNotIn(vs ...int16) predicate.Host {
	return predicate.Host(sql.FieldNotIn(FieldNetType, vs...))
}

// NetTypeGT applies the GT predicate on the "net_type" field.
func NetTypeGT(v int16) predicate.Host {
	return predicate.Host(sql.FieldGT(FieldNetType, v))
}

// NetTypeGTE applies the GTE predicate on the "net_type" field.
func NetTypeGTE(v int16) predicate.Host {
	return predicate.Host(sql.FieldGTE(FieldNetType, v))
}

// NetTypeLT applies the LT predicate on the "net_type" field.
func NetTypeLT(v int16) predicate.Host {
	return predicate.Host(sql.FieldLT(FieldNetType, v))
}

// NetTypeLTE applies the LTE predicate on the "net_type" field.
func NetTypeLTE(v int16) predicate.Host {
	return predicate.Host(sql.FieldLTE(FieldNetType, v))
}

// NetTypeIsNil applies the IsNil predicate on the "net_type" field.
func NetTypeIsNil() predicate.Host {
	return predicate.Host(sql.FieldIsNull(FieldNetType))
}

// NetTypeNotNil applies the NotNil predicate on the "net_type" field.
func NetTypeNotNil() predicate.Host {
	return predicate.Host(sql.FieldNotNull(FieldNetType))
}

// AdditionsIsNil applies the IsNil predicate on the "additions" field.
func AdditionsIsNil() predicate.Host {
	return predicate.Host(sql.FieldIsNull(FieldAdditions))
}

// AdditionsNotNil applies the NotNil predicate on the "additions" field.
func AdditionsNotNil() predicate.Host {
	return predicate.Host(sql.FieldNotNull(FieldAdditions))
}

// HasHostBlacklist applies the HasEdge predicate on the "host_blacklist" edge.
func HasHostBlacklist() predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HostBlacklistTable, HostBlacklistColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostBlacklistWith applies the HasEdge predicate on the "host_blacklist" edge with a given conditions (other predicates).
func HasHostBlacklistWith(preds ...predicate.HostBlacklist) predicate.Host {
	return predicate.Host(func(s *sql.Selector) {
		step := newHostBlacklistStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Host) predicate.Host {
	return predicate.Host(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Host) predicate.Host {
	return predicate.Host(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Host) predicate.Host {
	return predicate.Host(sql.NotPredicates(p))
}
