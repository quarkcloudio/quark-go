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
	I18nKey    string `json:"i18nKey,omitempty"`
	Icon       string `json:"icon,omitempty"`
	Order      int    `json:"order"`
	KeepAlive  bool   `json:"keepAlive,omitempty"`
	HideInMenu bool   `json:"hideInMenu,omitempty"`
	ActiveMenu string `json:"activeMenu,omitempty"`
}

// 用户路由
type UserRoute struct {
	Id        int       `json:"id"`
	Pid       int       `json:"pid"`
	Name      string    `json:"name"`
	Type      int       `json:"type"`
	IsEngine  int       `json:"is_engine"`
	IsLink    int       `json:"is_link"`
	IsFrame   int       `json:"is_frame"`
	Path      string    `json:"path"`
	Component string    `json:"component"`
	Meta      RouteMeta `json:"meta"`
	Query     string    `json:"query,omitempty"`
}

// 用户路由列表
type UserRoutesResp struct {
	Routes interface{} `json:"routes"`
	Home   string      `json:"home"`
}
