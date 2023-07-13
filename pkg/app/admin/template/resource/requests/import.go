package requests

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/space"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/tpl"
	models "github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"github.com/quarkcms/quark-go/v2/pkg/rand"
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
		FieldByName("Model").
		Interface()
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

		// 获取表单数据
		formValues := p.transformFormValues(fields, item)

		// 验证表单条件
		validator := ctx.Template.(interface {
			ValidatorForImport(ctx *builder.Context, data map[string]interface{}) error
		}).ValidatorForImport(ctx, formValues)
		if validator != nil {
			importResult = false
			importFailedNum = importFailedNum + 1
			item = append(item, validator.Error())
			importFailedData = append(importFailedData, item)

			// 跳出本次循环
			continue
		}

		// 验证保存前回调条件
		submitData, err := ctx.Template.(interface {
			BeforeSaving(ctx *builder.Context, data map[string]interface{}) (map[string]interface{}, error)
		}).BeforeSaving(ctx, formValues)
		if err != nil {
			importResult = false
			importFailedNum = importFailedNum + 1
			item = append(item, err.Error())
			importFailedData = append(importFailedData, item)

			// 跳出本次循环
			continue
		}

		// 插入数据库
		data := p.getSubmitData(fields, submitData)
		result := model.Create(data)
		if result.Error != nil {
			importResult = false
			importFailedNum = importFailedNum + 1
			item = append(item, result.Error.Error())
			importFailedData = append(importFailedData, item)

			// 跳出本次循环
			continue
		}

		getLastData := map[string]interface{}{}
		model.Order("id desc").First(&getLastData)

		// 保存后回调
		err = ctx.Template.(interface {
			AfterSaved(ctx *builder.Context, id int, data map[string]interface{}, result *gorm.DB) error
		}).AfterSaved(ctx, getLastData["id"].(int), data, result)
		if err != nil {
			importResult = false
			importFailedNum = importFailedNum + 1
			item = append(item, err.Error())
			importFailedData = append(importFailedData, item)

			// 跳出本次循环
			continue
		}

		importSuccessedNum = importSuccessedNum + 1
	}

	// 返回导入失败错误数据
	if !importResult {
		filePath := ctx.Engine.GetConfig().StaticPath + "/app/storage/failImports/"
		fileName := rand.MakeAlphanumeric(40) + ".xlsx"
		fileUrl := "//" + ctx.Host() + "/storage/failImports/" + fileName

		// 不存在路径，则创建
		if isExist(filePath) == false {
			err := os.MkdirAll(filePath, 0666)
			if err != nil {
				return ctx.JSONError(err.Error())
			}
		}

		f := excelize.NewFile()

		// 创建Sheet
		index, _ := f.NewSheet("Sheet1")

		// 创建表头
		importHead = append(importHead, "错误信息")
		a := 'a'
		for i := 1; i <= len(importHead); i++ {
			f.SetCellValue("Sheet1", string(a)+"1", importHead[i-1])
			a++
		}

		// 创建数据
		for k, v := range importFailedData {
			a := 'a'
			for i := 1; i <= len(v); i++ {
				f.SetCellValue("Sheet1", string(a)+strconv.Itoa(k+2), v[i-1])
				a++
			}
		}

		f.SetActiveSheet(index)
		if err := f.SaveAs(filePath + fileName); err != nil {
			return ctx.JSONError(err.Error())
		}

		tpl1 := (&tpl.Component{}).
			Init().
			SetBody("导入总量: " + strconv.Itoa(importTotalNum))

		tpl2 := (&tpl.Component{}).
			Init().
			SetBody("成功数量: " + strconv.Itoa(importSuccessedNum))

		tpl3 := (&tpl.Component{}).
			Init().
			SetBody("失败数量: <span style='color:#ff4d4f'>" + strconv.Itoa(importFailedNum) + "</span> <a href='" + fileUrl + "' target='_blank'>下载失败数据</a>")

		component := (&space.Component{}).
			Init().
			SetBody([]interface{}{
				tpl1,
				tpl2,
				tpl3,
			}).
			SetDirection("vertical").
			SetSize("small").
			SetStyle(map[string]interface{}{
				"marginLeft":   "50px",
				"marginBottom": "20px",
			})

		return ctx.JSON(200, component)
	}

	return ctx.JSONOk("操作成功！", strings.Replace("/layout/index?api="+indexRoute, ":resource", ctx.Param("resource"), -1))
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
			FieldByName("Component").
			String()

		name := reflect.
			ValueOf(field).
			Elem().
			FieldByName("Name").
			String()

		switch component {
		case "inputNumberField":

			result[name] = submitData.(map[string]interface{})[name]
		case "textField":

			result[name] = strings.Trim(submitData.(map[string]interface{})[name].(string), "\n")
		case "selectField":

			result[name] = field.(interface {
				GetOptionValue(label string) interface{}
			}).GetOptionValue(submitData.(map[string]interface{})[name].(string))
		case "checkboxField":

			result[name] = field.(interface {
				GetOptionValue(label string) interface{}
			}).GetOptionValue(submitData.(map[string]interface{})[name].(string))
		case "radioField":

			result[name] = field.(interface {
				GetOptionValue(label string) interface{}
			}).GetOptionValue(submitData.(map[string]interface{})[name].(string))
		case "switchField":

			result[name] = field.(interface {
				GetOptionValue(label string) bool
			}).GetOptionValue(submitData.(map[string]interface{})[name].(string))
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

// 判断文件是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
