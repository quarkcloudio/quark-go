package image

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Src                 string `json:"src"`
	Mode                string `json:"mode"`
	LazyLoad            bool   `json:"lazyLoad"`
	FadeShow            bool   `json:"fadeShow"`
	Webp                bool   `json:"webp"`
	ShowMenuByLongpress bool   `json:"showMenuByLongpress"`
	Draggable           bool   `json:"draggable"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "image"
	p.SetKey("image", component.DEFAULT_CRYPT)
	p.Mode = "scaleToFill"
	p.FadeShow = true
	p.Draggable = true

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 图片资源地址
func (p *Component) SetSrc(src string) *Component {
	p.Src = src
	return p
}

// 图片裁剪、缩放的模式
// scaleToFill:缩放，不保持纵横比缩放图片，使图片的宽高完全拉伸至填满 image 元素
// aspectFit:缩放，保持纵横比缩放图片，使图片的长边能完全显示出来。也就是说，可以完整地将图片显示出来。
// aspectFill:缩放，保持纵横比缩放图片，只保证图片的短边能完全显示出来。也就是说，图片通常只在水平或垂直方向是完整的，另一个方向将会发生截取。
// widthFix:缩放，宽度不变，高度自动变化，保持原图宽高比不变
// heightFix:缩放，高度不变，宽度自动变化，保持原图宽高比不变 App 和 H5 平台 HBuilderX 2.9.3+ 支持、微信小程序需要基础库 2.10.3
// top:裁剪，不缩放图片，只显示图片的顶部区域
// bottom:裁剪，不缩放图片，只显示图片的底部区域
// center:裁剪，不缩放图片，只显示图片的中间区域
// left:裁剪，不缩放图片，只显示图片的左边区域
// right:裁剪，不缩放图片，只显示图片的右边区域
// top left:裁剪，不缩放图片，只显示图片的左上边区域
// top right:裁剪，不缩放图片，只显示图片的右上边区域
// bottom left:裁剪，不缩放图片，只显示图片的左下边区域
// bottom right:裁剪，不缩放图片，只显示图片的右下边区域
func (p *Component) SetMode(mode string) *Component {
	p.Mode = mode

	return p
}

// 图片懒加载。只针对page与scroll-view下的image有效
func (p *Component) SetLazyLoad(lazyLoad bool) *Component {
	p.LazyLoad = lazyLoad

	return p
}

// 图片显示动画效果
func (p *Component) SetFadeShow(fadeShow bool) *Component {
	p.FadeShow = fadeShow

	return p
}

// 在系统不支持webp的情况下是否单独启用webp。默认false，只支持网络资源。webp支持详见下面说明
func (p *Component) SetWebp(webp bool) *Component {
	p.Webp = webp

	return p
}

// 开启长按图片显示识别小程序码菜单
func (p *Component) SetShowMenuByLongpress(showMenuByLongpress bool) *Component {
	p.ShowMenuByLongpress = showMenuByLongpress

	return p
}

// 是否能拖动图片
func (p *Component) SetDraggable(draggable bool) *Component {
	p.Draggable = draggable

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "image"

	return p
}
