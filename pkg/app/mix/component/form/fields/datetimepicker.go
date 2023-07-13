package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type DatetimePicker struct {
	Item
	Type             interface{} `json:"type"`
	Value            interface{} `json:"value"`
	Start            interface{} `json:"start"`
	End              interface{} `json:"end"`
	ReturnType       string      `json:"returnType"`
	Border           bool        `json:"border"`
	RangeSeparator   string      `json:"rangeSeparator"`
	Placeholder      string      `json:"placeholder"`
	StartPlaceholder string      `json:"startPlaceholder"`
	EndPlaceholder   string      `json:"endPlaceholder"`
	Disabled         bool        `json:"disabled"`
	ClearIcon        bool        `json:"clearIcon"`
	HideSecond       bool        `json:"hideSecond"`
}

// 初始化
func (p *DatetimePicker) Init() *DatetimePicker {
	p.Component = "datetimePickerField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.Type = "datetime"
	p.ReturnType = "string"
	p.Border = true
	p.RangeSeparator = "-"
	p.Placeholder = "-"
	p.StartPlaceholder = "-"
	p.EndPlaceholder = "-"
	p.ClearIcon = true

	return p
}

// 已选择数据的 value 值
func (p *DatetimePicker) SetValue(value interface{}) *DatetimePicker {
	p.Value = value

	return p
}

// 选择器类型
func (p *DatetimePicker) SetType(pickerType interface{}) *DatetimePicker {
	p.Type = pickerType

	return p
}

// 最小值，可以使用日期的字符串(String)、时间戳(Number)
func (p *DatetimePicker) SetStart(start bool) *DatetimePicker {
	p.Start = start

	return p
}

// 最大值，可以使用日期的字符串(String)、时间戳(Number)
func (p *DatetimePicker) SetEnd(end string) *DatetimePicker {
	p.End = end

	return p
}

// 返回值格式
func (p *DatetimePicker) SetReturnType(returnType string) *DatetimePicker {
	p.ReturnType = returnType

	return p
}

// 是否有边框
func (p *DatetimePicker) SetBorder(border bool) *DatetimePicker {
	p.Border = border

	return p
}

// 选择范围时的分隔符
func (p *DatetimePicker) SetRangeSeparator(rangeSeparator string) *DatetimePicker {
	p.RangeSeparator = rangeSeparator

	return p
}

// 非范围选择时的占位内容
func (p *DatetimePicker) SetPlaceholder(placeholder string) *DatetimePicker {
	p.Placeholder = placeholder

	return p
}

// 范围选择时开始日期的占位内容
func (p *DatetimePicker) SetStartPlaceholder(startPlaceholder string) *DatetimePicker {
	p.StartPlaceholder = startPlaceholder

	return p
}

// 范围选择时结束日期的占位内容
func (p *DatetimePicker) SetEndPlaceholder(endPlaceholder string) *DatetimePicker {
	p.EndPlaceholder = endPlaceholder

	return p
}

// 是否不可选择
func (p *DatetimePicker) SetDisabled(disabled bool) *DatetimePicker {
	p.Disabled = disabled

	return p
}

// 是否显示清除按钮
func (p *DatetimePicker) SetClearIcon(clearIcon bool) *DatetimePicker {
	p.ClearIcon = clearIcon

	return p
}

// 是否显示秒，只显示时分
func (p *DatetimePicker) SetHideSecond(hideSecond bool) *DatetimePicker {
	p.HideSecond = hideSecond

	return p
}

// 组件json序列化
func (p *DatetimePicker) JsonSerialize() *DatetimePicker {
	p.Component = "datetimePickerField"

	return p
}
