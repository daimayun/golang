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
	file = "../conv"
	if actual := IsFile(file); actual == true {
		t.Error("IsFile() error.")
	}
}

func TestIsDir(t *testing.T) {
	file := "./file_test.go"
	if actual := IsDir(file); actual == true {
		t.Error("IsDir() error.")
	}
	file = "../conv"
	if actual := IsDir(file); actual == false {
		t.Error("IsDir() error.")
	}
}

func TestFilePutContents(t *testing.T) {
	fileName := "./test.txt"
	data := "hello world!"
	if actual := FilePutContents(fileName, data); actual != nil {
		t.Error("FilePutContents() error.")
	}
}

func TestFilePutContentsToAppend(t *testing.T) {
	fileName := "./test.txt"
	data := "你好，世界！"
	if actual := FilePutContentsToAppend(fileName, data); actual != nil {
		t.Error("FilePutContentsToAppend() error.")
	}
}

func TestFileGetContents(t *testing.T) {
	fileName := "./test.txt"
	data := "hello world!你好，世界！"
	if actualData, actualOk := FileGetContents(fileName); actualData != data || actualOk != nil {
		t.Error("FileGetContents() error.")
	}
}
