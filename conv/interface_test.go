package conv

import (
	"testing"
)

func TestInterfaceToString(t *testing.T) {
	var str = "10"
	data := struct {
		s    string
		i    int
		i8   int8
		i16  int16
		i32  int32
		i64  int64
		ui   uint
		ui8  uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
		b    []byte
	}{
		"10",
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		[]byte("10"),
	}
	var actualStr string
	var actualErr error
	if actualStr, actualErr = InterfaceToString(data.s); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.i); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.i8); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.i16); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.i32); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.i64); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.ui); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.ui8); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.ui16); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.ui32); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.ui64); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	if actualStr, actualErr = InterfaceToString(data.b); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	str = `{"name":"zs","age":18,"sex":true,"province":"河南"}`
	user := struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Sex      bool   `json:"sex"`
		Province string `json:"province"`
	}{
		"zs",
		18,
		true,
		"河南",
	}
	if actualStr, actualErr = InterfaceToString(user); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	str = `{"province":"河南"}`
	um := map[string]interface{}{"province": "河南"}
	if actualStr, actualErr = InterfaceToString(um); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
	str = `["1","2","3"]`
	slice := []string{"1", "2", "3"}
	if actualStr, actualErr = InterfaceToString(slice); actualStr != str || actualErr != nil {
		t.Error("InterfaceToString() error.")
	}
}
