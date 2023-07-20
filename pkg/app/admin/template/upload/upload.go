package upload

import (
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"github.com/quarkcms/quark-go/v2/pkg/storage"
)

// 文件上传
type Template struct {
	builder.Template
	LimitSize        int64                // 限制文件大小
	LimitType        []string             // 限制文件类型
	LimitImageWidth  int                  // 限制图片宽度
	LimitImageHeight int                  // 限制图片高度
	Driver           string               // 存储驱动
	SavePath         string               // 保存路径
	OSSConfig        *storage.OSSConfig   // OSS配置
	MinioConfig      *storage.MinioConfig // Minio配置
}

// 初始化
func (p *Template) Init(ctx *builder.Context) interface{} {
	return p
}

// 初始化模板
func (p *Template) TemplateInit(ctx *builder.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 默认本地上传
	p.Driver = storage.LocalDriver

	return p
}

// 初始化路由映射
func (p *Template) RouteInit() interface{} {
	p.POST("/api/admin/upload/:resource/handle", p.Handle)
	p.POST("/api/admin/upload/:resource/base64Handle", p.HandleFromBase64)

	return p
}

// 获取限制文件大小
func (p *Template) GetLimitSize() int64 {
	return p.LimitSize
}

// 获取限制文件类型
func (p *Template) GetLimitType() []string {
	return p.LimitType
}

// 获取限制图片宽度
func (p *Template) GetLimitImageWidth() int {
	return p.LimitImageWidth
}

// 获取限制图片高度
func (p *Template) GetLimitImageHeight() int {
	return p.LimitImageHeight
}

// 获取存储驱动
func (p *Template) GetDriver() string {
	return p.Driver
}

// 获取保存路径
func (p *Template) GetSavePath() string {
	return p.SavePath
}

// 获取OSS配置
func (p *Template) GetOSSConfig() *storage.OSSConfig {
	return p.OSSConfig
}

// 获取Minio配置
func (p *Template) GetMinioConfig() *storage.MinioConfig {
	return p.MinioConfig
}

// 执行上传
func (p *Template) Handle(ctx *builder.Context) error {
	var (
		result *storage.FileInfo
		err    error
	)

	limitW := ctx.Query("limitW", "")
	limitH := ctx.Query("limitH", "")

	contentTypes := strings.Split(ctx.Header("Content-Type"), "; ")
	if len(contentTypes) != 2 {
		return ctx.JSON(200, message.Error("Content-Type error"))

	}
	if contentTypes[0] != "multipart/form-data" {
		return ctx.JSON(200, message.Error("Content-Type must use multipart/form-data"))
	}

	template := ctx.Template.(Uploader)

	limitSize := template.GetLimitSize()
	limitType := template.GetLimitType()

	limitImageWidth := template.GetLimitImageWidth()
	if limitW.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitW.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	limitImageHeight := template.GetLimitImageHeight()
	if limitH.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitH.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	driver := template.GetDriver()
	ossConfig := template.GetOSSConfig()
	minioConfig := template.GetMinioConfig()
	savePath := template.GetSavePath()

	byteReader := bytes.NewReader(ctx.Body())
	multipartReader := multipart.NewReader(byteReader, strings.TrimLeft(contentTypes[1], "boundary="))
	for p, err := multipartReader.NextPart(); err != io.EOF; p, err = multipartReader.NextPart() {
		if p.FormName() == "file" {
			fileData, _ := ioutil.ReadAll(p)
			fileSystem := storage.
				New(&storage.Config{
					LimitSize:        limitSize,
					LimitType:        limitType,
					LimitImageWidth:  limitImageWidth,
					LimitImageHeight: limitImageHeight,
					Driver:           driver,
					CheckFileExist:   true,
					OSSConfig:        ossConfig,
					MinioConfig:      minioConfig,
				}).
				Reader(&storage.File{
					Header:  p.Header,
					Name:    p.FileName(),
					Content: fileData,
				})

			// 上传前回调
			getFileSystem, fileInfo, err := template.BeforeHandle(ctx, fileSystem)
			if err != nil {
				return ctx.JSON(200, message.Error(err.Error()))
			}
			if fileInfo != nil {
				return template.AfterHandle(ctx, fileInfo)
			}

			result, err = getFileSystem.
				WithImageWH().
				RandName().
				Path(savePath).
				Save()
			if err != nil {
				return ctx.JSON(200, message.Error(err.Error()))
			}
		}
	}

	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return template.AfterHandle(ctx, result)
}

// 通过Base64执行上传
func (p *Template) HandleFromBase64(ctx *builder.Context) error {
	var (
		result *storage.FileInfo
		err    error
	)

	limitW := ctx.Query("limitW", "")
	limitH := ctx.Query("limitH", "")

	data := map[string]interface{}{}
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if data["file"] == nil {
		return ctx.JSON(200, message.Error("参数错误"))
	}

	files := strings.Split(data["file"].(string), ",")
	if len(files) != 2 {
		return ctx.JSON(200, message.Error("格式错误"))
	}

	fileData, err := base64.StdEncoding.DecodeString(files[1]) // 把文件写入到buffer
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	template := ctx.Template.(Uploader)

	limitSize := template.GetLimitSize()
	limitType := template.GetLimitType()

	limitImageWidth := template.GetLimitImageWidth()
	if limitW.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitW.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	limitImageHeight := template.GetLimitImageHeight()
	if limitH.(string) != "" {
		getLimitImageWidth, err := strconv.Atoi(limitH.(string))
		if err == nil {
			limitImageWidth = getLimitImageWidth
		}
	}

	savePath := template.GetSavePath()
	driver := template.GetDriver()
	ossConfig := template.GetOSSConfig()
	minioConfig := template.GetMinioConfig()

	fileSystem := storage.
		New(&storage.Config{
			LimitSize:        limitSize,
			LimitType:        limitType,
			LimitImageWidth:  limitImageWidth,
			LimitImageHeight: limitImageHeight,
			Driver:           driver,
			CheckFileExist:   true,
			OSSConfig:        ossConfig,
			MinioConfig:      minioConfig,
		}).
		Reader(&storage.File{
			Content: fileData,
		})

	// 上传前回调
	getFileSystem, fileInfo, err := template.BeforeHandle(ctx, fileSystem)
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if fileInfo != nil {
		return template.AfterHandle(ctx, fileInfo)
	}

	result, err = getFileSystem.
		WithImageWH().
		RandName().
		Path(savePath).
		Save()

	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return template.AfterHandle(ctx, result)
}

// 上传前回调
func (p *Template) BeforeHandle(ctx *builder.Context, fileSystem *storage.FileSystem) (*storage.FileSystem, *storage.FileInfo, error) {
	return fileSystem, nil, nil
}

// 上传后回调
func (p *Template) AfterHandle(ctx *builder.Context, result *storage.FileInfo) error {
	return ctx.JSON(200, message.Success("上传成功", "", result))
}
