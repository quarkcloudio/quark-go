package datetime

import (
	"strings"
	"time"
)

// 将字符串解析为日期
// ParseDate("2024-05-20") | ParseDate("2024/05/20", "2006/01/02")
func ParseDate(values ...string) (Date, error) {

	date, err := parseInLocation("2006-01-02", time.Local, values...)

	return Date{Time: date}, err
}

// 将字符串解析为日期
// ParseDateInLocation(time.Local, "2024-05-20") | ParseDateInLocation(time.Local, "2024/05/20", "2006/01/02")
func ParseDateInLocation(location *time.Location, values ...string) (Date, error) {

	date, err := parseInLocation("2006-01-02", location, values...)

	return Date{Time: date}, err
}

// 将字符串解析为时间
// ParseTime("14:00:00") | ParseTime("14-00-00", "15-04-05")
func ParseTime(values ...string) (Time, error) {

	date, err := parseInLocation("15:04:05", time.Local, values...)

	return Time{Time: date}, err
}

// 将字符串解析为时间
// ParseTimeInLocation(time.Local, "14:00:00") | ParseTimeInLocation(time.Local, "14-00-00", "15-04-05")
func ParseTimeInLocation(location *time.Location, values ...string) (Time, error) {

	date, err := parseInLocation("15:04:05", location, values...)

	return Time{Time: date}, err
}

// 将字符串解析为日期时间
// ParseDatetime("2024-05-20 11:22:33") | ParseDatetime("2024/05/20 11:22:33", "2006/01/02 15:04:05")
func ParseDatetime(values ...string) (Datetime, error) {

	time, err := parseInLocation("2006-01-02 15:04:05", time.Local, values...)

	return Datetime{Time: time}, err
}

// 将字符串解析为时间
// ParseDatetimeInLocation(time.Local, "2024-05-20 11:22:33") | ParseDatetimeInLocation(time.Local, "2024/05/20 11:22:33", "2006/01/02 15:04:05")
func ParseDatetimeInLocation(location *time.Location, values ...string) (Datetime, error) {

	time, err := parseInLocation("2006-01-02 15:04:05", location, values...)

	return Datetime{Time: time}, err
}

// 用于解析带有可选自定义格式的日期或时间字符串
func parseInLocation(defaultLayout string, location *time.Location, values ...string) (time.Time, error) {

	now := time.Now().Format(defaultLayout)
	layout := defaultLayout

	if len(values) > 0 {
		now = strings.TrimSpace(values[0])
	}

	if len(values) > 1 {
		layout = strings.TrimSpace(values[1])
	}

	return time.ParseInLocation(layout, now, location)
}
