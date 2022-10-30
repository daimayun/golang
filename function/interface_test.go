package function

import (
	"testing"
)

type getValueByInterface struct {
	Name string
	Age  int
	Sex  bool
}

func TestGetValueByInterface(t *testing.T) {
	var val getValueByInterface
	valPtr := &val
	v1 := GetValueByInterface(val)
	if v1.(getValueByInterface) != val {
		t.Error("GetValueByInterface() error.")
	}
	v2 := GetValueByInterface(valPtr)
	if v2.(getValueByInterface) != val {
		t.Error("GetValueByInterface() error.")
	}
}
