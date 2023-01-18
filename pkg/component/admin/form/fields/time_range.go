package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type TimeRange struct {
	Item
	Format string `json:"format"`
}

// 初始化
func (p *TimeRange) Init() *TimeRange {
	p.Component = "timeRangeField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Format = "HH:mm"
	p.DefaultValue = []interface{}{nil, nil}

	return p
}

// 使用 format 属性，可以自定义日期显示格式
func (p *TimeRange) SetFormat(format string) *TimeRange {
	p.Format = format

	return p
}
