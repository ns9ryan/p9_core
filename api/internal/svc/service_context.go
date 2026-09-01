// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"github.com/ns9ryan/p9_core/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// CoreRpc coreclient.Core
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
