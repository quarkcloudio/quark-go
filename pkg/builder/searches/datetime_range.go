package searches

type DatetimeRange struct {
	Search
}

// 初始化
func (p *DatetimeRange) ParentInit() interface{} {
	p.Component = "datetimeRangeField"

	return p
}
