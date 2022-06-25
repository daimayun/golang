package function

import (
	"testing"
)

func TestInSlice(t *testing.T) {
	testIntSlice := struct {
		slice  []int
		target int
		result bool
	}{
		[]int{1, 3, 5, 7, 9},
		1,
		true,
	}
	if actual := InSlice(testIntSlice.slice, testIntSlice.target); actual != testIntSlice.result {
		t.Error("InSlice() error.")
	}
	testFloatSlice := struct {
		slice  []float64
		target float64
		result bool
	}{
		[]float64{1.1, 3.3, 5.5, 7.7, 9.9},
		9.9,
		true,
	}
	if actual := InSlice(testFloatSlice.slice, testFloatSlice.target); actual != testFloatSlice.result {
		t.Error("InSlice() error.")
	}
	testStringSlice := struct {
		slice  []string
		target string
		result bool
	}{
		[]string{"a", "b", "c", "d"},
		"a",
		true,
	}
	if actual := InSlice(testStringSlice.slice, testStringSlice.target); actual != testStringSlice.result {
		t.Error("InSlice() error.")
	}
	testBoolSlice := struct {
		slice          []bool
		target, result bool
	}{
		[]bool{true, false, false, true, false},
		true,
		true,
	}
	if actual := InSlice(testBoolSlice.slice, testBoolSlice.target); actual != testBoolSlice.result {
		t.Error("InSlice() error.")
	}
}

func TestSliceUnique(t *testing.T) {
	testStringSlice := struct {
		slice, result []string
	}{
		[]string{"a", "b", "a", "d", "b", "c", "d", "e", "f", "f"},
		[]string{"a", "b", "c", "d", "e", "f"},
	}
	if actual := SliceUnique(testStringSlice.slice); len(actual) != len(testStringSlice.result) {
		t.Error("SliceUnique() error.")
	}
	testFloatSlice := struct {
		slice, result []float64
	}{
		[]float64{3.14, 520, 1314, 3.14, 6.28, 520, 1, 100, 3.14},
		[]float64{3.14, 520, 1314, 6.28, 1, 100},
	}
	if actual := SliceUnique(testFloatSlice.slice); len(actual) != len(testFloatSlice.result) {
		t.Error("SliceUnique() error.")
	}
	testBoolSlice := struct {
		slice, result []bool
	}{
		[]bool{true, false, false, false, true, false},
		[]bool{true, false},
	}
	if actual := SliceUnique(testBoolSlice.slice); len(actual) != len(testBoolSlice.result) {
		t.Error("SliceUnique() error.")
	}
	testIntSlice := struct {
		slice, result []int
	}{
		[]int{1, 3, 5, 3, 2, 6, 7, 9, 0, 4, 6, 8, 3, 2, 1, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	if actual := SliceUnique(testIntSlice.slice); len(actual) != len(testIntSlice.result) {
		t.Error("SliceUnique() error.")
	}
}

func TestSliceSum(t *testing.T) {
	TestIntSlice := struct {
		slice  []int
		result int
	}{
		[]int{1, 3, 5, 100},
		109,
	}
	if actual := SliceSum(TestIntSlice.slice); actual != TestIntSlice.result {
		t.Error("SliceSum() error.")
	}
	TestIntSlice = struct {
		slice  []int
		result int
	}{
		[]int{1, -3, 5, 100},
		103,
	}
	if actual := SliceSum(TestIntSlice.slice); actual != TestIntSlice.result {
		t.Error("SliceSum() error.")
	}
	TestIntFloat := struct {
		slice  []float64
		result float64
	}{
		[]float64{1, 3, 5, 100.99},
		109.99,
	}
	if actual := SliceSum(TestIntFloat.slice); actual != TestIntFloat.result {
		t.Error("SliceSum() error.")
	}
	TestIntFloat = struct {
		slice  []float64
		result float64
	}{
		[]float64{-1, 3, 5, 100.99},
		107.99,
	}
	if actual := SliceSum(TestIntFloat.slice); actual != TestIntFloat.result {
		t.Error("SliceSum() error.")
	}
}
