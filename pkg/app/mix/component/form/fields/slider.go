package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Slider struct {
	Item
	Value           int    `json:"value"`
	Min             int    `json:"min"`
	Max             int    `json:"max"`
	Step            int    `json:"step"`
	Disabled        bool   `json:"disabled"`
	ActiveColor     string `json:"activeColor"`
	BackgroundColor string `json:"backgroundColor"`
	BlockSize       int    `json:"blockSize"`
	BlockColor      string `json:"blockColor"`
	ShowValue       bool   `json:"showValue"`
}

// 初始化
func (p *Slider) Init() *Slider {
	p.Component = "sliderField"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.Min = 0
	p.Max = 100
	p.Step = 1
	p.Value = 0
	p.BackgroundColor = "#e9e9e9"
	p.BlockSize = 28
	p.BlockColor = "#ffffff"

	return p
}

// 默认值
func (p *Slider) SetValue(value int) *Slider {
	p.Value = value

	return p
}

// 最小值
func (p *Slider) SetMin(min int) *Slider {
	p.Min = min

	return p
}

// 最大值
func (p *Slider) SetMax(max int) *Slider {
	p.Max = max

	return p
}

// 步长，取值必须大于 0，并且可被(max - min)整除
func (p *Slider) SetStep(step int) *Slider {
	p.Step = step

	return p
}

// 是否禁用
func (p *Slider) SetDisabled(disabled bool) *Slider {
	p.Disabled = disabled

	return p
}

// 滑块左侧已选择部分的线条颜色
func (p *Slider) SetActiveColor(activeColor string) *Slider {
	p.ActiveColor = activeColor

	return p
}

// 滑块右侧背景条的颜色
func (p *Slider) SetBackgroundColor(backgroundColor string) *Slider {
	p.BackgroundColor = backgroundColor

	return p
}

// 滑块的大小，取值范围为 12 - 28
func (p *Slider) SetBlockSize(blockSize int) *Slider {
	p.BlockSize = blockSize

	return p
}

// 滑块的颜色
func (p *Slider) SetBlockColor(blockColor string) *Slider {
	p.BlockColor = blockColor

	return p
}

// 是否显示当前 value
func (p *Slider) SetShowValue(showValue bool) *Slider {
	p.ShowValue = showValue

	return p
}

// 组件json序列化
func (p *Slider) JsonSerialize() *Slider {
	p.Component = "sliderField"

	return p
}
