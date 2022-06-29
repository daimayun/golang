package conv

import "strconv"

// StringToInt string转int
func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// StringToInt64 string转int64
func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// StringToFloat64 string转float64
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
