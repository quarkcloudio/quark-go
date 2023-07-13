package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/quarkcms/quark-go/v2/pkg/adapter/ginadapter"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/install"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/middleware"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
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
	config := &builder.Config{
		AppKey:    "123456",
		Providers: service.Providers,
		DBConfig: &builder.DBConfig{
			Dialector: mysql.Open(dsn),
			Opts:      &gorm.Config{},
		},
	}

	// 创建对象
	b := builder.New(config)

	// 初始化安装
	install.Handle()

	// 中间件
	b.Use(middleware.Handle)

	// 适配gin
	ginadapter.Adapter(b, r)

	r.Run(":3000")
}
