package str

import "unicode/utf8"

// StringLength 获取字符串长度[包含中文字符串,一个中文字符串的长度为1]
func StringLength(str string) int {
	return utf8.RuneCountInString(str)
}

// StringLengthCutAndSplitJoint 字符串长度截取并拼接处理
func StringLengthCutAndSplitJoint(str string, cutLength int, splitJointStr ...string) string {
	if cutLength == 0 {
		return str
	}
	if StringLength(str) <= cutLength {
		return str
	}
	strRune := []rune(str)
	s := string(strRune[0:cutLength])
	if len(splitJointStr) > 0 {
		s += splitJointStr[0]
	}
	return s
}

// CheckStringLength 判断字符串长度是否在规定范围内
func CheckStringLength(str string, length int) bool {
	if StringLength(str) <= length {
		return true
	}
	return false
}
