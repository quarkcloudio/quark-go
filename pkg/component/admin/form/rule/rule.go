package rule

type Rule struct {
	Max               int    `json:"max,omitempty"`      // 必须设置 type：string 类型为字符串最大长度；number 类型时为最大值；array 类型时为数组最大长度
	Message           string `json:"message"`            // 错误信息，不设置时会通过模板自动生成
	Min               int    `json:"min,omitempty"`      // 必须设置 type：string 类型为字符串最小长度；number 类型时为最小值；array 类型时为数组最小长度
	Required          bool   `json:"required,omitempty"` // 是否为必选字段
	UniqueTable       string `json:"-"`                  // type：unique时，指定验证的表名
	UniqueTableField  string `json:"-"`                  // type：unique时，指定需验证表中的字段
	UniqueIgnoreValue string `json:"-"`                  // type：unique时，忽略符合条件验证的列，例如：{id}
	Type              string `json:"type"`               // 类型，string |number |boolean |url | email | unique
}

// 初始化
func New() *Rule {
	p := &Rule{}

	p.Type = "string"

	return p
}

// 必须设置 type：string 类型；为字符串最大长度；number 类型时为最大值；array 类型时为数组最大长度
func (p *Rule) SetMax(max int) *Rule {
	p.Max = max

	return p
}

// 错误信息，不设置时会通过模板自动生成
func (p *Rule) SetMessage(message string) *Rule {
	p.Message = message

	return p
}

// 必须设置 type：string 类型为字符串最小长度；number 类型时为最小值；array 类型时为数组最小长度
func (p *Rule) SetMin(min int) *Rule {
	p.Min = min

	return p
}

// 是否为必选字段
func (p *Rule) SetRequired(required bool) *Rule {
	p.Required = required

	return p
}

// 设置unique验证类型，SetUnique("admins","username")|SetUnique("admins","username","{id}")
func (p *Rule) SetUnique(unique ...string) *Rule {
	p.Type = "unique"

	if len(unique) == 2 {
		p.UniqueTable = unique[0]
		p.UniqueTableField = unique[1]
	}

	if len(unique) == 3 {
		p.UniqueTable = unique[0]
		p.UniqueTableField = unique[1]
		p.UniqueIgnoreValue = unique[2]
	}

	return p
}

// type：unique时，指定验证的表名
func (p *Rule) SetUniqueTable(uniqueTable string) *Rule {
	p.UniqueTable = uniqueTable

	return p
}

// type：unique时，指定验证的表名
func (p *Rule) SetUniqueTableField(uniqueTableField string) *Rule {
	p.UniqueTableField = uniqueTableField

	return p
}

// type：unique时，忽略符合条件验证的列，例如：{id}
func (p *Rule) SetUniqueIgnoreValue(uniqueIgnoreValue string) *Rule {
	p.UniqueIgnoreValue = uniqueIgnoreValue

	return p
}

// 类型，string | number | boolean | url | email | unique
func (p *Rule) SetType(ruleType string) *Rule {
	p.Type = ruleType

	return p
}
