package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.Int("price"),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("bills").Unique(),
		edge.From("src_user", User.Type).Ref("src_bills").Unique().Required(),
		edge.From("dst_user", User.Type).Ref("dst_bills").Unique().Required(),
	}
}
