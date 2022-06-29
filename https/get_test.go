package https

import "testing"

func TestGet(t *testing.T) {
	host := "https://www.baidu.com"
	if _, actual := Get(host); actual != nil {
		t.Error("Get() error.")
	}
}
