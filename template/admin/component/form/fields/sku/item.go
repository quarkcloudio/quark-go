package sku

type Item struct {
	Title      string      `json:"title,omitempty"`     // 列标题
	DataIndex  string      `json:"dataIndex,omitempty"` // 列索引
	Width      interface{} `json:"width,omitempty"`     // 列宽
	Fixed      interface{} `json:"fixed,omitempty"`     // （IE 下无效）列是否固定，可选 true (等效于 left) left rightr
	Editable   interface{} `json:"editable,omitempty"`
	ValueEnum  interface{} `json:"valueEnum,omitempty"`
	ValueType  string      `json:"valueType,omitempty"`
	FieldProps interface{} `json:"fieldProps,omitempty"`
}

func NewItem() *Item {
	return &Item{}
}

func (p *Item) SetTitle(title string) *Item {
	p.Title = title
	return p
}

func (p *Item) SetDataIndex(dataIndex string) *Item {
	p.DataIndex = dataIndex
	return p
}

func (p *Item) SetWidth(width interface{}) *Item {
	p.Width = width
	return p
}

func (p *Item) SetFixed(fixed interface{}) *Item {
	p.Fixed = fixed
	return p
}

func (p *Item) SetEditable(editable interface{}) *Item {
	p.Editable = editable
	return p
}

func (p *Item) SetValueEnum(valueEnum interface{}) *Item {
	p.ValueEnum = valueEnum
	return p
}

func (p *Item) SetValueType(valueType string) *Item {
	p.ValueType = valueType
	return p
}

func (p *Item) SetFieldProps(pieldProps interface{}) *Item {
	p.FieldProps = pieldProps
	return p
}
