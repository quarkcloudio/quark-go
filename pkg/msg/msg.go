package msg

const DEFAULT_MSG string = ""
const DEFAULT_URL string = ""
const DEFAULT_DATA string = ""

// 返回错误信息
func Error(msg string, url string) map[string]interface{} {

	return map[string]interface{}{
		"component": "message",
		"status":    "error",
		"msg":       msg,
		"url":       url,
	}
}

// 返回正确信息
func Success(msg string, url string, data interface{}) map[string]interface{} {

	return map[string]interface{}{
		"component": "message",
		"status":    "success",
		"msg":       msg,
		"url":       url,
		"data":      data,
	}
}
