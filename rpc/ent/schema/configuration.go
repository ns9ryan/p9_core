package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/ns9ryan/common/orm/ent/mixins"
)

// Configuration 定义动态配置表结构
type Configuration struct {
	ent.Schema
}

// Fields 定义动态配置表字段
func (Configuration) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("配置名称"),

		field.String("key").
			Comment("配置键名"),

		field.String("value").
			Comment("配置值"),

		field.String("category").
			Comment("配置分类"),

		field.String("remark").
			Optional().
			Comment("备注"),
	}
}

// Indexes 定义动态配置表索引
func (Configuration) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key"),
	}
}

// Mixin 定义动态配置表公共字段
func (Configuration) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.SortMixin{},
		mixins.StateMixin{},
	}
}

// Annotations 定义动态配置表数据库注解
func (Configuration) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),                     // 启用数据库字段注释
		schema.Comment("动态配置表"),                       // 设置数据库表注释
		entsql.Annotation{Table: "sys_configuration"}, // 设置数据库表名
	}
}
