package main

import (
	"os"
	"path"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/glebarez/sqlite"
	"github.com/quarkcloudio/quark-go/v3/pkg/adapter/gfadapter"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/install"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/middleware"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"gorm.io/gorm"
)

func main() {
	currentDir, _ := os.Getwd()
	s := g.Server()

	// 配置资源
	config := &builder.Config{

		// JWT加密密串
		AppKey: "123456",

		// 加载服务
		Providers: service.Providers,

		// 数据库配置
		DBConfig: &builder.DBConfig{
			Dialector: sqlite.Open("./examples/gfadmin/data.db"),
			Opts:      &gorm.Config{},
		},
	}

	// 创建对象
	b := builder.New(config)

	// 初始化安装
	install.Handle()

	// 中间件
	b.Use(middleware.Handle)

	// WEB根目录
	s.SetServerRoot("./web/app")
	s.AddSearchPath(path.Join(currentDir, "web/app/admin"))

	// 适配goframe
	gfadapter.Adapter(b, s)

	s.SetPort(3000)
	s.Run()
}
