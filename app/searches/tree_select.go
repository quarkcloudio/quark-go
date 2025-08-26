package searches

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/treeselect"
	"github.com/quarkcloudio/quark-go/v4/template/resource/searches"
	"gorm.io/gorm"
)

type TreeSelectField struct {
	searches.TreeSelect
	TreeSelectOptions []treeselect.TreeData
}

// 树形下拉框
func TreeSelect(column string, name string) *TreeSelectField {
	field := &TreeSelectField{}

	field.Column = column
	field.Name = name

	return field
}

// 执行查询
func (p *TreeSelectField) Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where(p.Column+" = ?", value)
}

// 属性
func (p *TreeSelectField) Options(ctx *quark.Context) interface{} {
	return p.TreeSelectOptions
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
func (p *TreeSelectField) SetTreeData(treeData ...interface{}) *TreeSelectField {
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
