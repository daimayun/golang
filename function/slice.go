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

// SliceDiff 两个切片的差集[php:array_diff]
func SliceDiff[T comparable](s1, s2 []T) (s3 []T) {
	m1 := make(map[T]struct{})
	m2 := make(map[T]struct{})
	for _, v := range s1 {
		m1[v] = struct{}{}
	}
	for _, v := range s2 {
		m2[v] = struct{}{}
		if _, ok := m1[v]; !ok {
			s3 = append(s3, v)
		}
	}
	for _, v := range s1 {
		if _, ok := m2[v]; !ok {
			s3 = append(s3, v)
		}
	}
	return
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

// ValueOfSlice 返回下标对应切片中的元素
func ValueOfSlice[T any](slice []T, index int) (val T, ok bool) {
	if len(slice) >= index {
		return slice[index], true
	}
	return
}

// FirstAndLastValueOfSlice 返回切片中第一个和最后元素
func FirstAndLastValueOfSlice[T any](slice []T) (T, T) {
	return slice[0], slice[len(slice)-1]
}

// SliceIsEqual 判断两个切片是否相等[长度、元素][forceChecks:是否强制校验切片元素的顺序和容量大小,默认强制校验]
func SliceIsEqual[T comparable](s1, s2 []T, forceChecks ...bool) bool {
	check := true
	if len(forceChecks) > 0 {
		check = forceChecks[0]
	}
	if len(s1) != len(s2) {
		return false
	}
	if check {
		if cap(s1) != cap(s2) {
			return false
		}
		for k, v := range s1 {
			if s2[k] != v {
				return false
			}
		}
	} else {
		for _, v := range s1 {
			if InSlice(s2, v) == false {
				return false
			}
		}
	}
	return true
}
