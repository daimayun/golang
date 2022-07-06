package function

import (
	"testing"
)

func TestGetPointer(t *testing.T) {
	var str string = "hello world"
	if actual := GetPointer(str); GetPointerValue(actual) != GetPointerValue(&str) {
		t.Error("GetPointer() error.")
	}
	var i int64 = 123
	if actual := GetPointer(i); GetPointerValue(actual) != GetPointerValue(&i) {
		t.Error("GetPointer() error.")
	}
	var ui uint = 123
	if actual := GetPointer(ui); GetPointerValue(actual) != GetPointerValue(&ui) {
		t.Error("GetPointer() error.")
	}
	var f float64 = 123.321
	if actual := GetPointer(f); GetPointerValue(actual) != GetPointerValue(&f) {
		t.Error("GetPointer() error.")
	}
	var b bool = true
	if actual := GetPointer(b); GetPointerValue(actual) != GetPointerValue(&b) {
		t.Error("GetPointer() error.")
	}
}

func TestGetPointerValue(t *testing.T) {
	var str string = "hello world"
	if actual := GetPointerValue(&str); actual != str {
		t.Error("GetPointerValue() error.")
	}
	var i int64 = 123
	if actual := GetPointerValue(&i); actual != i {
		t.Error("GetPointerValue() error.")
	}
	var ui uint = 123
	if actual := GetPointerValue(&ui); actual != ui {
		t.Error("GetPointerValue() error.")
	}
	var f float64 = 123.321
	if actual := GetPointerValue(&f); actual != f {
		t.Error("GetPointerValue() error.")
	}
	var b bool = true
	if actual := GetPointerValue(&b); actual != b {
		t.Error("GetPointerValue() error.")
	}
}
