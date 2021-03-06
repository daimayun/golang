package function

import (
	"io/ioutil"
	"os"
)

// FileGetContents 把整个文件读入一个字符串中[PHP:file_get_contents]
func FileGetContents(filename string) (str string, err error) {
	var data []byte
	data, err = ioutil.ReadFile(filename)
	str = string(data)
	return
}

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

// FileExists 判断文件或目录是否存在[PHP:file_exists]
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// IsDir 判断所给路径是否为文件夹[PHP:is_dir]
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err == nil {
		return f.IsDir()
	}
	return false
}

// IsFile 判断所给路径是否为文件[PHP:is_file]
func IsFile(path string) bool {
	f, err := os.Stat(path)
	if err == nil {
		return !f.IsDir()
	}
	return false
}
