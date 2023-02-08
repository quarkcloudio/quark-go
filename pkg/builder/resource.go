package builder

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/quarkcms/quark-go/pkg/dal"
	"github.com/quarkcms/quark-go/pkg/github"
	"gorm.io/gorm"
)

type Route struct {
	Path        string
	HandlerName string
}

type Resource struct {
	Providers        []interface{}                  // 服务列表
	Request          *Request                       // 请求数据
	TemplateInstance interface{}                    // 资源模板实例
	UseHandlers      []func(request *Request) error // 中间件方法
}

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

type AdminLayout struct {
	Title        string                   // layout 的左上角 的 title
	Logo         interface{}              // layout 的左上角 的 logo
	Actions      interface{}              // layout 的头部行为
	Layout       string                   // layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
	SplitMenus   bool                     // layout 的菜单模式为mix时，是否自动分割菜单
	ContentWidth string                   // layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
	PrimaryColor string                   // 主题色,"#1890ff"
	FixedHeader  bool                     // 是否固定 header 到顶部
	FixSiderbar  bool                     // 是否固定导航
	IconfontUrl  string                   // 使用 IconFont 的图标配置
	Locale       string                   // 当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
	SiderWidth   int                      // 侧边菜单宽度
	Copyright    string                   // 网站版权 time.Now().Format("2006") + " QuarkGo"
	Links        []map[string]interface{} // 友情链接
}

// 静态文件URL
const RespositoryURL = "https://github.com/quarkcms/quark-go/tree/main/website/"

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

// 获取后台默认布局
func getDefaultAdminLayout(adminLayout *AdminLayout) *AdminLayout {
	defalutAdminLayout := &AdminLayout{
		"QuarkGo",
		false,
		nil,
		"mix",
		false,
		"Fluid",
		"#1890ff",
		true,
		true,
		"//at.alicdn.com/t/font_1615691_3pgkh5uyob.js",
		"zh-CN",
		208,
		time.Now().Format("2006") + " QuarkGo",
		[]map[string]interface{}{
			{
				"key":   "1",
				"title": "Quark",
				"href":  "http://www.quarkcms.com/",
			},
			{
				"key":   "2",
				"title": "爱小圈",
				"href":  "http://www.ixiaoquan.com",
			},
			{
				"key":   "3",
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

// 解析UseHandler方法
func (p *Resource) UseHandlerParser() error {
	var err error

	// 执行本资源的方法
	for _, Handler := range p.UseHandlers {
		err = Handler(p.Request)
		if err != nil {
			return err
		}
	}

	return err
}

// 解析路由方法
func (p *Resource) RouteParser() (interface{}, error) {
	var (
		result           interface{}
		err              error
		templateInstance interface{}
	)

	// 获取模板实例
	templateInstance, err = p.GetTemplateInstance()
	if err != nil {
		return nil, err
	}

	// 设置模板实例
	p.SetTemplateInstance(templateInstance)

	// 执行挂载的方法
	templateInstanceRoutes := templateInstance.(interface {
		GetRoutes() []*Route
	}).GetRoutes()

	for _, v := range templateInstanceRoutes {
		if v.Path == p.Request.FullPath() {
			handlerResult := reflect.
				ValueOf(templateInstance).
				MethodByName(v.HandlerName).
				Call([]reflect.Value{
					reflect.ValueOf(p.Request),
					reflect.ValueOf(p),
					reflect.ValueOf(templateInstance),
				})

			if len(handlerResult) == 1 {
				result = handlerResult[0].Interface()
			}
		}
	}

	return result, err
}

// 替换路由中的资源参数
//
//	url := p.RouteToResourceUrl("/api/admin/login/:resource/captchaId") // url = "/api/admin/login/index/captchaId"
func (p *Resource) RouteToResourceUrl(route string) string {
	resourceName := p.Request.ResourceName()

	return strings.ReplaceAll(route, ":resource", resourceName)
}

// 根据路由判断是否为当前加载实例
func (p *Resource) IsCurrentTemplateInstance(provider interface{}) bool {
	providerName := reflect.TypeOf(provider).String()
	getNames := strings.Split(providerName, ".")
	structName := getNames[len(getNames)-1]
	resourceName := p.Request.ResourceName()

	// fmt.Println(providerName)
	// fmt.Println(resourceName)

	return strings.EqualFold(strings.ToLower(structName), strings.ToLower(resourceName))
}

// 获取当前模板实例
func (p *Resource) GetTemplateInstance() (interface{}, error) {
	var templateInstance interface{}

	for _, provider := range p.Providers {

		// 初始化
		getTemplateInstance := provider.(interface {
			Init() interface{}
		}).Init()

		// 获取模板定义的路由
		templateInstanceRoutes := getTemplateInstance.(interface {
			GetRoutes() []*Route
		}).GetRoutes()

		for _, v := range templateInstanceRoutes {
			if v.Path == p.Request.FullPath() {
				if p.IsCurrentTemplateInstance(provider) {
					// 设置实例
					templateInstance = getTemplateInstance
				}
			}
		}
	}

	if templateInstance == nil {
		return nil, errors.New("未获取到实例")
	}

	return templateInstance, nil
}

// 设置当前模板实例
func (p *Resource) SetTemplateInstance(templateInstance interface{}) {
	// 设置实例
	p.TemplateInstance = templateInstance
}

// 处理执行
func (p *Resource) Run() (interface{}, error) {

	// 解析UseHandler方法
	err := p.UseHandlerParser()
	if err != nil {
		return nil, err
	}

	// 解析路由
	return p.RouteParser()
}
