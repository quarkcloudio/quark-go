package searches

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/searches"
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
func (p *Status) Options(ctx *builder.Context) interface{} {

	return []*searches.SelectOption{
		p.Option(0, "禁用"),
		p.Option(1, "正常"),
	}
}
