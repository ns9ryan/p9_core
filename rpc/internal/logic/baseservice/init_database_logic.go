package baseservicelogic

import (
	"context"
	"fmt"

	"github.com/ns9ryan/p9_core/rpc/internal/svc"
	"github.com/ns9ryan/p9_core/rpc/pb/core/base"
	"github.com/zeromicro/go-zero/core/color"
	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *base.InitDatabaseRequest) (*base.InitDatabaseResponse, error) {
	fmt.Println("进来了")
	logx.AddGlobalFields(logx.Field("service", "my-service"))
	text := "Hello, World!"
	coloredText := logx.WithColor(text, color.FgRed)
	fmt.Println(coloredText)

	ctx := context.Background()
	ctx = logx.ContextWithFields(ctx, logx.Field("request_id", "12345"))

	logger := logx.WithContext(ctx)

	logger.Debugf("Debug logs with value: %d", 42)

	logx.Debugf("Debug logs with value: %d", 42)

	fmt.Println("进来了2")
	// If your mysql speed is high, comment the code below.Because the context deadline will reach if the database is too slow
	l.ctx = context.Background()

	// add lock to avoid duplicate initialization
	//locker := redislock.New(l.svcCtx.Redis)
	return &base.InitDatabaseResponse{}, nil
}
