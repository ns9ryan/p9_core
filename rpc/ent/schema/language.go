package schema

import "entgo.io/ent"

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return nil
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return nil
}
