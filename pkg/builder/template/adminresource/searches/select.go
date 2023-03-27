package searches

type SelectOption struct {
	Label    string      `json:"label"`
	Value    interface{} `json:"value"`
	Disabled bool        `json:"disabled,omitempty"`
}

type Select struct {
	Search
}

// 初始化
func (p *Select) ParentInit() interface{} {
	p.Component = "selectField"

	return p
}

// 设置Option
func (p *Select) Option(value interface{}, label string) *SelectOption {

	return &SelectOption{
		Value: value,
		Label: label,
	}
}
