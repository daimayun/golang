package encrypt

import "testing"

func TestMd5(t *testing.T) {
	if actual := Md5("123456"); actual != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("Md5() error.")
	}
}

func TestMd5S16(t *testing.T) {
	if actual := Md5S16("123456"); actual != "49ba59abbe56e057" {
		t.Error("Md5S16() error.")
	}
}
