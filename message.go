package quark

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	Url  string      `json:"url,omitempty"`
}

type CodeMap struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	StatusOk           = 200   // 执行成功
	StatusUnauthorized = 401   // 未授权
	StatusForbidden    = 403   // 无权限
	StatusError        = 10001 // 自定义错误信息
	StatusParamError   = 10002 // 参数错误
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
	{
		Code: StatusUnauthorized,
		Msg:  "Unauthorized",
	},
	{
		Code: StatusForbidden,
		Msg:  "Forbidden",
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

// 输出模版引擎URL跳转，RedirectTo("/home/index") | RedirectTo("成功", "/home/index")  | RedirectTo("失败", "/home/index", 10001)
func RedirectTo(params ...interface{}) *Message {
	var (
		msg  = ""
		url  = ""
		code = 200
	)
	if len(params) == 1 {
		url = params[0].(string)
	}
	if len(params) == 2 {
		msg = params[0].(string)
		url = params[1].(string)
	}
	if len(params) >= 3 {
		msg = params[0].(string)
		url = params[1].(string)
		code = params[2].(int)
	}

	return &Message{
		Code: code,
		Msg:  msg,
		Url:  url,
	}
}
