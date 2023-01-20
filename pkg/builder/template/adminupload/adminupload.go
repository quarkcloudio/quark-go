package adminupload

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
	"github.com/quarkcms/quark-go/pkg/storage"
)

// 文件上传
type Template struct {
	template.AdminTemplate
	LimitSize        int64       // 限制文件大小
	LimitType        interface{} // 限制文件类型
	LimitImageWidth  int64       // 限制图片宽度
	LimitImageHeight int64       // 限制图片高度
	SavePath         string      // 保存路径
}

// 初始化
func (p *Template) Init() interface{} {
	p.TemplateInit()

	return p
}

// 初始化模板
func (p *Template) TemplateInit() interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 清空路由
	p.Routes = nil

	// 注册路由
	p.AddRoute("/api/admin/upload/:resource/handle", "Handle")

	return p
}

// 执行上传
func (p *Template) Handle(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {
	var (
		result *storage.FileInfo
		err    error
	)

	contentTypes := strings.Split(request.Header("Content-Type"), "; ")
	if len(contentTypes) != 2 {
		return msg.Error("Content-Type error", "")
	}
	if contentTypes[0] != "multipart/form-data" {
		return msg.Error("Content-Type must use multipart/form-data", "")
	}

	limitSize := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitSize").Int()

	limitType := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitType").Interface()

	limitImageWidth := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitImageWidth").Int()

	limitImageHeight := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitImageHeight").Int()

	savePath := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("SavePath").String()

	byteReader := bytes.NewReader(request.Body())
	multipartReader := multipart.NewReader(byteReader, strings.TrimLeft(contentTypes[1], "boundary="))
	for p, err := multipartReader.NextPart(); err != io.EOF; p, err = multipartReader.NextPart() {
		if p.FormName() == "file" {
			fileData, _ := ioutil.ReadAll(p)
			result, err = storage.
				New(&storage.Config{
					LimitSize:        limitSize,
					LimitType:        limitType.([]string),
					LimitImageWidth:  int(limitImageWidth),
					LimitImageHeight: int(limitImageHeight),
				}).
				Reader(&storage.File{
					Header:  p.Header,
					Name:    p.FileName(),
					Content: fileData,
				}).
				WithImageWH().
				RandName().
				Path(savePath).
				Save()
		}
	}

	if err != nil {
		return msg.Error(err.Error(), "")
	}

	return templateInstance.(interface {
		AfterHandle(request *builder.Request, result *storage.FileInfo) interface{}
	}).AfterHandle(request, result)
}

// 执行上传
func (p *Template) AfterHandle(request *builder.Request, result *storage.FileInfo) interface{} {

	// 重写url
	result.Url = strings.ReplaceAll(result.Url, "./website/", "/")

	return msg.Success("上传成功", "", result)
}
