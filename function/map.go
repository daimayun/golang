package function

// MapKeys 获取MAP中所有的KEY[php:array_keys]
func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
	keys, _ := MapToSlice(m)
	return keys
}

// MapValues 获取MAP中所有的VALUE[php:array_values]
func MapValues[Key comparable, Val any](m map[Key]Val) []Val {
	_, values := MapToSlice(m)
	return values
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

// MapMerge 两个MAP合并[php:array_merge]
func MapMerge[Key comparable, Val any](m1 map[Key]Val, mn ...map[Key]Val) map[Key]Val {
	for _, m := range mn {
		for k, v := range m {
			m1[k] = v
		}
	}
	return m1
}

// MapMergeRecursive 把多个MAP合并为一个MAP[php:array_merge_recursive]
func MapMergeRecursive[Key, Val comparable](m1 map[Key]Val, mn ...map[Key]Val) map[Key][]Val {
	ms := make(map[Key][]Val)
	for k, v := range m1 {
		if InSlice(ms[k], v) == false {
			ms[k] = append(ms[k], v)
		}
	}
	for _, m := range mn {
		for k, v := range m {
			if InSlice(ms[k], v) == false {
				ms[k] = append(ms[k], v)
			}
		}
	}
	return ms
}

// MapSum MAP的VALUE求和[php:array_sum]
func MapSum[Key comparable, Val int | int8 | int32 | int64 | float32 | float64](m map[Key]Val) Val {
	var sum Val = 0
	for _, v := range m {
		sum += v
	}
	return sum
}

// MapIsEqual 判断两个MAP是否相等[长度、Key、Value]
func MapIsEqual[T comparable](m1, m2 map[T]T) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		val, ok := m2[k]
		if !ok || v != val {
			return false
		}
	}
	return true
}

// MapToSlice 将MAP转换成两个切片
func MapToSlice[Key comparable, Val any](m map[Key]Val) (keys []Key, values []Val) {
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return
}
