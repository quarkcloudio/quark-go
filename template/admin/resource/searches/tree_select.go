package searches

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/treeselect"
)

type TreeSelect struct {
	Search
	TreeSelectOptions []treeselect.TreeData
}

// 初始化模板
func (p *TreeSelect) LoadInitData(ctx *quark.Context) interface{} {
	p.Component = "treeSelectField"

	return p
}

// 设置Option
func (p *TreeSelect) Option(title string, value interface{}) treeselect.TreeData {

	return treeselect.TreeData{
		Value: value,
		Title: title,
	}
}

// 可选项数据源
//
//	SetTreeData([]treeselect.TreeData {{Value :"zhejiang", Title:"Zhejiang"}})
//
// 或者
//
// SetTreeData(options, "parent_key_name", "title_name", "value_name")
//
// 或者
//
// SetTreeData(options, 0, "parent_key_name", "title_name", "value_name")
func (p *TreeSelect) SetTreeData(treeData ...interface{}) *TreeSelect {
	if len(treeData) == 1 {
		getOptions, ok := treeData[0].([]treeselect.TreeData)
		if ok {
			p.TreeSelectOptions = getOptions
			return p
		}
	}
	if len(treeData) == 4 {
		p.TreeSelectOptions = treeselect.New().ListToTreeData(treeData[0], 0, treeData[1].(string), treeData[2].(string), treeData[3].(string))
	}
	if len(treeData) == 5 {
		p.TreeSelectOptions = treeselect.New().ListToTreeData(treeData[0], treeData[1].(int), treeData[2].(string), treeData[3].(string), treeData[4].(string))
	}
	return p
}
