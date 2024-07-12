package datetime

import (
	"database/sql/driver"
	"errors"
	"time"
)

// 日期
type Date struct {
	time.Time
}

// 当前日期
func DateNow() Date {
	return Date{
		Time: time.Now(),
	}
}

// 编码为自定义的Json格式
func (t Date) MarshalJSON() ([]byte, error) {

	// 时间为零返回null
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte("\"" + t.Format("2006-01-02") + "\""), nil
}

// 将Json格式解码
func (t *Date) UnmarshalJSON(data []byte) error {

	var err error

	if len(data) == 2 || string(data) == "null" {
		return err
	}

	var now Date

	// 自定义格式解析
	if now, err = ParseDate(string(data), "2006-01-02"); err == nil {
		*t = now
		return err
	}

	// 带引号的自定义格式解析
	if now, err = ParseDate(string(data), "\"2006-01-02\""); err == nil {
		*t = now
		return err
	}

	return err
}

// 转换为数据库值
func (t Date) Value() (driver.Value, error) {

	if t.IsZero() {
		return nil, nil
	}

	return t.Time, nil
}

// 数据库值转换为Time
func (t *Date) Scan(i interface{}) error {

	if value, ok := i.(time.Time); ok {
		*t = Date{Time: value}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}

// 将Date类型转换为DateTime类型
func (t Date) ToDateTime() Datetime {
	return Datetime{
		Time: t.Time,
	}
}

// 格式化为字符串
func (t Date) ToString() string {
	return t.Format("2006-01-02")
}

// 自定义格式字符串
func (t Date) FormatToString(format string) string {
	return t.Format(format)
}
