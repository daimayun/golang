package function

import "testing"

func TestInSlice(t *testing.T) {
	testInt := struct {
		slice  []int
		target int
		res    bool
	}{
		[]int{1, 3, 5, 7, 9},
		1,
		true,
	}
	if actual := InSlice(testInt.slice, testInt.target); actual != testInt.res {
		t.Error("InSlice() error.")
	}
	testFloat := struct {
		slice  []float64
		target float64
		res    bool
	}{
		[]float64{1.1, 3.3, 5.5, 7.7, 9.9},
		9.9,
		true,
	}
	if actual := InSlice(testFloat.slice, testFloat.target); actual != testFloat.res {
		t.Error("InSlice() error.")
	}
	testString := struct {
		slice  []string
		target string
		res    bool
	}{
		[]string{"a", "b", "c", "d"},
		"a",
		true,
	}
	if actual := InSlice(testString.slice, testString.target); actual != testString.res {
		t.Error("InSlice() error.")
	}
}
