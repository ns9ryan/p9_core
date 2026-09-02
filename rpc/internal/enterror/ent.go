package enterror

import (
	"github.com/ns9ryan/p9_core/pkg/grpcerror"
	"github.com/ns9ryan/p9_core/rpc/ent"
	"github.com/zeromicro/go-zero/core/logx"
)

// HandleEnt 处理 Ent 错误
func HandleEnt(logger logx.Logger, err error) error {
	switch {
	case ent.IsNotFound(err):
		logger.Errorw("数据不存在", logx.Field("error", err.Error()))
		return grpcerror.NotFound("目标不存在")

	case ent.IsConstraintError(err):
		logger.Errorw("数据约束错误", logx.Field("error", err.Error()))
		return grpcerror.InvalidArgument("数据约束错误")

	case ent.IsValidationError(err):
		logger.Errorw("数据校验错误", logx.Field("error", err.Error()))
		return grpcerror.InvalidArgument("数据校验错误")

	case ent.IsNotSingular(err):
		logger.Errorw("查询结果不唯一", logx.Field("error", err.Error()))
		return grpcerror.Internal("数据异常")

	default:
		logger.Errorw("数据库操作失败", logx.Field("error", err.Error()))
		return grpcerror.Internal("数据库操作失败")
	}
}
