package orm

import (
	"errors"

	"gorm.io/gorm"
)

// IsEmpty 判断查询是否为空
func IsEmpty(err error) bool {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}
