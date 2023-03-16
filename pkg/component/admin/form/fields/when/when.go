package when

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/go-basic/uuid"
	"github.com/quarkcms/quark-go/pkg/component/admin/component"
)

type Item struct {
	Component         string      `json:"component"`          // 组件名称
	Condition         string      `json:"condition"`          // 条件：js表达式语句
	ConditionName     string      `json:"condition_name"`     // 需要对比的字段名称
	ConditionOperator string      `json:"condition_operator"` // 操作符，= <>
	Option            interface{} `json:"option"`             // 条件符合的属性值
	Body              interface{} `json:"body"`               // 内容
}

type When struct {
	ComponentKey string  `json:"componentkey"` // 组件标识
	Component    string  `json:"component"`    // 组件名称
	Items        []*Item `json:"items"`        // When组件中需要解析的元素
}

// 初始化组件
func New() *When {
	return (&When{}).Init()
}

// 获取Item
func NewItem() *Item {
	return &Item{}
}

// 初始化
func (p *When) Init() *When {
	p.Component = "when"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *When) SetKey(key string, crypt bool) *When {
	if key == "" {
		key = uuid.New()
	}

	if crypt {
		h := md5.New()
		h.Write([]byte(key))
		key = hex.EncodeToString(h.Sum(nil))
	}

	p.ComponentKey = key

	return p
}

// 设置When组件中需要解析的元素
func (p *When) SetItems(items []*Item) *When {
	p.Items = items
	return p
}
