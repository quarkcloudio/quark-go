package order

type Table struct {
	DataIndex  string        `json:"dataIndex"`
	Columns    []Column      `json:"columns"`
	DataSource []interface{} `json:"dataSource"`
}

// 初始化组件
func NewTable() *Table {
	return &Table{}
}

// 设置数据源
func (p *Table) SetColumn(columns []Column) *Table {
	p.Columns = columns
	return p
}

// 设置数据索引
func (p *Table) SetDataIndex(dataIndex string) *Table {
	p.DataIndex = dataIndex
	return p
}

// 设置数据源
func (p *Table) SetDataSource(dataSource []interface{}) *Table {
	p.DataSource = dataSource
	return p
}
