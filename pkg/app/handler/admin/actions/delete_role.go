package actions

import (
	"strconv"
	"strings"

	"github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/actions"
	"gorm.io/gorm"
)

type DeleteRole struct {
	actions.Action
}

// 初始化
func (p *DeleteRole) Init(name string) *DeleteRole {
	// 初始化父结构
	p.ParentInit()

	// 行为名称，当行为在表格行展示时，支持js表达式
	p.Name = name

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要删除吗？", "删除后数据将无法恢复，请谨慎操作！", "modal")

	if name == "删除" {
		p.SetOnlyOnIndexTableRow(true)
	}

	if name == "批量删除" {
		p.SetOnlyOnIndexTableAlert(true)
	}

	return p
}

/**
 * 行为接口接收的参数，当行为在表格行展示的时候，可以配置当前行的任意字段
 *
 * @return array
 */
func (p *DeleteRole) GetApiParams() []string {
	return []string{
		"id",
	}
}

// 执行行为句柄
func (p *DeleteRole) Handle(ctx *builder.Context, query *gorm.DB) error {
	id := ctx.Query("id")
	if id == "" {
		return ctx.JSONError("参数错误")
	}

	err := query.Delete("").Error
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	ids := strings.Split(id.(string), ",")
	if len(ids) > 0 {
		for _, v := range ids {
			idInt, err := strconv.Atoi(v)
			if err != nil {
				return ctx.JSONError(err.Error())
			}

			// 清理casbin里的角色
			(&model.CasbinRule{}).RemoveRoleMenuAndPermissions(idInt)
		}
	} else {
		idInt, err := strconv.Atoi(id.(string))
		if err != nil {
			return ctx.JSONError(err.Error())
		}

		// 清理casbin里的角色
		(&model.CasbinRule{}).RemoveRoleMenuAndPermissions(idInt)
	}

	return ctx.JSONOk("操作成功")
}
