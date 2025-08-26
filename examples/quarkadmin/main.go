package main

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/app"
	"github.com/quarkcloudio/quark-go/v4/template"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 定义服务
	var providers []interface{}

	// 数据库配置信息
	dsn := "root:fK7xPGJi1gJfIief@tcp(localhost:3306)/quarkgo?charset=utf8&parseTime=True&loc=Local"

	// 加载后台服务
	providers = append(providers, app.Providers...)

	// 配置资源
	config := &quark.Config{
		AppKey:    "123456",
		Providers: providers,
		DBConfig: &quark.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
		// RedisConfig: &quark.RedisConfig{
		// 	Host:     "127.0.0.1",
		// 	Port:     "6379",
		// 	Password: "",
		// 	Database: 0,
		// },
	}

	// 实例化对象
	b := quark.New(config)

	// WEB根目录
	b.Static("/", "./web/app")

	// 初始化安装
	template.Install()

	// 中间件
	b.Use(template.Middleware)

	// 响应Get请求
	b.GET("/", func(ctx *quark.Context) error {
		return ctx.String(200, "Hello World!")
	})

	// 启动服务
	b.Run(":3000")
}
