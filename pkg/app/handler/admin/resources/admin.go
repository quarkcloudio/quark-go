package resources

import (
	"strconv"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/radio"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/hash"
	"gorm.io/gorm"
)

type Admin struct {
	adminresource.Template
}

// 初始化
func (p *Admin) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "管理员"

	// 模型
	p.Model = &model.Admin{}

	// 分页
	p.PerPage = 10

	p.WithExport = true

	return p
}

// 字段
func (p *Admin) Fields(ctx *builder.Context) []interface{} {

	field := &adminresource.Field{}

	// 角色列表
	roles, _ := (&model.Role{}).List()

	return []interface{}{
		field.ID("id", "ID"),

		field.Image("avatar", "头像").OnlyOnForms(),

		field.Text("username", "用户名", func() interface{} {

			return "<a href='#/layout/index?api=/api/admin/admin/edit&id=" + strconv.Itoa(p.Field["id"].(int)) + "'>" + p.Field["username"].(string) + "</a>"
		}).
			SetRules([]*rule.Rule{
				rule.Required(true, "用户名必须填写"),
				rule.Min(6, "用户名不能少于6个字符"),
				rule.Max(20, "用户名不能超过20个字符"),
			}).
			SetCreationRules([]*rule.Rule{
				rule.Unique("admins", "username", "用户名已存在"),
			}).
			SetUpdateRules([]*rule.Rule{
				rule.Unique("admins", "username", "{id}", "用户名已存在"),
			}),

		field.Checkbox("role_ids", "角色").
			SetOptions(roles).
			OnlyOnForms(),

		field.Text("nickname", "昵称").
			SetEditable(true).
			SetRules([]*rule.Rule{
				rule.Required(true, "昵称必须填写"),
			}),

		field.Text("email", "邮箱").
			SetRules([]*rule.Rule{
				rule.Required(true, "邮箱必须填写"),
			}).
			SetCreationRules([]*rule.Rule{
				rule.Unique("admins", "email", "邮箱已存在"),
			}).
			SetUpdateRules([]*rule.Rule{
				rule.Unique("admins", "email", "{id}", "邮箱已存在"),
			}),

		field.Text("phone", "手机号").
			SetRules([]*rule.Rule{
				rule.Required(true, "手机号必须填写"),
			}).
			SetCreationRules([]*rule.Rule{
				rule.Unique("admins", "phone", "手机号已存在"),
			}).
			SetUpdateRules([]*rule.Rule{
				rule.Unique("admins", "phone", "{id}", "手机号已存在"),
			}),

		field.Radio("sex", "性别").
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

		field.Datetime("last_login_time", "最后登录时间", func() interface{} {
			if p.Field["last_login_time"] == nil {
				return p.Field["last_login_time"]
			}

			if p.Field["last_login_time"].(time.Time).Format("2006-01-02 15:04:05") == "0001-01-01 00:00:00" {
				return nil
			}

			return p.Field["last_login_time"].(time.Time).Format("2006-01-02 15:04:05")
		}).OnlyOnIndex(),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Admin) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("username", "用户名"),
		(&searches.Input{}).Init("nickname", "昵称"),
		(&searches.Status{}).Init(),
		(&searches.DateTimeRange{}).Init("last_login_time", "登录时间"),
	}
}

// 行为
func (p *Admin) Actions(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&actions.Import{}).Init(),
		(&actions.CreateLink{}).Init(p.Title),
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Disable{}).Init("批量禁用"),
		(&actions.Enable{}).Init("批量启用"),
		(&actions.DetailLink{}).Init("详情"),
		(&actions.MoreActions{}).Init("更多").SetActions([]interface{}{
			(&actions.EditLink{}).Init("编辑"),
			(&actions.Delete{}).Init("删除"),
		}),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}

// 编辑页面显示前回调
func (p *Admin) BeforeEditing(ctx *builder.Context, data map[string]interface{}) map[string]interface{} {
	delete(data, "password")

	roleIds := []int{}
	db.Client.
		Model(&model.ModelHasRole{}).
		Where("model_id = ?", data["id"]).
		Where("model_type = ?", "admin").
		Distinct().
		Pluck("role_id", &roleIds)

	data["role_ids"] = roleIds

	return data
}

// 保存数据前回调
func (p *Admin) BeforeSaving(ctx *builder.Context, submitData map[string]interface{}) (map[string]interface{}, error) {

	// 加密密码
	if submitData["password"] != nil {
		submitData["password"] = hash.Make(submitData["password"].(string))
	}

	return submitData, nil
}

// 保存后回调
func (p *Admin) AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) error {

	// 导入操作，直接返回
	if ctx.IsImport() {
		return result.Error
	}

	// 返回错误信息
	if result.Error != nil {
		return ctx.JSONError(result.Error.Error())
	}

	// 编辑操作，先清空用户对应的角色
	if ctx.IsEditing() {
		db.Client.Model(&model.ModelHasRole{}).Where("model_id = ?", id).Where("model_type = ?", "admin").Delete("")
	}

	if data["role_ids"] == nil {
		return ctx.JSONOk("操作成功！", strings.Replace("/layout/index?api="+adminresource.IndexPath, ":resource", ctx.Param("resource"), -1))
	}

	roleData := []map[string]interface{}{}
	for _, v := range data["role_ids"].([]interface{}) {
		item := map[string]interface{}{
			"role_id":    v,
			"model_type": "admin",
			"model_id":   id,
		}
		roleData = append(roleData, item)
	}
	if len(roleData) > 0 {
		// 同步角色
		err := db.Client.Model(&model.ModelHasRole{}).Create(roleData).Error
		if err != nil {
			return ctx.JSONError(err.Error())
		}
	}

	return ctx.JSONOk("操作成功！", strings.Replace("/layout/index?api="+adminresource.IndexPath, ":resource", ctx.Param("resource"), -1))
}
