package https

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
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

// PostFileByForm 通过form表单提交上传文件
func PostFileByForm(url string) (b []byte, err error) {
	return
}
