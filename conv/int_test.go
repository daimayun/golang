package conv

import "testing"

func TestInt64ToString(t *testing.T) {
	var (
		i64 int64  = 1234567890
		str string = "1234567890"
	)
	if actual := Int64ToString(i64); actual != str {
		t.Error("Int64ToString() error.")
	}
}

func TestInt64ToFloat64(t *testing.T) {
	var (
		i64 int64   = 1234567890
		f64 float64 = 1234567890
	)
	if actual := Int64ToFloat64(i64); actual != f64 {
		t.Error("Int64ToFloat64() error.")
	}
}

func TestInt64ToInt(t *testing.T) {
	var (
		i64 int64 = 1234567890
		i   int   = 1234567890
	)
	if actual := Int64ToInt(i64); actual != i {
		t.Error("Int64ToInt() error.")
	}
}

func TestIntToInt64(t *testing.T) {
	var (
		i64 int64 = 1234567890
		i   int   = 1234567890
	)
	if actual := IntToInt64(i); actual != i64 {
		t.Error("IntToInt64() error.")
	}
}

func TestIntToString(t *testing.T) {
	var (
		i   int    = 1234567890
		str string = "1234567890"
	)
	if actual := IntToString(i); actual != str {
		t.Error("IntToString() error.")
	}
}
