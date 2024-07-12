package datetime

import (
	"database/sql/driver"
	"errors"
	"time"
)

// 日期时间
type Datetime struct {
	time.Time
}

// 当前日期时间
func Now() Datetime {
	return Datetime{
		Time: time.Now(),
	}
}

// 编码为自定义的Json格式
func (t Datetime) MarshalJSON() ([]byte, error) {

	// 时间为零返回null
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte("\"" + t.Format("2006-01-02 15:04:05") + "\""), nil
}

// 将Json格式解码
func (t *Datetime) UnmarshalJSON(data []byte) error {

	var err error

	if len(data) == 2 || string(data) == "null" {
		return err
	}

	var now Datetime

	// 自定义格式解析
	if now, err = ParseDatetime(string(data), "2006-01-02 15:04:05"); err == nil {
		*t = now
		return err
	}

	// 带引号的自定义格式解析
	if now, err = ParseDatetime(string(data), "\"2006-01-02 15:04:05\""); err == nil {
		*t = now
		return err
	}

	// 默认格式解析
	if now, err = ParseDatetime(string(data), time.RFC3339); err == nil {
		*t = now
		return err
	}

	if now, err = ParseDatetime(string(data), "\""+time.RFC3339+"\""); err == nil {
		*t = now
		return err
	}

	return err
}

// 转换为数据库值
func (t Datetime) Value() (driver.Value, error) {

	if t.IsZero() {
		return nil, nil
	}

	return t.Time, nil
}

// 数据库值转换为Time
func (t *Datetime) Scan(i interface{}) error {

	if value, ok := i.(time.Time); ok {
		*t = Datetime{Time: value}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}

// 将Datetime类型转换为Date类型
func (t Datetime) ToDate() Date {
	return Date{
		Time: t.Time,
	}
}

// 将Datetime类型转换为Time类型
func (t Datetime) ToTime() Time {
	return Time{
		Time: t.Time,
	}
}

// 格式化为字符串
func (t Datetime) ToString() string {
	return t.Format("2006-01-02 15:04:05")
}

// 自定义格式字符串
func (t Datetime) FormatToString(format string) string {
	return t.Format(format)
}
