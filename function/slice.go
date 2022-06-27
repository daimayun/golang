package function

// InSlice 判断元素是否在切片内[php:in_array]
func InSlice[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// SliceUnique 切片去重[php:array_unique]
func SliceUnique[T comparable](slice []T) []T {
	if len(slice) <= 1 {
		return slice
	}
	var (
		newSlice = make([]T, 0)
		m        = make(map[T]interface{})
	)
	for _, v := range slice {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = nil
		newSlice = append(newSlice, v)
	}
	return newSlice
}

// SliceSum 切片元素求和[php:array_sum]
func SliceSum[T uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64](slice []T) T {
	var sum T = 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// IndexOfSlice 返回元素所在切片中的下标
func IndexOfSlice[T comparable](slice []T, target T) (n int, ok bool) {
	for k, v := range slice {
		if v == target {
			return k, true
		}
	}
	return
}

// FirstAndLastValueBySlice 返回切片中第一个和最后元素
func FirstAndLastValueBySlice[T any](slice []T) (T, T) {
	return slice[0], slice[len(slice)-1]
}
