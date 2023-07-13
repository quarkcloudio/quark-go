package calendar

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Name              string      `json:"name"`
	Required          bool        `json:"required,omitempty"`
	Prop              string      `json:"prop,omitempty"`
	Rules             interface{} `json:"rules,omitempty"`
	Label             string      `json:"label,omitempty"`
	LabelWidth        int         `json:"labelWidth,omitempty"`
	LabelAlign        string      `json:"labelAlign,omitempty"`
	BodyAlign         string      `json:"bodyAlign,omitempty"`
	ErrorMessageAlign string      `json:"errorMessageAlign,omitempty"`
	ShowErrorLine     bool        `json:"showErrorLine,omitempty"`
	ShowErrorMessage  bool        `json:"showErrorMessage,omitempty"`

	Value           interface{} `json:"value,omitempty"`
	Type            string      `json:"type,omitempty"`
	Poppable        bool        `json:"poppable,omitempty"`
	IsAutoBackFill  bool        `json:"isAutoBackFill,omitempty"`
	Title           string      `json:"title,omitempty"`
	StartDate       string      `json:"startDate,omitempty"`
	EndDate         string      `json:"endDate,omitempty"`
	ShowToday       bool        `json:"showToday,omitempty"`
	StartText       string      `json:"startText,omitempty"`
	EndText         string      `json:"endText,omitempty"`
	ConfirmText     string      `json:"confirmText,omitempty"`
	ShowTitle       bool        `json:"showTitle,omitempty"`
	ShowSubTitle    bool        `json:"showSubTitle,omitempty"`
	ToDateAnimation bool        `json:"toDateAnimation,omitempty"`
	FirstDayOfWeek  int         `json:"firstDayOfWeek,omitempty"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "calendarField"
	p.SetKey("calendar", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 表单域 v-model 字段， 在使用表单校验功能的情况下，该属性是必填的
func (p *Component) SetProp(prop string) *Component {
	p.Name = prop
	p.Prop = prop

	return p
}

// 表单域 v-model 字段， 在使用表单校验功能的情况下，该属性是必填的
func (p *Component) SetName(name string) *Component {
	p.Name = name
	p.Prop = name

	return p
}

// 定义校验规则
func (p *Component) SetRules(rules interface{}) *Component {
	p.Rules = rules

	return p
}

// 是否显示必填字段的标签旁边的红色星号
func (p *Component) SetRequired(required bool) *Component {
	p.Required = required

	return p
}

// 输入框左边的文字提示
func (p *Component) SetLabel(label string) *Component {
	p.Label = label

	return p
}

// 表单项 label 宽度，默认单位为px
func (p *Component) SetLabelWidth(labelWidth int) *Component {
	p.LabelWidth = labelWidth

	return p
}

// 表单项 label 对齐方式，可选值为 center right
func (p *Component) SetLabelAlign(labelAlign string) *Component {
	p.LabelAlign = labelAlign

	return p
}

// 右侧插槽对齐方式，可选值为 center right
func (p *Component) SetBodyAlign(bodyAlign string) *Component {
	p.BodyAlign = bodyAlign

	return p
}

// 错误提示文案对齐方式，可选值为 center right
func (p *Component) SetErrorMessageAlign(errorMessageAlign string) *Component {
	p.ErrorMessageAlign = errorMessageAlign

	return p
}

// 是否在校验不通过时标红输入框
func (p *Component) SetShowErrorLine(showErrorLine bool) *Component {
	p.ShowErrorLine = showErrorLine

	return p
}

// 是否在校验不通过时在输入框下方展示错误提示
func (p *Component) SetShowErrorMessage(showErrorMessage bool) *Component {
	p.ShowErrorMessage = showErrorMessage

	return p
}

// 默认值
func (p *Component) SetValue(value []interface{}) *Component {
	p.Value = value

	return p
}

// 类型，日期单择one，区间选择range,日期多选multiple,周选择week(v4.0.1)
func (p *Component) SetCalendarType(calendarType string) *Component {
	p.Type = calendarType

	return p
}

// 是否弹窗状态展示
func (p *Component) SetPoppable(poppable bool) *Component {
	p.Poppable = poppable

	return p
}

// 自动回填
func (p *Component) SeIsAutoBackFill(isAutoBackFill bool) *Component {
	p.IsAutoBackFill = isAutoBackFill

	return p
}

// 显示标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title

	return p
}

// 开始日期
func (p *Component) SetStartDate(startDate string) *Component {
	p.StartDate = startDate

	return p
}

// 结束日期
func (p *Component) SetEndDate(endDate string) *Component {
	p.EndDate = endDate

	return p
}

// 是否展示今天标记
func (p *Component) SetShowToday(showToday bool) *Component {
	p.ShowToday = showToday

	return p
}

// 范围选择，开始信息文案
func (p *Component) SetStartText(startText string) *Component {
	p.StartText = startText

	return p
}

// 范围选择，结束信息文案
func (p *Component) SetEndText(endText string) *Component {
	p.EndText = endText

	return p
}

// 底部确认按钮文案
func (p *Component) SetConfirmText(confirmText string) *Component {
	p.ConfirmText = confirmText

	return p
}

// 是否在展示日历标题
func (p *Component) SetShowTitle(showTitle bool) *Component {
	p.ShowTitle = showTitle

	return p
}

// 是否展示日期标题
func (p *Component) SetShowSubTitle(showSubTitle bool) *Component {
	p.ShowSubTitle = showSubTitle

	return p
}

// 是否启动滚动动画
func (p *Component) SetToDateAnimation(toDateAnimation bool) *Component {
	p.ToDateAnimation = toDateAnimation

	return p
}

// 是否启动滚动动画
func (p *Component) SetFirstDayOfWeek(firstDayOfWeek int) *Component {
	p.FirstDayOfWeek = firstDayOfWeek

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "calendarField"

	return p
}
