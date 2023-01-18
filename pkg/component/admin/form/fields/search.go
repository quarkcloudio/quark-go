package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Search struct {
	Item
}

// 初始化
func (p *Search) Init() *Search {
	p.Component = "searchField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.SetWidth(200)
	p.Placeholder = "请输入要搜索的内容"
	p.AllowClear = true

	return p
}

// api
func (p *Search) SetApi(api string) *Search {
	p.Api = api

	return p
}

// 设置单选属性，[]map[string]interface{}{{"label": "Title1","value": "value1"},{"label": "Title2","value": "value2"}}
// 或者 map[interface{}]interface{}{"value1":"Title1","value2":"Title2"}
func (p *Search) SetOptions(options interface{}) *Search {
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

// 设置 Search 的模式为多选或标签，multiple | tags
func (p *Search) SetMode(mode string) *Search {
	p.Mode = mode

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Search) SetSize(size string) *Search {
	p.Size = size

	return p
}

// 可以点击清除图标删除内容
func (p *Search) SetAllowClear(allowClear bool) *Search {
	p.AllowClear = allowClear

	return p
}
