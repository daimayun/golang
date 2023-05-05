package https

import "testing"

func TestGet(t *testing.T) {
	host := "https://www.baidu.com"
	if _, code, actual := Get(host); actual != nil || code != 200 {
		t.Error("Get() error.")
	}
}
