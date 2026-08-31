package baseservicelogic

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/bsm/redislock"
	"github.com/ns9ryan/common/rpcerror"
	"github.com/ns9ryan/p9_core/rpc/internal/svc"
	"github.com/ns9ryan/p9_core/rpc/pb/core/base"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"
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
	// 如果你的 MySQL 速度很快，注释下面的代码。
	// 因为如果数据库太慢，上下文的截止时间就会到
	l.ctx = context.Background()

	// 加锁以避免重复初始化
	locker := redislock.New(l.svcCtx.Redis)
	lock, err := locker.Obtain(l.ctx, "INIT:DATABASE:LOCK", 10*time.Minute, nil)
	if errors.Is(err, redislock.ErrNotObtained) {
		logx.Error("数据库初始化任务正在执行")
		return nil, rpcerror.NewInternal("数据库初始化任务正在执行")
	}

	if err != nil {
		logx.Errorw("获取 Redis 锁失败", logx.Field("detail", err.Error()))
		return nil, rpcerror.NewInternal("获取 Redis 锁失败")
	}

	defer lock.Release(l.ctx)

	// 初始化表结构
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false), schema.WithDropColumn(true), schema.WithDropIndex(true)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		_ = l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:ERROR", err.Error(), 300*time.Second).Err()
		return nil, rpcerror.NewInternal(err.Error())
	}

	// 强制更新 Casbin 策略，以避免在策略更新失败时超级管理员无法登录
	//err = l.insertCasbinPoliciesData()
	if err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		_ = l.svcCtx.Redis.Set(l.ctx, "INIT:DATABASE:ERROR", err.Error(), 300*time.Second).Err()
		return nil, rpcerror.NewInternal(err.Error())
	}

	return &base.InitDatabaseResponse{}, nil
}

// 插入初始 Casbin 策略
//func (l *InitDatabaseLogic) insertCasbinPoliciesData() error {
//	apis, err := l.svcCtx.DB.API.Query().All(l.ctx)
//}
