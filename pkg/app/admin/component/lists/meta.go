package lists

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Meta struct {
	component.Element
	Title     string      `json:"title"`
	Attribute string      `json:"attribute"`
	DataIndex string      `json:"dataIndex"`
	Ellipsis  bool        `json:"ellipsis"`
	Copyable  bool        `json:"copyable"`
	ValueEnum interface{} `json:"valueEnum"`
	ValueType string      `json:"valueType"`
	Search    bool        `json:"search"`
	Actions   bool        `json:"actions"`
}

// 初始化
func (p *Meta) Init() *Meta {
	p.Component = "meta"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Meta) SetStyle(style map[string]interface{}) *Meta {
	p.Style = style

	return p
}

/**
 * 列头显示文字，既字段的列名
 *
 * @param  string  title
 * @return p
 */
func (p *Meta) SetTitle(title string) *Meta {
	p.Title = title
	return p
}

/**
 * 字段名称|字段的列名
 *
 * @param  string  attribute
 * @return p
 */
func (p *Meta) SetAttribute(attribute string) *Meta {
	p.DataIndex = attribute
	p.Attribute = attribute

	return p
}

/**
 * 是否自动缩略
 *
 * @param  bool  ellipsis
 * @return p
 */
func (p *Meta) SetEllipsis(ellipsis bool) *Meta {
	p.Ellipsis = ellipsis
	return p
}

/**
 * 是否支持复制
 *
 * @param  bool  copyable
 * @return p
 */
func (p *Meta) SetCopyable(copyable bool) *Meta {
	p.Copyable = copyable
	return p
}

/**
 * 值的枚举，会自动转化把值当成 key 来取出要显示的内容
 *
 * @param  array  valueEnum
 * @return p
 */
func (p *Meta) SetValueEnum(valueEnum interface{}) *Meta {
	valueEnumStr := map[string]interface{}{}
	valueEnumInt := map[int]interface{}{}

	for k, v := range valueEnum.(map[interface{}]interface{}) {
		if value, ok := k.(string); ok == true {
			valueEnumStr[value] = v
		} else if value, ok := k.(int); ok == true {
			valueEnumInt[value] = v
		}
	}

	if len(valueEnumStr) > 0 {
		p.ValueEnum = valueEnumStr
	}

	if len(valueEnumInt) > 0 {
		p.ValueEnum = valueEnumInt
	}

	return p
}

/**
 * 值的类型,"money" | "option" | "date" | "dateTime" | "time" | "text"| "index" | "indexBorder"
 *
 * @param  string  valueType
 * @return p
 */
func (p *Meta) SetValueType(valueType string) *Meta {
	p.ValueType = valueType
	return p
}

// 在查询表单中不展示此项
func (p *Meta) SetSearch(search bool) *Meta {
	p.Search = search

	return p
}
