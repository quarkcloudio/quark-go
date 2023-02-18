package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type DateRange struct {
	Item
	Picker string `json:"picker"`
}

// 初始化
func (p *DateRange) Init() *DateRange {
	p.Component = "dateRangeField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Picker = "date"
	p.Format = "YYYY-MM-DD"
	p.DefaultValue = []interface{}{nil, nil}

	return p
}

// 设置选择器类型,date | week | month | quarter | year
func (p *DateRange) SetPicker(picker string) *DateRange {

	limits := []string{
		"date", "week", "month", "quarter", "year",
	}

	inSlice := false
	for _, limit := range limits {
		if limit == picker {
			inSlice = true
		}
	}

	if inSlice == false {
		panic("argument must be in 'date', 'week', 'month', 'quarter', 'year'!")
	}

	switch picker {
	case "date":
		p.SetFormat("YYYY-MM-DD")
	case "week":
		p.SetFormat("MM-DD")
	case "month":
		p.SetFormat("YYYY-MM")
	case "quarter":
		p.SetFormat("YYYY-MM")
	case "year":
		p.SetFormat("YYYY")
	}

	p.Picker = picker

	return p
}

// 使用 format 属性，可以自定义日期显示格式
func (p *DateRange) SetFormat(format string) *DateRange {
	p.Format = format

	return p
}
