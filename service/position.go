package service

import (
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/checkbox"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
)

type PositionService struct{}

// 初始化
func NewPositionService() *PositionService {
	return &PositionService{}
}

// 获取职位列表
func (p PositionService) List() (list []checkbox.Option, Error error) {
	positions := []model.Position{}
	err := db.Client.Where("status = ?", 1).Find(&positions).Error
	if err != nil {
		return list, err
	}
	for _, v := range positions {
		list = append(list, checkbox.Option{
			Label: v.Name,
			Value: v.Id,
		})
	}
	return list, nil
}
