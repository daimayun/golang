package function

import "testing"

func TestMax(t *testing.T) {
	var x, y, v int
	x = 100
	y = 101
	v = Max(x, y)
	if v != y {
		t.Error("Max() error.")
	}
}

func TestMin(t *testing.T) {
	var x, y, v float64
	x = 100
	y = 101.64
	v = Min(x, y)
	if v != x {
		t.Error("Min() error.")
	}
}
