package response

// 上传文件返回
type UploadFileResp struct {
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

// 上传图片返回
type UploadImageResp struct {
	Id          int         `json:"id"`
	ContentType string      `json:"contentType"`
	Ext         string      `json:"ext"`
	Hash        string      `json:"hash"`
	Height      int         `json:"height"`
	Width       int         `json:"width"`
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Size        int64       `json:"size"`
	Url         string      `json:"url"`
	Extra       interface{} `json:"extra"`
}
