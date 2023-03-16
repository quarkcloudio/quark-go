package adminresource

import (
	"strings"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/action"
	"github.com/quarkcms/quark-go/pkg/component/admin/drawer"
	"github.com/quarkcms/quark-go/pkg/component/admin/dropdown"
	"github.com/quarkcms/quark-go/pkg/component/admin/modal"
	"github.com/quarkcms/quark-go/pkg/component/admin/space"
)

// 列表行为
func (p *Template) IndexActions(ctx *builder.Context) interface{} {
	actions := ctx.Template.(interface {
		Actions(ctx *builder.Context) []interface{}
	}).Actions(ctx)

	var items []interface{}
	for _, v := range actions {
		shownOnIndex := v.(interface {
			ShownOnIndex() bool
		}).ShownOnIndex()

		if shownOnIndex {
			getAction := p.buildAction(ctx, v)
			items = append(items, getAction)
		}
	}

	return (&space.Component{}).Init().SetBody(items)
}

//表格行内行为
func (p *Template) IndexTableRowActions(ctx *builder.Context) interface{} {
	actions := ctx.Template.(interface {
		Actions(ctx *builder.Context) []interface{}
	}).Actions(ctx)

	var items []interface{}
	for _, v := range actions {
		shownOnIndexTableRow := v.(interface {
			ShownOnIndexTableRow() bool
		}).ShownOnIndexTableRow()

		if shownOnIndexTableRow {
			getAction := p.buildAction(ctx, v)
			items = append(items, getAction)
		}
	}

	return items
}

//表格多选弹出层行为
func (p *Template) IndexTableAlertActions(ctx *builder.Context) interface{} {
	actions := ctx.Template.(interface {
		Actions(ctx *builder.Context) []interface{}
	}).Actions(ctx)

	var items []interface{}
	for _, v := range actions {
		shownOnIndexTableAlert := v.(interface {
			ShownOnIndexTableAlert() bool
		}).ShownOnIndexTableAlert()

		if shownOnIndexTableAlert {
			getAction := p.buildAction(ctx, v)
			items = append(items, getAction)
		}
	}

	return items
}

//表单页行为
func (p *Template) FormActions(ctx *builder.Context) []interface{} {
	actions := ctx.Template.(interface {
		Actions(ctx *builder.Context) []interface{}
	}).Actions(ctx)

	var items []interface{}
	for _, v := range actions {
		shownOnForm := v.(interface {
			ShownOnForm() bool
		}).ShownOnForm()

		if shownOnForm {
			getAction := p.buildAction(ctx, v)
			items = append(items, getAction)
		}
	}

	return items
}

//表单页右上角自定义区域行为
func (p *Template) FormExtraActions(ctx *builder.Context) interface{} {
	actions := ctx.Template.(interface {
		Actions(ctx *builder.Context) []interface{}
	}).Actions(ctx)

	var items []interface{}
	for _, v := range actions {
		shownOnFormExtra := v.(interface {
			ShownOnFormExtra() bool
		}).ShownOnFormExtra()

		if shownOnFormExtra {
			getAction := p.buildAction(ctx, v)
			items = append(items, getAction)
		}
	}

	return items
}

//详情页行为
func (p *Template) DetailActions(ctx *builder.Context) []interface{} {
	actions := ctx.Template.(interface {
		Actions(ctx *builder.Context) []interface{}
	}).Actions(ctx)

	var items []interface{}
	for _, v := range actions {
		shownOnDetail := v.(interface {
			ShownOnDetail() bool
		}).ShownOnDetail()

		if shownOnDetail {
			getAction := p.buildAction(ctx, v)
			items = append(items, getAction)
		}
	}

	return items
}

//详情页右上角自定义区域行为
func (p *Template) DetailExtraActions(ctx *builder.Context) interface{} {
	actions := ctx.Template.(interface {
		Actions(ctx *builder.Context) []interface{}
	}).Actions(ctx)

	var items []interface{}
	for _, v := range actions {
		shownOnDetailExtra := v.(interface {
			ShownOnDetailExtra() bool
		}).ShownOnDetailExtra()

		if shownOnDetailExtra {
			getAction := p.buildAction(ctx, v)
			items = append(items, getAction)
		}
	}

	return items
}

//创建行为组件
func (p *Template) buildAction(ctx *builder.Context, item interface{}) interface{} {
	name := item.(interface{ GetName() string }).GetName()
	withLoading := item.(interface{ GetWithLoading() bool }).GetWithLoading()
	reload := item.(interface{ GetReload() string }).GetReload()

	// uri唯一标识
	uriKey := item.(interface {
		GetUriKey(interface{}) string
	}).GetUriKey(item)

	// 获取api
	api := item.(interface {
		GetApi(ctx *builder.Context) string
	}).GetApi(ctx)

	// 获取api替换参数
	params := item.(interface {
		GetApiParams() []string
	}).GetApiParams()

	if api == "" {
		api = p.buildActionApi(ctx, params, uriKey)
	}

	actionType := item.(interface{ GetActionType() string }).GetActionType()
	buttonType := item.(interface{ GetType() string }).GetType()
	size := item.(interface{ GetSize() string }).GetSize()
	icon := item.(interface{ GetIcon() string }).GetIcon()
	confirmTitle := item.(interface{ GetConfirmTitle() string }).GetConfirmTitle()
	confirmText := item.(interface{ GetConfirmText() string }).GetConfirmText()
	confirmType := item.(interface{ GetConfirmType() string }).GetConfirmType()

	if actionType == "dropdown" {
		overlay := item.(interface {
			GetMenu(ctx *builder.Context) interface{}
		}).GetMenu(ctx)

		overlayStyle := item.(interface {
			GetOverlayStyle() map[string]interface{}
		}).GetOverlayStyle()

		placement := item.(interface {
			GetPlacement() string
		}).GetPlacement()

		trigger := item.(interface {
			GetTrigger() []string
		}).GetTrigger()

		arrow := item.(interface {
			GetArrow() bool
		}).GetArrow()

		getAction := (&dropdown.Component{}).
			Init().
			SetLabel(name).
			SetMenu(overlay).
			SetOverlayStyle(overlayStyle).
			SetPlacement(placement).
			SetTrigger(trigger).
			SetArrow(arrow).
			SetType(buttonType, false).
			SetSize(size)

		if icon != "" {
			getAction = getAction.
				SetIcon(icon)
		}

		return getAction
	}

	getAction := (&action.Component{}).
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
			GetHref(ctx *builder.Context) string
		}).GetHref(ctx)
		target := item.(interface {
			GetTarget(ctx *builder.Context) string
		}).GetTarget(ctx)

		getAction = getAction.
			SetLink(href, target)
	case "modal":
		formWidth := item.(interface {
			GetWidth() int
		}).GetWidth()

		formDestroyOnClose := item.(interface {
			GetDestroyOnClose() bool
		}).GetDestroyOnClose()

		formBody := item.(interface {
			GetBody(ctx *builder.Context) interface{}
		}).GetBody(ctx)

		formActions := item.(interface {
			GetActions(ctx *builder.Context) []interface{}
		}).GetActions(ctx)

		getAction = getAction.SetModal(func(modal *modal.Component) interface{} {
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
			GetBody(ctx *builder.Context) interface{}
		}).GetBody(ctx)

		formActions := item.(interface {
			GetActions(ctx *builder.Context) []interface{}
		}).GetActions(ctx)

		getAction = getAction.SetDrawer(func(drawer *drawer.Component) interface{} {
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

//创建行为接口
func (p *Template) buildActionApi(ctx *builder.Context, params []string, uriKey string) string {
	paramsUri := ""

	for _, v := range params {
		paramsUri = paramsUri + v + "=${" + v + "}&"
	}

	// 自动构建列表页接口
	api := strings.Replace(ctx.Path(), "/index", "/action/"+uriKey, -1)

	// 自动构建创建页接口
	api = strings.Replace(api, "/create", "/action/"+uriKey, -1)

	// 自动构建编辑页接口
	api = strings.Replace(api, "/edit", "/action/"+uriKey, -1)

	// 自动构建详情页接口
	api = strings.Replace(api, "/detail", "/action/"+uriKey, -1)

	if paramsUri != "" {
		api = api + "?" + paramsUri
	}

	return api
}
