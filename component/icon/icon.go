package icon

import (
	"github.com/quarkcloudio/quark-go/v4/component/component"
	"github.com/quarkcloudio/quark-go/v4/utils/hex"
)

type Component struct {
	component.Element
	ClassName string      `json:"className"` // 计算后的 svg 类名
	Height    interface{} `json:"height"`    // svg 元素高度
	Width     interface{} `json:"width"`     // svg 元素宽度
	Type      string      `json:"type"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "icon"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *Component) SetKey(key string, crypt bool) *Component {
	p.ComponentKey = hex.Make(key, crypt)

	return p
}

// 计算后的 svg 类名
func (p *Component) SetClassName(className string) *Component {
	p.ClassName = className
	return p
}

// svg 元素高度
func (p *Component) SetHeight(height interface{}) *Component {
	p.Height = height
	return p
}

// svg 元素宽度
func (p *Component) SetWidth(width interface{}) *Component {
	p.Width = width
	return p
}

// 可以在 https://icon-sets.iconify.design/ 查看图标，图标类型：{集合名}:{图标名}，例如：ant-design:user-outlined
func (p *Component) SetType(icontType string) *Component {
	p.Type = icontType
	return p
}
