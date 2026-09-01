package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ns9ryan/common/orm/ent/mixins"
)

// DictionaryDetail 定义字典键值表结构
type DictionaryDetail struct {
	ent.Schema
}

// Fields 定义字典键值表字段
func (DictionaryDetail) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("展示名称"),

		field.String("key").
			Comment("键"),

		field.String("value").
			Comment("值"),

		field.Uint64("dictionary_id").
			Optional().
			Comment("字典 ID"),
	}
}

// Edges 定义字典键值表关联关系
func (DictionaryDetail) Edges() []ent.Edge {
	return []ent.Edge{
		// 字典明细与字典关系
		edge.From("dictionaries", Dictionary.Type).
			Field("dictionary_id").
			Ref("dictionary_details").
			Unique(),
	}
}

// Mixin 定义字典键值表公共字段
func (DictionaryDetail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

// Annotations 定义字典键值表数据库注解
func (DictionaryDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),                          // 启用数据库字段注释
		schema.Comment("字典键值表"),                            // 设置数据库表注释
		entsql.Annotation{Table: "sys_dictionary_details"}, // 设置数据库表名
	}
}
