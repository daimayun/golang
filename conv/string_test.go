package conv

import "testing"

func TestStringToFloat64(t *testing.T) {
	var (
		f64 float64 = 1234567.8909
		str string  = "1234567.8909"
	)
	if actualFloat, actualOk := StringToFloat64(str); actualFloat != f64 || actualOk != nil {
		t.Error("StringToFloat64() error.")
	}
}

func TestStringToInt64(t *testing.T) {
	var (
		i64 int64  = 12345678909
		str string = "12345678909"
	)
	if actualInt, actualOk := StringToInt64(str); actualInt != i64 || actualOk != nil {
		t.Error("StringToInt64() error.")
	}
}

func TestStringToInt(t *testing.T) {
	var (
		i   int    = 1234567
		str string = "1234567"
	)
	if actualInt, actualOk := StringToInt(str); actualInt != i || actualOk != nil {
		t.Error("StringToInt64() error.")
	}
}
