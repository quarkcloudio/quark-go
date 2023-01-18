package builder

import (
	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/quarkcms/quark-go/pkg/dal"
	"github.com/quarkcms/quark-go/pkg/github"
	"gorm.io/gorm"
)

// 静态文件URL
const RespositoryURL = "https://github.com/quarkcms/quark-go/tree/2.0/website/"

type DBConfig struct {
	Dialector gorm.Dialector
	Opts      gorm.Option
}

type Config struct {
	AppKey      string        // 应用加密Key，用于JWT认证
	AppName     string        // 应用名称
	DBConfig    *DBConfig     // 数据库配置
	StaticPath  string        // 静态文件目录
	Providers   []interface{} // 服务列表
	AdminLayout *AdminLayout  // 后台布局
	Routes      []string      // 路由列表
}

// 全局配置
var AppConfig *Config

// 初始化对象
func New(config *Config) *Resource {

	// 初始化数据库
	if config.DBConfig != nil {
		dal.InitDB(config.DBConfig.Dialector, config.DBConfig.Opts)
	}

	// 初始化配置
	SetConfig(config)

	// 定义结构体
	resource := &Resource{
		Providers: AppConfig.Providers,
	}

	// 下载静态文件
	github.Download(RespositoryURL, config.StaticPath)

	// 调用初始化方法
	return resource
}

// 转换Request对象
func (p *Resource) TransformRequest(request *Request) *Resource {
	requestInstance := request.Init()

	// 定义结构体
	resource := &Resource{
		Providers:   p.Providers,
		UseHandlers: p.UseHandlers,
		Request:     requestInstance,
	}

	// 初始化路由列表
	p.InitRoutes()

	// 调用初始化方法
	return resource
}

// 初始化路由列表
func (p *Resource) InitRoutes() {
	if AppConfig.Routes != nil {
		return
	}

	var routes []string
	for _, provider := range AppConfig.Providers {

		// 初始化
		getTemplateInstance := provider.(interface {
			Init() interface{}
		}).Init()

		// 获取模板定义的路由
		templateInstanceRoutes := getTemplateInstance.(interface {
			GetRoutes() []*Route
		}).GetRoutes()

		for _, v := range templateInstanceRoutes {
			providerName := reflect.TypeOf(provider).String()
			getNames := strings.Split(providerName, ".")
			structName := getNames[len(getNames)-1]

			if strings.Contains(v.Path, ":resource") {
				path := strings.Replace(v.Path, ":resource", strings.ToLower(structName), -1)
				//处理行为
				if strings.Contains(path, ":uriKey") {
					actions := getTemplateInstance.(interface {
						Actions(request *Request) []interface{}
					}).Actions(p.Request)

					for _, av := range actions {

						// uri唯一标识
						uriKey := av.(interface {
							GetUriKey(interface{}) string
						}).GetUriKey(av)

						actionType := av.(interface{ GetActionType() string }).GetActionType()
						if actionType == "dropdown" {
							dropdownActions := av.(interface{ GetActions() []interface{} }).GetActions()
							for _, dropdownAction := range dropdownActions {
								uriKey := dropdownAction.(interface {
									GetUriKey(interface{}) string
								}).GetUriKey(dropdownAction) // uri唯一标识

								path = strings.Replace(path, ":uriKey", uriKey, -1)
							}
						} else {
							path = strings.Replace(path, ":uriKey", uriKey, -1)
						}
					}
				}

				routes = append(routes, path)
			}
		}
	}

	AppConfig.Routes = routes
}

// 获取后台默认布局
func getDefaultAdminLayout(adminLayout *AdminLayout) *AdminLayout {
	defalutAdminLayout := &AdminLayout{
		"QuarkGo",
		false,
		[]map[string]interface{}{
			{
				"component": "icon",
				"icon":      "icon-question-circle",
				"tooltip":   "使用文档",
				"href":      "https://www.quarkcms.com/",
				"target":    "_blank",
				"style": map[string]interface{}{
					"color": "#000",
				},
			},
		},
		"side",
		false,
		"dark",
		"Fluid",
		"dark",
		"#1890ff",
		true,
		true,
		"//at.alicdn.com/t/font_1615691_3pgkh5uyob.js",
		"zh-CN",
		208,
		time.Now().Format("2006") + " QuarkGo",
		[]map[string]interface{}{
			{
				"title": "Quark",
				"href":  "http://www.quarkcms.com/",
			},
			{
				"title": "爱小圈",
				"href":  "http://www.ixiaoquan.com",
			},
			{
				"title": "Github",
				"href":  "https://github.com/quarkcms",
			},
		},
	}

	// 设置布局
	copier.CopyWithOption(defalutAdminLayout, adminLayout, copier.Option{IgnoreEmpty: true})

	return defalutAdminLayout
}

// 设置配置
func SetConfig(config *Config) {

	// 初始化后台布局
	config.AdminLayout = getDefaultAdminLayout(config.AdminLayout)

	// 赋给全局变量
	AppConfig = config
}

// 获取当前配置
func GetConfig() *Config {

	return AppConfig
}

// 获取当前AdminLayout配置
func GetAdminLayoutConfig() *AdminLayout {

	return AppConfig.AdminLayout
}

// 获取路由列表
func GetRoutes() []string {

	return AppConfig.Routes
}

// 通用调用方法
func (p *Resource) Use(args interface{}) *Resource {
	argsName := reflect.TypeOf(args).String()

	switch argsName {
	case "*builder.AdminLayout":
		AppConfig.AdminLayout = getDefaultAdminLayout(args.(*AdminLayout))
	case "func(*builder.Request) error":
		p.UseHandlers = append(p.UseHandlers, args.(func(request *Request) error))
	default:
		panic(argsName + " arguments was not found")
	}

	return p
}
