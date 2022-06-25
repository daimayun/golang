package function

import (
	"testing"
)

func TestInSlice(t *testing.T) {
	testInt := struct {
		slice  []int
		target int
		result bool
	}{
		[]int{1, 3, 5, 7, 9},
		1,
		true,
	}
	if actual := InSlice(testInt.slice, testInt.target); actual != testInt.result {
		t.Error("InSlice() error.")
	}
	testFloat := struct {
		slice  []float64
		target float64
		result bool
	}{
		[]float64{1.1, 3.3, 5.5, 7.7, 9.9},
		9.9,
		true,
	}
	if actual := InSlice(testFloat.slice, testFloat.target); actual != testFloat.result {
		t.Error("InSlice() error.")
	}
	testString := struct {
		slice  []string
		target string
		result bool
	}{
		[]string{"a", "b", "c", "d"},
		"a",
		true,
	}
	if actual := InSlice(testString.slice, testString.target); actual != testString.result {
		t.Error("InSlice() error.")
	}
	testBool := struct {
		slice          []bool
		target, result bool
	}{
		[]bool{true, false, false, true, false},
		true,
		true,
	}
	if actual := InSlice(testBool.slice, testBool.target); actual != testBool.result {
		t.Error("InSlice() error.")
	}
}

func TestSliceUnique(t *testing.T) {
	testString := struct {
		slice, result []string
	}{
		[]string{"a", "b", "a", "d", "b", "c", "d", "e", "f", "f"},
		[]string{"a", "b", "c", "d", "e", "f"},
	}
	if actual := SliceUnique(testString.slice); len(actual) != len(testString.result) {
		t.Error("SliceUnique() error.")
	}
	testFloat := struct {
		slice, result []float64
	}{
		[]float64{3.14, 520, 1314, 3.14, 6.28, 520, 1, 100, 3.14},
		[]float64{3.14, 520, 1314, 6.28, 1, 100},
	}
	if actual := SliceUnique(testFloat.slice); len(actual) != len(testFloat.result) {
		t.Error("SliceUnique() error.")
	}
	testBool := struct {
		slice, result []bool
	}{
		[]bool{true, false, false, false, true, false},
		[]bool{true, false},
	}
	if actual := SliceUnique(testBool.slice); len(actual) != len(testBool.result) {
		t.Error("SliceUnique() error.")
	}
}
