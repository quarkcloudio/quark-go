package component

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/go-basic/uuid"
)

type Element struct {
	Key          string      `json:"key"`
	ComponentKey string      `json:"componentKey"`
	Component    string      `json:"component"`
	Style        interface{} `json:"style"`
}

const DEFAULT_KEY = ""
const DEFAULT_CRYPT = true

// 设置Key
func (p *Element) SetKey(key string, crypt bool) *Element {

	if key == "" {
		key = uuid.New()
	}

	if crypt {
		h := md5.New()
		h.Write([]byte(key))
		key = hex.EncodeToString(h.Sum(nil))
	}

	p.Key = key
	p.ComponentKey = key

	return p
}
