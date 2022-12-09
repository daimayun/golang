package file

import "path/filepath"

// PathExist 路径是否存在
func PathExist(filePath string) bool {
	indexMatches, err := filepath.Glob(filePath)
	return err == nil && len(indexMatches) > 0
}
