package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type DataCheckbox struct {
	Item
	Value             interface{} `json:"value"`
	Localdata         interface{} `json:"localdata"`
	Mode              string      `json:"mode"`
	Multiple          bool        `json:"multiple"`
	Min               interface{} `json:"min"`
	Max               interface{} `json:"max"`
	Wrap              bool        `json:"wrap"`
	Icon              string      `json:"icon"`
	SelectedColor     string      `json:"selectedColor"`
	SelectedTextColor string      `json:"selectedTextColor"`
	EmptyText         string      `json:"emptyText"`
	Map               interface{} `json:"map"`
}

// 初始化
func (p *DataCheckbox) Init() *DataCheckbox {
	p.Component = "dataCheckboxField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.Mode = "default"
	p.Icon = "left"
	p.SelectedColor = "#007aff"
	p.SelectedTextColor = "#333"
	p.EmptyText = "暂无数据"
	p.Map = map[string]string{
		"text":  "text",
		"value": "value",
	}

	return p
}

// 默认值，multiple=true时为 Array类型，否则为 String或Number类型
func (p *DataCheckbox) SetValue(value interface{}) *DataCheckbox {
	p.Value = value

	return p
}

// 本地渲染数据
func (p *DataCheckbox) SetLocaldata(localdata interface{}) *DataCheckbox {
	p.Localdata = localdata

	return p
}

// 显示模式
func (p *DataCheckbox) SetMode(mode string) *DataCheckbox {
	p.Mode = mode

	return p
}

// 是否多选
func (p *DataCheckbox) SetMultiple(multiple bool) *DataCheckbox {
	p.Multiple = multiple

	return p
}

// 最小选择个数 ，multiple为true时生效
func (p *DataCheckbox) SetMin(min interface{}) *DataCheckbox {
	p.Min = min

	return p
}

// 最大选择个数 ，multiple为true时生效
func (p *DataCheckbox) SetMax(max interface{}) *DataCheckbox {
	p.Max = max

	return p
}

// 是否换行显示
func (p *DataCheckbox) SetWrap(wrap bool) *DataCheckbox {
	p.Wrap = wrap

	return p
}

// list 列表模式下 icon 显示的位置
func (p *DataCheckbox) SetIcon(icon string) *DataCheckbox {
	p.Icon = icon

	return p
}

// 选中颜色
func (p *DataCheckbox) SetSelectedColor(selectedColor string) *DataCheckbox {
	p.SelectedColor = selectedColor

	return p
}

// 选中文本颜色，如不填写则自动显示
func (p *DataCheckbox) SetSelectedTextColor(selectedTextColor string) *DataCheckbox {
	p.SelectedTextColor = selectedTextColor

	return p
}

// 没有数据时显示的文字 ，本地数据无效
func (p *DataCheckbox) SetEmptyText(emptyText string) *DataCheckbox {
	p.EmptyText = emptyText

	return p
}

// 字段映射，将text/value映射到数据中的其他字段
func (p *DataCheckbox) SetMap(dataCheckboxMap interface{}) *DataCheckbox {
	p.Map = dataCheckboxMap

	return p
}

// 组件json序列化
func (p *DataCheckbox) JsonSerialize() *DataCheckbox {
	p.Component = "dataCheckboxField"

	return p
}
