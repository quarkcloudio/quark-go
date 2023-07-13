package divider

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	Dashed      bool        `json:"dashed"`
	Orientation string      `json:"orientation"`
	Plain       bool        `json:"plain"`
	Type        string      `json:"type"`
	Body        interface{} `json:"body"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "divider"
	p.Type = "horizontal"
	p.Orientation = "center"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 是否虚线
func (p *Component) SetDashed(dashed bool) *Component {
	p.Dashed = dashed

	return p
}

// 间距方向,'left', 'right', 'center'
func (p *Component) SetOrientation(orientation string) *Component {

	limits := []string{
		"left", "right", "center",
	}

	inSlice := false
	for _, limit := range limits {
		if limit == orientation {
			inSlice = true
		}
	}

	if inSlice == false {
		panic("Argument must be in 'left', 'right', 'center'!")
	}

	p.Orientation = orientation

	return p
}

// 文字是否显示为普通正文样式
func (p *Component) SetPlain(plain bool) *Component {
	p.Plain = plain

	return p
}

// 水平还是垂直类型,horizontal | vertical
func (p *Component) SetType(dividerType string) *Component {

	limits := []string{
		"vertical", "horizontal",
	}

	inSlice := false
	for _, limit := range limits {
		if limit == dividerType {
			inSlice = true
		}
	}

	if inSlice == false {
		panic("Argument must be in 'vertical', 'horizontal'!")
	}

	p.Type = dividerType

	return p
}

// 容器控件里面的内容
func (p *Component) SetBody(body interface{}) *Component {
	p.Body = body

	return p
}
