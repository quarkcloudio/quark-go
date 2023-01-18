package fields

import (
	"github.com/quarkcms/quark-go/pkg/component/admin/component"
)

type Switch struct {
	Item
}

// 初始化
func (p *Switch) Init() *Switch {
	p.Component = "switchField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

/**
 * 设置开状态文字
 *
 * @param  string $value
 * @return $this
 */
func (p *Switch) SetTrueValue(value string) *Switch {
	var off interface{}
	options, ok := p.Options.(map[string]interface{})

	if ok {
		off = options["off"]
	}

	p.Options = map[string]interface{}{
		"on":  value,
		"off": off,
	}

	return p
}

/**
 * 设置关状态文字
 *
 * @param  string $value
 * @return $this
 */
func (p *Switch) SetFalseValue(value string) *Switch {
	var on interface{}
	options, ok := p.Options.(map[string]interface{})

	if ok {
		on = options["on"]
	}

	p.Options = map[string]interface{}{
		"on":  on,
		"off": value,
	}

	return p
}
