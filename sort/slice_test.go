package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	var sInt, sIntResult []int
	sInt = []int{1, 8, 4, 6, 3, 9, 5, 0, 7, 2}
	sIntResult = BubbleSort(sInt)
	for i := 0; i < 10; i++ {
		if sIntResult[i] != i {
			t.Error("BubbleSort() error.")
		}
	}
	var sFloat, sFloatResult []float64
	sFloat = []float64{1, 8, 4, 6, 3, 9, 5, 0, 7, 2}
	sFloatResult = BubbleSort(sFloat)
	var i float64
	var v int
	for i = 0; i < 10; i++ {
		if sFloatResult[v] != i {
			t.Error("BubbleSort() error.")
		}
		v++
	}
}
