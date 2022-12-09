package file

// IsDir 判断所给路径是否为文件夹[PHP:is_dir]
func IsDir(path string) bool {
	f, err := Info(path)
	if err == nil {
		return f.IsDir()
	}
	return false
}
