// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"github.com/ns9ryan/p9_core/api/internal/config"
	"github.com/ns9ryan/p9_core/rpc/client/roleservice"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	RoleRpc roleservice.RoleService
	Trans   *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		RoleRpc: roleservice.NewRoleService(zrpc.MustNewClient(c.CoreRpc)),
	}
}
