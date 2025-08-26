package uploads

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/dto/response"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/quarkcloudio/quark-go/v4/template/upload"
)

type File struct {
	upload.Template
}

// 初始化
func (p *File) Init(ctx *quark.Context) interface{} {

	// 限制文件大小
	p.LimitSize = 1024 * 1024 * 1024 * 2

	// 限制文件类型
	p.LimitType = []string{
		"image/png",
		"image/gif",
		"image/jpeg",
		"video/mp4",
		"video/mpeg",
		"application/x-xls",
		"application/x-ppt",
		"application/msword",
		"application/zip",
		"application/pdf",
		"application/vnd.ms-excel",
		"application/vnd.ms-powerpoint",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation",
	}

	// 设置文件上传路径
	p.SavePath = "./web/app/storage/files/" + time.Now().Format("20060102") + "/"

	return p
}

// 上传前回调
func (p *File) BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error) {
	fileHash, err := fileSystem.GetFileHash()
	if err != nil {
		return fileSystem, nil, err
	}

	fileInfo, err := service.NewAttachmentService().GetInfoByHash(fileHash)
	if err != nil {
		return fileSystem, nil, err
	}

	if fileInfo.Id != 0 {
		var extra map[string]interface{}
		if fileInfo.Extra != "" {
			_ = json.Unmarshal([]byte(fileInfo.Extra), &extra)
		}
		fileInfo := &quark.FileInfo{
			Name:  fileInfo.Name,
			Size:  fileInfo.Size,
			Ext:   fileInfo.Ext,
			Path:  fileInfo.Path,
			Url:   fileInfo.Url,
			Hash:  fileInfo.Hash,
			Extra: extra,
		}
		return fileSystem, fileInfo, err
	}

	return fileSystem, nil, err
}

// 上传完成后回调
func (p *File) AfterHandle(ctx *quark.Context, result *quark.FileInfo) error {
	driver := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Driver").
		String()

	// 重写url
	if driver == quark.LocalStorage {
		result.Url = service.NewAttachmentService().GetFileUrl(result.Url)
	}

	adminInfo, err := service.NewAuthService(ctx).GetAdmin()
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	extra := ""
	if result.Extra != nil {
		extraData, err := json.Marshal(result.Extra)
		if err == nil {
			extra = string(extraData)
		}
	}

	// 插入数据库
	id, err := service.NewAttachmentService().InsertGetId(model.Attachment{
		Source: "ADMIN",
		Uid:    adminInfo.Id,
		Name:   result.Name,
		Type:   "FILE",
		Size:   result.Size,
		Ext:    result.Ext,
		Path:   result.Path,
		Url:    result.Url,
		Hash:   result.Hash,
		Extra:  extra,
		Status: 1,
	})
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	return ctx.CJSONOk("上传成功", response.UploadResp{
		Id:          id,
		ContentType: result.ContentType,
		Ext:         result.Ext,
		Hash:        result.Hash,
		Name:        result.Name,
		Path:        result.Path,
		Size:        result.Size,
		Url:         result.Url,
		Extra:       result.Extra,
	})
}
