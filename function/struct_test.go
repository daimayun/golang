package function

import "testing"

func TestGetStructTag(t *testing.T) {
	type tagExample struct {
		Name string `json:"name" form:"name"`
		Age  int64  `json:"age" form:"age"`
	}
	actualVal, actualOk := GetStructTag(tagExample{}, "Name", "json")
	if !actualOk || actualVal != "name" {
		t.Error("GetStructTag() error.")
	}
}
