package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
)

type AttachmentCategoryService struct{}

// 初始化
func NewAttachmentCategoryService() *AttachmentCategoryService {
	return &AttachmentCategoryService{}
}

// 获取列表
func (p *AttachmentCategoryService) GetList(adminId interface{}) (list []model.AttachmentCategory, Error error) {
	categorys := []model.AttachmentCategory{}
	err := db.Client.
		Where("source = ?", "ADMIN").
		Where("uid", adminId).
		Find(&categorys).Error
	if err != nil {
		return categorys, err
	}

	return categorys, nil
}
