package actions

import (
	"strings"

	"github.com/gobeam/stringy"
	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/actions"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

type SyncPermission struct {
	actions.Action
}

// 初始化
func (p *SyncPermission) Init() *SyncPermission {
	// 初始化父结构
	p.ParentInit()

	// 行为名称
	p.Name = "同步权限"

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	p.WithLoading = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	// 行为类型
	p.ActionType = "ajax"

	return p
}

// 执行行为句柄
func (p *SyncPermission) Handle(ctx *builder.Context, query *gorm.DB) error {
	// 获取当前权限
	permissions := ctx.Engine.GetUrlPaths()
	data := []model.Permission{}

	var names []string
	db.Client.Model(&model.Permission{}).Pluck("name", &names)
	for _, v := range permissions {
		if strings.Contains(v.Url, "/api/admin/") {
			url := strings.Replace(v.Url, "/api/admin/", "", -1)
			name := stringy.
				New(url).
				CamelCase("?", "")

			has := false
			for _, nv := range names {
				if nv == name {
					has = true
				}
			}

			if has == false {
				permission := model.Permission{
					MenuId:    0,
					Name:      name,
					Path:      v.Url,
					Method:    v.Method,
					GuardName: "admin",
				}
				data = append(data, permission)
			}
		}
	}
	if len(data) == 0 {
		return ctx.JSONError("暂无新增权限！")
	}

	err := query.Create(data).Error
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	err = db.Client.Model(&model.Permission{}).Where("name NOT IN ?", permissions).Delete("").Error
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	return ctx.JSONOk("操作成功")
}
