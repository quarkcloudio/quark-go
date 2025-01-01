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
