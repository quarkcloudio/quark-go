package file

import "os"

// 判断文件路径是否存在
func IsExist(path string) bool {

	// os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}

	return true
}
