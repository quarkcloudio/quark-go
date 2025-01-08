package resources

import (
	"encoding/json"
	"strconv"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-go/v3/utils/hash"
	"gorm.io/gorm"
)

type User struct {
	resource.Template
}

// 初始化
func (p *User) Init(ctx *quark.Context) interface{} {

	// 列表
	departments, _ := service.NewDepartmentService().GetList()

	// 树形下拉框
	p.TableTreeBar.
		SetName("departmentIds").
		SetTreeData(departments, "pid", "id", "name")

	// 标题
	p.Title = "用户"

	// 模型
	p.Model = &model.User{}

	// 分页
	p.PageSize = 10

	// 是否具有导出功能
	p.WithExport = true

	return p
}

// 列表查询
func (p *User) IndexQuery(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	departmentIds := ctx.Query("departmentIds")
	if departmentIds == nil || departmentIds == "" {
		return query
	}

	var ids []int
	err := json.Unmarshal([]byte(departmentIds.(string)), &ids)
	if err != nil {
		return query
	}

	if len(ids) == 0 {
		return query
	}

	for _, v := range ids {
		childrenIds := service.NewDepartmentService().GetChildrenIds(v)
		ids = append(ids, childrenIds...)
	}

	return query.Where("department_id in ?", ids)
}

// 字段
func (p *User) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 角色列表
	roles, _ := service.NewRoleService().List()

	// 列表
	departments, _ := service.NewDepartmentService().GetList()

	// 职位列表
	positions, _ := (&model.Position{}).List()

	return []interface{}{
		field.ID("id", "ID"),
		field.Image("avatar", "头像"),
		field.Text("username", "用户名", func(row map[string]interface{}) interface{} {
			return "<a href='#/layout/index?api=/api/admin/user/edit&id=" + strconv.Itoa(row["id"].(int)) + "'>" + row["username"].(string) + "</a>"
		}).
			SetRules([]rule.Rule{
				rule.Required("用户名必须填写"),
				rule.Min(6, "用户名不能少于6个字符"),
				rule.Max(20, "用户名不能超过20个字符"),
			}).
			SetCreationRules([]rule.Rule{
				rule.Unique("users", "username", "用户名已存在"),
			}).
			SetUpdateRules([]rule.Rule{
				rule.Unique("users", "username", "{id}", "用户名已存在"),
			}),
		field.Text("nickname", "昵称").
			SetEditable(true).
			SetRules([]rule.Rule{
				rule.Required("昵称必须填写"),
			}),
		field.Checkbox("role_ids", "角色").
			SetOptions(roles).
			OnlyOnForms().
			HideWhenImporting(true),
		field.TreeSelect("department_id", "部门").
			SetTreeData(departments, "pid", "name", "id").
			OnlyOnForms().
			HideWhenImporting(true),
		field.Checkbox("position_ids", "职位").
			SetOptions(positions).
			OnlyOnForms().
			HideWhenImporting(true),
		field.Text("email", "邮箱").
			SetRules([]rule.Rule{
				rule.Required("邮箱必须填写"),
				rule.Email("邮箱格式错误"),
			}).
			SetCreationRules([]rule.Rule{
				rule.Unique("users", "email", "邮箱已存在"),
			}).
			SetUpdateRules([]rule.Rule{
				rule.Unique("users", "email", "{id}", "邮箱已存在"),
			}).OnlyOnForms(),
		field.Text("phone", "手机号").
			SetRules([]rule.Rule{
				rule.Required("手机号必须填写"),
				rule.Phone("手机号格式错误"),
			}).
			SetCreationRules([]rule.Rule{
				rule.Unique("users", "phone", "手机号已存在"),
			}).
			SetUpdateRules([]rule.Rule{
				rule.Unique("users", "phone", "{id}", "手机号已存在"),
			}),
		field.Radio("sex", "性别").
			SetRules([]rule.Rule{
				rule.Required("请选择性别"),
			}).
			SetOptions([]radio.Option{
				field.RadioOption("男", 1),
				field.RadioOption("女", 2),
			}).
			SetFilters(true).
			SetDefault(1),
		field.Password("password", "密码").
			SetCreationRules([]rule.Rule{
				rule.Required("密码必须填写"),
			}).
			SetRules([]rule.Rule{
				rule.Regexp(`/^.{6,}$/`, "密码不少于六位"),
				rule.Regexp(`/[A-Z]/`, "至少包含一个大写字母"),
				rule.Regexp(`/[a-z]/`, "至少包含一个小写字母"),
				rule.Regexp(`/[0-9]/`, "至少包含一个数字"),
				rule.Regexp(`/[!@#\$%\^&\*\(\)_\+\-\=\\\|\[\]\{\};':",\.<>\/?]/`, "至少包含一个特殊字符"),
			}).
			SetHelp("密码不少于六位，且至少包含一个大写字母、小写字母、数字和特殊字符").
			OnlyOnForms().
			ShowOnImporting(true),
		field.Datetime("last_login_time", "最后登录时间").OnlyOnIndex(),
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
func (p *User) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("username", "用户名"),
		searches.Input("nickname", "昵称"),
		searches.Status(),
		searches.DatetimeRange("last_login_time", "登录时间"),
	}
}

// 行为
func (p *User) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.Import(),
		actions.CreateLink(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
		actions.DetailLink(),
		actions.More().
			SetActions([]interface{}{
				actions.EditLink(),
				actions.DeleteSpecial(),
			}),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

// 编辑页面显示前回调
func (p *User) BeforeEditing(ctx *quark.Context, data map[string]interface{}) map[string]interface{} {
	delete(data, "password")
	roles, err := service.NewCasbinService().GetUserRoles(data["id"].(int))
	if err == nil {
		roleIds := []int{}
		for _, role := range roles {
			roleIds = append(roleIds, role.Id)
		}
		data["role_ids"] = roleIds
	}
	return data
}

// 保存数据前回调
func (p *User) BeforeSaving(ctx *quark.Context, submitData map[string]interface{}) (map[string]interface{}, error) {
	if submitData["password"] != nil {
		submitData["password"] = hash.Make(submitData["password"].(string))
	}
	return submitData, nil
}

// 保存后回调
func (p *User) AfterSaved(ctx *quark.Context, id int, data map[string]interface{}, result *gorm.DB) (err error) {
	if result.Error != nil {
		return result.Error
	}
	if data["role_ids"] != nil {
		if roleIds, ok := data["role_ids"].([]interface{}); ok {
			ids := []int{}
			for _, v := range roleIds {
				roleId := int(v.(float64))
				ids = append(ids, roleId)
			}
			err := service.NewCasbinService().AddUserRole(id, ids)
			if err != nil {
				return err
			}
		}
	}
	return err
}
