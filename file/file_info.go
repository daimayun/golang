package file

import "os"

// Info 文件信息
func Info(filePath string) (os.FileInfo, error) {
	return os.Stat(filePath)
}
