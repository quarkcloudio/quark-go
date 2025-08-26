package order

import "github.com/quarkcloudio/quark-go/v4/component/component"

type Info struct {
	component.Element
	Title      string      `json:"title"`
	Column     int         `json:"column"`
	Layout     string      `json:"layout"`
	Colon      bool        `json:"colon"`
	DataIndex  string      `json:"dataIndex"`
	Items      []Item      `json:"items"`
	DataSource interface{} `json:"dataSource"`
}

// 初始化组件
func NewInfo() *Info {
	return (&Info{}).Init()
}

// 初始化
func (p *Info) Init() *Info {
	p.Component = "info"
	p.Column = 5
	p.Colon = true
	p.Layout = "horizontal"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	return p
}

// 设置标题
func (p *Info) SetTitle(title string) *Info {
	p.Title = title
	return p
}

// 设置列数量
func (p *Info) SetColumn(column int) *Info {
	p.Column = column
	return p
}

// 设置布局 vertical 或者 horizontal
func (p *Info) SetLayout(layout string) *Info {
	p.Layout = layout
	return p
}

// 是否有冒号分割
func (p *Info) SetColon(colon bool) *Info {
	p.Colon = colon
	return p
}

// 设置数据索引
func (p *Info) SetDataIndex(dataIndex string) *Info {
	p.DataIndex = dataIndex
	return p
}

// 设置展示项
func (p *Info) SetItems(items []Item) *Info {
	p.Items = items
	return p
}

// 设置数据源
func (p *Info) SetDataSource(dataSource interface{}) *Info {
	p.DataSource = dataSource
	return p
}
