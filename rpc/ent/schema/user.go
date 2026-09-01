package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().
			Comment("User's login name | 登录名"),
		field.String("password").
			Comment("Password | 密码"),
		field.String("nickname").Unique().
			Comment("Nickname | 昵称"),
		field.String("description").Optional().
			Comment("The description of user | 用户的描述信息"),
		field.String("home_path").Default("/dashboard").
			Comment("The home page that the user enters after logging in | 用户登陆后进入的首页"),
		field.String("mobile").Optional().
			Comment("Mobile number | 手机号"),
		field.String("email").Optional().
			Comment("Email | 邮箱号"),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("Avatar | 头像路径"),
		field.Uint64("department_id").Optional().Default(1).
			Comment("Department ID | 部门ID"),
		field.Time("expired_at").
			Comment("The expired time | 到期时间").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email").
			Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		schema.Comment("User Table | 用户信息表"),
		entsql.Annotation{Table: "sys_users"},
	}
}
