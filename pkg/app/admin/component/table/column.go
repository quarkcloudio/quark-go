package table

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"
)

type Column struct {
	component.Element
	ComponentKey  string      `json:"componentKey"`
	Title         string      `json:"title,omitempty"`
	Attribute     string      `json:"attribute,omitempty"`
	Align         string      `json:"align,omitempty"`
	DataIndex     string      `json:"dataIndex,omitempty"`
	Fixed         interface{} `json:"fixed,omitempty"`
	Tooltip       string      `json:"tooltip,omitempty"`
	Ellipsis      bool        `json:"ellipsis,omitempty"`
	Copyable      bool        `json:"copyable,omitempty"`
	ValueEnum     interface{} `json:"valueEnum,omitempty"`
	ValueType     string      `json:"valueType,omitempty"`
	HideInSearch  bool        `json:"hideInSearch,omitempty"`
	HideInTable   bool        `json:"hideInTable,omitempty"`
	HideInForm    bool        `json:"hideInForm,omitempty"`
	Filters       interface{} `json:"filters,omitempty"`
	Order         int         `json:"order,omitempty"`
	Sorter        interface{} `json:"sorter,omitempty"`
	Span          int         `json:"span,omitempty"`
	Width         int         `json:"width,omitempty"`
	Editable      interface{} `json:"editable,omitempty"`
	Actions       interface{} `json:"actions,omitempty"`
	FormItemProps interface{} `json:"formItemProps,omitempty"`
	FieldProps    interface{} `json:"fieldProps,omitempty"`
}

// 初始化
func (p *Column) Init() *Column {
	p.Component = "column"
	p.Align = "left"
	p.Editable = false
	p.Actions = false
	p.Filters = false
	p.HideInSearch = true

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Column) SetStyle(style map[string]interface{}) *Column {
	p.Style = style

	return p
}

// 列头显示文字，既字段的列名
func (p *Column) SetTitle(title string) *Column {
	p.Title = title
	return p
}

// 字段名称|字段的列名
func (p *Column) SetAttribute(attribute string) *Column {
	p.SetKey(attribute, false)
	p.ComponentKey = attribute
	p.DataIndex = attribute
	p.Attribute = attribute

	return p
}

// 设置列的对齐方式,left | right | center
func (p *Column) SetAlign(align string) *Column {
	p.Align = align
	return p
}

// （IE 下无效）列是否固定，可选 true (等效于 left) left right
func (p *Column) SetFixed(fixed interface{}) *Column {
	p.Fixed = fixed
	return p
}

// 会在 title 之后展示一个 icon，hover 之后提示一些信息
func (p *Column) SetTooltip(tooltip string) *Column {
	p.Tooltip = tooltip
	return p
}

// 是否自动缩略
func (p *Column) SetEllipsis(ellipsis bool) *Column {
	p.Ellipsis = ellipsis
	return p
}

// 是否支持复制
func (p *Column) SetCopyable(copyable bool) *Column {
	p.Copyable = copyable
	return p
}

// 值的枚举，会自动转化把值当成 key 来取出要显示的内容
func (p *Column) SetValueEnum(valueEnum interface{}) *Column {
	var (
		valueEnumStr = map[string]interface{}{}
		valueEnumInt = map[int]interface{}{}
	)

	if getValueEnum, ok := valueEnum.(map[interface{}]interface{}); ok {
		for k, v := range getValueEnum {
			switch k.(type) {
			case string:
				valueEnumStr[k.(string)] = v
			case int:
				valueEnumInt[k.(int)] = v
			case int64:
				valueEnumInt[int(k.(int64))] = v
			case float32:
				valueEnumInt[int(k.(float32))] = v
			case float64:
				valueEnumInt[int(k.(float64))] = v
			}
		}

		if len(valueEnumStr) > 0 {
			p.ValueEnum = valueEnumStr
		}

		if len(valueEnumInt) > 0 {
			p.ValueEnum = valueEnumInt
		}
	} else {
		p.ValueEnum = valueEnum
	}

	return p
}

// 值的类型,"money" | "option" | "date" | "dateTime" | "time" | "text"| "index" | "indexBorder"
func (p *Column) SetValueType(valueType string) *Column {
	p.ValueType = valueType
	return p
}

// 在查询表单中不展示此项
func (p *Column) SetHideInSearch(hideInSearch bool) *Column {
	p.HideInSearch = hideInSearch
	return p
}

// 在 Table 中不展示此列
func (p *Column) SetHideInTable(hideInTable bool) *Column {
	p.HideInTable = hideInTable
	return p
}

// 表头的筛选菜单项，当值为 true 时，自动使用 valueEnum 生成
func (p *Column) SetFilters(filters interface{}) *Column {
	getFilters, ok := filters.(map[string]string)

	if ok {
		tmpFilters := []map[string]string{}
		for k, v := range getFilters {
			tmpFilters = append(tmpFilters, map[string]string{
				"text":  v,
				"value": k,
			})
		}
		p.Filters = tmpFilters
	} else {
		p.Filters = filters
	}

	return p
}

// 查询表单中的权重，权重大排序靠前
func (p *Column) SetOrder(order int) *Column {
	p.Order = order
	return p
}

// 可排序列
func (p *Column) SetSorter(sorter interface{}) *Column {
	p.Sorter = sorter
	return p
}

// 包含列的数量
func (p *Column) SetSpan(span int) *Column {
	p.Span = span
	return p
}

// 设置列宽
func (p *Column) SetWidth(width int) *Column {
	p.Width = width
	return p
}

// 设置为可编辑列
func (p *Column) SetEditable(name string, options interface{}, action string) *Column {
	var getOptions interface{}

	if name == "select" {
		for k, v := range options.([]interface{}) {
			item := map[string]interface{}{
				"label": v,
				"value": k,
			}
			getOptions = append(getOptions.([]map[string]interface{}), item)
		}
	} else {
		getOptions = options
	}

	p.Editable = map[string]interface{}{
		"name":    name,
		"options": getOptions,
		"action":  action,
	}

	return p
}

// 设置为数据操作列
func (p *Column) SetActions(actions interface{}) *Column {
	p.Actions = actions

	return p
}

// 传递给 Form.Item 的配置,可以配置 rules，但是默认的查询表单 rules 是不生效的。需要配置 ignoreRules
func (p *Column) SetFormItemProps(formItemProps interface{}) *Column {
	p.FormItemProps = formItemProps
	return p
}

// 传给渲染的组件的 props，自定义的时候也会传递
func (p *Column) SetFieldProps(fieldProps interface{}) *Column {
	p.FieldProps = fieldProps
	return p
}

// 组件json序列化
func (p *Column) JsonSerialize() *Column {
	p.Component = "column"

	return p
}
