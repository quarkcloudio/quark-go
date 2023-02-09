package builder

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/julienschmidt/httprouter"
	"github.com/quarkcms/quark-go/pkg/dal"
	"github.com/quarkcms/quark-go/pkg/github"
	"gorm.io/gorm"
)

// Version of current package
const Version = "1.1.1"

// 静态文件URL
const RespositoryURL = "https://github.com/quarkcms/quark-go/tree/main/website/"

type Engine struct {
	Router      *httprouter.Router         // 路由
	UseHandlers []func(ctx *Context) error // 中间件方法
	config      *Config                    // 配置
	providers   []interface{}              // 服务列表
	routerPaths []string                   // 路由路径列表
}

type RouteMapping struct {
	Method      string
	Path        string
	HandlerName string
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

// Handle is a function that can be registered to a route to handle HTTP
// requests. Like http.HandlerFunc, but has a third parameter for the values of
// wildcards (path variables).
type Handle func(ctx *Context)

// 初始化对象
func New(config *Config) *Engine {

	// 初始化路由对象
	router := httprouter.New()
	router.RedirectTrailingSlash = true

	// 初始化数据库
	if config.DBConfig != nil {
		dal.InitDB(config.DBConfig.Dialector, config.DBConfig.Opts)
	}

	// 初始化后台布局
	config.AdminLayout = initAdminLayout(config.AdminLayout)

	// 定义结构体
	engine := &Engine{
		Router:    router,
		providers: config.Providers,
		config:    config,
	}

	// 下载静态文件
	github.Download(RespositoryURL, config.StaticPath)

	// 初始化路由列表
	engine.initRouterPaths()

	// 处理模版上的路由映射关系
	engine.routeMappingParser()

	// 调用初始化方法
	return engine
}

// 初始化后台布局
func initAdminLayout(adminLayout *AdminLayout) *AdminLayout {
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

// 获取当前配置
func (p *Engine) GetConfig() *Config {
	return p.config
}

// 获取当前AdminLayout配置
func (p *Engine) GetAdminLayout() *AdminLayout {
	return p.config.AdminLayout
}

// 创建上下文
func (p *Engine) NewContext(Writer http.ResponseWriter, Request *http.Request) *Context {
	return &Context{
		Engine:  p,
		Request: Request,
		Writer:  Writer,
	}
}

// 转换Request、Response对象
func (p *Engine) TransformContext(fullPath string, method string, url string, body io.Reader, writer io.Writer) *Context {
	// 转换为http.ResponseWriter
	w := NewResponse(writer)

	// 转换为http.Request
	r := NewRequest(method, url, body)

	// 创建上下文
	ctx := p.NewContext(w, r)

	// 设置当前路由
	ctx.SetFullPath(fullPath)

	// 返回对象
	return ctx
}

// 初始化路由路径列表
func (p *Engine) initRouterPaths() {
	if p.routerPaths != nil {
		return
	}

	var routerPaths []string
	for _, provider := range p.providers {

		// 初始化
		getTemplateInstance := provider.(interface {
			Init() interface{}
		}).Init()

		// 获取模板定义的路由
		templateInstanceRoutes := getTemplateInstance.(interface {
			GetRouteMapping() []*RouteMapping
		}).GetRouteMapping()

		for _, v := range templateInstanceRoutes {
			providerName := reflect.TypeOf(provider).String()
			getNames := strings.Split(providerName, ".")
			structName := getNames[len(getNames)-1]

			if strings.Contains(v.Path, ":resource") {
				path := strings.Replace(v.Path, ":resource", strings.ToLower(structName), -1)
				//处理行为
				if strings.Contains(path, ":uriKey") {
					actions := getTemplateInstance.(interface {
						Actions(ctx *Context) []interface{}
					}).Actions(&Context{})

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
				routerPaths = append(routerPaths, path)
			}
		}
	}

	p.routerPaths = routerPaths
}

// 获取路由路径列表
func (p *Engine) GetRouterPaths() []string {
	return p.routerPaths
}

// 通用调用方法
func (p *Engine) Use(args interface{}) {
	argsName := reflect.TypeOf(args).String()

	switch argsName {
	case "*builder.AdminLayout":
		p.config.AdminLayout = initAdminLayout(args.(*AdminLayout))
	case "func(*builder.Context) error":
		p.UseHandlers = append(p.UseHandlers, args.(func(ctx *Context) error))
	default:
		panic(argsName + " arguments was not found")
	}
}

// 解析模版方法
func (p *Engine) templateHandleParser(ctx *Context) (interface{}, error) {
	var (
		result           interface{}
		err              error
		templateInstance interface{}
	)

	// 获取模板实例
	templateInstance = ctx.Template

	// 执行挂载的方法
	templateInstanceRoutes := templateInstance.(interface {
		GetRouteMapping() []*RouteMapping
	}).GetRouteMapping()

	for _, v := range templateInstanceRoutes {
		handle, params, result := p.Router.Lookup(v.Method, ctx.FullPath())

		fmt.Println(result)
		if result {
			handle(ctx.Writer, ctx.Request, params)
		}
	}

	return result, err
}

// 渲染
func (p *Engine) Render(ctx *Context) (interface{}, error) {
	// 初始化模板
	err := ctx.InitTemplateInstance()
	if err != nil {
		return nil, err
	}

	// 解析UseHandler方法
	err = ctx.useHandlerParser()
	if err != nil {
		return nil, err
	}

	// 解析模版方法
	return p.templateHandleParser(ctx)
}

// 处理模版上的路由映射关系
func (p *Engine) routeMappingParser() {
	for _, provider := range p.providers {

		// 初始化
		getTemplateInstance := provider.(interface {
			Init() interface{}
		}).Init()

		// 获取模板定义的路由
		templateInstanceRoutes := getTemplateInstance.(interface {
			GetRouteMapping() []*RouteMapping
		}).GetRouteMapping()

		for _, v := range templateInstanceRoutes {

			// 匿名函数
			handle := func(ctx *Context) {

				handlerResult := reflect.
					ValueOf(provider).
					MethodByName(v.HandlerName).
					Call([]reflect.Value{
						reflect.ValueOf(ctx),
					})

				if len(handlerResult) == 1 {
					handlerResult[0].Interface()
				}
			}

			switch v.Method {
			case "GET":
				p.GET(v.Path, handle)
			case "HEAD":
				p.HEAD(v.Path, handle)
			case "OPTIONS":
				p.OPTIONS(v.Path, handle)
			case "POST":
				p.POST(v.Path, handle)
			case "PUT":
				p.PUT(v.Path, handle)
			case "PATCH":
				p.PATCH(v.Path, handle)
			case "DELETE":
				p.DELETE(v.Path, handle)
			}
		}
	}
}

// GET is a shortcut for router.Handle(http.MethodGet, path, handle)
func (p *Engine) GET(path string, handle Handle) {
	p.Router.Handle(http.MethodGet, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := p.NewContext(w, r)

		// 设置路由路径
		ctx.SetFullPath(path)

		// 设置Params
		ctx.SetParams(ps)

		// 初始化模板
		ctx.InitTemplateInstance()

		// 解析UseHandler方法
		err := ctx.useHandlerParser()
		if err != nil {
			panic(err)
		}

		// 执行方法
		handle(ctx)
	})
}

// HEAD is a shortcut for router.Handle(http.MethodHead, path, handle)
func (p *Engine) HEAD(path string, handle Handle) {
	p.Router.Handle(http.MethodHead, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := p.NewContext(w, r)

		// 设置路由路径
		ctx.SetFullPath(path)

		// 设置Params
		ctx.SetParams(ps)

		// 初始化模板
		ctx.InitTemplateInstance()

		// 解析UseHandler方法
		err := ctx.useHandlerParser()
		if err != nil {
			panic(err)
		}

		// 执行方法
		handle(ctx)
	})
}

// OPTIONS is a shortcut for router.Handle(http.MethodOptions, path, handle)
func (p *Engine) OPTIONS(path string, handle Handle) {
	p.Router.Handle(http.MethodOptions, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := p.NewContext(w, r)

		// 设置路由路径
		ctx.SetFullPath(path)

		// 设置Params
		ctx.SetParams(ps)

		// 初始化模板
		ctx.InitTemplateInstance()

		// 解析UseHandler方法
		err := ctx.useHandlerParser()
		if err != nil {
			panic(err)
		}

		// 执行方法
		handle(ctx)
	})
}

// POST is a shortcut for router.Handle(http.MethodPost, path, handle)
func (p *Engine) POST(path string, handle Handle) {
	p.Router.Handle(http.MethodPost, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := p.NewContext(w, r)

		// 设置路由路径
		ctx.SetFullPath(path)

		// 设置Params
		ctx.SetParams(ps)

		// 初始化模板
		ctx.InitTemplateInstance()

		// 解析UseHandler方法
		err := ctx.useHandlerParser()
		if err != nil {
			panic(err)
		}

		// 执行方法
		handle(ctx)
	})
}

// PUT is a shortcut for router.Handle(http.MethodPut, path, handle)
func (p *Engine) PUT(path string, handle Handle) {
	p.Router.Handle(http.MethodPut, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := p.NewContext(w, r)

		// 设置路由路径
		ctx.SetFullPath(path)

		// 设置Params
		ctx.SetParams(ps)

		// 初始化模板
		ctx.InitTemplateInstance()

		// 解析UseHandler方法
		err := ctx.useHandlerParser()
		if err != nil {
			panic(err)
		}

		// 执行方法
		handle(ctx)
	})
}

// PATCH is a shortcut for router.Handle(http.MethodPatch, path, handle)
func (p *Engine) PATCH(path string, handle Handle) {
	p.Router.Handle(http.MethodPatch, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := p.NewContext(w, r)

		// 设置路由路径
		ctx.SetFullPath(path)

		// 设置Params
		ctx.SetParams(ps)

		// 初始化模板
		ctx.InitTemplateInstance()

		// 解析UseHandler方法
		err := ctx.useHandlerParser()
		if err != nil {
			panic(err)
		}

		// 执行方法
		handle(ctx)
	})
}

// DELETE is a shortcut for router.Handle(http.MethodDelete, path, handle)
func (p *Engine) DELETE(path string, handle Handle) {
	p.Router.Handle(http.MethodDelete, path, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := p.NewContext(w, r)

		// 设置路由路径
		ctx.SetFullPath(path)

		// 设置Params
		ctx.SetParams(ps)

		// 初始化模板
		ctx.InitTemplateInstance()

		// 解析UseHandler方法
		err := ctx.useHandlerParser()
		if err != nil {
			panic(err)
		}

		// 执行方法
		handle(ctx)
	})
}

// ANY is a shortcut for router.Handle(http.MethodGet, path, handle)
func (p *Engine) ANY(path string, handle Handle) {
	p.GET(path, handle)
	p.HEAD(path, handle)
	p.OPTIONS(path, handle)
	p.POST(path, handle)
	p.PUT(path, handle)
	p.PATCH(path, handle)
	p.DELETE(path, handle)
}

// Run Server
func (p *Engine) Run(addr string) {
	http.ListenAndServe(addr, p.Router)
}
