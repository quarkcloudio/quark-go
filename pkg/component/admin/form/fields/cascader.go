package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Cascader struct {
	Item
}

// 初始化
func (p *Cascader) Init() *Cascader {
	p.Component = "cascaderField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Placeholder = "请选择"

	return p
}

// 设置属性
func (p *Cascader) SetOptions(options interface{}) *Cascader {
	p.Options = options

	return p
}

// 控件大小。注：标准表单内的输入框大小限制为 large。可选 large default small
func (p *Cascader) SetSize(size string) *Cascader {
	p.Size = size

	return p
}

// 可以点击清除图标删除内容
func (p *Cascader) SetAllowClear(allowClear bool) *Cascader {
	p.AllowClear = allowClear

	return p
}

// 获取数据接口
func (p *Cascader) SetApi(api string) *Cascader {
	p.Api = api
	return p
}
