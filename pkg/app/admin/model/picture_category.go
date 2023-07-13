package model

import "github.com/quarkcms/quark-go/v2/pkg/dal/db"

// 字段
type PictureCategory struct {
	Id          int    `json:"id" gorm:"autoIncrement"`
	ObjType     string `json:"obj_type" gorm:"size:100"`
	ObjId       int    `json:"obj_id" gorm:"size:11;default:0"`
	Title       string `json:"title" gorm:"size:255;not null"`
	Sort        int    `json:"sort" gorm:"size:11;default:0"`
	Description string `json:"description" gorm:"size:255"`
}

// 获取列表
func (model *PictureCategory) GetAuthList(appKey string, tokenString string) (list []*PictureCategory, Error error) {
	categorys := []*PictureCategory{}

	adminInfo, err := (&Admin{}).GetAuthUser(appKey, tokenString)
	if err != nil {
		return categorys, err
	}

	err = db.Client.
		Where("obj_type = ?", "ADMINID").
		Where("obj_id", adminInfo.Id).
		Find(&categorys).Error
	if err != nil {
		return categorys, err
	}

	return categorys, nil
}
