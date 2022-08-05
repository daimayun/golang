package file

import "os"

// Exists 判断文件或目录是否存在[PHP:file_exists]
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}