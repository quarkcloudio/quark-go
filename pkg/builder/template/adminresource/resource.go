package adminresource

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 后台增删改查模板
type Template struct {
	template.AdminTemplate
	Title        string // 标题
	SubTitle     string // 子标题
	PerPage      interface{}
	IndexPolling int
	IndexOrder   string
	Model        interface{}
	Field        map[string]interface{}
	WithExport   bool
}

const (
	IndexRoute          = "/api/admin/:resource/index"           // 后台增删改查，列表路由
	EditableRoute       = "/api/admin/:resource/editable"        // 后台增删改查，表格行内编辑路由
	ActionRoute         = "/api/admin/:resource/action/:uriKey"  // 后台增删改查，执行行为路由
	CreateRoute         = "/api/admin/:resource/create"          // 后台增删改查，创建页面路由
	StoreRoute          = "/api/admin/:resource/store"           // 后台增删改查，创建方法路由
	EditRoute           = "/api/admin/:resource/edit"            // 后台增删改查，编辑页面路由
	EditValuesRoute     = "/api/admin/:resource/edit/values"     // 后台增删改查，获取编辑表单值路由
	SaveRoute           = "/api/admin/:resource/save"            // 后台增删改查，保存编辑值路由
	ImportRoute         = "/api/admin/:resource/import"          // 后台增删改查，详情页面路由
	ExportRoute         = "/api/admin/:resource/export"          // 后台增删改查，导出数据路由
	DetailRoute         = "/api/admin/:resource/detail"          // 后台增删改查，导入数据路由
	ImportTemplateRoute = "/api/admin/:resource/import/template" // 后台增删改查，导入模板路由
	FormRoute           = "/api/admin/:resource/:uriKey/form"    // 后台增删改查，通用表单资源
)

// 初始化
func (p *Template) Init() interface{} {
	p.TemplateInit()

	return p
}

// 初始化模板
func (p *Template) TemplateInit() interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 清空路由映射
	p.ClearRouteMapping()

	// 注册路由映射
	p.GET(IndexRoute, "Render")          // 后台增删改查，列表路由
	p.GET(EditableRoute, "Render")       // 后台增删改查，表格行内编辑路由
	p.ANY(ActionRoute, "Render")         // 后台增删改查，执行行为路由
	p.GET(CreateRoute, "Render")         // 后台增删改查，创建页面路由
	p.POST(StoreRoute, "Render")         // 后台增删改查，创建方法路由
	p.GET(EditRoute, "Render")           // 后台增删改查，编辑页面路由
	p.GET(EditValuesRoute, "Render")     // 后台增删改查，获取编辑表单值路由
	p.POST(SaveRoute, "Render")          // 后台增删改查，保存编辑值路由
	p.GET(DetailRoute, "Render")         // 后台增删改查，详情页面路由
	p.GET(ExportRoute, "Render")         // 后台增删改查，导出数据路由
	p.POST(ImportRoute, "Render")        // 后台增删改查，导入数据路由
	p.GET(ImportTemplateRoute, "Render") // 后台增删改查，导入模板路由
	p.GET(FormRoute, "Render")           // 后台增删改查，通用表单资源

	return p
}

// 设置单列字段
func (p *Template) SetField(fieldData map[string]interface{}) interface{} {
	p.Field = fieldData

	return p
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

// 组件渲染
func (p *Template) Render(ctx *builder.Context) interface{} {
	var result interface{}
	switch ctx.FullPath() {
	case IndexRoute: // 列表
		data := (&IndexRequest{}).QueryData(ctx)
		body := p.IndexComponentRender(ctx, data)

		result = ctx.Template.(interface {
			PageComponentRender(ctx *builder.Context, body interface{}) interface{}
		}).PageComponentRender(ctx, body)
	case EditableRoute: // 表格行内编辑
		result = (&EditableRequest{}).Handle(ctx)
	case ActionRoute: // 执行行为路由
		result = (&ActionRequest{}).Handle(ctx)
	case CreateRoute: // 创建页面路由

		// 断言BeforeCreating方法，获取初始数据
		data := ctx.Template.(interface {
			BeforeCreating(ctx *builder.Context) map[string]interface{}
		}).BeforeCreating(ctx)

		// 断言CreationComponentRender方法
		body := ctx.Template.(interface {
			CreationComponentRender(*builder.Context, map[string]interface{}) interface{}
		}).CreationComponentRender(ctx, data)

		result = ctx.Template.(interface {
			PageComponentRender(ctx *builder.Context, body interface{}) interface{}
		}).PageComponentRender(ctx, body)
	case StoreRoute: // 创建方法路由
		result = (&StoreRequest{}).Handle(ctx)
	case EditRoute: // 编辑页面路由

		// 获取数据
		data := (&EditRequest{}).FillData(ctx)

		// 断言BeforeEditing方法，获取初始数据
		data = ctx.Template.(interface {
			BeforeEditing(*builder.Context, map[string]interface{}) map[string]interface{}
		}).BeforeEditing(ctx, data)

		// 断言UpdateComponentRender方法
		body := ctx.Template.(interface {
			UpdateComponentRender(*builder.Context, map[string]interface{}) interface{}
		}).UpdateComponentRender(ctx, data)

		result = ctx.Template.(interface {
			PageComponentRender(ctx *builder.Context, body interface{}) interface{}
		}).PageComponentRender(ctx, body)
	case EditValuesRoute: // 获取编辑表单值

		result = (&EditRequest{}).Values(ctx)
	case SaveRoute: // 保存编辑值

		result = (&UpdateRequest{}).Handle(ctx)
	case DetailRoute: // 详情页面
		data := (&DetailRequest{}).FillData(ctx)

		// 断言方法，获取初始数据
		data = ctx.Template.(interface {
			BeforeDetailShowing(*builder.Context, map[string]interface{}) map[string]interface{}
		}).BeforeDetailShowing(ctx, data)

		// 断言方法
		body := ctx.Template.(interface {
			DetailComponentRender(*builder.Context, map[string]interface{}) interface{}
		}).DetailComponentRender(ctx, data)

		result = ctx.Template.(interface {
			PageComponentRender(ctx *builder.Context, body interface{}) interface{}
		}).PageComponentRender(ctx, body)
	case ExportRoute: // 导出数据
		result = (&ExportRequest{}).Handle(ctx)
	case ImportRoute: // 导入数据
		result = (&ImportRequest{}).Handle(ctx)
	case ImportTemplateRoute: // 导入模板
		result = (&ImportTemplateRequest{}).Handle(ctx)
	case FormRoute: // 通用表单资源
		// 断言BeforeCreating方法，获取初始数据
		data := ctx.Template.(interface {
			BeforeCreating(ctx *builder.Context) map[string]interface{}
		}).BeforeCreating(ctx)

		// 断言CreationComponentRender方法
		body := ctx.Template.(interface {
			CreationComponentRender(*builder.Context, map[string]interface{}) interface{}
		}).CreationComponentRender(ctx, data)

		result = ctx.Template.(interface {
			PageComponentRender(ctx *builder.Context, body interface{}) interface{}
		}).PageComponentRender(ctx, body)
	}

	return ctx.JSON(200, result)
}
