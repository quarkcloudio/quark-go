package main

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin"
	"github.com/quarkcms/quark-go/pkg/app/handler/miniapp"
	"github.com/quarkcms/quark-go/pkg/app/handler/mix"
	"github.com/quarkcms/quark-go/pkg/app/handler/tool"
	"github.com/quarkcms/quark-go/pkg/app/install"
	"github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 定义服务
	var providers []interface{}

	// 数据库配置信息
	dsn := "root:fK7xPGJi1gJfIief@tcp(127.0.0.1:3306)/quarkgo?charset=utf8&parseTime=True&loc=Local"

	// 加载后台服务
	providers = append(providers, admin.Providers...)

	// 加载Mix服务
	providers = append(providers, mix.Providers...)

	// 加载MiniApp服务
	providers = append(providers, miniapp.Providers...)

	// 加载工具服务
	providers = append(providers, tool.Providers...)

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
			Password: "",
			Port:     "6379",
			Database: 0,
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
