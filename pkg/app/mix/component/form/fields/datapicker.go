package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type DataPicker struct {
	Item
	Value      interface{} `json:"value"`
	Localdata  interface{} `json:"localdata"`
	Preload    bool        `json:"preload"`
	Readonly   bool        `json:"readonly"`
	ClearIcon  bool        `json:"clearIcon"`
	Ellipsis   bool        `json:"ellipsis"`
	PopupTitle string      `json:"popupTitle"`
	Map        interface{} `json:"map"`
}

// 初始化
func (p *DataPicker) Init() *DataPicker {
	p.Component = "dataPickerField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.ClearIcon = true
	p.Ellipsis = true
	p.Map = map[string]string{
		"text":  "text",
		"value": "value",
	}

	return p
}

// 绑定数据
func (p *DataPicker) SetValue(value interface{}) *DataPicker {
	p.Value = value

	return p
}

// 数据
func (p *DataPicker) SetLocaldata(localdata interface{}) *DataPicker {
	p.Localdata = localdata

	return p
}

// 预加载数据
func (p *DataPicker) SetPreload(preload bool) *DataPicker {
	p.Preload = preload

	return p
}

// 是否禁用
func (p *DataPicker) SetReadonly(readonly bool) *DataPicker {
	p.Readonly = readonly

	return p
}

// 是否显示清除按钮
func (p *DataPicker) SetClearIcon(clearIcon bool) *DataPicker {
	p.ClearIcon = clearIcon

	return p
}

// 是否隐藏 tab 标签过长的文本
func (p *DataPicker) SetEllipsis(ellipsis bool) *DataPicker {
	p.Ellipsis = ellipsis

	return p
}

// 弹出层标题
func (p *DataPicker) SetPopupTitle(popupTitle string) *DataPicker {
	p.PopupTitle = popupTitle

	return p
}

// 字段映射，将text/value映射到数据中的其他字段
func (p *DataPicker) SetMap(dataCheckboxMap interface{}) *DataPicker {
	p.Map = dataCheckboxMap

	return p
}

// 组件json序列化
func (p *DataPicker) JsonSerialize() *DataPicker {
	p.Component = "dataPickerField"

	return p
}
