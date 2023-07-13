package dropdown

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/drawer"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/modal"
)

type Item struct {
	component.Element
	Label        string      `json:"label"`
	Block        bool        `json:"block"`
	Danger       bool        `json:"danger"`
	Disabled     bool        `json:"disabled"`
	Ghost        bool        `json:"ghost"`
	Icon         string      `json:"icon"`
	Shape        string      `json:"shape"`
	Size         string      `json:"size"`
	Type         string      `json:"type"`
	ActionType   string      `json:"actionType"`
	SubmitForm   any         `json:"submitForm"`
	Href         string      `json:"href"`
	Target       string      `json:"target"`
	Modal        interface{} `json:"modal"`
	Drawer       interface{} `json:"drawer"`
	ConfirmTitle string      `json:"confirmTitle"`
	ConfirmText  string      `json:"confirmText"`
	ConfirmType  string      `json:"confirmType"`
	Api          string      `json:"api"`
	Reload       string      `json:"reload"`
	WithLoading  bool        `json:"withLoading"`
}

// 初始化
func (p *Item) Init() *Item {
	p.Component = "itemStyle"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Item) SetStyle(style map[string]interface{}) *Item {
	p.Style = style

	return p
}

// 设置按钮文字
func (p *Item) SetLabel(label string) *Item {
	p.Label = label

	return p
}

// 将按钮宽度调整为其父宽度的选项
func (p *Item) SetBlock(block bool) *Item {
	p.Block = block

	return p
}

// 设置危险按钮
func (p *Item) SetDanger(danger bool) *Item {
	p.Danger = danger

	return p
}

// 按钮失效状态
func (p *Item) SetDisabled(disabled bool) *Item {
	p.Disabled = disabled

	return p
}

// 幽灵属性，使按钮背景透明
func (p *Item) SetGhost(ghost bool) *Item {
	p.Ghost = ghost

	return p
}

// 设置按钮图标
func (p *Item) SetIcon(icon string) *Item {
	p.Icon = "icon-" + icon

	return p
}

// 设置按钮形状，可选值为 circle、 round 或者不设
func (p *Item) SetShape(shape string) *Item {
	p.Shape = shape

	return p
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Item) SetType(buttonType string, danger bool) *Item {
	p.Type = buttonType
	p.Danger = danger

	return p
}

// 设置按钮大小，large | middle | small | default
func (p *Item) SetSize(size string) *Item {
	p.Size = size

	return p
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Item) SetActionType(actionType string) *Item {
	p.ActionType = actionType

	return p
}

// 当action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Item) SetSubmitForm(formKey string) *Item {
	p.SubmitForm = formKey

	return p
}

// 点击跳转的地址，指定此属性 button 的行为和 a 链接一致
func (p *Item) SetHref(href string) *Item {
	p.Href = href

	return p
}

// 相当于 a 链接的 target 属性，href 存在时生效
func (p *Item) SetTarget(target string) *Item {
	p.Target = target

	return p
}

// 设置跳转链接
func (p *Item) SetLink(href string, target string) *Item {
	p.SetHref(href)
	p.SetTarget(target)
	p.ActionType = "link"

	return p
}

// 弹窗
func (p *Item) SetModal(callback interface{}) *Item {
	component := (&modal.Component{}).Init()
	getCallback := callback.(func(modal *modal.Component) interface{})

	p.Modal = getCallback(component)

	return p
}

// 抽屉
func (p *Item) SetDrawer(callback interface{}) *Item {
	component := (&drawer.Component{}).Init()
	getCallback := callback.(func(drawer *drawer.Component) interface{})

	p.Drawer = getCallback(component)

	return p
}

// 设置行为前的确认操作
func (p *Item) SetWithConfirm(title string, text string, confirmType string) *Item {
	p.ConfirmTitle = title
	p.ConfirmText = text
	p.ConfirmType = confirmType

	return p
}

// 执行行为的接口链接
func (p *Item) SetApi(api string) *Item {
	p.Api = api
	p.ActionType = "ajax"

	return p
}

// 执行成功后刷新的组件
func (p *Item) SetReload(reload string) *Item {
	p.Reload = reload

	return p
}

// 是否具有loading
func (p *Item) SetWithLoading(loading bool) *Item {
	p.WithLoading = loading

	return p
}
