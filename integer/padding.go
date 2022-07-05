package integer

import (
	"fmt"
	"strconv"
)

// Int64LengthPadding 数字长度不够左侧填补0并返回字符串
func Int64LengthPadding(i64 int64, length int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"d", i64)
}
