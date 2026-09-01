package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/gofrs/uuid/v5"
	"github.com/ns9ryan/common/orm/ent/mixins"
)

// Token 定义令牌信息表结构
type Token struct {
	ent.Schema
}

// Fields 定义令牌信息表字段
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).
			Comment("用户 UUID"),

		field.String("username").
			Default("unknown").
			Comment("用户名"),

		field.String("token").
			Comment("Token 字符串"),

		field.String("source").
			Comment("Token 来源"),

		field.Time("expired_at").
			Comment("过期时间"),
	}
}

// Indexes 定义令牌信息表索引
func (Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid"),

		index.Fields("expired_at"),
	}
}

// Mixin 定义令牌信息表公共字段
func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

// Annotations 定义令牌信息表数据库注解
func (Token) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),              // 启用数据库字段注释
		schema.Comment("令牌信息表"),                // 设置数据库表注释
		entsql.Annotation{Table: "sys_tokens"}, // 设置数据库表名
	}
}
