package main

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/install"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/middleware"
	adminservice "github.com/quarkcms/quark-go/v2/pkg/app/admin/service"
	mixservice "github.com/quarkcms/quark-go/v2/pkg/app/mix/service"
	toolservice "github.com/quarkcms/quark-go/v2/pkg/app/tool/service"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// 定义服务
	var providers []interface{}

	// 数据库配置信息
	dsn := "./data.db"

	// 加载后台服务
	providers = append(providers, adminservice.Providers...)

	// 加载Mix服务
	providers = append(providers, mixservice.Providers...)

	// 加载工具服务
	providers = append(providers, toolservice.Providers...)

	// 配置资源
	config := &builder.Config{
		AppKey:    "123456",
		Providers: providers,
		DBConfig: &builder.DBConfig{
			Dialector: sqlite.Open(dsn),
			Opts:      &gorm.Config{},
		},
	}

	// 实例化对象
	b := builder.New(config)

	// WEB根目录
	b.Static("/", "./web/app")

	// 自动构建数据库、拉取静态文件
	install.Handle()

	// 后台中间件
	b.Use(middleware.Handle)

	// 响应Get请求
	b.GET("/", func(ctx *builder.Context) error {
		return ctx.String(200, "Hello World!")
	})

	// 启动服务
	b.Run(":3000")
}
