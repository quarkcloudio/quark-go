package actions

import (
	"reflect"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcloudio/quark-go/v4"
	"gorm.io/gorm"
)

type Action struct {
	Name       interface{} `json:"name"`       // 设置按钮文字；支持js表达式例如：<%= (status==1 ? '禁用' : '启用') %>，行为在表格行时，可以使用当前行的任意字段值，示例中status即为当前行的“状态”字段
	Reload     string      `json:"reload"`     // 执行成功后刷新的组件
	ApiParams  []string    `json:"apiParams"`  // 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
	Api        string      `json:"api"`        // 行为接口
	ActionType string      `json:"actionType"` // 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
	SubmitForm string      `json:"submitForm"` // 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
	Block      bool        `json:"block"`
	Batch      bool        `json:"batch"`
	Danger     bool        `json:"danger"`
	Disabled   bool        `json:"disabled"`
	Ghost      bool        `json:"ghost"`
	// 设置按钮的图标组件：
	// "ant-design:database-outlined", "ant-design:server-outlined", "ant-design:mobile-outlined", "ant-design:tablet-outlined", "ant-design:redenvelope-outlined",
	// "ant-design:book-outlined", "ant-design:filedone-outlined", "ant-design:reconciliation-outlined", "ant-design:file-exception-outlined",
	// "ant-design:filesync-outlined", "ant-design:filesearch-outlined", "ant-design:solution-outlined", "ant-design:fileprotect-outlined",
	// "ant-design:file-add-outlined", "ant-design:file-excel-outlined", "ant-design:file-exclamation-outlined", "ant-design:file-pdf-outlined",
	// "ant-design:file-image-outlined", "ant-design:file-markdown-outlined", "ant-design:file-unknown-outlined", "ant-design:file-ppt-outlined",
	// "ant-design:file-word-outlined", "ant-design:file-outlined", "ant-design:file-zip-outlined", "ant-design:file-text-outlined", "ant-design:file-copy-outlined",
	// "ant-design:snippets-outlined", "ant-design:audit-outlined", "ant-design:diff-outlined", "ant-design:Batchfolding-outlined", "ant-design:securityscan-outlined",
	// "ant-design:propertysafety-outlined", "ant-design:insurance-outlined", "ant-design:alert-outlined", "ant-design:delete-outlined", "ant-design:hourglass-outlined",
	// "ant-design:bulb-outlined", "ant-design:experiment-outlined", "ant-design:bell-outlined", "ant-design:trophy-outlined", "ant-design:rest-outlined",
	// "ant-design:USB-outlined", "ant-design:skin-outlined", "ant-design:home-outlined", "ant-design:bank-outlined", "ant-design:filter-outlined",
	// "ant-design:funnelplot-outlined", "ant-design:like-outlined", "ant-design:unlike-outlined", "ant-design:unlock-outlined", "ant-design:lock-outlined",
	// "ant-design:customerservice-outlined", "ant-design:flag-outlined", "ant-design:moneycollect-outlined", "ant-design:medicinebox-outlined",
	// "ant-design:shop-outlined", "ant-design:rocket-outlined", "ant-design:shopping-outlined", "ant-design:folder-outlined", "ant-design:folder-open-outlined",
	// "ant-design:folder-add-outlined", "ant-design:deploymentunit-outlined", "ant-design:accountbook-outlined", "ant-design:contacts-outlined",
	// "ant-design:carryout-outlined", "ant-design:calendar-check-outlined", "ant-design:calendar-outlined", "ant-design:scan-outlined",
	// "ant-design:select-outlined", "ant-design:boxplot-outlined", "ant-design:build-outlined", "ant-design:sliders-outlined", "ant-design:laptop-outlined",
	// "ant-design:barcode-outlined", "ant-design:camera-outlined", "ant-design:cluster-outlined", "ant-design:gateway-outlined", "ant-design:car-outlined",
	// "ant-design:printer-outlined", "ant-design:read-outlined", "ant-design:cloud-server-outlined", "ant-design:cloud-upload-outlined", "ant-design:cloud-outlined",
	// "ant-design:cloud-download-outlined", "ant-design:cloud-sync-outlined", "ant-design:video-outlined", "ant-design:notification-outlined", "ant-design:sound-outlined",
	// "ant-design:radarchart-outlined", "ant-design:qrcode-outlined", "ant-design:fund-outlined", "ant-design:image-outlined", "ant-design:mail-outlined",
	// "ant-design:table-outlined", "ant-design:idcard-outlined", "ant-design:creditcard-outlined", "ant-design:heart-outlined", "ant-design:block-outlined",
	// "ant-design:error-outlined", "ant-design:star-outlined", "ant-design:gold-outlined", "ant-design:heatmap-outlined", "ant-design:wifi-outlined",
	// "ant-design:attachment-outlined", "ant-design:edit-outlined", "ant-design:key-outlined", "ant-design:api-outlined", "ant-design:disconnect-outlined",
	// "ant-design:highlight-outlined", "ant-design:monitor-outlined", "ant-design:link-outlined", "ant-design:man-outlined", "ant-design:percentage-outlined",
	// "ant-design:pushpin-outlined", "ant-design:phone-outlined", "ant-design:shake-outlined", "ant-design:tag-outlined", "ant-design:wrench-outlined",
	// "ant-design:tags-outlined", "ant-design:scissor-outlined", "ant-design:mr-outlined", "ant-design:share-outlined", "ant-design:branches-outlined",
	// "ant-design:fork-outlined", "ant-design:shrink-outlined", "ant-design:arrawsalt-outlined", "ant-design:verticalright-outlined", "ant-design:verticalleft-outlined",
	// "ant-design:right-outlined", "ant-design:left-outlined", "ant-design:up-outlined", "ant-design:down-outlined", "ant-design:fullscreen-outlined",
	// "ant-design:fullscreen-exit-outlined", "ant-design:doubleleft-outlined", "ant-design:doubleright-outlined", "ant-design:arrowright-outlined",
	// "ant-design:arrowup-outlined", "ant-design:arrowleft-outlined", "ant-design:arrowdown-outlined", "ant-design:upload-outlined",
	// "ant-design:colum-height-outlined", "ant-design:vertical-align-botto-outlined", "ant-design:vertical-align-middl-outlined", "ant-design:totop-outlined",
	// "ant-design:vertical-align-top-outlined", "ant-design:download-outlined", "ant-design:sort-descending-outlined", "ant-design:sort-ascending-outlined",
	// "ant-design:fall-outlined", "ant-design:swap-outlined", "ant-design:stock-outlined", "ant-design:rise-outlined", "ant-design:indent-outlined",
	// "ant-design:outdent-outlined", "ant-design:menu-outlined", "ant-design:unorderedlist-outlined", "ant-design:orderedlist-outlined", "ant-design:align-right-outlined",
	// "ant-design:align-center-outlined", "ant-design:align-left-outlined", "ant-design:pic-center-outlined", "ant-design:pic-right-outlined", "ant-design:pic-left-outlined",
	// "ant-design:bold-outlined", "ant-design:font-colors-outlined", "ant-design:exclaimination-outlined", "ant-design:font-size-outlined", "ant-design:check-circle-outlined",
	// "ant-design:infomation-outlined", "ant-design:CI-outlined", "ant-design:line-height-outlined", "ant-design:Dollar-outlined", "ant-design:strikethrough-outlined",
	// "ant-design:compass-outlined", "ant-design:underline-outlined", "ant-design:close-circle-outlined", "ant-design:number-outlined", "ant-design:frown-outlined",
	// "ant-design:italic-outlined", "ant-design:info-circle-outlined", "ant-design:code-outlined", "ant-design:left-circle-outlined", "ant-design:column-width-outlined",
	// "ant-design:down-circle-outlined", "ant-design:check-outlined", "ant-design:EURO-outlined", "ant-design:ellipsis-outlined", "ant-design:copyright-outlined",
	// "ant-design:dash-outlined", "ant-design:minus-circle-outlined", "ant-design:close-outlined", "ant-design:meh-outlined", "ant-design:enter-outlined",
	// "ant-design:plus-circle-outlined", "ant-design:line-outlined", "ant-design:play-circle-outlined", "ant-design:minus-outlined", "ant-design:question-circle-outlined",
	// "ant-design:question-outlined", "ant-design:Pound-outlined", "ant-design:rollback-outlined", "ant-design:right-circle-outlined", "ant-design:small-dash-outlined",
	// "ant-design:smile-outlined", "ant-design:pause-outlined", "ant-design:trademark-outlined", "ant-design:bg-colors-outlined", "ant-design:time-circle-outlined",
	// "ant-design:crown-outlined", "ant-design:timeout-outlined", "ant-design:drag-outlined", "ant-design:earth-outlined", "ant-design:desktop-outlined",
	// "ant-design:YUAN-outlined", "ant-design:gift-outlined", "ant-design:up-circle-outlined", "ant-design:stop-outlined", "ant-design:warning-circle-outlined",
	// "ant-design:fire-outlined", "ant-design:sync-outlined", "ant-design:thunderbolt-outlined", "ant-design:transaction-outlined", "ant-design:alipay-outlined",
	// "ant-design:undo-outlined", "ant-design:taobao-outlined", "ant-design:redo-outlined", "ant-design:wechat-fill-outlined", "ant-design:reload-outlined",
	// "ant-design:comment-outlined", "ant-design:reloadtime-outlined", "ant-design:login-outlined", "ant-design:message-outlined", "ant-design:clear-outlined",
	// "ant-design:dashboard-outlined", "ant-design:issuesclose-outlined", "ant-design:poweroff-outlined", "ant-design:logout-outlined", "ant-design:piechart-outlined",
	// "ant-design:setting-outlined", "ant-design:eye-outlined", "ant-design:location-outlined", "ant-design:edit-square-outlined", "ant-design:export-outlined",
	// "ant-design:save-outlined", "ant-design:Import-outlined", "ant-design:appstore-outlined", "ant-design:close-square-outlined", "ant-design:down-square-outlined",
	// "ant-design:layout-outlined", "ant-design:left-square-outlined", "ant-design:play-square-outlined", "ant-design:control-outlined", "ant-design:codelibrary-outlined",
	// "ant-design:detail-outlined", "ant-design:minus-square-outlined", "ant-design:plus-square-outlined", "ant-design:right-square-outlined", "ant-design:project-outlined",
	// "ant-design:wallet-outlined", "ant-design:up-square-outlined", "ant-design:calculator-outlined", "ant-design:interation-outlined", "ant-design:check-square-outlined",
	// "ant-design:border-outlined", "ant-design:border-outer-outlined", "ant-design:border-top-outlined", "ant-design:border-bottom-outlined", "ant-design:border-left-outlined",
	// "ant-design:border-right-outlined", "ant-design:border-inner-outlined", "ant-design:border-verticle-outlined", "ant-design:border-horizontal-outlined",
	// "ant-design:radius-bottomleft-outlined", "ant-design:radius-bottomright-outlined", "ant-design:radius-upleft-outlined", "ant-design:radius-upright-outlined",
	// "ant-design:radius-setting-outlined", "ant-design:adduser-outlined", "ant-design:deleteteam-outlined", "ant-design:deleteuser-outlined", "ant-design:addteam-outlined",
	// "ant-design:user-outlined", "ant-design:team-outlined", "ant-design:areachart-outlined", "ant-design:linechart-outlined", "ant-design:barchart-outlined",
	// "ant-design:pointmap-outlined", "ant-design:container-outlined", "ant-design:atom-outlined", "ant-design:zanwutupian-outlined", "ant-design:safetycertificate-outlined",
	// "ant-design:password-outlined", "ant-design:article-outlined", "ant-design:page-outlined", "ant-design:plugin-outlined", "ant-design:admin-outlined",
	// "ant-design:banner-outlined"
	Icon                interface{} `json:"icon"`
	Type                string      `json:"type"`                // 设置按钮类型，primary | ghost | dashed | link | text | default
	Size                string      `json:"size"`                // 设置按钮大小,large | middle | small | default
	WithLoading         bool        `json:"withLoading"`         // 是否具有loading，当action 的作用类型为ajax,submit时有效
	Fields              interface{} `json:"fields"`              // 行为表单字段
	ConfirmTitle        string      `json:"confirmTitle"`        // 确认标题
	ConfirmText         string      `json:"confirmText"`         // 确认文字描述
	ConfirmType         string      `json:"confirmType"`         // 确认类型
	OnlyOnIndex         bool        `json:"onlyOnIndex"`         // 只在列表页展示
	OnlyOnForm          bool        `json:"onlyOnForm"`          // 只在表单页展示
	OnlyOnDetail        bool        `json:"onlyOnDetail"`        // 只在详情页展示
	ShowOnIndex         bool        `json:"showOnIndex"`         // 在列表页展示
	ShowOnIndexTableRow bool        `json:"showOnIndexTableRow"` // 在列表页行展示
	ShowOnForm          bool        `json:"showOnForm"`          // 在表单页展示
	ShowOnFormExtra     bool        `json:"showOnFormExtra"`     // 在表单页扩展栏展示
	ShowOnDetail        bool        `json:"showOnDetail"`        // 在详情页展示
	ShowOnDetailExtra   bool        `json:"showOnDetailExtra"`   // 在详情页扩展栏展示
}

// 加载初始化数据
func (p *Action) New(ctx *quark.Context) interface{} {
	p.ActionType = "ajax"
	p.Type = "default"
	p.Size = "small"

	return p
}

// 初始化
func (p *Action) Init(ctx *quark.Context) interface{} {
	return p
}

// 执行行为句柄
func (p *Action) Handle(ctx *quark.Context, query *gorm.DB) error {
	return ctx.CJSONError("method not implemented")
}

// 行为key
func (p *Action) GetUriKey(action interface{}) string {
	uriKey := reflect.TypeOf(action).String()
	uriKeys := strings.Split(uriKey, ".")
	uriKey = stringy.New(uriKeys[1]).KebabCase("?", "").ToLower()

	return uriKey
}

// 获取名称
func (p *Action) GetName() interface{} {
	return p.Name
}

// 执行成功后刷新的组件
func (p *Action) GetReload() string {
	return p.Reload
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *Action) GetApiParams() []string {
	return p.ApiParams
}

// 执行行为的接口
func (p *Action) GetApi() string {
	return p.Api
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Action) GetActionType() string {
	return p.ActionType
}

// 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Action) GetSubmitForm() string {
	return p.SubmitForm
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Action) GetType() string {
	return p.Type
}

// 设置按钮大小,large | middle | small | default
func (p *Action) GetSize() string {
	return p.Size
}

// 是否具有loading，当action 的作用类型为ajax,submit时有效
func (p *Action) GetWithLoading() bool {
	return p.WithLoading
}

// 设置按钮的图标组件
func (p *Action) GetIcon() interface{} {
	return p.Icon
}

// 设置按钮是否为块级元素
func (p *Action) GetBlock() bool {
	return p.Block
}

// 是否批量操作
func (p *Action) GetBatch() bool {
	return p.Batch
}

// 危险按钮
func (p *Action) GetDanger() bool {
	return p.Danger
}

// 禁用按钮
func (p *Action) GetDisabled() bool {
	return p.Disabled
}

func (p *Action) GetGhost() bool {
	return p.Ghost
}

// 行为表单字段
func (p *Action) GetFields() interface{} {
	return p.Fields
}

// 确认标题
func (p *Action) GetConfirmTitle() string {
	return p.ConfirmTitle
}

// 确认文字
func (p *Action) GetConfirmText() string {
	return p.ConfirmText
}

// 确认类型
func (p *Action) GetConfirmType() string {
	return p.ConfirmType
}

// 设置名称
func (p *Action) SetName(name string) {
	p.Name = name
}

// 设置执行成功后刷新的组件
func (p *Action) SetReload(componentKey string) {
	p.Reload = componentKey
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *Action) SetApiParams(apiParams []string) {
	p.ApiParams = apiParams
}

// 执行行为的接口
func (p *Action) SetApi(api string) {
	p.Api = api
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Action) SetActionType(actionType string) {
	p.ActionType = actionType
}

// 当 action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Action) SetSubmitForm(submitForm string) {
	p.SubmitForm = submitForm
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Action) SetType(buttonType string) {
	p.Type = buttonType
}

// 设置按钮大小,large | middle | small | default
func (p *Action) SetSize(size string) {
	p.Size = size
}

// 是否具有loading，当action 的作用类型为ajax,submit时有效
func (p *Action) SetWithLoading(loading bool) {
	p.WithLoading = loading
}

// 配置按钮是否为块级元素
func (p *Action) SetBlock(block bool) {
	p.Block = block
}

// 配置批量操作
func (p *Action) SetBatch(batch bool) {
	p.Batch = batch
}

// 危险按钮
func (p *Action) SetDanger(danger bool) {
	p.Danger = danger
}

// 配置按钮是否禁用
func (p *Action) SetDisabled(disabled bool) {
	p.Disabled = disabled
}

// 配置按钮是否为幽灵按钮
func (p *Action) SetGhost(ghost bool) {
	p.Ghost = ghost
}

// 设置按钮的图标组件
func (p *Action) SetIcon(icon string) {
	p.Icon = icon
}

// 行为表单字段
func (p *Action) SetFields(fields interface{}) {
	p.Fields = fields
}

// 确认标题
func (p *Action) SetConfirmTitle(confirmTitle string) {
	p.ConfirmTitle = confirmTitle
}

// 确认文字
func (p *Action) SetConfirmText(confirmText string) {
	p.ConfirmText = confirmText
}

// 确认类型
func (p *Action) SetConfirmType(confirmType string) {
	p.ConfirmType = confirmType
}

// 设置行为前的确认操作
func (p *Action) WithConfirm(title string, text string, confirmType string) {

	p.ConfirmTitle = title
	p.ConfirmText = text
	p.ConfirmType = confirmType
}

// 只在列表页展示
func (p *Action) SetOnlyOnIndex(value bool) {
	p.OnlyOnIndex = value
	p.ShowOnIndex = value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value
}

// 除了列表页外展示
func (p *Action) SetExceptOnIndex() {
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
	p.ShowOnIndex = false
}

// 只在表单页展示
func (p *Action) SetOnlyOnForm(value bool) {
	p.ShowOnForm = value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value
}

// 除了表单页外展示
func (p *Action) SetExceptOnForm() {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = false
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
}

// 只在表单页右上角自定义区域展示
func (p *Action) SetOnlyOnFormExtra(value bool) {
	p.ShowOnForm = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value
}

// 除了表单页右上角自定义区域外展示
func (p *Action) SetExceptOnFormExtra() {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = false
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
}

// 只在详情页展示
func (p *Action) SetOnlyOnDetail(value bool) {
	p.OnlyOnDetail = value
	p.ShowOnDetail = value
	p.ShowOnIndex = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetailExtra = !value
}

// 除了详情页外展示
func (p *Action) SetExceptOnDetail() {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetailExtra = true
}

// 只在详情页右上角自定义区域展示
func (p *Action) SetOnlyOnDetailExtra(value bool) {
	p.ShowOnForm = !value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnIndexTableRow = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = value
}

// 除了详情页右上角自定义区域外展示
func (p *Action) SetExceptOnDetailExtra() {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnIndexTableRow = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = false
}

// 在表格行内展示
func (p *Action) SetOnlyOnIndexTableRow(value bool) {
	p.ShowOnIndexTableRow = value
	p.ShowOnIndex = !value
	p.ShowOnDetail = !value
	p.ShowOnForm = !value
	p.ShowOnFormExtra = !value
	p.ShowOnDetail = !value
	p.ShowOnDetailExtra = !value
}

// 除了表格行内外展示
func (p *Action) SetExceptOnIndexTableRow() {
	p.ShowOnIndexTableRow = false
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnForm = true
	p.ShowOnFormExtra = true
	p.ShowOnDetail = true
	p.ShowOnDetailExtra = true
}

// 在列表页展示
func (p *Action) SetShowOnIndex() {
	p.ShowOnIndex = true
}

// 在表单页展示
func (p *Action) SetShowOnForm() {
	p.ShowOnForm = true
}

// 在表单页右上角自定义区域展示
func (p *Action) SetShowOnFormExtra() {
	p.ShowOnFormExtra = true
}

// 在详情页展示
func (p *Action) SetShowOnDetail() {
	p.ShowOnDetail = true
}

// 在详情页右上角自定义区域展示
func (p *Action) SetShowOnDetailExtra() {
	p.ShowOnDetailExtra = true
}

// 在表格行内展示
func (p *Action) SetShowOnIndexTableRow() {
	p.ShowOnIndexTableRow = true
}

// 判断是否在列表页展示
func (p *Action) ShownOnIndex() bool {
	if p.OnlyOnIndex {
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

// 判断是否在表单页展示
func (p *Action) ShownOnForm() bool {
	if p.OnlyOnForm {
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

// 判断是否在详情页展示
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

// 判断是否在表格行内展示
func (p *Action) ShownOnIndexTableRow() bool {
	return p.ShowOnIndexTableRow
}

// 判断是否在表单页右上角自定义区域展示
func (p *Action) ShownOnFormExtra() bool {
	return p.ShowOnFormExtra
}

// 判断是否在详情页右上角自定义区域展示
func (p *Action) ShownOnDetailExtra() bool {
	return p.ShowOnDetailExtra
}
