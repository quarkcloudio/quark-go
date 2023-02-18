package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Datetime struct {
	Item
}

// 初始化
func (p *Datetime) Init() *Datetime {
	p.Component = "datetimeField"
	p.Format = "YYYY-MM-DD HH:mm:ss"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 使用 format 属性，可以自定义日期显示格式
func (p *Datetime) SetFormat(format string) *Datetime {
	p.Format = format

	return p
}
