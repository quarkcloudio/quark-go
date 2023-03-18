package geofence

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type Geofence struct {
	ComponentKey string `json:"componentkey"` // 组件标识
	Component    string `json:"component"`    // 组件名称

	Colon         bool        `json:"colon,omitempty"`        // 配合 label 属性使用，表示是否显示 label 后面的冒号
	Extra         string      `json:"extra,omitempty"`        // 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
	HasFeedback   bool        `json:"hasFeedback,omitempty"`  // 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
	Help          string      `json:"help,omitempty"`         // 提示信息，如不设置，则会根据校验规则自动生成
	Hidden        bool        `json:"hidden,omitempty"`       // 是否隐藏字段（依然会收集和校验字段）
	InitialValue  interface{} `json:"initialValue,omitempty"` // 设置子元素默认值，如果与 Form 的 initialValues 冲突则以 Form 为准
	Label         string      `json:"label,omitempty"`        // label 标签的文本
	LabelAlign    string      `json:"labelAlign,omitempty"`   // 标签文本对齐方式
	LabelCol      interface{} `json:"labelCol,omitempty"`     // label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。你可以通过 Form 的 labelCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准
	Name          string      `json:"name,omitempty"`         // 字段名，支持数组
	NoStyle       bool        `json:"noStyle,omitempty"`      // 为 true 时不带样式，作为纯字段控件使用
	Required      bool        `json:"required,omitempty"`     // 必填样式设置。如不设置，则会根据校验规则自动生成
	Tooltip       string      `json:"tooltip,omitempty"`      // 会在 label 旁增加一个 icon，悬浮后展示配置的信息
	ValuePropName string      `json:"valuePropName"`          // 子节点的值的属性，如 Switch 的是 'checked'。该属性为 getValueProps 的封装，自定义 getValueProps 后会失效
	WrapperCol    interface{} `json:"wrapperCol"`             // 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。你可以通过 Form 的 wrapperCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准

	Api            string        `json:"api,omitempty"` // 获取数据接口
	Ignore         bool          `json:"ignore"`        // 是否忽略保存到数据库，默认为 false
	Rules          []*rule.Rule  `json:"-"`             // 全局校验规则
	CreationRules  []*rule.Rule  `json:"-"`             // 创建页校验规则
	UpdateRules    []*rule.Rule  `json:"-"`             // 编辑页校验规则
	FrontendRules  []*rule.Rule  `json:"frontendRules"` // 前端校验规则，设置字段的校验逻辑
	When           *when.When    `json:"when"`          //
	WhenItem       []*when.Item  `json:"-"`             //
	ShowOnIndex    bool          `json:"-"`             // 在列表页展示
	ShowOnDetail   bool          `json:"-"`             // 在详情页展示
	ShowOnCreation bool          `json:"-"`             // 在创建页面展示
	ShowOnUpdate   bool          `json:"-"`             // 在编辑页面展示
	ShowOnExport   bool          `json:"-"`             // 在导出的Excel上展示
	ShowOnImport   bool          `json:"-"`             // 在导入Excel上展示
	Editable       bool          `json:"-"`             // 表格上是否可编辑
	Column         *table.Column `json:"-"`             // 表格列
	Callback       interface{}   `json:"-"`             // 回调函数

	DefaultValue interface{}            `json:"defaultValue,omitempty"` // 默认选中的选项
	Disabled     bool                   `json:"disabled,omitempty"`     // 整组失效
	Style        map[string]interface{} `json:"style,omitempty"`        // 自定义样式
	Value        interface{}            `json:"value,omitempty"`        // 指定选中项,string[] | number[]
	Zoom         int                    `json:"zoom"`                   // 缩放级别
	MapKey       string                 `json:"mapKey"`                 // 地图Key
}

// 初始化组件
func New() *Geofence {
	return (&Geofence{}).Init()
}

// 初始化
func (p *Geofence) Init() *Geofence {
	p.Component = "geofenceField"
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true
	p.Column = (&table.Column{}).Init()
	p.Value = map[string]interface{}{
		"center": map[string]interface{}{
			"longitude": "116.397724",
			"latitude":  "39.903755",
		},
		"points": []interface{}{},
	}
	p.Zoom = 14
	p.MapKey = "788e08def03f95c670944fe2c78fa76f"
	p.Style = map[string]interface{}{
		"height":    500,
		"width":     "100%",
		"marginTop": "10px",
	}

	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *Geofence) SetKey(key string, crypt bool) *Geofence {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Geofence) SetTooltip(tooltip string) *Geofence {
	p.Tooltip = tooltip

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Geofence) SetColon(colon bool) *Geofence {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Geofence) SetExtra(extra string) *Geofence {
	p.Extra = extra
	return p
}

// 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Geofence) SetHasFeedback(hasFeedback bool) *Geofence {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Geofence) SetHelp(help string) *Geofence {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Geofence) SetNoStyle() *Geofence {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Geofence) SetLabel(label string) *Geofence {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Geofence) SetLabelAlign(align string) *Geofence {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Geofence) SetLabelCol(col interface{}) *Geofence {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Geofence) SetName(name string) *Geofence {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Geofence) SetRequired() *Geofence {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *Geofence) GetFrontendRules(path string) *Geofence {
	var (
		frontendRules []*rule.Rule
		rules         []*rule.Rule
		creationRules []*rule.Rule
		updateRules   []*rule.Rule
	)

	uri := strings.Split(path, "/")
	isCreating := (uri[len(uri)-1] == "create") || (uri[len(uri)-1] == "store")
	isEditing := (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "update")

	if len(p.Rules) > 0 {
		rules = rule.ConvertToFrontendRules(p.Rules)
	}
	if isCreating && len(p.CreationRules) > 0 {
		creationRules = rule.ConvertToFrontendRules(p.CreationRules)
	}
	if isEditing && len(p.UpdateRules) > 0 {
		updateRules = rule.ConvertToFrontendRules(p.UpdateRules)
	}
	if len(rules) > 0 {
		frontendRules = append(frontendRules, rules...)
	}
	if len(creationRules) > 0 {
		frontendRules = append(frontendRules, creationRules...)
	}
	if len(updateRules) > 0 {
		frontendRules = append(frontendRules, updateRules...)
	}

	p.FrontendRules = frontendRules

	return p
}

// 校验规则，设置字段的校验逻辑
func (p *Geofence) SetRules(rules []*rule.Rule) *Geofence {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Geofence) SetCreationRules(rules []*rule.Rule) *Geofence {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Geofence) SetUpdateRules(rules []*rule.Rule) *Geofence {
	p.UpdateRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Geofence) SetValuePropName(valuePropName string) *Geofence {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Geofence) SetWrapperCol(col interface{}) *Geofence {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *Geofence) SetValue(value interface{}) *Geofence {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Geofence) SetDefault(value interface{}) *Geofence {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Geofence) SetDisabled(disabled bool) *Geofence {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Geofence) SetIgnore(ignore bool) *Geofence {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *Geofence) SetWhen(value ...any) *Geofence {
	w := when.New()
	i := when.NewItem()
	var operator string
	var option any

	if len(value) == 2 {
		operator = "="
		option = value[0]
		callback := value[1].(func() interface{})

		i.Body = callback()
	}

	if len(value) == 3 {
		operator = value[0].(string)
		option = value[1]
		callback := value[2].(func() interface{})

		i.Body = callback()
	}

	getOption := untils.InterfaceToString(option)

	switch operator {
	case "=":
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	case ">":
		i.Condition = "<%=String(" + p.Name + ") > '" + getOption + "' %>"
		break
	case "<":
		i.Condition = "<%=String(" + p.Name + ") < '" + getOption + "' %>"
		break
	case "<=":
		i.Condition = "<%=String(" + p.Name + ") <= '" + getOption + "' %>"
		break
	case ">=":
		i.Condition = "<%=String(" + p.Name + ") => '" + getOption + "' %>"
		break
	case "has":
		i.Condition = "<%=(String(" + p.Name + ").indexOf('" + getOption + "') !=-1) %>"
		break
	case "in":
		jsonStr, _ := json.Marshal(option)
		i.Condition = "<%=(" + string(jsonStr) + ".indexOf(" + p.Name + ") !=-1) %>"
		break
	default:
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	}

	i.ConditionName = p.Name
	i.ConditionOperator = operator
	i.Option = option
	p.WhenItem = append(p.WhenItem, i)
	p.When = w.SetItems(p.WhenItem)

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Geofence) HideFromIndex(callback bool) *Geofence {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Geofence) HideFromDetail(callback bool) *Geofence {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Geofence) HideWhenCreating(callback bool) *Geofence {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Geofence) HideWhenUpdating(callback bool) *Geofence {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Geofence) HideWhenExporting(callback bool) *Geofence {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Geofence) HideWhenImporting(callback bool) *Geofence {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Geofence) OnIndexShowing(callback bool) *Geofence {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Geofence) OnDetailShowing(callback bool) *Geofence {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Geofence) ShowOnCreating(callback bool) *Geofence {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Geofence) ShowOnUpdating(callback bool) *Geofence {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Geofence) ShowOnExporting(callback bool) *Geofence {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Geofence) ShowOnImporting(callback bool) *Geofence {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Geofence) OnlyOnIndex() *Geofence {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Geofence) OnlyOnDetail() *Geofence {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Geofence) OnlyOnForms() *Geofence {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Geofence) OnlyOnExport() *Geofence {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Geofence) OnlyOnImport() *Geofence {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Geofence) ExceptOnForms() *Geofence {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Geofence) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *Geofence) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Geofence) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Geofence) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Geofence) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Geofence) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Geofence) SetEditable(editable bool) *Geofence {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Geofence) SetColumn(f func(column *table.Column) *table.Column) *Geofence {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *Geofence) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *Geofence) SetCallback(closure func() interface{}) *Geofence {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Geofence) GetCallback() interface{} {
	return p.Callback
}

// 获取数据接口
func (p *Geofence) SetApi(api string) *Geofence {
	p.Api = api
	return p
}

// zoom
func (p *Geofence) SetZoom(zoom int) *Geofence {
	p.Zoom = zoom
	return p
}

// 高德地图key
func (p *Geofence) SetMapKey(key string) *Geofence {
	p.MapKey = key
	return p
}

// 地图宽度
func (p *Geofence) SetWidth(width interface{}) *Geofence {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 地图高度
func (p *Geofence) SetHeight(height interface{}) *Geofence {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["height"] = height
	p.Style = style

	return p
}

// 中心点
func (p *Geofence) SetCenter(longitude string, latitude string) *Geofence {
	p.Value.(map[string]interface{})["center"] = map[string]interface{}{
		"longitude": longitude,
		"latitude":  latitude,
	}

	return p
}

// 多边形围栏坐标点
func (p *Geofence) SetPoints(points []interface{}) *Geofence {
	p.Value.(map[string]interface{})["points"] = points

	return p
}
