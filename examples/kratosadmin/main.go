package main

import (
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/adapter/kratosadapter"
	"github.com/quarkcloudio/quark-go/v4/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	hs := http.NewServer(
		http.Address(":3000"),
	)

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

	// 适配kratos
	kratosadapter.Adapter(b, hs)

	// WEB根目录，只能放在后面，否则与其他路由有冲突
	hs.HandlePrefix("/", stdhttp.FileServer(stdhttp.Dir("./web/app")))

	// 创建服务
	app := kratos.New(
		kratos.Server(hs),
	)

	// 启动服务
	app.Run()
}
