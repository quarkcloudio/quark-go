package actions

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
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
func (p *DataScopeAction) Init(ctx *builder.Context) interface{} {

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
func (p *DataScopeAction) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	departments, _ := (&model.Department{}).GetList()

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("name", "名称").SetDisabled(true),

		field.Select("data_scope", "数据范围").SetOptions(
			[]*selectfield.Option{
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
			SetRules([]*rule.Rule{
				rule.Required(true, "请选择数据范围"),
			}).
			SetDefault(1),

		field.Dependency().
			SetWhen("data_scope", 2, func() interface{} {
				return []interface{}{
					field.Tree("department_ids", "数据权限").
						SetDefaultExpandAll(true).
						SetData(departments, "pid", "id", "name"),
				}
			}),
	}
}

// 表单数据（异步获取）
func (p *DataScopeAction) Data(ctx *builder.Context) map[string]interface{} {
	id := ctx.Query("id")
	idInt, err := strconv.Atoi(id.(string))
	if err != nil {
		return nil
	}

	role, err := (&model.Role{}).GetInfoById(idInt)
	if err != nil {
		return nil
	}

	departmentIds, err := (&model.CasbinRule{}).GetRoleDepartmentIds(idInt)
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
func (p *DataScopeAction) Handle(ctx *builder.Context, query *gorm.DB) error {
	type Form struct {
		Id            int   `form:"id" json:"id"`
		DataScope     int   `form:"data_scope" json:"data_scope"`
		DepartmentIds []int `form:"department_ids" json:"department_ids"`
	}
	var form Form
	err := ctx.Bind(&form)
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	// 更新角色数据权限
	err = (&model.Role{}).UpdateRoleDataScope(form.Id, form.DataScope, form.DepartmentIds)
	if err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}

	return ctx.JSON(200, message.Success("操作成功"))
}
