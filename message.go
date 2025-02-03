package quark

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
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

// 返回正确信息
func Success(message string, data interface{}) *Message {
	return &Message{
		Code: StatusOk,
		Msg:  message,
		Data: data,
	}
}
