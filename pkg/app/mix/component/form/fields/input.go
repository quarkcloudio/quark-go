package fields

import "github.com/quarkcms/quark-go/v2/pkg/app/mix/component/component"

type Input struct {
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
}

// 初始化
func (p *Input) Init() *Input {
	p.Component = "inputField"
	p.Type = "text"
	p.SetKey("formItem", component.DEFAULT_CRYPT)
	p.PlaceholderClass = "input-placeholder"
	p.Maxlength = 140
	p.ConfirmType = "done"
	p.SelectionStart = -1
	p.SelectionEnd = -1
	p.AdjustPosition = true
	p.IgnoreCompositionEvent = true

	return p
}

// 输入框的初始内容
func (p *Input) SetValue(value interface{}) *Input {
	p.Value = value

	return p
}

// input 的类型
func (p *Input) SetType(textType string) *Input {
	p.Type = textType

	return p
}

// 是否是密码类型
func (p *Input) SetPassword(password bool) *Input {
	p.Password = password

	return p
}

// 输入框为空时占位符
func (p *Input) SetPlaceholder(placeholder string) *Input {
	p.Placeholder = placeholder

	return p
}

// 指定 placeholder 的样式
func (p *Input) SetPlaceholderStyle(placeholderStyle string) *Input {
	p.PlaceholderStyle = placeholderStyle

	return p
}

// 指定 placeholder 的样式类，注意页面或组件的style中写了scoped时，需要在类名前写/deep/
func (p *Input) SetPlaceholderClass(placeholderClass string) *Input {
	p.PlaceholderClass = placeholderClass

	return p
}

// 是否禁用
func (p *Input) SetDisabled(disabled bool) *Input {
	p.Disabled = disabled

	return p
}

// 最大输入长度，设置为 -1 的时候不限制最大长度
func (p *Input) SetMaxlength(maxlength int) *Input {
	p.Maxlength = maxlength

	return p
}

// 指定光标与键盘的距离，单位 px 。取 input 距离底部的距离和 cursor-spacing 指定的距离的最小值作为光标与键盘的距离
func (p *Input) SetCursorSpacing(cursorSpacing int) *Input {
	p.CursorSpacing = cursorSpacing

	return p
}

// 获取焦点
func (p *Input) SetFocus(focus bool) *Input {
	p.Focus = focus

	return p
}

// 设置键盘右下角按钮的文字，仅在 type="text" 时生效。
func (p *Input) SetConfirmType(confirmType string) *Input {
	p.ConfirmType = confirmType

	return p
}

// 点击键盘右下角按钮时是否保持键盘不收起
func (p *Input) SetConfirmHold(confirmHold bool) *Input {
	p.ConfirmHold = confirmHold

	return p
}

// 指定focus时的光标位置
func (p *Input) SetCursor(cursor int) *Input {
	p.Cursor = cursor

	return p
}

// 光标起始位置，自动聚集时有效，需与selection-end搭配使用
func (p *Input) SetSelectionStart(selectionStart int) *Input {
	p.SelectionStart = selectionStart

	return p
}

// 光标结束位置，自动聚集时有效，需与selection-start搭配使用
func (p *Input) SetSelectionEnd(selectionEnd int) *Input {
	p.SelectionEnd = selectionEnd

	return p
}

// 键盘弹起时，是否自动上推页面
func (p *Input) SetAdjustPosition(adjustPosition bool) *Input {
	p.AdjustPosition = adjustPosition

	return p
}

// 键盘收起时，是否自动失去焦点
func (p *Input) SetAutoBlur(autoBlur bool) *Input {
	p.AutoBlur = autoBlur

	return p
}

// 是否忽略组件内对文本合成系统事件的处理。为 false 时将触发 compositionstart、compositionend、compositionupdate 事件，且在文本合成期间会触发 input 事件
func (p *Input) SetIgnoreCompositionEvent(ignoreCompositionEvent bool) *Input {
	p.IgnoreCompositionEvent = ignoreCompositionEvent

	return p
}

// 强制 input 处于同层状态，默认 focus 时 input 会切到非同层状态 (仅在 iOS 下生效)
func (p *Input) SetAlwaysEmbed(alwaysEmbed bool) *Input {
	p.AlwaysEmbed = alwaysEmbed

	return p
}

// focus时，点击页面的时候不收起键盘
func (p *Input) SetHoldKeyboard(holdKeyboard bool) *Input {
	p.HoldKeyboard = holdKeyboard

	return p
}

// 	安全键盘加密公钥的路径，只支持包内路径
func (p *Input) SetSafePasswordCertPath(safePasswordCertPath string) *Input {
	p.SafePasswordCertPath = safePasswordCertPath

	return p
}

// 安全键盘输入密码长度
func (p *Input) SetSafePasswordLength(safePasswordLength int) *Input {
	p.SafePasswordLength = safePasswordLength

	return p
}

// 安全键盘加密时间戳
func (p *Input) SetSafePasswordTimeStamp(safePasswordTimeStamp int) *Input {
	p.SafePasswordTimeStamp = safePasswordTimeStamp

	return p
}

// 安全键盘加密盐值
func (p *Input) SetSafePasswordNonce(safePasswordNonce string) *Input {
	p.SafePasswordNonce = safePasswordNonce

	return p
}

// 安全键盘计算 hash 盐值，若指定custom-hash 则无效
func (p *Input) SetSafePasswordSalt(safePasswordSalt string) *Input {
	p.SafePasswordSalt = safePasswordSalt

	return p
}

// 安全键盘计算 hash 的算法表达式，如 md5(sha1('foo' + sha256(sm3(password + 'bar'))))
func (p *Input) SetSafePasswordCustomHash(safePasswordCustomHash string) *Input {
	p.SafePasswordCustomHash = safePasswordCustomHash

	return p
}

// 当 type 为 number, digit, idcard 数字键盘是否随机排列
func (p *Input) SetRandomNumber(randomNumber bool) *Input {
	p.RandomNumber = randomNumber

	return p
}

// 是否为受控组件。为 true 时，value 内容会完全受 setData 控制
func (p *Input) SetControlled(controlled bool) *Input {
	p.Controlled = controlled

	return p
}

// 组件json序列化
func (p *Input) JsonSerialize() *Input {
	p.Component = "inputField"

	return p
}
