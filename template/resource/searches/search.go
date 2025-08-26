package searches

import (
	"reflect"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcloudio/quark-go/v4"
	"gorm.io/gorm"
)

type Search struct {
	Column    string `json:"column"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Api       string `json:"api"`
}

// 加载初始化数据
func (p *Search) New(ctx *quark.Context) interface{} {
	p.Component = "textField"

	return p
}

// 初始化
func (p *Search) Init(ctx *quark.Context) interface{} {
	return p
}

// 获取字段名
func (p *Search) GetColumn(search interface{}) string {
	if p.Column == "" {
		column := reflect.TypeOf(search).String()
		column = strings.Replace(column, "*searches.", "", -1)
		return stringy.New(column).ToLower()
	}

	return p.Column
}

// 获取名称
func (p *Search) GetName() string {
	return p.Name
}

// 获取组件名称
func (p *Search) GetComponent() string {
	return p.Component
}

// 获取接口
func (p *Search) GetApi() string {
	return p.Api
}

// 默认值
func (p *Search) GetDefault() interface{} {
	return true
}

// 执行查询
func (p *Search) Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query
}

// 属性
func (p *Search) Options(ctx *quark.Context) interface{} {
	return nil
}

// 单向联动,返回数据类型：map[string]string{"field": "you_want_load_field","api": "/api/admin/resource_name/action/select-options"}
func (p *Search) Load(ctx *quark.Context) map[string]string {
	return nil
}
