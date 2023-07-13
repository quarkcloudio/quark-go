package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Radio struct {
	Item
	Value    interface{} `json:"value"`
	Disabled bool        `json:"disabled"`
	Checked  bool        `json:"checked"`
	Color    string      `json:"color"`
	Options  interface{} `json:"options"`
}

// 初始化
func (p *Radio) Init() *Radio {
	p.Component = "radioField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.Color = "rgb(0, 122, 255)"

	return p
}

// 默认值
func (p *Radio) SetValue(value interface{}) *Radio {
	p.Value = value

	return p
}

// 本地渲染数据
func (p *Radio) SetDisabled(disabled bool) *Radio {
	p.Disabled = disabled

	return p
}

// 本地渲染数据
func (p *Radio) SetChecked(checked bool) *Radio {
	p.Checked = checked

	return p
}

// list 列表模式下 icon 显示的位置
func (p *Radio) SetColor(color string) *Radio {
	p.Color = color

	return p
}

// 设置单选属性，[]map[string]interface{}{{"text": "Title1","value": "value1"},{"text": "Title2","value": "value2"}}
// 或者 map[interface{}]interface{}{"value1":"Title1","value2":"Title2"}
func (p *Radio) SetOptions(options interface{}) *Radio {
	var data []map[string]interface{}

	if mapOptions, ok := options.(map[interface{}]interface{}); ok {
		for k, v := range mapOptions {
			option := map[string]interface{}{
				"text":  v,
				"value": k,
			}

			data = append(data, option)
		}
	} else if sliceOptions, ok := options.([]map[string]interface{}); ok {
		data = sliceOptions
	}

	p.Options = data

	return p
}

// 组件json序列化
func (p *Radio) JsonSerialize() *Radio {
	p.Component = "radioField"

	return p
}
