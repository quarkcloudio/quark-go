package resources

import (
	"strconv"
	"strings"

	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/form/rule"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/actions"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/service/searches"
	"github.com/quarkcloudio/quark-go/v3/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v3/pkg/builder"
	"github.com/quarkcloudio/quark-go/v3/pkg/utils/hash"
	"gorm.io/gorm"
)

type User struct {
	resource.Template
}

// 初始化
func (p *User) Init(ctx *builder.Context) interface{} {

	// 树形下拉框
	p.TableTreeBar.SetTreeData((&model.Department{}).TableTree())

	// 标题
	p.Title = "用户"

	// 模型
	p.Model = &model.User{}

	// 分页
	p.PerPage = 10

	// 是否具有导出功能
	p.WithExport = true

	return p
}

// 字段
func (p *User) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	// 角色列表
	roles, _ := (&model.Role{}).List()

	return []interface{}{
		field.ID("id", "ID"),

		field.Image("avatar", "头像").OnlyOnForms(),

		field.Text("username", "用户名", func() interface{} {

			return "<a href='#/layout/index?api=/api/admin/user/edit&id=" + strconv.Itoa(p.Field["id"].(int)) + "'>" + p.Field["username"].(string) + "</a>"
		}).
			SetRules([]*rule.Rule{
				rule.Required(true, "用户名必须填写"),
				rule.Min(6, "用户名不能少于6个字符"),
				rule.Max(20, "用户名不能超过20个字符"),
			}).
			SetCreationRules([]*rule.Rule{
				rule.Unique("users", "username", "用户名已存在"),
			}).
			SetUpdateRules([]*rule.Rule{
				rule.Unique("users", "username", "{id}", "用户名已存在"),
			}),

		field.Checkbox("role_ids", "角色").
			SetOptions(roles).
			OnlyOnForms().
			HideWhenImporting(true),

		field.Text("nickname", "昵称").
			SetEditable(true).
			SetRules([]*rule.Rule{
				rule.Required(true, "昵称必须填写"),
			}),

		field.Text("email", "邮箱").
			SetRules([]*rule.Rule{
				rule.Required(true, "邮箱必须填写"),
				rule.Email("邮箱格式错误"),
			}).
			SetCreationRules([]*rule.Rule{
				rule.Unique("admins", "email", "邮箱已存在"),
			}).
			SetUpdateRules([]*rule.Rule{
				rule.Unique("admins", "email", "{id}", "邮箱已存在"),
			}).OnlyOnForms(),

		field.Text("phone", "手机号").
			SetRules([]*rule.Rule{
				rule.Required(true, "手机号必须填写"),
				rule.Phone("手机号格式错误"),
			}).
			SetCreationRules([]*rule.Rule{
				rule.Unique("admins", "phone", "手机号已存在"),
			}).
			SetUpdateRules([]*rule.Rule{
				rule.Unique("admins", "phone", "{id}", "手机号已存在"),
			}),

		field.Radio("sex", "性别").
			SetRules([]*rule.Rule{
				rule.Required(true, "请选择性别"),
			}).
			SetOptions([]*radio.Option{
				{
					Value: 1,
					Label: "男",
				},
				{
					Value: 2,
					Label: "女",
				},
			}).
			SetFilters(true).
			SetDefault(1),

		field.Password("password", "密码").
			SetCreationRules([]*rule.Rule{
				rule.Required(true, "密码必须填写"),
			}).
			OnlyOnForms().
			ShowOnImporting(true),

		field.Datetime("last_login_time", "最后登录时间").OnlyOnIndex(),

		field.Switch("status", "状态").
			SetRules([]*rule.Rule{
				rule.Required(true, "请选择状态"),
			}).
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *User) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		searches.Input("username", "用户名"),
		searches.Input("nickname", "昵称"),
		searches.Status(),
		searches.DatetimeRange("last_login_time", "登录时间"),
	}
}

// 行为
func (p *User) Actions(ctx *builder.Context) []interface{} {

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
func (p *User) BeforeEditing(ctx *builder.Context, data map[string]interface{}) map[string]interface{} {

	// 编辑页面清理password
	delete(data, "password")

	roles, err := (&model.CasbinRule{}).GetUserRoles(data["id"].(int))
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
func (p *User) BeforeSaving(ctx *builder.Context, submitData map[string]interface{}) (map[string]interface{}, error) {

	// 加密密码
	if submitData["password"] != nil {
		submitData["password"] = hash.Make(submitData["password"].(string))
	}

	return submitData, nil
}

// 保存后回调
func (p *User) AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) error {

	// 导入操作，直接返回
	if ctx.IsImport() {
		return result.Error
	}

	// 返回错误信息
	if result.Error != nil {
		return ctx.JSON(200, message.Error(result.Error.Error()))
	}

	if data["role_ids"] != nil {
		if roleIds, ok := data["role_ids"].([]interface{}); ok {
			ids := []int{}
			for _, v := range roleIds {
				roleId := int(v.(float64))
				ids = append(ids, roleId)
			}

			err := (&model.CasbinRule{}).AddUserRole(id, ids)
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
