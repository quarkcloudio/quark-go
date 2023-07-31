package types

import (
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
	SetName(name string)

	// 设置执行成功后刷新的组件
	SetReload(componentKey string)

	// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
	SetApiParams(apiParams []string)

	// 执行行为的接口
	SetApi(api string)

	// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
	SetActionType(actionType string)

	// 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
	SetSubmitForm(submitForm string)

	// 设置按钮类型，primary | ghost | dashed | link | text | default
	SetType(buttonType string)

	// 设置按钮大小,large | middle | small | default
	SetSize(size string)

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	SetWithLoading(loading bool)

	// 设置按钮的图标组件
	SetIcon(icon string)

	// 行为表单字段
	SetFields(fields interface{})

	// 确认标题
	SetConfirmTitle(confirmTitle string)

	// 确认文字
	SetConfirmText(confirmText string)

	// 确认类型
	SetConfirmType(confirmType string)

	// 设置行为前的确认操作
	WithConfirm(title string, text string, confirmType string)

	// 只在列表页展示
	SetOnlyOnIndex(value bool)

	// 除了列表页外展示
	SetExceptOnIndex()

	// 只在表单页展示
	SetOnlyOnForm(value bool)

	// 除了表单页外展示
	SetExceptOnForm()

	// 除了表单页右上角自定义区域外展示
	SetOnlyOnFormExtra(value bool)

	// 只在详情页展示
	SetOnlyOnDetail(value bool)

	// 除了详情页外展示
	SetExceptOnDetail()

	// 只在详情页右上角自定义区域展示
	SetOnlyOnDetailExtra(value bool)

	// 除了详情页右上角自定义区域外展示
	SetExceptOnDetailExtra()

	// 在表格行内展示
	SetOnlyOnIndexTableRow(value bool)

	// 除了表格行内外展示
	SetExceptOnIndexTableRow()

	// 在表格多选弹出层展示
	SetOnlyOnIndexTableAlert(value bool)

	// 除了表格多选弹出层外展示
	SetExceptOnIndexTableAlert()

	// 在列表页展示
	SetShowOnIndex()

	// 在表单页展示
	SetShowOnForm()

	// 在表单页右上角自定义区域展示
	SetShowOnFormExtra()

	// 在详情页展示
	SetShowOnDetail()

	// 在详情页右上角自定义区域展示
	SetShowOnDetailExtra()

	// 在表格行内展示
	SetShowOnIndexTableRow()

	// 在多选弹出层展示
	SetShowOnIndexTableAlert()

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
