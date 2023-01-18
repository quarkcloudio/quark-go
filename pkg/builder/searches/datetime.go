package searches

type Datetime struct {
	Search
}

// 初始化
func (p *Datetime) ParentInit() interface{} {
	p.Component = "datetime"

	return p
}
