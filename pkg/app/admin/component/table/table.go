package table

import "github.com/quarkcloudio/quark-go/v3/pkg/app/admin/component/component"

type Expandable struct {
	ChildrenColumnName     string        `json:"childrenColumnName,omitempty"`
	ColumnTitle            interface{}   `json:"columnTitle,omitempty"`
	ColumnWidth            interface{}   `json:"columnWidth,omitempty"`
	DefaultExpandAllRows   bool          `json:"defaultExpandAllRows,omitempty"`
	DefaultExpandedRowKeys []interface{} `json:"defaultExpandedRowKeys,omitempty"`
	ExpandedRowClassName   string        `json:"expandedRowClassName,omitempty"`
	ExpandedRowKeys        []interface{} `json:"expandedRowKeys,omitempty"`
	ExpandIcon             interface{}   `json:"expandIcon,omitempty"`
	ExpandRowByClick       bool          `json:"expandRowByClick,omitempty"`
	Fixed                  interface{}   `json:"fixed,omitempty"`
	IndentSize             int           `json:"indentSize,omitempty"`
	RowExpandable          bool          `json:"rowExpandable,omitempty"`
	ShowExpandColumn       bool          `json:"showExpandColumn,omitempty"`
}

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
	TreeBar          interface{}     `json:"treeBar"`
	TableExtraRender interface{}     `json:"tableExtraRender"`
	Expandable       *Expandable     `json:"expandable,omitempty"`
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

// 获取工具栏
func NewToolBar() *ToolBar {

	return (&ToolBar{}).Init()
}

// 获取TreeBar
func NewTreeBar() *TreeBar {

	return (&TreeBar{}).Init()
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

// 获取表格数据接口
func (p *Component) SetApi(api string) *Component {
	p.Api = api

	return p
}

// 获取表格数据接口类型
func (p *Component) SetApiType(apiType string) *Component {
	p.ApiType = apiType

	return p
}

// 表格元素的 table-layout 属性，设为 fixed 表示内容不会影响列的布局,- | auto | fixed
func (p *Component) SetTableLayout(tableLayout string) *Component {
	p.TableLayout = tableLayout

	return p
}

//  表头标题
func (p *Component) SetTitle(title string) *Component {
	p.HeaderTitle = title

	return p
}

// 表头标题
func (p *Component) SetHeaderTitle(headerTitle string) *Component {
	p.HeaderTitle = headerTitle

	return p
}

// 是否显示搜索表单，传入对象时为搜索表单的配置
func (p *Component) SetSearch(callback interface{}) interface{} {
	// callback(p.Search)

	return p.Search
}

// 搜索表单的配置
func (p *Component) SetSearches(search interface{}) *Component {
	p.Search = search

	return p
}

// 批量设置表格列
func (p *Component) SetColumns(columns interface{}) *Component {
	p.Columns = columns

	return p
}

// 批量操作选择项
func (p *Component) SetRowSelection(rowSelection interface{}) *Component {
	p.RowSelection = rowSelection

	return p
}

// table 工具栏，设为 false 时不显示,{ fullScreen: true, reload: true ,setting: true}
func (p *Component) SetOptions(options map[string]bool) *Component {
	p.Options = options

	return p
}

// 转化 moment 格式数据为特定类型，false 不做转化,"string" | "number" | false
func (p *Component) SetDateFormatter(dateFormatter string) *Component {
	p.DateFormatter = dateFormatter

	return p
}

// 空值时的显示，不设置 则默认显示 -
func (p *Component) SetColumnEmptyText(columnEmptyText string) *Component {
	p.ColumnEmptyText = columnEmptyText

	return p
}

// 透传 ProUtils 中的 ListToolBar 配置项
func (p *Component) SetToolBar(toolBar interface{}) *Component {
	p.ToolBar = toolBar

	return p
}

// 树形栏
func (p *Component) SetTreeBar(treeBar interface{}) *Component {
	p.TreeBar = treeBar

	return p
}

// 表格的批量操作
func (p *Component) SetBatchActions(batchActions interface{}) *Component {
	p.BatchActions = batchActions

	return p
}

// 自定义表格的主体函数
func (p *Component) SetTableExtraRender(tableExtraRender interface{}) *Component {
	p.TableExtraRender = tableExtraRender

	return p
}

// 配置展开行
func (p *Component) SetExpandable(expandable *Expandable) *Component {
	p.Expandable = expandable

	return p
}

// 设置表格滚动
func (p *Component) SetScroll(scroll interface{}) *Component {
	p.Scroll = scroll

	return p
}

// 设置表格滚动
func (p *Component) SetStriped(striped bool) *Component {
	p.Striped = striped

	return p
}

// 表格数据
func (p *Component) SetDatasource(datasource interface{}) *Component {
	p.Datasource = datasource

	return p
}

// 表格分页
func (p *Component) SetPagination(current int, pageSize int, total int, defaultCurrent int) *Component {
	p.Pagination = map[string]int{
		"current":        current,
		"pageSize":       pageSize,
		"total":          total,
		"defaultCurrent": defaultCurrent,
	}

	return p
}

// 是否轮询
func (p *Component) SetPolling(polling int) *Component {
	p.Polling = polling

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "table"

	return p
}
