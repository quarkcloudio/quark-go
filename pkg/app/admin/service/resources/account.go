package resources

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/radio"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

type Account struct {
	resource.Template
}

// 初始化
func (p *Account) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "个人设置"

	// 模型
	p.Model = &model.Admin{}

	return p
}

// 表单接口
func (p *Account) FormApi(ctx *builder.Context) string {
	return "/api/admin/account/action/change-account"
}

// 字段
func (p *Account) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{

		field.Image("avatar", "头像").OnlyOnForms(),

		field.Text("nickname", "昵称").
			SetEditable(true).
			SetRules([]*rule.Rule{
				rule.New().SetRequired().SetMessage("昵称必须填写"),
			}),

		field.Text("email", "邮箱").
			SetRules([]*rule.Rule{
				rule.New().SetRequired().SetMessage("邮箱必须填写"),
			}).
			SetCreationRules([]*rule.Rule{
				rule.New().SetUnique("admins", "email").SetMessage("邮箱已存在"),
			}).
			SetUpdateRules([]*rule.Rule{
				rule.New().SetUnique("admins", "email", "{id}").SetMessage("邮箱已存在"),
			}),

		field.Text("phone", "手机号").
			SetRules([]*rule.Rule{
				rule.New().SetRequired().SetMessage("手机号必须填写"),
			}).
			SetCreationRules([]*rule.Rule{
				rule.New().SetUnique("admins", "phone").SetMessage("手机号已存在"),
			}).
			SetUpdateRules([]*rule.Rule{
				rule.New().SetUnique("admins", "phone", "{id}").SetMessage("手机号已存在"),
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
			SetDefault(1),

		field.Password("password", "密码").
			SetCreationRules([]*rule.Rule{
				rule.New().SetRequired().SetMessage("密码必须填写"),
			}).
			OnlyOnForms(),
	}
}

// 行为
func (p *Account) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.ChangeAccount(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

// 创建页面显示前回调
func (p *Account) BeforeCreating(ctx *builder.Context) map[string]interface{} {
	data := map[string]interface{}{}
	adminInfo, _ := (&model.Admin{}).GetAuthUser(ctx.Engine.GetConfig().AppKey, ctx.Token())
	db.Client.
		Model(p.Model).
		Where("id = ?", adminInfo.Id).
		First(&data)

	delete(data, "password")

	return data
}
