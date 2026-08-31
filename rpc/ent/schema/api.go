package schema

import "entgo.io/ent"

// Api holds the schema definition for the Api entity.
type Api struct {
	ent.Schema
}

// Fields of the Api.
func (Api) Fields() []ent.Field {
	return nil
}

// Edges of the Api.
func (Api) Edges() []ent.Edge {
	return nil
}
