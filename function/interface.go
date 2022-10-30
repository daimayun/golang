package function

import (
	"reflect"
)

// GetValueByInterface 获取interface类型的非指针类型的值
func GetValueByInterface(i interface{}) interface{} {
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		return reflect.ValueOf(i).Elem().Interface()
	}
	return i
}
