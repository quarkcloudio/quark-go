package adminresource

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/xuri/excelize/v2"
)

type ImportTemplateRequest struct{}

// 导入数据模板
func (p *ImportTemplateRequest) Handle(ctx *builder.Context) interface{} {
	fields := ctx.Template.(interface {
		ImportFields(ctx *builder.Context) interface{}
	}).ImportFields(ctx)

	exportTitles := []string{}
	for _, v := range fields.([]interface{}) {
		label := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Label").String()

		exportTitles = append(exportTitles, label+p.getFieldRemark(v))
	}

	f := excelize.NewFile()
	// 创建一个工作表
	index, _ := f.NewSheet("Sheet1")

	//定义一个字符 变量a 是一个byte类型的 表示单个字符
	var a = 'a'

	//生成26个字符
	for i := 1; i <= len(exportTitles); i++ {
		// 设置单元格的值
		f.SetCellValue("Sheet1", string(a)+"1", exportTitles[i-1])
		a++
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	buf, _ := f.WriteToBuffer()

	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=data_"+time.Now().Format("20060102150405")+".xlsx")
	ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
	ctx.Writer.Write(buf.Bytes())

	return nil
}

// 导入字段提示信息
func (p *ImportTemplateRequest) getFieldRemark(field interface{}) string {
	remark := ""

	component := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Component").String()

	switch component {
	case "inputNumberField":

	}

	switch component {
	case "inputNumberField":
		remark = "数字格式"

	case "textField":
		remark = ""

	case "selectField":

		options := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Options").Interface()

		mode := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Mode").String()

		fieldOptionLabel := p.getFieldOptionLabels(options)

		if mode != "" {
			remark = "可多选：" + fieldOptionLabel + "；多值请用“,”分割"
		} else {
			remark = "可选：" + fieldOptionLabel
		}

	case "cascaderField":
		remark = "级联格式，例如：省，市，县"

	case "checkboxField":
		options := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Options").Interface()

		remark = "可多选项：" + p.getFieldOptionLabels(options) + "；多值请用“,”分割"

	case "radioField":
		options := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Options").Interface()

		remark = "可选项：" + p.getFieldOptionLabels(options)

	case "switchField":
		options := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Options").Interface()

		remark = "可选项：" + p.getSwitchLabels(options)

	case "dateField":
		remark = "日期格式，例如：1987-02-15"

	case "datetimeField":
		remark = "日期时间格式，例如：1987-02-15 20:00:00"
	}

	rules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Rules").Interface()

	creationRules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("CreationRules").Interface()

	items := []interface{}{}

	for _, v := range rules.([]string) {
		items = append(items, v)
	}

	for _, v := range creationRules.([]string) {
		items = append(items, v)
	}

	ruleMessage := p.getFieldRuleMessage(items)

	if ruleMessage != "" {
		remark = remark + " 条件：" + ruleMessage
	}

	if remark != "" {
		remark = "（" + remark + "）"
	}

	return remark
}

// 导入字段的规则
func (p *ImportTemplateRequest) getFieldRuleMessage(rules []interface{}) string {
	var message []string
	rule := ""

	for _, v := range rules {
		var arr []string

		if strings.Contains(v.(string), ":") {
			arr = strings.Split(v.(string), ":")
			rule = arr[0]
		} else {
			rule = v.(string)
		}

		switch rule {

		case "required":
			// 必填
			message = append(message, "必填")

		case "min":
			// 最小字符串数
			message = append(message, "大于"+arr[1]+"个字符")

		case "max":
			// 最大字符串数
			message = append(message, "小于"+arr[1]+"个字符")

		case "email":
			// 必须为邮箱
			message = append(message, "必须为邮箱格式")

		case "numeric":
			// 必须为数字
			message = append(message, "必须为数字格式")

		case "url":
			// 必须为url
			message = append(message, "必须为链接格式")

		case "integer":
			// 必须为整数
			message = append(message, "必须为整数格式")

		case "date":
			// 必须为日期
			message = append(message, "必须为日期格式")

		case "boolean":
			// 必须为布尔值
			message = append(message, "必须为布尔格式")

		case "unique":
			// 必须为布尔值
			message = append(message, "不可重复")
		}
	}

	if len(message) > 0 {
		return strings.Replace(strings.Trim(fmt.Sprint(message), "/"), " ", "，", -1)
	} else {
		return ""
	}
}

// 获取字段的可选值
func (p *ImportTemplateRequest) getFieldOptionLabels(options interface{}) string {
	result := []string{}

	for _, v := range options.([]map[string]interface{}) {
		result = append(result, v["label"].(string))
	}

	return strings.Replace(strings.Trim(fmt.Sprint(result), "[]"), " ", "，", -1)
}

// 获取开关组件值
func (p *ImportTemplateRequest) getSwitchLabels(options interface{}) string {
	return options.(map[string]interface{})["on"].(string) + "，" + options.(map[string]interface{})["off"].(string)
}
