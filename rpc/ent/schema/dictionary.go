package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ns9ryan/common/orm/ent/mixins"
)

// Dictionary 定义字典信息表结构
type Dictionary struct {
	ent.Schema
}

// Fields 定义字典信息表字段
func (Dictionary) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("展示名称"),

		field.String("name").
			Unique().
			Comment("字典名称"),

		field.String("desc").
			Optional().
			Comment("字典描述"),

		field.Bool("is_public").
			Default(false).
			Comment("是否公开，无需登录即可访问"),
	}
}

// Edges 定义字典信息表关联关系
func (Dictionary) Edges() []ent.Edge {
	return []ent.Edge{
		// 字典与字典明细关系
		edge.To("dictionary_details", DictionaryDetail.Type),
	}
}

// Mixin 定义字典信息表公共字段
func (Dictionary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}

// Annotations 定义字典信息表数据库注解
func (Dictionary) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),                    // 启用数据库字段注释
		schema.Comment("字典信息表"),                      // 设置数据库表注释
		entsql.Annotation{Table: "sys_dictionaries"}, // 设置数据库表名
	}
}
