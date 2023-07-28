package types

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

type Actioner interface {

	// 初始化
	Init(ctx *builder.Context) interface{}

	// 初始化模板
	TemplateInit(ctx *builder.Context) interface{}

	// 行为key
	GetUriKey(action interface{}) string

	// 获取名称
	GetName() string

	// 执行成功后刷新的组件
	GetReload() string

	// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
	GetApiParams() []string

	// 执行行为的接口
	GetApi() string

	// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
	GetActionType() string

	// 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
	GetSubmitForm() string

	// 设置按钮类型，primary | ghost | dashed | link | text | default
	GetType() string

	// 设置按钮大小,large | middle | small | default
	GetSize() string

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	GetWithLoading() bool

	// 设置按钮的图标组件
	GetIcon() string

	// 行为表单字段
	GetFields() interface{}

	// 确认标题
	GetConfirmTitle() string

	// 确认文字
	GetConfirmText() string

	// 确认类型
	GetConfirmType() string

	// 设置名称
	SetName(name string) *actions.Action

	// 设置执行成功后刷新的组件
	SetReload(componentKey string) *actions.Action

	// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
	SetApiParams(apiParams []string) *actions.Action

	// 执行行为的接口
	SetApi(api string) *actions.Action

	// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
	SetActionType(actionType string) *actions.Action

	// 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
	SetSubmitForm(submitForm string) *actions.Action

	// 设置按钮类型，primary | ghost | dashed | link | text | default
	SetType(buttonType string) *actions.Action

	// 设置按钮大小,large | middle | small | default
	SetSize(size string) *actions.Action

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	SetWithLoading(loading bool) *actions.Action

	// 设置按钮的图标组件
	SetIcon(icon string) *actions.Action

	// 行为表单字段
	SetFields(fields interface{}) *actions.Action

	// 确认标题
	SetConfirmTitle(confirmTitle string) *actions.Action

	// 确认文字
	SetConfirmText(confirmText string) *actions.Action

	// 确认类型
	SetConfirmType(confirmType string) *actions.Action

	// 设置行为前的确认操作
	WithConfirm(title string, text string, confirmType string) *actions.Action

	// 只在列表页展示
	SetOnlyOnIndex(value bool) *actions.Action

	// 除了列表页外展示
	SetExceptOnIndex() *actions.Action

	// 只在表单页展示
	SetOnlyOnForm(value bool) *actions.Action

	// 除了表单页外展示
	SetExceptOnForm() *actions.Action

	// 除了表单页右上角自定义区域外展示
	SetOnlyOnFormExtra(value bool) *actions.Action

	// 只在详情页展示
	SetOnlyOnDetail(value bool) *actions.Action

	// 除了详情页外展示
	SetExceptOnDetail() *actions.Action

	// 只在详情页右上角自定义区域展示
	SetOnlyOnDetailExtra(value bool) *actions.Action

	// 除了详情页右上角自定义区域外展示
	SetExceptOnDetailExtra() *actions.Action

	// 在表格行内展示
	SetOnlyOnIndexTableRow(value bool) *actions.Action

	// 除了表格行内外展示
	SetExceptOnIndexTableRow() *actions.Action

	// 在表格多选弹出层展示
	SetOnlyOnIndexTableAlert(value bool) *actions.Action

	// 除了表格多选弹出层外展示
	SetExceptOnIndexTableAlert() *actions.Action

	// 在列表页展示
	SetShowOnIndex() *actions.Action

	// 在表单页展示
	SetShowOnForm() *actions.Action

	// 在表单页右上角自定义区域展示
	SetShowOnFormExtra() *actions.Action

	// 在详情页展示
	SetShowOnDetail() *actions.Action

	// 在详情页右上角自定义区域展示
	SetShowOnDetailExtra() *actions.Action

	// 在表格行内展示
	SetShowOnIndexTableRow() *actions.Action

	// 在多选弹出层展示
	SetShowOnIndexTableAlert() *actions.Action

	// 判断是否在列表页展示
	ShownOnIndex() bool

	// 判断是否在表单页展示
	ShownOnForm() bool

	// 判断是否在详情页展示
	ShownOnDetail() bool

	// 判断是否在表格行内展示
	ShownOnIndexTableRow() bool

	// 判断是否在多选弹出层展示
	ShownOnIndexTableAlert() bool

	// 判断是否在表单页右上角自定义区域展示
	ShownOnFormExtra() bool

	// 判断是否在详情页右上角自定义区域展示
	ShownOnDetailExtra() bool
}
