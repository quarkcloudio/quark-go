package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type TreeSelect struct {
	Item
	AllowClear              bool        `json:"allowClear"`
	AutoClearSearchValue    bool        `json:"autoClearSearchValue"`
	Bordered                bool        `json:"bordered"`
	DefaultValue            interface{} `json:"defaultValue"`
	Disabled                bool        `json:"disabled"`
	PopupClassName          string      `json:"popupClassName"`
	ListHeight              int         `json:"listHeight"`
	MaxTagCount             int         `json:"maxTagCount,omitempty"`
	MaxTagPlaceholder       interface{} `json:"maxTagPlaceholder,omitempty"`
	MaxTagTextLength        int         `json:"maxTagTextLength,omitempty"`
	Multiple                bool        `json:"multiple"`
	NotFoundContent         string      `json:"notFoundContent,omitempty"`
	Placeholder             string      `json:"placeholder"`
	Placement               string      `json:"placement"`
	ShowArrow               bool        `json:"showArrow"`
	ShowSearch              bool        `json:"showSearch"`
	Size                    string      `json:"size"`
	Status                  string      `json:"status"`
	SuffixIcon              interface{} `json:"suffixIcon,omitempty"`
	TreeCheckable           bool        `json:"treeCheckable"`
	TreeData                interface{} `json:"treeData"`
	TreeDataSimpleMode      bool        `json:"treeDataSimpleMode"`
	TreeDefaultExpandAll    bool        `json:"treeDefaultExpandAll"`
	TreeDefaultExpandedKeys []string    `json:"treeDefaultExpandedKeys,omitempty"`
	TreeExpandAction        interface{} `json:"treeExpandAction,omitempty"`
	TreeExpandedKeys        []string    `json:"treeExpandedKeys,omitempty"`
	TreeIcon                bool        `json:"treeIcon,omitempty"`
	TreeLine                bool        `json:"treeLine,omitempty"`
	Value                   interface{} `json:"value,omitempty"`
	Virtual                 bool        `json:"virtual,omitempty"`
	Width                   interface{} `json:"width,omitempty"`
	MaxLength               int         `json:"maxLength,omitempty"`
}

// 初始化
func (p *TreeSelect) Init() *TreeSelect {
	p.Component = "treeSelectField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Width = 400
	p.TreeDefaultExpandAll = true
	p.ShowArrow = true
	p.TreeLine = true

	return p
}

// 显示清除按钮
func (p *TreeSelect) SetAllowClear(allowClear bool) *TreeSelect {
	p.AllowClear = allowClear

	return p
}

// 当多选模式下值被选择，自动清空搜索框
func (p *TreeSelect) SetAutoClearSearchValue(autoClearSearchValue bool) *TreeSelect {
	p.AutoClearSearchValue = autoClearSearchValue

	return p
}

// 是否显示边框
func (p *TreeSelect) SetBordered(bordered bool) *TreeSelect {
	p.Bordered = bordered

	return p
}

// 指定默认选中的条目
func (p *TreeSelect) SetDefaultValue(defaultValue interface{}) *TreeSelect {
	p.DefaultValue = defaultValue

	return p
}

// 是否禁用
func (p *TreeSelect) SetDisabled(disabled bool) *TreeSelect {
	p.Disabled = disabled

	return p
}

// 下拉菜单的 className 属性
func (p *TreeSelect) SetPopupClassName(popupClassName string) *TreeSelect {
	p.PopupClassName = popupClassName

	return p
}

// 设置弹窗滚动高度
func (p *TreeSelect) SetListHeight(listHeight int) *TreeSelect {
	p.ListHeight = listHeight

	return p
}

// 最多显示多少个 tag，响应式模式会对性能产生损耗
func (p *TreeSelect) SetMaxTagCount(maxTagCount int) *TreeSelect {
	p.MaxTagCount = maxTagCount

	return p
}

// 隐藏 tag 时显示的内容
func (p *TreeSelect) SetMaxTagPlaceholder(maxTagPlaceholder interface{}) *TreeSelect {
	p.MaxTagPlaceholder = maxTagPlaceholder

	return p
}

// 最大显示的 tag 文本长度
func (p *TreeSelect) SetMaxTagTextLength(maxTagTextLength int) *TreeSelect {
	p.MaxTagTextLength = maxTagTextLength

	return p
}

// 支持多选（当设置 treeCheckable 时自动变为 true）
func (p *TreeSelect) SetMultiple(multiple bool) *TreeSelect {
	p.Multiple = multiple

	return p
}

// 当下拉列表为空时显示的内容
func (p *TreeSelect) SetNotFoundContent(notFoundContent string) *TreeSelect {
	p.NotFoundContent = notFoundContent

	return p
}

// 选择框默认文字
func (p *TreeSelect) SetPlaceholder(placeholder string) *TreeSelect {
	p.Placeholder = placeholder

	return p
}

// 选择框弹出的位置, bottomLeft bottomRight topLeft topRight
func (p *TreeSelect) SetPlacement(placement string) *TreeSelect {
	p.Placement = placement

	return p
}

// 是否显示 suffixIcon，单选模式下默认 true
func (p *TreeSelect) SetShowArrow(showArrow bool) *TreeSelect {
	p.ShowArrow = showArrow

	return p
}

// 是否支持搜索框
func (p *TreeSelect) SetShowSearch(showSearch bool) *TreeSelect {
	p.ShowSearch = showSearch

	return p
}

// 选择框大小
func (p *TreeSelect) SetSize(size string) *TreeSelect {
	p.Size = size

	return p
}

// 设置校验状态,'error' | 'warning'
func (p *TreeSelect) SetStatus(status string) *TreeSelect {
	p.Status = status

	return p
}

// 自定义的选择框后缀图标, 多选模式下必须同时设置 showArrow 为 true
func (p *TreeSelect) SetSuffixIcon(suffixIcon interface{}) *TreeSelect {
	p.SuffixIcon = suffixIcon

	return p
}

// 显示 Checkbox
func (p *TreeSelect) SetTreeCheckable(treeCheckable bool) *TreeSelect {
	p.TreeCheckable = treeCheckable

	return p
}

// 设置树选择数据
func (p *TreeSelect) SetData(data interface{}) *TreeSelect {
	p.TreeData = data

	return p
}

// 使用简单格式的 treeData，具体设置参考可设置的类型 (此时 treeData 应变为这样的数据结构: [{id:1, pId:0, value:'1', title:"test1",...},...]， pId 是父节点的 id)
func (p *TreeSelect) SetTreeDataSimpleMode(treeDataSimpleMode bool) *TreeSelect {
	p.TreeDataSimpleMode = treeDataSimpleMode

	return p
}

// 默认展开所有树节点
func (p *TreeSelect) SetTreeDefaultExpandAll(treeDefaultExpandAll bool) *TreeSelect {
	p.TreeDefaultExpandAll = treeDefaultExpandAll

	return p
}

// 默认展开所有树节点
func (p *TreeSelect) SetTreeDefaultExpandedKeys(treeDefaultExpandedKeys []string) *TreeSelect {
	p.TreeDefaultExpandedKeys = treeDefaultExpandedKeys

	return p
}

// 是否展示 TreeNode title 前的图标，没有默认样式，如设置为 true，需要自行定义图标相关样式
func (p *TreeSelect) SetTreeIcon(treeIcon bool) *TreeSelect {
	p.TreeIcon = treeIcon

	return p
}

// 是否展示线条样式
func (p *TreeSelect) SetTreeLine(treeLine bool) *TreeSelect {
	p.TreeLine = treeLine

	return p
}

// 指定当前选中的条目
func (p *TreeSelect) SetValue(value interface{}) *TreeSelect {
	p.Value = value

	return p
}

// 设置 false 时关闭虚拟滚动
func (p *TreeSelect) SetVirtual(virtual bool) *TreeSelect {
	p.Virtual = virtual

	return p
}

// 设置宽度
func (p *TreeSelect) SetWidth(width interface{}) *TreeSelect {
	p.Width = width

	return p
}

// 最大字符数
func (p *TreeSelect) SetMaxLength(maxLength int) *TreeSelect {
	p.MaxLength = maxLength

	return p
}
