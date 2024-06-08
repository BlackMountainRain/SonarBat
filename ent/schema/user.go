package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Bool("status").Default(true),
		field.String("username").NotEmpty().MaxLen(255),
		field.String("email").NotEmpty().MaxLen(255),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", RbacRole.Type).
			Ref("users"),
		edge.To("tokens", Token.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}

//CREATE TABLE IF NOT EXISTS users (
//id bigserial PRIMARY KEY,
//status boolean NOT NULL DEFAULT true,
//
//username varchar(255) NOT NULL,
//email varchar(255) NOT NULL,
//
//roles BIGINT[],
//
//created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
//updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL
//);
//CREATE INDEX idx_users_username ON users USING btree (username);
//CREATE UNIQUE INDEX idx_users_email ON users USING btree (email);
