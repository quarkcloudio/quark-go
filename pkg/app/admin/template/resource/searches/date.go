package searches

type Date struct {
	Search
}

// 初始化
func (p *Date) ParentInit() interface{} {
	p.Component = "dateField"

	return p
}
