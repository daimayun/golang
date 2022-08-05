package file

import "io/ioutil"

// FileGetContents 把整个文件读入一个字符串中[PHP:file_get_contents]
func FileGetContents(filename string) (str string, err error) {
	var data []byte
	data, err = FileGetContentsToByte(filename)
	if err != nil {
		return
	}
	str = string(data)
	return
}

// FileGetContentsToByte 把整个文件读入byte中[PHP:file_get_contents]
func FileGetContentsToByte(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
