package str

import (
	"fmt"
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

func TestLcWords(t *testing.T) {
	str := "Hello world, 123 奔 and Ben. Where are you from. ."
	if actual := LcWords(str); actual != "hello world, 123 奔 and ben. where are you from. ." {
		fmt.Println(actual)
		t.Error("LcWords() error.")
	}
}
