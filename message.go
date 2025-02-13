package quark

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type ComponentMessage struct {
	Type     string      `json:"type"`
	Content  interface{} `json:"content"`
	Duration int         `json:"duration"`
	Icon     string      `json:"icon"`
	Data     interface{} `json:"data"`
	Url      string      `json:"url"`
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

// 返回错误信息，Error("内部服务调用异常") | Error("错误", map[string]interface{}{"title":"标题"})
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

	return &Message{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// 返回错误信息，ErrorByCode(10001) | ErrorByCode(10001, map[string]interface{}{"title":"标题"})
func ErrorByCode(params ...interface{}) *Message {
	var (
		code = StatusError
		msg  = ""
		data interface{}
	)
	if len(params) == 1 {
		code = params[0].(int)
	}
	if len(params) == 2 {
		code = params[0].(int)
		data = params[1]
	}
	msg = GetMsgByCode(code)

	return &Message{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// 返回组件引擎成功信息，ComponentSuccess("成功") | ComponentSuccess("成功", map[string]interface{}{"title":"标题"})
func ComponentSuccess(params ...interface{}) *ComponentMessage {
	var (
		content = ""
		data    interface{}
	)
	if len(params) == 1 {
		content = params[0].(string)
	}
	if len(params) == 2 {
		content = params[0].(string)
		data = params[1]
	}

	return &ComponentMessage{
		Type:    "success",
		Content: content,
		Data:    data,
	}
}

// 返回组件引擎失败信息，ComponentError("错误") | ComponentError("成功", map[string]interface{}{"title":"标题"})
func ComponentError(params ...interface{}) *ComponentMessage {
	var (
		content = ""
		data    interface{}
	)
	if len(params) == 1 {
		content = params[0].(string)
	}
	if len(params) == 2 {
		content = params[0].(string)
		data = params[1]
	}

	return &ComponentMessage{
		Type:    "error",
		Content: content,
		Data:    data,
	}
}

// 返回错误信息，ComponentErrorByCode(10001) | ComponentErrorByCode(10001, map[string]interface{}{"title":"标题"})
func ComponentErrorByCode(params ...interface{}) *ComponentMessage {
	var (
		code    = StatusError
		content = ""
		data    interface{}
	)
	if len(params) == 1 {
		code = params[0].(int)
	}
	if len(params) == 2 {
		code = params[0].(int)
		data = params[1]
	}
	content = GetMsgByCode(code)

	return &ComponentMessage{
		Type:    "error",
		Content: content,
		Data:    data,
	}
}

// 输出模版引擎URL跳转，ComponentRedirectTo("/home/index") | ComponentRedirectTo("成功", "/home/index")  | ComponentRedirectTo("成功", "/home/index", "error")
func ComponentRedirectTo(params ...interface{}) *ComponentMessage {
	var (
		content = ""
		url     = ""
		msgType = "success"
	)
	if len(params) == 1 {
		url = params[0].(string)
	}
	if len(params) == 2 {
		content = params[0].(string)
		url = params[1].(string)
	}
	if len(params) >= 3 {
		content = params[0].(string)
		url = params[1].(string)
		msgType = params[2].(string)
	}

	return &ComponentMessage{
		Type:    msgType,
		Content: content,
		Url:     url,
	}
}
