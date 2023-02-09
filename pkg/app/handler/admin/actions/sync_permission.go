package actions

import (
	models "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/actions"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
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
func (p *SyncPermission) Handle(ctx *builder.Context, model *gorm.DB) interface{} {
	// 获取当前权限
	permissions := ctx.Engine.GetRouterPaths()
	data := []models.Permission{}

	var names []string
	db.Client.Model(&models.Permission{}).Pluck("name", &names)
	for _, v := range permissions {
		has := false
		for _, nv := range names {
			if nv == v {
				has = true
			}
		}
		if has == false {
			permission := models.Permission{
				MenuId:    0,
				Name:      v,
				GuardName: "admin",
			}
			data = append(data, permission)
		}
	}
	if len(data) == 0 {
		return msg.Error("暂无新增权限！", "")
	}

	err := model.Create(data).Error
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	err = db.Client.Model(&models.Permission{}).Where("name NOT IN ?", permissions).Delete("").Error
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	return msg.Success("操作成功", "", "")
}
