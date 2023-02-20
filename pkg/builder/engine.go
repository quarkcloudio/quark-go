package builder

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/quarkcms/quark-go/pkg/dal"
	"github.com/quarkcms/quark-go/pkg/github"
	"gorm.io/gorm"
)

const (
	// Name of current package
	AppName = "QuarkGo"

	// Version of current package
	Version = "1.1.22"

	// 静态文件URL
	RespositoryURL = "https://github.com/quarkcms/quark-go/tree/main/website/"
)

type Engine struct {
	echo        *echo.Echo                 // Echo框架实例
	useHandlers []func(ctx *Context) error // 中间件方法
	config      *Config                    // 配置
	providers   []interface{}              // 服务列表
	urlPaths    []string                   // 请求路径列表
	routePaths  []*RouteMapping            // 路由路径列表
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
type Handle func(ctx *Context) error

// 初始化对象
func New(config *Config) *Engine {

	// 初始化echo引擎
	e := echo.New()

	// 隐藏banner
	e.HideBanner = true

	// 初始化数据库
	if config.DBConfig != nil {
		dal.InitDB(config.DBConfig.Dialector, config.DBConfig.Opts)
	}

	// 初始化后台布局
	config.AdminLayout = initAdminLayout(config.AdminLayout)

	// 定义结构体
	engine := &Engine{
		echo:      e,
		providers: config.Providers,
		config:    config,
	}

	// 下载静态文件
	github.Download(RespositoryURL, config.StaticPath)

	// 初始化请求列表
	engine.initPaths()

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

// 获取所有服务
func (p *Engine) GetProviders() []interface{} {
	return p.providers
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
func (p *Engine) TransformContext(fullPath string, header map[string][]string, method string, url string, body io.Reader, writer io.Writer) *Context {
	// 转换为http.ResponseWriter
	w := NewResponse(writer)

	// 转换为http.Request
	r := NewRequest(method, url, body)

	// 转换Header
	r.Header = header

	// 创建上下文
	ctx := p.NewContext(w, r)

	// 设置当前路由
	ctx.SetFullPath(fullPath)

	// 返回对象
	return ctx
}

// 初始化请求列表
func (p *Engine) initPaths() {
	var (
		urlPaths   []string
		routePaths []*RouteMapping
	)

	if p.urlPaths != nil && p.routePaths != nil {
		return
	}

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
				urlPaths = append(urlPaths, path)
			}

			if !hasRoutePath(routePaths, v.Method, v.Path) {
				routePaths = append(routePaths, &RouteMapping{v.Method, v.Path, v.HandlerName})
			}
		}
	}

	p.urlPaths = urlPaths
	p.routePaths = routePaths
}

// 判断是否存在RoutePath
func hasRoutePath(routePaths []*RouteMapping, method string, path string) bool {
	var has bool
	for _, v := range routePaths {
		if v.Method == method && v.Path == path {
			has = true
		}
	}

	return has
}

// 获取请求列表
func (p *Engine) GetUrlPaths() []string {
	return p.urlPaths
}

// 获取路由列表
func (p *Engine) GetRoutePaths() []*RouteMapping {
	return p.routePaths
}

// 通用调用方法
func (p *Engine) Use(args interface{}) {
	argsName := reflect.TypeOf(args).String()

	switch argsName {
	case "*builder.AdminLayout":
		p.config.AdminLayout = initAdminLayout(args.(*AdminLayout))
	case "func(*builder.Context) error":
		p.useHandlers = append(p.useHandlers, args.(func(ctx *Context) error))
	default:
		panic(argsName + " arguments was not found")
	}
}

// 获取通用调用方法
func (p *Engine) UseHandlers() []func(ctx *Context) error {
	return p.useHandlers
}

// 解析模版方法
func (p *Engine) handleParser(ctx *Context) error {
	var (
		result           interface{}
		err              error
		templateInstance interface{}
	)

	// 获取模板实例
	templateInstance = ctx.Template
	if templateInstance == nil {
		return ctx.String(200, "未获取到资源实例")
	}

	// 执行挂载的方法
	for _, v := range p.routePaths {
		if v.Path == ctx.FullPath() {
			value := reflect.ValueOf(templateInstance)
			if value.IsValid() {
				method := value.MethodByName(v.HandlerName)
				if method.IsValid() {
					getResult := method.Call([]reflect.Value{
						reflect.ValueOf(ctx),
					})
					if len(getResult) == 1 {
						result = getResult[0].Interface()
						if v, ok := result.(error); ok {
							err = v
						}
					}
				}
			}
		}
	}

	return err
}

// 渲染
func (p *Engine) Render(ctx *Context) error {
	// 初始化模板
	err := ctx.InitTemplate()
	if err != nil {
		return err
	}

	// 解析UseHandler方法
	err = ctx.useHandlerParser()
	if err.Error() != ctx.Next().Error() {
		return err
	}
	if err != nil {
		if err.Error() == ctx.Next().Error() {
			// 解析模版方法
			return p.handleParser(ctx)
		}
	}

	// 解析模版方法
	return err
}

// 处理模版上的路由映射关系
func (p *Engine) routeMappingParser() {
	for _, routePath := range p.routePaths {

		switch routePath.Method {
		case "GET":
			p.GET(routePath.Path, func(ctx *Context) error {
				return p.handleParser(ctx)
			})
		case "HEAD":
			p.HEAD(routePath.Path, func(ctx *Context) error {
				return p.handleParser(ctx)
			})
		case "OPTIONS":
			p.OPTIONS(routePath.Path, func(ctx *Context) error {
				return p.handleParser(ctx)
			})
		case "POST":
			p.POST(routePath.Path, func(ctx *Context) error {
				return p.handleParser(ctx)
			})
		case "PUT":
			p.PUT(routePath.Path, func(ctx *Context) error {
				return p.handleParser(ctx)
			})
		case "PATCH":
			p.PATCH(routePath.Path, func(ctx *Context) error {
				return p.handleParser(ctx)
			})
		case "DELETE":
			p.DELETE(routePath.Path, func(ctx *Context) error {
				return p.handleParser(ctx)
			})
		case "Any":
			p.Any(routePath.Path, func(ctx *Context) error {
				return p.handleParser(ctx)
			})
		}
	}
}

// 获取Echo框架实例
func (p *Engine) Echo() *echo.Echo {
	return p.echo
}

// 适配Echo框架方法
func (p *Engine) echoHandle(path string, handle Handle, c echo.Context) error {
	ctx := p.NewContext(c.Response().Writer, c.Request())

	// 设置路由路径
	ctx.SetFullPath(path)

	// 初始化模板
	ctx.InitTemplate()

	// 解析UseHandler方法
	err := ctx.useHandlerParser()
	if err != nil {
		if err.Error() == ctx.Next().Error() {
			// 执行方法
			return handle(ctx)
		}
	}

	return err
}

// tatic registers a new route with path prefix to serve static files from the provided root directory.
func (p *Engine) Static(pathPrefix string, fsRoot string) {
	p.echo.Static(pathPrefix, fsRoot)
}

// Group creates a new router group with prefix and optional group-level middleware.
func (p *Engine) Group(path string, handlers ...Handle) error {
	p.echo.Group(path, func(next echo.HandlerFunc) echo.HandlerFunc {
		if len(handlers) > 0 {
			for _, handle := range handlers {
				newHandle := func(c echo.Context) error {
					return p.echoHandle(path, handle, c)
				}
				return newHandle
			}
		}
		return next
	})
	return nil
}

// GET is a shortcut for router.Handle(http.MethodGet, path, handle)
func (p *Engine) GET(path string, handle Handle) error {
	p.echo.GET(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// HEAD is a shortcut for router.Handle(http.MethodHead, path, handle)
func (p *Engine) HEAD(path string, handle Handle) error {
	p.echo.HEAD(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// OPTIONS is a shortcut for router.Handle(http.MethodOptions, path, handle)
func (p *Engine) OPTIONS(path string, handle Handle) error {
	p.echo.OPTIONS(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// POST is a shortcut for router.Handle(http.MethodPost, path, handle)
func (p *Engine) POST(path string, handle Handle) error {
	p.echo.POST(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// PUT is a shortcut for router.Handle(http.MethodPut, path, handle)
func (p *Engine) PUT(path string, handle Handle) error {
	p.echo.PUT(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// PATCH is a shortcut for router.Handle(http.MethodPatch, path, handle)
func (p *Engine) PATCH(path string, handle Handle) error {
	p.echo.PATCH(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// DELETE is a shortcut for router.Handle(http.MethodDelete, path, handle)
func (p *Engine) DELETE(path string, handle Handle) error {
	p.echo.DELETE(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// Any is a shortcut for router.Handle(http.MethodGet, path, handle)
func (p *Engine) Any(path string, handle Handle) error {
	p.echo.Any(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// Run Server
func (p *Engine) Run(addr string) {

	// 处理模版上的路由映射关系
	p.routeMappingParser()

	// 启动服务
	p.echo.Logger.Fatal(p.echo.Start(addr))
}
