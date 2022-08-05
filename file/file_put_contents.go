package file

import (
	"io/ioutil"
	"os"
)

// FilePutContents 把一个字符串写入文件中[覆盖原文件内容][PHP:file_put_contents]
func FilePutContents(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}

// FilePutContentsToAppend 把一个字符串写入文件中[追加至原文件][PHP:file_put_contents,FILE_APPEND]
func FilePutContentsToAppend(filename string, data string) (err error) {
	var f *os.File
	f, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write([]byte(data))
	return
}

// FilePutContentsByByte 把byte写入文件中[覆盖原文件内容][PHP:file_put_contents]
func FilePutContentsByByte(filename string, b []byte) error {
	return ioutil.WriteFile(filename, b, 0644)
}

// FilePutContentsByByteToAppend 把byte写入文件中[追加至原文件][PHP:file_put_contents,FILE_APPEND]
func FilePutContentsByByteToAppend(filename string, b []byte) (err error) {
	var f *os.File
	f, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write(b)
	return
}
