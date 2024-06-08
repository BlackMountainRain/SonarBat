package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.Bool("status").Default(true),
		field.String("name").NotEmpty().MaxLen(100),
		field.String("comment").Default("").MaxLen(255),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subtasks", Subtask.Type),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		OperatorMixin{},
	}
}

func (Task) Index() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}
