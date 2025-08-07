package uploads

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dto/request"
	"github.com/quarkcloudio/quark-go/v3/dto/response"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/upload"
)

type Image struct {
	upload.Template
	GetListPath string // 获取文件列表路由路径
	DeletePath  string // 文件删除路由路径
	CropPath    string // 图片裁剪路由路径
}

// 启动模版
func (p *Image) Bootstrap() interface{} {
	p.GetListPath = "/api/admin/upload/:resource/getList"           // 获取文件列表路由路径
	p.DeletePath = "/api/admin/upload/:resource/delete"             // 文件删除路由路径
	p.CropPath = "/api/admin/upload/:resource/crop"                 // 图片裁剪路由路径
	p.HandlePath = "/api/admin/upload/:resource/handle"             // 文件上传路由路径
	p.Base64HandlePath = "/api/admin/upload/:resource/base64Handle" // Base64文件上传路由路径

	return p
}

// 加载初始化路由
func (p *Image) LoadInitRoute() interface{} {
	p.GET(p.GetListPath, p.GetList)
	p.Any(p.DeletePath, p.Delete)
	p.POST(p.CropPath, p.Crop)
	p.POST(p.HandlePath, p.Handle)
	p.POST(p.Base64HandlePath, p.HandleFromBase64)

	return p
}

// 初始化
func (p *Image) Init(ctx *quark.Context) interface{} {

	// 限制文件大小
	p.LimitSize = 1024 * 1024 * 1024 * 2

	// 限制文件类型
	p.LimitType = []string{
		"image/png",
		"image/gif",
		"image/jpeg",
		"image/svg+xml",
	}

	// 设置文件上传路径
	p.SavePath = "./web/app/storage/images/" + time.Now().Format("20060102") + "/"

	return p
}

// 获取文件列表
func (p *Image) GetList(ctx *quark.Context) error {
	imageListReq := request.ImageListReq{}
	err := ctx.Bind(&imageListReq)
	if err != nil {
		return ctx.CJSONError("参数错误")
	}

	adminInfo, _ := service.NewAuthService(ctx).GetAdmin()
	images, total, err := service.NewAttachmentService().GetListBySearch(
		adminInfo.Id,
		"IMAGE",
		imageListReq.CategoryId,
		imageListReq.Name,
		imageListReq.Createtime,
		imageListReq.Page,
	)
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	categorys, err := service.NewAttachmentCategoryService().GetList(adminInfo.Id)
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	return ctx.CJSONOk("获取成功", response.ImageListResp{
		Pagination: response.Pagination{
			Current:        imageListReq.Page,
			DefaultCurrent: 1,
			PageSize:       8,
			Total:          total,
		},
		List:      images,
		Categorys: categorys,
	})
}

// 图片删除
func (p *Image) Delete(ctx *quark.Context) error {
	imageDeleteReq := request.ImageDeleteReq{}
	if err := ctx.Bind(&imageDeleteReq); err != nil {
		return ctx.CJSONError("参数错误")
	}

	err := service.NewAttachmentService().DeleteById(imageDeleteReq.Id)
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	return ctx.CJSONOk("操作成功")
}

// 图片裁剪
func (p *Image) Crop(ctx *quark.Context) error {
	var (
		result *quark.FileInfo
		err    error
	)

	imageCropReq := request.ImageCropReq{}
	if err := ctx.Bind(&imageCropReq); err != nil {
		return ctx.CJSONError(err.Error())
	}

	imageInfo, err := service.NewAttachmentService().GetInfoById(imageCropReq.Id)
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	adminInfo, err := service.NewAuthService(ctx).GetAdmin()
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	limitW := ctx.Query("limitW", "")
	limitH := ctx.Query("limitH", "")

	files := strings.Split(imageCropReq.File, ",")
	if len(files) != 2 {
		return ctx.CJSONError("格式错误")
	}

	fileData, err := base64.StdEncoding.DecodeString(files[1]) //成图片文件并把文件写入到buffer
	if err != nil {
		return ctx.CJSONError(err.Error())
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
			OSSConfig:        ossConfig.(*quark.OSSConfig),
			MinioConfig:      minioConfig.(*quark.MinioConfig),
		}).
		Reader(&quark.File{
			Content: fileData,
		})

	// 上传前回调
	getFileSystem, _, err := ctx.Template.(interface {
		BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error)
	}).BeforeHandle(ctx, fileSystem)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.CJSONError(err.Error())
	}

	filePaths := strings.Split(imageInfo.Path, "/")
	fileName := filePaths[len(filePaths)-1]
	result, err = getFileSystem.
		WithImageExtra().
		FileName(fileName).
		Path(savePath).
		Save()
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	// 重写url
	if driver == quark.LocalStorage {
		result.Url = service.NewAttachmentService().GetImageUrl(result.Url)
	}

	extra := ""
	if result.Extra != nil {
		extraData, err := json.Marshal(result.Extra)
		if err == nil {
			extra = string(extraData)
		}
	}

	// 更新数据库
	err = service.NewAttachmentService().UpdateById(imageInfo.Id, model.Attachment{
		Source: "ADMIN",
		Uid:    adminInfo.Id,
		Name:   result.Name,
		Type:   "IMAGE",
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

	return ctx.CJSONOk("操作成功", result)
}

// 上传前回调
func (p *Image) BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error) {
	fileHash, err := fileSystem.GetFileHash()
	if err != nil {
		return fileSystem, nil, err
	}

	imageInfo, err := service.NewAttachmentService().GetInfoByHash(fileHash)
	if err != nil {
		return fileSystem, nil, err
	}
	if imageInfo.Id != 0 {
		var extra map[string]interface{}
		if imageInfo.Extra != "" {
			_ = json.Unmarshal([]byte(imageInfo.Extra), &extra)
		}

		fileInfo := &quark.FileInfo{
			Name:  imageInfo.Name,
			Size:  imageInfo.Size,
			Ext:   imageInfo.Ext,
			Path:  imageInfo.Path,
			Url:   imageInfo.Url,
			Hash:  imageInfo.Hash,
			Extra: extra,
		}

		return fileSystem, fileInfo, err
	}

	return fileSystem, nil, err
}

// 上传完成后回调
func (p *Image) AfterHandle(ctx *quark.Context, result *quark.FileInfo) error {
	driver := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Driver").
		String()

	// 重写url
	if driver == quark.LocalStorage {
		result.Url = service.NewAttachmentService().GetImageUrl(result.Url)
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
		Type:   "IMAGE",
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
