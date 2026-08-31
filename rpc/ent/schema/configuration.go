package schema

import "entgo.io/ent"

// Configuration holds the schema definition for the Configuration entity.
type Configuration struct {
	ent.Schema
}

// Fields of the Configuration.
func (Configuration) Fields() []ent.Field {
	return nil
}

// Edges of the Configuration.
func (Configuration) Edges() []ent.Edge {
	return nil
}
