package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/ns9ryan/common/orm/ent/mixins"
)

// Position 定义职位信息表结构
type Position struct {
	ent.Schema
}

// Fields 定义职位信息表字段
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("职位名称"),

		field.String("code").
			Comment("职位编码"),

		field.String("remark").
			Optional().
			Comment("备注"),
	}
}

// Edges 定义职位信息表关联关系
func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		// 职位与用户关系
		edge.From("users", User.Type).
			Ref("positions"),
	}
}

// Indexes 定义职位信息表索引
func (Position) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").
			Unique(),
	}
}

// Mixin 定义职位信息表公共字段
func (Position) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

// Annotations 定义职位信息表数据库注解
func (Position) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),                 // 启用数据库字段注释
		schema.Comment("职位信息表"),                   // 设置数据库表注释
		entsql.Annotation{Table: "sys_positions"}, // 设置数据库表名
	}
}
