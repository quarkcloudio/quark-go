package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Tree struct {
	Item
}

// 初始化
func (p *Tree) Init() *Tree {
	p.Component = "treeField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置树形组件数据
func (p *Tree) SetData(data interface{}) *Tree {
	p.TreeData = data

	return p
}
