package when

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"
	"github.com/quarkcms/quark-go/v2/pkg/utils/hex"
)

type Item struct {
	Component         string      `json:"component"`          // 组件名称
	Condition         string      `json:"condition"`          // 条件：js表达式语句
	ConditionName     string      `json:"condition_name"`     // 需要对比的字段名称
	ConditionOperator string      `json:"condition_operator"` // 操作符，= <>
	Option            interface{} `json:"option"`             // 条件符合的属性值
	Body              interface{} `json:"body"`               // 内容
}

type Component struct {
	ComponentKey string  `json:"componentkey"` // 组件标识
	Component    string  `json:"component"`    // 组件名称
	Items        []*Item `json:"items"`        // When组件中需要解析的元素
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 获取Item
func NewItem() *Item {
	return &Item{}
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "when"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *Component) SetKey(key string, crypt bool) *Component {
	p.ComponentKey = hex.Make(key, crypt)

	return p
}

// 设置When组件中需要解析的元素
func (p *Component) SetItems(items []*Item) *Component {
	p.Items = items
	return p
}
