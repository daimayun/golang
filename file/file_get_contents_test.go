package file

import "testing"

func TestGetContents(t *testing.T) {
	TestPutContents(t)
	TestPutContentsToAppend(t)
	fileName := "./test.txt"
	data := "hello world!你好，世界！"
	if actualData, actualOk := GetContents(fileName); actualData != data || actualOk != nil {
		t.Error("GetContents() error.")
	}
}
