package main

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 实例化对象
	b := app.New(&app.Config{
		AppKey: "123456",
		DBConfig: &app.DBConfig{
			Dialector: mysql.Open("root:fK7xPGJi1gJfIief@tcp(localhost:3306)/quarkcloud?charset=utf8&parseTime=True&loc=Local"),
			Opts:      &gorm.Config{},
		},
	})

	// 响应Get请求
	b.GET("/", func(ctx *quark.Context) error {
		return ctx.String(200, "Hello World!")
	})

	// 启动服务
	b.Run(":4000")
}
