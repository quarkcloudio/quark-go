package order

import "github.com/quarkcloudio/quark-go/v4/component/component"

type Component struct {
	component.Element
	InitApi         string  `json:"initApi"`
	Icon            string  `json:"icon"`
	OrderNoText     string  `json:"orderNoText"`
	OrderNoName     string  `json:"orderNoName"`
	OrderDetailText string  `json:"orderDetailText"`
	OrderItemText   string  `json:"orderItemText"`
	OrderStatusText string  `json:"orderStatusText"`
	Info            *Info   `json:"info"`
	DetailInfo      []*Info `json:"detailInfo"`
	ItemInfo        *Table  `json:"itemInfo"`
	StatusInfo      *Table  `json:"statusInfo"`
}

type Column struct {
	Title     string `json:"title"`
	DataIndex string `json:"dataIndex"`
	ValueType string `json:"valueType"`
}

type Item struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// NewColumn("name", "标题") 或者 NewColumn("name", "标题", "image")
func NewColumn(params ...string) Column {
	column := Column{}
	if len(params) == 2 {
		column.DataIndex = params[0]
		column.Title = params[1]
		column.ValueType = "text"
	}
	if len(params) == 3 {
		column.DataIndex = params[0]
		column.Title = params[1]
		column.ValueType = params[2]
	}
	return column
}

func NewItem(key string, label string) Item {
	return Item{key, label}
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "order"
	p.OrderNoName = "order_no"
	p.OrderNoText = "订单号"
	p.OrderDetailText = "订单信息"
	p.OrderItemText = "商品信息"
	p.OrderStatusText = "订单记录"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style
	return p
}

// 初始化数据接口
func (p *Component) SetInitApi(initApi string) *Component {
	p.InitApi = initApi
	return p
}

// 设置图标地址
func (p *Component) SetIcon(icon string) *Component {
	p.Icon = icon
	return p
}

// 设置订单号展示文本
func (p *Component) SetOrderNoText(orderNoText string) *Component {
	p.OrderNoText = orderNoText
	return p
}

// 设置订单信息展示文本
func (p *Component) SetOrderDetailText(orderDetailText string) *Component {
	p.OrderDetailText = orderDetailText
	return p
}

// 设置商品信息展示文本
func (p *Component) SetOrderItemText(orderItemText string) *Component {
	p.OrderItemText = orderItemText
	return p
}

// 设置订单记录展示文本
func (p *Component) SetOrderStatusText(orderStatusText string) *Component {
	p.OrderStatusText = orderStatusText
	return p
}

// 设置订单号name
func (p *Component) SetOrderNoName(orderNoName string) *Component {
	p.OrderNoName = orderNoName
	return p
}

// 设置订单基本信息
func (p *Component) SetInfo(info *Info) *Component {
	p.Info = info
	return p
}

// 设置订单信息
func (p *Component) SetDetailInfo(detailInfo []*Info) *Component {
	p.DetailInfo = detailInfo
	return p
}

// 设置商品信息
func (p *Component) SetItemInfo(itemInfo *Table) *Component {
	p.ItemInfo = itemInfo
	return p
}

// 设置订单记录
func (p *Component) SetStatusInfo(statusInfo *Table) *Component {
	p.StatusInfo = statusInfo
	return p
}
