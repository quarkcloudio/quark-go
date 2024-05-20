package datetime

import (
	"database/sql/driver"
	"errors"
	"time"
)

// 时间
type Time struct {
	time.Time
}

// 当前时间
func TimeNow() Time {
	return Time{
		Time: time.Now(),
	}
}

// 编码为自定义的Json格式
func (t Time) MarshalJSON() ([]byte, error) {

	// 时间为零返回null
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte("\"" + t.Format("2006-01-02 15:04:05") + "\""), nil
}

// 将Json格式解码
func (t *Time) UnmarshalJSON(data []byte) error {

	var err error

	if len(data) == 2 || string(data) == "null" {
		return err
	}

	var now time.Time

	// 自定义格式解析
	if now, err = time.ParseInLocation("2006-01-02 15:04:05", string(data), time.Local); err == nil {
		*t = Time{now}
		return err
	}

	// 带引号的自定义格式解析
	if now, err = time.ParseInLocation("\"2006-01-02 15:04:05\"", string(data), time.Local); err == nil {
		*t = Time{now}
		return err
	}

	// 默认格式解析
	if now, err = time.ParseInLocation(time.RFC3339, string(data), time.Local); err == nil {
		*t = Time{now}
		return err
	}

	if now, err = time.ParseInLocation("\""+time.RFC3339+"\"", string(data), time.Local); err == nil {
		*t = Time{now}
		return err
	}

	return err
}

// 转换为数据库值
func (t Time) Value() (driver.Value, error) {

	if t.IsZero() {
		return nil, nil
	}

	return t.Time, nil
}

// 数据库值转换为Time
func (t *Time) Scan(i interface{}) error {

	if value, ok := i.(time.Time); ok {
		*t = Time{Time: value}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}

// 返回字符串
func (t *Time) String() string {
	return t.Format("2006-01-02 15:04:05")
}
