package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RbacPolicy holds the schema definition for the RbacPolicy entity.
type RbacPolicy struct {
	ent.Schema
}

// Fields of the RbacPolicy.
func (RbacPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("role").MaxLen(20).Optional().Comment("role or user"),
		field.String("obj").MaxLen(20).Optional().Comment("module or api"),
		field.String("act").MaxLen(20).Optional().Comment("action or method"),
		field.String("uri").MaxLen(20).Optional().Comment("uri or path"),
	}
}

// Edges of the RbacPolicy.
func (RbacPolicy) Edges() []ent.Edge {
	return nil
}

func (RbacPolicy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}
