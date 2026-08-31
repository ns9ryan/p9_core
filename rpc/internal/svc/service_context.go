package svc

import (
	"github.com/ns9ryan/p9_core/rpc/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client           // Ent 数据库客户端
	Redis  redis.UniversalClient // Redis 客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 创建 Ent 客户端配置
	entOpts := []ent.Option{
		ent.Log(logx.Info), // 使用 go-zero 日志输出 SQL
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()), // 设置数据库驱动
	}

	// 开启 Ent 调试模式
	if c.DatabaseConf.Debug {
		entOpts = append(entOpts, ent.Debug())
	}

	// 创建 Ent 数据库客户端
	db := ent.NewClient(entOpts...)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  c.RedisConf.MustNewUniversalRedis(),
	}
}
