package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type EasyInput struct {
	Item
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
	AlwaysSystem           bool        `json:"alwaysSystem"`
	Clearable              bool        `json:"clearable"`
	AutoHeight             bool        `json:"autoHeight"`
	ClearSize              int         `json:"clearSize"`
	PrefixIcon             string      `json:"prefixIcon"`
	SuffixIcon             string      `json:"suffixIcon"`
	Trim                   interface{} `json:"trim"`
	InputBorder            bool        `json:"inputBorder"`
	Styles                 interface{} `json:"styles"`
	PasswordIcon           bool        `json:"passwordIcon"`
}

// 初始化
func (p *EasyInput) Init() *EasyInput {
	p.Component = "easyInputField"
	p.Type = "text"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.PlaceholderClass = "input-placeholder"
	p.Maxlength = 140
	p.ConfirmType = "done"
	p.SelectionStart = -1
	p.SelectionEnd = -1
	p.AdjustPosition = true
	p.IgnoreCompositionEvent = true
	p.InputBorder = true
	p.PasswordIcon = true
	p.Styles = map[string]string{
		"color":        "#333",
		"disableColor": "#eee",
		"borderColor":  "#e5e5e5",
	}

	return p
}

// 输入框的初始内容
func (p *EasyInput) SetValue(value interface{}) *EasyInput {
	p.Value = value

	return p
}

// input 的类型
func (p *EasyInput) SetType(textType string) *EasyInput {
	p.Type = textType

	return p
}

// 是否是密码类型
func (p *EasyInput) SetPassword(password bool) *EasyInput {
	p.Password = password

	return p
}

// 输入框为空时占位符
func (p *EasyInput) SetPlaceholder(placeholder string) *EasyInput {
	p.Placeholder = placeholder

	return p
}

// 指定 placeholder 的样式
func (p *EasyInput) SetPlaceholderStyle(placeholderStyle string) *EasyInput {
	p.PlaceholderStyle = placeholderStyle

	return p
}

// 指定 placeholder 的样式类，注意页面或组件的style中写了scoped时，需要在类名前写/deep/
func (p *EasyInput) SetPlaceholderClass(placeholderClass string) *EasyInput {
	p.PlaceholderClass = placeholderClass

	return p
}

// 是否禁用
func (p *EasyInput) SetDisabled(disabled bool) *EasyInput {
	p.Disabled = disabled

	return p
}

// 最大输入长度，设置为 -1 的时候不限制最大长度
func (p *EasyInput) SetMaxlength(maxlength int) *EasyInput {
	p.Maxlength = maxlength

	return p
}

// 指定光标与键盘的距离，单位 px 。取 input 距离底部的距离和 cursor-spacing 指定的距离的最小值作为光标与键盘的距离
func (p *EasyInput) SetCursorSpacing(cursorSpacing int) *EasyInput {
	p.CursorSpacing = cursorSpacing

	return p
}

// 获取焦点
func (p *EasyInput) SetFocus(focus bool) *EasyInput {
	p.Focus = focus

	return p
}

// 设置键盘右下角按钮的文字，仅在 type="text" 时生效。
func (p *EasyInput) SetConfirmType(confirmType string) *EasyInput {
	p.ConfirmType = confirmType

	return p
}

// 点击键盘右下角按钮时是否保持键盘不收起
func (p *EasyInput) SetConfirmHold(confirmHold bool) *EasyInput {
	p.ConfirmHold = confirmHold

	return p
}

// 指定focus时的光标位置
func (p *EasyInput) SetCursor(cursor int) *EasyInput {
	p.Cursor = cursor

	return p
}

// 光标起始位置，自动聚集时有效，需与selection-end搭配使用
func (p *EasyInput) SetSelectionStart(selectionStart int) *EasyInput {
	p.SelectionStart = selectionStart

	return p
}

// 光标结束位置，自动聚集时有效，需与selection-start搭配使用
func (p *EasyInput) SetSelectionEnd(selectionEnd int) *EasyInput {
	p.SelectionEnd = selectionEnd

	return p
}

// 键盘弹起时，是否自动上推页面
func (p *EasyInput) SetAdjustPosition(adjustPosition bool) *EasyInput {
	p.AdjustPosition = adjustPosition

	return p
}

// 键盘收起时，是否自动失去焦点
func (p *EasyInput) SetAutoBlur(autoBlur bool) *EasyInput {
	p.AutoBlur = autoBlur

	return p
}

// 是否忽略组件内对文本合成系统事件的处理。为 false 时将触发 compositionstart、compositionend、compositionupdate 事件，且在文本合成期间会触发 input 事件
func (p *EasyInput) SetIgnoreCompositionEvent(ignoreCompositionEvent bool) *EasyInput {
	p.IgnoreCompositionEvent = ignoreCompositionEvent

	return p
}

// 强制 input 处于同层状态，默认 focus 时 input 会切到非同层状态 (仅在 iOS 下生效)
func (p *EasyInput) SetAlwaysEmbed(alwaysEmbed bool) *EasyInput {
	p.AlwaysEmbed = alwaysEmbed

	return p
}

// focus时，点击页面的时候不收起键盘
func (p *EasyInput) SetHoldKeyboard(holdKeyboard bool) *EasyInput {
	p.HoldKeyboard = holdKeyboard

	return p
}

// 	安全键盘加密公钥的路径，只支持包内路径
func (p *EasyInput) SetSafePasswordCertPath(safePasswordCertPath string) *EasyInput {
	p.SafePasswordCertPath = safePasswordCertPath

	return p
}

// 安全键盘输入密码长度
func (p *EasyInput) SetSafePasswordLength(safePasswordLength int) *EasyInput {
	p.SafePasswordLength = safePasswordLength

	return p
}

// 安全键盘加密时间戳
func (p *EasyInput) SetSafePasswordTimeStamp(safePasswordTimeStamp int) *EasyInput {
	p.SafePasswordTimeStamp = safePasswordTimeStamp

	return p
}

// 安全键盘加密盐值
func (p *EasyInput) SetSafePasswordNonce(safePasswordNonce string) *EasyInput {
	p.SafePasswordNonce = safePasswordNonce

	return p
}

// 安全键盘计算 hash 盐值，若指定custom-hash 则无效
func (p *EasyInput) SetSafePasswordSalt(safePasswordSalt string) *EasyInput {
	p.SafePasswordSalt = safePasswordSalt

	return p
}

// 安全键盘计算 hash 的算法表达式，如 md5(sha1('foo' + sha256(sm3(password + 'bar'))))
func (p *EasyInput) SetSafePasswordCustomHash(safePasswordCustomHash string) *EasyInput {
	p.SafePasswordCustomHash = safePasswordCustomHash

	return p
}

// 当 type 为 number, digit, idcard 数字键盘是否随机排列
func (p *EasyInput) SetRandomNumber(randomNumber bool) *EasyInput {
	p.RandomNumber = randomNumber

	return p
}

// 是否为受控组件。为 true 时，value 内容会完全受 setData 控制
func (p *EasyInput) SetControlled(controlled bool) *EasyInput {
	p.Controlled = controlled

	return p
}

// 是否强制使用系统键盘和 Web-view 创建的 input 元素。为 true 时，confirm-type、confirm-hold 可能失效
func (p *EasyInput) SetAlwaysSystem(alwaysSystem bool) *EasyInput {
	p.AlwaysSystem = alwaysSystem

	return p
}

// 是否显示右侧清空内容的图标控件(输入框有内容且不禁用时显示)，点击可清空输入框内容
func (p *EasyInput) SetClearable(clearable bool) *EasyInput {
	p.Clearable = clearable

	return p
}

// 是否自动增高输入区域，type为textarea时有效
func (p *EasyInput) SetAutoHeight(autoHeight bool) *EasyInput {
	p.AutoHeight = autoHeight

	return p
}

// 清除图标的大小，单位px
func (p *EasyInput) SetClearSize(clearSize int) *EasyInput {
	p.ClearSize = clearSize

	return p
}

// 输入框头部图标
func (p *EasyInput) SetPrefixIcon(prefixIcon string) *EasyInput {
	p.PrefixIcon = prefixIcon

	return p
}

// 输入框尾部图标
func (p *EasyInput) SetSuffixIcon(suffixIcon string) *EasyInput {
	p.SuffixIcon = suffixIcon

	return p
}

// 是否自动去除空格，传入类型为 Boolean 时，自动去除前后空格
func (p *EasyInput) SetTrim(trim interface{}) *EasyInput {
	p.Trim = trim

	return p
}

// 是否显示input输入框的边框
func (p *EasyInput) SetInputBorder(inputBorder bool) *EasyInput {
	p.InputBorder = inputBorder

	return p
}

// 样式自定义
func (p *EasyInput) SetStyles(styles interface{}) *EasyInput {
	p.Styles = styles

	return p
}

// type=password 时，是否显示小眼睛图标
func (p *EasyInput) SetPasswordIcon(passwordIcon bool) *EasyInput {
	p.PasswordIcon = passwordIcon

	return p
}

// 组件json序列化
func (p *EasyInput) JsonSerialize() *EasyInput {
	p.Component = "easyInputField"

	return p
}
