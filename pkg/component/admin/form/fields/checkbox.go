package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Checkbox struct {
	Item
	Layout string
}

// 初始化
func (p *Checkbox) Init() *Checkbox {
	p.Component = "checkboxField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置单选属性，[]map[string]interface{}{{"label": "Title1","value": "value1"},{"label": "Title2","value": "value2"}}
// 或者 map[interface{}]interface{}{"value1":"Title1","value2":"Title2"}
func (p *Checkbox) SetOptions(options interface{}) *Checkbox {
	var data []map[string]interface{}

	if mapOptions, ok := options.(map[interface{}]interface{}); ok {
		for k, v := range mapOptions {
			option := map[string]interface{}{
				"label": v,
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

// 配置 checkbox 的样子，支持垂直vertical 和 horizontal
func (p *Checkbox) SetLayout(layout string) *Checkbox {
	p.Layout = layout

	return p
}

// 组件json序列化
func (p *Checkbox) JsonSerialize() *Checkbox {
	p.Component = "checkboxField"

	return p
}
