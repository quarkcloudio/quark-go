package service

import (
	"sync"

	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
)

type ConfigService struct{}

// 存储配置
var webConfig = make(map[string]string)
var mu sync.Mutex

// 初始化
func NewConfigService() *ConfigService {
	return &ConfigService{}
}

// 刷新配置
func (p *ConfigService) Refresh() {
	configs := []model.Config{}
	db.Client.Where("status", 1).Find(&configs)
	// 确保对 webConfig map 的写入操作是互斥的，防止并发写入导致的数据竞争
	mu.Lock()
	defer mu.Unlock()
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
