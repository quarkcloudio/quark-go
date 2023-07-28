package types

import (
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type Searcher interface {

	// 初始化
	Init(ctx *builder.Context) interface{}

	// 初始化模板
	TemplateInit(ctx *builder.Context) interface{}

	// 获取字段名
	GetColumn(search interface{}) string

	// 获取名称
	GetName() string

	// 获取组件名称
	GetComponent() string

	// 获取接口
	GetApi() string

	// 默认值
	GetDefault() interface{}

	// 执行查询
	Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB

	// 属性
	Options(ctx *builder.Context) interface{}

	// 单向联动,返回数据类型：map[string]string{"field": "you_want_load_field","api": "admin/resource_name/action/select-options"}
	Load(ctx *builder.Context) map[string]string
}
