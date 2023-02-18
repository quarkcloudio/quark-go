package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type DatetimeRange struct {
	Item
}

// 初始化
func (p *DatetimeRange) Init() *DatetimeRange {
	p.Component = "datetimeRangeField"
	p.Format = "YYYY-MM-DD HH:mm:ss"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.DefaultValue = []interface{}{nil, nil}

	return p
}

// 使用 format 属性，可以自定义日期显示格式
func (p *DatetimeRange) SetFormat(format string) *DatetimeRange {
	p.Format = format

	return p
}
