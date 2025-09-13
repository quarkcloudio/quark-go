package resources

import (
	"encoding/json"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/app/actions"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v4/component/form/rule"
	"github.com/quarkcloudio/quark-go/v4/dal/db"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/quarkcloudio/quark-go/v4/template/resource"
	"github.com/quarkcloudio/quark-go/v4/utils/hash"
	"gorm.io/gorm"
)

type Account struct {
	resource.Template
}

// 初始化
func (p *Account) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "个人设置"

	// 模型
	p.Model = &model.User{}

	return p
}

// 字段
func (p *Account) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Image("avatar", "头像"),
		field.Text("nickname", "昵称").
			SetRules([]rule.Rule{
				rule.Required("昵称必须填写"),
			}),
		field.Text("email", "邮箱").
			SetRules([]rule.Rule{
				rule.Required("邮箱必须填写"),
			}),
		field.Text("phone", "手机号").
			SetRules([]rule.Rule{
				rule.Required("手机号必须填写"),
			}),
		field.Radio("sex", "性别").
			SetOptions([]radio.Option{
				field.RadioOption("男", 1),
				field.RadioOption("女", 2),
			}).
			SetDefault(1),
		field.
			Password("password", "密码").
			SetRules([]rule.Rule{
				rule.Regexp(`/^.{6,}$/`, "密码不少于六位"),
				rule.Regexp(`/[A-Z]/`, "至少包含一个大写字母"),
				rule.Regexp(`/[a-z]/`, "至少包含一个小写字母"),
				rule.Regexp(`/[0-9]/`, "至少包含一个数字"),
				rule.Regexp(`/[!@#\$%\^&\*\(\)_\+\-\=\\\|\[\]\{\};':",\.<>\/?]/`, "至少包含一个特殊字符"),
			}).
			SetHelp("密码不少于六位，且至少包含一个大写字母、小写字母、数字和特殊字符"),
	}
}

// 行为
func (p *Account) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

// 表单显示前回调
func (p *Account) BeforeFormShowing(ctx *quark.Context) map[string]interface{} {
	data := map[string]interface{}{}
	adminInfo, _ := service.NewAuthService(ctx).GetAdmin()
	db.Client.
		Model(p.Model).
		Where("id = ?", adminInfo.Id).
		First(&data)
	delete(data, "password")
	return data
}

func (p *Account) FormHandle(ctx *quark.Context, query *gorm.DB, data map[string]interface{}) error {
	if data["avatar"] != "" && data["avatar"] != nil {
		data["avatar"], _ = json.Marshal(data["avatar"])
	}
	// 加密密码
	if data["password"] != nil {
		data["password"] = hash.Make(data["password"].(string))
	}
	// 获取登录管理员信息
	adminInfo, err := service.NewAuthService(ctx).GetAdmin()
	if err != nil {
		return ctx.JSONError(err.Error())
	}
	err = query.Where("id", adminInfo.Id).Updates(data).Error
	if err != nil {
		return ctx.JSONError(err.Error())
	}
	return ctx.JSONOk("操作成功")
}
