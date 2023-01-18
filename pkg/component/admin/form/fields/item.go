package fields

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/go-basic/uuid"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
)

type Item struct {
	Key                  string                 `json:"-"`
	ComponentKey         string                 `json:"componentKey"`
	Component            string                 `json:"component"`
	Style                map[string]interface{} `json:"style"`
	Tooltip              string                 `json:"tooltip"`
	Colon                bool                   `json:"colon"`
	Value                interface{}            `json:"value"`
	DefaultValue         interface{}            `json:"defaultValue"`
	Extra                string                 `json:"extra"`
	HasFeedback          bool                   `json:"hasFeedback"`
	Help                 string                 `json:"help"`
	NoStyle              bool                   `json:"noStyle"`
	Label                string                 `json:"label"`
	LabelAlign           string                 `json:"labelAlign"`
	LabelCol             interface{}            `json:"labelCol"`
	Name                 string                 `json:"name"`
	Required             bool                   `json:"required"`
	Disabled             bool                   `json:"disabled"`
	Ignore               bool                   `json:"ignore"`
	Rules                []string               `json:"-"`
	RuleMessages         map[string]string      `json:"-"`
	CreationRules        []string               `json:"-"`
	CreationRuleMessages map[string]string      `json:"-"`
	UpdateRules          []string               `json:"-"`
	UpdateRuleMessages   map[string]string      `json:"-"`
	FrontendRules        interface{}            `json:"frontendRules"`
	ValuePropName        string                 `json:"valuePropName"`
	WrapperCol           interface{}            `json:"wrapperCol"`
	When                 interface{}            `json:"when"`
	WhenItem             []interface{}          `json:"-"`
	ShowOnIndex          bool                   `json:"-"`
	ShowOnDetail         bool                   `json:"-"`
	ShowOnCreation       bool                   `json:"-"`
	ShowOnUpdate         bool                   `json:"-"`
	ShowOnExport         bool                   `json:"-"`
	ShowOnImport         bool                   `json:"-"`
	Editable             bool                   `json:"-"`
	Options              interface{}            `json:"options"`
	Column               *table.Column          `json:"-"`
	Callback             interface{}            `json:"-"`
	Placeholder          string                 `json:"placeholder"`
	TreeData             interface{}            `json:"treeData"`
	Mode                 string                 `json:"mode"`
	Size                 string                 `json:"size"`
	AllowClear           bool                   `json:"allowClear"`
	Load                 map[string]string      `json:"load"`
	Button               string                 `json:"button"`
	LimitSize            int                    `json:"limitSize"`
	LimitType            []string               `json:"limitType"`
	LimitNum             int                    `json:"limitNum"`
	LimitWH              map[string]int         `json:"limitWH"`
	Api                  string                 `json:"api"`
	AutoSize             interface{}            `json:"autoSize"`
}

const DEFAULT_KEY = ""
const DEFAULT_CRYPT = true

// 初始化
func (p *Item) InitItem() *Item {
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true
	p.Column = (&table.Column{}).Init()

	return p
}

// 设置Key
func (p *Item) SetKey(key string, crypt bool) *Item {

	if key == "" {
		key = uuid.New()
	}

	if crypt {
		h := md5.New()
		h.Write([]byte(key))
		key = hex.EncodeToString(h.Sum(nil))
	}

	p.Key = key
	p.ComponentKey = key

	return p
}

// Set style.
func (p *Item) SetStyle(style map[string]interface{}) *Item {
	p.Style = style

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *Item) SetTooltip(tooltip string) *Item {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *Item) SetWidth(width interface{}) *Item {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *Item) SetColon(colon bool) *Item {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *Item) SetExtra(extra string) *Item {
	p.Extra = extra
	return p
}

// 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Item) SetHasFeedback(hasFeedback bool) *Item {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *Item) SetHelp(help string) *Item {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *Item) SetNoStyle() *Item {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *Item) SetLabel(label string) *Item {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *Item) SetLabelAlign(align string) *Item {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *Item) SetLabelCol(col interface{}) *Item {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *Item) SetName(name string) *Item {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *Item) SetRequired() *Item {
	p.Required = true
	return p
}

// 解析成前端验证规则
func (p *Item) parseFrontendRules(rules []string, messages map[string]string) []map[string]interface{} {
	result := []map[string]interface{}{}
	values := []string{}
	rule := ""

	for _, v := range rules {
		if strings.Contains(v, ":") {
			values = strings.Split(v, ":")
			rule = values[0]
		} else {
			rule = v
		}

		data := map[string]interface{}{}

		switch rule {
		case "required":
			data = map[string]interface{}{
				"required": true,
				"message":  messages["required"],
			}
		case "min":
			min, _ := strconv.Atoi(values[1])

			data = map[string]interface{}{
				"min":     min,
				"message": messages["min"],
			}
		case "max":
			max, _ := strconv.Atoi(values[1])

			data = map[string]interface{}{
				"max":     max,
				"message": messages["max"],
			}

		case "email":
			data = map[string]interface{}{
				"type":    "email",
				"message": messages["email"],
			}

		case "numeric":
			data = map[string]interface{}{
				"type":    "number",
				"message": messages["numeric"],
			}

		case "url":
			data = map[string]interface{}{
				"type":    "url",
				"message": messages["url"],
			}

		case "integer":
			data = map[string]interface{}{
				"type":    "integer",
				"message": messages["integer"],
			}

		case "date":
			data = map[string]interface{}{
				"type":    "date",
				"message": messages["date"],
			}
		case "boolean":
			data = map[string]interface{}{
				"type":    "boolean",
				"message": messages["boolean"],
			}
		}

		if len(data) > 0 {
			result = append(result, data)
		}
	}

	return result
}

// 自动创建前端验证规则
func (p *Item) BuildFrontendRules(path string) interface{} {
	frontendRules := []map[string]interface{}{}

	var (
		rules         []map[string]interface{}
		creationRules []map[string]interface{}
		updateRules   []map[string]interface{}
	)

	uri := strings.Split(path, "/")
	isCreating := (uri[len(uri)-1] == "create") || (uri[len(uri)-1] == "store")
	isEditing := (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "update")

	if len(p.Rules) > 0 {
		rules = p.parseFrontendRules(p.Rules, p.RuleMessages)
	}
	if isCreating && len(p.CreationRules) > 0 {
		creationRules = p.parseFrontendRules(p.CreationRules, p.CreationRuleMessages)
	}
	if isEditing && len(p.UpdateRules) > 0 {
		updateRules = p.parseFrontendRules(p.UpdateRules, p.UpdateRuleMessages)
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
func (p *Item) SetRules(rules []string, messages map[string]string) *Item {
	p.Rules = rules
	p.RuleMessages = messages

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *Item) SetCreationRules(rules []string, messages map[string]string) *Item {
	p.CreationRules = rules
	p.CreationRuleMessages = messages

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *Item) SetUpdateRules(rules []string, messages map[string]string) *Item {
	p.UpdateRules = rules
	p.UpdateRuleMessages = messages

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *Item) SetValuePropName(valuePropName string) *Item {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *Item) SetWrapperCol(col interface{}) *Item {
	p.WrapperCol = col
	return p
}

// 设置保存值。
func (p *Item) SetValue(value interface{}) *Item {
	p.Value = value
	return p
}

// 设置默认值。
func (p *Item) SetDefault(value interface{}) *Item {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *Item) SetDisabled(disabled bool) *Item {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *Item) SetIgnore(ignore bool) *Item {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *Item) SetWhen(value ...any) *Item {
	whenItem := map[string]any{}
	when := map[string]any{}
	var operator string
	var option any

	if len(value) == 2 {
		operator = "="
		option = value[0]
		callback := value[1].(func() interface{})

		whenItem["body"] = callback()
	}

	if len(value) == 3 {
		operator = value[0].(string)
		option = value[1]
		callback := value[2].(func() interface{})

		whenItem["body"] = callback()
	}

	switch operator {
	case "=":
		whenItem["condition"] = "<%=String(" + p.Name + ") === '" + option.(string) + "' %>"
		break
	case ">":
		whenItem["condition"] = "<%=String(" + p.Name + ") > '" + option.(string) + "' %>"
		break
	case "<":
		whenItem["condition"] = "<%=String(" + p.Name + ") < '" + option.(string) + "' %>"
		break
	case "<=":
		whenItem["condition"] = "<%=String(" + p.Name + ") <= '" + option.(string) + "' %>"
		break
	case ">=":
		whenItem["condition"] = "<%=String(" + p.Name + ") => '" + option.(string) + "' %>"
		break
	case "has":
		whenItem["condition"] = "<%=(String(" + p.Name + ").indexOf('" + option.(string) + "') !=-1) %>"
		break
	case "in":
		jsonStr, _ := json.Marshal(option)
		whenItem["condition"] = "<%=(" + string(jsonStr) + ".indexOf(" + p.Name + ") !=-1) %>"
		break
	default:
		whenItem["condition"] = "<%=String(" + p.Name + ") === '" + option.(string) + "' %>"
		break
	}

	whenItem["condition_name"] = p.Name
	whenItem["condition_operator"] = operator
	whenItem["condition_option"] = option
	p.WhenItem = append(p.WhenItem, whenItem)

	when["component"] = "when"
	when["items"] = whenItem
	p.When = when

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Item) HideFromIndex(callback bool) *Item {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Item) HideFromDetail(callback bool) *Item {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Item) HideWhenCreating(callback bool) *Item {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Item) HideWhenUpdating(callback bool) *Item {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Item) HideWhenExporting(callback bool) *Item {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Item) HideWhenImporting(callback bool) *Item {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *Item) OnIndexShowing(callback bool) *Item {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *Item) OnDetailShowing(callback bool) *Item {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *Item) ShowOnCreating(callback bool) *Item {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *Item) ShowOnUpdating(callback bool) *Item {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *Item) ShowOnExporting(callback bool) *Item {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *Item) ShowOnImporting(callback bool) *Item {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *Item) OnlyOnIndex() *Item {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *Item) OnlyOnDetail() *Item {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *Item) OnlyOnForms() *Item {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *Item) OnlyOnExport() *Item {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *Item) OnlyOnImport() *Item {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *Item) ExceptOnForms() *Item {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *Item) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *Item) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *Item) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *Item) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *Item) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *Item) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *Item) SetEditable(editable bool) *Item {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *Item) SetColumn(f func(column *table.Column) *table.Column) *Item {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *Item) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	if options, ok := p.Options.([]map[string]interface{}); ok {
		for _, v := range options {

			data[v["value"]] = v["label"]
		}
	}

	return data
}

// Switch组件值的枚举 valueEnum
func (p *Item) GetSwitchValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}
	for k, v := range p.Options.(map[string]interface{}) {
		var key int
		if k == "on" {
			key = 1
		} else {
			key = 0
		}
		data[key] = v
	}

	return data
}

// 设置回调函数
func (p *Item) SetCallback(closure func() interface{}) *Item {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *Item) GetCallback() interface{} {
	return p.Callback
}

// 控件占位符
func (p *Item) SetPlaceholder(placeholder string) *Item {
	p.Placeholder = placeholder

	return p
}
