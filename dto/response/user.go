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
