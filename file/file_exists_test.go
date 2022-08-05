package file

import "testing"

func TestFileExists(t *testing.T) {
	file := "./file_exists.go"
	if actual := FileExists(file); actual == false {
		t.Error("FileExists() error.")
	}
	file = "../file"
	if actual := FileExists(file); actual == false {
		t.Error("FileExists() error.")
	}
}
