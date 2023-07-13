package descriptions

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Title      string      `json:"title"`
	Tooltip    string      `json:"tooltip"`
	Extra      interface{} `json:"extra"`
	Bordered   bool        `json:"bordered"`
	Column     interface{} `json:"column"`
	Size       string      `json:"size"`
	Layout     string      `json:"layout"`
	Colon      bool        `json:"colon"`
	Columns    interface{} `json:"columns"`
	DataSource interface{} `json:"dataSource"`
	Items      interface{} `json:"items"`
	Actions    interface{} `json:"actions"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 获取description组件Field
func NewField() *Field {
	return (&Field{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "descriptions"
	p.Column = 1
	p.Layout = "horizontal"
	p.Colon = true
	p.Size = "default"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 内容的补充描述，hover 后显示
func (p *Component) SetTooltip(tooltip string) *Component {
	p.Tooltip = tooltip
	return p
}

// 描述列表的操作区域，显示在右上方
func (p *Component) SetExtra(extra string) *Component {
	p.Extra = extra
	return p
}

// 是否展示边框
func (p *Component) SetBordered(bordered bool) *Component {
	p.Bordered = bordered
	return p
}

// 布局，horizontal|vertical
func (p *Component) SetLayout(layout string) *Component {
	p.Layout = layout
	return p
}

// 一行的 ProDescriptionsItems 数量，可以写成像素值或支持响应式的对象写法 { xs: 8, sm: 16, md: 24}
func (p *Component) SetColumn(column interface{}) *Component {
	p.Column = column
	return p
}

// 列
func (p *Component) SetColumns(columns interface{}) *Component {
	p.Columns = columns
	return p
}

// 数据
func (p *Component) SetDataSource(dataSource interface{}) *Component {
	p.DataSource = dataSource
	return p
}

// 设置尺寸
func (p *Component) SetSize(size string) *Component {
	p.Size = size
	return p
}

// 配置 ProDescriptions.Item 的 colon 的默认值
func (p *Component) SetColon(colon bool) *Component {
	p.Colon = colon
	return p
}

// 数据项
func (p *Component) SetItems(items interface{}) *Component {
	p.Items = items
	return p
}

// 行为
func (p *Component) SetActions(actions interface{}) *Component {
	p.Actions = actions
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "descriptions"

	return p
}
