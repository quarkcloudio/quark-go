package searches

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/searches"
	"gorm.io/gorm"
)

type Status struct {
	searches.Select
}

// 初始化
func (p *Status) Init() *Status {
	p.ParentInit()
	p.Name = "状态"

	return p
}

// 执行查询
func (p *Status) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {

	var status int

	if value.(string) == "on" {
		status = 1
	} else {
		status = 0
	}

	return query.Where("status = ?", status)
}

// 属性
func (p *Status) Options(ctx *builder.Context) map[interface{}]interface{} {
	return map[interface{}]interface{}{
		"on":  "正常",
		"off": "禁用",
	}
}
