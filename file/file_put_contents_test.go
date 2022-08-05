package file

import "testing"

func TestPutContents(t *testing.T) {
	fileName := "./test.txt"
	data := "hello world!"
	if actual := PutContents(fileName, data); actual != nil {
		t.Error("PutContents() error.")
	}
}

func TestPutContentsToAppend(t *testing.T) {
	fileName := "./test.txt"
	data := "你好，世界！"
	if actual := PutContentsToAppend(fileName, data); actual != nil {
		t.Error("PutContentsToAppend() error.")
	}
}
