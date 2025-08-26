package service

import (
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
)

type ActionLogService struct{}

// 初始化
func NewActionLogService() *ActionLogService {
	return &ActionLogService{}
}

// 插入数据
func (p *ActionLogService) InsertGetId(data model.ActionLog) (id int, Error error) {
	err := db.Client.Create(&data).Error

	return data.Id, err
}
