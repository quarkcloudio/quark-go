package quark

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type ComponentMessage struct {
	ClassName string      `json:"className"`
	Type      string      `json:"type"`
	Content   interface{} `json:"content"`
	Duration  int         `json:"duration"`
	Icon      string      `json:"icon"`
	Style     interface{} `json:"style"`
	Data      interface{} `json:"data"`
	Url       string      `json:"url"`
}

type CodeMap struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	StatusOk         = 200   // 执行成功
	StatusError      = 10001 // 自定义错误信息
	StatusParamError = 10002 // 参数错误
)

var CodeMaps = []*CodeMap{
	{
		Code: StatusOk,
		Msg:  "ok",
	},
	{
		Code: StatusError,
		Msg:  "Internal Server Error",
	},
	{
		Code: StatusParamError,
		Msg:  "Param Error",
	},
}

// 根据code获取错误信息
func GetMsgByCode(code int) string {
	for _, v := range CodeMaps {
		if v.Code == code {
			return v.Msg
		}
	}
	return ""
}

// 返回正确信息
func Success(message string, data interface{}) *Message {
	return &Message{
		Code: StatusOk,
		Msg:  message,
		Data: data,
	}
}

// 返回错误信息，Error("内部服务调用异常") | Error("错误", map[string]interface{}{"title":"标题"}) | Error(10001, "错误", map[string]interface{}{"title":"标题"})
func Error(params ...interface{}) *Message {
	var (
		code = StatusError
		msg  = ""
		data interface{}
	)
	if len(params) == 1 {
		msg = params[0].(string)
	}
	if len(params) == 2 {
		msg = params[0].(string)
		data = params[1]
	}
	if len(params) == 3 {
		code = params[0].(int)
		msg = params[1].(string)
		if msg == "" {
			msg = GetMsgByCode(code)
		}
		data = params[2]
	}

	return &Message{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// 返回组件引擎成功信息，ComponentSuccess("成功") | ComponentSuccess("成功", "/home/index", map[string]interface{}{"title":"标题"})
func ComponentSuccess(message ...interface{}) *ComponentMessage {
	var (
		content = ""
		url     = ""
		data    interface{}
	)
	if len(message) == 1 {
		content = message[0].(string)
	}
	if len(message) == 2 {
		content = message[0].(string)
		url = message[1].(string)
	}
	if len(message) >= 3 {
		content = message[0].(string)
		url = message[1].(string)
		data = message[2]
	}

	return &ComponentMessage{
		Type:    "message",
		Content: content,
		Url:     url,
		Data:    data,
	}
}

// 返回组件引擎失败信息，Error("错误") | Error("操作失败", "/home/index")
func ComponentError(message ...interface{}) *ComponentMessage {
	var (
		content = ""
		url     = ""
	)
	if len(message) == 1 {
		content = message[0].(string)
	}
	if len(message) == 2 {
		content = message[0].(string)
		url = message[1].(string)
	}

	return &ComponentMessage{
		Type:    "error",
		Content: content,
		Url:     url,
	}
}
