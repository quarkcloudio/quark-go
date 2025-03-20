package searches

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/selectfield"
)

type Select struct {
	Search
	SelectOptions []selectfield.Option
}

// 初始化模板
func (p *Select) LoadInitData(ctx *quark.Context) interface{} {
	p.Component = "selectField"
	return p
}

// 设置Option
func (p *Select) Option(label string, value interface{}) selectfield.Option {
	return selectfield.Option{
		Value: value,
		Label: label,
	}
}

func (p *Select) Options(ctx *quark.Context) interface{} {
	return []selectfield.Option{}
}

// 单向联动,返回数据类型：map[string]string{"field": "you_want_load_field","api": "admin/resource_name/action/select-options"}
func (p *Select) Load(ctx *quark.Context) map[string]string {
	return nil
}

//	[]selectfield.Option{
//			{Value: 1, Label: "新闻"},
//			{Value: 2, Label: "音乐"},
//			{Value: 3, Label: "体育"},
//		}
//
// 或者
//
// SetOptions(options, "label_name", "value_name")
func (p *Select) SetOptions(options ...interface{}) *Select {
	if len(options) == 1 {
		getOptions, ok := options[0].([]selectfield.Option)
		if ok {
			p.SelectOptions = getOptions
			return p
		}
	}
	if len(options) == 3 {
		p.SelectOptions = selectfield.New().ListToOptions(options[0], options[1].(string), options[2].(string))
	}
	return p
}
