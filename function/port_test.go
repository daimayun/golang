package function

import "testing"

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	if err != nil || !(port > 0) {
		t.Error("GetFreePort() error.")
	}
}
