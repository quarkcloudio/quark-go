package builder

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/derekstavis/go-qs"
	"github.com/gobeam/stringy"
)

// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
type Context struct {
	Engine   *Engine                // 引擎实例
	Request  *http.Request          // Request
	Writer   http.ResponseWriter    // ResponseWriter
	Template interface{}            // 资源模板实例
	fullPath string                 // 路由
	Params   map[string]string      // URL param
	Querys   map[string]interface{} // URL querys
}

type ParamValue struct {
	Key   int
	Value string
}

// 设置SetFullPath
func (c *Context) SetFullPath(fullPath string) *Context {
	c.fullPath = fullPath

	return c
}

// FullPath returns a matched route full path. For not found routes
// returns an empty string.
//
//	router.GET("/user/:id", func(c *gin.Context) {
//	    c.FullPath() == "/user/:id" // true
//	})
func (c *Context) FullPath() string {
	return c.fullPath
}

// Method return request method.
//
// Returned value is valid until returning from RequestHandler.
func (c *Context) Method() string {
	return c.Request.Method
}

// Host returns requested host.
//
// The host is valid until returning from RequestHandler.
func (p *Context) Host() string {
	return p.Request.Host
}

// Path returns requested path.
//
// The path is valid until returning from RequestHandler.
func (p *Context) Path() string {
	return p.Request.URL.Path
}

// IP tries to parse the headers in [X-Real-Ip, X-Forwarded-For]. It calls RemoteIP() under the hood
func (p *Context) ClientIP() string {
	return p.Request.RemoteAddr
}

// OriginalURL returns url query data
func (p *Context) OriginalURL() string {
	return p.Request.URL.RawQuery
}

// Body returns body data
func (p *Context) Body() []byte {
	body, err := ioutil.ReadAll(p.Request.Body)
	if err != nil {
		return nil
	}

	// 重新赋值
	p.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body
}

// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
func (p *Context) parseParams() {
	params := map[string]string{}
	fullPaths := strings.Split(p.FullPath(), "/")
	paramValues := []*ParamValue{}
	for k, v := range fullPaths {
		if strings.Contains(v, ":") {
			v = stringy.New(v).ReplaceFirst(":", "")
			mapValue := &ParamValue{
				Key:   k,
				Value: v,
			}
			paramValues = append(paramValues, mapValue)
		}
	}

	paths := strings.Split(p.Path(), "/")
	for _, v := range paramValues {
		params[v.Value] = paths[v.Key]
	}

	p.SetParams(params)
}

// 设置Params
func (p *Context) SetParams(params map[string]string) *Context {
	p.Params = params

	return p
}

// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
//
//	router.GET("/user/:id", func(c *hertz.RequestContext) {
//	    // a GET request to /user/john
//	    id := c.Param("id") // id == "john"
//	})
func (p *Context) Param(key string) string {
	if p.Params == nil {
		p.parseParams()
	}

	return p.Params[key]
}

// Query returns the value of the URL Querys.
func (p *Context) parseQuerys() {
	querys, err := qs.Unmarshal(p.Request.URL.RawQuery)
	if err == nil {
		p.Querys = querys
	}
}

// Query returns the value of the URL query.
func (p *Context) Query(params ...interface{}) interface{} {
	if p.Querys == nil {
		p.parseQuerys()
	}

	if len(params) == 2 {
		if p.Querys[params[0].(string)] == nil {
			return params[1]
		}
	}

	return p.Querys[params[0].(string)]
}

// AllQuerys returns all query arguments from RequestURI.
func (p *Context) AllQuerys() map[string]interface{} {
	if p.Querys == nil {
		p.parseQuerys()
	}

	return p.Querys
}

// ResourceName returns the value of the URL Resource param.
// If request path is "/api/admin/login/index" and route is "/api/admin/login/:resource"
//
//	resourceName := p.ResourceName() // resourceName = "index"
func (p *Context) ResourceName() string {
	return p.Param("resource")
}

// 替换路由中的资源参数
//
//	url := p.RouteToEngineUrl("/api/admin/login/:resource/captchaId") // url = "/api/admin/login/index/captchaId"
func (p *Context) RouterPathToUrl(routerPath string) string {
	name := p.ResourceName()

	return strings.ReplaceAll(routerPath, ":resource", name)
}

// BodyParser binds the request body to a struct.
// It supports decoding the following content types based on the Content-Type header:
// application/json, application/xml, application/x-www-form-urlencoded, multipart/form-data
// If none of the content types above are matched, it will return a ErrUnprocessableEntity error
func (p *Context) BodyParser(out interface{}) error {
	return json.Unmarshal(p.Body(), out)
}

// 获取请求头数据
func (p *Context) Header(key string) string {
	if len(p.Request.Header[key]) > 0 {
		return p.Request.Header[key][0]
	}

	return ""
}

// 获取Header中的token
func (p *Context) Token() string {
	authorization := p.Header("Authorization")
	authorizations := strings.Split(authorization, " ")
	if len(authorizations) == 2 {
		return authorizations[1]
	}

	queryToken := p.Query("token", "")
	if queryToken.(string) != "" {
		return queryToken.(string)
	}

	return ""
}

// 判断当前页面是否为列表页面 todo
func (p *Context) IsIndex() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "index")
}

// 判断当前页面是否为创建页面
func (p *Context) IsCreating() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "create") || (uri[len(uri)-1] == "store")
}

// 判断当前页面是否为编辑页面
func (p *Context) IsEditing() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "update")
}

// 判断当前页面是否为详情页面
func (p *Context) IsDetail() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "detail")
}

// 判断当前页面是否为导出页面
func (p *Context) isExport() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "export")
}

// 判断当前页面是否为导入页面
func (p *Context) isImport() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "import")
}

// 根据路由判断是否为当前加载实例
func (p *Context) isCurrentTemplate(provider interface{}) bool {
	providerName := reflect.TypeOf(provider).String()
	getNames := strings.Split(providerName, ".")
	structName := getNames[len(getNames)-1]
	ResourceName := p.ResourceName()

	return strings.EqualFold(strings.ToLower(structName), strings.ToLower(ResourceName))
}

// 解析UseHandler方法
func (p *Context) useHandlerParser() error {
	var err error
	for _, Handler := range p.Engine.UseHandlers() {
		err = Handler(p)
		if err == nil {
			return nil
		}
		if err.Error() != p.Next().Error() {
			return err
		}
	}

	return err
}

// 初始化模板实例
func (p *Context) InitTemplate() error {
	var (
		err              error
		templateInstance interface{}
	)

	// 获取模板实例
	templateInstance, err = p.getTemplate()
	if err != nil {
		return err
	}

	// 设置模板实例
	p.setTemplate(templateInstance)

	return nil
}

// 获取当前模板实例
func (p *Context) getTemplate() (interface{}, error) {
	var templateInstance interface{}

	for _, provider := range p.Engine.providers {

		// 初始化
		template := provider.(interface {
			Init() interface{}
		}).Init()

		// 获取模板定义的路由
		templateInstanceRoutes := template.(interface {
			GetRouteMapping() []*RouteMapping
		}).GetRouteMapping()

		for _, v := range templateInstanceRoutes {
			if v.Path == p.FullPath() {
				if p.isCurrentTemplate(provider) {
					// 设置实例
					templateInstance = template
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
func (p *Context) setTemplate(templateInstance interface{}) {
	// 设置实例
	p.Template = templateInstance
}

// Method return request method.
//
// Returned value is valid until returning from RequestHandler.
func (p *Context) Write(data []byte) {

	p.Writer.Write(data)
}

// 输出Json数据
func (p *Context) JSON(code int, data interface{}) error {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(code)
	err := json.NewEncoder(p.Writer).Encode(data)
	if err != nil {
		return err
	}

	return nil
}

// 输出String数据
func (p *Context) String(code int, data string) error {
	p.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	p.Writer.WriteHeader(code)
	p.Writer.Write([]byte(data))

	return nil
}

// 执行下一个Use方法，TODO
func (p *Context) Next() error {
	return errors.New("NextUseHandler")
}
