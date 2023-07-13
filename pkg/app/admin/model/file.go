package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"github.com/xuri/excelize/v2"
)

// 字段
type File struct {
	Id             int       `json:"id" gorm:"autoIncrement"`
	ObjType        string    `json:"obj_type" gorm:"size:255"`
	ObjId          int       `json:"obj_id" gorm:"size:11;default:0"`
	FileCategoryId int       `json:"file_category_id" gorm:"size:11;default:0"`
	Sort           int       `json:"sort" gorm:"size:11;default:0"`
	Name           string    `json:"name" gorm:"size:255;not null"`
	Size           int64     `json:"size" gorm:"size:20;default:0"`
	Ext            string    `json:"ext" gorm:"size:255"`
	Path           string    `json:"path" gorm:"size:255;not null"`
	Url            string    `json:"url" gorm:"size:255;not null"`
	Hash           string    `json:"hash" gorm:"size:255;not null"`
	Status         int       `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// 插入数据并返回ID
func (model *File) InsertGetId(data *File) (id int, Error error) {
	err := db.Client.Create(&data).Error

	return data.Id, err
}

// 根据hash查询文件信息
func (model *File) GetInfoByHash(hash string) (file *File, Error error) {
	err := db.Client.Where("status = ?", 1).Where("hash = ?", hash).First(&file).Error

	return file, err
}

// 获取文件路径
func (model *File) GetPath(id interface{}) string {
	http, path := "", ""
	webSiteDomain := (&Config{}).GetValue("WEB_SITE_DOMAIN")
	WebConfig := (&Config{}).GetValue("SSL_OPEN")
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
			// 如果为map
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
			return http + webSiteDomain + path
		}
	}

	file := &File{}
	db.Client.Where("id", id).Where("status", 1).First(&file)
	if file.Id != 0 {
		path = file.Url
		if strings.Contains(path, "//") {
			return path
		}
		if strings.Contains(path, "./") {
			path = strings.Replace(path, "./web/app/", "/", -1)
		}
		if path != "" {
			return http + webSiteDomain + path
		}
	}

	return ""
}

// 获取多文件路径
func (model *File) GetPaths(id interface{}) []string {
	var paths []string
	http, path := "", ""
	webSiteDomain := (&Config{}).GetValue("WEB_SITE_DOMAIN")
	WebConfig := (&Config{}).GetValue("SSL_OPEN")
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
func (model *File) GetExcelData(fileId int) (data [][]interface{}, Error error) {
	file := &File{}
	err := db.Client.Where("id", fileId).Where("status", 1).First(&file).Error
	if err != nil {
		return data, err
	}
	if file.Id == 0 {
		return data, errors.New("参数错误！")
	}

	f, err := excelize.OpenFile(file.Path + file.Name)
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
