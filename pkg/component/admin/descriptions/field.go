package descriptions

import (
	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/descriptions/fields"
)

type Field struct {
	component.Element
}

// 初始化
func (p *Field) Init() *Field {
	p.Component = "descriptionField"

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// text组件
func (p *Field) Text(params ...string) *fields.Text {

	fields := &fields.Text{}

	if len(params) == 1 {
		fields = fields.Init().SetDataIndex(params[0]).SetLabel(params[0])
	} else {
		fields = fields.Init().SetDataIndex(params[0]).SetLabel(params[1])
	}

	return fields
}
