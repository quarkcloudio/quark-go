package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Image struct {
	Item
}

// 初始化
func (p *Image) Init() *Image {
	p.Component = "imageField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Mode = "single"
	p.Button = "上传图片"
	p.LimitSize = 2
	p.LimitNum = 3
	p.LimitType = []string{
		"image/jpeg",
		"image/png",
	}
	p.Api = "/api/admin/picture/upload"
	p.LimitWH = map[string]int{
		"width":  0,
		"height": 0,
	}

	return p
}

// 上传模式，单图或多图，single|multiple
func (p *Image) SetMode(mode string) *Image {
	if mode == "s" {
		mode = "single"
	}

	if mode == "m" {
		mode = "multiple"
	}

	limits := []string{
		"single", "multiple",
	}

	inSlice := false
	for _, limit := range limits {
		if limit == mode {
			inSlice = true
		}
	}

	if inSlice == false {
		panic("argument must be in 'single', 'multiple'!")
	}

	p.Mode = mode
	return p
}

// 上传文件大小限制
func (p *Image) SetLimitSize(limitSize int) *Image {
	p.LimitSize = limitSize
	return p
}

// 上传文件类型限制
func (p *Image) SetLimitType(limitType []string) *Image {
	p.LimitType = limitType
	return p
}

// 上传文件数量限制
func (p *Image) SetLimitNum(limitNum int) *Image {
	p.LimitNum = limitNum
	return p
}

// 上传图片限制尺寸
func (p *Image) SetLimitWH(width int, height int) *Image {
	p.LimitWH = map[string]int{
		"width":  width,
		"height": height,
	}

	return p
}

// 上传的api接口
func (p *Image) SetApi(api string) *Image {
	p.Api = api
	return p
}

// 上传按钮的标题
func (p *Image) SetButton(text string) *Image {
	p.Button = text
	return p
}
