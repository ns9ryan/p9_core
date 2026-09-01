package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/ns9ryan/common/orm/ent/mixins"
)

// OauthProvider 定义第三方登录配置表结构
type OauthProvider struct {
	ent.Schema
}

// Fields 定义第三方登录配置表字段
func (OauthProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Comment("提供商名称"),

		field.String("client_id").
			Comment("客户端 ID"),

		field.String("client_secret").
			Comment("客户端密钥"),

		field.String("redirect_url").
			Comment("回调地址"),

		field.String("scopes").
			Comment("权限范围"),

		field.String("auth_url").
			Comment("认证地址"),

		field.String("token_url").
			Comment("Token 获取地址"),

		field.Uint64("auth_style").
			Comment("鉴权方式：0 自动，1 第三方登录，2 用户名密码登录"),

		field.String("info_url").
			Comment("用户信息请求地址"),
	}
}

// Mixin 定义第三方登录配置表公共字段
func (OauthProvider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Annotations 定义第三方登录配置表数据库注解
func (OauthProvider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),                       // 启用数据库字段注释
		schema.Comment("第三方登录配置表"),                      // 设置数据库表注释
		entsql.Annotation{Table: "sys_oauth_providers"}, // 设置数据库表名
	}
}
