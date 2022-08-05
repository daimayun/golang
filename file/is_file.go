package file

import "os"

// IsFile 判断所给路径是否为文件[PHP:is_file]
func IsFile(path string) bool {
	f, err := os.Stat(path)
	if err == nil {
		return !f.IsDir()
	}
	return false
}
