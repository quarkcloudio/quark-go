package resources

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

type Account struct {
	adminresource.Template
}

// 初始化
func (p *Account) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "个人设置"

	// 模型
	p.Model = &model.Admin{}

	return p
}

// 表单接口
func (p *Account) FormApi(request *builder.Request) string {

	return "admin/account/action/change-account"
}

// 字段
func (p *Account) Fields(request *builder.Request) []interface{} {
	field := &builder.AdminField{}

	return []interface{}{

		field.Image("avatar", "头像").OnlyOnForms(),

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
			}).SetDefault(1),

		field.Password("password", "密码").
			SetCreationRules(
				[]string{
					"required",
				},
				map[string]string{
					"required": "密码必须填写",
				},
			).OnlyOnForms(),
	}
}

// 行为
func (p *Account) Actions(request *builder.Request) []interface{} {
	return []interface{}{
		(&actions.ChangeAccount{}),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}

// 创建页面显示前回调
func (p *Account) BeforeCreating(request *builder.Request) map[string]interface{} {
	data := map[string]interface{}{}
	adminInfo, _ := (&model.Admin{}).GetAuthUser(request.Token())
	db.Client.
		Model(p.Model).
		Where("id = ?", adminInfo.Id).
		First(&data)

	delete(data, "password")

	return data
}
