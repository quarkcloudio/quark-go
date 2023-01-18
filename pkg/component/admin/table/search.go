package table

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Search struct {
	component.Element
	DateFormatter     string        `json:"dateFormatter"`
	LabelAlign        string        `json:"labelAlign"`
	Size              string        `json:"size"`
	DefaultCollapsed  bool          `json:"defaultCollapsed"`
	HideRequiredMark  bool          `json:"hideRequiredMark"`
	DefaultColsNumber int           `json:"defaultColsNumber"`
	LabelWidth        int           `json:"labelWidth"`
	Span              int           `json:"span"`
	Split             bool          `json:"split"`
	ShowSubmitButton  bool          `json:"showSubmitButton"`
	ShowResetButton   bool          `json:"showResetButton"`
	ShowExportButton  bool          `json:"showExportButton"`
	ResetButton       string        `json:"resetButton"`
	SubmitButton      string        `json:"submitButton"`
	ExportButton      string        `json:"exportButton"`
	ExportApi         string        `json:"exportApi"`
	Items             []interface{} `json:"items"`
}

// 初始化
func (p *Search) Init() *Search {
	p.Component = "search"
	p.DateFormatter = "string"
	p.LabelAlign = "right"
	p.Size = "default"
	p.DefaultCollapsed = true
	p.HideRequiredMark = true
	p.DefaultColsNumber = 2
	p.LabelWidth = 98
	p.ShowSubmitButton = true
	p.ShowResetButton = true
	p.ResetButton = "重置"
	p.SubmitButton = "查询"
	p.ExportButton = "导出数据"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Search) SetStyle(style map[string]interface{}) *Search {
	p.Style = style

	return p
}

/**
 * 自动格式数据，例如 moment 的表单,支持 string 和 number 两种模式
 *
 * @param string dateFormatter
 * @return p
 */
func (p *Search) SetDateFormatter(dateFormatter string) *Search {
	p.DateFormatter = dateFormatter

	return p
}

/**
 * 默认状态下是否折叠超出的表单项
 *
 * @param bool collapsed
 * @return p
 */
func (p *Search) SetCollapsed(collapsed bool) *Search {
	p.DefaultCollapsed = collapsed

	return p
}

/**
 * label 标签的文本对齐方式，left | right
 *
 * @param string labelAlign
 * @return p
 */
func (p *Search) SetLabelAlign(labelAlign string) *Search {
	p.LabelAlign = labelAlign

	return p
}

/**
 * 设置字段组件的尺寸（仅限 antd 组件）,small | middle | large
 *
 * @param string size
 * @return p
 */
func (p *Search) SetSize(size string) *Search {
	p.Size = size

	return p
}

/**
 * 隐藏所有表单项的必选标记，默认隐藏
 *
 * @param bool hideRequiredMark
 * @return p
 */
func (p *Search) SetHideRequiredMark(hideRequiredMark bool) *Search {
	p.HideRequiredMark = hideRequiredMark

	return p
}

/**
 * 自定义折叠状态下默认显示的表单控件数量，没有设置或小于 0，则显示一行控件 数量大于等于控件数量则隐藏展开按钮
 *
 * @param number defaultColsNumber
 * @return p
 */
func (p *Search) SetDefaultColsNumber(defaultColsNumber int) *Search {
	p.DefaultColsNumber = defaultColsNumber

	return p
}

/**
 * label 宽度,number | 'auto'
 *
 * @param number|string labelWidth
 * @return p
 */
func (p *Search) SetLabelWidth(labelWidth int) *Search {
	p.LabelWidth = labelWidth

	return p
}

/**
 * 表单项宽度,number[0 - 24]
 *
 * @param number span
 * @return p
 */
func (p *Search) SetSpan(span int) *Search {
	p.Span = span

	return p
}

/**
 * 每一行是否有分割线
 *
 * @param bool split
 * @return p
 */
func (p *Search) SetSplit(split bool) *Search {
	p.Split = split

	return p
}

/**
 * 是否显示提交按钮
 *
 * @param bool showSubmitButton
 * @return p
 */
func (p *Search) SetShowSubmitButton(showSubmitButton bool) *Search {
	p.ShowSubmitButton = showSubmitButton

	return p
}

/**
 * 是否显示重置按钮
 *
 * @param bool showResetButton
 * @return p
 */
func (p *Search) SetShowResetButton(showResetButton bool) *Search {
	p.ShowResetButton = showResetButton

	return p
}

/**
 * 是否显示导出按钮
 *
 * @param bool showExportButton
 * @return p
 */
func (p *Search) SetShowExportButton(showExportButton bool) *Search {
	p.ShowExportButton = showExportButton

	return p
}

/**
 * 提交按钮文字
 *
 * @param string submitButton
 * @return p
 */
func (p *Search) SetSubmitButton(submitButton string) *Search {
	p.SubmitButton = submitButton

	return p
}

/**
 * 重置按钮文字
 *
 * @param string resetButton
 * @return p
 */
func (p *Search) SetResetButton(resetButton string) *Search {
	p.ResetButton = resetButton

	return p
}

/**
 * 导出按钮文字
 *
 * @param string exportButton
 * @return p
 */
func (p *Search) SetExportButton(exportButton string) *Search {
	p.ExportButton = exportButton

	return p
}

/**
 * 导出数据接口
 *
 * @param string exportApi
 * @return p
 */
func (p *Search) SetExportApi(exportApi string) *Search {
	p.ExportApi = exportApi

	return p
}

/**
 * 设置搜索表单项
 *
 * @param string exportApi
 * @return p
 */
func (p *Search) SetItems(item interface{}) *Search {
	p.Items = append(p.Items, item)

	return p
}

// 组件json序列化
func (p *Search) JsonSerialize() *Search {
	p.Component = "search"

	return p
}
