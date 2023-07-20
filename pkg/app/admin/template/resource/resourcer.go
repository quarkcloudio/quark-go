package resource

import (
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type Resourcer interface {

	// 模版接口
	builder.Templater

	// 获取标题
	GetTitle() string

	// 获取子标题
	GetSubTitle() string

	// 获取分页配置
	GetPerPage() interface{}

	// 获取轮询数据
	GetIndexPolling() int

	// 获取排序规则
	GetIndexOrder() string

	// 获取注入的字段数据
	GetField(store string) map[string]interface{}

	// 获取是否具有导出功能
	GetWithExport(ctx *builder.Context) bool

	// 设置单列字段
	SetField(fieldData map[string]interface{}) interface{}

	// 数据导出前回调
	BeforeExporting(ctx *builder.Context, list []map[string]interface{}) []interface{}

	// 数据导入前回调
	BeforeImporting(ctx *builder.Context, list [][]interface{}) [][]interface{}

	// 表格行内编辑执行完之后回调
	AfterEditable(ctx *builder.Context, id interface{}, field string, value interface{}) error

	// 行为执行完之后回调
	AfterAction(ctx *builder.Context, uriKey string, query *gorm.DB) error

	// 列表页渲染
	IndexRender(ctx *builder.Context) error

	// 表格行内编辑
	EditableRender(ctx *builder.Context) error

	// 执行行为
	ActionRender(ctx *builder.Context) error

	// 创建页面渲染
	CreationRender(ctx *builder.Context) error

	// 创建方法
	StoreRender(ctx *builder.Context) error

	// 编辑页面渲染
	EditRender(ctx *builder.Context) error

	// 获取编辑表单值
	EditValuesRender(ctx *builder.Context) error

	// 保存编辑值
	SaveRender(ctx *builder.Context) error

	// 详情页渲染
	DetailRender(ctx *builder.Context) error

	// 导出数据
	ExportRender(ctx *builder.Context) error

	// 导入数据
	ImportRender(ctx *builder.Context) error

	// 导入数据模板
	ImportTemplateRender(ctx *builder.Context) error

	// 通用表单资源
	FormRender(ctx *builder.Context) error

	// 页面组件渲染
	PageComponentRender(ctx *builder.Context, body interface{}) interface{}

	// 页面容器组件渲染
	PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{}

	// 全局查询
	Query(ctx *builder.Context, query *gorm.DB) *gorm.DB

	// 列表查询
	IndexQuery(ctx *builder.Context, query *gorm.DB) *gorm.DB

	// 详情查询
	DetailQuery(ctx *builder.Context, query *gorm.DB) *gorm.DB

	// 导出查询
	ExportQuery(ctx *builder.Context, query *gorm.DB) *gorm.DB

	// 字段
	Fields(ctx *builder.Context) []interface{}

	// 搜索
	Searches(ctx *builder.Context) []interface{}

	// 行为
	Actions(ctx *builder.Context) []interface{}

	// 列表行为
	IndexActions(ctx *builder.Context) interface{}

	// 表格行内行为
	IndexTableRowActions(ctx *builder.Context) interface{}

	// 表格多选弹出层行为
	IndexTableAlertActions(ctx *builder.Context) interface{}

	// 表单页行为
	FormActions(ctx *builder.Context) []interface{}

	// 表单页右上角自定义区域行为
	FormExtraActions(ctx *builder.Context) interface{}

	// 详情页行为
	DetailActions(ctx *builder.Context) []interface{}

	// 详情页行为
	DetailExtraActions(ctx *builder.Context) interface{}

	// 创建表单的接口
	CreationApi(ctx *builder.Context) string

	// 渲染创建页组件
	CreationComponentRender(ctx *builder.Context, data map[string]interface{}) interface{}

	// 创建页面显示前回调
	BeforeCreating(ctx *builder.Context) map[string]interface{}

	// 详情页标题
	DetailTitle(ctx *builder.Context) string

	// 渲染详情页组件
	DetailComponentRender(ctx *builder.Context, data map[string]interface{}) interface{}

	// 详情页页面显示前回调
	BeforeDetailShowing(ctx *builder.Context, data map[string]interface{}) map[string]interface{}

	// 更新表单的接口
	UpdateApi(ctx *builder.Context) string

	// 编辑页面获取表单数据接口
	EditValueApi(request *builder.Context) string

	// 渲染编辑页组件
	UpdateComponentRender(ctx *builder.Context, data map[string]interface{}) interface{}

	// 编辑页面显示前回调
	BeforeEditing(request *builder.Context, data map[string]interface{}) map[string]interface{}

	// 表单接口
	FormApi(ctx *builder.Context) string

	// 表单标题
	FormTitle(ctx *builder.Context) string

	// 保存数据前回调
	BeforeSaving(ctx *builder.Context, submitData map[string]interface{}) (map[string]interface{}, error)

	// 保存数据后回调
	AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) error

	// 列表页表格主体
	IndexExtraRender(ctx *builder.Context) interface{}

	// 列表页工具栏
	IndexToolBar(ctx *builder.Context) interface{}

	// 列表标题
	IndexTitle(ctx *builder.Context) string

	// 列表页组件渲染
	IndexComponentRender(ctx *builder.Context, data interface{}) interface{}

	// 列表页面显示前回调
	BeforeIndexShowing(ctx *builder.Context, list []map[string]interface{}) []interface{}
}
