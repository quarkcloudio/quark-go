package hex

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/go-basic/uuid"
)

// 生成Hex字符串
func Make(key string, crypt bool) string {
	if key == "" {
		key = uuid.New()
	}

	if crypt {
		h := md5.New()
		h.Write([]byte(key))
		key = hex.EncodeToString(h.Sum(nil))
	}

	return key
}
