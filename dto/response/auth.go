package response

// 用户信息
type UserInfoResp struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Avatar   string   `json:"avatar"`
	Roles    []string `json:"roles"`
	Buttons  []string `json:"buttons"`
}

// 路由元信息
type RouteMeta struct {
	Title      string `json:"title"`
	I18nKey    string `json:"i18nKey"`
	Icon       string `json:"icon"`
	Order      int    `json:"order"`
	KeepAlive  bool   `json:"keepAlive"`
	HideInMenu bool   `json:"hideInMenu"`
	ActiveMenu string `json:"activeMenu"`
}

// 用户路由
type UserRoute struct {
	Id           int       `json:"id"`
	Pid          int       `json:"pid"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Component    string    `json:"component"`
	Meta         RouteMeta `json:"meta"`
	Handle       RouteMeta `json:"handle"`
	MatchedFiles []string  `json:"matchedFiles"`
}

// 用户路由列表
type UserRoutesResp struct {
	Routes interface{} `json:"routes"`
	Home   string      `json:"home"`
}
