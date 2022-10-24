package function

import "testing"

func TestReturnMaxAndMinNum(t *testing.T) {
	var max, min float64
	var vMax, vMin float64
	max = 9
	min = 1
	vMax, vMin = ReturnMaxAndMinNum(1)
	if max != vMax || min != vMin {
		t.Error("TestReturnMaxAndMinNum() error.")
	}
	max = 99
	min = 10
	vMax, vMin = ReturnMaxAndMinNum(2)
	if max != vMax || min != vMin {
		t.Error("TestReturnMaxAndMinNum() error.")
	}
	max = 99999
	min = 10000
	vMax, vMin = ReturnMaxAndMinNum(5)
	if max != vMax || min != vMin {
		t.Error("TestReturnMaxAndMinNum() error.")
	}
	max = 999999999
	min = 100000000
	vMax, vMin = ReturnMaxAndMinNum(9)
	if max != vMax || min != vMin {
		t.Error("TestReturnMaxAndMinNum() error.")
	}
}
