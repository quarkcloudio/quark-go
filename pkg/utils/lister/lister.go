package lister

import (
	"encoding/json"
	"errors"
	"strings"
)

// 数据集转换成Tree
func ListToTree(list interface{}, pk string, pid string, child string, root int) (treeList []interface{}, Error error) {
	result := []map[string]interface{}{}
	jsonBytes, err := json.Marshal(list)
	if err != nil {
		return treeList, err
	}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return treeList, err
	}

	for _, v := range result {
		if int(v[pid].(float64)) == root {
			childNode, err := ListToTree(list, pk, pid, child, int(v[pk].(float64)))
			if err != nil {
				return treeList, err
			}
			if childNode != nil {
				v[child] = childNode
			}
			treeList = append(treeList, v)
		}
	}

	return treeList, nil
}

// Tree转换为有序列表
func TreeToOrderedList(tree []interface{}, level int, field string, child string) (list []interface{}, Error error) {
	for _, v := range tree {
		if value, ok := v.(map[string]interface{}); ok {
			value[field] = strings.Repeat("—", level) + value[field].(string)
			list = append(list, value)
			if value[child] != nil {
				if childValue, ok := value[child].([]interface{}); ok {
					children, err := TreeToOrderedList(childValue, level+1, field, child)
					if err != nil {
						return list, err
					}
					list = append(list, children...)
				} else {
					return list, errors.New("格式错误")
				}
			}
		} else {
			return list, errors.New("格式错误")
		}
	}

	return list, nil
}
