package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/ns9ryan/common/orm/ent/mixins"
)

// API 定义 API 接口表结构
type API struct {
	ent.Schema
}

// Fields 定义 API 接口表字段
func (API) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").
			Comment("API 路径"),

		field.String("method").
			Default("POST").
			Comment("HTTP 请求方法"),

		field.String("description").
			Comment("API 描述"),

		field.String("api_group").
			Comment("API 分组"),

		field.String("service_name").
			Default("Other").
			Comment("服务名称"),

		field.Bool("is_required").
			Default(false).
			Comment("是否为必要接口"),
	}
}

// Edges 定义 API 接口表关联关系
func (API) Edges() []ent.Edge {
	return nil
}

// Indexes 定义 API 接口表索引
func (API) Indexes() []ent.Index {
	return []ent.Index{
		// API 路径和请求方法唯一
		index.Fields("path", "method").
			Unique(),
	}
}

// Mixin 定义 API 接口表公共字段
func (API) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Annotations 定义 API 接口表数据库注解
func (API) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),            // 启用数据库字段注释
		schema.Comment("API 接口表"),            // 设置数据库表注释
		entsql.Annotation{Table: "sys_apis"}, // 设置数据库表名
	}
}
