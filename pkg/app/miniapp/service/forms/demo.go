package forms

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/miniapp/template/form"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

type Demo struct {
	form.Template
}

// 初始化
func (p *Demo) Init(ctx *builder.Context) interface{} {

	return p
}

// 字段
func (p *Demo) Fields(ctx *builder.Context) []interface{} {
	return []interface{}{
		p.Field().Input("username", "姓名"),
		p.Field().Input("password", "密码"),
	}
}
