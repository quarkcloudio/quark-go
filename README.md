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

## 示例

```go
package main

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin"
	"github.com/quarkcms/quark-go/pkg/app/install"
	"github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
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

	// 实例化对象
	b := builder.New(config)

	// 自动构建数据库、拉取静态文件
	b.Use(install.Handle)

	// 使用后台中间件
	b.Use(middleware.Handle)

	// 响应Get请求
	b.GET("/", func(ctx *builder.Context) error {
		ctx.Write([]byte("hello world!"))

		return nil
	})

	// 启动服务
	b.Run(":3000")
}
```

后台地址： http://127.0.0.1:3000/admin/
```
账号：administrator
密码：123456
```


## 技术支持
为了避免打扰作者日常工作，你可以在Github上提交 [Issues](https://github.com/quarkcms/quark-go/issues)

相关教程，你可以查看 [在线文档](http://www.quarkcms.com/quark-go/)

## License
QuarkGo is licensed under The MIT License (MIT).