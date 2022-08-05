package file

import "testing"

func TestIsDir(t *testing.T) {
	file := "./is_dir_test.go"
	if actual := IsDir(file); actual == true {
		t.Error("IsDir() error.")
	}
	file = "../file"
	if actual := IsDir(file); actual == false {
		t.Error("IsDir() error.")
	}
}
