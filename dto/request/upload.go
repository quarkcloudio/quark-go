package request

// 获取图片列表请求
type ImageListReq struct {
	Page       int      `query:"page"`
	CategoryId int      `query:"categoryId"`
	Name       string   `query:"name"`
	Createtime []string `query:"createtime"`
}
