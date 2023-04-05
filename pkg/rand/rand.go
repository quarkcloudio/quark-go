package rand

import (
	"math/rand"
	"strings"
	"time"
)

// 生成随机字符串
func Make(strtype string, length int) string {
	var result string
	switch strtype {
	case "numeric":
		result = MakeNumeric(length)
	case "letter":
		result = MakeLetter(length)
	case "alphanumeric":
		result = MakeAlphanumeric(length)
	default:
		result = MakeAlphanumeric(length)
	}

	return result
}

// 生成字母类型字符串
func MakeLetter(length int) string {
	seed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return makeRand(seed, length)
}

// 生成数字类型字符串
func MakeNumeric(length int) string {
	seed := "1234567890"
	return makeRand(seed, length)
}

// 生成数字类型字符串
func MakeAlphanumeric(length int) string {
	seed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	return makeRand(seed, length)
}

// 随机函数
func makeRand(seed string, length int) string {
	if length < 1 {
		return ""
	}
	charArr := strings.Split(seed, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))

	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}

	return rchar
}
