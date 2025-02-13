package response

// 上传文件返回
type UploadResp struct {
	Id          int         `json:"id"`
	ContentType string      `json:"contentType"`
	Ext         string      `json:"ext"`
	Hash        string      `json:"hash"`
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Size        int64       `json:"size"`
	Url         string      `json:"url"`
	Extra       interface{} `json:"extra"`
}

// 获取图片列表请求
type Pagination struct {
	DefaultCurrent int         `json:"defaultCurrent"`
	Current        interface{} `json:"current"`
	PageSize       int         `json:"pageSize"`
	Total          interface{} `json:"total"`
}

// 获取图片列表请求
type ImageListResp struct {
	Pagination Pagination  `json:"pagination"`
	List       interface{} `json:"list"`
	Categorys  interface{} `json:"categorys"`
}
