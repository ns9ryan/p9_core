// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"github.com/ns9ryan/p9_core/api/internal/config"
	"github.com/ns9ryan/p9_core/api/internal/locales"
	"github.com/ns9ryan/p9_core/api/internal/middleware"
	"github.com/ns9ryan/p9_core/api/internal/rpcerror"
	"github.com/ns9ryan/p9_core/pkg/i18n"
	"github.com/ns9ryan/p9_core/rpc/client/roleservice"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	RoleRpc  roleservice.RoleService
	Trans    *i18n.Translator
	Language rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 创建翻译器
	trans, err := i18n.New(c.I18n, locales.FS)
	logx.Must(err)

	// 创建Core RPC客户端
	coreRpc := zrpc.MustNewClient(
		c.CoreRpc,
		zrpc.WithUnaryClientInterceptor(rpcerror.UnaryClientInterceptor),
	)

	return &ServiceContext{
		Config: c,
		//RoleRpc:  roleservice.NewRoleService(zrpc.MustNewClient(c.CoreRpc)),
		RoleRpc:  roleservice.NewRoleService(coreRpc),
		Trans:    trans,
		Language: middleware.NewLanguageMiddleware().Handle,
	}
}
