package roleservicelogic

import (
	"context"

	"github.com/ns9ryan/p9_core/rpc/internal/svc"
	"github.com/ns9ryan/p9_core/rpc/pb/core/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get 获取角色
func (l *GetLogic) Get(in *role.GetRoleRequest) (*role.GetRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &role.GetRoleResponse{}, nil
}
