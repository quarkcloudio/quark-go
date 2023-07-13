package video

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Component struct {
	component.Element
	Src                       string      `json:"src"`
	Autoplay                  bool        `json:"autoplay"`
	Loop                      bool        `json:"loop"`
	Muted                     bool        `json:"muted"`
	InitialTime               int         `json:"initialTime"`
	Duration                  int         `json:"duration"`
	Controls                  bool        `json:"controls"`
	DanmuList                 interface{} `json:"danmuList"`
	DanmuBtn                  bool        `json:"danmuBtn"`
	EnableDanmu               bool        `json:"enableDanmu"`
	PageGesture               bool        `json:"pageGesture               "`
	Direction                 int         `json:"direction"`
	ShowProgress              bool        `json:"showProgress"`
	ShowFullscreenBtn         bool        `json:"showFullscreenBtn"`
	ShowPlayBtn               bool        `json:"showPlayBtn"`
	ShowCenterPlayBtn         bool        `json:"showCenterPlayBtn"`
	ShowLoading               bool        `json:"showLoading"`
	EnableProgressGesture     bool        `json:"enableProgressGesture"`
	ObjectFit                 string      `json:"objectFit"`
	Poster                    string      `json:"poster"`
	ShowMuteBtn               bool        `json:"showMuteBtn"`
	Title                     string      `json:"title"`
	PlayBtnPosition           string      `json:"playBtnPosition"`
	MobilenetHintType         int         `json:"mobilenetHintType"`
	EnablePlayGesture         bool        `json:"enablePlayGesture"`
	AutoPauseIfNavigate       bool        `json:"autoPauseIfNavigate"`
	AutoPauseIfOpenNative     bool        `json:"autoPauseIfOpenNative"`
	VslideGesture             bool        `json:"vslideGesture"`
	VslideGestureInFullscreen bool        `json:"vslideGestureInFullscreen"`
	AdUnitId                  string      `json:"adUnitId"`
	PosterForCrawler          string      `json:"posterForCrawler"`
	Codec                     string      `json:"codec"`
	HttpCache                 bool        `json:"httpCache"`
	PlayStrategy              int         `json:"playStrategy"`
	Header                    interface{} `json:"header"`
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "video"
	p.Controls = true
	p.ShowProgress = true
	p.ShowFullscreenBtn = true
	p.ShowPlayBtn = true
	p.ShowCenterPlayBtn = true
	p.ShowLoading = true
	p.EnableProgressGesture = true
	p.ObjectFit = "contain"
	p.PlayBtnPosition = "bottom"
	p.MobilenetHintType = 1
	p.AutoPauseIfNavigate = true
	p.AutoPauseIfOpenNative = true
	p.VslideGestureInFullscreen = true
	p.Codec = "hardware"
	p.HttpCache = true
	p.PlayStrategy = 0
	p.SetKey("video", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 要播放视频的资源地址
func (p *Component) SetSrc(src string) *Component {
	p.Src = src
	return p
}

// 是否自动播放
func (p *Component) SetAutoplay(autoplay bool) *Component {
	p.Autoplay = autoplay

	return p
}

// 是否循环播放
func (p *Component) SetLoop(loop bool) *Component {
	p.Loop = loop

	return p
}

// 是否静音播放
func (p *Component) SetMuted(muted bool) *Component {
	p.Muted = muted

	return p
}

// 指定视频初始播放位置，单位为秒（s）。
func (p *Component) SetInitialTime(initialTime int) *Component {
	p.InitialTime = initialTime

	return p
}

// 指定视频时长，单位为秒（s）。
func (p *Component) SetDuration(duration int) *Component {
	p.Duration = duration

	return p
}

// 是否显示默认播放控件（播放/暂停按钮、播放进度、时间）
func (p *Component) SetControls(controls bool) *Component {
	p.Controls = controls

	return p
}

// 弹幕列表
func (p *Component) SetDanmuList(danmuList interface{}) *Component {
	p.DanmuList = danmuList

	return p
}

// 是否显示弹幕按钮，只在初始化时有效，不能动态变更
func (p *Component) SetDanmuBtn(danmuBtn bool) *Component {
	p.DanmuBtn = danmuBtn

	return p
}

// 是否展示弹幕，只在初始化时有效，不能动态变更
func (p *Component) SetEnableDanmu(enableDanmu bool) *Component {
	p.EnableDanmu = enableDanmu

	return p
}

// 在非全屏模式下，是否开启亮度与音量调节手势
func (p *Component) SetPageGesture(pageGesture bool) *Component {
	p.PageGesture = pageGesture

	return p
}

// 设置全屏时视频的方向，不指定则根据宽高比自动判断。有效值为 0（正常竖向）, 90（屏幕逆时针90度）, -90（屏幕顺时针90度）
func (p *Component) SetDirection(direction int) *Component {
	p.Direction = direction

	return p
}

// 若不设置，宽度大于240时才会显示
func (p *Component) SetShowProgress(showProgress bool) *Component {
	p.ShowFullscreenBtn = showProgress

	return p
}

// 是否显示全屏按钮
func (p *Component) SetShowFullscreenBtn(showFullscreenBtn bool) *Component {
	p.ShowFullscreenBtn = showFullscreenBtn

	return p
}

// 是否显示视频底部控制栏的播放按钮
func (p *Component) SetShowPlayBtn(showPlayBtn bool) *Component {
	p.ShowPlayBtn = showPlayBtn

	return p
}

// 是否显示视频中间的播放按钮
func (p *Component) SetShowCenterPlayBtn(showCenterPlayBtn bool) *Component {
	p.ShowCenterPlayBtn = showCenterPlayBtn

	return p
}

// 是否显示loading控件
func (p *Component) SetShowLoading(showLoading bool) *Component {
	p.ShowLoading = showLoading

	return p
}

// 是否开启控制进度的手势
func (p *Component) SetEnableProgressGesture(enableProgressGesture bool) *Component {
	p.EnableProgressGesture = enableProgressGesture

	return p
}

// 当视频大小与 video 容器大小不一致时，视频的表现形式。contain：包含，fill：填充，cover：覆盖
func (p *Component) SetObjectFit(objectFit string) *Component {
	p.ObjectFit = objectFit

	return p
}

// 视频封面的图片网络资源地址，如果 controls 属性值为 false 则设置 poster 无效
func (p *Component) SetPoster(poster string) *Component {
	p.Poster = poster

	return p
}

// 是否显示静音按钮
func (p *Component) SetShowMuteBtn(showMuteBtn bool) *Component {
	p.ShowMuteBtn = showMuteBtn

	return p
}

// 视频的标题，全屏时在顶部展示
func (p *Component) SetTitle(title string) *Component {
	p.Title = title

	return p
}

// 播放按钮的位置
func (p *Component) SetPlayBtnPosition(playBtnPosition string) *Component {
	p.PlayBtnPosition = playBtnPosition

	return p
}

// 移动网络提醒样式：0是不提醒，1是提醒，默认值为1
func (p *Component) SetMobilenetHintType(mobilenetHintType int) *Component {
	p.MobilenetHintType = mobilenetHintType

	return p
}

// 是否开启播放手势，即双击切换播放/暂停
func (p *Component) SetEnablePlayGesture(enablePlayGesture bool) *Component {
	p.EnablePlayGesture = enablePlayGesture

	return p
}

// 当跳转到其它小程序页面时，是否自动暂停本页面的视频
func (p *Component) SetAutoPauseIfNavigate(autoPauseIfNavigate bool) *Component {
	p.AutoPauseIfNavigate = autoPauseIfNavigate

	return p
}

// 当跳转到其它微信原生页面时，是否自动暂停本页面的视频
func (p *Component) SetAutoPauseIfOpenNative(autoPauseIfOpenNative bool) *Component {
	p.AutoPauseIfOpenNative = autoPauseIfOpenNative

	return p
}

// 在非全屏模式下，是否开启亮度与音量调节手势（同 page-gesture）
func (p *Component) SetVslideGesture(vslideGesture bool) *Component {
	p.VslideGesture = vslideGesture

	return p
}

// 在全屏模式下，是否开启亮度与音量调节手势
func (p *Component) SetVslideGestureInFullscreen(vslideGestureInFullscreen bool) *Component {
	p.VslideGestureInFullscreen = vslideGestureInFullscreen

	return p
}

// 视频前贴广告单元ID，更多详情可参考开放能力视频前贴广告
func (p *Component) SetAdUnitId(adUnitId string) *Component {
	p.AdUnitId = adUnitId

	return p
}

// 用于给搜索等场景作为视频封面展示，建议使用无播放 icon 的视频封面图，只支持网络地址
func (p *Component) SetPosterForCrawler(posterForCrawler string) *Component {
	p.PosterForCrawler = posterForCrawler

	return p
}

// 解码器选择，hardware：硬解码（硬解码可以增加解码算力，提高视频清晰度。少部分老旧硬件可能存在兼容性问题）；software：ffmpeg 软解码；
func (p *Component) SetCodec(codec string) *Component {
	p.Codec = codec

	return p
}

// 是否对 http、https 视频源开启本地缓存。缓存策略:开启了此开关的视频源，在视频播放时会在本地保存缓存文件，如果本地缓存池已超过100M，在进行缓存前会清空之前的缓存（不适用于m3u8等流媒体协议）
func (p *Component) SetHttpCache(httpCache bool) *Component {
	p.HttpCache = httpCache

	return p
}

// 播放策略，0：普通模式，适合绝大部分视频播放场景；1：平滑播放模式（降级），增加缓冲区大小，采用open sl解码音频，避免音视频脱轨的问题，可能会降低首屏展现速度、视频帧率，出现开屏音频延迟等。 适用于高码率视频的极端场景；2： M3U8优化模式，增加缓冲区大小，提升视频加载速度和流畅度，可能会降低首屏展现速度。 适用于M3U8在线播放的场景
func (p *Component) SetPlayStrategy(playStrategy int) *Component {
	p.PlayStrategy = playStrategy

	return p
}

// HTTP 请求 Header
func (p *Component) SetHeader(header interface{}) *Component {
	p.Header = header

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "video"

	return p
}
