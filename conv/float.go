package conv

import "strconv"

// Float64ToString float64转string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Float64ToInt64 float64转int64
func Float64ToInt64(f float64) int64 {
	return int64(f)
}
