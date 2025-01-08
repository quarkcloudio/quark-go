package miniappmodule

import (
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/service"
)

// 中间件
func Middleware(ctx *quark.Context) error {

	// 排除非后台路由
	if !strings.Contains(ctx.Path(), "api/miniapp/user") {
		return ctx.Next()
	}

	// 获取用户信息
	_, err := service.NewAuthService(ctx).GetUser()
	if err != nil {
		return ctx.JSON(401, quark.Error(err.Error()))
	}

	return ctx.Next()
}
