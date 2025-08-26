package rule

type Rule struct {
	Name              string        `json:"-"`                      // 需要验证的字段名称
	RuleType          string        `json:"-"`                      // 规则类型，max | min | unique | required
	DefaultField      interface{}   `json:"defaultField,omitempty"` // 仅在 type 为 array 类型时有效，用于指定数组元素的校验规
	Enum              []interface{} `json:"enum,omitempty"`         // 是否匹配枚举中的值（需要将 type 设置为 enum）
	Fields            interface{}   `json:"fields,omitempty"`       // 仅在 type 为 array 或 object 类型时有效，用于指定子元素的校验规则
	Len               int           `json:"len,omitempty"`          // string 类型时为字符串长度；number 类型时为确定数字； array 类型时为数组长度
	Max               int           `json:"max,omitempty"`          // 必须设置 type：string 类型为字符串最大长度；number 类型时为最大值；array 类型时为数组最大长度
	Message           string        `json:"message"`                // 错误信息，不设置时会通过模板自动生成
	Min               int           `json:"min,omitempty"`          // 必须设置 type：string 类型为字符串最小长度；number 类型时为最小值；array 类型时为数组最小长度
	Pattern           string        `json:"pattern,omitempty"`      // 正则表达式匹配
	Required          bool          `json:"required,omitempty"`     // 是否为必选字段
	UniqueTable       string        `json:"-"`                      // type：unique时，指定验证的表名
	UniqueTableField  string        `json:"-"`                      // type：unique时，指定需验证表中的字段
	UniqueIgnoreValue string        `json:"-"`                      // type：unique时，忽略符合条件验证的列，例如：{id}
	Type              string        `json:"type,omitempty"`         // 字段类型，string | number | boolean | method | regexp | integer | float | array | object | enum | date | url | hex | email | any
}

// 初始化
func New() *Rule {
	p := &Rule{}

	return p
}

// 转换前端验证规则，剔除前端不支持的unique
func ConvertToFrontendRules(rules []Rule) []Rule {
	var newRules []Rule

	for _, rule := range rules {
		if rule.RuleType != "unique" {
			newRules = append(newRules, rule)
		}
	}

	return newRules
}

// 必须设置 type：string 类型；为字符串最大长度；number 类型时为最大值；array 类型时为数组最大长度
func Max(max int, message string) Rule {
	p := &Rule{}

	return *p.SetMax(max).SetMessage(message)
}

// 必须设置 type：string 类型为字符串最小长度；number 类型时为最小值；array 类型时为数组最小长度
func Min(min int, message string) Rule {
	p := &Rule{}

	return *p.SetMin(min).SetMessage(message)
}

// 正则表达式匹配
func Regexp(pattern string, message string) Rule {
	p := &Rule{}

	return *p.SetRegexp(pattern).SetMessage(message)
}

// 必须为字符串
func String(message string) Rule {
	p := &Rule{}

	return *p.SetString().SetMessage(message)
}

// 必须为数字
func Number(message string) Rule {
	p := &Rule{}

	return *p.SetNumber().SetMessage(message)
}

// 必须为布尔类型
func Boolean(message string) Rule {
	p := &Rule{}

	return *p.SetBoolean().SetMessage(message)
}

// 必须为整型
func Integer(message string) Rule {
	p := &Rule{}

	return *p.SetInteger().SetMessage(message)
}

// 必须为浮点型
func Float(message string) Rule {
	p := &Rule{}

	return *p.SetFloat().SetMessage(message)
}

// 必须为邮箱字段
func Email(message string) Rule {
	p := &Rule{}

	return *p.SetEmail().SetMessage(message)
}

// 必须为链接
func Url(message string) Rule {
	p := &Rule{}

	return *p.SetUrl().SetMessage(message)
}

// 必须为手机号
func Phone(message string) Rule {
	p := &Rule{}

	return *p.SetPhone().SetMessage(message)
}

// 是否为必选字段
func Required(message string) Rule {
	p := &Rule{}

	return *p.SetRequired().SetMessage(message)
}

// 设置unique验证类型，插入数据时：Unique("admins", "username", "用户名已存在")，更新数据时：Unique("admins", "username", "{id}", "用户名已存在")
func Unique(unique ...string) Rule {
	var (
		uniqueTable       string
		uniqueTableField  string
		uniqueIgnoreValue string
		message           string
	)

	p := &Rule{}
	if len(unique) == 3 {
		uniqueTable = unique[0]
		uniqueTableField = unique[1]
		message = unique[2]

		p.SetUnique(uniqueTable, uniqueTableField)
	}

	if len(unique) == 4 {
		uniqueTable = unique[0]
		uniqueTableField = unique[1]
		uniqueIgnoreValue = unique[2]
		message = unique[3]

		p.SetUnique(uniqueTable, uniqueTableField, uniqueIgnoreValue)
	}

	p.SetMessage(message)

	return *p
}

// 需要验证的字段名称
func (p *Rule) SetName(name string) *Rule {
	p.Name = name

	return p
}

// 必须设置 type：string 类型；为字符串最大长度；number 类型时为最大值；array 类型时为数组最大长度
func (p *Rule) SetMax(max int) *Rule {
	p.Max = max

	return p.SetRuleType("max")
}

// 错误信息，不设置时会通过模板自动生成
func (p *Rule) SetMessage(message string) *Rule {
	p.Message = message

	return p
}

// 必须设置 type：string 类型为字符串最小长度；number 类型时为最小值；array 类型时为数组最小长度
func (p *Rule) SetMin(min int) *Rule {
	p.Min = min

	return p.SetRuleType("min")
}

// 正则表达式匹配
func (p *Rule) SetRegexp(pattern string) *Rule {
	p.Pattern = pattern

	return p.SetRuleType("regexp")
}

// 是否为必选字段
func (p *Rule) SetRequired() *Rule {
	p.Required = true

	return p.SetRuleType("required")
}

// 必须为字符串
func (p *Rule) SetString() *Rule {
	p.Type = "string"

	return p.SetRuleType("string")
}

// 必须为数字
func (p *Rule) SetNumber() *Rule {
	p.Type = "number"

	return p.SetRuleType("number")
}

// 必须为布尔类型
func (p *Rule) SetBoolean() *Rule {
	p.Type = "boolean"

	return p.SetRuleType("boolean")
}

// 必须为整型
func (p *Rule) SetInteger() *Rule {
	p.Type = "integer"

	return p.SetRuleType("integer")
}

// 必须为浮点型
func (p *Rule) SetFloat() *Rule {
	p.Type = "float"

	return p.SetRuleType("float")
}

// 必须为邮箱字段
func (p *Rule) SetEmail() *Rule {
	p.Type = "email"

	return p.SetRuleType("email")
}

// 必须为链接
func (p *Rule) SetUrl() *Rule {
	p.Type = "url"

	return p.SetRuleType("url")
}

// 必须为手机号
func (p *Rule) SetPhone() *Rule {

	return p.SetRegexp(`/^1[3-9]\d{9}$/`)
}

// 设置unique验证类型，插入数据：SetUnique("admins","username")，更新数据：SetUnique("admins","username","{id}")
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

	return p.SetRuleType("unique")
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

// 字段类型，string | number | boolean | url | email
func (p *Rule) SetType(ruleType string) *Rule {
	p.Type = ruleType

	return p
}

// 规则类型，max | min | unique | required
func (p *Rule) SetRuleType(ruleType string) *Rule {
	p.RuleType = ruleType

	return p
}
