package pages

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/tpl"
	"github.com/quarkcloudio/quark-go/v4/template/page"
)

type Index struct {
	page.Template
}

// 初始化
func (p *Index) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "页面标题"

	return p
}

// 自定义内容页内容
func (p *Index) Content(ctx *quark.Context) interface{} {
	return tpl.New().SetBody("<a href='https://www.baidu.com'>页面内容</a>")
}
