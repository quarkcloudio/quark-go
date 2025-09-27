package actions

import (
	"strconv"
	"strings"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/quarkcloudio/quark-go/v4/template/resource/actions"
	"gorm.io/gorm"
)

type BatchDeleteRoleAction struct {
	actions.Batch
}

// 批量删除角色
func BatchDeleteRole() *BatchDeleteRoleAction {
	return &BatchDeleteRoleAction{}
}

// 初始化
func (p *BatchDeleteRoleAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮文字
	p.Name = "批量删除"

	// 图标
	p.Icon = "ant-design:delete-outlined"

	// 危险按钮
	p.Danger = true

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要删除吗？", "删除后数据将无法恢复，请谨慎操作！", "modal")

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	return p
}

// 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
func (p *BatchDeleteRoleAction) GetApiParams() []string {
	return []string{
		"id",
	}
}

// 执行行为句柄
func (p *BatchDeleteRoleAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	id := ctx.Query("id")
	if id == "" {
		return ctx.CJSONError("参数错误")
	}

	err := query.Delete("").Error
	if err != nil {
		return ctx.CJSONError(err.Error())
	}

	ids := strings.Split(id.(string), ",")
	if len(ids) > 0 {
		for _, v := range ids {
			idInt, err := strconv.Atoi(v)
			if err != nil {
				return ctx.CJSONError(err.Error())
			}
			// 清理casbin里的角色
			service.NewCasbinService().RemoveRoleMenuAndPermissions(idInt)
		}
	} else {
		idInt, err := strconv.Atoi(id.(string))
		if err != nil {
			return ctx.CJSONError(err.Error())
		}
		// 清理casbin里的角色
		service.NewCasbinService().RemoveRoleMenuAndPermissions(idInt)
	}
	return ctx.CJSONOk("操作成功")
}
