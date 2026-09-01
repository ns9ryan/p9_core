package schematype

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	gen "github.com/ns9ryan/p9_core/rpc/ent"
	"github.com/ns9ryan/p9_core/rpc/ent/hook"
	"github.com/ns9ryan/p9_core/rpc/ent/intercept"
)

const softDeleteField = "deleted_at"

// SoftDeleteMixin 通用软删除字段及处理逻辑
type SoftDeleteMixin struct {
	mixin.Schema
}

// softDeleteKey 软删除跳过标识
type softDeleteKey struct{}

// Fields 定义通用软删除字段
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time(softDeleteField).
			Optional().
			Comment("删除时间"),
	}
}

// SkipSoftDelete 创建跳过软删除限制的上下文
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors 自动过滤已软删除的数据
func (m SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, query intercept.Query) error {
			// 跳过软删除过滤时包含已删除数据
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return nil
			}

			m.P(query)

			return nil
		}),
	}
}

// Hooks 将删除操作转换为软删除
func (m SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
					// 跳过软删除时执行物理删除
					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
						return next.Mutate(ctx, mutation)
					}

					// 获取支持软删除的 Mutation
					mx, ok := mutation.(interface {
						SetOp(ent.Op)
						Client() *gen.Client
						SetDeletedAt(time.Time)
						WhereP(...func(*sql.Selector))
					})
					if !ok {
						return nil, fmt.Errorf("不支持的 Ent Mutation 类型：%T", mutation)
					}

					// 只处理未删除的数据
					m.P(mx)

					// 将删除操作转换为更新 deleted_at
					mx.SetOp(ent.OpUpdate)
					mx.SetDeletedAt(time.Now())

					return mx.Client().Mutate(ctx, mutation)
				})
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}

// P 添加未删除数据过滤条件
func (m SoftDeleteMixin) P(target interface {
	WhereP(...func(*sql.Selector))
}) {
	target.WhereP(
		sql.FieldIsNull(softDeleteField),
	)
}
