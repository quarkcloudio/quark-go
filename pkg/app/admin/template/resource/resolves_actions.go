package resource

import (
	"strings"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/action"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/drawer"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/dropdown"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/modal"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/space"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
)

// 列表行为
func (p *Template) IndexActions(ctx *builder.Context) interface{} {
	var items []interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取行为
	actions := template.Actions(ctx)

	// 解析行为
	for _, v := range actions {

		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.TemplateInit(ctx)

		// 初始化
		actionInstance.Init(ctx)

		// 判断是否在列表页展示
		if actionInstance.ShownOnIndex() {
			items = append(items, p.BuildAction(ctx, actionInstance))
		}
	}

	return (&space.Component{}).
		Init().
		SetBody(items)
}

// 表格行内行为
func (p *Template) IndexTableRowActions(ctx *builder.Context) interface{} {
	var items []interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取行为
	actions := template.Actions(ctx)

	// 解析行为
	for _, v := range actions {

		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.TemplateInit(ctx)

		// 初始化
		actionInstance.Init(ctx)

		// 判断是否在表格行内展示
		if actionInstance.ShownOnIndexTableRow() {
			items = append(items, p.BuildAction(ctx, actionInstance))
		}
	}

	return items
}

// 表格多选弹出层行为
func (p *Template) IndexTableAlertActions(ctx *builder.Context) interface{} {
	var items []interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取行为
	actions := template.Actions(ctx)

	// 解析行为
	for _, v := range actions {

		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.TemplateInit(ctx)

		// 初始化
		actionInstance.Init(ctx)

		// 判断是否在多选弹出层展示
		if actionInstance.ShownOnIndexTableAlert() {
			items = append(items, p.BuildAction(ctx, actionInstance))
		}
	}

	return items
}

// 表单页行为
func (p *Template) FormActions(ctx *builder.Context) []interface{} {
	var items []interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取行为
	actions := template.Actions(ctx)

	// 解析行为
	for _, v := range actions {

		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.TemplateInit(ctx)

		// 初始化
		actionInstance.Init(ctx)

		// 判断是否在表单页展示
		if actionInstance.ShownOnForm() {
			items = append(items, p.BuildAction(ctx, actionInstance))
		}
	}

	return items
}

// 表单页右上角自定义区域行为
func (p *Template) FormExtraActions(ctx *builder.Context) interface{} {
	var items []interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取行为
	actions := template.Actions(ctx)

	// 解析行为
	for _, v := range actions {

		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.TemplateInit(ctx)

		// 初始化
		actionInstance.Init(ctx)

		// 判断是否在表单页右上角自定义区域展示
		if actionInstance.ShownOnFormExtra() {
			items = append(items, p.BuildAction(ctx, actionInstance))
		}
	}

	return items
}

// 详情页行为
func (p *Template) DetailActions(ctx *builder.Context) []interface{} {
	var items []interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取行为
	actions := template.Actions(ctx)

	// 解析行为
	for _, v := range actions {

		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.TemplateInit(ctx)

		// 初始化
		actionInstance.Init(ctx)

		// 判断是否在详情页展示
		if actionInstance.ShownOnDetail() {
			items = append(items, p.BuildAction(ctx, actionInstance))
		}
	}

	return items
}

// 详情页右上角自定义区域行为
func (p *Template) DetailExtraActions(ctx *builder.Context) interface{} {
	var items []interface{}

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取行为
	actions := template.Actions(ctx)

	// 解析行为
	for _, v := range actions {

		actionInstance := v.(types.Actioner)

		// 初始化模版
		actionInstance.TemplateInit(ctx)

		// 初始化
		actionInstance.Init(ctx)

		// 判断是否在详情页右上角自定义区域展示
		if actionInstance.ShownOnDetailExtra() {
			items = append(items, p.BuildAction(ctx, actionInstance))
		}
	}

	return items
}

// 创建行为组件
func (p *Template) BuildAction(ctx *builder.Context, item interface{}) interface{} {
	actionInstance := item.(types.Actioner)

	// 行为名称
	name := actionInstance.GetName()

	// 是否携带Loading
	withLoading := actionInstance.GetWithLoading()

	// 行为执行完成后刷新的组件
	reload := actionInstance.GetReload()

	// uri唯一标识
	uriKey := actionInstance.GetUriKey(item)

	// 获取api
	api := actionInstance.GetApi()

	// 获取api替换参数
	params := actionInstance.GetApiParams()
	if api == "" {
		api = p.BuildActionApi(ctx, params, uriKey)
	}

	// 行为类型
	actionType := actionInstance.GetActionType()

	// 按钮类型
	buttonType := actionInstance.GetType()

	// 按钮大小
	size := actionInstance.GetSize()

	// 按钮图标
	icon := actionInstance.GetIcon()

	// 确认操作标题
	confirmTitle := actionInstance.GetConfirmTitle()

	// 确认操作提示信息
	confirmText := actionInstance.GetConfirmText()

	// 确认操作类型
	confirmType := actionInstance.GetConfirmType()

	// 构建行为
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
		getAction = getAction.SetIcon(icon)
	}

	switch actionType {
	case "link":
		linkActioner := item.(types.Linker)

		// 是否显示箭头图标
		href := linkActioner.GetHref(ctx)

		// 相当于 a 链接的 target 属性，href 存在时生效
		target := linkActioner.GetTarget(ctx)

		// 设置跳转链接
		getAction = getAction.SetLink(href, target)
	case "modal":
		modalActioner := item.(types.Modaler)

		// 宽度
		modalWidth := modalActioner.GetWidth()

		// 关闭时销毁 Modal 里的子元素
		modalDestroyOnClose := modalActioner.GetDestroyOnClose()

		// 内容
		modalBody := modalActioner.GetBody(ctx)

		// 弹窗行为
		modalActions := modalActioner.GetActions(ctx)

		// 设置弹窗
		getAction = getAction.SetModal(func(modal *modal.Component) interface{} {
			return modal.
				SetTitle(name).
				SetWidth(modalWidth).
				SetBody(modalBody).
				SetActions(modalActions).
				SetDestroyOnClose(modalDestroyOnClose)
		})
	case "drawer":
		drawerActioner := item.(types.Drawer)

		// 宽度
		drawerWidth := drawerActioner.GetWidth()

		// 关闭时销毁 Drawer 里的子元素
		drawerDestroyOnClose := drawerActioner.GetDestroyOnClose()

		// 内容
		drawerBody := drawerActioner.GetBody(ctx)

		// 弹窗行为
		drawerActions := drawerActioner.GetActions(ctx)

		// 构建弹窗
		getAction = getAction.SetDrawer(func(drawer *drawer.Component) interface{} {
			return drawer.
				SetTitle(name).
				SetWidth(drawerWidth).
				SetBody(drawerBody).
				SetActions(drawerActions).
				SetDestroyOnClose(drawerDestroyOnClose)
		})
	case "modalForm":
		modalFormerActioner := item.(types.ModalFormer)

		// 表单数据接口
		initApi := p.BuildFormInitApi(ctx, params, uriKey)

		// 字段
		modalFormFields := modalFormerActioner.Fields(ctx)

		// 解析表单组件内的字段
		formFields := p.FormFieldsParser(ctx, modalFormFields)

		// 表单初始数据
		modalFormData := map[string]interface{}{}

		// 宽度
		modalFormWidth := modalFormerActioner.GetWidth()

		// 关闭时销毁 Modal 里的子元素
		modalFormDestroyOnClose := modalFormerActioner.GetDestroyOnClose()

		// 构建表单组件
		formComponent := form.
			New().
			SetKey(uriKey, false).
			SetStyle(map[string]interface{}{
				"paddingTop": "24px",
			}).
			SetApi(api).
			SetInitApi(initApi).
			SetBody(formFields).
			SetInitialValues(modalFormData).
			SetLabelCol(map[string]interface{}{
				"span": 6,
			}).
			SetWrapperCol(map[string]interface{}{
				"span": 18,
			})

		// 取消按钮文案
		modalFormCancelText := modalFormerActioner.GetCancelText()

		// 提交按钮文案
		modalFormSubmitText := modalFormerActioner.GetSubmitText()

		// 弹窗行为
		modalFormlActions := []interface{}{
			(&action.Component{}).
				Init().
				SetLabel(modalFormCancelText).
				SetActionType("cancel"),

			(&action.Component{}).
				Init().
				SetLabel(modalFormSubmitText).
				SetWithLoading(true).
				SetReload(reload).
				SetActionType("submit").
				SetType("primary", false).
				SetSubmitForm(uriKey),
		}

		// 设置弹窗
		getAction = getAction.
			SetActionType("modal").
			SetModal(func(modal *modal.Component) interface{} {
				return modal.
					SetTitle(name).
					SetWidth(modalFormWidth).
					SetBody(formComponent).
					SetActions(modalFormlActions).
					SetDestroyOnClose(modalFormDestroyOnClose)
			})
	case "drawerForm":
		drawerFormerActioner := item.(types.DrawerFormer)

		// 表单数据接口
		initApi := p.BuildFormInitApi(ctx, params, uriKey)

		// 字段
		drawerFormFields := drawerFormerActioner.Fields(ctx)

		// 解析表单组件内的字段
		formFields := p.FormFieldsParser(ctx, drawerFormFields)

		// 表单初始数据
		drawerFormData := map[string]interface{}{}

		// 宽度
		drawerFormWidth := drawerFormerActioner.GetWidth()

		// 关闭时销毁 Modal 里的子元素
		drawerFormDestroyOnClose := drawerFormerActioner.GetDestroyOnClose()

		// 构建表单组件
		formComponent := form.
			New().
			SetKey(uriKey, false).
			SetApi(api).
			SetInitApi(initApi).
			SetBody(formFields).
			SetInitialValues(drawerFormData).
			SetLabelCol(map[string]interface{}{
				"span": 6,
			}).
			SetWrapperCol(map[string]interface{}{
				"span": 18,
			})

		// 取消按钮文案
		drawerFormCancelText := drawerFormerActioner.GetCancelText()

		// 提交按钮文案
		drawerFormSubmitText := drawerFormerActioner.GetSubmitText()

		// 弹窗行为
		drawerFormlActions := []interface{}{
			(&action.Component{}).
				Init().
				SetLabel(drawerFormCancelText).
				SetActionType("cancel"),

			(&action.Component{}).
				Init().
				SetLabel(drawerFormSubmitText).
				SetWithLoading(true).
				SetReload(reload).
				SetActionType("submit").
				SetType("primary", false).
				SetSubmitForm(uriKey),
		}

		// 设置弹窗
		getAction = getAction.
			SetActionType("drawer").
			SetDrawer(func(drawer *drawer.Component) interface{} {
				return drawer.
					SetTitle(name).
					SetWidth(drawerFormWidth).
					SetBody(formComponent).
					SetActions(drawerFormlActions).
					SetDestroyOnClose(drawerFormDestroyOnClose)
			})
	case "dropdown":
		dropdownActioner := item.(types.Dropdowner)

		// 获取下拉菜单
		overlay := dropdownActioner.GetMenu(ctx)

		// 下拉根元素的样式
		overlayStyle := dropdownActioner.GetOverlayStyle()

		// 菜单弹出位置：bottomLeft bottomCenter bottomRight topLeft topCenter topRight
		placement := dropdownActioner.GetPlacement()

		// 触发下拉的行为, 移动端不支持 hover,Array<click|hover|contextMenu>
		trigger := dropdownActioner.GetTrigger()

		// 是否显示箭头图标
		arrow := dropdownActioner.GetArrow()

		// 构建行为
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
			getAction = getAction.SetIcon(icon)
		}

		return getAction
	}

	if confirmTitle != "" {
		getAction = getAction.SetWithConfirm(confirmTitle, confirmText, confirmType)
	}

	return getAction
}

// 创建行为接口
func (p *Template) BuildActionApi(ctx *builder.Context, params []string, uriKey string) string {
	var (
		paramsUri = ""
		api       = ""
	)

	for _, v := range params {
		paramsUri = paramsUri + v + "=${" + v + "}&"
	}

	// 列表页接口
	api = strings.Replace(ctx.Path(), "/index", "/action/"+uriKey, -1)

	// 创建页接口
	api = strings.Replace(api, "/create", "/action/"+uriKey, -1)

	// 编辑页接口
	api = strings.Replace(api, "/edit", "/action/"+uriKey, -1)

	// 详情页接口
	api = strings.Replace(api, "/detail", "/action/"+uriKey, -1)

	// 追加参数
	if paramsUri != "" {
		api = api + "?" + paramsUri
	}

	return api
}

// 创建表单初始化数据接口
func (p *Template) BuildFormInitApi(ctx *builder.Context, params []string, uriKey string) string {
	var (
		paramsUri = ""
		api       = ""
	)

	for _, v := range params {
		paramsUri = paramsUri + v + "=${" + v + "}&"
	}

	// 列表页接口
	api = strings.Replace(ctx.Path(), "/index", "/action/"+uriKey+"/values", -1)

	// 创建页接口
	api = strings.Replace(api, "/create", "/action/"+uriKey+"/values", -1)

	// 编辑页接口
	api = strings.Replace(api, "/edit", "/action/"+uriKey+"/values", -1)

	// 详情页接口
	api = strings.Replace(api, "/detail", "/action/"+uriKey+"/values", -1)

	// 追加参数
	if paramsUri != "" {
		api = api + "?" + paramsUri
	}

	return api
}
