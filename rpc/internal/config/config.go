package config

import (
	"github.com/ns9ryan/common/config"
	"github.com/ns9ryan/common/plugins/casbin"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	RedisConf    config.RedisConf
	CasbinConf   casbin.CasbinConf
}
