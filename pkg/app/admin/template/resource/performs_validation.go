package resource

import (
	"encoding/json"
	"errors"
	"strings"
	"unicode/utf8"

	"github.com/derekstavis/go-qs"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/fields/when"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/component/form/rule"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
)

// 创建请求的验证器
func (p *Template) ValidatorForCreation(ctx *builder.Context, data map[string]interface{}) error {

	// 获取创建数据验证规则
	rules := p.RulesForCreation(ctx)

	// 验证数据是否合法
	validator := p.Validator(rules, data)

	// 验证成功后回调
	p.AfterValidation(ctx, validator)

	// 创建请求验证完成后回调
	p.AfterCreationValidation(ctx, validator)

	return validator
}

// 验证规则
func (p *Template) Validator(rules []*rule.Rule, data map[string]interface{}) error {
	var result error

	for _, rule := range rules {
		fieldValue := data[rule.Name]
		switch rule.RuleType {
		case "required":
			if fieldValue == nil {
				errMsg := rule.Message
				if errMsg != "" {
					result = errors.New(errMsg)
				}
			}
		case "min":
			if fieldValue, ok := fieldValue.(string); ok {
				strNum := utf8.RuneCountInString(fieldValue)
				if strNum < rule.Min {
					errMsg := rule.Message
					if errMsg != "" {
						result = errors.New(errMsg)
					}
				}
			}
		case "max":
			if fieldValue, ok := fieldValue.(string); ok {
				strNum := utf8.RuneCountInString(fieldValue)
				if strNum > rule.Max {
					errMsg := rule.Message
					if errMsg != "" {
						result = errors.New(errMsg)
					}
				}
			}
		case "unique":
			var (
				table  string
				field  string
				except string
				count  int64
			)

			table = rule.UniqueTable
			field = rule.UniqueTableField
			except = rule.UniqueIgnoreValue
			if except != "" {
				ignoreField := strings.ReplaceAll(except, "{", "")
				ignoreField = strings.ReplaceAll(ignoreField, "}", "")
				ignoreValue := data[ignoreField]

				db.Client.Table(table).Where(ignoreField+" <> ?", ignoreValue).Where(field+" = ?", fieldValue).Count(&count)
			} else {
				db.Client.Table(table).Where(field+" = ?", fieldValue).Count(&count)
			}

			if count > 0 {
				errMsg := rule.Message
				if errMsg != "" {
					result = errors.New(errMsg)
				}
			}
		}
	}

	return result
}

// 创建请求的验证规则
func (p *Template) RulesForCreation(ctx *builder.Context) (rules []*rule.Rule) {
	fields := ctx.Template.(interface {
		CreationFieldsWithoutWhen(*builder.Context) interface{}
	}).CreationFieldsWithoutWhen(ctx)

	for _, v := range fields.([]interface{}) {
		rules = append(rules, p.getRulesForCreation(v)...)

		if whenComponent, ok := v.(interface {
			GetWhen() *when.Component
		}); ok {
			getWhen := whenComponent.GetWhen()
			if getWhen != nil {
				// 获取When组件中的验证规则
				whenItems := getWhen.Items
				if len(whenItems) > 0 {
					for _, vi := range whenItems {
						if p.needValidateWhenRules(ctx, vi) {
							body := vi.Body
							if body != nil {
								// 如果为数组
								getBody, ok := body.([]interface{})
								if ok {
									for _, bv := range getBody {
										rules = append(rules, p.getRulesForCreation(bv)...)
									}
								} else {
									rules = append(rules, p.getRulesForCreation(getBody)...)
								}
							}
						}
					}
				}
			}
		}
	}

	return rules
}

// 判断是否需要验证When组件里的规则
func (p *Template) needValidateWhenRules(ctx *builder.Context, when *when.Item) bool {
	conditionName := when.ConditionName
	conditionOption := when.Option
	conditionOperator := when.ConditionOperator
	result := false

	data, error := qs.Unmarshal(ctx.OriginalURL())
	if error != nil {
		return false
	}

	value, ok := data[conditionName]
	if !ok {
		return false
	}

	valueString, isString := value.(string)
	if isString {
		if valueString == "" {
			return false
		}
	}

	switch conditionOperator {
	case "=":
		result = (value.(string) == conditionOption.(string))
	case ">":
		result = (value.(string) > conditionOption.(string))
	case "<":
		result = (value.(string) < conditionOption.(string))
	case "<=":
		result = (value.(string) <= conditionOption.(string))
	case ">=":
		result = (value.(string) >= conditionOption.(string))
	case "has":
		_, isArray := value.([]string)
		if isArray {
			getJson, err := json.Marshal(value)
			if err != nil {
				result = strings.Contains(string(getJson), conditionOption.(string))
			}
		} else {
			result = strings.Contains(value.(string), conditionOption.(string))
		}
	case "in":
		conditionOptionArray, isArray := conditionOption.([]string)
		if isArray {
			for _, v := range conditionOptionArray {
				if v == value.(string) {
					result = true
				}
			}
		}
	default:
		result = (value.(string) == conditionOption)
	}

	return result
}

// 获取创建请求资源规则
func (p *Template) getRulesForCreation(field interface{}) (rules []*rule.Rule) {
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

	return rules
}

// 更新请求的验证器
func (p *Template) ValidatorForUpdate(ctx *builder.Context, data map[string]interface{}) error {

	// 获取更新数据验证规则
	rules := p.RulesForUpdate(ctx)

	// 验证数据是否合法
	validator := p.Validator(rules, data)

	// 验证成功后回调
	p.AfterValidation(ctx, validator)

	// 编辑请求验证完成后回调
	p.AfterUpdateValidation(ctx, validator)

	return validator
}

// 更新请求的验证规则
func (p *Template) RulesForUpdate(ctx *builder.Context) (rules []*rule.Rule) {
	fields := ctx.Template.(interface {
		UpdateFieldsWithoutWhen(*builder.Context) interface{}
	}).UpdateFieldsWithoutWhen(ctx)

	for _, v := range fields.([]interface{}) {
		rules = append(rules, p.getRulesForUpdate(v)...)

		if whenComponent, ok := v.(interface {
			GetWhen() *when.Component
		}); ok {
			getWhen := whenComponent.GetWhen()
			if getWhen != nil {

				// 获取When组件中的验证规则
				whenItems := getWhen.Items
				if len(whenItems) > 0 {
					for _, vi := range whenItems {
						if p.needValidateWhenRules(ctx, vi) {
							body := vi.Body
							if body != nil {
								// 如果为数组
								getBody, ok := body.([]interface{})
								if ok {
									for _, bv := range getBody {
										rules = append(rules, p.getRulesForUpdate(bv)...)
									}
								} else {
									rules = append(rules, p.getRulesForUpdate(getBody)...)
								}
							}
						}
					}
				}
			}
		}
	}

	return rules
}

// 获取更新请求资源规则
func (p *Template) getRulesForUpdate(field interface{}) (rules []*rule.Rule) {
	if v, ok := field.(interface {
		GetRules() []*rule.Rule
	}); ok {
		rules = append(rules, v.GetRules()...)
	}

	if v, ok := field.(interface {
		GetUpdateRules() []*rule.Rule
	}); ok {
		rules = append(rules, v.GetUpdateRules()...)
	}

	return rules
}

// 导入请求的验证器
func (p *Template) ValidatorForImport(ctx *builder.Context, data map[string]interface{}) error {

	// 获取更新数据验证规则
	rules := p.RulesForImport(ctx)

	// 验证数据是否合法
	validator := p.Validator(rules, data)

	// 验证成功后回调
	p.AfterValidation(ctx, validator)

	// 	导入请求验证完成后回调
	p.AfterCreationValidation(ctx, validator)

	return validator
}

// 创建请求的验证规则
func (p *Template) RulesForImport(ctx *builder.Context) (rules []*rule.Rule) {

	fields := ctx.Template.(interface {
		ImportFieldsWithoutWhen(*builder.Context) interface{}
	}).ImportFieldsWithoutWhen(ctx)

	for _, v := range fields.([]interface{}) {
		rules = append(rules, p.getRulesForCreation(v)...)

		if whenComponent, ok := v.(interface {
			GetWhen() *when.Component
		}); ok {
			getWhen := whenComponent.GetWhen()
			if getWhen != nil {
				// 获取When组件中的验证规则
				whenItems := getWhen.Items
				if len(whenItems) > 0 {
					for _, vi := range whenItems {
						if p.needValidateWhenRules(ctx, vi) {
							body := vi.Body
							if body != nil {
								// 如果为数组
								getBody, ok := body.([]interface{})
								if ok {
									for _, bv := range getBody {
										rules = append(rules, p.getRulesForCreation(bv)...)
									}
								} else {
									rules = append(rules, p.getRulesForCreation(getBody)...)
								}
							}
						}
					}
				}
			}
		}
	}

	return rules
}

// 验证完成后回调
func (p *Template) AfterValidation(ctx *builder.Context, validator interface{}) {
	//
}

// 创建请求验证完成后回调
func (p *Template) AfterCreationValidation(ctx *builder.Context, validator interface{}) {
	//
}

// 更新请求验证完成后回调
func (p *Template) AfterUpdateValidation(ctx *builder.Context, validator interface{}) {
	//
}

// 创建请求验证完成后回调
func (p *Template) AfterImportValidation(ctx *builder.Context, validator interface{}) {
	//
}
