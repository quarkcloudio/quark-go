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
	LimitSize        int      // 限制文件大小
	LimitType        []string // 限制文件类型
	LimitImageWidth  int      // 限制图片宽度
	LimitImageHeight int      // 限制图片高度
	SavePath         string   // 保存路径
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
	contentTypes := strings.Split(request.Header("Content-Type"), "; ")
	if len(contentTypes) != 2 {
		return msg.Error("Content-Type error", "")
	}
	if contentTypes[0] != "multipart/form-data" {
		return msg.Error("Content-Type must use multipart/form-data", "")
	}

	savePath := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("SavePath").String()

	byteReader := bytes.NewReader(request.Body())
	multipartReader := multipart.NewReader(byteReader, strings.TrimLeft(contentTypes[1], "boundary="))
	for p, err := multipartReader.NextPart(); err != io.EOF; p, err = multipartReader.NextPart() {
		if p.FormName() == "file" {
			fileData, _ := ioutil.ReadAll(p)
			storage.
				New(&storage.Config{}).
				Reader(&storage.File{
					Header:  p.Header,
					Name:    p.FileName(),
					Content: fileData,
				}).
				SetSavePath(savePath).
				SetSaveRandName(true).
				Save()
		}
	}

	return msg.Success("上传成功", "", "")
}
