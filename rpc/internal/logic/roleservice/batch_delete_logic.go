package roleservicelogic

import (
	"context"

	"github.com/ns9ryan/p9_core/rpc/internal/svc"
	"github.com/ns9ryan/p9_core/rpc/pb/core/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteLogic {
	return &BatchDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BatchDelete 批量删除角色
func (l *BatchDeleteLogic) BatchDelete(in *role.BatchDeleteRolesRequest) (*role.BatchDeleteRolesResponse, error) {
	// todo: add your logic here and delete this line

	return &role.BatchDeleteRolesResponse{}, nil
}
