package requests

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"
	"strings"

	models "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/component/admin/space"
	"github.com/quarkcms/quark-go/pkg/component/admin/tpl"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/msg"
	"github.com/quarkcms/quark-go/pkg/rand"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ImportRequest struct{}

// 执行行为
func (p *ImportRequest) Handle(ctx *builder.Context, indexRoute string) error {
	data := map[string]interface{}{}
	json.Unmarshal(ctx.Body(), &data)
	fileId := data["fileId"]
	getFileId := 0

	if value, ok := fileId.(map[string]interface{}); ok {
		if value, ok := value["0"]; ok {
			if value, ok := value.(map[string]interface{})["id"]; ok {
				getFileId = int(value.(float64))
			}
		}
	}

	if getFileId == 0 {
		return ctx.JSONError("参数错误！")
	}

	modelInstance := reflect.
		ValueOf(ctx.Template).
		Elem().
		FieldByName("Model").Interface()
	model := db.Client.Model(&modelInstance)

	importData, err := (&models.File{}).GetExcelData(getFileId)
	if err != nil {
		return ctx.JSONError(err.Error())
	}

	// 表格头部
	importHead := importData[0]

	// 去除表格头部
	importData = importData[1:]

	// 导入前回调
	lists := ctx.Template.(interface {
		BeforeImporting(ctx *builder.Context, list [][]interface{}) [][]interface{}
	}).BeforeImporting(ctx, importData)

	importResult := true
	importTotalNum := len(lists)
	importSuccessedNum := 0
	importFailedNum := 0
	importFailedData := [][]interface{}{}

	// 获取字段
	fields := ctx.Template.(interface {
		ImportFields(ctx *builder.Context) interface{}
	}).ImportFields(ctx)

	for _, item := range lists {

		formValues := p.transformFormValues(fields, item)

		validator := ctx.Template.(interface {
			ValidatorForImport(ctx *builder.Context, data map[string]interface{}) error
		}).ValidatorForImport(ctx, formValues)

		if validator != nil {

			importResult = false
			importFailedNum = importFailedNum + 1  // 记录错误条数
			item = append(item, validator.Error()) // 记录错误信息

			//记录错误数据
			importFailedData = append(importFailedData, item)
		}

		submitData, err := ctx.Template.(interface {
			BeforeSaving(ctx *builder.Context, data map[string]interface{}) (map[string]interface{}, error)
		}).BeforeSaving(ctx, formValues)
		if err != nil {
			importResult = false
			importFailedNum = importFailedNum + 1 // 记录错误条数
			item = append(item, err.Error())      // 记录错误信息

			//记录错误数据
			importFailedData = append(importFailedData, item)
		}

		data := p.getSubmitData(
			fields,
			submitData,
		)

		// 获取对象
		getModel := model.Create(data)

		if getModel.Error != nil {
			importResult = false
			importFailedNum = importFailedNum + 1       // 记录错误条数
			item = append(item, getModel.Error.Error()) // 记录错误信息

			//记录错误数据
			importFailedData = append(importFailedData, item)
		} else {

			lastData := map[string]interface{}{}
			getModel.Order("id desc").First(&lastData)

			ctx.Template.(interface {
				AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) interface{}
			}).AfterSaved(ctx, lastData["id"].(int), data, getModel)

			importSuccessedNum = importSuccessedNum + 1
		}
	}

	if importResult {
		return ctx.JSONOk("操作成功！", strings.Replace("/layout/index?api="+indexRoute, ":resource", ctx.Param("resource"), -1))
	} else {
		importHead = append(importHead, "错误信息")

		f := excelize.NewFile()
		// Create a new sheet.
		index, _ := f.NewSheet("Sheet1")

		//定义一个字符 变量a 是一个byte类型的 表示单个字符
		var a = 'a'

		//生成26个字符
		for i := 1; i <= len(importHead); i++ {
			// 设置单元格的值
			f.SetCellValue("Sheet1", string(a)+"1", importHead[i-1])
			a++
		}

		for k, v := range importFailedData {
			//定义一个字符 变量a 是一个byte类型的 表示单个字符
			var a = 'a'

			//生成26个字符
			for i := 1; i <= len(v); i++ {
				// 设置单元格的值
				f.SetCellValue("Sheet1", string(a)+strconv.Itoa(k+2), v[i-1])
				a++
			}
		}

		f.SetActiveSheet(index)

		filePath := "./website/storage/failImports/"
		fileName := rand.MakeAlphanumeric(40) + ".xlsx"

		// 不存在路径，则创建
		if isExist(filePath) == false {
			err := os.MkdirAll(filePath, 0666)
			if err != nil {
				return ctx.JSON(200, msg.Error(err.Error(), ""))
			}
		}

		if err := f.SaveAs(filePath + fileName); err != nil {
			return ctx.JSON(200, msg.Error(err.Error(), ""))
		}

		importTotalNumTpl := (&tpl.Component{}).
			Init().
			SetBody("导入总量: " + strconv.Itoa(importTotalNum))

		importSuccessedNumTpl := (&tpl.Component{}).
			Init().
			SetBody("成功数量: " + strconv.Itoa(importSuccessedNum))

		importFailedNumTpl := (&tpl.Component{}).
			Init().
			SetBody("失败数量: <span style='color:#ff4d4f'>" + strconv.Itoa(importFailedNum) + "</span> <a href='//" + ctx.Host() + "/storage/failImports/" + fileName + "' target='_blank'>下载失败数据</a>")

		component := (&space.Component{}).
			Init().
			SetBody([]interface{}{
				importTotalNumTpl,
				importSuccessedNumTpl,
				importFailedNumTpl,
			}).
			SetDirection("vertical").
			SetSize("small").
			SetStyle(map[string]interface{}{
				"marginLeft":   "50px",
				"marginBottom": "20px",
			})

		return ctx.JSON(200, component)
	}
}

// 将表格数据转换成表单数据
func (p *ImportRequest) transformFormValues(fields interface{}, data []interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range fields.([]interface{}) {
		if data[k] != nil {
			name := reflect.
				ValueOf(v).
				Elem().
				FieldByName("Name").String()
			result[name] = data[k]
		}
	}

	return result
}

// 获取提交表单的数据
func (p *ImportRequest) getSubmitData(fields interface{}, submitData interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, field := range fields.([]interface{}) {
		component := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Component").String()

		name := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Name").String()

		switch component {
		case "inputNumberField":
			result[name] = submitData.(map[string]interface{})[name]
		case "textField":
			result[name] = strings.Trim(submitData.(map[string]interface{})[name].(string), "\n")
		case "selectField":
			optionValue := field.(interface {
				GetOptionValue(label string) interface{}
			}).GetOptionValue(submitData.(map[string]interface{})[name].(string))

			result[name] = optionValue
		case "checkboxField":
			optionValue := field.(interface {
				GetOptionValue(label string) interface{}
			}).GetOptionValue(submitData.(map[string]interface{})[name].(string))

			result[name] = optionValue
		case "radioField":
			optionValue := field.(interface {
				GetOptionValue(label string) interface{}
			}).GetOptionValue(submitData.(map[string]interface{})[name].(string))

			result[name] = optionValue
		case "switchField":
			optionValue := field.(interface {
				GetOptionValue(label string) bool
			}).GetOptionValue(submitData.(map[string]interface{})[name].(string))

			result[name] = optionValue
		default:
			result[name] = submitData.(map[string]interface{})[name]
		}

		if getValue, ok := result[name].([]interface{}); ok {
			result[name], _ = json.Marshal(getValue)
		}

		if getValue, ok := result[name].([]map[string]interface{}); ok {
			result[name], _ = json.Marshal(getValue)
		}

		if getValue, ok := result[name].(map[string]interface{}); ok {
			result[name], _ = json.Marshal(getValue)
		}
	}

	return result
}

// if file or directory exits
func isExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
