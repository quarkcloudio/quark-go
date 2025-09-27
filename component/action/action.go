package action

import (
	"github.com/quarkcloudio/quark-go/v4/component/component"
	"github.com/quarkcloudio/quark-go/v4/component/drawer"
	"github.com/quarkcloudio/quark-go/v4/component/modal"
)

type Component struct {
	component.Element
	Label             interface{} `json:"label"`
	Block             bool        `json:"block"`
	Batch             bool        `json:"batch"`
	Danger            bool        `json:"danger"`
	Disabled          bool        `json:"disabled"`
	Ghost             bool        `json:"ghost"`
	Icon              interface{} `json:"icon"`
	Shape             string      `json:"shape"`
	Size              string      `json:"size"`
	Type              string      `json:"type"`
	ActionType        string      `json:"actionType"`
	SubmitForm        any         `json:"submitForm"`
	Href              string      `json:"href"`
	Target            string      `json:"target"`
	Modal             interface{} `json:"modal"`
	Drawer            interface{} `json:"drawer"`
	CheckedChildren   interface{} `json:"checkedChildren,omitempty"`   // 选中时的内容
	UnCheckedChildren interface{} `json:"unCheckedChildren,omitempty"` // 自定义的选择框后缀图标
	FieldName         interface{} `json:"fieldName,omitempty"`         // 字段名称
	FieldValue        interface{} `json:"fieldValue,omitempty"`        // 字段值
	ConfirmTitle      string      `json:"confirmTitle"`
	ConfirmText       string      `json:"confirmText"`
	ConfirmType       string      `json:"confirmType"`
	Api               string      `json:"api"`
	Reload            string      `json:"reload"`
	WithLoading       bool        `json:"withLoading"`
}

// 初始化组件
func New() *Component {

	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "action"
	p.Size = "default"
	p.Type = "default"
	p.SetKey("action", component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 设置按钮文字
func (p *Component) SetLabel(label interface{}) *Component {
	p.Label = label

	return p
}

// 将按钮宽度调整为其父宽度的选项
func (p *Component) SetBlock(block bool) *Component {
	p.Block = block

	return p
}

// 设置批量操作
func (p *Component) SetBatch(batch bool) *Component {
	p.Batch = batch

	return p
}

// 设置危险按钮
func (p *Component) SetDanger(danger bool) *Component {
	p.Danger = danger

	return p
}

// 按钮失效状态
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 幽灵属性，使按钮背景透明
func (p *Component) SetGhost(ghost bool) *Component {
	p.Ghost = ghost

	return p
}

// 设置按钮图标
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
func (p *Component) SetIcon(icon interface{}) *Component {
	if value, ok := icon.(string); ok {
		p.Icon = value
	}
	if value, ok := icon.([]string); ok {
		p.Icon = []string{value[0], value[1]}
	}

	return p
}

// 设置按钮形状，可选值为 circle、 round 或者不设
func (p *Component) SetShape(shape string) *Component {
	p.Shape = shape

	return p
}

// 设置按钮类型，primary | ghost | dashed | link | text | default
func (p *Component) SetType(buttonType string, danger bool) *Component {
	p.Type = buttonType
	p.Danger = danger

	return p
}

// 设置按钮大小，large | middle | small | default
func (p *Component) SetSize(size string) *Component {
	p.Size = size

	return p
}

// 【必填】这是 action 最核心的配置，来指定该 action 的作用类型，支持：ajax、link、url、drawer、dialog、confirm、cancel、prev、next、copy、close。
func (p *Component) SetActionType(actionType string) *Component {
	p.ActionType = actionType

	return p
}

// 当action 的作用类型为submit的时候，可以指定提交哪个表格，submitForm为提交表单的key值，为空时提交当前表单
func (p *Component) SetSubmitForm(formKey string) *Component {
	p.SubmitForm = formKey

	return p
}

// 点击跳转的地址，指定此属性 button 的行为和 a 链接一致
func (p *Component) SetHref(href string) *Component {
	p.Href = href

	return p
}

// 相当于 a 链接的 target 属性，href 存在时生效
func (p *Component) SetTarget(target string) *Component {
	p.Target = target

	return p
}

// 设置跳转链接
func (p *Component) SetLink(href string, target string) *Component {
	p.SetHref(href)
	p.SetTarget(target)
	p.ActionType = "link"

	return p
}

// 弹窗
func (p *Component) SetModal(callback interface{}) *Component {
	component := (&modal.Component{}).Init()
	getCallback := callback.(func(modal *modal.Component) interface{})

	p.Modal = getCallback(component)

	return p
}

// 抽屉
func (p *Component) SetDrawer(callback interface{}) *Component {
	component := (&drawer.Component{}).Init()
	getCallback := callback.(func(drawer *drawer.Component) interface{})

	p.Drawer = getCallback(component)

	return p
}

// 选中时的内容
func (p *Component) SetCheckedChildren(checkedChildren interface{}) *Component {
	p.CheckedChildren = checkedChildren

	return p
}

// 未选中时的内容
func (p *Component) SetUnCheckedChildren(unCheckedChildren interface{}) *Component {
	p.UnCheckedChildren = unCheckedChildren

	return p
}

// 获取字段名称
func (p *Component) SetFieldName(fieldName interface{}) *Component {
	p.FieldName = fieldName

	return p
}

// 获取字段值
func (p *Component) SetFieldValue(fieldValue interface{}) *Component {
	p.FieldValue = fieldValue

	return p
}

// 设置行为前的确认操作
func (p *Component) SetWithConfirm(title string, text string, confirmType string) *Component {
	p.ConfirmTitle = title
	p.ConfirmText = text
	p.ConfirmType = confirmType

	return p
}

// 执行行为的接口链接
func (p *Component) SetApi(api string) *Component {
	p.Api = api
	p.ActionType = "ajax"

	return p
}

// 执行成功后刷新的组件
func (p *Component) SetReload(reload string) *Component {
	p.Reload = reload

	return p
}

// 是否具有loading
func (p *Component) SetWithLoading(loading bool) *Component {
	p.WithLoading = loading

	return p
}
