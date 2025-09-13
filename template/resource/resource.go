package resource

import (
	"reflect"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/form"
	"github.com/quarkcloudio/quark-go/v4/component/pagecontainer"
	"github.com/quarkcloudio/quark-go/v4/component/table"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/template/resource/requests"
	"github.com/quarkcloudio/quark-go/v4/template/resource/types"
	"gorm.io/gorm"
)

// 增删改查模板
type Template struct {
	quark.Template
	IndexPath              string           // 列表路径
	EditablePath           string           // 表格行内编辑路径
	ActionPath             string           // 执行行为路径
	ActionValuesPath       string           // 行为表单值路径
	CreatePath             string           // 创建页面路径
	StorePath              string           // 创建方法路径
	EditPath               string           // 编辑页面路径
	EditValuesPath         string           // 获取编辑表单值路径
	SavePath               string           // 保存编辑值路径
	ImportPath             string           // 导入页面路径
	ExportPath             string           // 导出数据路径
	DetailPath             string           // 详情页面路径
	DetailValuesPath       string           // 详情页面值路径
	ImportTemplatePath     string           // 导入模板路径
	FormPath               string           // 表单页面路径
	ContentPath            string           // 自定义内容页面路径
	Title                  string           // 页面标题
	SubTitle               string           // 页面子标题
	BackIcon               bool             // 页面是否携带返回Icon
	Form                   *form.Component  // 表单页Form实例
	Table                  *table.Component // 列表页Table实例
	TableSearch            *table.Search    // 列表Table组件中的搜索实例
	TableColumn            *table.Column    // 列表Table组件中的Column实例
	TableToolBar           *table.ToolBar   // 列表Table组件中的ToolBar实例
	TableTreeBar           *table.TreeBar   // 列表Table组件中的TreeBar实例
	TableTitleSuffix       string           // 列表页表格标题后缀
	TableActionColumnTitle string           // 列表页表格行为列显示文字，既字段的列名
	TableActionColumnWidth int              // 列表页表格行为列的宽度
	TablePolling           int              // 列表页表格是否轮询数据
	TableListToTree        interface{}      // 列表页数据转换为树形结构, true 或者 map[string]interface{}{"pkName": "id",""pidName": "pid","childrenName": "children","rootId":0}
	Export                 bool             // 列表是否具有导出功能
	ExportText             string           // 列表导出按钮文字内容
	PageSize               interface{}      // 列表页分页配置
	PageSizeOptions        []int            // 指定每页可以显示多少条，[10, 20, 50, 100]
	QueryOrder             string           // 全局排序规则
	IndexQueryOrder        string           // 列表页排序规则
	ExportQueryOrder       string           // 导出数据排序规则
	Model                  interface{}      // 挂载模型
}

// 启动模版
func (p *Template) Bootstrap() interface{} {
	p.IndexPath = "/api/admin/:resource/index"                        // 列表路径
	p.EditablePath = "/api/admin/:resource/editable"                  // 表格行内编辑路径
	p.ActionPath = "/api/admin/:resource/action/:uriKey"              // 执行行为路径
	p.ActionValuesPath = "/api/admin/:resource/action/:uriKey/values" // 行为表单值路径
	p.CreatePath = "/api/admin/:resource/create"                      // 创建页面路径
	p.StorePath = "/api/admin/:resource/store"                        // 创建方法路径
	p.EditPath = "/api/admin/:resource/edit"                          // 编辑页面路径
	p.EditValuesPath = "/api/admin/:resource/edit/values"             // 获取编辑表单值路径
	p.SavePath = "/api/admin/:resource/save"                          // 保存编辑值路径
	p.ImportPath = "/api/admin/:resource/import"                      // 导入数据路径
	p.ExportPath = "/api/admin/:resource/export"                      // 导出数据路径
	p.DetailPath = "/api/admin/:resource/detail"                      // 详情页路径
	p.DetailValuesPath = "/api/admin/:resource/detail/values"         // 获取详情页值路径
	p.ImportTemplatePath = "/api/admin/:resource/import/template"     // 导入模板路径
	p.FormPath = "/api/admin/:resource/form"                          // 表单页路径

	return p
}

// 加载初始化路由
func (p *Template) LoadInitRoute() interface{} {
	p.GET(p.IndexPath, p.IndexRender)                   // 列表
	p.GET(p.EditablePath, p.EditableRender)             // 表格行内编辑
	p.Any(p.ActionPath, p.ActionRender)                 // 执行行为
	p.Any(p.ActionValuesPath, p.ActionValuesRender)     // 获取行为表单值
	p.GET(p.CreatePath, p.CreationRender)               // 创建页面
	p.POST(p.StorePath, p.StoreRender)                  // 创建方法
	p.GET(p.EditPath, p.EditRender)                     // 编辑页面
	p.GET(p.EditValuesPath, p.EditValuesRender)         // 获取编辑表单值
	p.POST(p.SavePath, p.SaveRender)                    // 保存编辑值
	p.GET(p.DetailPath, p.DetailRender)                 // 详情页面
	p.GET(p.DetailValuesPath, p.DetailValuesRender)     // 获取详情页值
	p.GET(p.ExportPath, p.ExportRender)                 // 导出数据
	p.POST(p.ImportPath, p.ImportRender)                // 导入数据
	p.GET(p.ImportTemplatePath, p.ImportTemplateRender) // 导入模板
	p.GET(p.FormPath, p.FormRender)                     // 通用表单资源

	return p
}

// 加载初始化数据
func (p *Template) LoadInitData(ctx *quark.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 初始化Form实例
	p.Form = (&form.Component{}).Init()

	// 初始化Table实例
	p.Table = (&table.Component{}).Init()

	// 初始化TableSearch实例
	p.TableSearch = (&table.Search{}).Init()

	// 初始化TableColumn实例
	p.TableColumn = (&table.Column{}).Init()

	// 初始化TableToolBar实例
	p.TableToolBar = (&table.ToolBar{}).Init()

	// 初始化TableTreeBar实例
	p.TableTreeBar = (&table.TreeBar{}).Init()

	// 列表页表格行为列显示文字，既字段的列名
	p.TableActionColumnTitle = "操作"

	// 列表页表格标题后缀
	p.TableTitleSuffix = "列表"

	// 列表导出按钮文字内容
	p.ExportText = "导出"

	// 页面是否携带返回Icon
	p.BackIcon = true

	return p
}

// 模版初始化
func (p *Template) Init(ctx *quark.Context) interface{} {
	return p
}

// 自定义路由
func (p *Template) Route() interface{} {
	// p.GET("/api/admin/:resource/demo", p.Demo)
	return p
}

// 获取Model结构体
func (p *Template) GetModel() interface{} {
	return p.Model
}

// 获取标题
func (p *Template) GetTitle() string {
	return p.Title
}

// 获取子标题
func (p *Template) GetSubTitle() string {
	return p.SubTitle
}

// 页面是否携带返回Icon
func (p *Template) GetBackIcon() bool {
	return p.BackIcon
}

// 获取表单页Form实例
func (p *Template) GetForm() *form.Component {
	return p.Form
}

// 获取列表页Table实例
func (p *Template) GetTable() *table.Component {
	return p.Table
}

// 获取TableSearch实例
func (p *Template) GetTableSearch(ctx *quark.Context) *table.Search {
	return p.TableSearch
}

// 获取TableColumn实例
func (p *Template) GetTableColumn(ctx *quark.Context) *table.Column {
	return p.TableColumn
}

// 获取工具栏实例
func (p *Template) GetTableToolBar(ctx *quark.Context) *table.ToolBar {
	return p.TableToolBar
}

// 获取树形实例
func (p *Template) GetTableTreeBar(ctx *quark.Context) *table.TreeBar {
	return p.TableTreeBar
}

// 列表页表格标题后缀
func (p *Template) GetTableTitleSuffix() string {
	return p.TableTitleSuffix
}

// 列表页表格行为列显示文字，既字段的列名
func (p *Template) GetTableActionColumnTitle() string {
	return p.TableActionColumnTitle
}

// 列表页表格行为列的宽度
func (p *Template) GetTableActionColumnWidth() int {
	return p.TableActionColumnWidth
}

// 获取轮询数据
func (p *Template) GetTablePolling() int {
	return p.TablePolling
}

// 获取分页配置
func (p *Template) GetPageSize() interface{} {
	return p.PageSize
}

// 指定每页可以显示多少条，[10, 20, 50, 100]
func (p *Template) GetPageSizeOptions() []int {
	return p.PageSizeOptions
}

// 列表页列表数据转换为树形结构
func (p *Template) GetTableListToTree() interface{} {
	return p.TableListToTree
}

// 获取全局排序规则
func (p *Template) GetQueryOrder() string {
	return p.QueryOrder
}

// 获取列表页排序规则
func (p *Template) GetIndexQueryOrder() string {
	return p.IndexQueryOrder
}

// 获取导出数据排序规则
func (p *Template) GetExportQueryOrder() string {
	return p.ExportQueryOrder
}

// 获取是否具有导出功能
func (p *Template) GetExport() bool {
	return p.Export
}

// 获取导出按钮文字内容
func (p *Template) GetExportText() string {
	return p.ExportText
}

// 字段
func (p *Template) Fields(ctx *quark.Context) []interface{} {
	return []interface{}{}
}

// 搜索
func (p *Template) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{}
}

// 行为
func (p *Template) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{}
}

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
func (p *Template) MenuItems(ctx *quark.Context) []map[string]string {
	return []map[string]string{}
}

// 数据导出前回调
func (p *Template) BeforeExporting(ctx *quark.Context, list []map[string]interface{}) []interface{} {
	result := []interface{}{}
	for _, v := range list {
		result = append(result, v)
	}

	return result
}

// 数据导入前回调
func (p *Template) BeforeImporting(ctx *quark.Context, list [][]interface{}) [][]interface{} {
	return list
}

// 行内编辑执行完之前回调
func (p *Template) BeforeEditable(ctx *quark.Context, id interface{}, field string, value interface{}) error {
	return nil
}

// 行内编辑执行完之后回调
func (p *Template) AfterEditable(ctx *quark.Context, id interface{}, field string, value interface{}) error {
	return nil
}

// 行为执行完之前回调
func (p *Template) BeforeAction(ctx *quark.Context, uriKey string, query *gorm.DB) error {
	return nil
}

// 行为执行完之后回调
func (p *Template) AfterAction(ctx *quark.Context, uriKey string, query *gorm.DB) error {
	return nil
}

// 列表页渲染
func (p *Template) IndexRender(ctx *quark.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	data := (&requests.IndexRequest{}).QueryData(ctx)

	// 组件渲染
	body := template.IndexComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSONOk("操作成功", result)
}

// 表格行内编辑
func (p *Template) EditableRender(ctx *quark.Context) error {
	return (&requests.EditableRequest{}).Handle(ctx)
}

// 执行行为
func (p *Template) ActionRender(ctx *quark.Context) error {
	return (&requests.ActionRequest{}).Handle(ctx)
}

// 行为表单值
func (p *Template) ActionValuesRender(ctx *quark.Context) error {
	return (&requests.ActionRequest{}).Values(ctx)
}

// 创建页面渲染
func (p *Template) CreationRender(ctx *quark.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 展示前回调
	data := template.BeforeCreating(ctx)

	// 组件渲染
	body := template.CreationComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSONOk("操作成功", result)
}

// 创建方法
func (p *Template) StoreRender(ctx *quark.Context) error {
	data := map[string]interface{}{}
	ctx.Bind(&data)

	template := ctx.Template.(types.Resourcer)

	// 模型结构体
	modelInstance := template.GetModel()

	// 数据库实例
	model := db.Client.Model(modelInstance)

	return template.FormHandle(ctx, model, data)
}

// 编辑页面渲染
func (p *Template) EditRender(ctx *quark.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	data := (&requests.EditRequest{}).FillData(ctx)

	// 展示前回调
	data = template.BeforeEditing(ctx, data)

	// 组件渲染
	body := template.UpdateComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSONOk("操作成功", result)
}

// 获取编辑表单值
func (p *Template) EditValuesRender(ctx *quark.Context) error {
	return (&requests.EditRequest{}).Values(ctx)
}

// 保存编辑值
func (p *Template) SaveRender(ctx *quark.Context) error {
	return (&requests.UpdateRequest{}).Handle(ctx)
}

// 详情页渲染
func (p *Template) DetailRender(ctx *quark.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	data := (&requests.DetailRequest{}).FillData(ctx)

	// 显示前回调
	data = template.BeforeDetailShowing(ctx, data)

	// 组件渲染
	body := template.DetailComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSONOk("操作成功", result)
}

// 获取详情页值
func (p *Template) DetailValuesRender(ctx *quark.Context) error {
	return (&requests.DetailRequest{}).Values(ctx)
}

// 导出数据
func (p *Template) ExportRender(ctx *quark.Context) error {
	return (&requests.ExportRequest{}).Handle(ctx)
}

// 导入数据
func (p *Template) ImportRender(ctx *quark.Context) error {
	indexPath := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("IndexPath").
		String()

	return (&requests.ImportRequest{}).Handle(ctx, indexPath)
}

// 导入数据模板
func (p *Template) ImportTemplateRender(ctx *quark.Context) error {
	return (&requests.ImportTemplateRequest{}).Handle(ctx)
}

// 通用表单资源
func (p *Template) FormRender(ctx *quark.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	data := template.BeforeFormShowing(ctx)

	// 组件渲染
	body := template.CreationComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSONOk("操作成功", result)
}

// 页面组件渲染
func (p *Template) PageComponentRender(ctx *quark.Context, body interface{}) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 页面容器组件渲染
	return template.PageContainerComponentRender(ctx, body)
}

// 页面容器组件渲染
func (p *Template) PageContainerComponentRender(ctx *quark.Context, body interface{}) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 页面标题
	title := template.GetTitle()

	// 页面子标题
	subTitle := template.GetSubTitle()

	// 页面是否携带返回Icon
	backIcon := template.GetBackIcon()

	// 设置头部
	header := (&pagecontainer.PageHeader{}).
		Init().
		SetTitle(title).
		SetSubTitle(subTitle)

	if !backIcon {
		header.SetBackIcon(false)
	}

	return (&pagecontainer.Component{}).
		Init().
		SetHeader(header).
		SetBody(body)
}
