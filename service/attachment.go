package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/xuri/excelize/v2"
)

type AttachmentService struct{}

// 初始化
func NewAttachmentService() *AttachmentService {
	return &AttachmentService{}
}

// 获取列表
func (p *AttachmentService) GetListBySearch(adminId interface{}, attachmentType string, categoryId interface{}, name interface{}, createtime []string, page int) (list []model.Attachment, total int64, Error error) {
	attachments := []model.Attachment{}

	query := db.Client.Model(&model.Attachment{}).
		Where("status =?", 1).
		Where("source = ?", "ADMIN").
		Where("uid", adminId)

	if categoryId != "" {
		query.Where("category_id =?", categoryId)
	}
	if name != "" {
		query.Where("name LIKE %?%", name)
	}
	if createtime != nil && len(createtime) == 2 {
		query.Where("created_at BETWEEN ? AND ?", createtime[0], createtime[1])
	}

	query.Count(&total)
	query.
		Order("id desc").
		Limit(8).
		Offset((page - 1) * 8).
		Find(&attachments)

	for k, v := range attachments {
		v.Url = p.GetPath(v.Url) + "?timestamp=" + strconv.Itoa(int(time.Now().Unix()))
		attachments[k] = v
	}

	return attachments, total, nil
}

// 插入数据并返回ID
func (p *AttachmentService) InsertGetId(attachment model.Attachment) (id int, Error error) {
	err := db.Client.Create(&attachment).Error
	if err != nil {
		return id, err
	}

	return attachment.Id, nil
}

// 通过Id删除记录
func (p *AttachmentService) DeleteById(id interface{}) error {

	return db.Client.Model(model.Attachment{}).Where("id =?", id).Delete("").Error
}

// 根据id查询文件信息
func (p *AttachmentService) GetInfoById(id interface{}) (attachment model.Attachment, Error error) {
	err := db.Client.Where("status = ?", 1).Where("id = ?", id).First(&attachment).Error

	return attachment, err
}

// 根据id更新文件信息
func (p *AttachmentService) UpdateById(id interface{}, data model.Attachment) (Error error) {
	err := db.Client.Where("status = ?", 1).Where("id = ?", id).Updates(&data).Error

	return err
}

// 根据hash查询文件信息
func (p *AttachmentService) GetInfoByHash(hash string) (attachment model.Attachment, Error error) {
	err := db.Client.Where("status = ?", 1).Where("hash = ?", hash).First(&attachment).Error

	return attachment, err
}

// 获取附件路径，GetPath(1) 或者 GetPath("FILE", 1)
func (p *AttachmentService) GetPath(params ...interface{}) string {
	var id, attachmentType interface{}
	if len(params) == 1 {
		id = params[0]
	}
	if len(params) == 2 {
		attachmentType = params[0].(string)
		id = params[1]
	}

	http, path := "", ""
	webSiteDomain := NewConfigService().GetValue("WEB_SITE_DOMAIN")
	WebConfig := NewConfigService().GetValue("SSL_OPEN")
	if webSiteDomain != "" {
		if WebConfig == "1" {
			http = "https://"
		} else {
			http = "http://"
		}
	}

	if getId, ok := id.(string); ok {
		if strings.Contains(getId, "//") && !strings.Contains(getId, "{") {
			return getId
		}
		if strings.Contains(getId, "./") && !strings.Contains(getId, "{") {
			return http + webSiteDomain + strings.Replace(getId, "./web/app/", "/", -1)
		}
		if strings.Contains(getId, "/") && !strings.Contains(getId, "{") {
			return http + webSiteDomain + getId
		}

		// json字符串
		if strings.Contains(getId, "{") {
			var jsonData interface{}
			json.Unmarshal([]byte(getId), &jsonData)
			if mapData, ok := jsonData.(map[string]interface{}); ok {
				path = mapData["url"].(string)
			}

			// 如果为数组，返回第一个key的path
			if arrayData, ok := jsonData.([]map[string]interface{}); ok {
				path = arrayData[0]["url"].(string)
			}
		}
		if strings.Contains(path, "//") {
			return path
		}
		if strings.Contains(path, "./") {
			path = strings.Replace(path, "./web/app/", "/", -1)
		}
		if path != "" {
			// 如果设置域名，则加上域名前缀
			return http + webSiteDomain + path
		}
	}

	attachment := model.Attachment{}
	db.Client.Where("id", id).Where("status", 1).First(&attachment)
	if attachment.Id != 0 {
		path = attachment.Url
		if strings.Contains(path, "//") {
			return path
		}
		if strings.Contains(path, "./") {
			path = strings.Replace(path, "./web/app/", "/", -1)
		}
	}
	if path != "" {
		// 如果设置域名，则加上域名前缀
		return http + webSiteDomain + path
	}
	if attachmentType == "IMAGE" {
		return http + webSiteDomain + "/admin/default.png"
	}
	return ""
}

// 获取文件路径
func (p *AttachmentService) GetFilePath(id interface{}) string {
	return p.GetPath("FILE", id)
}

// 获取图片路径
func (p *AttachmentService) GetImagePath(id interface{}) string {
	return p.GetPath("IMAGE", id)
}

// 获取多文件路径
func (p *AttachmentService) GetPaths(id interface{}) []string {
	var paths []string
	http, path := "", ""
	webSiteDomain := NewConfigService().GetValue("WEB_SITE_DOMAIN")
	WebConfig := NewConfigService().GetValue("SSL_OPEN")
	if webSiteDomain != "" {
		if WebConfig == "1" {
			http = "https://"
		} else {
			http = "http://"
		}
	}

	if getId, ok := id.(string); ok {
		// json字符串
		if strings.Contains(getId, "{") {
			var jsonData []map[string]interface{}
			err := json.Unmarshal([]byte(getId), &jsonData)
			if err == nil {
				for _, v := range jsonData {
					path = v["url"].(string)
					if strings.Contains(path, "//") {
						paths = append(paths, v["url"].(string))
					} else {
						if strings.Contains(path, "./") {
							path = strings.Replace(path, "./web/app/", "/", -1)
						}
						if path != "" {
							path = http + webSiteDomain + path
						}
						paths = append(paths, path)
					}
				}
			}
		}
	}

	return paths
}

// 获取Excel文件数据
func (p *AttachmentService) GetExcelData(fileId int) (data [][]interface{}, Error error) {
	file := model.Attachment{}
	err := db.Client.Where("id", fileId).Where("status", 1).First(&file).Error
	if err != nil {
		return data, err
	}
	if file.Id == 0 {
		return data, errors.New("param error")
	}

	f, err := excelize.OpenFile(file.Path)
	if err != nil {
		return data, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return data, err
	}

	for _, row := range rows {
		getRows := []interface{}{}
		for _, colCell := range row {
			getRows = append(getRows, colCell)
		}
		data = append(data, getRows)
	}

	return data, err
}
