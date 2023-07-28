package builder

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
	StatusOk    = 200   // 执行成功
	StatusError = 10001 // 自定义错误类型
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
}

// 返回错误信息，Error("内部服务调用异常")
func Error(message string) *Message {
	return &Message{
		Code: StatusError,
		Msg:  message,
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
