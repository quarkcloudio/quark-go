package resources

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/app/actions"
	"github.com/quarkcloudio/quark-go/v4/app/searches"
	"github.com/quarkcloudio/quark-go/v4/component/form/rule"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/quarkcloudio/quark-go/v4/template/resource"
	"gorm.io/gorm"
)

type Role struct {
	resource.Template
}

// 初始化
func (p *Role) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "角色"

	// 模型
	p.Model = &model.Role{}

	// 分页
	p.PageSize = 10

	// 默认排序
	p.IndexQueryOrder = "id asc"

	return p
}

// 字段
func (p *Role) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}
	treeData, _ := service.NewMenuService().GetList()

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("name", "名称").
			SetRules([]rule.Rule{
				rule.Required("名称必须填写"),
			}).
			SetCreationRules([]rule.Rule{
				rule.Unique("roles", "name", "名称已存在"),
			}).
			SetUpdateRules([]rule.Rule{
				rule.Unique("roles", "name", "{id}", "名称已存在"),
			}),
		field.Text("guard_name", "守卫").
			SetDefault("admin"),
		field.Tree("menu_ids", "权限").
			SetTreeData(treeData, "pid", "id", "name").
			OnlyOnForms(),
		field.Datetime("created_at", "创建时间").
			OnlyOnIndex(),
		field.Datetime("updated_at", "更新时间").
			OnlyOnIndex(),
		field.Switch("status", "状态").
			SetRules([]rule.Rule{
				rule.Required("请选择状态"),
			}).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Role) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
	}
}

// 行为
func (p *Role) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.CreateLink(),
		actions.BatchDeleteRole(),
		actions.DataScope(),
		actions.EditLink(),
		actions.DeleteRole(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

// 编辑页面显示前回调
func (p *Role) BeforeEditing(ctx *quark.Context, data map[string]interface{}) map[string]interface{} {
	id := ctx.Query("id")
	idInt, err := strconv.Atoi(id.(string))
	if err == nil {
		menus, _ := service.NewCasbinService().GetRoleMenus(idInt)
		ids := []int{}
		for _, v := range menus {
			ids = append(ids, v.Id)
		}
		data["menu_ids"] = ids
	}
	return data
}

// 保存后回调
func (p *Role) AfterSaved(ctx *quark.Context, id int, data map[string]interface{}, result *gorm.DB) (err error) {
	if data["menu_ids"] != nil {
		if menuIds, ok := data["menu_ids"].([]interface{}); ok {
			ids := []int{}
			for _, v := range menuIds {
				menuId := int(v.(float64))
				ids = append(ids, menuId)
			}
			err = service.NewCasbinService().AddMenuAndPermissionToRole(id, ids)
			if err != nil {
				return err
			}
		}
	}
	return err
}
