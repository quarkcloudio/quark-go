package login

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/component"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/form/fields/when"
)

type ActivityConfig struct {
	Title    string                 `json:"title,omitempty"`
	SubTitle string                 `json:"subTitle,omitempty"`
	Action   interface{}            `json:"action,omitempty"`
	Style    map[string]interface{} `json:"style,omitempty"`
}

type Component struct {
	component.Element
	Component          string                 `json:"component"`
	Api                string                 `json:"api,omitempty"`
	Redirect           string                 `json:"redirect,omitempty"`
	Logo               interface{}            `json:"logo,omitempty"`
	Title              string                 `json:"title,omitempty"`
	SubTitle           string                 `json:"subTitle,omitempty"`
	BackgroundImageUrl string                 `json:"backgroundImageUrl,omitempty"`
	CaptchaIdUrl       string                 `json:"captchaIdUrl,omitempty"`
	CaptchaUrl         string                 `json:"captchaUrl,omitempty"`
	LoginType          []string               `json:"loginType"`
	ActivityConfig     *ActivityConfig        `json:"activityConfig,omitempty"`
	Values             map[string]interface{} `json:"values,omitempty"`
	InitialValues      map[string]interface{} `json:"initialValues,omitempty"`
	Body               interface{}            `json:"body,omitempty"`
	Actions            []interface{}          `json:"actions,omitempty"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "login"
	p.LoginType = []string{"account"}
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style

	return p
}

// 登录接口
func (p *Component) SetApi(api string) *Component {
	p.Api = api
	return p
}

// 登录后跳转地址
func (p *Component) SetRedirect(redirect string) *Component {
	p.Redirect = redirect
	return p
}

// Logo
func (p *Component) SetLogo(logo interface{}) *Component {
	p.Logo = logo
	return p
}

// 标题
func (p *Component) SetTitle(title string) *Component {
	p.Title = title
	return p
}

// 子标题
func (p *Component) SetSubTitle(subTitle string) *Component {
	p.SubTitle = subTitle
	return p
}

// 整个区域的背景图片配置，手机端不会展示
func (p *Component) SeBackgroundImageUrl(backgroundImageUrl string) *Component {
	p.BackgroundImageUrl = backgroundImageUrl
	return p
}

// 获取验证码ID链接
func (p *Component) SetCaptchaIdUrl(captchaIdUrl string) *Component {
	p.CaptchaIdUrl = captchaIdUrl
	return p
}

// 验证码链接
func (p *Component) SetCaptchaUrl(captchaUrl string) *Component {
	p.CaptchaUrl = captchaUrl
	return p
}

// 登录类型，{"account","phone"}
func (p *Component) SetLoginType(loginType []string) *Component {
	p.LoginType = loginType
	return p
}

// 活动的配置，包含 title，subTitle，action，分别代表标题，次标题和行动按钮，也可配置 style 来控制区域的样式
func (p *Component) SetActivityConfig(activityConfig *ActivityConfig) *Component {
	p.ActivityConfig = activityConfig
	return p
}

// 解析initialValue
func (p *Component) parseInitialValue(item interface{}, initialValues map[string]interface{}) interface{} {
	var value any

	// 数组直接返回
	if _, ok := item.([]interface{}); ok {
		return nil
	}

	reflectElem := reflect.
		ValueOf(item).
		Elem()

	name := reflectElem.
		FieldByName("Name").
		String()
	if name == "" {
		return nil
	}

	issetDefaultValue := reflectElem.
		FieldByName("DefaultValue").
		IsValid()

	if issetDefaultValue {
		defaultValue := reflectElem.
			FieldByName("DefaultValue").
			Interface()

		if defaultValue != nil {
			value = defaultValue
		}
	}

	issetValue := reflectElem.
		FieldByName("Value").
		IsValid()

	if issetValue {
		getValue := reflectElem.
			FieldByName("Value").
			Interface()

		if getValue != nil {
			value = getValue
		}
	}

	if initialValues[name] != nil {
		value = initialValues[name]
	}

	return value
}

// 查找字段
func (p *Component) findFields(fields interface{}, when bool) interface{} {
	var items []interface{}

	if getFields, ok := fields.([]interface{}); ok {
		for _, v := range getFields {
			items = append(items, p.fieldParser(v, when)...)
		}
	} else {
		items = append(items, p.fieldParser(fields, when)...)
	}

	return items
}

// 解析字段
func (p *Component) fieldParser(v interface{}, when bool) []interface{} {
	var items []interface{}

	// 数组直接返回
	if _, ok := v.([]interface{}); ok {
		return items
	}

	vKind := reflect.
		ValueOf(v).
		Kind()
	if !(vKind.String() == "interface" || vKind.String() == "ptr") {
		return items
	}

	hasBody := reflect.
		ValueOf(v).
		Elem().
		FieldByName("Body").
		IsValid()

	// 存在body的情况下
	if hasBody {
		body := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Body").
			Interface()

		getItems := p.findFields(body, true)
		if getItems, ok := getItems.([]interface{}); ok {
			if len(getItems) > 0 {
				items = append(items, getItems...)
			}
		}

		return items
	}

	hasTabPanes := reflect.
		ValueOf(v).
		Elem().
		FieldByName("TabPanes").
		IsValid()

	// 存在TabPanes情况下
	if hasTabPanes {
		body := reflect.
			ValueOf(v).
			Elem().
			FieldByName("TabPanes").
			Interface()

		getItems := p.findFields(body, true)
		if getItems, ok := getItems.([]interface{}); ok {
			if len(getItems) > 0 {
				items = append(items, getItems...)
			}
		}

		return items
	}

	// 默认情况
	component := reflect.
		ValueOf(v).
		Elem().
		FieldByName("Component").
		String()
	if strings.Contains(component, "Field") {
		items = append(items, v)
		if when {
			whenFields := p.getWhenFields(v)
			if len(whenFields) > 0 {
				items = append(items, whenFields...)
			}
		}
	}

	return items
}

// 获取When组件中的字段
func (p *Component) getWhenFields(item interface{}) []interface{} {
	var items []interface{}
	whenIsValid := reflect.
		ValueOf(item).
		Elem().
		FieldByName("When").
		IsValid()
	if !whenIsValid {
		return items
	}

	getWhen := item.(interface {
		GetWhen() *when.Component
	}).GetWhen()

	if getWhen == nil {
		return items
	}
	whenItems := getWhen.Items
	if whenItems == nil {
		return items
	}

	for _, v := range whenItems {
		if v.Body != nil {
			if body, ok := v.Body.([]interface{}); ok {
				if len(body) > 0 {
					items = append(items, body...)
				}
			} else {
				items = append(items, body)
			}
		}
	}

	return items
}

// 表单默认值，只有初始化以及重置时生效
func (p *Component) SetInitialValues(initialValues map[string]interface{}) *Component {
	data := initialValues

	fields := p.findFields(p.Body, true)
	if body, ok := fields.([]any); ok {
		for _, v := range body {
			value := p.parseInitialValue(v, initialValues)
			if value != nil {
				name := reflect.
					ValueOf(v).
					Elem().
					FieldByName("Name").
					String()

				data[name] = value
			}
		}
	}

	for k, v := range data {
		getV, ok := v.(string)
		if ok {
			if strings.Contains(getV, "[") {
				var m []interface{}
				err := json.Unmarshal([]byte(getV), &m)
				if err == nil {
					v = m
				} else {
					if strings.Contains(getV, "{") {
						var m map[string]interface{}
						err := json.Unmarshal([]byte(getV), &m)
						if err == nil {
							v = m
						}
					}
				}
			} else {
				if strings.Contains(getV, "{") {
					var m map[string]interface{}
					err := json.Unmarshal([]byte(getV), &m)
					if err == nil {
						v = m
					}
				}
			}
		}

		data[k] = v
	}

	p.InitialValues = data

	return p
}

// 表单项
func (p *Component) SetBody(items interface{}) *Component {
	p.Body = items

	return p
}

// 设置表单行为
func (p *Component) SetActions(actions []interface{}) *Component {
	p.Actions = actions

	return p
}
