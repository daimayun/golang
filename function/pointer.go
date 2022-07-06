package function

// GetPointer 获取指针地址
func GetPointer[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string | float32 | float64 | bool](i T) *T {
	return &i
}
