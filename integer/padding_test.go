package integer

import "testing"

func TestInt64LengthPadding(t *testing.T) {
	var i64 int64 = 123456
	if actual := Int64LengthPadding(i64, 12); actual != "000000123456" {
		t.Error("Int64LengthPadding() error.")
	}
	if actual := Int64LengthPadding(i64, 4); actual != "123456" {
		t.Error("Int64LengthPadding() error.")
	}
}
