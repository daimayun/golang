package function

import "testing"

func TestFileExists(t *testing.T) {
	file := "./file_test.go"
	if actual := FileExists(file); actual == false {
		t.Error("FileExists() error.")
	}
	file = "../conv"
	if actual := FileExists(file); actual == false {
		t.Error("FileExists() error.")
	}
}

func TestIsFile(t *testing.T) {
	file := "./file_test.go"
	if actual := IsFile(file); actual == false {
		t.Error("IsFile() error.")
	}
}
