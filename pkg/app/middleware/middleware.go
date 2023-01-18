package middleware

import (
	"errors"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin/login"
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
)

// 中间件
func Handle(request *builder.Request) error {
	loginIndex := (&login.Index{}).Init()

	// 获取登录模板定义的路由
	loginIndexRoutes := loginIndex.(interface {
		GetRoutes() []*builder.Route
	}).GetRoutes()

	inLoginRoute := false
	for _, v := range loginIndexRoutes {
		if v.Path == request.FullPath() {
			inLoginRoute = true
		}
	}

	// 排除登录路由
	if inLoginRoute {
		return nil
	}

	// 获取登录管理员信息
	adminInfo, err := (&model.Admin{}).GetAuthUser(request.Token())
	if err != nil {
		return err
	}

	guardName := adminInfo.GuardName
	if guardName != "admin" {
		return errors.New("401 Unauthozied")
	}

	// 管理员id
	if adminInfo.Id != 1 {
		permissions, err := (&model.Permission{}).GetListByAdminId(adminInfo.Id)
		if err != nil {
			return errors.New("403 Forbidden")
		}

		hasPermission := false
		for _, v := range permissions {
			if "/"+v.Name == request.Path() {
				hasPermission = true
			}
		}

		if !hasPermission {
			return errors.New("403 Forbidden")
		}
	}

	// 记录操作日志
	(&model.ActionLog{}).InsertGetId(&model.ActionLog{
		ObjectId: adminInfo.Id,
		Url:      request.Path(),
		Ip:       request.IP(),
		Type:     "admin",
	})

	return nil
}
