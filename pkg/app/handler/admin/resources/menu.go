package resources

import (
	"strconv"
	"strings"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	models "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/radio"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/lister"
	"gorm.io/gorm"
)

type Menu struct {
	adminresource.Template
}

// 初始化
func (p *Menu) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "菜单"

	// 模型
	p.Model = &models.Menu{}

	// 分页
	p.PerPage = false

	// 默认排序
	p.IndexOrder = "sort asc"

	return p
}

// 字段
func (p *Menu) Fields(ctx *builder.Context) []interface{} {
	field := &adminresource.Field{}

	// 权限列表
	permissions, _ := (&models.Permission{}).List()

	// 菜单列表
	menus, _ := (&models.Menu{}).TreeSelect(true)

	return []interface{}{
		field.Hidden("id", "ID"), // 列表读取且不展示的字段

		field.Hidden("pid", "PID").OnlyOnIndex(), // 列表读取且不展示的字段

		field.Text("name", "名称").
			SetRules([]*rule.Rule{
				rule.Required(true, "名称必须填写"),
			}),

		field.Text("guard_name", "GuardName").
			SetDefault("admin").
			OnlyOnForms(),

		field.Icon("icon", "图标").OnlyOnForms(),

		field.TreeSelect("pid", "父节点").
			SetData(menus).
			SetDefault(0).
			OnlyOnForms(),

		field.Radio("type", "类型").
			SetOptions([]*radio.Option{
				{
					Value: 1,
					Label: "目录",
				},
				{
					Value: 2,
					Label: "菜单",
				},
				{
					Value: 3,
					Label: "按钮",
				},
			}).
			SetWhen(1, func() interface{} {
				return []interface{}{
					field.Text("path", "路由").
						SetEditable(true).
						SetHelp("前端路由"),
				}
			}).
			SetWhen(2, func() interface{} {
				return []interface{}{
					field.Switch("is_engine", "引擎组件").
						SetTrueValue("是").
						SetFalseValue("否").
						SetDefault(true),

					field.Switch("is_link", "外部链接").
						SetTrueValue("是").
						SetFalseValue("否").
						SetDefault(false),

					field.Text("path", "路由").
						SetEditable(true).
						SetHelp("前端路由或后端api").
						OnlyOnForms(),
				}
			}).
			SetWhen(3, func() interface{} {
				return []interface{}{
					field.Select("permission_ids", "绑定权限").
						SetMode("tags").
						SetOptions(permissions).
						OnlyOnForms(),
				}
			}).
			SetDefault(1),

		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Menu) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		(&searches.Input{}).Init("name", "名称"),
		(&searches.Input{}).Init("path", "路由"),
		(&searches.Status{}).Init(),
	}
}

// 行为
func (p *Menu) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		(&actions.CreateDrawer{}).Init(p.Title),
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Disable{}).Init("批量禁用"),
		(&actions.Enable{}).Init("批量启用"),
		(&actions.ChangeStatus{}).Init(),
		(&actions.EditDrawer{}).Init("编辑"),
		(&actions.Delete{}).Init("删除"),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}

// 列表页面显示前回调
func (p *Menu) BeforeIndexShowing(ctx *builder.Context, list []map[string]interface{}) []interface{} {
	data := ctx.AllQuerys()
	if search, ok := data["search"].(map[string]interface{}); ok == true && search != nil {
		result := []interface{}{}
		for _, v := range list {
			result = append(result, v)
		}

		return result
	}

	// 转换成树形表格
	tree, _ := lister.ListToTree(list, "id", "pid", "children", 0)

	return tree
}

// 编辑页面显示前回调
func (p *Menu) BeforeEditing(ctx *builder.Context, data map[string]interface{}) map[string]interface{} {
	id := ctx.Query("id", "")
	idInt, err := strconv.Atoi(id.(string))

	if id != "" && err == nil {
		permissionIds := []int{}
		permissions, err := (&models.CasbinRule{}).GetMenuPermissions(idInt)
		if err == nil {
			for _, v := range permissions {
				permissionIds = append(permissionIds, v.Id)
			}
		}
		data["permission_ids"] = permissionIds
	}

	return data
}

// 保存后回调
func (p *Menu) AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) error {
	if data["permission_ids"] != nil {
		err := (&models.CasbinRule{}).AddMenuPermission(id, data["permission_ids"])
		if err != nil {
			return ctx.JSONError(err.Error())
		}
	}

	if result.Error != nil {
		return ctx.JSONError(result.Error.Error())
	}

	return ctx.JSONOk("操作成功！", strings.Replace("/layout/index?api="+adminresource.IndexPath, ":resource", ctx.Param("resource"), -1))
}
