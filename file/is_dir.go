package file

import "os"

// IsDir 判断所给路径是否为文件夹[PHP:is_dir]
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err == nil {
		return f.IsDir()
	}
	return false
}
