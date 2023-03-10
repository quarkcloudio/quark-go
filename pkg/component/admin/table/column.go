package table

import (
	"github.com/quarkcms/quark-go/pkg/component/admin/component"
)

type Column struct {
	component.Element
	ComponentKey  string      `json:"componentKey"`
	Title         string      `json:"title"`
	Attribute     string      `json:"attribute"`
	Align         string      `json:"align"`
	DataIndex     string      `json:"dataIndex"`
	Fixed         interface{} `json:"fixed"`
	Tooltip       string      `json:"tooltip"`
	Ellipsis      bool        `json:"ellipsis"`
	Copyable      bool        `json:"copyable"`
	ValueEnum     interface{} `json:"valueEnum"`
	ValueType     string      `json:"valueType"`
	HideInSearch  bool        `json:"hideInSearch"`
	HideInTable   bool        `json:"hideInTable"`
	HideInForm    bool        `json:"hideInForm"`
	Filters       interface{} `json:"filters"`
	Order         int         `json:"order"`
	Sorter        interface{} `json:"sorter"`
	Width         int         `json:"width"`
	Editable      interface{} `json:"editable"`
	Actions       interface{} `json:"actions"`
	FormItemProps interface{} `json:"formItemProps"`
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

/**
 * 列头显示文字，既字段的列名
 *
 * @param  string  title
 * @return p
 */
func (p *Column) SetTitle(title string) *Column {
	p.Title = title
	return p
}

/**
 * 字段名称|字段的列名
 *
 * @param  string  attribute
 * @return p
 */
func (p *Column) SetAttribute(attribute string) *Column {
	p.SetKey(attribute, false)
	p.ComponentKey = attribute
	p.DataIndex = attribute
	p.Attribute = attribute

	return p
}

/**
 * 设置列的对齐方式,left | right | center
 *
 * @param  string  align
 * @return p
 */
func (p *Column) SetAlign(align string) *Column {
	p.Align = align
	return p
}

/**
 * （IE 下无效）列是否固定，可选 true (等效于 left) left right
 *
 * @param  bool|string  fixed
 * @return p
 */
func (p *Column) SetFixed(fixed interface{}) *Column {
	p.Fixed = fixed
	return p
}

/**
 * 会在 title 之后展示一个 icon，hover 之后提示一些信息
 *
 * @param  bool|string  tooltip
 * @return p
 */
func (p *Column) SetTooltip(tooltip string) *Column {
	p.Tooltip = tooltip
	return p
}

/**
 * 是否自动缩略
 *
 * @param  bool  ellipsis
 * @return p
 */
func (p *Column) SetEllipsis(ellipsis bool) *Column {
	p.Ellipsis = ellipsis
	return p
}

/**
 * 是否支持复制
 *
 * @param  bool  copyable
 * @return p
 */
func (p *Column) SetCopyable(copyable bool) *Column {
	p.Copyable = copyable
	return p
}

/**
 * 值的枚举，会自动转化把值当成 key 来取出要显示的内容
 *
 * @param  array  valueEnum
 * @return p
 */
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

/**
 * 值的类型,"money" | "option" | "date" | "dateTime" | "time" | "text"| "index" | "indexBorder"
 *
 * @param  string  valueType
 * @return p
 */
func (p *Column) SetValueType(valueType string) *Column {
	p.ValueType = valueType
	return p
}

/**
 * 在查询表单中不展示此项
 *
 * @param  bool  hideInSearch
 * @return p
 */
func (p *Column) SetHideInSearch(hideInSearch bool) *Column {
	p.HideInSearch = hideInSearch
	return p
}

/**
 * 在 Table 中不展示此列
 *
 * @param  bool  hideInTable
 * @return p
 */
func (p *Column) SetHideInTable(hideInTable bool) *Column {
	p.HideInTable = hideInTable
	return p
}

/**
 * 表头的筛选菜单项，当值为 true 时，自动使用 valueEnum 生成
 *
 * @param  array|bool  filters
 * @return p
 */
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

/**
 * 查询表单中的权重，权重大排序靠前
 *
 * @param  number  order
 * @return p
 */
func (p *Column) SetOrder(order int) *Column {
	p.Order = order
	return p
}

/**
 * 可排序列
 *
 * @param  bool  sorter
 * @return p
 */
func (p *Column) SetSorter(sorter bool) *Column {
	p.Sorter = sorter
	return p
}

/**
 * 设置列宽
 *
 * @param  string|number  width
 * @return p
 */
func (p *Column) SetWidth(width int) *Column {
	p.Width = width
	return p
}

/**
 * 设置为可编辑列
 *
 * @param  string  name
 * @param  array|bool  options
 * @param  string  action
 * @return p
 */
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

/**
 * 设置为数据操作列
 *
 * @param  array  actions
 * @return p
 */
func (p *Column) SetActions(actions interface{}) *Column {
	p.Actions = actions

	return p
}

/**
 * 传递给 Form.Item 的配置,可以配置 rules，但是默认的查询表单 rules 是不生效的。需要配置 ignoreRules
 *
 * @param  formItemProps
 * @return p
 */
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
