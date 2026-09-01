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

// Role 定义角色信息表结构
type Role struct {
	ent.Schema
}

// Fields 定义角色信息表字段
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("角色名称"),

		field.String("code").
			NotEmpty().
			Comment("角色编码"),

		field.String("remark").
			Default("").
			Comment("备注"),
	}
}

// Edges 定义角色信息表关联关系
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		// 角色与菜单关系
		edge.To("menus", Menu.Type),

		// 角色与用户关系
		edge.From("users", User.Type).
			Ref("roles"),
	}
}

// Indexes 定义角色信息表索引
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").
			Unique(),
	}
}

// Mixin 定义角色信息表公共字段
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

// Annotations 定义角色信息表数据库注解
func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),             // 启用数据库字段注释
		schema.Comment("角色信息表"),               // 设置数据库表注释
		entsql.Annotation{Table: "sys_roles"}, // 设置数据库表名
	}
}
