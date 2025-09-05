package app

import (
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/app/auths"
	"github.com/quarkcloudio/quark-go/v4/app/dashboards"
	"github.com/quarkcloudio/quark-go/v4/app/layouts"
	"github.com/quarkcloudio/quark-go/v4/app/pages"
	"github.com/quarkcloudio/quark-go/v4/app/resources"
	"github.com/quarkcloudio/quark-go/v4/app/uploads"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/quarkcloudio/quark-go/v4/utils/file"
	"gorm.io/gorm"
)

// 注册服务
var Providers = []interface{}{
	&auths.Index{},
	&layouts.Index{},
	&dashboards.Index{},
	&resources.User{},
	&resources.Role{},
	&resources.Permission{},
	&resources.Department{},
	&resources.Position{},
	&resources.Menu{},
	&resources.ActionLog{},
	&resources.Config{},
	&resources.File{},
	&resources.Image{},
	&resources.WebConfig{},
	&resources.Account{},
	&pages.Index{},
	&uploads.File{},
	&uploads.Image{},
}

type DBConfig struct {
	Dialector gorm.Dialector
	Opts      gorm.Option
}

type RedisConfig struct {
	Host     string // 地址
	Password string // 密码
	Port     string // 端口
	Database int    // 数据库
}

type Config struct {
	AppKey      string                // 应用加密Key，用于JWT认证
	DBConfig    *DBConfig             // 数据库配置
	RedisConfig *RedisConfig          // Redis配置
	CookieStore *sessions.CookieStore // Cookie存储，用于保存Session
	StaticPath  string                // 静态文件目录
	Providers   []interface{}         // 服务列表
}

func New(config *Config) *quark.Engine {
	providers := append(config.Providers, Providers...)

	quarkConfig := &quark.Config{
		AppKey:      config.AppKey,
		CookieStore: config.CookieStore,
		StaticPath:  config.StaticPath,
		Providers:   providers,
	}

	if config.DBConfig != nil {
		quarkConfig.DBConfig = &quark.DBConfig{
			Dialector: config.DBConfig.Dialector,
			Opts:      config.DBConfig.Opts,
		}
	}

	if config.RedisConfig != nil {
		quarkConfig.RedisConfig = &quark.RedisConfig{
			Host:     config.RedisConfig.Host,
			Password: config.RedisConfig.Password,
			Port:     config.RedisConfig.Port,
			Database: config.RedisConfig.Database,
		}
	}

	// 实例化对象
	b := quark.New(quarkConfig)

	// WEB根目录
	b.Static("/", "./web/app")

	// 初始化安装
	Install()

	// 中间件
	b.Use(Middleware)

	return b
}

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
	loginInstance := &auths.Index{}

	// 启动模版
	loginInstance.Bootstrap()

	// 初始化路由
	loginInstance.LoadInitRoute()

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

	adminInfo, err := service.NewAuthService(ctx).GetAdmin()
	if err != nil {
		return ctx.JSON(200, quark.ErrorByCode(quark.StatusUnauthorized))
	}

	casbinService := service.NewCasbinService()
	if adminInfo.Id != 1 {
		result1, err := casbinService.Enforce("admin|"+strconv.Itoa(adminInfo.Id), ctx.FullPath(), "Any")
		if err != nil {
			return ctx.JSON(200, quark.ErrorByCode(quark.StatusForbidden))
		}

		result2, err := casbinService.Enforce("admin|"+strconv.Itoa(adminInfo.Id), ctx.FullPath(), ctx.Method())
		if err != nil {
			return ctx.JSON(200, quark.ErrorByCode(quark.StatusForbidden))
		}

		result3, err := casbinService.Enforce("admin|"+strconv.Itoa(adminInfo.Id), ctx.Path(), "Any")
		if err != nil {
			return ctx.JSON(200, quark.ErrorByCode(quark.StatusForbidden))
		}

		result4, err := casbinService.Enforce("admin|"+strconv.Itoa(adminInfo.Id), ctx.Path(), ctx.Method())
		if err != nil {
			return ctx.JSON(200, quark.ErrorByCode(quark.StatusForbidden))
		}

		if !(result1 || result2 || result3 || result4) {
			return ctx.JSON(200, quark.ErrorByCode(quark.StatusForbidden))
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
