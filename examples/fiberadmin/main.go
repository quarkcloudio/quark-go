package main

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/adapter/fiberadapter"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin"
	"github.com/quarkcms/quark-go/pkg/app/install"
	"github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// 将/admin重定向到/admin/
	app.Use("/admin", func(c *fiber.Ctx) error {
		originalUrl := c.OriginalURL()

		if !strings.HasSuffix(originalUrl, "/") && !strings.Contains("originalUrl", ".") {
			return c.Redirect(originalUrl + "/")
		}

		return c.Next()
	})

	// 静态资源
	app.Static("/", "./website", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        false,
		Index:         "index.html",
		CacheDuration: 1 * time.Second,
		MaxAge:        3600,
	})

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
	install.Handle()

	// 中间件
	b.Use(middleware.Handle)

	// 适配fiber
	fiberadapter.Adapter(b, app)

	app.Listen(":3000")
}
