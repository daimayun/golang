package https

import "net/http"

// Get get请求
func Get(url string, headers ...map[string]string) (b []byte, code int, err error) {
	return Request(http.MethodGet, url, nil, headers...)
}
