package table

import (
	"reflect"

	"github.com/gobeam/stringy"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/component"
)

type FieldNames struct {
	Title    string `json:"title"`
	Key      string `json:"key"`
	Children string `json:"children"`
}

type TreeData struct {
	Checkable       bool        `json:"checkable,omitempty"`       // 当树为 checkable 时，设置独立节点是否展示 Checkbox
	DisableCheckbox bool        `json:"disableCheckbox,omitempty"` // 禁掉 checkbox
	Disabled        bool        `json:"disabled,omitempty"`        // 禁掉响应
	Icon            interface{} `json:"icon,omitempty"`            // 自定义图标。可接收组件，props 为当前节点 props
	IsLeaf          bool        `json:"isLeaf,omitempty"`          // 设置为叶子节点 (设置了 loadData 时有效)。为 false 时会强制将其作为父节点
	Key             interface{} `json:"key"`                       // 被树的 (default)ExpandedKeys / (default)CheckedKeys / (default)SelectedKeys 属性所用。注意：整个树范围内的所有节点的 key 值不能重复！
	Selectable      bool        `json:"selectable,omitempty"`      // 设置节点是否可被选中
	Title           string      `json:"title"`                     // 标题
	Children        []TreeData  `json:"children,omitempty"`        // 子节点
}

type TreeBar struct {
	component.Element
	Component           string                 `json:"component"`                     // 组件名称
	Name                string                 `json:"name,omitempty"`                // 字段名，支持数组
	AutoExpandParent    bool                   `json:"autoExpandParent,omitempty"`    // 是否自动展开父节点
	BockNode            bool                   `json:"blockNode,omitempty"`           // 是否节点占据一行
	Checkable           bool                   `json:"checkable,omitempty"`           // 节点前添加 Checkbox 复选框
	CheckedKeys         []interface{}          `json:"checkedKeys,omitempty"`         // （受控）选中复选框的树节点（注意：父子节点有关联，如果传入父节点 key，则子节点自动选中；相应当子节点 key 都传入，父节点也自动选中。当设置 checkable 和 checkStrictly，它是一个有checked和halfChecked属性的对象，并且父子节点的选中与否不再关联
	CheckStrictly       bool                   `json:"checkStrictly,omitempty"`       // checkable 状态下节点选择完全受控（父子节点选中状态不再关联）
	DefaultCheckedKeys  []interface{}          `json:"defaultCheckedKeys,omitempty"`  // 默认选中复选框的树节点
	DefaultExpandAll    bool                   `json:"defaultExpandAll,omitempty"`    // 默认展开所有树节点
	DefaultExpandedKeys []interface{}          `json:"defaultExpandedKeys,omitempty"` // 默认展开指定的树节点
	DefaultExpandParent bool                   `json:"defaultExpandParent,omitempty"` // 默认展开父节点
	DefaultSelectedKeys []interface{}          `json:"defaultSelectedKeys,omitempty"` // 默认选中的树节点
	DefaultValue        interface{}            `json:"defaultValue,omitempty"`        // 默认选中的选项
	Disabled            bool                   `json:"disabled,omitempty"`            // 整组失效
	Draggable           bool                   `json:"draggable,omitempty"`           // 设置节点可拖拽，可以通过 icon: false 关闭拖拽提示图标
	ExpandedKeys        []interface{}          `json:"expandedKeys,omitempty"`        // （受控）展开指定的树节点
	FieldNames          *FieldNames            `json:"fieldNames,omitempty"`          // 自定义 options 中 label value children 的字段
	Height              int                    `json:"height,omitempty"`              // 设置虚拟滚动容器高度，设置后内部节点不再支持横向滚动
	Icon                interface{}            `json:"icon,omitempty"`                // 自定义树节点图标
	Multiple            bool                   `json:"multiple,omitempty"`            // 支持点选多个节点（节点本身）
	Placeholder         string                 `json:"placeholder,omitempty"`         // 占位文本
	RootClassName       string                 `json:"rootClassName,omitempty"`       // 添加在 Tree 最外层的 className
	RootStyle           interface{}            `json:"rootStyle,omitempty"`           // 添加在 Tree 最外层的 style
	Selectable          bool                   `json:"selectable,omitempty"`          // 是否可选中
	SelectedKeys        []interface{}          `json:"selectedKeys,omitempty"`        // （受控）设置选中的树节点
	ShowIcon            bool                   `json:"showIcon,omitempty"`            // 是否展示 TreeNode title 前的图标，没有默认样式，如设置为 true，需要自行定义图标相关样式
	ShowLine            bool                   `json:"showLine,omitempty"`            // 是否展示连接线
	SwitcherIcon        interface{}            `json:"switcherIcon,omitempty"`        // 自定义树节点的展开/折叠图标
	TreeData            []TreeData             `json:"treeData,omitempty"`            // treeNodes 数据，如果设置则不需要手动构造 TreeNode 节点（value 在整个树范围内唯一）
	Value               interface{}            `json:"value,omitempty"`               // 指定当前选中的条目，多选时为一个数组。（value 数组引用未变化时，Select 不会更新）
	Virtual             bool                   `json:"virtual,omitempty"`             // 设置 false 时关闭虚拟滚动
	Style               map[string]interface{} `json:"style,omitempty"`               // 自定义样式
}

// 初始化
func (p *TreeBar) Init() *TreeBar {
	p.Component = "treeBar"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Name = "treeBarSelectedKeys"
	p.Placeholder = "请输入搜索内容"
	p.DefaultExpandAll = true
	p.ShowLine = true
	return p
}

// 字段名，支持数组
func (p *TreeBar) SetName(name string) *TreeBar {
	p.Name = name
	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *TreeBar) SetWidth(width interface{}) *TreeBar {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 是否自动展开父节点
func (p *TreeBar) SetAutoExpandParent(autoExpandParent bool) *TreeBar {
	p.AutoExpandParent = autoExpandParent

	return p
}

// 是否节点占据一行
func (p *TreeBar) SetBockNode(blockNode bool) *TreeBar {
	p.BockNode = blockNode

	return p
}

// 节点前添加 Checkbox 复选框
func (p *TreeBar) SetCheckable(checkable bool) *TreeBar {
	p.Checkable = checkable

	return p
}

// （受控）选中复选框的树节点（注意：父子节点有关联，如果传入父节点 key，则子节点自动选中；相应当子节点 key 都传入，父节点也自动选中。当设置 checkable 和 checkStrictly，它是一个有checked和halfChecked属性的对象，并且父子节点的选中与否不再关联
func (p *TreeBar) SetCheckedKeys(checkedKeys []interface{}) *TreeBar {
	p.CheckedKeys = checkedKeys

	return p
}

// checkable 状态下节点选择完全受控（父子节点选中状态不再关联）
func (p *TreeBar) SetCheckStrictly(checkStrictly bool) *TreeBar {
	p.CheckStrictly = checkStrictly

	return p
}

// 默认选中复选框的树节点
func (p *TreeBar) SetDefaultCheckedKeys(defaultCheckedKeys []interface{}) *TreeBar {
	p.DefaultCheckedKeys = defaultCheckedKeys

	return p
}

// 默认展开所有树节点
func (p *TreeBar) SetDefaultExpandAll(defaultExpandAll bool) *TreeBar {
	p.DefaultExpandAll = defaultExpandAll

	return p
}

// 默认展开指定的树节点
func (p *TreeBar) SetDefaultExpandedKeys(defaultExpandedKeys []interface{}) *TreeBar {
	p.DefaultExpandedKeys = defaultExpandedKeys

	return p
}

// 默认展开父节点
func (p *TreeBar) SetDefaultExpandParent(defaultExpandParent bool) *TreeBar {
	p.DefaultExpandParent = defaultExpandParent

	return p
}

// 默认选中的树节点
func (p *TreeBar) SetDefaultSelectedKeys(defaultSelectedKeys []interface{}) *TreeBar {
	p.DefaultSelectedKeys = defaultSelectedKeys

	return p
}

// 设置节点可拖拽，可以通过 icon: false 关闭拖拽提示图标
func (p *TreeBar) SetDraggable(draggable bool) *TreeBar {
	p.Draggable = draggable

	return p
}

// （受控）展开指定的树节点
func (p *TreeBar) SetExpandedKeys(expandedKeys []interface{}) *TreeBar {
	p.ExpandedKeys = expandedKeys

	return p
}

// 自定义 options 中 label value children 的字段
func (p *TreeBar) SetFieldNames(fieldNames *FieldNames) *TreeBar {
	p.FieldNames = fieldNames

	return p
}

// 设置虚拟滚动容器高度，设置后内部节点不再支持横向滚动
func (p *TreeBar) SetHeight(height int) *TreeBar {
	p.Height = height

	return p
}

// 自定义树节点图标
func (p *TreeBar) SetIcon(icon interface{}) *TreeBar {
	p.Icon = icon

	return p
}

// 支持点选多个节点（节点本身）
func (p *TreeBar) SetMultiple(multiple bool) *TreeBar {
	p.Multiple = multiple

	return p
}

// 占位文本
func (p *TreeBar) SetPlaceholder(placeholder string) *TreeBar {
	p.Placeholder = placeholder

	return p
}

// 添加在 Tree 最外层的 className
func (p *TreeBar) SetRootClassName(rootClassName string) *TreeBar {
	p.RootClassName = rootClassName

	return p
}

// 添加在 Tree 最外层的 style
func (p *TreeBar) SetRootStyle(rootStyle interface{}) *TreeBar {
	p.RootStyle = rootStyle

	return p
}

// 是否可选中
func (p *TreeBar) SetSelectable(selectable bool) *TreeBar {
	p.Selectable = selectable

	return p
}

// 设置选中的树节点
func (p *TreeBar) SetSelectedKeys(selectedKeys []interface{}) *TreeBar {
	p.SelectedKeys = selectedKeys

	return p
}

// 是否展示 TreeNode title 前的图标，没有默认样式，如设置为 true，需要自行定义图标相关样式
func (p *TreeBar) SetShowIcon(showIcon bool) *TreeBar {
	p.ShowIcon = showIcon

	return p
}

// 是否展示连接线
func (p *TreeBar) SetShowLine(showLine bool) *TreeBar {
	p.ShowLine = showLine

	return p
}

// 自定义树节点的展开/折叠图标
func (p *TreeBar) SetSwitcherIcon(switcherIcon interface{}) *TreeBar {
	p.SwitcherIcon = switcherIcon

	return p
}

// buildTree 使用反射构建树结构
func (p *TreeBar) buildTree(items interface{}, pid int, parentKeyName string, keyName string, titleName string) []TreeData {
	var tree []TreeData

	// 通过反射获取切片的值
	v := reflect.ValueOf(items)

	// 确保传入的是切片类型
	if v.Kind() != reflect.Slice {
		return nil
	}

	// 遍历切片中的每个元素
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)

		if item.Kind() == reflect.Ptr && item.Elem().IsValid() {
			item = item.Elem() // 解引用指针
		}

		// 使用反射获取字段
		keyField := item.FieldByName(
			stringy.
				New(keyName).
				CamelCase("?", "").Get(),
		)

		parentKeyField := item.FieldByName(
			stringy.
				New(parentKeyName).
				CamelCase("?", "").Get(),
		)
		titleField := item.FieldByName(
			stringy.
				New(titleName).
				CamelCase("?", "").Get(),
		)

		// 确保字段存在并且类型正确
		if !keyField.IsValid() || !parentKeyField.IsValid() || !titleField.IsValid() {
			continue
		}

		// 断言字段类型
		key := int(keyField.Int())
		parentKey := int(parentKeyField.Int())
		title := titleField.String()

		// 如果当前项的 Pid 与传入的 pid 匹配
		if parentKey == pid {
			// 递归查找子节点
			children := p.buildTree(items, key, parentKeyName, keyName, titleName)

			// 构建级联选择框的选项
			option := TreeData{
				Key:      key,
				Title:    title,
				Children: children,
			}

			tree = append(tree, option)
		}
	}
	return tree
}

func (p *TreeBar) ListToTreeData(list interface{}, rootId int, parentKeyName string, keyName string, titleName string) []TreeData {
	return p.buildTree(list, rootId, parentKeyName, keyName, titleName)
}

// 可选项数据源
//
//	SetTreeData([]tree.TreeData {{ Key :"zhejiang", Title:"Zhejiang"}})
//
// 或者
//
// SetTreeData(options, "parent_key_name", "key_name", "title_name")
//
// 或者
//
// SetTreeData(options, 0, "parent_key_name", "key_name", "title_name")
func (p *TreeBar) SetTreeData(treeData ...interface{}) *TreeBar {
	if len(treeData) == 1 {
		getOptions, ok := treeData[0].([]TreeData)
		if ok {
			p.TreeData = getOptions
			return p
		}
	}
	if len(treeData) == 4 {
		p.TreeData = p.ListToTreeData(treeData[0], 0, treeData[1].(string), treeData[2].(string), treeData[3].(string))
	}
	if len(treeData) == 5 {
		p.TreeData = p.ListToTreeData(treeData[0], treeData[1].(int), treeData[2].(string), treeData[3].(string), treeData[4].(string))
	}
	return p
}

// 自定义样式
func (p *TreeBar) SetStyle(style map[string]interface{}) *TreeBar {
	p.Style = style

	return p
}
