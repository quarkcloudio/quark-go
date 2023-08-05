package builder

import (
	"io"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/quarkcms/quark-go/v2/pkg/dal"
	"github.com/quarkcms/quark-go/v2/pkg/gopkg"
	"github.com/quarkcms/quark-go/v2/pkg/utils/file"
	"gorm.io/gorm"
)

const (

	// 应用名称
	AppName = "QuarkGo"

	// 版本号
	Version = "2.1.3"

	// 包名
	PkgName = "github.com/quarkcms/quark-go/v2"
)

type Engine struct {
	echo        *echo.Echo                 // Echo框架实例
	useHandlers []func(ctx *Context) error // 中间件方法
	config      *Config                    // 配置
	providers   []interface{}              // 服务列表
	urlPaths    []*UrlPath                 // 请求路径列表
	routePaths  []*RouteMapping            // 路由路径列表
}

type RouteMapping struct {
	Method  string
	Path    string
	Handler func(ctx *Context) error
}

type UrlPath struct {
	Method string
	Url    string
}

type DBConfig struct {
	Dialector gorm.Dialector
	Opts      gorm.Option
}

type RedisConfig struct {
	Host     string // 地址
	Password string // 密码
	Port     string // 端口
	Database int    // 数据库
}

type Config struct {
	AppKey      string        // 应用加密Key，用于JWT认证
	DBConfig    *DBConfig     // 数据库配置
	RedisConfig *RedisConfig  // Redis配置
	StaticPath  string        // 静态文件目录
	Providers   []interface{} // 服务列表
}

// 定义路由组
type Group struct {
	engine    *Engine
	echoGroup *echo.Group
}

// 定义路由方法类型
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

	// 初始化Redis
	if config.RedisConfig != nil {
		dal.InitRedis(&redis.Options{
			Addr:     config.RedisConfig.Host + ":" + config.RedisConfig.Port,
			Password: config.RedisConfig.Password,
			DB:       config.RedisConfig.Database,
		})
	}

	// 定义结构体
	engine := &Engine{
		echo:      e,
		providers: config.Providers,
		config:    config,
	}

	// 默认WEB资源目录
	if config.StaticPath == "" {
		config.StaticPath = "./web"
	}

	// 下载静态文件
	if !file.IsExist(config.StaticPath) {
		err := gopkg.New(PkgName, Version).Save("web", config.StaticPath)
		if err != nil {
			panic(err)
		}
	}

	// 初始化请求列表
	engine.initPaths()

	// 调用初始化方法
	return engine
}

// 获取当前配置
func (p *Engine) GetConfig() *Config {
	return p.config
}

// 获取所有服务
func (p *Engine) GetProviders() []interface{} {
	return p.providers
}

// 创建上下文
func (p *Engine) NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	echoContext := p.echo.NewContext(request, writer)

	return &Context{
		Engine:      p,
		EchoContext: echoContext,
		Request:     request,
		Writer:      writer,
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
		urlPaths   []*UrlPath
		routePaths []*RouteMapping
	)
	if p.urlPaths != nil && p.routePaths != nil {
		return
	}
	for _, provider := range p.providers {

		// 初始化路由
		provider.(interface {
			RouteInit() interface{}
		}).RouteInit()

		// 加载自定义路由
		provider.(interface {
			Route() interface{}
		}).Route()

		// 获取模板定义的路由
		templateRoutes := provider.(interface {
			GetRouteMapping() []*RouteMapping
		}).GetRouteMapping()

		for _, v := range templateRoutes {
			providerName := reflect.TypeOf(provider).String()
			getNames := strings.Split(providerName, ".")
			structName := getNames[len(getNames)-1]

			if strings.Contains(v.Path, ":resource") {
				url := strings.Replace(v.Path, ":resource", strings.ToLower(structName), -1)
				//处理行为
				if strings.Contains(url, ":uriKey") {
					actions := provider.(interface {
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
								url = strings.Replace(url, ":uriKey", uriKey, -1)
							}
						} else {
							url = strings.Replace(url, ":uriKey", uriKey, -1)
						}
					}
				}
				urlPaths = append(urlPaths, &UrlPath{
					Method: v.Method,
					Url:    url,
				})
			}

			if !hasRoutePath(routePaths, v.Method, v.Path) {
				routePaths = append(routePaths, &RouteMapping{v.Method, v.Path, v.Handler})
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
func (p *Engine) GetUrlPaths() []*UrlPath {
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
		result           []reflect.Value
		err              error
		templateInstance interface{}
	)

	// 获取模板实例
	templateInstance = ctx.Template
	if templateInstance == nil {
		return ctx.String(200, "Unable to find resource instance")
	}

	// 执行挂载的方法
	for _, v := range p.routePaths {
		if v.Path == ctx.FullPath() {

			// 反射实例值
			value := reflect.ValueOf(templateInstance)
			if !value.IsValid() {
				continue
			}

			// 获取指针
			pc := reflect.ValueOf(v.Handler).Pointer()

			// 获取func全路径
			fn := runtime.FuncForPC(pc)
			fullPaths := strings.Split(fn.Name(), ".")
			lastPathName := fullPaths[len(fullPaths)-1]

			// 获取方法名称
			funcNames := strings.Split(lastPathName, "-")
			if len(funcNames) <= 0 {
				continue
			}
			funcName := funcNames[0]

			// 获取实例上方法
			method := value.MethodByName(funcName)
			if !method.IsValid() {
				continue
			}

			// 反射执行结果
			result = method.Call([]reflect.Value{
				reflect.ValueOf(ctx),
			})
			if len(result) != 1 {
				continue
			}

			// 执行结果
			if v, ok := result[0].Interface().(error); ok {
				err = v
			}
		}
	}

	return err
}

// 渲染
func (p *Engine) Render(ctx *Context) error {
	// 初始化模板
	err := ctx.InitTemplate(ctx)
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
	// 创建上下文
	ctx := p.NewContext(c.Response().Writer, c.Request())

	// 设置路由路径
	ctx.SetFullPath(path)

	// 初始化模板
	ctx.InitTemplate(ctx)

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

// 加载静态文件
func (p *Engine) Static(pathPrefix string, fsRoot string) {
	p.echo.Static(pathPrefix, fsRoot)
}

// GET请求
func (p *Engine) GET(path string, handle Handle) error {
	p.echo.GET(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// HEAD请求
func (p *Engine) HEAD(path string, handle Handle) error {
	p.echo.HEAD(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// OPTIONS请求
func (p *Engine) OPTIONS(path string, handle Handle) error {
	p.echo.OPTIONS(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// POST请求
func (p *Engine) POST(path string, handle Handle) error {
	p.echo.POST(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// PUT请求
func (p *Engine) PUT(path string, handle Handle) error {
	p.echo.PUT(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// PATCH请求
func (p *Engine) PATCH(path string, handle Handle) error {
	p.echo.PATCH(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// DELETE请求
func (p *Engine) DELETE(path string, handle Handle) error {
	p.echo.DELETE(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// Any请求
func (p *Engine) Any(path string, handle Handle) error {
	p.echo.Any(path, func(c echo.Context) error {
		return p.echoHandle(path, handle, c)
	})

	return nil
}

// 路由组
func (p *Engine) Group(path string, handlers ...Handle) *Group {
	echoGroup := p.echo.Group(path, func(next echo.HandlerFunc) echo.HandlerFunc {
		if len(handlers) > 0 {
			for _, handle := range handlers {
				newHandle := func(c echo.Context) error {
					err := p.echoHandle(path, handle, c)
					if err != nil {
						// 执行下一步操作，这里应该放到数组里面执行，暂时用NextUseHandler
						if err.Error() == "NextUseHandler" {
							return next(c)
						}
					}

					return err
				}
				return newHandle
			}
		}
		return next
	})

	return &Group{engine: p, echoGroup: echoGroup}
}

// GET请求
func (p *Group) GET(path string, handle Handle) error {
	p.echoGroup.GET(path, func(c echo.Context) error {
		return p.engine.echoHandle(path, handle, c)
	})

	return nil
}

// HEAD请求
func (p *Group) HEAD(path string, handle Handle) error {
	p.echoGroup.HEAD(path, func(c echo.Context) error {
		return p.engine.echoHandle(path, handle, c)
	})

	return nil
}

// OPTIONS请求
func (p *Group) OPTIONS(path string, handle Handle) error {
	p.echoGroup.OPTIONS(path, func(c echo.Context) error {
		return p.engine.echoHandle(path, handle, c)
	})

	return nil
}

// POST请求
func (p *Group) POST(path string, handle Handle) error {
	p.echoGroup.POST(path, func(c echo.Context) error {
		return p.engine.echoHandle(path, handle, c)
	})

	return nil
}

// PUT请求
func (p *Group) PUT(path string, handle Handle) error {
	p.echoGroup.PUT(path, func(c echo.Context) error {
		return p.engine.echoHandle(path, handle, c)
	})

	return nil
}

// PATCH请求
func (p *Group) PATCH(path string, handle Handle) error {
	p.echoGroup.PATCH(path, func(c echo.Context) error {
		return p.engine.echoHandle(path, handle, c)
	})

	return nil
}

// DELETE请求
func (p *Group) DELETE(path string, handle Handle) error {
	p.echoGroup.DELETE(path, func(c echo.Context) error {
		return p.engine.echoHandle(path, handle, c)
	})

	return nil
}

// Any请求
func (p *Group) Any(path string, handle Handle) error {
	p.echoGroup.Any(path, func(c echo.Context) error {
		return p.engine.echoHandle(path, handle, c)
	})

	return nil
}

// 路由组
func (p *Group) Group(path string, handlers ...Handle) *Group {
	echoGroup := p.engine.echo.Group(path, func(next echo.HandlerFunc) echo.HandlerFunc {
		if len(handlers) > 0 {
			for _, handle := range handlers {
				newHandle := func(c echo.Context) error {
					err := p.engine.echoHandle(path, handle, c)
					if err != nil {
						// 执行下一步操作，这里应该放到数组里面执行，暂时用NextUseHandler
						if err.Error() == "NextUseHandler" {
							return next(c)
						}
					}

					return err
				}
				return newHandle
			}
		}
		return next
	})

	return &Group{engine: p.engine, echoGroup: echoGroup}
}

// Run Server
func (p *Engine) Run(addr string) {

	// 处理模版上的路由映射关系
	p.routeMappingParser()

	// 启动服务
	p.echo.Logger.Fatal(p.echo.Start(addr))
}
