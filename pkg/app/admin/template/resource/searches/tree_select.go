package searches

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/treeselect"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type TreeSelect struct {
	Search
	TreeSelectOptions []*treeselect.TreeData
}

// 初始化模板
func (p *TreeSelect) TemplateInit(ctx *builder.Context) interface{} {
	p.Component = "treeSelectField"

	return p
}

// 设置Option
func (p *TreeSelect) Option(value interface{}, title string) *treeselect.TreeData {

	return &treeselect.TreeData{
		Value: value,
		Title: title,
	}
}
