package str

import (
	"testing"
)

func TestGetFirstAndEndString(t *testing.T) {
	str := "hello world! 你好，世界"
	if actualFirst, actualSecond := GetFirstAndEndString(str); actualFirst != "h" || actualSecond != "界" {
		t.Error("GetFirstAndEndString() error.")
	}
	str = "我"
	if actualFirst, actualSecond := GetFirstAndEndString(str); actualFirst != "我" || actualSecond != "我" {
		t.Error("GetFirstAndEndString() error.")
	}
	str = ""
	if actualFirst, actualSecond := GetFirstAndEndString(str); actualFirst != "" || actualSecond != "" {
		t.Error("GetFirstAndEndString() error.")
	}
}
