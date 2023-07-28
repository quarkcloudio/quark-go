package resources

import (
	"strconv"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type Role struct {
	resource.Template
}

// 初始化
func (p *Role) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "角色"

	// 模型
	p.Model = &model.Role{}

	// 分页
	p.PerPage = 10

	return p
}

// 字段
func (p *Role) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}
	treeData, _ := (&model.Menu{}).Tree()

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("name", "名称").
			SetRules([]*rule.Rule{
				rule.Required(true, "名称必须填写"),
			}),

		field.Text("guard_name", "GuardName").SetDefault("admin"),
		field.Tree("menu_ids", "权限").SetData(treeData).OnlyOnForms(),
		field.Datetime("created_at", "创建时间", func() interface{} {
			if p.Field["created_at"] == nil {
				return p.Field["created_at"]
			}

			return p.Field["created_at"].(time.Time).Format("2006-01-02 15:04:05")
		}).OnlyOnIndex(),
		field.Datetime("updated_at", "更新时间", func() interface{} {
			if p.Field["updated_at"] == nil {
				return p.Field["updated_at"]
			}

			return p.Field["updated_at"].(time.Time).Format("2006-01-02 15:04:05")
		}).OnlyOnIndex(),
	}
}

// 搜索
func (p *Role) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
	}
}

// 行为
func (p *Role) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.CreateLink(),
		actions.BatchDeleteRole(),
		actions.EditLink(),
		actions.DeleteRole(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

// 编辑页面显示前回调
func (p *Role) BeforeEditing(ctx *builder.Context, data map[string]interface{}) map[string]interface{} {
	id := ctx.Query("id")
	idInt, err := strconv.Atoi(id.(string))
	if err == nil {
		menus, _ := (&model.CasbinRule{}).GetRoleMenus(idInt)
		ids := []int{}
		for _, v := range menus {
			ids = append(ids, v.Id)
		}
		data["menu_ids"] = ids
	}

	return data
}

// 保存后回调
func (p *Role) AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) error {
	if data["menu_ids"] != nil {
		if menuIds, ok := data["menu_ids"].([]interface{}); ok {
			ids := []int{}
			for _, v := range menuIds {
				menuId := int(v.(float64))
				ids = append(ids, menuId)
			}

			err := (&model.CasbinRule{}).AddMenuAndPermissionToRole(id, ids)
			if err != nil {
				return ctx.JSON(200, message.Error(err.Error()))
			}
		}
	}

	return ctx.JSON(200, message.Success(
		"操作成功",
		strings.Replace("/layout/index?api="+resource.IndexPath, ":resource", ctx.Param("resource"), -1),
	))
}
