package types

import (
	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/form"
	"github.com/quarkcloudio/quark-go/v4/component/table"
	"gorm.io/gorm"
)

type Resourcer interface {

	// 模版接口
	quark.Templater

	// 获取Model结构体
	GetModel() interface{}

	// 获取标题
	GetTitle() string

	// 获取子标题
	GetSubTitle() string

	// 页面是否携带返回Icon
	GetBackIcon() bool

	// 获取表单页Form实例
	GetForm() *form.Component

	// 获取列表页Table实例
	GetTable() *table.Component

	// 获取TableSearch实例
	GetTableSearch(ctx *quark.Context) *table.Search

	// 获取TableColumn实例
	GetTableColumn(ctx *quark.Context) *table.Column

	// 获取工具栏实例
	GetTableToolBar(ctx *quark.Context) *table.ToolBar

	// 获取树形实例
	GetTableTreeBar(ctx *quark.Context) *table.TreeBar

	// 列表页表格标题后缀
	GetTableTitleSuffix() string

	// 列表页表格行为列显示文字，既字段的列名
	GetTableActionColumnTitle() string

	// 列表页表格行为列的宽度
	GetTableActionColumnWidth() int

	// 获取轮询数据
	GetTablePolling() int

	// 列表页列表数据转换为树形结构
	GetTableListToTree() interface{}

	// 获取是否具有导出功能
	GetExport() bool

	// 获取导出按钮文字内容
	GetExportText() string

	// 获取分页配置
	GetPageSize() interface{}

	// 指定每页可以显示多少条，[10, 20, 50, 100]
	GetPageSizeOptions() []int

	// 获取全局排序规则
	GetQueryOrder() string

	// 获取列表页排序规则
	GetIndexQueryOrder() string

	// 获取导出数据排序规则
	GetExportQueryOrder() string

	// 数据导出前回调
	BeforeExporting(ctx *quark.Context, list []map[string]interface{}) []interface{}

	// 数据导入前回调
	BeforeImporting(ctx *quark.Context, list [][]interface{}) [][]interface{}

	// 表格行内编辑执行完之前回调
	BeforeEditable(ctx *quark.Context, id interface{}, field string, value interface{}) error

	// 表格行内编辑执行完之后回调
	AfterEditable(ctx *quark.Context, id interface{}, field string, value interface{}) error

	// 行为执行完之前回调
	BeforeAction(ctx *quark.Context, uriKey string, query *gorm.DB) error

	// 行为执行完之后回调
	AfterAction(ctx *quark.Context, uriKey string, query *gorm.DB) error

	// 列表页渲染
	IndexRender(ctx *quark.Context) error

	// 表格行内编辑
	EditableRender(ctx *quark.Context) error

	// 执行行为
	ActionRender(ctx *quark.Context) error

	// 创建页面渲染
	CreationRender(ctx *quark.Context) error

	// 创建方法
	StoreRender(ctx *quark.Context) error

	// 编辑页面渲染
	EditRender(ctx *quark.Context) error

	// 获取编辑表单值
	EditValuesRender(ctx *quark.Context) error

	// 保存编辑值
	SaveRender(ctx *quark.Context) error

	// 详情页渲染
	DetailRender(ctx *quark.Context) error

	// 获取详情页数据
	DetailValuesRender(ctx *quark.Context) error

	// 导出数据
	ExportRender(ctx *quark.Context) error

	// 导入数据
	ImportRender(ctx *quark.Context) error

	// 导入数据后回调
	AfterImported(ctx *quark.Context, id int, data map[string]interface{}, result *gorm.DB) (err error)

	// 导入数据模板
	ImportTemplateRender(ctx *quark.Context) error

	// 通用表单资源
	FormRender(ctx *quark.Context) error

	// 全局查询
	Query(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 行为查询
	ActionQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 详情查询
	DetailQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 编辑查询
	EditQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 表格行内编辑查询
	EditableQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 导出查询
	ExportQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 列表查询
	IndexQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 更新查询
	UpdateQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 字段
	Fields(ctx *quark.Context) []interface{}

	// 搜索
	Searches(ctx *quark.Context) []interface{}

	// 行为
	Actions(ctx *quark.Context) []interface{}

	// 列表工具栏菜单项，示例如下：
	//
	//	[]map[string]string{
	//		{
	//			"key":   "day",
	//			"label": "日账单",
	//		},
	//		{
	//			"key":   "week",
	//			"label": "周账单",
	//		},
	//	}
	MenuItems(ctx *quark.Context) []map[string]string

	// 创建行为组件
	BuildAction(ctx *quark.Context, item interface{}) interface{}

	// 创建行为接口
	BuildActionApi(ctx *quark.Context, params []string, uriKey string) string

	// 表格行为
	IndexTableActions(ctx *quark.Context) interface{}

	// 表格行内行为
	IndexTableRowActions(ctx *quark.Context) interface{}

	// 表单页行为
	FormActions(ctx *quark.Context) []interface{}

	// 表单页右上角自定义区域行为
	FormExtraActions(ctx *quark.Context) interface{}

	// 详情页行为
	DetailActions(ctx *quark.Context) []interface{}

	// 详情页行为
	DetailExtraActions(ctx *quark.Context) interface{}

	// 创建表单的接口
	CreationApi(ctx *quark.Context) string

	// 渲染创建页组件
	CreationComponentRender(ctx *quark.Context, data map[string]interface{}) interface{}

	// 创建页面显示前回调
	BeforeCreating(ctx *quark.Context) map[string]interface{}

	// 详情页标题
	DetailTitle(ctx *quark.Context) string

	// 渲染详情页组件
	DetailComponentRender(ctx *quark.Context, data map[string]interface{}) interface{}

	// 详情页页面显示前回调
	BeforeDetailShowing(ctx *quark.Context, data map[string]interface{}) map[string]interface{}

	// 更新表单的接口
	UpdateApi(ctx *quark.Context) string

	// 编辑页面获取表单数据接口
	EditValueApi(request *quark.Context) string

	// 渲染编辑页组件
	UpdateComponentRender(ctx *quark.Context, data map[string]interface{}) interface{}

	// 编辑页面显示前回调
	BeforeEditing(request *quark.Context, data map[string]interface{}) map[string]interface{}

	// 详情页面获取表单数据接口
	DetailValueApi(request *quark.Context) string

	// 表单接口
	FormApi(ctx *quark.Context) string

	// 表单标题
	FormTitle(ctx *quark.Context) string

	// 表单页面显示前回调
	BeforeFormShowing(ctx *quark.Context) map[string]interface{}

	// 表单执行
	FormHandle(ctx *quark.Context, query *gorm.DB, data map[string]interface{}) error

	// 保存数据前回调
	BeforeSaving(ctx *quark.Context, submitData map[string]interface{}) (map[string]interface{}, error)

	// 保存数据后回调
	AfterSaved(ctx *quark.Context, id int, data map[string]interface{}, result *gorm.DB) error

	// 保存数据后跳转回调
	AfterSavedRedirectTo(ctx *quark.Context, id int, data map[string]interface{}, err error) error

	// 列表页表格主体
	IndexTableExtraRender(ctx *quark.Context) interface{}

	// 列表页工具栏
	IndexTableToolBar(ctx *quark.Context) interface{}

	// 列表工具栏菜单
	IndexTableMenu(ctx *quark.Context) interface{}

	// 列表工具栏菜单项，示例如下：
	//
	//	[]map[string]string{
	//		{
	//			"key":   "day",
	//			"label": "日账单",
	//		},
	//		{
	//			"key":   "week",
	//			"label": "周账单",
	//		},
	//	}
	IndexTableMenuItems(ctx *quark.Context) []map[string]string

	// 列表页树形表格
	IndexTableTreeBar(ctx *quark.Context) interface{}

	// 列表标题
	IndexTableTitle(ctx *quark.Context) string

	// 列表页组件渲染
	IndexComponentRender(ctx *quark.Context, data interface{}) interface{}

	// 列表页面显示前回调
	BeforeIndexShowing(ctx *quark.Context, list []map[string]interface{}) []interface{}

	// 列表页字段
	IndexFields(ctx *quark.Context) interface{}

	// 创建页字段
	CreationFields(ctx *quark.Context) interface{}

	// 不包含When组件内字段的创建页字段
	CreationFieldsWithoutWhen(ctx *quark.Context) interface{}

	// 包裹在组件内的创建页字段
	CreationFieldsWithinComponents(ctx *quark.Context) interface{}

	// 编辑页字段
	UpdateFields(ctx *quark.Context) interface{}

	// 不包含When组件内字段的编辑页字段
	UpdateFieldsWithoutWhen(ctx *quark.Context) interface{}

	// 包裹在组件内的编辑页字段
	UpdateFieldsWithinComponents(ctx *quark.Context) interface{}

	// 解析表单组件内的字段
	FormFieldsParser(ctx *quark.Context, fields interface{}) interface{}

	// 详情页字段
	DetailFields(ctx *quark.Context) interface{}

	// 包裹在组件内的详情页字段
	DetailFieldsWithinComponents(ctx *quark.Context, initApi interface{}, data map[string]interface{}) interface{}

	// 导出字段
	ExportFields(ctx *quark.Context) interface{}

	// 导入字段
	ImportFields(ctx *quark.Context) interface{}

	// 筛选表单
	Filters(ctx *quark.Context) []interface{}

	// 创建执行行为查询
	BuildActionQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 创建详情页查询
	BuildDetailQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 创建编辑查询
	BuildEditQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 创建表格行内编辑查询
	BuildEditableQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 创建导出查询
	BuildExportQuery(ctx *quark.Context, query *gorm.DB, search []interface{}, filters []interface{}, columnFilters map[string]interface{}, orderings map[string]interface{}) *gorm.DB

	// 创建列表查询
	BuildIndexQuery(ctx *quark.Context, query *gorm.DB, search []interface{}, filters []interface{}, columnFilters map[string]interface{}, orderings map[string]interface{}) *gorm.DB

	// 创建更新查询
	BuildUpdateQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB

	// 创建请求的验证器
	ValidatorForCreation(ctx *quark.Context, data map[string]interface{}) error

	// 更新请求的验证器
	ValidatorForUpdate(ctx *quark.Context, data map[string]interface{}) error

	// 导入请求的验证器
	ValidatorForImport(ctx *quark.Context, data map[string]interface{}) error
}
