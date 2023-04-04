package sort

import (
	"github.com/daimayun/golang/base"
	systemSort "sort"
)

// BubbleSort 冒泡排序[默认升序]
func BubbleSort[T base.Number](arr []T, desc ...bool) []T {
	asc := true
	if len(desc) > 0 {
		if desc[0] {
			asc = false
		}
	}
	if asc {
		systemSort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	} else {
		systemSort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
	}
	return arr
}
