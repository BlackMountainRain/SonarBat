package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HostBlacklist holds the schema definition for the HostBlacklist entity.
type HostBlacklist struct {
	ent.Schema
}

// Fields of the HostBlacklist.
func (HostBlacklist) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.UUID("host_id", uuid.UUID{}),
		field.String("reason").MaxLen(50).NotEmpty(),
	}
}

// Edges of the HostBlacklist.
func (HostBlacklist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("host", Host.Type).
			Ref("host_blacklist").
			Field("host_id").
			Unique().Required(),
	}
}

func (HostBlacklist) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}
