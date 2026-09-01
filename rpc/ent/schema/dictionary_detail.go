package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/ns9ryan/common/orm/ent/mixins"
)

// DictionaryDetail holds the schema definition for the DictionaryDetail entity.
type DictionaryDetail struct {
	ent.Schema
}

// Fields of the DictionaryDetail.
func (DictionaryDetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("The title shown in the ui | 展示名称 （建议配合i18n）"),
		field.String("key").
			Comment("key | 键"),
		field.String("value").
			Comment("value | 值"),
		field.Uint64("dictionary_id").Optional().
			Comment("Dictionary ID | 字典ID"),
	}
}

// Edges of the DictionaryDetail.
func (DictionaryDetail) Edges() []ent.Edge {
	return nil
}

func (DictionaryDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (DictionaryDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("Dictionary Key/Value Table | 字典键值表"),
		entsql.Annotation{Table: "sys_dictionary_details"},
	}
}
