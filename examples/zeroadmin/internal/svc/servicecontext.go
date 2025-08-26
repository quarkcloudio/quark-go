package svc

import (
	"github.com/quarkcloudio/quark-go/v4/examples/zeroadmin/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
