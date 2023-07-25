package convert

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// interface转string
func AnyToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value := value.(type) {
	case float64:
		key = strconv.FormatFloat(value, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(value), 'f', -1, 64)
	case int:
		key = strconv.Itoa(value)
	case uint:
		key = strconv.Itoa(int(value))
	case int8:
		key = strconv.Itoa(int(value))
	case uint8:
		key = strconv.Itoa(int(value))
	case int16:
		key = strconv.Itoa(int(value))
	case uint16:
		key = strconv.Itoa(int(value))
	case int32:
		key = strconv.Itoa(int(value))
	case uint32:
		key = strconv.Itoa(int(value))
	case int64:
		key = strconv.FormatInt(value, 10)
	case uint64:
		key = strconv.FormatUint(value, 10)
	case string:
		key = value
	case time.Time:
		key = value.String()

		// 2022-11-23 11:29:07 +0800 CST  这类格式把尾巴去掉
		key = strings.Replace(key, " +0800 CST", "", 1)
		key = strings.Replace(key, " +0000 UTC", "", 1)
	case []byte:
		key = string(value)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
