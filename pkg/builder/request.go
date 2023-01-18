package builder

import (
	"encoding/json"
	"strings"

	"github.com/derekstavis/go-qs"
	"github.com/gobeam/stringy"
)

type Request struct {
	IPString       string                 `json:"IPString"`       // 请求的ip地址
	HeaderString   string                 `json:"headerString"`   // 请求的Header字符串
	MethodString   string                 `json:"methodString"`   // 请求方法
	FullPathString string                 `json:"fullPathString"` // 路由
	HostString     string                 `json:"hostString"`     // 主机地址
	PathString     string                 `json:"pathString"`     // URL路径
	QueryString    string                 `json:"queryString"`    // 请求参数
	Headers        map[string]string      `json:"headers"`        // 请求的Headers
	Params         map[string]string      `json:"params"`         // URL param
	Querys         map[string]interface{} `json:"querys"`         // URL querys
	BodyBuffer     []byte                 `json:"bodyBuffer"`     // 请求的Body数据
}

type ParamValue struct {
	Key   int
	Value string
}

// 初始化
func (p *Request) Init() *Request {
	p.parseParams()
	p.parseHeaders()
	p.parseQuerys()

	return p
}

// Method return request method.
//
// Returned value is valid until returning from RequestHandler.
func (p *Request) Method() string {
	return p.MethodString
}

// FullPath returns a matched route full path. For not found routes
// returns an empty string.
//
//	router.GET("/user/:id", func(c *hertz.RequestContext) {
//	    c.FullPath() == "/user/:id" // true
//	})
func (p *Request) FullPath() string {
	return p.FullPathString
}

// Host returns requested host.
//
// The host is valid until returning from RequestHandler.
func (p *Request) Host() string {
	return p.HostString
}

// Path returns requested path.
//
// The path is valid until returning from RequestHandler.
func (p *Request) Path() string {
	return p.PathString
}

// IP tries to parse the headers in [X-Real-Ip, X-Forwarded-For]. It calls RemoteIP() under the hood
func (p *Request) IP() string {
	return p.IPString
}

// OriginalURL returns url query data
func (p *Request) OriginalURL() string {
	return p.QueryString
}

// Body returns body data
func (p *Request) Body() []byte {
	return p.BodyBuffer
}

// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
func (p *Request) parseParams() {
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

	p.Params = params
}

// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
//
//	router.GET("/user/:id", func(c *hertz.RequestContext) {
//	    // a GET request to /user/john
//	    id := c.Param("id") // id == "john"
//	})
func (p *Request) Param(key string) string {
	return p.Params[key]
}

// Query returns the value of the URL Querys.
func (p *Request) parseQuerys() {
	querys, err := qs.Unmarshal(p.QueryString)
	if err == nil {
		p.Querys = querys
	}
}

// Query returns the value of the URL query.
func (p *Request) Query(params ...interface{}) interface{} {
	if len(params) == 2 {
		if p.Querys[params[0].(string)] == nil {
			return params[1]
		}
	}

	return p.Querys[params[0].(string)]
}

// AllQuerys returns all query arguments from RequestURI.
func (p *Request) AllQuerys() map[string]interface{} {
	return p.Querys
}

// ResourceName returns the value of the URL Resource param.
// If request path is "/api/admin/login/index" and route is "/api/admin/login/:resource"
//
//	resourceName := p.ResourceName() // resourceName = "index"
func (p *Request) ResourceName() string {
	return p.Param("resource")
}

// BodyParser binds the request body to a struct.
// It supports decoding the following content types based on the Content-Type header:
// application/json, application/xml, application/x-www-form-urlencoded, multipart/form-data
// If none of the content types above are matched, it will return a ErrUnprocessableEntity error
func (p *Request) BodyParser(out interface{}) error {
	return json.Unmarshal(p.BodyBuffer, out)
}

// 解析头部数据
func (p *Request) parseHeaders() {
	params := map[string]string{}
	headers := strings.Split(p.HeaderString, "\r\n")
	for _, v := range headers {
		header := strings.Split(v, ": ")
		if len(header) == 2 {
			params[header[0]] = header[1]
		}
	}

	p.Headers = params
}

// 获取请求头数据
func (p *Request) Header(key string) string {
	return p.Headers[key]
}

// 获取Header中的token
func (p *Request) Token() string {
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
func (p *Request) IsIndex() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "index")
}

//判断当前页面是否为创建页面
func (p *Request) IsCreating() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "create") || (uri[len(uri)-1] == "store")
}

// 判断当前页面是否为编辑页面
func (p *Request) IsEditing() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "update")
}

// 判断当前页面是否为详情页面
func (p *Request) IsDetail() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "detail")
}

// 判断当前页面是否为导出页面
func (p *Request) isExport() bool {
	uri := strings.Split(p.Path(), "/")

	return (uri[len(uri)-1] == "export")
}
