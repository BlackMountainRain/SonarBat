package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RbacObject holds the schema definition for the RbacObject entity.
type RbacObject struct {
	ent.Schema
}

// Fields of the RbacObject.
func (RbacObject) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Bool("status").Default(true),
		field.String("value").NotEmpty().MaxLen(50),
	}
}

// Edges of the RbacObject.
func (RbacObject) Edges() []ent.Edge {
	return nil
}

func (RbacObject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}
