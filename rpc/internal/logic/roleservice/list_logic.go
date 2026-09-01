package roleservicelogic

import (
	"context"

	"github.com/ns9ryan/p9_core/rpc/internal/svc"
	"github.com/ns9ryan/p9_core/rpc/pb/core/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// List 获取角色列表
func (l *ListLogic) List(in *role.ListRolesRequest) (*role.ListRolesResponse, error) {
	// todo: add your logic here and delete this line

	return &role.ListRolesResponse{}, nil
}
