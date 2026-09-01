package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	commonmixins "github.com/ns9ryan/common/orm/ent/mixins"
	"github.com/ns9ryan/p9_core/rpc/ent/schema/schematype"
)

// User 定义用户信息表结构
type User struct {
	ent.Schema
}

// Fields 定义用户信息表字段
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			Comment("登录名"),

		field.String("password").
			Comment("密码"),

		field.String("nickname").
			Unique().
			Comment("昵称"),

		field.String("description").
			Optional().
			Comment("用户描述"),

		field.String("home_path").
			Default("/dashboard").
			Comment("登录后首页路径"),

		field.String("mobile").
			Optional().
			Comment("手机号"),

		field.String("email").
			Optional().
			Comment("邮箱"),

		field.String("avatar").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(512)",
			}).
			Optional().
			Comment("头像路径"),

		field.Uint64("department_id").
			Optional().
			Default(1).
			Comment("部门 ID"),

		field.Time("expired_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Optional().
			Comment("到期时间"),
	}
}

// Edges 定义用户信息表关联关系
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// 用户与部门关系
		edge.To("departments", Department.Type).
			Unique().
			Field("department_id"),

		// 用户与职位关系
		edge.To("positions", Position.Type),

		// 用户与角色关系
		edge.To("roles", Role.Type),
	}
}

// Indexes 定义用户信息表索引
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email").
			Unique(),
	}
}

// Mixin 定义用户信息表公共字段及通用行为
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		commonmixins.UUIDMixin{},
		commonmixins.StatusMixin{},
		schematype.SoftDeleteMixin{},
	}
}

// Annotations 定义用户信息表数据库注解
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),             // 启用数据库字段注释
		schema.Comment("用户信息表"),               // 设置数据库表注释
		entsql.Annotation{Table: "sys_users"}, // 设置数据库表名
	}
}
