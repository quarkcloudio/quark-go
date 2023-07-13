package table

import "github.com/quarkcms/quark-go/v2/pkg/app/admin/component/component"

type Component struct {
	component.Element
	RowKey           string          `json:"rowKey"`
	Api              string          `json:"api"`
	ApiType          string          `json:"apiType"`
	TableLayout      string          `json:"tableLayout"`
	HeaderTitle      string          `json:"headerTitle"`
	Columns          interface{}     `json:"columns"`
	RowSelection     interface{}     `json:"rowSelection"`
	Options          map[string]bool `json:"options"`
	Search           interface{}     `json:"search"`
	BatchActions     interface{}     `json:"batchActions"`
	DateFormatter    string          `json:"dateFormatter"`
	ColumnEmptyText  string          `json:"columnEmptyText"`
	ToolBar          interface{}     `json:"toolBar"`
	TableExtraRender interface{}     `json:"tableExtraRender"`
	Scroll           interface{}     `json:"scroll"`
	Striped          bool            `json:"striped"`
	Datasource       interface{}     `json:"datasource"`
	Pagination       interface{}     `json:"pagination"`
	Polling          int             `json:"polling"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 获取Column
func NewColumn() *Column {

	return (&Column{}).Init()
}

// 获取Search
func NewSearch() *Search {

	return (&Search{}).Init()
}

// 获取SearchItem
func NewToolBar() *ToolBar {

	return (&ToolBar{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "table"
	p.RowKey = "id"
	p.Api = ""
	p.ApiType = "GET"
	p.ColumnEmptyText = "-"
	p.DateFormatter = "string"
	p.Options = map[string]bool{
		"fullScreen": true, "reload": true, "setting": true,
	}
	p.RowSelection = []interface{}{}
	p.SetKey("table", false)

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
 * 表格元素的 table-layout 属性，设为 fixed 表示内容不会影响列的布局,- | auto | fixed
 *
 * @param  string  tableLayout
 * @return p
 */
func (p *Component) SetTableLayout(tableLayout string) *Component {
	p.TableLayout = tableLayout

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
 * 是否显示搜索表单，传入对象时为搜索表单的配置
 *
 * @param  Closure  callback
 * @return p
 */
func (p *Component) SetSearch(callback interface{}) interface{} {
	// callback(p.Search)

	return p.Search
}

/**
 * 搜索表单的配置
 *
 * @param  array  search
 * @return p
 */
func (p *Component) SetSearches(search interface{}) *Component {
	p.Search = search

	return p
}

/**
 * 批量设置表格列
 *
 * @param array columns
 * @return p
 */
func (p *Component) SetColumns(columns interface{}) *Component {
	p.Columns = columns

	return p
}

/**
 * 批量操作选择项
 *
 * @param array rowSelection
 * @return p
 */
func (p *Component) SetRowSelection(rowSelection interface{}) *Component {
	p.RowSelection = rowSelection

	return p
}

/**
 * table 工具栏，设为 false 时不显示,{ fullScreen: true, reload: true ,setting: true}
 *
 * @param  array|bool  options
 * @return p
 */
func (p *Component) SetOptions(options map[string]bool) *Component {
	p.Options = options

	return p
}

/**
 * 转化 moment 格式数据为特定类型，false 不做转化,"string" | "number" | false
 *
 * @param  string  dateFormatter
 * @return p
 */
func (p *Component) SetDateFormatter(dateFormatter string) *Component {
	p.DateFormatter = dateFormatter

	return p
}

/**
 * 空值时的显示，不设置 则默认显示 -
 *
 * @param  string  columnEmptyText
 * @return p
 */
func (p *Component) SetColumnEmptyText(columnEmptyText string) *Component {
	p.ColumnEmptyText = columnEmptyText

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
 * 表格的批量操作
 *
 * @param  array  batchActions
 * @return p
 */
func (p *Component) SetBatchActions(batchActions interface{}) *Component {
	p.BatchActions = batchActions

	return p
}

/**
 * 自定义表格的主体函数
 *
 * @param  array  tableExtraRender
 * @return p
 */
func (p *Component) SetTableExtraRender(tableExtraRender interface{}) *Component {
	p.TableExtraRender = tableExtraRender

	return p
}

/**
 * 设置表格滚动
 *
 * @param  array  scroll
 * @return p
 */
func (p *Component) scroll(scroll interface{}) *Component {
	p.Scroll = scroll

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

/**
 * 是否轮询
 *
 * @param  null|number  polling
 * @return p
 */
func (p *Component) SetPolling(polling int) *Component {
	p.Polling = polling

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "table"

	return p
}
