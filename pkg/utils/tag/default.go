package tag

import (
	"reflect"
	"strconv"
	"strings"
)

// 设置默认值
func SetDefaults(value reflect.Value) {

	// 检查传递的值是否是结构体或结构体指针类型
	if value.Kind() != reflect.Struct && (value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct) {
		// 如果不是结构体或结构体指针类型，直接返回
		return
	}
	
	// 解引用指针类型
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		value = value.Elem()
	}
	
	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		// 嵌套结构体递归处理
		if field.Kind() == reflect.Struct {
			SetDefaults(field)
		}
		// 检查字段标签中存在default标签并且当前值是该字段类型的零值
		defaultValue := typ.Field(i).Tag.Get("default")
		if defaultValue != "" && reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			setValue(field, defaultValue)
		}
	}
}

func setSliceFieldValue(field reflect.Value, defaultValue string) {

	values := strings.Split(defaultValue, ",")
	slice := reflect.MakeSlice(field.Type(), len(values), len(values))

	for key, value := range values {
		setValue(slice.Index(key), value)
	}
	field.Set(slice)
}

func setArrayFieldValue(field reflect.Value, defaultValue string) {

	for key, value := range strings.Split(defaultValue, ",") {
		if key >= field.Len() {
			continue
		}
		setValue(field.Index(key), value)
	}
}

func setMapFieldValue(field reflect.Value, defaultValue string) {

	mapType := field.Type()
	mapInstance := reflect.MakeMap(mapType)

	for _, pair := range strings.Split(defaultValue, ",") {
		kv := strings.Split(pair, ":")
		if len(kv) != 2 {
			continue
		}
		key := reflect.New(mapType.Key()).Elem()
		value := reflect.New(mapType.Elem()).Elem()

		setValue(key, kv[0])
		setValue(value, kv[1])

		mapInstance.SetMapIndex(key, value)
	}

	field.Set(mapInstance)
}

// 设置字段值
func setValue(field reflect.Value, defaultValue string) {

	switch field.Kind() {
	case reflect.String:
		field.SetString(defaultValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intValue, err := strconv.ParseInt(defaultValue, 10, 64); err == nil {
			field.SetInt(intValue)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if uintValue, err := strconv.ParseUint(defaultValue, 10, 64); err == nil {
			field.SetUint(uintValue)
		}
	case reflect.Float32, reflect.Float64:
		if floatValue, err := strconv.ParseFloat(defaultValue, 64); err == nil {
			field.SetFloat(floatValue)
		}
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(defaultValue); err == nil {
			field.SetBool(boolValue)
		}
	case reflect.Array:
		setArrayFieldValue(field, defaultValue)
	case reflect.Slice:
		setSliceFieldValue(field, defaultValue)
	case reflect.Map:
		setMapFieldValue(field, defaultValue)
	default:
		field.Set(reflect.ValueOf(defaultValue))
	}
}
