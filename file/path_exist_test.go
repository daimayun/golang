package file

import "testing"

func TestPathExist(t *testing.T) {
	file := "../file/path_exist.go"
	if actual := Exists(file); actual == false {
		t.Error("PathExist() error.")
	}
}
