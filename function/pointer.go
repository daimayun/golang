package function

// GetPointer 获取变量的指针地址
func GetPointer[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string | float32 | float64 | bool](v T) *T {
	return &v
}

// GetPointerValue 根据指针地址获取变量的值
func GetPointerValue[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string | float32 | float64 | bool](vp *T) (v T) {
	if vp == nil {
		return
	}
	return *vp
}
