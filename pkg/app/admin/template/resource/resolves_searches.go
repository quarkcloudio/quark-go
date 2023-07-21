package resource

import (
	"strings"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/cascader"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/treeselect"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/table"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

// 列表页搜索表单
func (p *Template) IndexSearches(ctx *builder.Context) interface{} {
	template := ctx.Template.(types.Resourcer)
	searches := template.Searches(ctx)

	search := (&table.Search{}).Init()
	withExport := template.GetWithExport()
	if withExport {
		search = search.SetExportText("导出").SetExportApi(strings.Replace(ExportPath, ":resource", ctx.Param("resource"), -1))
	}

	for _, v := range searches {

		// 获取组件名称
		component := v.(interface{ GetComponent() string }).GetComponent()

		// label 标签的文本
		label := v.(interface{ GetName() string }).GetName()

		// 字段名，支持数组
		name := v.(interface {
			GetColumn(search interface{}) string
		}).GetColumn(v)

		// 获取接口
		api := v.(interface{ GetApi() string }).GetApi()

		// 获取属性
		options := v.(interface {
			Options(ctx *builder.Context) interface{}
		}).Options(ctx)

		// 搜索栏表单项
		var item interface{}
		var field = &Field{}

		switch component {
		case "textField":
			item = field.
				Text(name, label).
				SetWidth(nil)
		case "selectField":
			item = field.
				Select(name, label).
				SetWidth(nil).
				SetOptions(options.([]*selectfield.Option))
		case "multipleSelectField":
			item = field.
				Select(name, label).
				SetMode("multiple").
				SetWidth(nil).
				SetOptions(options.([]*selectfield.Option))
		case "dateField":
			item = field.
				Date(name, label).
				SetWidth(nil)
		case "datetimeField":
			item = field.
				Datetime(name, label).
				SetWidth(nil)
		case "dateRangeField":
			item = field.
				DateRange(name, label).
				SetWidth(nil)
		case "datetimeRangeField":
			item = field.
				DatetimeRange(name, label).
				SetWidth(nil)
		case "cascaderField":
			item = field.
				Cascader(name, label).
				SetApi(api).
				SetWidth(nil).
				SetOptions(options.([]*cascader.Option))

		case "treeSelectField":
			item = field.
				TreeSelect(name, label).
				SetWidth(nil).
				SetData(options.([]*treeselect.TreeData))
		}

		search = search.SetItems(item)
	}

	return search
}
