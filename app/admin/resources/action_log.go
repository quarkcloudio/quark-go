package resources

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
)

type ActionLog struct {
	resource.Template
}

// 初始化
func (p *ActionLog) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "操作日志"

	// 模型
	p.Model = &model.ActionLog{}

	// 分页
	p.PageSize = 10

	// 是否具有导出功能
	p.Export = true

	return p
}

// 字段
func (p *ActionLog) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),
		field.Text("username", "用户信息", func(row map[string]interface{}) interface{} {
			userInfo, err := service.NewUserService().GetInfoById(row["uid"])
			if err != nil {
				return ""
			}
			return "账号：<a href='#/layout/index?api=/api/admin/user/detail&id=" + strconv.Itoa(userInfo.Id) + "'>" + userInfo.Username + "</a><br/>昵称：" + userInfo.Nickname
		}),
		field.Text("url", "行为").SetEllipsis(true),
		field.Text("ip", "IP"),
		field.Datetime("created_at", "发生时间"),
	}
}

// 搜索
func (p *ActionLog) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("username", "账号"),
		searches.Input("url", "行为"),
		searches.Input("ip", "IP"),
	}
}

// 行为
func (p *ActionLog) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.BatchDelete(),
		actions.Delete(),
	}
}
