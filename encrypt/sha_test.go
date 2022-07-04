package encrypt

import "testing"

func TestSha1(t *testing.T) {
	if actual := Sha1("123456"); actual != "7c4a8d09ca3762af61e59520943dc26494f8941b" {
		t.Error("Sha1() error.")
	}
}

func TestSha256(t *testing.T) {
	if actual := Sha256("123456"); actual != "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92" {
		t.Error("Sha256() error.")
	}
}

func TestSha512(t *testing.T) {
	if actual := Sha512("123456"); actual != "ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413" {
		t.Error("Sha512() error.")
	}
}
