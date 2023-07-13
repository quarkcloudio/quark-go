package descriptions

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/descriptions/fields"
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
