package builder

type RouteMapping struct {
	Method  string
	Path    string
	Handler func(ctx *Context) error
}

var RouteMappings []*RouteMapping

// 获取路由
func GetRouteMapping() []*RouteMapping {
	return RouteMappings
}

// 是否存在路由
func hasRouteMapping(method string, path string, handler func(ctx *Context) error) bool {
	has := false
	for _, v := range RouteMappings {
		if v.Method == method && v.Path == path {
			has = true
		}
	}
	return has
}

// 注册路由
func AddRouteMapping(method string, path string, handler func(ctx *Context) error) {
	if !hasRouteMapping(method, path, handler) {
		getRoute := &RouteMapping{
			Method:  method,
			Path:    path,
			Handler: handler,
		}

		RouteMappings = append(RouteMappings, getRoute)
	}
}
