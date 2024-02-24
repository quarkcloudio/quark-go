package main

import (
	admininstall "github.com/quarkcloudio/quark-go/v2/pkg/app/admin/install"
	adminmiddleware "github.com/quarkcloudio/quark-go/v2/pkg/app/admin/middleware"
	adminservice "github.com/quarkcloudio/quark-go/v2/pkg/app/admin/service"
	miniappinstall "github.com/quarkcloudio/quark-go/v2/pkg/app/miniapp/install"
	miniappmiddleware "github.com/quarkcloudio/quark-go/v2/pkg/app/miniapp/middleware"
	miniappservice "github.com/quarkcloudio/quark-go/v2/pkg/app/miniapp/service"
	toolservice "github.com/quarkcloudio/quark-go/v2/pkg/app/tool/service"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 定义服务
	var providers []interface{}

	// 数据库配置信息
	dsn := "root:fK7xPGJi1gJfIief@tcp(127.0.0.1:3306)/quarkgo?charset=utf8&parseTime=True&loc=Local"

	// 加载后台服务
	providers = append(providers, adminservice.Providers...)

	// 加载MiniApp服务
	providers = append(providers, miniappservice.Providers...)

	// 加载工具服务
	providers = append(providers, toolservice.Providers...)

	// 配置资源
	config := &builder.Config{
		AppKey:    "123456",
		Providers: providers,
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
		RedisConfig: &builder.RedisConfig{
			Host:     "127.0.0.1",
			Port:     "6379",
			Password: "",
			Database: 0,
		},
	}

	// 实例化对象
	b := builder.New(config)

	// WEB根目录
	b.Static("/", "./web/app")

	// 构建管理后台数据库
	admininstall.Handle()

	// 管理后台中间件
	b.Use(adminmiddleware.Handle)

	// 构建MiniApp数据库
	miniappinstall.Handle()

	// MiniApp中间件
	b.Use(miniappmiddleware.Handle)

	// 响应Get请求
	b.GET("/", func(ctx *builder.Context) error {
		return ctx.String(200, "Hello World!")
	})

	// 启动服务
	b.Run(":3000")
}
