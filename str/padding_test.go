package str

import "testing"

func TestStringLengthPadding(t *testing.T) {
	str := "123456"
	if actual := StringLengthPadding(str, 12); actual != "000000123456" {
		t.Error("StringLengthPadding() error.")
	}
	if actual := StringLengthPadding(str, 5); actual != "123456" {
		t.Error("StringLengthPadding() error.")
	}
}
