package adminresource

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/derekstavis/go-qs"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
)

// 创建请求的验证器
func (p *Template) ValidatorForCreation(request *builder.Request, templateInstance interface{}, data map[string]interface{}) error {
	rules, messages := p.RulesForCreation(request, templateInstance)

	validator := p.Validator(rules, messages, data)

	p.afterValidation(request, validator)
	p.afterCreationValidation(request, validator)

	return validator
}

// 验证规则
func (p *Template) Validator(rules []interface{}, messages []interface{}, data map[string]interface{}) error {
	var result error

	for _, rule := range rules {
		for k, v := range rule.(map[string]interface{}) {
			fieldValue := data[k]
			for _, item := range v.([]interface{}) {
				getItem, ok := item.(string)
				if ok {
					getItems := strings.Split(getItem, ":")
					getOption := ""
					if len(getItems) == 2 {
						getItem = getItems[0]
						getOption = getItems[1]
					}

					switch getItem {
					case "required":
						if fieldValue == nil {
							errMsg := p.getRuleMessage(messages, k+"."+getItem)

							if errMsg != "" {
								result = errors.New(errMsg)
							}
						}
					case "min":
						strNum := utf8.RuneCountInString(fieldValue.(string))
						minOption, _ := strconv.Atoi(getOption)

						if strNum < minOption {
							errMsg := p.getRuleMessage(messages, k+"."+getItem)
							if errMsg != "" {
								result = errors.New(errMsg)
							}
						}
					case "max":
						strNum := utf8.RuneCountInString(fieldValue.(string))
						maxOption, _ := strconv.Atoi(getOption)

						if strNum > maxOption {
							errMsg := p.getRuleMessage(messages, k+"."+getItem)
							if errMsg != "" {
								result = errors.New(errMsg)
							}
						}
					case "unique":
						var (
							table  string
							field  string
							except string
							count  int
						)

						uniqueOptions := strings.Split(getOption, ",")

						if len(uniqueOptions) == 2 {
							table = uniqueOptions[0]
							field = uniqueOptions[1]
						}

						if len(uniqueOptions) == 3 {
							table = uniqueOptions[0]
							field = uniqueOptions[1]
							except = uniqueOptions[2]
						}

						if except != "" {
							db.Client.Raw("SELECT Count("+field+") FROM "+table+" WHERE id <> "+except+" AND "+field+" = ?", fieldValue).Scan(&count)
						} else {
							db.Client.Raw("SELECT Count("+field+") FROM "+table+" WHERE "+field+" = ?", fieldValue).Scan(&count)
						}

						if count > 0 {
							errMsg := p.getRuleMessage(messages, k+"."+getItem)
							if errMsg != "" {
								result = errors.New(errMsg)
							}
						}
					}
				}
			}
		}
	}

	return result
}

// 获取规则错误信息
func (p *Template) getRuleMessage(messages []interface{}, key string) string {
	msg := ""

	for _, v := range messages {
		getMsg := v.(map[string]interface{})[key]
		if getMsg != nil {
			msg = getMsg.(string)
		}
	}

	return msg
}

// 创建请求的验证规则
func (p *Template) RulesForCreation(request *builder.Request, templateInstance interface{}) ([]interface{}, []interface{}) {

	fields := templateInstance.(interface {
		CreationFieldsWithoutWhen(*builder.Request, interface{}) interface{}
	}).CreationFieldsWithoutWhen(request, templateInstance)

	rules := []interface{}{}
	ruleMessages := []interface{}{}

	for _, v := range fields.([]interface{}) {
		getResult := p.getRulesForCreation(request, v)

		if len(getResult["rules"].(map[string]interface{})) > 0 {
			rules = append(rules, getResult["rules"])
		}

		if len(getResult["messages"].(map[string]interface{})) > 0 {
			ruleMessages = append(ruleMessages, getResult["messages"])
		}

		when := reflect.
			ValueOf(v).
			Elem().
			FieldByName("When").Interface()

		if when != nil {
			whenItems := reflect.
				ValueOf(when).
				Elem().
				FieldByName("Items").Interface()

			if whenItems != nil {
				for _, vi := range whenItems.([]map[string]interface{}) {
					if p.needValidateWhenRules(request, vi) {
						body := vi["body"]
						if body != nil {
							// 如果为数组
							getBody, ok := body.([]interface{})
							if ok {
								for _, bv := range getBody {
									whenItemResult := p.getRulesForCreation(request, bv)

									if len(whenItemResult["rules"].(map[string]interface{})) > 0 {
										rules = append(rules, whenItemResult["rules"])
									}

									if len(whenItemResult["messages"].(map[string]interface{})) > 0 {
										ruleMessages = append(ruleMessages, whenItemResult["messages"])
									}
								}
							} else {
								whenItemResult := p.getRulesForCreation(request, getBody)

								if len(whenItemResult["rules"].(map[string]interface{})) > 0 {
									rules = append(rules, whenItemResult["rules"])
								}

								if len(whenItemResult["messages"].(map[string]interface{})) > 0 {
									ruleMessages = append(ruleMessages, whenItemResult["messages"])
								}
							}
						}
					}
				}
			}
		}

	}

	return rules, ruleMessages
}

// 判断是否需要验证When组件里的规则
func (p *Template) needValidateWhenRules(request *builder.Request, when map[string]interface{}) bool {
	conditionName := when["condition_name"].(string)
	conditionOption := when["condition_option"]
	conditionOperator := when["condition_operator"].(string)
	result := false

	data, error := qs.Unmarshal(request.OriginalURL())
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
func (p *Template) getRulesForCreation(request *builder.Request, field interface{}) map[string]interface{} {
	getRules := map[string]interface{}{}
	getRuleMessages := map[string]interface{}{}

	name := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Name").String()

	rules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Rules").Interface()

	ruleMessages := reflect.
		ValueOf(field).
		Elem().
		FieldByName("RuleMessages").Interface()

	creationRules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("CreationRules").Interface()

	creationRuleMessages := reflect.
		ValueOf(field).
		Elem().
		FieldByName("CreationRuleMessages").Interface()

	items := []interface{}{}

	for _, v := range p.formatRules(request, rules.([]string)) {
		items = append(items, v)
	}

	for key, v := range ruleMessages.(map[string]string) {
		getRuleMessages[name+"."+key] = v
	}

	for _, v := range p.formatRules(request, creationRules.([]string)) {
		items = append(items, v)
	}

	for key, v := range creationRuleMessages.(map[string]string) {
		getRuleMessages[name+"."+key] = v
	}

	if len(items) > 0 {
		getRules[name] = items
	}

	return map[string]interface{}{
		"rules":    getRules,
		"messages": getRuleMessages,
	}
}

// 更新请求的验证器
func (p *Template) ValidatorForUpdate(request *builder.Request, templateInstance interface{}, data map[string]interface{}) error {
	rules, messages := p.RulesForUpdate(request, templateInstance)

	validator := p.Validator(rules, messages, data)

	p.afterValidation(request, validator)
	p.afterCreationValidation(request, validator)

	return validator
}

// 更新请求的验证规则
func (p *Template) RulesForUpdate(request *builder.Request, templateInstance interface{}) ([]interface{}, []interface{}) {

	fields := templateInstance.(interface {
		UpdateFieldsWithoutWhen(*builder.Request, interface{}) interface{}
	}).UpdateFieldsWithoutWhen(request, templateInstance)

	rules := []interface{}{}
	ruleMessages := []interface{}{}

	for _, v := range fields.([]interface{}) {
		getResult := p.getRulesForUpdate(request, v)

		if len(getResult["rules"].(map[string]interface{})) > 0 {
			rules = append(rules, getResult["rules"])
		}

		if len(getResult["messages"].(map[string]interface{})) > 0 {
			ruleMessages = append(ruleMessages, getResult["messages"])
		}

		when := reflect.
			ValueOf(v).
			Elem().
			FieldByName("When").Interface()

		if when != nil {
			whenItems := reflect.
				ValueOf(when).
				Elem().
				FieldByName("Items").Interface()

			if whenItems != nil {
				for _, vi := range whenItems.([]map[string]interface{}) {
					if p.needValidateWhenRules(request, vi) {
						body := vi["body"]

						if body != nil {

							// 如果为数组
							getBody, ok := body.([]interface{})
							if ok {
								for _, bv := range getBody {
									whenItemResult := p.getRulesForUpdate(request, bv)

									if len(whenItemResult["rules"].(map[string]interface{})) > 0 {
										rules = append(rules, whenItemResult["rules"])
									}

									if len(whenItemResult["messages"].(map[string]interface{})) > 0 {
										ruleMessages = append(ruleMessages, whenItemResult["messages"])
									}
								}
							} else {
								whenItemResult := p.getRulesForUpdate(request, getBody)

								if len(whenItemResult["rules"].(map[string]interface{})) > 0 {
									rules = append(rules, whenItemResult["rules"])
								}

								if len(whenItemResult["messages"].(map[string]interface{})) > 0 {
									ruleMessages = append(ruleMessages, whenItemResult["messages"])
								}
							}
						}
					}
				}
			}
		}

	}

	return rules, ruleMessages
}

// 获取更新请求资源规则
func (p *Template) getRulesForUpdate(request *builder.Request, field interface{}) map[string]interface{} {

	getRules := map[string]interface{}{}
	getRuleMessages := map[string]interface{}{}

	name := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Name").String()

	rules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("Rules").Interface()

	ruleMessages := reflect.
		ValueOf(field).
		Elem().
		FieldByName("RuleMessages").Interface()

	updateRules := reflect.
		ValueOf(field).
		Elem().
		FieldByName("UpdateRules").Interface()

	updateRuleMessages := reflect.
		ValueOf(field).
		Elem().
		FieldByName("UpdateRuleMessages").Interface()

	items := []interface{}{}

	for _, v := range p.formatRules(request, rules.([]string)) {
		items = append(items, v)
	}

	for key, v := range ruleMessages.(map[string]string) {
		getRuleMessages[name+"."+key] = v
	}

	for _, v := range p.formatRules(request, updateRules.([]string)) {
		items = append(items, v)
	}

	for key, v := range updateRuleMessages.(map[string]string) {
		getRuleMessages[name+"."+key] = v
	}

	if len(items) > 0 {
		getRules[name] = items
	}

	return map[string]interface{}{
		"rules":    getRules,
		"messages": getRuleMessages,
	}
}

// 导入请求的验证器
func (p *Template) ValidatorForImport(request *builder.Request, templateInstance interface{}, data map[string]interface{}) error {
	rules, messages := p.RulesForImport(request, templateInstance)

	validator := p.Validator(rules, messages, data)

	p.afterValidation(request, validator)
	p.afterCreationValidation(request, validator)

	return validator
}

// 创建请求的验证规则
func (p *Template) RulesForImport(request *builder.Request, templateInstance interface{}) ([]interface{}, []interface{}) {

	fields := templateInstance.(interface {
		ImportFieldsWithoutWhen(*builder.Request, interface{}) interface{}
	}).ImportFieldsWithoutWhen(request, templateInstance)

	rules := []interface{}{}
	ruleMessages := []interface{}{}

	for _, v := range fields.([]interface{}) {
		getResult := p.getRulesForCreation(request, v)

		if len(getResult["rules"].(map[string]interface{})) > 0 {
			rules = append(rules, getResult["rules"])
		}

		if len(getResult["messages"].(map[string]interface{})) > 0 {
			ruleMessages = append(ruleMessages, getResult["messages"])
		}

		when := reflect.
			ValueOf(v).
			Elem().
			FieldByName("When").Interface()

		if when != nil {
			whenItems := reflect.
				ValueOf(when).
				Elem().
				FieldByName("Items").Interface()

			if whenItems != nil {
				for _, vi := range whenItems.([]map[string]interface{}) {
					if p.needValidateWhenRules(request, vi) {
						body := vi["body"]

						if body != nil {

							// 如果为数组
							getBody, ok := body.([]interface{})
							if ok {
								for _, bv := range getBody {
									whenItemResult := p.getRulesForCreation(request, bv)

									if len(whenItemResult["rules"].(map[string]interface{})) > 0 {
										rules = append(rules, whenItemResult["rules"])
									}

									if len(whenItemResult["messages"].(map[string]interface{})) > 0 {
										ruleMessages = append(ruleMessages, whenItemResult["messages"])
									}
								}
							} else {
								whenItemResult := p.getRulesForCreation(request, getBody)

								if len(whenItemResult["rules"].(map[string]interface{})) > 0 {
									rules = append(rules, whenItemResult["rules"])
								}

								if len(whenItemResult["messages"].(map[string]interface{})) > 0 {
									ruleMessages = append(ruleMessages, whenItemResult["messages"])
								}
							}
						}
					}
				}
			}
		}

	}

	return rules, ruleMessages
}

// 格式化规则
func (p *Template) formatRules(request *builder.Request, rules []string) []string {
	data := map[string]interface{}{}
	json.Unmarshal(request.Body(), &data)

	formId := data["id"]
	requestId := request.Query("id", "")
	if requestId.(string) == "" && formId == nil {
		return rules
	}

	if requestId != "" {
		for key, v := range rules {
			rules[key] = strings.Replace(v, "{id}", requestId.(string), -1)
		}
	} else if formId != nil {
		for key, v := range rules {
			requestId = strconv.FormatFloat(formId.(float64), 'E', -1, 64)
			rules[key] = strings.Replace(v, "{id}", requestId.(string), -1)
		}
	}

	return rules
}

// 验证完成后回调
func (p *Template) afterValidation(request *builder.Request, validator interface{}) {
	//
}

// 创建请求验证完成后回调
func (p *Template) afterCreationValidation(request *builder.Request, validator interface{}) {
	//
}

// 更新请求验证完成后回调
func (p *Template) afterUpdateValidation(request *builder.Request, validator interface{}) {
	//
}

// 创建请求验证完成后回调
func (p *Template) afterImportValidation(request *builder.Request, validator interface{}) {
	//
}
