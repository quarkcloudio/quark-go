package searches

type Select struct {
	Search
}

// 初始化
func (p *Select) ParentInit() interface{} {
	p.Component = "select"

	return p
}
