## 介绍
QuarkGO 是一个基于golang的低代码工具；它提供的丰富组件，能帮助您使用很少的代码就能搭建出功能完善的应用系统。

## 系统特性

- 用户管理
- 权限系统
- 菜单管理
- 系统配置
- 操作日志
- 附件管理
- 组件丰富

## 快速开始

1. 创建 demo 文件夹，进入该目录中执行如下命令，初始化项目：
``` bash
go mod init demo/hello
```
2. 创建 main.go 文件
3. 在 main.go 文件中添加如下代码：
```go
package main

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin"
	"github.com/quarkcms/quark-go/pkg/app/install"
	"github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// 定义服务
	var providers []interface{}

	// 数据库配置信息
	dsn := "./data.db"

	// 加载后台服务
	providers = append(providers, admin.Providers...)

	// 配置资源
	config := &builder.Config{
		AppKey:    "123456",
		Providers: providers,
		DBConfig: &builder.DBConfig{
			Dialector: sqlite.Open(dsn),
			Opts:      &gorm.Config{},
		},
	}

	// 实例化对象
	b := builder.New(config)

	// 静态文件
	b.Static("/", "./website")

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
```
4. 拉取依赖
``` bash
go mod tidy
```
5. 启动服务
``` bash
go run main.go
```

后台地址： ```http://127.0.0.1:3000/admin/```

账号：```administrator```
密码：```123456```

## 特别注意
1. **因为众所周知的原因，国内用户拉取静态文件可能会失败；建议您手动下载 [website](https://github.com/quarkcms/quark-go/tree/1.2/website) 静态文件后，将其复制到项目根目录，并在website目录内创建install.lock锁定文件。**
2. **后台用户认证使用了AppKey作为JWT的加密密串，生成环境请务必更改**

## 相关项目
- [QuarkSmart](https://github.com/quarkcms/quark-smart) 单体应用
- [QuarkMicro](https://github.com/quarkcms/quark-go) 微服务应用

## 演示站点
香港站点，页面加载可能比较缓慢

- 地址：http://smart.quarkcms.com/admin/#/
- 账号：```administrator```
- 密码：```123456```

## 技术支持
为了避免打扰作者日常工作，你可以在Github上提交 [Issues](https://github.com/quarkcms/quark-go/issues)

相关教程，你可以查看 [在线文档](http://www.quarkcms.com/quark-go/)

## License
QuarkGo is licensed under The MIT License (MIT).