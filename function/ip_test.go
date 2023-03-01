package function

import "testing"

func TestGetLocalIp(t *testing.T) {
	_, err := GetLocalIp()
	if err != nil {
		t.Error("TestGetLocalIp() error.")
	}
}
