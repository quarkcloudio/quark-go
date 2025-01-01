package adminmodule

import (
	"os"
	"strconv"
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/logins"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/dto"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/file"
	"gorm.io/gorm"
)

// 执行安装操作
func Install() {

	// 如果锁定文件存在则不执行安装步骤
	if file.IsExist("install.lock") {
		return
	}

	// 迁移数据
	db.Client.AutoMigrate(
		&model.ActionLog{},
		&model.User{},
		&model.Config{},
		&model.Menu{},
		&model.Attachment{},
		&model.AttachmentCategory{},
		&model.Permission{},
		&model.Role{},
		&model.Department{},
		&model.Position{},
		&model.CasbinRule{},
	)

	// 如果超级管理员不存在，初始化数据库数据
	adminInfo, err := service.NewUserService().GetInfoById(1)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}
	if adminInfo.Id == 0 {
		// 数据填充
		(&model.User{}).Seeder()
		(&model.Config{}).Seeder()
		(&model.Menu{}).Seeder()
		(&model.Role{}).Seeder()
		(&model.Department{}).Seeder()
		(&model.Position{}).Seeder()
	}

	// 创建锁定文件
	file, _ := os.Create("install.lock")
	file.Close()
}

// 中间件
func Middleware(ctx *quark.Context) error {

	// 获取登录实例
	loginInstance := &logins.Index{}

	// 初始化路由
	loginInstance.RouteInit()

	// 加载自定义路由
	loginInstance.Route()

	// 获取登录模板定义的路由
	loginIndexRoutes := loginInstance.GetRouteMapping()

	inLoginRoute := false
	for _, v := range loginIndexRoutes {
		if v.Path == ctx.FullPath() {
			inLoginRoute = true
		}
	}

	// 排除登录路由
	if inLoginRoute {
		return ctx.Next()
	}

	// 排除非后台路由
	if !strings.Contains(ctx.Path(), "api/admin") {
		return ctx.Next()
	}

	// 定义管理员结构体
	adminInfo := &dto.UserClaims{}

	// 获取登录管理员信息
	err := ctx.JwtAuthUser(adminInfo)
	if err != nil {
		return ctx.JSON(401, quark.Error(err.Error()))
	}

	guardName := adminInfo.GuardName
	if guardName != "admin" {
		return ctx.JSON(401, quark.Error("401 Unauthozied"))
	}

	casbinService := service.NewCasbinService()

	// 管理员id
	if adminInfo.Id != 1 {
		result1, err := casbinService.Enforce("admin|"+strconv.Itoa(adminInfo.Id), ctx.FullPath(), "Any")
		if err != nil {
			return ctx.JSON(500, quark.Error(err.Error()))
		}

		result2, err := casbinService.Enforce("admin|"+strconv.Itoa(adminInfo.Id), ctx.FullPath(), ctx.Method())
		if err != nil {
			return ctx.JSON(500, quark.Error(err.Error()))
		}

		result3, err := casbinService.Enforce("admin|"+strconv.Itoa(adminInfo.Id), ctx.Path(), "Any")
		if err != nil {
			return ctx.JSON(500, quark.Error(err.Error()))
		}

		result4, err := casbinService.Enforce("admin|"+strconv.Itoa(adminInfo.Id), ctx.Path(), ctx.Method())
		if err != nil {
			return ctx.JSON(500, quark.Error(err.Error()))
		}

		if !(result1 || result2 || result3 || result4) {
			return ctx.JSON(403, quark.Error("403 Forbidden"))
		}
	}

	// 记录操作日志
	service.NewActionLogService().InsertGetId(model.ActionLog{
		Uid:      adminInfo.Id,
		Username: adminInfo.Username,
		Url:      ctx.Path(),
		Ip:       ctx.ClientIP(),
		Type:     "ADMIN",
	})

	return ctx.Next()
}
