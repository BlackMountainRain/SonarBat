package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int64("user_id"),
		field.Bool("status").Default(true),
		field.String("name").NotEmpty().MaxLen(20),
		field.String("remark").Default("").MaxLen(200),
		field.String("token").NotEmpty().MaxLen(100),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tokens").
			Field("user_id").
			Unique().Required(),
	}
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}
