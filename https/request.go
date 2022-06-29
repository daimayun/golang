package https

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Request 执行请求
func Request(method, url string, body *string, headers ...map[string]string) (b []byte, err error) {
	var (
		req *http.Request
		res *http.Response
	)
	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, strings.NewReader(*body))
	}
	if err != nil {
		return
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			req.Header.Add(key, value)
		}
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}
