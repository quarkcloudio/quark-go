package svc

import (
	"github.com/quarkcms/quark-go/v2/examples/zeroadmin/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
