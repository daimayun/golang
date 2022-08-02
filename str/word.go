package str

import (
	"strings"
	"unicode"
)

// ToBigCamelCase 字符串转大驼峰格式
func ToBigCamelCase(str string, signs ...string) (res string) {
	sign := "_"
	if len(signs) > 0 {
		sign = signs[0]
	}
	strArr := strings.Split(str, sign)
	for _, v := range strArr {
		res += strings.ToUpper(string(v[0])) + v[1:]
	}
	return
}

// ToSmallCamelCase 字符串转小驼峰格式
func ToSmallCamelCase(str string, signs ...string) (res string) {
	sign := "_"
	if len(signs) > 0 {
		sign = signs[0]
	}
	strArr := strings.Split(str, sign)
	for k, v := range strArr {
		if k == 0 {
			res += strings.ToLower(string(v[0])) + v[1:]
		} else {
			res += strings.ToUpper(string(v[0])) + v[1:]
		}
	}
	return
}

// ToSnakeString XxYy to xx_yy , XxYY to xx_y_y
func ToSnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data))
}

// UcFirst 仅开头字母大写[将字符串中的第一个字母转换成大写]
func UcFirst(str string) string {
	for i, v := range str {
		if v >= 97 && v <= 122 {
			return string(unicode.ToUpper(v)) + str[i+1:]
		}
		return str
	}
	return ""
}

// LcFirst 仅开头字母小写[将字符串中的第一个字母转换成小写]
func LcFirst(str string) string {
	for i, v := range str {
		if v >= 65 && v <= 90 {
			return string(unicode.ToLower(v)) + str[i+1:]
		}
		return str
	}
	return ""
}

// UcWords 所有首字母大写[将字符串中每个单词的首字母转换成大写]
func UcWords(str string) (res string) {
	arr := strings.Split(str, " ")
	sign := ""
	for _, v := range arr {
		res += sign + UcFirst(v)
		sign = " "
	}
	return
}

// LcWords 所有首字母小写[将字符串中每个单词的首字母转换成小写]
func LcWords(str string) (res string) {
	arr := strings.Split(str, " ")
	sign := ""
	for _, v := range arr {
		res += sign + LcFirst(v)
		sign = " "
	}
	return
}

// GetFirstAndEndString 获取字符串的第一个和最后一个字符[不区分中英文]
func GetFirstAndEndString(str string) (string, string) {
	if len(str) == 0 {
		return "", ""
	}
	slice := []rune(str)
	return string(slice[0]), string(slice[len(slice)-1])
}
