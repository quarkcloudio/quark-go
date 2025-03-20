package resources

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
)

type Demo struct {
	resource.Template
}

// 初始化
func (p *Demo) Init(ctx *quark.Context) interface{} {
	return p
}

// 字段
func (p *Demo) Fields(ctx *quark.Context) []interface{} {

	return []interface{}{}
}

// 搜索
func (p *Demo) Searches(ctx *quark.Context) []interface{} {

	return []interface{}{}
}

// 行为
func (p *Demo) Actions(ctx *quark.Context) []interface{} {

	return []interface{}{}
}
