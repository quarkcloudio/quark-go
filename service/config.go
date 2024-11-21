package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/model"
)

type ConfigService struct{}

// 存储配置
var webConfig = make(map[string]string)

// 初始化
func NewConfigService() *ConfigService {
	return &ConfigService{}
}

// 刷新配置
func (p *ConfigService) Refresh() {
	configs := []model.Config{}
	db.Client.Where("status", 1).Find(&configs)
	for _, config := range configs {
		webConfig[config.Name] = config.Value
	}
}

// 设置配置信息
func (p *ConfigService) SetValue(key string, value string) {
	db.Client.Model(&model.Config{}).Where("name", key).Update("value", value)
	p.Refresh()
}

// 获取配置信息
func (p *ConfigService) GetValue(key string) string {
	if len(webConfig) == 0 {
		p.Refresh()
	}

	return webConfig[key]
}
