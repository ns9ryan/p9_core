// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package ping

import (
	"context"

	"github.com/ns9ryan/p9_core/api/internal/svc"
	"github.com/ns9ryan/p9_core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.PingResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
