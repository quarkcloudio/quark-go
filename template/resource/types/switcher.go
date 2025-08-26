package types

type Switcher interface {
	Actioner

	// 选中时的内容
	GetCheckedChildren() interface{}

	// 自定义的选择框后缀图标
	GetUnCheckedChildren() interface{}

	// 获取字段名称
	GetFieldName() interface{}

	// 获取字段值
	GetFieldValue() interface{}
}
