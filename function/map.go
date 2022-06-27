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

// MapMerge 两个MAP合并[php:array_merge]
func MapMerge[Key comparable, Val any](map1, map2 map[Key]Val) map[Key]Val {
	m := make(map[Key]Val)
	for i, v := range map1 {
		for j, w := range map2 {
			if i == j {
				m[i] = w
			} else {
				if _, ok := m[i]; !ok {
					m[i] = v
				}
				if _, ok := m[j]; !ok {
					m[j] = w
				}
			}
		}
	}
	return m
}

// MapSum MAP的VALUE求和[php:array_sum]
func MapSum[Key comparable, Val int | int8 | int32 | int64 | float32 | float64](m map[Key]Val) Val {
	var sum Val = 0
	for _, v := range m {
		sum += v
	}
	return sum
}

// MapDiff 两个MAP之间的差集[php:array_diff]
func MapDiff() {
	//
}

// MapIsEqual 判断两个MAP是否相等[长度、KEY、Value]
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
