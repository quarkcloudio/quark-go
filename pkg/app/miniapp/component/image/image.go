package image

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Src                 string `json:"src"`
	Mode                string `json:"mode,omitempty"`
	Webp                bool   `json:"webp,omitempty"`
	Svg                 bool   `json:"svg,omitempty"`
	LazyLoad            bool   `json:"lazyLoad,omitempty"`
	ShowMenuByLongpress bool   `json:"showMenuByLongpress,omitempty"`
	DefaultSource       string `json:"defaultSource,omitempty"`
	ImageMenuPrevent    string `json:"imageMenuPrevent,omitempty"`
	Preview             string `json:"preview,omitempty"`
	OriginalSrc         string `json:"originalSrc,omitempty"`
	AriaLabel           string `json:"ariaLabel,omitempty"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "image"
	p.SetKey("image", component.DEFAULT_CRYPT)

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
func (p *Component) SetMode(mode string) *Component {
	p.Mode = mode
	return p
}

// 默认不解析 webP 格式，只支持网络资源
func (p *Component) SetWebp(webp bool) *Component {
	p.Webp = webp
	return p
}

// 默认不解析 svg 格式，svg 图片只支持 aspectFit
func (p *Component) SetSvg(webp bool) *Component {
	p.Svg = webp
	return p
}

// 图片懒加载。只针对 page 与 scroll-view 下的 image 有效
func (p *Component) SetLazyLoad(lazyLoad bool) *Component {
	p.LazyLoad = lazyLoad
	return p
}

// 开启长按图片显示识别小程序码菜单
func (p *Component) SetShowMenuByLongpress(showMenuByLongpress bool) *Component {
	p.ShowMenuByLongpress = showMenuByLongpress
	return p
}

// 默认图片地址，若设置默认图片地址，会先显示默认图片，等 src 对应的图片加载成功后，再渲染对应的图片。
func (p *Component) SetDefaultSource(defaultSource string) *Component {
	p.DefaultSource = defaultSource
	return p
}

// 阻止长按图片时弹起默认菜单（即将该属性设置为image-menu-prevent="true"或image-menu-prevent），只在初始化时有效，不能动态变更；若不想阻止弹起默认菜单，则不需要设置此属性。注：长按菜单后的操作暂不支持 svg 格式
func (p *Component) SetImageMenuPrevent(imageMenuPrevent string) *Component {
	p.ImageMenuPrevent = imageMenuPrevent
	return p
}

// 点击后是否预览图片。在不设置的情况下，若 image 未监听点击事件且宽度大于 1/4 屏宽，则默认开启
func (p *Component) SetPreview(preview string) *Component {
	p.Preview = preview
	return p
}

// 预览时显示的图片地址
func (p *Component) SetOriginalSrc(originalSrc string) *Component {
	p.OriginalSrc = originalSrc
	return p
}

// 无障碍访问，（属性）元素的额外描述
func (p *Component) SetAriaLabel(ariaLabel string) *Component {
	p.AriaLabel = ariaLabel
	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "image"

	return p
}
