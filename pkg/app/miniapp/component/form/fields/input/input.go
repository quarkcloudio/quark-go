package input

import "github.com/quarkcms/quark-go/v2/pkg/app/miniapp/component/component"

type Component struct {
	component.Element
	Name              string      `json:"name"`
	Required          bool        `json:"required"`
	Prop              string      `json:"prop"`
	Rules             interface{} `json:"rules,omitempty"`
	Label             string      `json:"label,omitempty"`
	LabelWidth        int         `json:"labelWidth,omitempty"`
	LabelAlign        string      `json:"labelAlign,omitempty"`
	BodyAlign         string      `json:"bodyAlign,omitempty"`
	ErrorMessageAlign string      `json:"errorMessageAlign,omitempty"`
	ShowErrorLine     bool        `json:"showErrorLine,omitempty"`
	ShowErrorMessage  bool        `json:"showErrorMessage,omitempty"`

	Value                  interface{} `json:"value"`
	Type                   string      `json:"type"`
	Password               bool        `json:"password"`
	Placeholder            string      `json:"placeholder"`
	PlaceholderStyle       string      `json:"placeholderStyle"`
	PlaceholderClass       string      `json:"placeholderClass"`
	Disabled               bool        `json:"disabled"`
	Maxlength              int         `json:"maxlength"`
	CursorSpacing          int         `json:"cursorSpacing"`
	Focus                  bool        `json:"focus"`
	ConfirmType            string      `json:"confirmType"`
	ConfirmHold            bool        `json:"confirmHold"`
	Cursor                 int         `json:"cursor"`
	SelectionStart         int         `json:"selectionStart"`
	SelectionEnd           int         `json:"selectionEnd"`
	AdjustPosition         bool        `json:"adjustPosition"`
	AutoBlur               bool        `json:"autoBlur"`
	IgnoreCompositionEvent bool        `json:"ignoreCompositionEvent"`
	AlwaysEmbed            bool        `json:"alwaysEmbed"`
	HoldKeyboard           bool        `json:"holdKeyboard"`
	SafePasswordCertPath   string      `json:"safePasswordCertPath"`
	SafePasswordLength     int         `json:"safePasswordLength"`
	SafePasswordTimeStamp  int         `json:"safePasswordTimeStamp"`
	SafePasswordNonce      string      `json:"safePasswordNonce"`
	SafePasswordSalt       string      `json:"safePasswordSalt"`
	SafePasswordCustomHash string      `json:"safePasswordCustomHash"`
	RandomNumber           bool        `json:"randomNumber"`
	Controlled             bool        `json:"controlled"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "inputField"
	p.Type = "text"
	p.SetKey("input", component.DEFAULT_CRYPT)
	p.PlaceholderClass = "input-placeholder"
	p.Maxlength = 140
	p.ConfirmType = "done"
	p.SelectionStart = -1
	p.SelectionEnd = -1
	p.AdjustPosition = true
	p.IgnoreCompositionEvent = true

	return p
}

// Set style.
func (p *Component) SetStyle(style interface{}) *Component {
	p.Style = style

	return p
}

// 表单域 v-model 字段， 在使用表单校验功能的情况下，该属性是必填的
func (p *Component) SetProp(prop string) *Component {
	p.Name = prop
	p.Prop = prop

	return p
}

// 表单域 v-model 字段， 在使用表单校验功能的情况下，该属性是必填的
func (p *Component) SetName(name string) *Component {
	p.Name = name
	p.Prop = name

	return p
}

// 定义校验规则
func (p *Component) SetRules(rules interface{}) *Component {
	p.Rules = rules

	return p
}

// 是否显示必填字段的标签旁边的红色星号
func (p *Component) SetRequired(required bool) *Component {
	p.Required = required

	return p
}

// 输入框左边的文字提示
func (p *Component) SetLabel(label string) *Component {
	p.Label = label

	return p
}

// 表单项 label 宽度，默认单位为px
func (p *Component) SetLabelWidth(labelWidth int) *Component {
	p.LabelWidth = labelWidth

	return p
}

// 表单项 label 对齐方式，可选值为 center right
func (p *Component) SetLabelAlign(labelAlign string) *Component {
	p.LabelAlign = labelAlign

	return p
}

// 右侧插槽对齐方式，可选值为 center right
func (p *Component) SetBodyAlign(bodyAlign string) *Component {
	p.BodyAlign = bodyAlign

	return p
}

// 错误提示文案对齐方式，可选值为 center right
func (p *Component) SetErrorMessageAlign(errorMessageAlign string) *Component {
	p.ErrorMessageAlign = errorMessageAlign

	return p
}

// 是否在校验不通过时标红输入框
func (p *Component) SetShowErrorLine(showErrorLine bool) *Component {
	p.ShowErrorLine = showErrorLine

	return p
}

// 是否在校验不通过时在输入框下方展示错误提示
func (p *Component) SetShowErrorMessage(showErrorMessage bool) *Component {
	p.ShowErrorMessage = showErrorMessage

	return p
}

// 输入框的初始内容
func (p *Component) SetValue(value interface{}) *Component {
	p.Value = value

	return p
}

// input 的类型
func (p *Component) SetType(textType string) *Component {
	p.Type = textType

	return p
}

// 是否是密码类型
func (p *Component) SetPassword(password bool) *Component {
	p.Password = password

	return p
}

// 输入框为空时占位符
func (p *Component) SetPlaceholder(placeholder string) *Component {
	p.Placeholder = placeholder

	return p
}

// 指定 placeholder 的样式
func (p *Component) SetPlaceholderStyle(placeholderStyle string) *Component {
	p.PlaceholderStyle = placeholderStyle

	return p
}

// 指定 placeholder 的样式类，注意页面或组件的style中写了scoped时，需要在类名前写/deep/
func (p *Component) SetPlaceholderClass(placeholderClass string) *Component {
	p.PlaceholderClass = placeholderClass

	return p
}

// 是否禁用
func (p *Component) SetDisabled(disabled bool) *Component {
	p.Disabled = disabled

	return p
}

// 最大输入长度，设置为 -1 的时候不限制最大长度
func (p *Component) SetMaxlength(maxlength int) *Component {
	p.Maxlength = maxlength

	return p
}

// 指定光标与键盘的距离，单位 px 。取 input 距离底部的距离和 cursor-spacing 指定的距离的最小值作为光标与键盘的距离
func (p *Component) SetCursorSpacing(cursorSpacing int) *Component {
	p.CursorSpacing = cursorSpacing

	return p
}

// 获取焦点
func (p *Component) SetFocus(focus bool) *Component {
	p.Focus = focus

	return p
}

// 设置键盘右下角按钮的文字，仅在 type="text" 时生效。
func (p *Component) SetConfirmType(confirmType string) *Component {
	p.ConfirmType = confirmType

	return p
}

// 点击键盘右下角按钮时是否保持键盘不收起
func (p *Component) SetConfirmHold(confirmHold bool) *Component {
	p.ConfirmHold = confirmHold

	return p
}

// 指定focus时的光标位置
func (p *Component) SetCursor(cursor int) *Component {
	p.Cursor = cursor

	return p
}

// 光标起始位置，自动聚集时有效，需与selection-end搭配使用
func (p *Component) SetSelectionStart(selectionStart int) *Component {
	p.SelectionStart = selectionStart

	return p
}

// 光标结束位置，自动聚集时有效，需与selection-start搭配使用
func (p *Component) SetSelectionEnd(selectionEnd int) *Component {
	p.SelectionEnd = selectionEnd

	return p
}

// 键盘弹起时，是否自动上推页面
func (p *Component) SetAdjustPosition(adjustPosition bool) *Component {
	p.AdjustPosition = adjustPosition

	return p
}

// 键盘收起时，是否自动失去焦点
func (p *Component) SetAutoBlur(autoBlur bool) *Component {
	p.AutoBlur = autoBlur

	return p
}

// 是否忽略组件内对文本合成系统事件的处理。为 false 时将触发 compositionstart、compositionend、compositionupdate 事件，且在文本合成期间会触发 input 事件
func (p *Component) SetIgnoreCompositionEvent(ignoreCompositionEvent bool) *Component {
	p.IgnoreCompositionEvent = ignoreCompositionEvent

	return p
}

// 强制 input 处于同层状态，默认 focus 时 input 会切到非同层状态 (仅在 iOS 下生效)
func (p *Component) SetAlwaysEmbed(alwaysEmbed bool) *Component {
	p.AlwaysEmbed = alwaysEmbed

	return p
}

// focus时，点击页面的时候不收起键盘
func (p *Component) SetHoldKeyboard(holdKeyboard bool) *Component {
	p.HoldKeyboard = holdKeyboard

	return p
}

// 安全键盘加密公钥的路径，只支持包内路径
func (p *Component) SetSafePasswordCertPath(safePasswordCertPath string) *Component {
	p.SafePasswordCertPath = safePasswordCertPath

	return p
}

// 安全键盘输入密码长度
func (p *Component) SetSafePasswordLength(safePasswordLength int) *Component {
	p.SafePasswordLength = safePasswordLength

	return p
}

// 安全键盘加密时间戳
func (p *Component) SetSafePasswordTimeStamp(safePasswordTimeStamp int) *Component {
	p.SafePasswordTimeStamp = safePasswordTimeStamp

	return p
}

// 安全键盘加密盐值
func (p *Component) SetSafePasswordNonce(safePasswordNonce string) *Component {
	p.SafePasswordNonce = safePasswordNonce

	return p
}

// 安全键盘计算 hash 盐值，若指定custom-hash 则无效
func (p *Component) SetSafePasswordSalt(safePasswordSalt string) *Component {
	p.SafePasswordSalt = safePasswordSalt

	return p
}

// 安全键盘计算 hash 的算法表达式，如 md5(sha1('foo' + sha256(sm3(password + 'bar'))))
func (p *Component) SetSafePasswordCustomHash(safePasswordCustomHash string) *Component {
	p.SafePasswordCustomHash = safePasswordCustomHash

	return p
}

// 当 type 为 number, digit, idcard 数字键盘是否随机排列
func (p *Component) SetRandomNumber(randomNumber bool) *Component {
	p.RandomNumber = randomNumber

	return p
}

// 是否为受控组件。为 true 时，value 内容会完全受 setData 控制
func (p *Component) SetControlled(controlled bool) *Component {
	p.Controlled = controlled

	return p
}

// 组件json序列化
func (p *Component) JsonSerialize() *Component {
	p.Component = "inputField"

	return p
}
