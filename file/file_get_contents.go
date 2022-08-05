package file

import "io/ioutil"

// GetContents 把整个文件读入一个字符串中[PHP:file_get_contents]
func GetContents(filename string) (str string, err error) {
	var data []byte
	data, err = GetContentsToByte(filename)
	if err != nil {
		return
	}
	str = string(data)
	return
}

// GetContentsToByte 把整个文件读入byte中[PHP:file_get_contents]
func GetContentsToByte(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
