package https

import (
	"io/ioutil"
	"net/http"
	"net/url"
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
	return Request(http.MethodPost, url, &jsonString, headers...)
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
	str := data.Encode()
	return Request(http.MethodPost, url, &str, headers...)
}

// PostFormByNotAttachedHeaders post_form 不附带Header信息
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
