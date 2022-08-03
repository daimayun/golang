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

// Int8ToInt64 int8转int64
func Int8ToInt64(i8 int8) int64 {
	return int64(i8)
}

// Int16ToInt64 int16转int64
func Int16ToInt64(i16 int16) int64 {
	return int64(i16)
}

// Int32ToInt64 int32转int64
func Int32ToInt64(i32 int32) int64 {
	return int64(i32)
}

// UintToInt uint转int
func UintToInt(ui uint) int {
	return int(ui)
}

// Uint8ToInt64 uint8转int64
func Uint8ToInt64(ui8 uint8) int64 {
	return int64(ui8)
}

// Uint16ToInt64 uint16转int64
func Uint16ToInt64(ui16 uint16) int64 {
	return int64(ui16)
}

// Uint32ToInt64 uint32转int64
func Uint32ToInt64(ui32 uint32) int64 {
	return int64(ui32)
}

// Uint64ToInt64 uint64转int64
func Uint64ToInt64(ui64 uint64) int64 {
	return int64(ui64)
}
