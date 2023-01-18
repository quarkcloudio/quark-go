package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Editor struct {
	Item
}

// 初始化
func (p *Editor) Init() *Editor {
	p.Component = "editorField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Style = map[string]interface{}{
		"height": 500,
		"width":  "100%",
	}

	return p
}

// 高度
func (p *Editor) SetHeight(height interface{}) *Editor {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["height"] = height
	p.Style = style

	return p
}
