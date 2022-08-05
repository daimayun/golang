package file

import "testing"

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
