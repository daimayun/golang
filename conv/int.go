package conv

import "strconv"

// IntToInt64 int转int64
func IntToInt64(i int) int64 {
	return int64(i)
}

// IntToString int转string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int64ToInt int64转int
func Int64ToInt(i64 int64) int {
	i, _ := strconv.Atoi(strconv.FormatInt(i64, 10))
	return i
}

// Int64ToString int64转string
func Int64ToString(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

// Int64ToFloat64 int64转float64
func Int64ToFloat64(i64 int64) float64 {
	return float64(i64)
}
