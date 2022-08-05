package file

import (
	"io/ioutil"
	"os"
)

// PutContents 把一个字符串写入文件中[覆盖原文件内容][PHP:file_put_contents]
func PutContents(filename string, data string) error {
	return PutContentsByByte(filename, []byte(data))
}

// PutContentsToAppend 把一个字符串写入文件中[追加至原文件][PHP:file_put_contents,FILE_APPEND]
func PutContentsToAppend(filename string, data string) error {
	return PutContentsByByteToAppend(filename, []byte(data))
}

// PutContentsByByte 把byte写入文件中[覆盖原文件内容][PHP:file_put_contents]
func PutContentsByByte(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

// PutContentsByByteToAppend 把byte写入文件中[追加至原文件][PHP:file_put_contents,FILE_APPEND]
func PutContentsByByteToAppend(filename string, data []byte) (err error) {
	var f *os.File
	f, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write(data)
	return
}
