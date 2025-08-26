package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/adapter/ginadapter"
	"github.com/quarkcloudio/quark-go/v4/app"
	"github.com/quarkcloudio/quark-go/v4/template"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// WEB根目录
	r.Use(static.Serve("/", static.LocalFile("./web/app", false)))

	// 数据库配置信息
	dsn := "root:fK7xPGJi1gJfIief@tcp(127.0.0.1:3306)/quarkgo?charset=utf8&parseTime=True&loc=Local"

	// 配置资源
	config := &quark.Config{
		AppKey:    "123456",
		Providers: app.Providers,
		DBConfig: &quark.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
	}

	// 创建对象
	b := quark.New(config)

	// 初始化安装
	template.Install()

	// 中间件
	b.Use(template.Middleware)

	// 适配gin
	ginadapter.Adapter(b, r)

	r.Run(":3000")
}
