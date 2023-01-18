package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Number struct {
	Item
	Min       int `json:"min"`
	Max       int `json:"max"`
	Step      int `json:"step"`
	Precision int `json:"precision"`
}

// 初始化
func (p *Number) Init() *Number {
	p.Component = "inputNumberField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Min = -10000
	p.Max = 10000
	p.Step = 1
	p.Precision = 0
	p.SetWidth(200)

	return p
}

// 最小值
func (p *Number) SetMin(min int) *Number {
	p.Min = min

	return p
}

// 最大值
func (p *Number) SetMax(max int) *Number {
	p.Max = max

	return p
}

// 每次改变步数，可以为小数
func (p *Number) SetStep(step int) *Number {
	p.Step = step

	return p
}

// 数值精度
func (p *Number) SetPrecision(precision int) *Number {
	p.Precision = precision

	return p
}
