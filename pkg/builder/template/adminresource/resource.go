package adminresource

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/requests"
	"github.com/quarkcms/quark-go/pkg/component/admin/pagecontainer"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

// 路由路径常量
const (
	IndexPath          = "/api/admin/:resource/index"           // 列表路径
	EditablePath       = "/api/admin/:resource/editable"        // 表格行内编辑路径
	ActionPath         = "/api/admin/:resource/action/:uriKey"  // 执行行为路径
	CreatePath         = "/api/admin/:resource/create"          // 创建页面路径
	StorePath          = "/api/admin/:resource/store"           // 创建方法路径
	EditPath           = "/api/admin/:resource/edit"            // 编辑页面路径
	EditValuesPath     = "/api/admin/:resource/edit/values"     // 获取编辑表单值路径
	SavePath           = "/api/admin/:resource/save"            // 保存编辑值路径
	ImportPath         = "/api/admin/:resource/import"          // 详情页面路径
	ExportPath         = "/api/admin/:resource/export"          // 导出数据路径
	DetailPath         = "/api/admin/:resource/detail"          // 导入数据路径
	ImportTemplatePath = "/api/admin/:resource/import/template" // 导入模板路径
	FormPath           = "/api/admin/:resource/:uriKey/form"    // 通用表单资源路径
)

// 增删改查模板
type Template struct {
	template.Template
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
func (p *Template) Init() interface{} {
	p.TemplateInit()

	return p
}

// 初始化模板
func (p *Template) TemplateInit() interface{} {

	// 初始化数据对象
	p.DB = db.Client

	// 注册路由映射
	p.GET(IndexPath, p.IndexRender)                   // 列表
	p.GET(EditablePath, p.EditableRender)             // 表格行内编辑
	p.Any(ActionPath, p.ActionRender)                 // 执行行为
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
	// 获取数据
	data := (&requests.IndexRequest{}).QueryData(ctx)

	// 获取列表页组件
	body := ctx.Template.(interface {
		IndexComponentRender(ctx *builder.Context, data interface{}) interface{}
	}).IndexComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

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

// 创建页面渲染
func (p *Template) CreationRender(ctx *builder.Context) error {

	// 断言BeforeCreating方法，获取初始数据
	data := ctx.Template.(interface {
		BeforeCreating(ctx *builder.Context) map[string]interface{}
	}).BeforeCreating(ctx)

	// 断言CreationComponentRender方法
	body := ctx.Template.(interface {
		CreationComponentRender(*builder.Context, map[string]interface{}) interface{}
	}).CreationComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 创建方法
func (p *Template) StoreRender(ctx *builder.Context) error {
	return (&requests.StoreRequest{}).Handle(ctx)
}

// 编辑页面渲染
func (p *Template) EditRender(ctx *builder.Context) error {
	// 获取数据
	data := (&requests.EditRequest{}).FillData(ctx)

	// 断言BeforeEditing方法，获取初始数据
	data = ctx.Template.(interface {
		BeforeEditing(*builder.Context, map[string]interface{}) map[string]interface{}
	}).BeforeEditing(ctx, data)

	// 断言UpdateComponentRender方法
	body := ctx.Template.(interface {
		UpdateComponentRender(*builder.Context, map[string]interface{}) interface{}
	}).UpdateComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

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
	data := (&requests.DetailRequest{}).FillData(ctx)

	// 断言方法，获取初始数据
	data = ctx.Template.(interface {
		BeforeDetailShowing(*builder.Context, map[string]interface{}) map[string]interface{}
	}).BeforeDetailShowing(ctx, data)

	// 断言方法
	body := ctx.Template.(interface {
		DetailComponentRender(*builder.Context, map[string]interface{}) interface{}
	}).DetailComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

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

	// 断言BeforeCreating方法，获取初始数据
	data := ctx.Template.(interface {
		BeforeCreating(ctx *builder.Context) map[string]interface{}
	}).BeforeCreating(ctx)

	// 断言CreationComponentRender方法
	body := ctx.Template.(interface {
		CreationComponentRender(*builder.Context, map[string]interface{}) interface{}
	}).CreationComponentRender(ctx, data)

	result := ctx.Template.(interface {
		PageComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageComponentRender(ctx, body)

	return ctx.JSON(200, result)
}

// 页面组件渲染
func (p *Template) PageComponentRender(ctx *builder.Context, body interface{}) interface{} {

	// 页面容器组件渲染
	return ctx.Template.(interface {
		PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{}
	}).PageContainerComponentRender(ctx, body)
}

// 页面容器组件渲染
func (p *Template) PageContainerComponentRender(ctx *builder.Context, body interface{}) interface{} {
	value := reflect.ValueOf(ctx.Template).Elem()

	// 页面标题
	title := value.FieldByName("Title").String()

	// 页面子标题
	subTitle := value.FieldByName("SubTitle").String()

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
