package file

// IsFile 判断所给路径是否为文件[PHP:is_file]
func IsFile(path string) bool {
	f, err := Info(path)
	if err == nil {
		return !f.IsDir()
	}
	return false
}
