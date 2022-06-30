package https

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// PostJson post_json
func PostJson(url, jsonString string, headers ...map[string]string) (b []byte, err error) {
	if len(headers) > 0 {
		if _, ok := headers[0]["Content-Type"]; !ok {
			headers[0]["Content-Type"] = "application/json"
		}
	} else {
		headers = []map[string]string{{"Content-Type": "application/json"}}
	}
	return Request(http.MethodPost, url, strings.NewReader(jsonString), headers...)
}

// PostForm post_form
func PostForm(url string, data url.Values, headers ...map[string]string) (b []byte, err error) {
	if len(headers) > 0 {
		if _, ok := headers[0]["Content-Type"]; !ok {
			headers[0]["Content-Type"] = "application/x-www-form-urlencoded"
		}
	} else {
		headers = []map[string]string{{"Content-Type": "application/x-www-form-urlencoded"}}
	}
	return Request(http.MethodPost, url, strings.NewReader(data.Encode()), headers...)
}

// PostFormByNotAttachedHeaders post_form[不附带Header信息]
func PostFormByNotAttachedHeaders(url string, data url.Values) (b []byte, err error) {
	var res *http.Response
	res, err = http.PostForm(url, data)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}

// PostFileByByte 以二进制格式上传文件
func PostFileByByte(url, filePath string, headers ...map[string]string) (b []byte, err error) {
	if len(headers) > 0 {
		if _, ok := headers[0]["Content-Type"]; !ok {
			headers[0]["Content-Type"] = "multipart/form-data"
		}
	} else {
		headers = []map[string]string{{"Content-Type": "multipart/form-data"}}
	}
	var fileByte []byte
	fileByte, err = ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	b, err = Request(http.MethodPost, url, bytes.NewReader(fileByte), headers...)
	return
}

// 执行文件操作处理
func execFileHandle(fileKey, filePath string, writer *multipart.Writer) (err error) {
	var (
		file *os.File
		part io.Writer
	)
	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	part, err = writer.CreateFormFile(fileKey, filepath.Base(filePath))
	if err != nil {
		return
	}
	_, err = io.Copy(part, file)
	return
}

// PostFormWithFiles 通过form表单提交上传文件
func PostFormWithFiles(url string, fileData, paramData map[string]string, headers ...map[string]string) (b []byte, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for fileKey, filePath := range fileData {
		err = execFileHandle(fileKey, filePath, writer)
		if err != nil {
			return
		}
	}
	for k, v := range paramData {
		err = writer.WriteField(k, v)
		if err != nil {
			return
		}
	}
	err = writer.Close()
	if err != nil {
		return
	}
	if len(headers) > 0 {
		if _, ok := headers[0]["Content-Type"]; !ok {
			headers[0]["Content-Type"] = writer.FormDataContentType()
		}
	} else {
		headers = []map[string]string{{"Content-Type": writer.FormDataContentType()}}
	}
	return Request(http.MethodPost, url, body, headers...)
}
