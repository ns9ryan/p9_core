// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package role

import (
	"context"

	"github.com/ns9ryan/p9_core/api/internal/svc"
	"github.com/ns9ryan/p9_core/api/internal/types"
	"github.com/ns9ryan/p9_core/pkg/i18nkey"
	"github.com/ns9ryan/p9_core/rpc/pb/core/role"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRoleRequest) (resp *types.CreateRoleResponse, err error) {
	_, err = l.svcCtx.RoleRpc.Create(l.ctx,
		&role.CreateRoleRequest{
			Status: req.Status, // 状态
			Name:   req.Name,   // 角色名称
			Code:   req.Code,   // 角色编码
			Remark: req.Remark, // 备注
			Sort:   req.Sort,   // 排序
		})
	if err != nil {
		return nil, err
	}
	return &types.CreateRoleResponse{
		BaseResponse: types.BaseResponse{
			Msg: l.svcCtx.Trans.Trans(l.ctx, i18nkey.CreateSuccess),
		},
	}, nil
}
