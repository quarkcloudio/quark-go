package actions

import (
	"reflect"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/pkg/builder"
)

type Action struct {
	Name                  string      `json:"name"`
	Reload                string      `json:"reload"`
	ActionType            string      `json:"actionType"`
	SubmitForm            string      `json:"submitForm"`
	Icon                  string      `json:"icon"`
	Type                  string      `json:"type"`
	Size                  string      `json:"size"`
	WithLoading           bool        `json:"withLoading"`
	Fields                interface{} `json:"fields"`
	ConfirmTitle          string      `json:"confirmTitle"`
	ConfirmText           string      `json:"confirmText"`
	ConfirmType           string      `json:"confirmType"`
	OnlyOnIndex           bool        `json:"onlyOnIndex"`
	OnlyOnForm            bool        `json:"onlyOnForm"`
	OnlyOnDetail          bool        `json:"onlyOnDetail"`
	ShowOnIndex           bool        `json:"showOnIndex"`
	ShowOnIndexTableRow   bool        `json:"showOnIndexTableRow"`
	ShowOnIndexTableAlert bool        `json:"showOnIndexTableAlert"`
	ShowOnForm            bool        `json:"showOnForm"`
	ShowOnFormExtra       bool        `json:"showOnFormExtra"`
	ShowOnDetail          bool        `json:"showOnDetail"`
	ShowOnDetailExtra     bool        `json:"showOnDetailExtra"`
}

// 初始化
func (p *Action) ParentInit() interface{} {
	p.ActionType = "ajax"

	return p
}

/**
 * 行为key
 *
 * @return string
 */
func (p *Action) GetUriKey(action interface{}) string {
	uriKey := reflect.TypeOf(action).String()
	uriKey = strings.Replace(uriKey, "*actions.", "", -1)
	uriKey = stringy.New(uriKey).KebabCase("?", "").ToLower()

	return uriKey
}

/**
 * 获取名称
 *
 * @return string
 */
func (p *Action) GetName() string {
	return p.Name
}

/**
 * 执行成功后刷新的组件
 *
 * @return string
 */
func (p *Action) GetReload() string {
	return p.Reload
}

/**
 * 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
 *
 * @return array
 */
func (p *Action) GetApiParams() []string {
	return []string{}
}

/**
 * 执行行为的接口
 *
 * @return string
 */
func (p *Action) GetApi(ctx *builder.Context) string {

	return ""
}

/**
 * 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
 *
 * @return string
 */
func (p *Action) GetActionType() string {
	return p.ActionType
}

/**
 * 当action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
 *
 * @return string
 */
func (p *Action) GetSubmitForm() string {
	return p.SubmitForm
}

/**
 * 设置按钮类型，primary | ghost | dashed | link | text | default
 *
 * @return string
 */
func (p *Action) GetType() string {
	return p.Type
}

/**
 * 设置按钮大小,large | middle | small | default
 *
 * @return string
 */
func (p *Action) GetSize() string {
	return p.Size
}

/**
 * 是否具有loading，当action 的作用类型为ajax,submit时有效
 *
 * @return bool
 */
func (p *Action) GetWithLoading() bool {
	return p.WithLoading
}

/**
 * 设置按钮的图标组件
 *
 * @return string
 */
func (p *Action) GetIcon() string {
	return p.Icon
}

/**
 * 行为表单字段
 *
 * @return mixed
 */
func (p *Action) GetFields() interface{} {
	return p.Fields
}

/**
 * 确认标题
 *
 * @return mixed
 */
func (p *Action) GetConfirmTitle() string {
	return p.ConfirmTitle
}

/**
 * 确认文字
 *
 * @return mixed
 */
func (p *Action) GetConfirmText() string {
	return p.ConfirmText
}

/**
 * 确认类型
 *
 * @return mixed
 */
func (p *Action) GetConfirmType() string {
	return p.ConfirmType
}

/**
 * 设置行为前的确认操作
 *
 * @param  string  title
 * @param  string  text
 * @param  string  confirmType
 * @return p
 */
func (p *Action) WithConfirm(title string, text string, confirmType string) *Action {

	p.ConfirmTitle = title
	p.ConfirmText = text
	p.ConfirmType = confirmType

	return p
}

/**
 * 只在列表页展示
 *
 * @param  bool  value
 * @return p
 */
func (p *Action) SetOnlyOnIndex(value bool) *Action {
	p.OnlyOnIndex = value
	p.ShowOnIndex = value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

/**
 * 除了列表页外展示
 *
 * @return p
 */
func (p *Action) SetExceptOnIndex() *Action {
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnIndexTableAlert = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
	p.ShowOnIndex = false

	return p
}

/**
 * 只在表单页展示
 *
 * @param  bool  value
 * @return p
 */
func (p *Action) SetOnlyOnForm(value bool) *Action {
	p.ShowOnForm = value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

/**
 * 除了表单页外展示
 *
 * @return p
 */
func (p *Action) SetExceptOnForm() *Action {
	p.ShowOnIndexTableAlert = true
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = false
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true

	return p
}

/**
 * 只在表单页右上角自定义区域展示
 *
 * @param  bool  value
 * @return p
 */
func (p *Action) SetOnlyOnFormExtra(value bool) *Action {
	p.ShowOnForm = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

/**
 * 除了表单页右上角自定义区域外展示
 *
 * @return p
 */
func (p *Action) SetExceptOnFormExtra() *Action {
	p.ShowOnIndexTableAlert = true
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = false
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true

	return p
}

/**
 * 只在详情页展示
 *
 * @param  bool  value
 * @return p
 */
func (p *Action) SetOnlyOnDetail(value bool) *Action {
	p.OnlyOnDetail = value
	p.ShowOnDetail = value
	p.ShowOnIndex = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetailExtra = !value

	return p
}

/**
 * 除了详情页外展示
 *
 * @return p
 */
func (p *Action) SetExceptOnDetail() *Action {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnIndexTableRow = true
	p.ShowOnIndexTableAlert = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetailExtra = true

	return p
}

/**
 * 只在详情页右上角自定义区域展示
 *
 * @param  bool  value
 * @return p
 */
func (p *Action) SetOnlyOnDetailExtra(value bool) *Action {
	p.ShowOnForm = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = value

	return p
}

/**
 * 除了详情页右上角自定义区域外展示
 *
 * @return p
 */
func (p *Action) SetExceptOnDetailExtra() *Action {
	p.ShowOnIndexTableAlert = true
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = false

	return p
}

/**
 * 在表格行内展示
 *
 * @param  bool  value
 * @return p
 */
func (p *Action) SetOnlyOnIndexTableRow(value bool) *Action {
	p.ShowOnIndexTableRow = value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableAlert = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

/**
 * 除了表格行内外展示
 *
 * @return p
 */
func (p *Action) SetExceptOnIndexTableRow() *Action {
	p.ShowOnIndexTableRow = false
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableAlert = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true

	return p
}

/**
 * 在表格多选弹出层展示
 *
 * @param  bool  value
 * @return p
 */
func (p *Action) SetOnlyOnIndexTableAlert(value bool) *Action {
	p.ShowOnIndexTableAlert = value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value

	return p
}

/**
 * 除了表格多选弹出层外展示
 *
 * @return p
 */
func (p *Action) SetExceptOnIndexTableAlert() *Action {
	p.ShowOnIndexTableAlert = false
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true

	return p
}

/**
 * 在列表页展示
 *
 * @return p
 */
func (p *Action) SetShowOnIndex() *Action {
	p.ShowOnIndex = true

	return p
}

/**
 * 在表单页展示
 *
 * @return p
 */
func (p *Action) SetShowOnForm() *Action {
	p.ShowOnForm = true

	return p
}

/**
 * 在表单页右上角自定义区域展示
 *
 * @return p
 */
func (p *Action) SetShowOnFormExtra() *Action {
	p.ShowOnFormExtra = true

	return p
}

/**
 * 在详情页展示
 *
 * @return p
 */
func (p *Action) SetShowOnDetail() *Action {
	p.ShowOnDetail = true

	return p
}

/**
 * 在详情页右上角自定义区域展示
 *
 * @return p
 */
func (p *Action) SetShowOnDetailExtra() *Action {
	p.ShowOnDetailExtra = true

	return p
}

/**
 * 在表格行内展示
 *
 * @return p
 */
func (p *Action) SetShowOnIndexTableRow() *Action {
	p.ShowOnIndexTableRow = true

	return p
}

/**
 * 在多选弹出层展示
 *
 * @return p
 */
func (p *Action) SetShowOnIndexTableAlert() *Action {
	p.ShowOnIndexTableAlert = true

	return p
}

/**
 * 判断是否在列表页展示
 *
 * @return bool
 */
func (p *Action) ShownOnIndex() bool {
	if p.OnlyOnIndex == true {
		return true
	}

	if p.OnlyOnDetail {
		return false
	}

	if p.OnlyOnForm {
		return false
	}

	return p.ShowOnIndex
}

/**
 * 判断是否在表单页展示
 *
 * @return bool
 */
func (p *Action) ShownOnForm() bool {
	if p.OnlyOnForm == true {
		return true
	}

	if p.OnlyOnDetail {
		return false
	}

	if p.OnlyOnIndex {
		return false
	}

	return p.ShowOnForm
}

/**
 * 判断是否在详情页展示
 *
 * @return bool
 */
func (p *Action) ShownOnDetail() bool {
	if p.OnlyOnDetail {
		return true
	}

	if p.OnlyOnIndex {
		return false
	}

	if p.OnlyOnForm {
		return false
	}

	return p.ShowOnDetail
}

/**
 * 判断是否在表格行内展示
 *
 * @return bool
 */
func (p *Action) ShownOnIndexTableRow() bool {
	return p.ShowOnIndexTableRow
}

/**
 * 判断是否在多选弹出层展示
 *
 * @return bool
 */
func (p *Action) ShownOnIndexTableAlert() bool {
	return p.ShowOnIndexTableAlert
}

/**
 * 判断是否在表单页右上角自定义区域展示
 *
 * @return bool
 */
func (p *Action) ShownOnFormExtra() bool {
	return p.ShowOnFormExtra
}

/**
 * 判断是否在详情页右上角自定义区域展示
 *
 * @return bool
 */
func (p *Action) ShownOnDetailExtra() bool {
	return p.ShowOnDetailExtra
}
