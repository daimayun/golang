package file

import "testing"

func TestFileGetContents(t *testing.T) {
	TestFilePutContents(t)
	TestFilePutContentsToAppend(t)
	fileName := "./test.txt"
	data := "hello world!你好，世界！"
	if actualData, actualOk := FileGetContents(fileName); actualData != data || actualOk != nil {
		t.Error("FileGetContents() error.")
	}
}
