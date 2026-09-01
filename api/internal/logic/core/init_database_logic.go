// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package core

import (
	"context"

	"github.com/ns9ryan/p9_core/api/internal/svc"
	"github.com/ns9ryan/p9_core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.InitDatabaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
