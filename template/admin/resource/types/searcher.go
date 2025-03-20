package types

import (
	"github.com/quarkcloudio/quark-go/v3"
	"gorm.io/gorm"
)

type Searcher interface {
	// 加载初始化数据
	New(ctx *quark.Context) interface{}

	// 初始化
	Init(ctx *quark.Context) interface{}

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
	Apply(ctx *quark.Context, query *gorm.DB, value interface{}) *gorm.DB

	// 属性
	Options(ctx *quark.Context) interface{}

	// 单向联动,返回数据类型：map[string]string{"field": "you_want_load_field","api": "admin/resource_name/action/select-options"}
	Load(ctx *quark.Context) map[string]string
}
