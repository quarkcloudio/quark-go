package resources

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/quarkcloudio/quark-go/v4"
	"github.com/quarkcloudio/quark-go/v4/app/actions"
	"github.com/quarkcloudio/quark-go/v4/app/searches"
	"github.com/quarkcloudio/quark-go/v4/component/form/fields/radio"
	"github.com/quarkcloudio/quark-go/v4/component/form/rule"
	"github.com/quarkcloudio/quark-go/v4/model"
	"github.com/quarkcloudio/quark-go/v4/service"
	"github.com/quarkcloudio/quark-go/v4/template/resource"
	"gorm.io/gorm"
)

type Menu struct {
	resource.Template
}

// 初始化
func (p *Menu) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "菜单"

	// 模型
	p.Model = &model.Menu{}

	// 列表页数据转换成树
	p.TableListToTree = true

	// 默认排序
	p.QueryOrder = "sort asc"

	// 分页
	p.PageSize = false

	return p
}

// 字段
func (p *Menu) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	// 权限列表
	permissions, _ := service.NewPermissionService().DataSource()

	// 菜单列表
	menus, _ := service.NewMenuService().GetListWithRoot()

	return []interface{}{
		field.Hidden("id", "ID"),                 // 列表读取且不展示的字段
		field.Hidden("pid", "PID").OnlyOnIndex(), // 列表读取且不展示的字段
		field.Hidden("query", "查询参数"),
		field.Hidden("api", "API接口"),
		field.Hidden("url", "URL地址"),
		field.Hidden("component", "组件"),
		field.Group([]interface{}{
			field.Text("name", "名称").
				SetRules([]rule.Rule{
					rule.Required("名称必须填写"),
				}),
			field.Text("guard_name", "守卫").
				SetRules([]rule.Rule{
					rule.Required("守卫必须填写"),
				}).
				SetDefault("admin").
				OnlyOnForms(),
			field.Radio("type", "类型").
				SetOptions([]radio.Option{
					field.RadioOption("目录", 1),
					field.RadioOption("菜单", 2),
					field.RadioOption("按钮", 3),
				}).
				SetRules([]rule.Rule{
					rule.Required("类型必须选择"),
				}).
				SetDefault(1),
		}),
		field.Group([]interface{}{
			field.Icon("icon", "图标").OnlyOnForms(),
			field.Number("sort", "排序").
				SetEditable(true).
				SetDefault(0),
			field.TreeSelect("pid", "上级菜单").
				SetTreeData(menus, -1, "pid", "name", "id").
				SetDefault(0).
				OnlyOnForms(),
		}),
		field.Dependency().
			SetWhen("type", 1, func() interface{} {
				return []interface{}{
					field.Text("path", "路由").
						SetRules([]rule.Rule{
							rule.Required("路由必须填写"),
						}).
						SetEditable(true).
						SetHelp("访问的路由地址，如：`user`").
						SetWidth("400px"),
				}
			}),
		field.Dependency().
			SetWhen("type", 2, func() interface{} {
				return []interface{}{
					field.Radio("page_type", "页面类型").
						SetOptions([]radio.Option{
							field.RadioOption("默认", 1),
							field.RadioOption("引擎", 2),
							field.RadioOption("外链", 3),
							field.RadioOption("iframe", 4),
						}).
						SetRules([]rule.Rule{
							rule.Required("页面类型必须选择"),
						}).
						SetDefault(1).
						OnlyOnForms(),
					field.Dependency().
						SetWhen("page_type", 1, func() interface{} {
							return []interface{}{
								field.Text("path", "路由地址").
									SetRules([]rule.Rule{
										rule.Required("路由地址必须填写"),
									}).
									SetEditable(true).
									SetHelp("访问的路由地址，如：`user`").
									SetWidth("400px").
									OnlyOnForms(),
								field.Text("component", "组件路径").
									SetRules([]rule.Rule{
										rule.Required("组件路径必须填写"),
									}).
									SetEditable(true).
									SetHelp("访问的组件路径，如：`user/index`").
									SetWidth("400px").
									OnlyOnForms(),
							}
						}),
					field.Dependency().
						SetWhen("page_type", 2, func() interface{} {
							return []interface{}{
								field.Text("path", "路由地址").
									SetRules([]rule.Rule{
										rule.Required("路由地址必须填写"),
									}).
									SetEditable(true).
									SetHelp("访问的路由地址，如：`user`").
									SetWidth("400px").
									OnlyOnForms(),
								field.Text("api", "接口地址").
									SetRules([]rule.Rule{
										rule.Required("接口地址必须填写"),
									}).
									SetEditable(true).
									SetHelp("引擎接口地址，如：`/api/admin/user/index`").
									SetWidth("400px").
									OnlyOnForms(),
							}
						}),
					field.Dependency().
						SetWhen("page_type", 3, func() interface{} {
							return []interface{}{
								field.Text("path", "外链地址").
									SetRules([]rule.Rule{
										rule.Required("外链地址必须填写"),
									}).
									SetEditable(true).
									SetHelp("访问的外链地址，以`http(s)://`开头").
									SetWidth("400px").
									OnlyOnForms(),
							}
						}),
					field.Dependency().
						SetWhen("page_type", 4, func() interface{} {
							return []interface{}{
								field.Text("path", "路由地址").
									SetRules([]rule.Rule{
										rule.Required("路由地址必须填写"),
									}).
									SetEditable(true).
									SetHelp("访问的路由地址，如：`user`").
									SetWidth("400px").
									OnlyOnForms(),
								field.Text("url", "iframe地址").
									SetRules([]rule.Rule{
										rule.Required("iframe地址必须填写"),
									}).
									SetEditable(true).
									SetHelp("访问的iframe地址，以`http(s)://`开头").
									SetWidth("400px").
									OnlyOnForms(),
							}
						}),
				}
			}),
		field.Dependency().
			SetWhen("type", ">", 1, func() interface{} {
				return []interface{}{
					field.Text("permission", "权限标识").
						SetHelp("鉴权标识，如：`user:index`").
						SetWidth("400px"),
				}
			}),
		field.Group([]interface{}{
			field.Switch("visible", "显示").
				SetTrueValue("显示").
				SetFalseValue("隐藏").
				SetEditable(true).
				SetDefault(true),
			field.Switch("status", "状态").
				SetTrueValue("正常").
				SetFalseValue("禁用").
				SetEditable(true).
				SetDefault(true),
		}),
		field.Dependency().
			SetWhen("type", 3, func() interface{} {
				return []interface{}{
					field.Transfer("permission_ids", "绑定权限").
						SetDataSource(permissions).
						SetListStyle(map[string]interface{}{
							"width":  "320px",
							"height": "300px",
						}).
						SetShowSearch(true).
						OnlyOnForms(),
				}
			}),
	}
}

// 搜索
func (p *Menu) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
		searches.Input("path", "路由"),
		searches.Status(),
	}
}

// 行为
func (p *Menu) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.MenuCreateDrawer(),
		actions.BatchDelete(),
		actions.ChangeStatus(),
		actions.MenuEditDrawer(),
		actions.Delete(),
	}
}

// 编辑页面显示前回调
func (p *Menu) BeforeEditing(ctx *quark.Context, data map[string]interface{}) map[string]interface{} {
	id := ctx.Query("id", "")
	idInt, err := strconv.Atoi(id.(string))
	if id != "" && err == nil {
		permissionIds := []int{}
		permissions, err := service.NewPermissionService().GetMenuPermissions(idInt)
		if err == nil {
			for _, v := range permissions {
				permissionIds = append(permissionIds, v.Id)
			}
		}
		data["permission_ids"] = permissionIds
	}

	switch data["page_type"].(int) {
	case 2:
		var query map[string]interface{}
		json.NewDecoder(strings.NewReader(data["query"].(string))).Decode(&query)
		data["api"] = query["api"]
	case 4:
		var query map[string]interface{}
		json.NewDecoder(strings.NewReader(data["query"].(string))).Decode(&query)
		data["url"] = query["url"]
	}

	return data
}

// 保存数据前回调
func (p *Menu) BeforeSaving(ctx *quark.Context, submitData map[string]interface{}) (map[string]interface{}, error) {
	if submitData["page_type"] != nil {
		if submitData["page_type"].(float64) == 2 {
			submitData["query"] = fmt.Sprintf(`{"api":"%s"}`, submitData["api"])
		}
		if submitData["page_type"].(float64) == 4 {
			submitData["query"] = fmt.Sprintf(`{"url":"%s"}`, submitData["url"])
		}
	}

	return submitData, nil
}

// 保存后回调
func (p *Menu) AfterSaved(ctx *quark.Context, id int, data map[string]interface{}, result *gorm.DB) error {
	if data["permission_ids"] != nil {
		err := service.NewPermissionService().AddMenuPermission(id, data["permission_ids"])
		if err != nil {
			return err
		}
	}

	return result.Error
}
