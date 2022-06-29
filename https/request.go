package https

import (
	"io"
	"io/ioutil"
	"net/http"
)

func Request(method, url string, jsonStrReader io.Reader, headers ...map[string]string) (b []byte, err error) {
	var (
		req *http.Request
		res *http.Response
	)
	req, err = http.NewRequest(method, url, jsonStrReader)
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
