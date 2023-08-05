package resource

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/pagecontainer"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/requests"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"gorm.io/gorm"
)

// 路由路径常量
const (
	IndexPath          = "/api/admin/:resource/index"                 // 列表路径
	EditablePath       = "/api/admin/:resource/editable"              // 表格行内编辑路径
	ActionPath         = "/api/admin/:resource/action/:uriKey"        // 执行行为路径
	ActionValuesPath   = "/api/admin/:resource/action/:uriKey/values" // 行为表单值路径
	CreatePath         = "/api/admin/:resource/create"                // 创建页面路径
	StorePath          = "/api/admin/:resource/store"                 // 创建方法路径
	EditPath           = "/api/admin/:resource/edit"                  // 编辑页面路径
	EditValuesPath     = "/api/admin/:resource/edit/values"           // 获取编辑表单值路径
	SavePath           = "/api/admin/:resource/save"                  // 保存编辑值路径
	ImportPath         = "/api/admin/:resource/import"                // 详情页面路径
	ExportPath         = "/api/admin/:resource/export"                // 导出数据路径
	DetailPath         = "/api/admin/:resource/detail"                // 导入数据路径
	ImportTemplatePath = "/api/admin/:resource/import/template"       // 导入模板路径
	FormPath           = "/api/admin/:resource/:uriKey/form"          // 通用表单资源路径
)

// 增删改查模板
type Template struct {
	builder.Template
	Title        string                 // 标题
	SubTitle     string                 // 子标题
	PerPage      interface{}            // 分页配置
	IndexPolling int                    // 轮询数据
	IndexOrder   string                 // 排序规则
	Model        interface{}            // 挂载模型
	Field        map[string]interface{} // 注入的字段数据
	WithExport   bool                   // 是否具有导出功能
}

// 初始化
func (p *Template) Init(ctx *builder.Context) interface{} {
	return p
}

// 初始化模板
func (p *Template) TemplateInit(ctx *builder.Context) interface{} {

	// 初始化数据对象
	p.DB = db.Client

	return p
}

// 初始化路由映射
func (p *Template) RouteInit() interface{} {
	p.GET(IndexPath, p.IndexRender)                   // 列表
	p.GET(EditablePath, p.EditableRender)             // 表格行内编辑
	p.Any(ActionPath, p.ActionRender)                 // 执行行为
	p.Any(ActionValuesPath, p.ActionValuesRender)     // 获取行为表单值
	p.GET(CreatePath, p.CreationRender)               // 创建页面
	p.POST(StorePath, p.StoreRender)                  // 创建方法
	p.GET(EditPath, p.EditRender)                     // 编辑页面
	p.GET(EditValuesPath, p.EditValuesRender)         // 获取编辑表单值
	p.POST(SavePath, p.SaveRender)                    // 保存编辑值
	p.GET(DetailPath, p.DetailRender)                 // 详情页面
	p.GET(ExportPath, p.ExportRender)                 // 导出数据
	p.POST(ImportPath, p.ImportRender)                // 导入数据
	p.GET(ImportTemplatePath, p.ImportTemplateRender) // 导入模板
	p.GET(FormPath, p.FormRender)                     // 通用表单资源

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

// 获取分页配置
func (p *Template) GetPerPage() interface{} {
	return p.PerPage
}

// 获取轮询数据
func (p *Template) GetIndexPolling() int {
	return p.IndexPolling
}

// 获取排序规则
func (p *Template) GetIndexOrder() string {
	return p.IndexOrder
}

// 获取注入的字段数据
func (p *Template) GetField() map[string]interface{} {
	return p.Field
}

// 获取是否具有导出功能
func (p *Template) GetWithExport() bool {
	return p.WithExport
}

// 设置单列字段
func (p *Template) SetField(fieldData map[string]interface{}) interface{} {
	p.Field = fieldData

	return p
}

// 字段
func (p *Template) Fields(ctx *builder.Context) []interface{} {
	return []interface{}{}
}

// 搜索
func (p *Template) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{}
}

// 行为
func (p *Template) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{}
}

// 数据导出前回调
func (p *Template) BeforeExporting(ctx *builder.Context, list []map[string]interface{}) []interface{} {
	result := []interface{}{}
	for _, v := range list {
		result = append(result, v)
	}

	return result
}

// 数据导入前回调
func (p *Template) BeforeImporting(ctx *builder.Context, list [][]interface{}) [][]interface{} {
	return list
}

// 表格行内编辑执行完之后回调
func (p *Template) AfterEditable(ctx *builder.Context, id interface{}, field string, value interface{}) error {
	return nil
}

// 行为执行完之后回调
func (p *Template) AfterAction(ctx *builder.Context, uriKey string, query *gorm.DB) error {
	return nil
}

// 列表页渲染
func (p *Template) IndexRender(ctx *builder.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	data := (&requests.IndexRequest{}).QueryData(ctx)

	// 组件渲染
	body := template.IndexComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 表格行内编辑
func (p *Template) EditableRender(ctx *builder.Context) error {
	return (&requests.EditableRequest{}).Handle(ctx)
}

// 执行行为
func (p *Template) ActionRender(ctx *builder.Context) error {
	return (&requests.ActionRequest{}).Handle(ctx)
}

// 行为表单值
func (p *Template) ActionValuesRender(ctx *builder.Context) error {
	return (&requests.ActionRequest{}).Values(ctx)
}

// 创建页面渲染
func (p *Template) CreationRender(ctx *builder.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 展示前回调
	data := template.BeforeCreating(ctx)

	// 组件渲染
	body := template.CreationComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 创建方法
func (p *Template) StoreRender(ctx *builder.Context) error {
	return (&requests.StoreRequest{}).Handle(ctx)
}

// 编辑页面渲染
func (p *Template) EditRender(ctx *builder.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	data := (&requests.EditRequest{}).FillData(ctx)

	// 展示前回调
	data = template.BeforeEditing(ctx, data)

	// 组件渲染
	body := template.UpdateComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 获取编辑表单值
func (p *Template) EditValuesRender(ctx *builder.Context) error {
	return (&requests.EditRequest{}).Values(ctx)
}

// 保存编辑值
func (p *Template) SaveRender(ctx *builder.Context) error {
	return (&requests.UpdateRequest{}).Handle(ctx)
}

// 详情页渲染
func (p *Template) DetailRender(ctx *builder.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	data := (&requests.DetailRequest{}).FillData(ctx)

	// 显示前回调
	data = template.BeforeDetailShowing(ctx, data)

	// 组件渲染
	body := template.DetailComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 导出数据
func (p *Template) ExportRender(ctx *builder.Context) error {
	return (&requests.ExportRequest{}).Handle(ctx)
}

// 导入数据
func (p *Template) ImportRender(ctx *builder.Context) error {
	return (&requests.ImportRequest{}).Handle(ctx, IndexPath)
}

// 导入数据模板
func (p *Template) ImportTemplateRender(ctx *builder.Context) error {
	return (&requests.ImportTemplateRequest{}).Handle(ctx)
}

// 通用表单资源
func (p *Template) FormRender(ctx *builder.Context) error {
	template := ctx.Template.(types.Resourcer)

	// 获取数据
	data := template.BeforeCreating(ctx)

	// 组件渲染
	body := template.CreationComponentRender(ctx, data)

	// 页面渲染
	result := template.PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 页面组件渲染
func (p *Template) PageComponentRender(ctx *builder.Context, body interface{}) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 页面容器组件渲染
	return template.PageContainerComponentRender(ctx, body)
}

// 页面容器组件渲染
func (p *Template) PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{} {
	template := ctx.Template.(types.Resourcer)

	// 页面标题
	title := template.GetTitle()

	// 页面子标题
	subTitle := template.GetSubTitle()

	// 设置头部
	header := (&pagecontainer.PageHeader{}).
		Init().
		SetTitle(title).
		SetSubTitle(subTitle)

	return (&pagecontainer.Component{}).
		Init().
		SetHeader(header).
		SetBody(body)
}
