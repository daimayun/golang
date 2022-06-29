package https

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PostJson post_json
func PostJson(url string, jsonStrReader io.Reader, headers ...map[string]string) (b []byte, err error) {
	if len(headers) > 0 {
		if _, ok := headers[0]["Content-Type"]; !ok {
			headers[0]["Content-Type"] = "application/json"
		}
	} else {
		headers = []map[string]string{{"Content-Type": "application/json"}}
	}
	return Request(http.MethodPost, url, jsonStrReader, headers...)
}

// PostForm post_form
func PostForm(url string, data url.Values) (b []byte, err error) {
	var res *http.Response
	res, err = http.PostForm(url, data)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}
