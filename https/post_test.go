package https

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestPostJson(t *testing.T) {
	host := "https://www.baidu.com"
	jsonByte, _ := json.Marshal(struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Sex  bool   `json:"sex"`
	}{
		Name: "zs",
		Age:  18,
		Sex:  false,
	})
	if _, code, actual := PostJson(host, string(jsonByte)); actual != nil || code != 200 {
		t.Error("PostJson() error.")
	}
}

func TestPostForm(t *testing.T) {
	host := "https://www.baidu.com"
	var data url.Values = map[string][]string{"name": {"zs"}, "age": {"18"}, "sex": {"男"}}
	if _, code, actual := PostForm(host, data); actual != nil || code != 200 {
		t.Error("PostForm() error.")
	}
}

func TestPostFormByNotAttachedHeaders(t *testing.T) {
	host := "https://www.baidu.com"
	var data url.Values = map[string][]string{"name": {"zs"}, "age": {"18"}, "sex": {"男"}}
	if _, actual := PostFormByNotAttachedHeaders(host, data); actual != nil {
		t.Error("PostFormByNotAttachedHeaders() error.")
	}
}

func TestPostFormWithFiles(t *testing.T) {
	host := "https://www.baidu.com"
	if _, code, actual := PostFormWithFiles(host, map[string]string{"file": "../function/test.txt"}, map[string]string{"hello": "world"}); actual != nil || code != 200 {
		t.Error("PostFormWithFiles() error.")
	}
}
