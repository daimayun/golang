package function

import "math"

// ReturnMaxAndMinNum 返回最大和最小的X位数
func ReturnMaxAndMinNum(num int) (max, min float64) {
	if num < 1 {
		return
	}
	max = math.Pow10(num) - 1
	if num == 1 {
		min = 1
		return
	}
	min = math.Pow10(num - 1)
	return
}
