package str

import "testing"

func TestCheckStringLength(t *testing.T) {
	str := "hello world! 你好，世界！"
	if actual := CheckStringLength(str, 19); actual != true {
		t.Error("CheckStringLength() error.")
	}
	str = ""
	if actual := CheckStringLength(str, 0); actual != true {
		t.Error("CheckStringLength() error.")
	}
}

func TestStringLength(t *testing.T) {
	str := "hello world! 你好，世界！"
	if actual := StringLength(str); actual != 19 {
		t.Error("StringLength() error.")
	}
	str = ""
	if actual := StringLength(str); actual != 0 {
		t.Error("StringLength() error.")
	}
}

func TestStringLengthCutAndSplitJoint(t *testing.T) {
	str := "hello world! 你好，世界！"
	if actual := StringLengthCutAndSplitJoint(str, 5, ", Ben."); actual != "hello, Ben." {
		t.Error("StringLengthCutAndSplitJoint() error.")
	}
	if actual := StringLengthCutAndSplitJoint(str, 16, "奔！"); actual != "hello world! 你好，奔！" {
		t.Error("StringLengthCutAndSplitJoint() error.")
	}
}
