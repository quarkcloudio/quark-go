package request

// 获取图片列表请求
type ImageListReq struct {
	Page       int      `query:"page"`
	CategoryId int      `query:"categoryId"`
	Name       string   `query:"name"`
	Createtime []string `query:"createtime"`
}

// 图片删除请求
type ImageDeleteReq struct {
	Id int `json:"id" form:"id" query:"id"`
}

// 图片裁剪请求
type ImageCropReq struct {
	Id   int    `json:"id" form:"id"`
	File string `json:"file" form:"file"`
}
