package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type DataSelect struct {
	Item
	Value       interface{} `json:"value"`
	Localdata   interface{} `json:"localdata"`
	Clear       bool        `json:"clear"`
	Label       string      `json:"label"`
	Placeholder string      `json:"placeholder"`
	EmptyText   string      `json:"emptyText"`
}

// 初始化
func (p *DataSelect) Init() *DataSelect {
	p.Component = "dataSelectField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.Placeholder = "请选择"
	p.EmptyText = "暂无数据	"

	return p
}

// 已选择数据的 value 值
func (p *DataSelect) SetValue(value interface{}) *DataSelect {
	p.Value = value

	return p
}

// 本地渲染数据
func (p *DataSelect) SetLocaldata(localdata interface{}) *DataSelect {
	p.Localdata = localdata

	return p
}

// 是否可以清空已选项
func (p *DataSelect) SetClear(clear bool) *DataSelect {
	p.Clear = clear

	return p
}

// 左侧标题
func (p *DataSelect) SetLabel(label string) *DataSelect {
	p.Label = label

	return p
}

// 输入框的提示文字
func (p *DataSelect) SetPlaceholder(placeholder string) *DataSelect {
	p.Placeholder = placeholder

	return p
}

// 没有数据时显示的文字 ，本地数据无效
func (p *DataSelect) SetEmptyText(emptyText string) *DataSelect {
	p.EmptyText = emptyText

	return p
}

// 组件json序列化
func (p *DataSelect) JsonSerialize() *DataSelect {
	p.Component = "dataSelectField"

	return p
}
