package function

import "reflect"

// GetStructTag 获取结构体某个字段的标签值
func GetStructTag(s any, field, key string) (val string, ok bool) {
	var structField reflect.StructField
	structField, ok = reflect.TypeOf(s).FieldByName(field)
	if ok {
		val, ok = structField.Tag.Lookup(key)
	}
	return
}
