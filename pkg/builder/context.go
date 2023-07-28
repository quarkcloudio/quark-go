package builder

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/derekstavis/go-qs"
	"github.com/gobeam/stringy"
	"github.com/labstack/echo/v4"
)

// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
type Context struct {
	Engine      *Engine                // 引擎实例
	EchoContext echo.Context           // Echo框架上下文
	Request     *http.Request          // Request
	Writer      http.ResponseWriter    // ResponseWriter
	Template    interface{}            // 资源模板实例
	fullPath    string                 // 路由
	Params      map[string]string      // URL param
	Querys      map[string]interface{} // URL querys
}

type ParamValue struct {
	Key   int
	Value string
}

// ContentType returns the Content-Type header of the request.
func (p *Context) ContentType() string {

	return p.Request.Header.Get("Content-Type")
}

// IsTLS returns true if HTTP connection is TLS otherwise false.
func (p *Context) IsTLS() bool {
	return p.EchoContext.IsTLS()
}

// IsWebSocket returns true if HTTP connection is WebSocket otherwise false.
func (p *Context) IsWebSocket() bool {
	return p.EchoContext.IsWebSocket()
}

// Scheme returns the HTTP protocol scheme, `http` or `https`.
func (p *Context) Scheme() string {
	return p.EchoContext.Scheme()
}

// RealIP returns the client's network address based on `X-Forwarded-For`
// or `X-Real-IP` request header.
// The behavior can be configured using `Echo#IPExtractor`.
func (p *Context) RealIP() string {
	return p.EchoContext.RealIP()
}

// QueryParam returns the query param for the provided name.
func (p *Context) QueryParam(name string) string {
	return p.EchoContext.QueryParam(name)
}

// QueryParams returns the query parameters as `url.Values`.
func (p *Context) QueryParams() url.Values {
	return p.EchoContext.QueryParams()
}

// QueryString returns the URL query string.
func (p *Context) QueryString() string {
	return p.EchoContext.QueryString()
}

// FormValue returns the form field value for the provided name.
func (p *Context) FormValue(name string) string {
	return p.EchoContext.FormValue(name)
}

// FormParams returns the form parameters as `url.Values`.
func (p *Context) FormParams() (url.Values, error) {
	return p.EchoContext.FormParams()
}

// FormFile returns the multipart form file for the provided name.
func (p *Context) FormFile(name string) (*multipart.FileHeader, error) {
	return p.EchoContext.FormFile(name)
}

// MultipartForm returns the multipart form.
func (p *Context) MultipartForm() (*multipart.Form, error) {
	return p.EchoContext.MultipartForm()
}

// Cookie returns the named cookie provided in the request.
func (p *Context) Cookie(name string) (*http.Cookie, error) {
	return p.EchoContext.Cookie(name)
}

// SetCookie adds a `Set-Cookie` header in HTTP response.
func (p *Context) SetCookie(cookie *http.Cookie) {
	p.EchoContext.SetCookie(cookie)
}

// Cookies returns the HTTP cookies sent with the request.
func (p *Context) Cookies() []*http.Cookie {
	return p.EchoContext.Cookies()
}

// Get retrieves data from the context.
func (p *Context) Get(key string) interface{} {
	return p.EchoContext.Get(key)
}

// Set saves data in the context.
func (p *Context) Set(key string, val interface{}) {
	p.EchoContext.Set(key, val)
}

// Bind binds the request body into provided type `i`. The default binder
// does it based on Content-Type header.
func (p *Context) Bind(i interface{}) error {
	return p.EchoContext.Bind(i)
}

// Validate validates provided `i`. It is usually called after `Context#Bind()`.
// Validator must be registered using `Echo#Validator`.
func (p *Context) Validate(i interface{}) error {
	return p.EchoContext.Validate(i)
}

// Render renders a template with data and sends a text/html response with status
// code. Renderer must be registered using `Echo.Renderer`.
func (p *Context) Render(code int, name string, data interface{}) error {
	return p.EchoContext.Render(code, name, data)
}

// HTML sends an HTTP response with status code.
func (p *Context) HTML(code int, html string) error {
	return p.EchoContext.HTML(code, html)
}

// HTMLBlob sends an HTTP blob response with status code.
func (p *Context) HTMLBlob(code int, b []byte) error {
	return p.EchoContext.HTMLBlob(code, b)
}

// String sends a string response with status code.
func (p *Context) String(code int, s string) error {
	return p.EchoContext.String(code, s)
}

// JSON sends a JSON response with status code.
func (p *Context) JSON(code int, i interface{}) error {
	return p.EchoContext.JSON(code, i)
}

// JSONPretty sends a pretty-print JSON with status code.
func (p *Context) JSONPretty(code int, i interface{}, indent string) error {
	return p.EchoContext.JSONPretty(code, i, indent)
}

// JSONBlob sends a JSON blob response with status code.
func (p *Context) JSONBlob(code int, b []byte) error {
	return p.EchoContext.JSONBlob(code, b)
}

// JSONP sends a JSONP response with status code. It uses `callback` to construct
// the JSONP payload.
func (p *Context) JSONP(code int, callback string, i interface{}) error {
	return p.EchoContext.JSONP(code, callback, i)
}

// JSONPBlob sends a JSONP blob response with status code. It uses `callback`
// to construct the JSONP payload.
func (p *Context) JSONPBlob(code int, callback string, b []byte) error {
	return p.EchoContext.JSONPBlob(code, callback, b)
}

// XML sends an XML response with status code.
func (p *Context) XML(code int, i interface{}) error {
	return p.EchoContext.XML(code, i)
}

// XMLPretty sends a pretty-print XML with status code.
func (p *Context) XMLPretty(code int, i interface{}, indent string) error {
	return p.EchoContext.XMLPretty(code, i, indent)
}

// XMLBlob sends an XML blob response with status code.
func (p *Context) XMLBlob(code int, b []byte) error {
	return p.EchoContext.XMLBlob(code, b)
}

// Blob sends a blob response with status code and content type.
func (p *Context) Blob(code int, contentType string, b []byte) error {
	return p.EchoContext.Blob(code, contentType, b)
}

// Stream sends a streaming response with status code and content type.
func (p *Context) Stream(code int, contentType string, r io.Reader) error {
	return p.EchoContext.Stream(code, contentType, r)
}

// File sends a response with the content of the file.
func (p *Context) File(file string) error {
	return p.EchoContext.File(file)
}

// Attachment sends a response as attachment, prompting client to save the
// file.
func (p *Context) Attachment(file string, name string) error {
	return p.EchoContext.Attachment(file, name)
}

// Inline sends a response as inline, opening the file in the browser.
func (p *Context) Inline(file string, name string) error {
	return p.EchoContext.Inline(file, name)
}

// NoContent sends a response with no body and a status code.
func (p *Context) NoContent(code int) error {
	return p.EchoContext.NoContent(code)
}

// Redirect redirects the request to a provided URL with status code.
func (p *Context) Redirect(code int, url string) error {
	return p.EchoContext.Redirect(code, url)
}

// Error invokes the registered global HTTP error handler. Generally used by middleware.
// A side-effect of calling global error handler is that now Response has been committed (sent to the client) and
// middlewares up in chain can not change Response status code or Response body anymore.
//
// Avoid using this method in handlers as no middleware will be able to effectively handle errors after that.
func (p *Context) Error(err error) {
	p.EchoContext.Error(err)
}

// 设置SetFullPath
func (p *Context) SetFullPath(fullPath string) *Context {
	p.fullPath = fullPath

	return p
}

// FullPath returns a matched route full path. For not found routes
// returns an empty string.
//
//	router.GET("/user/:id", func(c *gin.Context) {
//	    c.FullPath() == "/user/:id" // true
//	})
func (p *Context) FullPath() string {
	return p.fullPath
}

// Method return request method.
//
// Returned value is valid until returning from RequestHandler.
func (p *Context) Method() string {
	return p.Request.Method
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

// OriginalURL returns url query data
func (p *Context) OriginalURL() string {
	return p.Request.URL.RawQuery
}

// IP tries to parse the headers in [X-Real-Ip, X-Forwarded-For]. It calls RemoteIP() under the hood
func (p *Context) ClientIP() string {
	return p.EchoContext.RealIP()
}

// BodyParser binds the request body to a struct.
// It supports decoding the following content types based on the Content-Type header:
// application/json, application/xml, application/x-www-form-urlencoded, multipart/form-data
// If none of the content types above are matched, it will return a ErrUnprocessableEntity error
func (p *Context) BodyParser(i interface{}) error {
	return p.EchoContext.Bind(i)
}

// 获取请求头数据
func (p *Context) Header(key string) string {
	if len(p.Request.Header[key]) > 0 {
		return p.Request.Header[key][0]
	}

	return ""
}

// Method return request method.
//
// Returned value is valid until returning from RequestHandler.
func (p *Context) Write(data []byte) {

	p.Writer.Write(data)
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
func (p *Context) InitTemplate(ctx *Context) error {
	var (
		err              error
		templateInstance interface{}
	)

	// 获取模板实例
	templateInstance, err = p.getTemplate(ctx)
	if err != nil {
		return err
	}

	// 设置模板实例
	p.setTemplate(templateInstance)

	return nil
}

// 获取当前模板实例
func (p *Context) getTemplate(ctx *Context) (interface{}, error) {
	var templateInstance interface{}
	for _, provider := range p.Engine.providers {

		// 模版参数初始化
		provider.(interface {
			TemplateInit(ctx *Context) interface{}
		}).TemplateInit(ctx)

		// 实例初始化
		provider.(interface {
			Init(ctx *Context) interface{}
		}).Init(ctx)

		// 初始化路由
		provider.(interface {
			RouteInit() interface{}
		}).RouteInit()

		// 加载自定义路由
		provider.(interface {
			Route() interface{}
		}).Route()

		// 获取模板定义的路由
		templateInstanceRoutes := provider.(interface {
			GetRouteMapping() []*RouteMapping
		}).GetRouteMapping()

		for _, v := range templateInstanceRoutes {
			if v.Path == p.FullPath() {
				if p.isCurrentTemplate(provider) {
					// 设置实例
					templateInstance = provider
				}
			}
		}
	}

	if templateInstance == nil {
		return nil, errors.New("Unable to find resource instance")
	}

	return templateInstance, nil
}

// 设置当前模板实例
func (p *Context) setTemplate(templateInstance interface{}) {
	// 设置实例
	p.Template = templateInstance
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

	return (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "save")
}

// 判断当前页面是否为详情页面
func (p *Context) IsDetail() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "detail")
}

// 判断当前页面是否为导出页面
func (p *Context) IsExport() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "export")
}

// 判断当前页面是否为导入页面
func (p *Context) IsImport() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "import")
}

// 输出成功状态的JSON数据，JSONOk("成功") | JSONOk("成功", map[string]interface{}{"title":"标题"})
func (p *Context) JSONOk(message ...interface{}) error {
	var (
		content = ""
		data    interface{}
	)

	if len(message) == 1 {
		content = message[0].(string)
	}

	if len(message) == 2 {
		content = message[0].(string)
		data = message[1]
	}

	return p.JSON(200, Success(content, data))
}

// 输出失败状态的JSON数据，JSONError("错误")
func (p *Context) JSONError(message string) error {

	return p.JSON(200, Error(message))
}

// 执行下一个Use方法，TODO
func (p *Context) Next() error {
	return errors.New("NextUseHandler")
}
