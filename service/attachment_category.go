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
func (p *AttachmentCategoryService) GetAuthList(appKey string, tokenString string) (list []model.AttachmentCategory, Error error) {
	categorys := []model.AttachmentCategory{}

	adminInfo, err := NewUserService().GetAuthUser(appKey, tokenString)
	if err != nil {
		return categorys, err
	}

	err = db.Client.
		Where("source = ?", "ADMIN").
		Where("uid", adminInfo.Id).
		Find(&categorys).Error
	if err != nil {
		return categorys, err
	}

	return categorys, nil
}
