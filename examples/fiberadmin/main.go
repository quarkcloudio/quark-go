package main

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/adapter/fiberadapter"
	"github.com/quarkcloudio/quark-go/v4/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	f := fiber.New()

	// 将/admin重定向到/admin/
	f.Use("/admin", func(c *fiber.Ctx) error {
		originalUrl := c.OriginalURL()

		if !strings.HasSuffix(originalUrl, "/") && !strings.Contains("originalUrl", ".") {
			return c.Redirect(originalUrl + "/")
		}

		return c.Next()
	})

	// WEB根目录
	f.Static("/", "./web/app", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        false,
		Index:         "index.html",
		CacheDuration: 1 * time.Second,
		MaxAge:        3600,
	})

	// 数据库配置信息
	dsn := "root:fK7xPGJi1gJfIief@tcp(127.0.0.1:3306)/quarkcloud?charset=utf8&parseTime=True&loc=Local"

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
	app.Install()

	// 中间件
	b.Use(app.Middleware)

	// 适配fiber
	fiberadapter.Adapter(b, f)

	f.Listen(":3000")
}
