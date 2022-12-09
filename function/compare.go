package function

import "github.com/daimayun/golang/base"

// Max 返回最大数
func Max[T base.Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Min 返回最小数
func Min[T base.Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}
