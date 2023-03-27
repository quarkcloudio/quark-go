package adminresource

import (
	"reflect"
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/descriptions"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/cascader"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/checkbox"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/radio"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/selectfield"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/treeselect"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/component/admin/tabs"
)

// 列表字段
func (p *Template) IndexFields(ctx *builder.Context) interface{} {
	fields := p.getFields(ctx)
	var items []interface{}
	for _, v := range fields.([]interface{}) {
		if v, ok := v.(interface {
			IsShownOnIndex() bool
		}); ok {
			isShownOnIndex := v.IsShownOnIndex()

			if isShownOnIndex {
				items = append(items, v)
			}
		}
	}

	return items
}

// 表格列
func (p *Template) IndexColumns(ctx *builder.Context) interface{} {
	fields := p.IndexFields(ctx)
	var columns []interface{}
	for _, v := range fields.([]interface{}) {
		isShownOnIndex := v.(interface {
			IsShownOnIndex() bool
		}).IsShownOnIndex()
		if isShownOnIndex {
			getColumn := p.fieldToColumn(ctx, v)

			if getColumn != nil {
				columns = append(columns, getColumn)
			}
		}
	}

	// 行内行为
	indexTableRowActions := ctx.Template.(interface {
		IndexTableRowActions(ctx *builder.Context) interface{}
	}).IndexTableRowActions(ctx)
	if len(indexTableRowActions.([]interface{})) > 0 {
		column := (&table.Column{}).
			Init().
			SetTitle("操作").
			SetAttribute("action").
			SetValueType("option").
			SetActions(indexTableRowActions).
			SetFixed("right")
		columns = append(columns, column)
	}

	return columns
}

// 将表单项转换为表格列
func (p *Template) fieldToColumn(ctx *builder.Context, field interface{}) interface{} {
	var options interface{}

	reflectElem := reflect.
		ValueOf(field).
		Elem()

	// 字段
	name := reflectElem.
		FieldByName("Name").
		String()

	// 文字
	label := reflectElem.
		FieldByName("Label").
		String()

	// 组件类型
	component := reflectElem.
		FieldByName("Component").
		String()

	// 是否可编辑
	editable := reflectElem.
		FieldByName("Editable").
		Bool()

	// 是否可编辑
	getColumn := reflectElem.
		FieldByName("Column").
		Interface()

	column := getColumn.(*table.Column).
		SetTitle(label).
		SetAttribute(name)

	switch component {
	case "idField":
		// 是否显示在列表
		onIndexDisplayed := reflectElem.
			FieldByName("OnIndexDisplayed").
			Bool()
		if onIndexDisplayed {
			column = column.SetValueType("text")
		} else {
			return nil
		}
	case "hiddenField":
		return nil
	case "textField":
		column = column.SetValueType("text")
	case "textAreaField":
		column = column.SetValueType("text")
	case "treeSelectField":
		// 获取属性
		options = field.(interface {
			GetOptions() []*treeselect.TreeData
		}).GetOptions()

		// 设置表格列
		column = column.
			SetValueType("treeSelect").
			SetFieldProps(map[string]interface{}{
				"options": options,
			})
	case "cascaderField":
		// 获取属性
		options = field.(interface {
			GetOptions() []*cascader.Option
		}).GetOptions()

		// 设置表格列
		column = column.
			SetValueType("cascader").
			SetFieldProps(map[string]interface{}{
				"options": options,
			})
	case "selectField":
		// 获取属性
		options = field.(interface {
			GetOptions() []*selectfield.Option
		}).GetOptions()

		// 设置表格列
		column = column.
			SetValueType("select").
			SetFieldProps(map[string]interface{}{
				"options": options,
			})

		// 是否设置了过滤项
		filters := column.Filters
		if getfilters, ok := filters.(bool); ok {
			if getfilters {
				// 获取值的枚举，会自动转化把值当成 key 来取出要显示的内容
				valueEnum := field.(interface {
					GetValueEnum() interface{}
				}).GetValueEnum()
				column.SetValueEnum(valueEnum)
			}
		}
	case "checkboxField":
		// 获取属性
		options = field.(interface {
			GetOptions() []*checkbox.Option
		}).GetOptions()

		// 设置表格列
		column = column.
			SetValueType("checkbox").
			SetFieldProps(map[string]interface{}{
				"options": options,
			})

		// 是否设置了过滤项
		filters := column.Filters
		if getfilters, ok := filters.(bool); ok {
			if getfilters {
				// 获取值的枚举，会自动转化把值当成 key 来取出要显示的内容
				valueEnum := field.(interface {
					GetValueEnum() interface{}
				}).GetValueEnum()
				column.SetValueEnum(valueEnum)
			}
		}
	case "radioField":
		// 获取属性
		options = field.(interface {
			GetOptions() []*radio.Option
		}).GetOptions()

		// 设置表格列
		column = column.
			SetValueType("radio").
			SetFieldProps(map[string]interface{}{
				"options": options,
			})

		// 是否设置了过滤项
		filters := column.Filters
		if getfilters, ok := filters.(bool); ok {
			if getfilters {
				// 获取值的枚举，会自动转化把值当成 key 来取出要显示的内容
				valueEnum := field.(interface {
					GetValueEnum() interface{}
				}).GetValueEnum()
				column.SetValueEnum(valueEnum)
			}
		}
	case "switchField":
		// 获取属性
		options = field.(interface {
			GetOptions() interface{}
		}).GetOptions()

		// 设置表格列
		column = column.
			SetValueType("select").
			SetValueEnum(options)

		// 是否设置了过滤项
		filters := column.Filters
		if getfilters, ok := filters.(bool); ok {
			if getfilters {
				// 获取值的枚举，会自动转化把值当成 key 来取出要显示的内容
				valueEnum := field.(interface {
					GetValueEnum() interface{}
				}).GetValueEnum()
				column.SetValueEnum(valueEnum)
			}
		}
	case "imageField":
		column = column.SetValueType("image")
	default:
		column = column.SetValueType(component)
	}

	if editable {

		// 可编辑api地址
		editableApi := strings.Replace(ctx.Path(), "/index", "/editable", -1)

		// 设置编辑项
		column = column.SetEditable(component, options, editableApi)
	}

	return column
}

// 创建页字段
func (p *Template) CreationFields(ctx *builder.Context) interface{} {
	var items []interface{}

	fields := p.getFields(ctx)
	for _, v := range fields.([]interface{}) {
		if v, ok := v.(interface {
			IsShownOnCreation() bool
		}); ok {
			isShownOnCreation := v.IsShownOnCreation()
			if isShownOnCreation {
				items = append(items, v)
			}
		}
	}

	return items
}

// 不包含When组件内字段的创建页字段
func (p *Template) CreationFieldsWithoutWhen(ctx *builder.Context) interface{} {
	fields := p.getFieldsWithoutWhen(ctx)
	var items []interface{}
	for _, v := range fields.([]interface{}) {
		if v, ok := v.(interface {
			IsShownOnCreation() bool
		}); ok {
			isShownOnCreation := v.IsShownOnCreation()
			if isShownOnCreation {
				items = append(items, v)
			}
		}
	}

	return items
}

// 包裹在组件内的创建页字段
func (p *Template) CreationFieldsWithinComponents(ctx *builder.Context) interface{} {
	fields := ctx.Template.(interface {
		Fields(ctx *builder.Context) []interface{}
	}).Fields(ctx)
	var items []interface{}
	for _, v := range fields {
		component := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Component").String()
		if component == "tabPane" {
			body := reflect.
				ValueOf(v).
				Elem().
				FieldByName("Body").Interface()
			var subItems []interface{}
			for _, sv := range body.([]interface{}) {
				if sv, ok := sv.(interface {
					IsShownOnCreation() bool
				}); ok {
					isShownOnCreation := sv.IsShownOnCreation()
					if isShownOnCreation {
						sv.(interface {
							GetFrontendRules(string) interface{}
						}).GetFrontendRules(ctx.Path())
						subItems = append(subItems, sv)
					}
				}
			}
			v.(interface {
				SetBody(interface{}) *tabs.TabPane
			}).SetBody(subItems)

			items = append(items, v)
		} else {
			if v, ok := v.(interface {
				IsShownOnCreation() bool
			}); ok {
				isShownOnCreation := v.IsShownOnCreation()
				if isShownOnCreation {
					v.(interface {
						GetFrontendRules(string) interface{}
					}).GetFrontendRules(ctx.Path())
					items = append(items, v)
				}
			}
		}
	}

	return items
}

// 编辑页字段
func (p *Template) UpdateFields(ctx *builder.Context) interface{} {
	var items []interface{}

	fields := p.getFields(ctx)
	for _, v := range fields.([]interface{}) {
		if v, ok := v.(interface {
			IsShownOnUpdate() bool
		}); ok {
			isShownOnUpdate := v.IsShownOnUpdate()
			if isShownOnUpdate {
				items = append(items, v)
			}
		}
	}

	return items
}

// 不包含When组件内字段的编辑页字段
func (p *Template) UpdateFieldsWithoutWhen(ctx *builder.Context) interface{} {
	fields := p.getFieldsWithoutWhen(ctx)
	var items []interface{}

	for _, v := range fields.([]interface{}) {
		if v, ok := v.(interface {
			IsShownOnUpdate() bool
		}); ok {
			isShownOnUpdate := v.IsShownOnUpdate()
			if isShownOnUpdate {
				items = append(items, v)
			}
		}
	}

	return items
}

// 包裹在组件内的编辑页字段
func (p *Template) UpdateFieldsWithinComponents(ctx *builder.Context) interface{} {
	fields := ctx.Template.(interface {
		Fields(ctx *builder.Context) []interface{}
	}).Fields(ctx)
	var items []interface{}

	for _, v := range fields {
		component := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Component").String()
		if component == "tabPane" {
			body := reflect.
				ValueOf(v).
				Elem().
				FieldByName("Body").Interface()
			var subItems []interface{}
			for _, sv := range body.([]interface{}) {
				if sv, ok := sv.(interface {
					IsShownOnUpdate() bool
				}); ok {
					isShownOnUpdate := sv.IsShownOnUpdate()
					if isShownOnUpdate {
						sv.(interface {
							GetFrontendRules(string) interface{}
						}).GetFrontendRules(ctx.Path())
						subItems = append(subItems, sv)
					}
				}
			}
			v.(interface {
				SetBody(interface{}) *tabs.TabPane
			}).SetBody(subItems)

			items = append(items, v)
		} else {
			if v, ok := v.(interface {
				IsShownOnUpdate() bool
			}); ok {
				isShownOnUpdate := v.IsShownOnUpdate()
				if isShownOnUpdate {
					v.(interface {
						GetFrontendRules(string) interface{}
					}).GetFrontendRules(ctx.Path())
					items = append(items, v)
				}
			}
		}
	}

	return items
}

// 详情页字段
func (p *Template) DetailFields(ctx *builder.Context) interface{} {
	fields := p.getFields(ctx)
	var items []interface{}
	for _, v := range fields.([]interface{}) {
		if v, ok := v.(interface {
			IsShownOnDetail() bool
		}); ok {
			isShownOnDetail := v.IsShownOnDetail()
			if isShownOnDetail {
				items = append(items, v)
			}
		}
	}

	return items
}

// 包裹在组件内的详情页字段
func (p *Template) DetailFieldsWithinComponents(ctx *builder.Context, data map[string]interface{}) interface{} {
	componentType := "description"

	fields := ctx.Template.(interface {
		Fields(ctx *builder.Context) []interface{}
	}).Fields(ctx)
	var items []interface{}

	for _, v := range fields {

		component := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Component").String()

		if component == "tabPane" {

			body := reflect.
				ValueOf(v).
				Elem().
				FieldByName("Body").Interface()

			var subItems []interface{}
			for _, sv := range body.([]interface{}) {
				if sv, ok := sv.(interface {
					IsShownOnDetail() bool
				}); ok {
					isShownOnDetail := sv.IsShownOnDetail()
					if isShownOnDetail {
						getColumn := p.fieldToColumn(ctx, sv)
						subItems = append(subItems, getColumn)
					}
				}
			}

			descriptions := (&descriptions.Component{}).Init().SetStyle(map[string]interface{}{
				"padding": "24px",
			}).
				SetTitle("").
				SetColumn(2).
				SetColumns(subItems).
				SetDataSource(data).
				SetActions(p.DetailActions(ctx))

			v.(interface{ SetBody(interface{}) interface{} }).SetBody(descriptions)
			items = append(items, v)
		} else {
			if v, ok := v.(interface {
				IsShownOnDetail() bool
			}); ok {
				isShownOnDetail := v.IsShownOnDetail()
				if isShownOnDetail {
					getColumn := p.fieldToColumn(ctx, v)
					items = append(items, getColumn)
				}
			}
		}
	}

	if componentType == "description" {
		return (&descriptions.Component{}).
			Init().
			SetStyle(map[string]interface{}{
				"padding": "24px",
			}).
			SetTitle("").
			SetColumn(2).
			SetColumns(items).
			SetDataSource(data).
			SetActions(p.DetailActions(ctx))
	} else {
		return items
	}
}

// 导出字段
func (p *Template) ExportFields(ctx *builder.Context) interface{} {
	fields := p.getFields(ctx)
	var items []interface{}
	for _, v := range fields.([]interface{}) {
		if v, ok := v.(interface {
			IsShownOnExport() bool
		}); ok {
			isShownOnExport := v.IsShownOnExport()
			if isShownOnExport {
				items = append(items, v)
			}
		}
	}

	return items
}

// 导入字段
func (p *Template) ImportFields(ctx *builder.Context) interface{} {
	fields := p.getFields(ctx)
	var items []interface{}
	for _, v := range fields.([]interface{}) {
		if v, ok := v.(interface {
			IsShownOnImport() bool
		}); ok {
			isShownOnImport := v.IsShownOnImport()
			if isShownOnImport {
				items = append(items, v)
			}
		}
	}

	return items
}

// 不包含When组件内字段的导入字段
func (p *Template) ImportFieldsWithoutWhen(ctx *builder.Context) interface{} {
	fields := p.getFieldsWithoutWhen(ctx)
	var items []interface{}
	for _, v := range fields.([]interface{}) {
		if v, ok := v.(interface {
			IsShownOnImport() bool
		}); ok {
			isShownOnImport := v.IsShownOnImport()
			if isShownOnImport {
				items = append(items, v)
			}
		}
	}

	return items
}

// 获取字段
func (p *Template) getFields(ctx *builder.Context) interface{} {
	fields := ctx.Template.(interface {
		Fields(ctx *builder.Context) []interface{}
	}).Fields(ctx)

	return p.findFields(fields, true)
}

// 获取不包含When组件的字段
func (p *Template) getFieldsWithoutWhen(ctx *builder.Context) interface{} {

	fields := ctx.Template.(interface {
		Fields(ctx *builder.Context) []interface{}
	}).Fields(ctx)

	return p.findFields(fields, false)
}

// 查找字段
func (p *Template) findFields(fields interface{}, when bool) interface{} {
	var items []interface{}
	if fields, ok := fields.([]interface{}); ok {
		for _, v := range fields {
			hasBody := reflect.
				ValueOf(v).
				Elem().
				FieldByName("Body").IsValid()
			if hasBody {
				body := reflect.
					ValueOf(v).
					Elem().
					FieldByName("Body").Interface()

				getItems := p.findFields(body, true)
				if getItems, ok := getItems.([]interface{}); ok {
					if len(getItems) > 0 {
						items = append(items, getItems...)
					}
				}
			} else {
				component := reflect.
					ValueOf(v).
					Elem().
					FieldByName("Component").String()
				if strings.Contains(component, "Field") {
					items = append(items, v)
					if when {
						whenFields := p.getWhenFields(v)
						if len(whenFields) > 0 {
							items = append(items, whenFields...)
						}
					}
				}
			}
		}
	}

	return items
}

// 获取When组件中的字段
func (p *Template) getWhenFields(item interface{}) []interface{} {
	var items []interface{}
	whenIsValid := reflect.
		ValueOf(item).
		Elem().
		FieldByName("When").IsValid()
	if !whenIsValid {
		return items
	}

	getWhen := item.(interface {
		GetWhen() *when.Component
	}).GetWhen()

	if getWhen == nil {
		return items
	}
	whenItems := getWhen.Items
	if whenItems == nil {
		return items
	}

	for _, v := range whenItems {
		if v.Body != nil {
			if body, ok := v.Body.([]interface{}); ok {
				if len(body) > 0 {
					items = append(items, body...)
				}
			}
			if body, ok := v.Body.(interface{}); ok {
				items = append(items, body)
			}
		}
	}

	return items
}
