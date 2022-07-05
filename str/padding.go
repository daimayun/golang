package str

import (
	"fmt"
	"strconv"
)

// StringLengthPadding 字符串长度不够左侧填补0
func StringLengthPadding(str string, length int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"s", str)
}
