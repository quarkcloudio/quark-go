package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Picker struct {
	Item
	Value        interface{}   `json:"value"`
	Start        string        `json:"start"`
	End          string        `json:"end"`
	Fields       string        `json:"fields"`
	CustomItem   string        `json:"customItem"`
	Mode         string        `json:"mode"`
	Range        []interface{} `json:"range"`
	RangeKey     string        `json:"RangeKey"`
	SelectorType string        `json:"SelectorType"`
	Disabled     bool          `json:"disabled"`
}

// 初始化
func (p *Picker) Init() *Picker {
	p.Component = "pickerField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.Mode = "selector"
	p.Value = 0
	p.SelectorType = "auto"

	return p
}

// value 的值表示选择了 range 中的第几个（下标从 0 开始）
func (p *Picker) SetValue(value int) *Picker {
	p.Value = value

	return p
}

// 选择器模式：selector,multiSelector,time,date,region
func (p *Picker) SetMode(mode string) *Picker {
	p.Mode = mode

	return p
}

// mode为 selector 或 multiSelector 时，range 有效
func (p *Picker) SetRange(pickerRange []interface{}) *Picker {
	p.Range = pickerRange

	return p
}

// 表示有效时间范围的开始
func (p *Picker) SetStart(start string) *Picker {
	p.Start = start

	return p
}

// 表示有效时间范围的结束
func (p *Picker) SetEnd(end string) *Picker {
	p.End = end

	return p
}

// 有效值 year、month、day，表示选择器的粒度，默认为 day，App 端未配置此项时使用系统 UI
func (p *Picker) SetFields(fields string) *Picker {
	p.Fields = fields

	return p
}

// 有效值 year、month、day，表示选择器的粒度，默认为 day，App 端未配置此项时使用系统 UI
func (p *Picker) SetCustomItem(customItem string) *Picker {
	p.CustomItem = customItem

	return p
}

// 当 range 是一个 Array＜Object＞ 时，通过 range-key 来指定 Object 中 key 的值作为选择器显示内容
func (p *Picker) SetRangeKey(rangeKey string) *Picker {
	p.RangeKey = rangeKey

	return p
}

// 大屏时UI类型，支持 picker、select、auto，默认在 iPad 以 picker 样式展示而在 PC 以 select 样式展示
func (p *Picker) SetSelectorType(selectorType string) *Picker {
	p.SelectorType = selectorType

	return p
}

// 是否禁用
func (p *Picker) SetDisabled(disabled bool) *Picker {
	p.Disabled = disabled

	return p
}

// 组件json序列化
func (p *Picker) JsonSerialize() *Picker {
	p.Component = "pickerField"

	return p
}
