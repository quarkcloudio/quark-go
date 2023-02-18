package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Time struct {
	Item
}

// 初始化
func (p *Time) Init() *Time {
	p.Component = "timeField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Format = "HH:mm"

	return p
}

// 使用 format 属性，可以自定义日期显示格式
func (p *Time) SetFormat(format string) *Time {
	p.Format = format

	return p
}
