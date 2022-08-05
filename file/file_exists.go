package file

import "os"

// FileExists 判断文件或目录是否存在[PHP:file_exists]
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
