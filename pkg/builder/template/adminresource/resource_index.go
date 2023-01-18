package adminresource

import (
	"reflect"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
)

// 列表页表格主体
func (p *Template) IndexExtraRender(request *builder.Request, templateInstance interface{}) interface{} {
	return nil
}

// 列表页工具栏
func (p *Template) IndexToolBar(request *builder.Request, templateInstance interface{}) interface{} {
	return (&table.ToolBar{}).Init().SetTitle(p.IndexTitle(request, templateInstance)).SetActions(p.IndexActions(request, templateInstance))
}

// 列表标题
func (p *Template) IndexTitle(request *builder.Request, templateInstance interface{}) string {
	return reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("Title").
		String() + "列表"
}

// 列表页组件渲染
func (p *Template) IndexComponentRender(request *builder.Request, templateInstance interface{}, data interface{}) interface{} {
	var component interface{}

	// 列表标题
	title := p.IndexTitle(request, templateInstance)

	// 反射获取参数
	value := reflect.ValueOf(templateInstance).Elem()
	indexPolling := value.FieldByName("IndexPolling").Int()

	// 列表页表格主体
	indexExtraRender := p.IndexExtraRender(request, templateInstance)

	// 列表页工具栏
	indexToolBar := p.IndexToolBar(request, templateInstance)

	// 列表页表格列
	indexColumns := p.IndexColumns(request, templateInstance)

	// 列表页批量操作
	indexTableAlertActions := p.IndexTableAlertActions(request, templateInstance)

	// 列表页搜索栏
	indexSearches := p.IndexSearches(request, templateInstance)

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
	perPage := reflect.
		ValueOf(templateInstance).
		Elem().
		FieldByName("PerPage").Interface()

	// 不分页，直接返回数据
	if reflect.TypeOf(perPage).String() != "int" {
		component = table.SetDatasource(data)
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
func (p *Template) BeforeIndexShowing(request *builder.Request, list []map[string]interface{}) []interface{} {
	result := []interface{}{}
	for _, v := range list {
		result = append(result, v)
	}

	return result
}
