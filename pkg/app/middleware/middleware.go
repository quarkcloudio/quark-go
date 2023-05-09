package middleware

import (
	"strings"

	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/msg"
)

// 中间件
func Handle(ctx *builder.Context) error {

	// 排除登录路由
	if strings.Contains(ctx.FullPath(), "api/admin/login") {
		return ctx.Next()
	}

	// 排除非后台路由
	if !strings.Contains(ctx.FullPath(), "api/admin") {
		return ctx.Next()
	}

	// 获取登录管理员信息
	adminInfo, err := (&model.Admin{}).GetAuthUser(ctx.Engine.GetConfig().AppKey, ctx.Token())
	if err != nil {
		return ctx.JSON(401, msg.Error(err.Error(), ""))
	}

	guardName := adminInfo.GuardName
	if guardName != "admin" {
		return ctx.JSON(401, msg.Error("401 Unauthozied", ""))
	}

	// 管理员id
	if adminInfo.Id != 1 {
		permissions, err := (&model.Permission{}).GetListByAdminId(adminInfo.Id)
		if err != nil {
			return ctx.JSON(403, msg.Error("403 Forbidden", ""))
		}

		hasPermission := false
		for _, v := range permissions {
			if strings.ToLower(v.Name) == strings.ToLower(ctx.FullPath()) {
				hasPermission = true
			}

			if strings.ToLower(v.Name) == strings.ToLower(ctx.Path()) {
				hasPermission = true
			}
		}

		if !hasPermission {
			return ctx.JSON(403, msg.Error("403 Forbidden", ""))
		}
	}

	// 记录操作日志
	(&model.ActionLog{}).InsertGetId(&model.ActionLog{
		ObjectId: adminInfo.Id,
		Url:      ctx.Path(),
		Ip:       ctx.ClientIP(),
		Type:     "admin",
	})

	return ctx.Next()
}
