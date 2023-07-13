package lists

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	RowKey       string      `json:"rowKey"`
	Api          string      `json:"api"`
	ApiType      string      `json:"apiType"`
	HeaderTitle  string      `json:"headerTitle"`
	Metas        interface{} `json:"metas"`
	RowSelection interface{} `json:"rowSelection"`
	Striped      bool        `json:"striped"`
	Datasource   interface{} `json:"datasource"`
	Pagination   interface{} `json:"pagination"`
	ToolBar      interface{} `json:"toolBar"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 获取Meta
func NewMeta() *Meta {
	return (&Meta{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "list"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// layout 的左上角 的 title
func (p *Component) SetRowKey(rowKey string) *Component {
	p.RowKey = rowKey
	return p
}

/**
 * 获取表格数据接口
 *
 * @param  string  api
 * @return p
 */
func (p *Component) SetApi(api string) *Component {
	p.Api = api

	return p
}

/**
 * 获取表格数据接口类型
 *
 * @param  string  apiType
 * @return p
 */
func (p *Component) SetApiType(apiType string) *Component {
	p.ApiType = apiType

	return p
}

/**
 * 表头标题
 *
 * @param  string  title
 * @return p
 */
func (p *Component) SetTitle(title string) *Component {
	p.HeaderTitle = title

	return p
}

/**
 * 表头标题
 *
 * @param  string  headerTitle
 * @return p
 */
func (p *Component) SetHeaderTitle(headerTitle string) *Component {
	p.HeaderTitle = headerTitle

	return p
}

/**
 * 批量设置表格列
 *
 * @param array metas
 * @return p
 */
func (p *Component) SetMetas(metas interface{}) *Component {

	limits := []string{
		"type",
		"title",
		"subTitle",
		"description",
		"avatar",
		"actions",
		"content",
		"extra",
	}

	for k, _ := range metas.(map[string]interface{}) {
		inSlice := false
		for _, limit := range limits {
			if limit == k {
				inSlice = true
			}
		}

		if inSlice == false {
			panic("meta index key must be in 'type','title','subTitle','description','avatar','actions','content','extra'!")
		}
	}

	p.Metas = metas

	return p
}

/**
 * 批量操作选择项
 *
 * @param array rowSelection
 * @return p
 */
func (p *Component) SetRowSelection(rowSelection []interface{}) *Component {
	p.RowSelection = rowSelection

	return p
}

/**
 * 透传 ProUtils 中的 ListToolBar 配置项
 *
 * @param  void
 * @return p
 */
func (p *Component) SetToolBar(toolBar interface{}) *Component {
	p.ToolBar = toolBar

	return p
}

/**
 * 设置表格滚动
 *
 * @param  bool  striped
 * @return p
 */
func (p *Component) striped(striped bool) *Component {
	p.Striped = striped

	return p
}

/**
 * 表格数据
 *
 * @param  array|string  datasource
 * @return p
 */
func (p *Component) SetDatasource(datasource interface{}) *Component {
	p.Datasource = datasource

	return p
}

/**
 * 表格分页
 *
 * @param  number  current
 * @param  number  pageSize
 * @param  number  total
 * @param  number  defaultCurrent
 * @return p
 */
func (p *Component) SetPagination(current int, pageSize int, total int, defaultCurrent int) *Component {
	p.Pagination = map[string]int{
		"current":        current,
		"pageSize":       pageSize,
		"total":          total,
		"defaultCurrent": defaultCurrent,
	}

	return p
}
