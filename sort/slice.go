package sort

import "github.com/daimayun/golang/base"

// BubbleSort 冒泡排序[默认升序]
func BubbleSort[T base.Number](arr []T, desc ...bool) []T {
	asc := true
	if len(desc) > 0 {
		if desc[0] {
			asc = false
		}
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if asc {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			} else {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
	return arr
}
