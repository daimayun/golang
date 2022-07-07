package function

// GetPointer 获取变量的指针地址
func GetPointer[T any](v T) *T {
	return &v
}

// GetPointerValue 根据指针地址获取变量的值
func GetPointerValue[T any](vp *T) (v T) {
	if vp == nil {
		return
	}
	return *vp
}
