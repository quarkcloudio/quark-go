package captcha

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/tool/captcha"
)

type Index struct {
	captcha.Template
}

// 初始化
func (p *Index) Init(ctx *quark.Context) interface{} {

	// 验证码长度
	p.Length = 4

	return p
}
