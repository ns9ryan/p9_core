package roleservicelogic

import (
	"context"

	"github.com/ns9ryan/p9_core/rpc/internal/svc"
	"github.com/ns9ryan/p9_core/rpc/pb/core/role"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Create 创建角色
func (l *CreateLogic) Create(in *role.CreateRoleRequest) (*role.CreateRoleResponse, error) {
	return &role.CreateRoleResponse{}, nil
	// result, err := l.svcCtx.DB.Role.Create().
	// 	SetNillableStatus(in.Status).
	// 	SetName(in.Name).
	// 	SetCode(in.Code).
	// 	SetNillableRemark(in.Remark).
	// 	SetNillableSort(in.Sort).
	// 	Save(l.ctx)
	// if err != nil {
	// 	return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	// }
	//
	// return &role.CreateRoleResponse{Id: result.ID}, nil
}
