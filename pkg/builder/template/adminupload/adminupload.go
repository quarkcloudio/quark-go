package adminupload

import (
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"mime/multipart"
	"reflect"
	"strconv"
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
	LimitSize        int64              // 限制文件大小
	LimitType        interface{}        // 限制文件类型
	LimitImageWidth  int64              // 限制图片宽度
	LimitImageHeight int64              // 限制图片高度
	Driver           string             // 存储驱动
	SavePath         string             // 保存路径
	OSSConfig        *storage.OSSConfig // OSS配置
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

	// 默认本地上传
	p.Driver = storage.LocalDriver

	// 注册路由
	p.AddRoute("/api/admin/upload/:resource/handle", "Handle")
	p.AddRoute("/api/admin/upload/:resource/base64Handle", "HandleFromBase64")

	return p
}

// 执行上传
func (p *Template) Handle(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {
	var (
		result *storage.FileInfo
		err    error
	)

	limitW := request.Query("limitW", "")
	limitH := request.Query("limitH", "")

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

	limitImageWidth := int(reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitImageWidth").Int())

	if limitW.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitW.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	limitImageHeight := int(reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitImageHeight").Int())

	if limitH.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitH.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	driver := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Driver").String()

	ossConfig := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("OSSConfig").Interface()

	savePath := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("SavePath").String()

	byteReader := bytes.NewReader(request.Body())
	multipartReader := multipart.NewReader(byteReader, strings.TrimLeft(contentTypes[1], "boundary="))
	for p, err := multipartReader.NextPart(); err != io.EOF; p, err = multipartReader.NextPart() {
		if p.FormName() == "file" {
			fileData, _ := ioutil.ReadAll(p)
			fileSystem := storage.
				New(&storage.Config{
					LimitSize:        limitSize,
					LimitType:        limitType.([]string),
					LimitImageWidth:  limitImageWidth,
					LimitImageHeight: limitImageHeight,
					Driver:           driver,
					OSSConfig:        ossConfig.(*storage.OSSConfig),
				}).
				Reader(&storage.File{
					Header:  p.Header,
					Name:    p.FileName(),
					Content: fileData,
				})

			// 上传前回调
			getFileSystem, fileInfo, err := templateInstance.(interface {
				BeforeHandle(request *builder.Request, templateInstance interface{}, fileSystem *storage.FileSystem) (*storage.FileSystem, *storage.FileInfo, error)
			}).BeforeHandle(request, templateInstance, fileSystem)
			if err != nil {
				return msg.Error(err.Error(), "")
			}
			if fileInfo != nil {
				return templateInstance.(interface {
					AfterHandle(request *builder.Request, templateInstance interface{}, result *storage.FileInfo) interface{}
				}).AfterHandle(request, templateInstance, fileInfo)
			}

			result, err = getFileSystem.
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
		AfterHandle(request *builder.Request, templateInstance interface{}, result *storage.FileInfo) interface{}
	}).AfterHandle(request, templateInstance, result)
}

// 通过Base64执行上传
func (p *Template) HandleFromBase64(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {
	var (
		result *storage.FileInfo
		err    error
	)

	limitW := request.Query("limitW", "")
	limitH := request.Query("limitH", "")

	data := map[string]interface{}{}
	if err := request.BodyParser(&data); err != nil {
		return msg.Error(err.Error(), "")
	}
	if data["file"] == nil {
		return msg.Error("参数错误", "")
	}

	files := strings.Split(data["file"].(string), ",")
	if len(files) != 2 {
		return msg.Error("格式错误", "")
	}

	fileData, err := base64.StdEncoding.DecodeString(files[1]) //成图片文件并把文件写入到buffer
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	limitSize := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitSize").Int()

	limitType := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitType").Interface()

	limitImageWidth := int(reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitImageWidth").Int())

	if limitW.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitW.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	limitImageHeight := int(reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("LimitImageHeight").Int())

	if limitH.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitH.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	savePath := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("SavePath").String()

	driver := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Driver").String()

	ossConfig := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("OSSConfig").Interface()

	fileSystem := storage.
		New(&storage.Config{
			LimitSize:        limitSize,
			LimitType:        limitType.([]string),
			LimitImageWidth:  limitImageWidth,
			LimitImageHeight: limitImageHeight,
			Driver:           driver,
			OSSConfig:        ossConfig.(*storage.OSSConfig),
		}).
		Reader(&storage.File{
			Content: fileData,
		})

	// 上传前回调
	getFileSystem, fileInfo, err := templateInstance.(interface {
		BeforeHandle(request *builder.Request, templateInstance interface{}, fileSystem *storage.FileSystem) (*storage.FileSystem, *storage.FileInfo, error)
	}).BeforeHandle(request, templateInstance, fileSystem)
	if err != nil {
		return msg.Error(err.Error(), "")
	}
	if fileInfo != nil {
		return templateInstance.(interface {
			AfterHandle(request *builder.Request, templateInstance interface{}, result *storage.FileInfo) interface{}
		}).AfterHandle(request, templateInstance, fileInfo)
	}

	result, err = getFileSystem.
		WithImageWH().
		RandName().
		Path(savePath).
		Save()

	if err != nil {
		return msg.Error(err.Error(), "")
	}

	return templateInstance.(interface {
		AfterHandle(request *builder.Request, templateInstance interface{}, result *storage.FileInfo) interface{}
	}).AfterHandle(request, templateInstance, result)
}

// 上传前回调
func (p *Template) BeforeHandle(request *builder.Request, templateInstance interface{}, fileSystem *storage.FileSystem) (*storage.FileSystem, *storage.FileInfo, error) {

	return fileSystem, nil, nil
}

// 执行上传
func (p *Template) AfterHandle(request *builder.Request, templateInstance interface{}, result *storage.FileInfo) interface{} {
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
