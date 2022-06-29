package https

import (
	"net/http"
	"net/url"
	"testing"
)

func TestRequest(t *testing.T) {
	host := "https://www.baidu.com"
	var data url.Values = map[string][]string{"name": {"zs"}, "age": {"18"}, "sex": {"男"}}
	str := data.Encode()
	if _, actual := Request(http.MethodPost, host, &str, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}); actual != nil {
		t.Error("Request() error.")
	}
}
