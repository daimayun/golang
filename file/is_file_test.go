package file

import "testing"

func TestIsFile(t *testing.T) {
	file := "./is_file_test.go"
	if actual := IsFile(file); actual == false {
		t.Error("IsFile() error.")
	}
	file = "../file"
	if actual := IsFile(file); actual == true {
		t.Error("IsFile() error.")
	}
}
