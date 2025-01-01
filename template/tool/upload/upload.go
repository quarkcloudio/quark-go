package upload

import (
	"bytes"
	"encoding/base64"
	"io"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
)

// 文件上传
type Template struct {
	quark.Template
	LimitSize        int64              // 限制文件大小
	LimitType        interface{}        // 限制文件类型
	LimitImageWidth  int64              // 限制图片宽度
	LimitImageHeight int64              // 限制图片高度
	Driver           string             // 存储驱动
	SavePath         string             // 保存路径
	OSSConfig        *quark.OSSConfig   // OSS配置
	MinioConfig      *quark.MinioConfig // Minio配置
}

// 初始化
func (p *Template) Init(ctx *quark.Context) interface{} {
	return p
}

// 初始化模板
func (p *Template) TemplateInit(ctx *quark.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 默认本地上传
	p.Driver = quark.LocalStorage

	return p
}

// 初始化路由映射
func (p *Template) RouteInit() interface{} {
	p.POST("/api/tool/upload/:resource/handle", p.Handle)
	p.POST("/api/tool/upload/:resource/base64Handle", p.HandleFromBase64)
	p.POST("/api/upload/:resource/handle", p.Handle)
	p.POST("/api/upload/:resource/base64Handle", p.HandleFromBase64)

	return p
}

// 执行上传
func (p *Template) Handle(ctx *quark.Context) error {
	var (
		result *quark.FileInfo
	)

	limitW := ctx.Query("limitW", "")
	limitH := ctx.Query("limitH", "")

	contentTypes := strings.Split(ctx.Header("Content-Type"), "; ")
	if len(contentTypes) != 2 {
		return ctx.JSONError("Content-Type error")
	}
	if contentTypes[0] != "multipart/form-data" {
		return ctx.JSONError("Content-Type must use multipart/form-data")
	}

	limitSize := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("LimitSize").Int()

	limitType := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("LimitType").Interface()

	limitImageWidth := int(reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("LimitImageWidth").Int())

	if limitW.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitW.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	limitImageHeight := int(reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("LimitImageHeight").Int())

	if limitH.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitH.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	driver := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Driver").String()

	ossConfig := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("OSSConfig").Interface()

	minioConfig := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("MinioConfig").Interface()

	savePath := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("SavePath").String()

	byteReader := bytes.NewReader(ctx.Body())
	multipartReader := multipart.NewReader(byteReader, strings.TrimLeft(contentTypes[1], "boundary="))
	for p, err := multipartReader.NextPart(); err != io.EOF; p, err = multipartReader.NextPart() {
		if p.FormName() == "file" {
			fileData, _ := io.ReadAll(p)
			fileSystem := quark.
				NewStorage(&quark.StorageConfig{
					LimitSize:        limitSize,
					LimitType:        limitType.([]string),
					LimitImageWidth:  limitImageWidth,
					LimitImageHeight: limitImageHeight,
					Driver:           driver,
					CheckFileExist:   true,
					OSSConfig:        ossConfig.(*quark.OSSConfig),
					MinioConfig:      minioConfig.(*quark.MinioConfig),
				}).
				Reader(&quark.File{
					Header:  p.Header,
					Name:    p.FileName(),
					Content: fileData,
				})

			// 上传前回调
			getFileSystem, fileInfo, err := ctx.Template.(interface {
				BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error)
			}).BeforeHandle(ctx, fileSystem)
			if err != nil {
				return ctx.JSONError(err.Error())
			}
			if fileInfo != nil {
				return ctx.Template.(interface {
					AfterHandle(ctx *quark.Context, result *quark.FileInfo) error
				}).AfterHandle(ctx, fileInfo)
			}

			result, err = getFileSystem.
				WithImageExtra().
				RandName().
				Path(savePath).
				Save()
			if err != nil {
				return ctx.JSONError(err.Error())
			}
		}
	}

	return ctx.Template.(interface {
		AfterHandle(ctx *quark.Context, result *quark.FileInfo) error
	}).AfterHandle(ctx, result)
}

// 通过Base64执行上传
func (p *Template) HandleFromBase64(ctx *quark.Context) error {
	var (
		result *quark.FileInfo
		err    error
	)

	limitW := ctx.Query("limitW", "")
	limitH := ctx.Query("limitH", "")

	data := map[string]interface{}{}
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.JSONError(err.Error())
	}
	if data["file"] == nil {
		return ctx.JSONError("参数错误")
	}

	files := strings.Split(data["file"].(string), ",")
	if len(files) != 2 {
		return ctx.JSONError("格式错误")
	}

	fileData, err := base64.StdEncoding.DecodeString(files[1]) //成图片文件并把文件写入到buffer
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	limitSize := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("LimitSize").Int()

	limitType := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("LimitType").Interface()

	limitImageWidth := int(reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("LimitImageWidth").Int())

	if limitW.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitW.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	limitImageHeight := int(reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("LimitImageHeight").Int())

	if limitH.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitH.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	savePath := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("SavePath").String()

	driver := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Driver").String()

	ossConfig := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("OSSConfig").Interface()

	minioConfig := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("MinioConfig").Interface()

	fileSystem := quark.
		NewStorage(&quark.StorageConfig{
			LimitSize:        limitSize,
			LimitType:        limitType.([]string),
			LimitImageWidth:  limitImageWidth,
			LimitImageHeight: limitImageHeight,
			Driver:           driver,
			CheckFileExist:   true,
			OSSConfig:        ossConfig.(*quark.OSSConfig),
			MinioConfig:      minioConfig.(*quark.MinioConfig),
		}).
		Reader(&quark.File{
			Content: fileData,
		})

	// 上传前回调
	getFileSystem, fileInfo, err := ctx.Template.(interface {
		BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error)
	}).BeforeHandle(ctx, fileSystem)
	if err != nil {
		return ctx.JSONError(err.Error())
	}
	if fileInfo != nil {
		return ctx.Template.(interface {
			AfterHandle(ctx *quark.Context, result *quark.FileInfo) error
		}).AfterHandle(ctx, fileInfo)
	}

	result, err = getFileSystem.
		WithImageExtra().
		RandName().
		Path(savePath).
		Save()

	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.Template.(interface {
		AfterHandle(ctx *quark.Context, result *quark.FileInfo) error
	}).AfterHandle(ctx, result)
}

// 上传前回调
func (p *Template) BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error) {

	return fileSystem, nil, nil
}

// 执行上传
func (p *Template) AfterHandle(ctx *quark.Context, result *quark.FileInfo) error {

	return ctx.JSONOk("上传成功", result)
}
