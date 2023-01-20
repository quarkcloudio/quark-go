package uploads

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminupload"
	"github.com/quarkcms/quark-go/pkg/msg"
	"github.com/quarkcms/quark-go/pkg/storage"
)

type Image struct {
	adminupload.Template
}

// 初始化
func (p *Image) Init() interface{} {

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

	// 设置文件上传路径
	p.SavePath = "./website/storage/images/" + time.Now().Format("20060102") + "/"

	// 注册路由
	p.AddRoute("/api/admin/upload/:resource/getList", "GetList")
	p.AddRoute("/api/admin/upload/:resource/delete", "Delete")
	p.AddRoute("/api/admin/upload/:resource/crop", "Crop")

	return p
}

// 获取文件列表
func (p *Image) GetList(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {
	page := request.Query("page", "1")
	categoryId := request.Query("pictureCategoryId", "")
	searchName := request.Query("pictureSearchName", "")
	searchDateStart := request.Query("pictureSearchDate[0]", "")
	searchDateEnd := request.Query("pictureSearchDate[1]", "")
	currentPage, _ := strconv.Atoi(page.(string))

	pictures, total, err := (&model.Picture{}).GetListBySearch(
		request.Token(),
		categoryId, searchName,
		searchDateStart,
		searchDateEnd,
		currentPage,
	)
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	pagination := map[string]interface{}{
		"defaultCurrent": 1,
		"current":        currentPage,
		"pageSize":       12,
		"total":          total,
	}

	categorys, err := (&model.PictureCategory{}).GetAuthList(request.Token())
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	return msg.Success("上传成功", "", map[string]interface{}{
		"pagination": pagination,
		"lists":      pictures,
		"categorys":  categorys,
	})
}

// 图片删除
func (p *Image) Delete(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {
	data := map[string]interface{}{}
	json.Unmarshal(request.Body(), &data)
	if data["id"] == "" {
		return msg.Error("参数错误！", "")
	}

	err := (&model.Picture{}).DeleteById(data["id"])
	if err != nil {
		return msg.Error(err.Error(), "")
	} else {
		return msg.Success("操作成功！", "", "")
	}
}

// 图片裁剪
func (p *Image) Crop(request *builder.Request, resource *builder.Resource, templateInstance interface{}) interface{} {
	var (
		result *storage.FileInfo
		err    error
	)

	data := map[string]interface{}{}
	if err := request.BodyParser(&data); err != nil {
		return msg.Error(err.Error(), "")
	}
	if data["id"] == "" {
		return msg.Error("参数错误！", "")
	}
	if data["file"] == "" {
		return msg.Error("参数错误！", "")
	}

	limitW := request.Query("limitW", "")
	limitH := request.Query("limitH", "")

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

	// 重写url
	if driver == storage.LocalDriver {
		result.Url = (&model.Picture{}).GetPath(result.Url)
	}

	adminInfo, err := (&model.Admin{}).GetAuthUser(request.Token())
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	// 插入数据库
	(&model.Picture{}).InsertGetId(&model.Picture{
		ObjType: "ADMINID",
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

	return msg.Success("裁剪成功", "", result)
}

// 上传前回调
func (p *Image) BeforeHandle(request *builder.Request, templateInstance interface{}, fileSystem *storage.FileSystem) (*storage.FileSystem, *storage.FileInfo, error) {
	fileHash, err := fileSystem.GetFileHash()
	if err != nil {
		return fileSystem, nil, err
	}

	pictureInfo, _ := (&model.Picture{}).GetInfoByHash(fileHash)
	if err != nil {
		return fileSystem, nil, err
	}
	if pictureInfo.Id != 0 {
		fileInfo := &storage.FileInfo{
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
func (p *Image) AfterHandle(request *builder.Request, templateInstance interface{}, result *storage.FileInfo) interface{} {
	driver := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Driver").String()

	// 重写url
	if driver == storage.LocalDriver {
		result.Url = (&model.Picture{}).GetPath(result.Url)
	}

	adminInfo, err := (&model.Admin{}).GetAuthUser(request.Token())
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	// 插入数据库
	(&model.Picture{}).InsertGetId(&model.Picture{
		ObjType: "ADMINID",
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

	return msg.Success("上传成功", "", result)
}
