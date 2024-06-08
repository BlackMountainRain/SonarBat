package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RbacRole holds the schema definition for the RbacRole entity.
type RbacRole struct {
	ent.Schema
}

// Fields of the RbacRole.
func (RbacRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Bool("status").Default(true),
		field.String("name").NotEmpty().MaxLen(20),
		field.String("description").Default("").MaxLen(255),
	}
}

// Edges of the RbacRole.
func (RbacRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}

func (RbacRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}
