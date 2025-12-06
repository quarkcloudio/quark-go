package actions

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v4/component/form/rule"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/quarkcloudio/quark-go/v4/template/resource"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
	"gorm.io/gorm"
)

type DataScopeAction struct {
	actions.ModalForm
}

// 数据权限，DataScope()
func DataScope(options ...interface{}) *DataScopeAction {
	action := &DataScopeAction{}

	// 文字
	action.Name = "数据权限"

	return action
}

// 初始化
func (p *DataScopeAction) Init(ctx *quark.Context) interface{} {

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 关闭时销毁 Modal 里的子元素
	p.DestroyOnClose = true

	// 类型
	p.Type = "link"

	// 关闭时销毁 Modal 里的子元素
	p.DestroyOnClose = true

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	// 行为接口接收的参数
	p.SetApiParams([]string{
		"id",
		"name",
	})

	return p
}

// 字段
func (p *DataScopeAction) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	departments, _ := service.NewDepartmentService().GetList()

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("name", "名称").SetDisabled(true),

		field.Select("data_scope", "数据范围").SetOptions(
			[]selectfield.Option{
				{
					Label: "全部数据权限",
					Value: 1,
				},
				{
					Label: "自定数据权限",
					Value: 2,
				},
				{
					Label: "本部门数据权限",
					Value: 3,
				},
				{
					Label: "本部门及以下数据权限",
					Value: 4,
				},
				{
					Label: "仅本人数据权限",
					Value: 5,
				},
			}).
			SetRules([]rule.Rule{
				rule.Required("请选择数据范围"),
			}).
			SetDefault(1),

		field.Dependency().
			SetWhen("data_scope", 2, func() interface{} {
				return []interface{}{
					field.Tree("department_ids", "数据权限").
						SetDefaultExpandAll(true).
						SetTreeData(departments, "pid", "id", "name"),
				}
			}),
	}
}

// 表单数据（异步获取）
func (p *DataScopeAction) Data(ctx *quark.Context) map[string]interface{} {
	id := ctx.Query("id")
	idInt, err := strconv.Atoi(id.(string))
	if err != nil {
		return nil
	}

	role, err := service.NewRoleService().GetInfoById(idInt)
	if err != nil {
		return nil
	}

	departmentIds, err := service.NewPermissionService().GetRoleDepartmentIds(idInt)
	if err != nil {
		return nil
	}

	return map[string]interface{}{
		"id":             role.Id,
		"name":           role.Name,
		"data_scope":     role.DataScope,
		"department_ids": departmentIds,
	}
}

// 执行行为句柄
func (p *DataScopeAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	type Form struct {
		Id            int   `form:"id" json:"id"`
		DataScope     int   `form:"data_scope" json:"data_scope"`
		DepartmentIds []int `form:"department_ids" json:"department_ids"`
	}
	var form Form
	err := ctx.Bind(&form)
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	// 更新角色数据权限
	err = service.NewRoleService().UpdateRoleDataScope(form.Id, form.DataScope, form.DepartmentIds)
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("操作成功")
}
