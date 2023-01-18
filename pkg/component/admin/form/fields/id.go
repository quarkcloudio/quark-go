package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type ID struct {
	Item
	OnIndexDisplayed  bool `json:"onIndexDisplayed"`
	OnDetailDisplayed bool `json:"onDetailDisplayed"`
	OnFormDisplayed   bool `json:"onFormDisplayed"`
	OnExportDisplayed bool `json:"onExportDisplayed"`
}

// 初始化
func (p *ID) Init() *ID {
	p.Component = "idField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.OnIndexDisplayed = true
	p.ShowOnImport = false

	return p
}

// 在列表页是否显示字段
func (p *ID) SetOnIndexDisplayed(displayed bool) *ID {
	p.OnIndexDisplayed = displayed
	return p
}

// 在详情页是否显示字段
func (p *ID) SetOnDetailDisplayed(displayed bool) *ID {
	p.OnDetailDisplayed = displayed
	return p
}

// 在表单页是否显示控件
func (p *ID) SetOnFormDisplayed(displayed bool) *ID {
	p.OnFormDisplayed = displayed
	return p
}

// 在导出页是否显示字段
func (p *ID) SetOnExportDisplayed(displayed bool) *ID {
	p.OnExportDisplayed = displayed
	return p
}
