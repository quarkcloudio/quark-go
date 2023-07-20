package pages

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/mix/template/page"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Index struct {
	page.Template
}

// 初始化
func (p *Index) Init(ctx *builder.Context) interface{} {
	return p
}

// 广告图
func (p *Index) Banner(ctx *builder.Context) []string {
	return []string{
		"https://img.zcool.cn/community/013d1d563828276ac7259e0f3b05b7.png@1280w_1l_2o_100sh.png",
		"https://img.zcool.cn/community/013d1d563828276ac7259e0f3b05b7.png@1280w_1l_2o_100sh.png",
	}
}

// 组件渲染
func (p *Index) Content(ctx *builder.Context) interface{} {
	return "Hello World!"
}
