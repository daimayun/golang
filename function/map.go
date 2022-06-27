package function

// MapKeys 获取MAP中所有的KEY[php:array_keys]
func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
	keys := make([]Key, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValues 获取MAP中所有的VALUE[php:array_values]
func MapValues[Key comparable, Val any](m map[Key]Val) []Val {
	keys := make([]Val, 0)
	for _, v := range m {
		keys = append(keys, v)
	}
	return keys
}

// MapKeyExists 判断MAP中是否含有某个KEY[php:array_key_exists]
func MapKeyExists[Key comparable, Val any](m map[Key]Val, key Key) (Val, bool) {
	v, ok := m[key]
	return v, ok
}

// MapValueExists 判断MAP中是否含有某个VALUE[php:array_search|array_value_exists]
func MapValueExists[Key, Val comparable](m map[Key]Val, value Val) (key Key, ok bool) {
	for k, v := range m {
		if v == value {
			return k, true
		}
	}
	return
}
