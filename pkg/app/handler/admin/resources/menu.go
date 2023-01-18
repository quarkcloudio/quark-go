package resources

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	models "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/lister"
	"github.com/quarkcms/quark-go/pkg/msg"
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
func (p *Menu) Fields(request *builder.Request) []interface{} {
	field := &builder.AdminField{}

	// 权限列表
	permissions, _ := (&models.Permission{}).List()

	// 菜单列表
	menus, _ := (&models.Menu{}).OrderedList()

	return []interface{}{
		field.Hidden("id", "ID"), // 列表读取且不展示的字段

		field.Hidden("pid", "PID").OnlyOnIndex(), // 列表读取且不展示的字段

		field.Text("name", "名称").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "名称必须填写",
				},
			),

		field.Text("guard_name", "GuardName").
			SetDefault("admin").
			OnlyOnForms(),

		field.Icon("icon", "图标").OnlyOnForms(),

		field.Radio("type", "渲染组件").
			SetOptions(map[interface{}]interface{}{
				"default": "无组件",
				"engine":  "引擎组件",
			}).SetDefault("engine"),

		field.Text("path", "路由").
			SetEditable(true).
			SetHelp("前端路由或后端api"),

		field.Select("pid", "父节点").
			SetOptions(menus).
			SetDefault(0).
			OnlyOnForms(),

		field.Number("sort", "排序").
			SetEditable(true).
			SetDefault(0),

		field.Select("permission_ids", "绑定权限").
			SetMode("tags").
			SetOptions(permissions).
			OnlyOnForms(),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Menu) Searches(request *builder.Request) []interface{} {
	return []interface{}{
		(&searches.Input{}).Init("name", "名称"),
		(&searches.Input{}).Init("path", "路由"),
		(&searches.Status{}).Init(),
	}
}

// 行为
func (p *Menu) Actions(request *builder.Request) []interface{} {
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
func (p *Menu) BeforeIndexShowing(request *builder.Request, list []map[string]interface{}) []interface{} {
	data := request.AllQuerys()
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
func (p *Menu) BeforeEditing(request *builder.Request, data map[string]interface{}) map[string]interface{} {
	id := request.Query("id", "")

	if id != "" {
		menus := []int{}

		db.Client.
			Model(&models.Permission{}).
			Where("menu_id = ?", id).
			Pluck("id", &menus)

		data["permission_ids"] = menus
	}

	return data
}

// 保存数据前回调
func (p *Menu) BeforeSaving(request *builder.Request, submitData map[string]interface{}) interface{} {

	// 暂时清理permission_ids
	delete(submitData, "permission_ids")

	return submitData
}

// 保存后回调
func (p *Menu) AfterSaved(request *builder.Request, model *gorm.DB) interface{} {
	var result *gorm.DB
	data := map[string]interface{}{}
	json.Unmarshal(request.Body(), &data)
	id := 0
	if request.IsCreating() {
		last := map[string]interface{}{}
		result = model.Order("id desc").First(&last) // 获取最后一条记录的id
		id = last["id"].(int)
	} else {
		id = int(data["id"].(float64))
		result = db.Client.
			Model(&models.Permission{}).
			Where("menu_id = ?", id).
			Update("menu_id", 0)
	}

	if data["permission_ids"] != nil {
		result = db.Client.
			Model(&models.Permission{}).
			Where("id In ?", data["permission_ids"]).
			Update("menu_id", id)
	}

	if result.Error != nil {
		return msg.Error(result.Error.Error(), "")
	}

	return msg.Success("操作成功！", strings.Replace("/index?api="+adminresource.IndexRoute, ":resource", request.Param("resource"), -1), "")
}
