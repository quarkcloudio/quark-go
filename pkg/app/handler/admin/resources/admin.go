package resources

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/app/model"
	models "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/hash"
	"github.com/quarkcms/quark-go/pkg/msg"
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
func (p *Admin) Fields(request *builder.Request) []interface{} {

	field := &builder.AdminField{}

	// 角色列表
	roles, _ := (&model.Role{}).List()

	return []interface{}{
		field.ID("id", "ID"),

		field.Image("avatar", "头像").OnlyOnForms(),

		field.Text("username", "用户名", func() interface{} {

			return "<a href='#/index?api=/api/admin/admin/edit&id=" + strconv.Itoa(p.Field["id"].(int)) + "'>" + p.Field["username"].(string) + "</a>"
		}).
			SetRules(
				[]string{
					"required",
					"min:6",
					"max:20",
				},
				map[string]string{
					"required": "用户名必须填写",
					"min":      "用户名不能少于6个字符",
					"max":      "用户名不能超过20个字符",
				},
			).
			SetCreationRules(
				[]string{
					"unique:admins,username",
				},
				map[string]string{
					"unique": "用户名已存在",
				},
			).
			SetUpdateRules(
				[]string{
					"unique:admins,username,{id}",
				},
				map[string]string{
					"unique": "用户名已存在",
				},
			),

		field.Checkbox("role_ids", "角色").
			SetOptions(roles).
			OnlyOnForms(),

		field.Text("nickname", "昵称").
			SetEditable(true).
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "昵称必须填写",
				},
			),

		field.Editor("abc", "内容"),

		field.Text("email", "邮箱").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "邮箱必须填写",
				},
			).
			SetCreationRules(
				[]string{
					"unique:admins,email",
				},
				map[string]string{
					"unique": "邮箱已存在",
				},
			).
			SetUpdateRules(
				[]string{
					"unique:admins,email,{id}",
				},
				map[string]string{
					"unique": "邮箱已存在",
				},
			),

		field.Text("phone", "手机号").
			SetRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "手机号必须填写",
				},
			).
			SetCreationRules(
				[]string{
					"unique:admins,phone",
				},
				map[string]string{
					"unique": "手机号已存在",
				},
			).
			SetUpdateRules(
				[]string{
					"unique:admins,phone,{id}",
				},
				map[string]string{
					"unique": "手机号已存在",
				},
			),

		field.Radio("sex", "性别").
			SetOptions(map[interface{}]interface{}{
				1: "男",
				2: "女",
			}).SetDefault(1).
			SetColumn(func(column *table.Column) *table.Column {
				return column.SetFilters(true)
			}),

		field.Password("password", "密码").
			SetCreationRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "密码必须填写",
				},
			).OnlyOnForms(),

		field.Datetime("last_login_time", "最后登录时间", func() interface{} {
			if p.Field["last_login_time"] == nil {
				return p.Field["last_login_time"]
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
func (p *Admin) Searches(request *builder.Request) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("username", "用户名"),
		(&searches.Input{}).Init("nickname", "昵称"),
		(&searches.Status{}).Init(),
		(&searches.DateTimeRange{}).Init("last_login_time", "登录时间"),
	}
}

// 行为
func (p *Admin) Actions(request *builder.Request) []interface{} {

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
func (p *Admin) BeforeEditing(request *builder.Request, data map[string]interface{}) map[string]interface{} {
	delete(data, "password")

	roleIds := []int{}
	db.Client.
		Model(&model.ModelHasRole{}).
		Where("model_id = ?", data["id"]).
		Where("model_type = ?", "admin").
		Pluck("role_id", &roleIds)

	data["role_ids"] = roleIds

	return data
}

// 保存数据前回调
func (p *Admin) BeforeSaving(request *builder.Request, submitData map[string]interface{}) (map[string]interface{}, error) {

	// 加密密码
	if submitData["password"] != nil {
		submitData["password"] = hash.Make(submitData["password"].(string))
	}

	// 暂时清理role_ids
	delete(submitData, "role_ids")

	return submitData, nil
}

// 保存后回调
func (p *Admin) AfterSaved(request *builder.Request, model *gorm.DB) interface{} {
	data := map[string]interface{}{}
	json.Unmarshal(request.Body(), &data)
	if data["role_ids"] == nil {
		if model.Error != nil {
			return msg.Error(model.Error.Error(), "")
		}

		return msg.Success("操作成功！", strings.Replace("/index?api="+adminresource.IndexRoute, ":resource", request.Param("resource"), -1), "")
	}

	var result *gorm.DB
	if request.IsCreating() {
		last := map[string]interface{}{}
		model.Order("id desc").First(&last) // hack
		roleData := []map[string]interface{}{}
		for _, v := range data["role_ids"].([]interface{}) {
			item := map[string]interface{}{
				"role_id":    v,
				"model_type": "admin",
				"model_id":   last["id"],
			}
			roleData = append(roleData, item)
		}
		if len(roleData) > 0 {
			// 同步角色
			result = db.Client.Model(&models.ModelHasRole{}).Create(roleData)
		}
	} else {

		// 同步角色
		id := data["id"].(float64)
		roleData := []map[string]interface{}{}

		// 先清空用户对应的角色
		db.Client.Model(&models.ModelHasRole{}).Where("model_id = ?", id).Where("model_type = ?", "admin").Delete("")
		for _, v := range data["role_ids"].([]interface{}) {
			item := map[string]interface{}{
				"role_id":    v,
				"model_type": "admin",
				"model_id":   int(id),
			}
			roleData = append(roleData, item)
		}
		if len(roleData) > 0 {
			// 同步角色
			result = db.Client.Model(&models.ModelHasRole{}).Create(roleData)
		}
	}

	if result.Error != nil {
		return msg.Error(result.Error.Error(), "")
	}

	return msg.Success("操作成功！", strings.Replace("/index?api="+adminresource.IndexRoute, ":resource", request.Param("resource"), -1), "")
}
