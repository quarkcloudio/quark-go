package chart

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Line struct {
	component.Element
	Api           string      `json:"api"`
	Width         int         `json:"width"`
	Height        int         `json:"height"`
	AutoFit       bool        `json:"autoFit"`
	Padding       interface{} `json:"padding"`
	AppendPadding interface{} `json:"appendPadding"`
	Renderer      string      `json:"renderer"`
	LimitInPlot   bool        `json:"limitInPlot"`
	Locale        string      `json:"locale"`
	Data          interface{} `json:"data"`
	XField        string      `json:"xField"`
	YField        string      `json:"yField"`
	Meta          interface{} `json:"meta"`
	Smooth        bool        `json:"smooth"`
}

// 折线图表
func NewLine(data interface{}) *Line {
	return (&Line{}).Init().SetData(data)
}

// 初始化
func (p *Line) Init() *Line {
	p.Component = "line"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 数据接口
func (p *Line) SetApi(api string) *Line {
	p.Api = api
	return p
}

// 设置图表宽度
func (p *Line) SetWidth(width int) *Line {
	p.Width = width
	return p
}

// 设置图表高度
func (p *Line) SetHeight(height int) *Line {
	p.Height = height
	return p
}

// 图表是否自适应容器宽高。当 autoFit 设置为 true 时，width 和 height 的设置将失效。
func (p *Line) SetAutoFit(autoFit bool) *Line {
	p.AutoFit = autoFit
	return p
}

// 画布的 padding 值，代表图表在上右下左的间距，可以为单个数字 16，或者数组 [16, 8, 16, 8] 代表四个方向，或者开启 auto，由底层自动计算间距。
func (p *Line) SetPadding(padding interface{}) *Line {
	p.Padding = padding
	return p
}

// 额外增加的 appendPadding 值，在 padding 的基础上，设置额外的 padding 数值，可以是单个数字 16，或者数组 [16, 8, 16, 8] 代表四个方向。
func (p *Line) SetAppendPadding(appendPadding interface{}) *Line {
	p.AppendPadding = appendPadding
	return p
}

// 设置图表渲染方式为 canvas 或 svg。
func (p *Line) SetRenderer(renderer string) *Line {
	p.Renderer = renderer
	return p
}

// 是否对超出坐标系范围的 Geometry 进行剪切。
func (p *Line) SetLimitInPlot(limitInPlot bool) *Line {
	p.LimitInPlot = limitInPlot
	return p
}

// 指定具体语言，目前内置 'zh-CN' and 'en-US' 两个语言，你也可以使用 G2Plot.registerLocale 方法注册新的语言。语言包格式参考：src/locales/en_US.ts
func (p *Line) SetLocale(locale string) *Line {
	p.Locale = locale
	return p
}

// 数据
func (p *Line) SetData(data interface{}) *Line {
	p.Data = data
	return p
}

// X轴字段
func (p *Line) SetXField(xField string) *Line {
	p.XField = xField
	return p
}

// y轴字段
func (p *Line) SetYField(yField string) *Line {
	p.YField = yField
	return p
}

// 通过 meta 可以全局化配置图表数据元信息，以字段为单位进行配置。在 meta 上的配置将同时影响所有组件的文本信息。传入以字段名为 key，MetaOption 为 value 的配置，同时设置多个字段的元信息。
func (p *Line) SetMeta(meta interface{}) *Line {
	p.Meta = meta

	return p
}

// 是否平滑
func (p *Line) SetSmooth(smooth bool) *Line {
	p.Smooth = smooth
	return p
}
