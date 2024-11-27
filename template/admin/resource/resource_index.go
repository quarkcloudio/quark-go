package resource

import (
	"reflect"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/types"
	"github.com/quarkcloudio/quark-go/v3/utils/lister"
)

// 列表页数据转换成树
func (p *Template) IndexTableListToTree(ctx *quark.Context, list []interface{}) []interface{} {
	data := ctx.AllQuerys()
	if search, ok := data["search"].(map[string]interface{}); ok && search != nil {
		return list
	}

	pkName := "id"
	pidName := "pid"
	childrenName := "children"
	rootId := 0
	template := ctx.Template.(types.Resourcer)
	tableListToTree := template.GetTableListToTree()
	if isTableListToTree, ok := tableListToTree.(bool); ok {
		if !isTableListToTree {
			return list
		}
	} else if tableListToTreeMap, ok := tableListToTree.(map[string]interface{}); ok {
		pkName = tableListToTreeMap["pkName"].(string)
		pidName = tableListToTreeMap["pidName"].(string)
		childrenName = tableListToTreeMap["childrenName"].(string)
		rootId = tableListToTreeMap["rootId"].(int)
	}

	tree, _ := lister.ListToTree(list, pkName, pidName, childrenName, rootId)
	return tree
}

// 列表页表格主体
func (p *Template) IndexTableExtraRender(ctx *quark.Context) interface{} {
	return nil
}

// 列表页工具栏
func (p *Template) IndexTableToolBar(ctx *quark.Context) interface{} {
	template := ctx.Template.(types.Resourcer)
	return template.GetTableToolBar(ctx).
		SetTitle(p.IndexTableTitle(ctx)).
		SetActions(p.IndexTableActions(ctx)).
		SetMenu(p.IndexTableMenus(ctx))
}

// 列表页树形表格
func (p *Template) IndexTableTreeBar(ctx *quark.Context) interface{} {
	template := ctx.Template.(types.Resourcer)
	return template.GetTableTreeBar(ctx)
}

// 列表标题
func (p *Template) IndexTableTitle(ctx *quark.Context) string {
	template := ctx.Template.(types.Resourcer)
	return template.GetTitle() + template.GetTableTitleSuffix()
}

// 列表页组件渲染
func (p *Template) IndexComponentRender(ctx *quark.Context, data interface{}) interface{} {
	var component interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取列表页Table实例
	table := template.GetTable()

	// 列表标题
	tableTitle := p.IndexTableTitle(ctx)

	// 列表页轮询数据
	tablePolling := template.GetTablePolling()

	// 列表页表格主体
	tableExtraRender := p.IndexTableExtraRender(ctx)

	// 列表页工具栏
	tableToolBar := p.IndexTableToolBar(ctx)

	// 列表页树形工具
	tableTreeBar := p.IndexTableTreeBar(ctx)

	// 列表页表格列
	tableColumns := p.IndexTableColumns(ctx)

	// 列表页批量操作
	indexTableAlertActions := p.IndexTableAlertActions(ctx)

	// 列表页搜索栏
	indexSearches := p.IndexSearches(ctx)

	// 是否开启树形表格
	tableListToTree := template.GetTableListToTree()
	if tableListToTree != nil {
		data = p.IndexTableListToTree(ctx, data.([]interface{}))
	}

	// 表格组件
	table = table.
		SetPolling(int(tablePolling)).
		SetTitle(tableTitle).
		SetTableExtraRender(tableExtraRender).
		SetToolBar(tableToolBar).
		SetTreeBar(tableTreeBar).
		SetColumns(tableColumns).
		SetBatchActions(indexTableAlertActions).
		SetSearches(indexSearches)

	// 获取分页
	pageSize := template.GetPageSize()
	if pageSize == nil {
		return table.SetDatasource(data)
	}

	// 不分页，直接返回数据
	if reflect.TypeOf(pageSize).String() != "int" {
		return table.SetDatasource(data)
	} else {
		current := data.(map[string]interface{})["page"]
		pageSize := data.(map[string]interface{})["pageSize"]
		pageSizeOptions := data.(map[string]interface{})["pageSizeOptions"]
		total := data.(map[string]interface{})["total"]
		items := data.(map[string]interface{})["items"]
		component = table.
			SetPagination(current.(int), pageSize.(int), int(total.(int64)), 1, pageSizeOptions.([]int)).
			SetDatasource(items)
	}

	return component
}

// 列表页面显示前回调
func (p *Template) BeforeIndexShowing(ctx *quark.Context, list []map[string]interface{}) []interface{} {
	result := []interface{}{}
	for _, v := range list {
		result = append(result, v)
	}

	return result
}
