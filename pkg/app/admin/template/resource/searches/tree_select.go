package searches

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/treeselect"

type TreeSelect struct {
	Search
}

// 初始化
func (p *TreeSelect) ParentInit() interface{} {
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
