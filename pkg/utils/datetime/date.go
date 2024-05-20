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

	var now time.Time

	// 自定义格式解析
	if now, err = time.ParseInLocation("2006-01-02", string(data), time.Local); err == nil {
		*t = Date{now}
		return err
	}

	// 带引号的自定义格式解析
	if now, err = time.ParseInLocation("\"2006-01-02\"", string(data), time.Local); err == nil {
		*t = Date{now}
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

// 返回字符串
func (t *Date) String() string {
	return t.Format("2006-01-02")
}
