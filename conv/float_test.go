package conv

import "testing"

func TestFloat64ToInt64(t *testing.T) {
	var (
		i64 int64   = 1234567890
		f64 float64 = 1234567890
	)
	if actual := Float64ToInt64(f64); actual != i64 {
		t.Error("Float64ToInt64() error.")
	}
}

func TestFloat64ToString(t *testing.T) {
	var (
		f64 float64 = 1234567.8909
		str string  = "1234567.8909"
	)
	if actual := Float64ToString(f64); actual != str {
		t.Error("Float64ToString() error.")
	}
}
