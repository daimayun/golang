package file

// Size 文件大小
func Size(filePath string) (int64, error) {
	fileInfo, err := Info(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
