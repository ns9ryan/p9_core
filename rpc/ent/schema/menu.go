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

// Menu 定义菜单表结构
type Menu struct {
	ent.Schema
}

// Fields 定义菜单表字段
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("parent_id").
			Default(100000).
			Optional().
			Comment("父菜单 ID"),

		field.String("name").
			Comment("菜单名称"),

		field.Uint32("menu_type").
			Comment("菜单类型：0 目录，1 菜单"),

		field.Uint32("menu_level").
			Comment("菜单层级"),

		field.String("path").
			Optional().
			Default("").
			Comment("菜单路由路径"),

		field.String("component").
			Optional().
			Default("").
			Comment("组件路径"),

		field.String("redirect").
			Optional().
			Default("").
			Comment("跳转路径"),

		field.String("service_name").
			Optional().
			Default("Other").
			Comment("服务名称"),

		field.String("permission").
			Optional().
			Comment("权限标识"),

		field.Bool("disabled").
			Optional().
			Default(false).
			Comment("是否停用"),

		// 菜单元信息
		field.String("title").
			Comment("菜单显示标题"),

		field.String("icon").
			Comment("菜单图标"),

		field.Bool("hide_menu").
			Optional().
			Default(false).
			Comment("是否隐藏菜单"),

		field.Bool("hide_breadcrumb").
			Optional().
			Default(false).
			Comment("是否隐藏面包屑"),

		field.Bool("ignore_keep_alive").
			Optional().
			Default(false).
			Comment("是否取消页面缓存"),

		field.Bool("hide_tab").
			Optional().
			Default(false).
			Comment("是否隐藏页签"),

		field.String("frame_src").
			Optional().
			Default("").
			Comment("内嵌页面地址"),

		field.Bool("carry_param").
			Optional().
			Default(false).
			Comment("是否携带路由参数"),

		field.Bool("hide_children_in_menu").
			Optional().
			Default(false).
			Comment("是否隐藏子菜单"),

		field.Bool("affix").
			Optional().
			Default(false).
			Comment("是否固定页签"),

		field.Uint32("dynamic_level").
			Optional().
			Default(20).
			Comment("最大可打开子页签数量"),

		field.String("real_path").
			Optional().
			Default("").
			Comment("不包含动态参数的实际路由路径"),
	}
}

// Edges 定义菜单表关联关系
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		// 菜单与角色关系
		edge.From("roles", Role.Type).
			Ref("menus"),

		// 菜单上下级关系
		edge.To("children", Menu.Type).
			From("parent").
			Unique().
			Field("parent_id"),
	}
}

// Indexes 定义菜单表索引
func (Menu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),

		index.Fields("path").
			Unique(),
	}
}

// Mixin 定义菜单表公共字段
func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.SortMixin{},
	}
}

// Annotations 定义菜单表数据库注解
func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),             // 启用数据库字段注释
		schema.Comment("菜单表"),                 // 设置数据库表注释
		entsql.Annotation{Table: "sys_menus"}, // 设置数据库表名
	}
}
