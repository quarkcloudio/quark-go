package resource

import (
	"reflect"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/table"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

// 列表页表格主体
func (p *Template) IndexExtraRender(ctx *builder.Context) interface{} {
	return nil
}

// 列表页工具栏
func (p *Template) IndexToolBar(ctx *builder.Context) interface{} {
	return (&table.ToolBar{}).
		Init().
		SetTitle(p.IndexTitle(ctx)).
		SetActions(p.IndexActions(ctx))
}

// 列表标题
func (p *Template) IndexTitle(ctx *builder.Context) string {

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	return template.GetTitle() + "列表"
}

// 列表页组件渲染
func (p *Template) IndexComponentRender(ctx *builder.Context, data interface{}) interface{} {
	var component interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 列表标题
	title := p.IndexTitle(ctx)

	// 列表页轮询数据
	indexPolling := template.GetIndexPolling()

	// 列表页表格主体
	indexExtraRender := p.IndexExtraRender(ctx)

	// 列表页工具栏
	indexToolBar := p.IndexToolBar(ctx)

	// 列表页表格列
	indexColumns := p.IndexColumns(ctx)

	// 列表页批量操作
	indexTableAlertActions := p.IndexTableAlertActions(ctx)

	// 列表页搜索栏
	indexSearches := p.IndexSearches(ctx)

	// 表格组件
	table := (&table.Component{}).
		Init().
		SetPolling(int(indexPolling)).
		SetTitle(title).
		SetTableExtraRender(indexExtraRender).
		SetToolBar(indexToolBar).
		SetColumns(indexColumns).
		SetBatchActions(indexTableAlertActions).
		SetSearches(indexSearches)

	// 获取分页
	perPage := template.GetPerPage()
	if perPage == nil {
		return table.SetDatasource(data)
	}

	// 不分页，直接返回数据
	if reflect.TypeOf(perPage).String() != "int" {
		return table.SetDatasource(data)
	} else {
		current := data.(map[string]interface{})["currentPage"]
		perPage := data.(map[string]interface{})["perPage"]
		total := data.(map[string]interface{})["total"]
		items := data.(map[string]interface{})["items"]

		component = table.SetPagination(current.(int), perPage.(int), int(total.(int64)), 1).SetDatasource(items)
	}

	return component
}

// 列表页面显示前回调
func (p *Template) BeforeIndexShowing(ctx *builder.Context, list []map[string]interface{}) []interface{} {
	result := []interface{}{}
	for _, v := range list {
		result = append(result, v)
	}

	return result
}
