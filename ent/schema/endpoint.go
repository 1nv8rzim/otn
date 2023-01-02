package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Endpoint holds the schema definition for the Endpoint entity.
type Endpoint struct {
	ent.Schema
}

// Fields of the Endpoint.
func (Endpoint) Fields() []ent.Field {
	return []ent.Field{
		field.String("uri").
			Unique().
			NotEmpty().
			Comment("The URI of the endpoint"),
		field.String("content").
			NotEmpty().
			Comment("The content of the endpoint"),
		field.Time("created_at").
			Default(time.Now).
			Comment("The time the endpoint was created"),
	}
}

// Edges of the Endpoint.
func (Endpoint) Edges() []ent.Edge {
	return nil
}
