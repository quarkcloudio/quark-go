package grid

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct{}

type Col struct {
	component.Element
	Flex   string      `json:"flex"`
	Offset int         `json:"offset"`
	Order  int         `json:"order"`
	Pull   int         `json:"pull"`
	Push   int         `json:"push"`
	Span   int         `json:"span"`
	Xs     interface{} `json:"xs"`
	Sm     interface{} `json:"sm"`
	Md     interface{} `json:"md"`
	Lg     interface{} `json:"lg"`
	Xl     interface{} `json:"xl"`
	Xxl    interface{} `json:"xxl"`
	Body   interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{})
}

// 获取Col
func (p *Component) Col() *Col {
	return (&Col{}).Init()
}

// 初始化
func (p *Col) Init() *Col {
	p.Component = "col"
	p.Offset = 0
	p.Order = 0
	p.Pull = 0
	p.Push = 0

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Col) SetStyle(style map[string]interface{}) *Col {
	p.Style = style

	return p
}

// 布局属性
func (p *Col) SetFlex(flex string) *Col {
	p.Flex = flex

	return p
}

// 栅格左侧的间隔格数，间隔内不可以有栅格
func (p *Col) SetOffset(offset int) *Col {
	p.Offset = offset

	return p
}

// 栅格顺序
func (p *Col) SetOrder(order int) *Col {
	p.Order = order

	return p
}

// 栅格向左移动格数
func (p *Col) SetPull(pull int) *Col {
	p.Pull = pull

	return p
}

// 栅格向右移动格数
func (p *Col) SetPush(push int) *Col {
	p.Push = push

	return p
}

// 栅格占位格数，为 0 时相当于 display: none
func (p *Col) SetSpan(span int) *Col {
	p.Span = span

	return p
}

// 屏幕 < 576px 响应式栅格，可为栅格数或一个包含其他属性的对象
func (p *Col) SetXs(xs interface{}) *Col {
	p.Xs = xs

	return p
}

// 屏幕 ≥ 576px 响应式栅格，可为栅格数或一个包含其他属性的对象
func (p *Col) SetSm(sm interface{}) *Col {
	p.Sm = sm

	return p
}

// 屏幕 ≥ 768px 响应式栅格，可为栅格数或一个包含其他属性的对象
func (p *Col) SetMd(md interface{}) *Col {
	p.Md = md

	return p
}

// 屏幕 ≥ 992px 响应式栅格，可为栅格数或一个包含其他属性的对象
func (p *Col) SetLg(lg interface{}) *Col {
	p.Lg = lg

	return p
}

// 屏幕 ≥ 1200px 响应式栅格，可为栅格数或一个包含其他属性的对象
func (p *Col) SetXl(xl interface{}) *Col {
	p.Xl = xl

	return p
}

// 屏幕 ≥ 1600px 响应式栅格，可为栅格数或一个包含其他属性的对象
func (p *Col) SetXxl(xxl interface{}) *Col {
	p.Xxl = xxl

	return p
}

// 卡牌内容
func (p *Col) SetBody(body interface{}) *Col {
	p.Body = body

	return p
}

// 组件json序列化
func (p *Col) JsonSerialize() *Col {
	p.Component = "col"

	return p
}
