package pages

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/mixpage"
)

type Index struct {
	mixpage.Template
}

// 初始化
func (p *Index) Init() interface{} {
	// 初始化模板
	p.TemplateInit()

	return p
}

// 组件渲染
func (p *Index) Content(ctx *builder.Context) interface{} {
	var items []interface{}

	imageStyle := "width: 100%; height: 200px;"

	item := p.SwiperItem(p.Image("https://img.zcool.cn/community/013d1d563828276ac7259e0f3b05b7.png@1280w_1l_2o_100sh.png").SetStyle(imageStyle))
	items = append(items, item)

	item = p.SwiperItem(p.Image("https://img.zcool.cn/community/013d1d563828276ac7259e0f3b05b7.png@1280w_1l_2o_100sh.png").SetStyle(imageStyle))
	items = append(items, item)

	return p.Swiper(items)
}
