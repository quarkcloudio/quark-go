package requests

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource/types"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/xuri/excelize/v2"
)

type ImportTemplateRequest struct{}

// 导入数据模板
func (p *ImportTemplateRequest) Handle(ctx *builder.Context) error {

	// 模版实例
	template := ctx.Template.(types.Resourcer)

	// 获取字段
	fields := template.ImportFields(ctx)

	exportTitles := []string{}
	for _, v := range fields.([]interface{}) {
		label := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Label").
			String()

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
		FieldByName("Component").
		String()

	switch component {
	case "inputNumberField":
		remark = "数字格式"
	case "textField":
		remark = ""
	case "selectField":
		mode := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Mode").
			String()

		optionLabels := field.(interface {
			GetOptionLabels() string
		}).GetOptionLabels()

		if mode != "" {
			remark = "可多选：" + optionLabels + "；多值请用“,”分割"
		} else {
			remark = "可选：" + optionLabels
		}
	case "cascaderField":
		remark = "级联格式，例如：省，市，县"
	case "checkboxField":
		optionLabels := field.(interface {
			GetOptionLabels() string
		}).GetOptionLabels()

		remark = "可多选项：" + optionLabels + "；多值请用“,”分割"
	case "radioField":
		optionLabels := field.(interface {
			GetOptionLabels() string
		}).GetOptionLabels()

		remark = "可选项：" + optionLabels
	case "switchField":
		optionLabels := field.(interface {
			GetOptionLabels() string
		}).GetOptionLabels()

		remark = "可选项：" + optionLabels
	case "dateField":
		remark = "日期格式，例如：1987-02-15"

	case "datetimeField":
		remark = "日期时间格式，例如：1987-02-15 20:00:00"
	}

	var rules []*rule.Rule
	if v, ok := field.(interface {
		GetRules() []*rule.Rule
	}); ok {
		rules = append(rules, v.GetRules()...)
	}

	if v, ok := field.(interface {
		GetCreationRules() []*rule.Rule
	}); ok {
		rules = append(rules, v.GetCreationRules()...)
	}

	ruleMessage := p.getFieldRuleMessage(rules)
	if ruleMessage != "" {
		remark = remark + " 条件：" + ruleMessage
	}

	if remark != "" {
		remark = "（" + remark + "）"
	}

	return remark
}

// 导入字段的规则
func (p *ImportTemplateRequest) getFieldRuleMessage(rules []*rule.Rule) string {
	var message []string

	for _, v := range rules {
		switch v.RuleType {
		case "required":
			// 必填
			message = append(message, "必填")
		case "min":
			// 最小字符串数
			message = append(message, "大于"+strconv.Itoa(v.Min)+"个字符")
		case "max":
			// 最大字符串数
			message = append(message, "小于"+strconv.Itoa(v.Max)+"个字符")
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
	}

	return ""
}
