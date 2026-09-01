package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ns9ryan/common/orm/ent/mixins"
)

// Department 定义部门表结构
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("部门名称"),

		field.Uint64("parent_id").
			Optional().
			Default(0).
			Comment("父级部门 ID"),

		field.String("ancestors").
			Optional().
			Comment("父级部门列表"),

		field.String("leader").
			Optional().
			Comment("部门负责人"),

		field.String("phone").
			Optional().
			Comment("负责人电话"),

		field.String("email").
			Optional().
			Comment("负责人邮箱"),

		field.String("remark").
			Optional().
			Comment("备注"),
	}
}

// Edges 定义部门表关联关系
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		// 部门上下级关系
		edge.To("children", Department.Type).
			From("parent").
			Unique().
			Field("parent_id"),

		// 部门与用户关系
		edge.From("users", User.Type).
			Ref("departments"),
	}
}

// Mixin 定义部门表公共字段
func (Department) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

// Annotations 定义部门表数据库注解
func (Department) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),                   // 启用数据库字段注释
		schema.Comment("部门表"),                       // 设置数据库表注释
		entsql.Annotation{Table: "sys_departments"}, // 设置数据库表名
	}
}
