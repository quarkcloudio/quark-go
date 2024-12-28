package uploads

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/dto/response"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/upload"
)

type Image struct {
	upload.Template
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

// 初始化路由映射
func (p *Image) RouteInit() interface{} {
	p.GET("/api/admin/upload/:resource/getList", p.GetList)
	p.Any("/api/admin/upload/:resource/delete", p.Delete)
	p.POST("/api/admin/upload/:resource/crop", p.Crop)
	p.POST("/api/admin/upload/:resource/handle", p.Handle)
	p.POST("/api/admin/upload/:resource/base64Handle", p.HandleFromBase64)

	return p
}

// 获取文件列表n
func (p *Image) GetList(ctx *quark.Context) error {
	page := ctx.Query("page", "1")
	categoryId := ctx.Query("categoryId", "")
	name := ctx.Query("name", "")
	startDate := ctx.Query("cteatetime[0]", "")
	endDate := ctx.Query("cteatetime[1]", "")
	currentPage, _ := strconv.Atoi(page.(string))

	pictures, total, err := service.NewPictureService().GetListBySearch(
		ctx.Engine.GetConfig().AppKey,
		ctx.Token(),
		categoryId,
		name,
		startDate,
		endDate,
		currentPage,
	)
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	pagination := map[string]interface{}{
		"defaultCurrent": 1,
		"current":        currentPage,
		"pageSize":       8,
		"total":          total,
	}

	categorys, err := service.NewPictureCategoryService().GetAuthList(ctx.Engine.GetConfig().AppKey, ctx.Token())
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("获取成功", "", map[string]interface{}{
		"pagination": pagination,
		"lists":      pictures,
		"categorys":  categorys,
	}))
}

// 图片删除
func (p *Image) Delete(ctx *quark.Context) error {
	data := map[string]interface{}{}
	json.Unmarshal(ctx.Body(), &data)
	if data["id"] == "" {
		return ctx.JSON(200, message.Error("参数错误！"))
	}

	err := service.NewPictureService().DeleteById(data["id"])
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("操作成功"))
}

// 图片裁剪
func (p *Image) Crop(ctx *quark.Context) error {
	var (
		result *quark.FileInfo
		err    error
	)

	data := map[string]interface{}{}
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if data["id"] == "" || data["file"] == "" {
		return ctx.JSON(200, message.Error("参数错误！"))
	}

	pictureInfo, err := service.NewPictureService().GetInfoById(data["id"])
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if pictureInfo.Id == 0 {
		return ctx.JSON(200, message.Error("文件不存在"))
	}

	adminInfo, err := service.NewUserService().GetAuthUser(ctx.Engine.GetConfig().AppKey, ctx.Token())
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	limitW := ctx.Query("limitW", "")
	limitH := ctx.Query("limitH", "")

	files := strings.Split(data["file"].(string), ",")
	if len(files) != 2 {
		return ctx.JSON(200, message.Error("格式错误"))
	}

	fileData, err := base64.StdEncoding.DecodeString(files[1]) //成图片文件并把文件写入到buffer
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
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
	getFileSystem, fileInfo, err := ctx.Template.(interface {
		BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error)
	}).BeforeHandle(ctx, fileSystem)
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if fileInfo != nil {
		// 更新数据库
		service.NewPictureService().UpdateById(pictureInfo.Id, model.Picture{
			ObjType: "ADMIN",
			ObjId:   adminInfo.Id,
			Name:    fileInfo.Name,
			Size:    fileInfo.Size,
			Width:   fileInfo.Width,
			Height:  fileInfo.Height,
			Ext:     fileInfo.Ext,
			Path:    fileInfo.Path,
			Url:     fileInfo.Url,
			Hash:    fileInfo.Hash,
			Status:  1,
		})
	}

	result, err = getFileSystem.
		WithImageWH().
		FileName(pictureInfo.Name).
		Path(savePath).
		Save()
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 重写url
	if driver == quark.LocalStorage {
		result.Url = service.NewPictureService().GetPath(result.Url)
	}

	// 更新数据库
	service.NewPictureService().UpdateById(pictureInfo.Id, model.Picture{
		ObjType: "ADMIN",
		ObjId:   adminInfo.Id,
		Name:    result.Name,
		Size:    result.Size,
		Width:   result.Width,
		Height:  result.Height,
		Ext:     result.Ext,
		Path:    result.Path,
		Url:     result.Url,
		Hash:    result.Hash,
		Status:  1,
	})

	return ctx.JSON(200, message.Success("操作成功", "", result))
}

// 上传前回调
func (p *Image) BeforeHandle(ctx *quark.Context, fileSystem *quark.FileSystem) (*quark.FileSystem, *quark.FileInfo, error) {
	fileHash, err := fileSystem.GetFileHash()
	if err != nil {
		return fileSystem, nil, err
	}

	pictureInfo, err := service.NewPictureService().GetInfoByHash(fileHash)
	if err != nil {
		return fileSystem, nil, err
	}
	if pictureInfo.Id != 0 {
		fileInfo := &quark.FileInfo{
			Name:   pictureInfo.Name,
			Size:   pictureInfo.Size,
			Width:  pictureInfo.Width,
			Height: pictureInfo.Height,
			Ext:    pictureInfo.Ext,
			Path:   pictureInfo.Path,
			Url:    pictureInfo.Url,
			Hash:   pictureInfo.Hash,
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
		FieldByName("Driver").String()

	// 重写url
	if driver == quark.LocalStorage {
		result.Url = service.NewPictureService().GetPath(result.Url)
	}

	adminInfo, err := service.NewUserService().GetAuthUser(ctx.Engine.GetConfig().AppKey, ctx.Token())
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 插入数据库
	id, err := service.NewPictureService().InsertGetId(model.Picture{
		ObjType: "ADMIN",
		ObjId:   adminInfo.Id,
		Name:    result.Name,
		Size:    result.Size,
		Width:   result.Width,
		Height:  result.Height,
		Ext:     result.Ext,
		Path:    result.Path,
		Url:     result.Url,
		Hash:    result.Hash,
		Status:  1,
	})

	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("上传成功", "", response.UploadImageResp{
		Id:          id,
		ContentType: result.ContentType,
		Ext:         result.Ext,
		Hash:        result.Hash,
		Name:        result.Name,
		Path:        result.Path,
		Size:        result.Size,
		Url:         result.Url,
		Height:      result.Height,
		Width:       result.Width,
	}))
}
