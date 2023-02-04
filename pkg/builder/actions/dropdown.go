package actions

import (
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/action"
	"github.com/quarkcms/quark-go/pkg/component/admin/menu"
)

type Dropdown struct {
	Action
	Arrow        bool                   `json:"arrow"`
	Placement    string                 `json:"placement"`
	Trigger      []string               `json:"trigger"`
	OverlayStyle map[string]interface{} `json:"overlayStyle"`
	Actions      []interface{}          `json:"actions"`
}

// 初始化
func (p *Dropdown) ParentInit() interface{} {
	p.ActionType = "dropdown"
	p.Placement = "bottomLeft"
	p.Trigger = append(p.Trigger, "hover")

	return p
}

// 是否显示箭头图标
func (p *Dropdown) GetArrow() bool {
	return p.Arrow
}

// 菜单弹出位置：bottomLeft bottomCenter bottomRight topLeft topCenter topRight
func (p *Dropdown) GetPlacement() string {
	return p.Placement
}

// 触发下拉的行为, 移动端不支持 hover,Array<click|hover|contextMenu>
func (p *Dropdown) GetTrigger() []string {
	return p.Trigger
}

// 下拉根元素的样式
func (p *Dropdown) GetOverlayStyle() map[string]interface{} {
	return p.OverlayStyle
}

// 菜单
func (p *Dropdown) GetMenu(request *builder.Request, templateInstance interface{}) interface{} {
	actions := p.GetActions()
	items := []interface{}{}

	for _, v := range actions {
		action := p.buildAction(request, v, templateInstance)
		items = append(items, action)
	}

	return (&menu.Component{}).Init().SetItems(items)
}

//创建行为组件
func (p *Dropdown) buildAction(request *builder.Request, item interface{}, templateInstance interface{}) interface{} {
	name := item.(interface{ GetName() string }).GetName()
	withLoading := item.(interface{ GetWithLoading() bool }).GetWithLoading()
	reload := item.(interface{ GetReload() string }).GetReload()

	// uri唯一标识
	uriKey := item.(interface {
		GetUriKey(interface{}) string
	}).GetUriKey(item)

	// 获取api
	api := item.(interface {
		GetApi(*builder.Request) string
	}).GetApi(request)

	// 获取api替换参数
	params := item.(interface {
		GetApiParams() []string
	}).GetApiParams()

	if api == "" {
		api = p.buildActionApi(request, params, uriKey)
	}

	actionType := item.(interface{ GetActionType() string }).GetActionType()
	buttonType := item.(interface{ GetType() string }).GetType()
	size := item.(interface{ GetSize() string }).GetSize()
	icon := item.(interface{ GetIcon() string }).GetIcon()
	confirmTitle := item.(interface{ GetConfirmTitle() string }).GetConfirmTitle()
	confirmText := item.(interface{ GetConfirmText() string }).GetConfirmText()
	confirmType := item.(interface{ GetConfirmType() string }).GetConfirmType()

	getAction := (&menu.Item{}).Init().
		Init().
		SetLabel(name).
		SetWithLoading(withLoading).
		SetReload(reload).
		SetApi(api).
		SetActionType(actionType).
		SetType(buttonType, false).
		SetSize(size)

	if icon != "" {
		getAction = getAction.
			SetIcon(icon)
	}

	switch actionType {
	case "link":
		href := item.(interface {
			GetHref(request *builder.Request) string
		}).GetHref(request)
		target := item.(interface {
			GetTarget(request *builder.Request) string
		}).GetTarget(request)

		getAction = getAction.
			SetLink(href, target).
			SetStyle(map[string]interface{}{
				"color": "#1890ff",
			})
	case "modal":
		formWidth := item.(interface {
			GetWidth() int
		}).GetWidth()

		formDestroyOnClose := item.(interface {
			GetDestroyOnClose() bool
		}).GetDestroyOnClose()

		formBody := item.(interface {
			GetBody(request *builder.Request, templateInstance interface{}) interface{}
		}).GetBody(request, templateInstance)

		formActions := item.(interface {
			GetActions(request *builder.Request, templateInstance interface{}) []interface{}
		}).GetActions(request, templateInstance)

		getAction = getAction.SetModal(func(modal *action.Modal) interface{} {
			return modal.
				SetTitle(name).
				SetWidth(formWidth).
				SetBody(formBody).
				SetActions(formActions).
				SetDestroyOnClose(formDestroyOnClose)
		})
	case "drawer":
		formWidth := item.(interface {
			GetWidth() int
		}).GetWidth()

		formDestroyOnClose := item.(interface {
			GetDestroyOnClose() bool
		}).GetDestroyOnClose()

		formBody := item.(interface {
			GetBody(request *builder.Request, templateInstance interface{}) interface{}
		}).GetBody(request, templateInstance)

		formActions := item.(interface {
			GetActions(request *builder.Request, templateInstance interface{}) []interface{}
		}).GetActions(request, templateInstance)

		getAction = getAction.SetDrawer(func(drawer *action.Drawer) interface{} {
			return drawer.
				SetTitle(name).
				SetWidth(formWidth).
				SetBody(formBody).
				SetActions(formActions).
				SetDestroyOnClose(formDestroyOnClose)
		})
	}

	if confirmTitle != "" {
		getAction = getAction.
			SetWithConfirm(confirmTitle, confirmText, confirmType)
	}

	return getAction
}

// 下拉菜单行为
func (p *Dropdown) SetActions(actions []interface{}) *Dropdown {
	p.Actions = actions

	return p
}

// 获取下拉菜单行为
func (p *Dropdown) GetActions() []interface{} {
	return p.Actions
}

//创建行为接口
func (p *Dropdown) buildActionApi(request *builder.Request, params []string, uriKey string) string {
	paramsUri := ""
	for _, v := range params {
		paramsUri = paramsUri + v + "=${" + v + "}&"
	}

	api := strings.Replace(request.Path(), "/index", "/action/"+uriKey, -1)
	if paramsUri != "" {
		api = api + "?" + paramsUri
	}

	return api
}
