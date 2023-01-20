package uploads

import (
	"reflect"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminupload"
	"github.com/quarkcms/quark-go/pkg/msg"
	"github.com/quarkcms/quark-go/pkg/storage"
)

type File struct {
	adminupload.Template
}

// 初始化
func (p *File) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 限制文件大小
	p.LimitSize = 1024 * 1024 * 1024 * 2

	// 限制文件类型
	p.LimitType = []string{
		"image/jpg",
		"image/jpeg",
		"image/png",
		"image/gif",
	}

	p.SavePath = "./website/storage/files/" + time.Now().Format("20060102") + "/"

	return p
}

// 上传完成后回调
func (p *File) AfterHandle(request *builder.Request, templateInstance interface{}, result *storage.FileInfo) interface{} {
	driver := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Driver").String()

	// 重写url
	if driver == storage.LocalDriver {
		result.Url = strings.ReplaceAll(result.Url, "./website/", "/")
	}

	return msg.Success("上传成功", "", result)
}
