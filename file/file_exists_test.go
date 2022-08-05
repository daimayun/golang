package file

import "testing"

func TestExists(t *testing.T) {
	file := "./file_exists.go"
	if actual := Exists(file); actual == false {
		t.Error("Exists() error.")
	}
	file = "../file"
	if actual := Exists(file); actual == false {
		t.Error("Exists() error.")
	}
}
