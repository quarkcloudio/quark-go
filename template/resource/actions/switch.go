package actions

import "github.com/quarkcloudio/quark-go/v4"

type Switch struct {
	Action
	CheckedChildren   interface{} `json:"checkedChildren,omitempty"`   // 选中时的内容
	UnCheckedChildren interface{} `json:"unCheckedChildren,omitempty"` // 自定义的选择框后缀图标
	FieldName         interface{} `json:"fieldName,omitempty"`         // 字段名称
	FieldValue        interface{} `json:"fieldValue,omitempty"`        // 字段值
}

// 初始化
func (p *Switch) New(ctx *quark.Context) interface{} {
	p.ActionType = "switch"
	return p
}

// 选中时的内容
func (p *Switch) GetCheckedChildren() interface{} {
	return p.CheckedChildren
}

// 自定义的选择框后缀图标
func (p *Switch) GetUnCheckedChildren() interface{} {
	return p.UnCheckedChildren
}

// 获取字段名称
func (p *Switch) GetFieldName() interface{} {
	return p.FieldName
}

// 获取字段值
func (p *Switch) GetFieldValue() interface{} {
	return p.FieldValue
}
