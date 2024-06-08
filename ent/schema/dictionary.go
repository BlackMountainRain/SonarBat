package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Dictionary holds the schema definition for the Dictionary entity.
type Dictionary struct {
	ent.Schema
}

// Fields of the Dictionary.
func (Dictionary) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("category").MaxLen(20),
		field.String("key").MaxLen(20),
		field.String("value").MaxLen(100),
	}
}

// Edges of the Dictionary.
func (Dictionary) Edges() []ent.Edge {
	return nil
}

func (Dictionary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
