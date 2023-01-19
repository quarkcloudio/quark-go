package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/quarkcms/quark-go/pkg/adapter/ginadapter"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin"
	"github.com/quarkcms/quark-go/pkg/app/install"
	"github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// 静态文件
	r.Use(static.Serve("/", static.LocalFile("./website", false)))

	// 数据库配置信息
	dsn := "root:Bc5HQFJc4bLjZCcC@tcp(127.0.0.1:3306)/quarkgo?charset=utf8&parseTime=True&loc=Local"

	// 配置资源
	config := &builder.Config{
		AppKey:    "123456",
		Providers: admin.Providers,
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
	}

	// 创建对象
	b := builder.New(config)

	// 初始化安装
	b.Use(install.Handle)

	// 中间件
	b.Use(middleware.Handle)

	// 适配gin
	ginadapter.Adapter(b, r)

	r.Run(":3000")
}
