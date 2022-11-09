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
	if actual := SliceUnique(testStringSlice.slice); SliceIsEqual(actual, testStringSlice.result, false) == false {
		t.Error("SliceUnique() error.")
	}
	testFloatSlice := struct {
		slice, result []float64
	}{
		[]float64{3.14, 520, 1314, 3.14, 6.28, 520, 1, 100, 3.14},
		[]float64{3.14, 520, 1314, 6.28, 1, 100},
	}
	if actual := SliceUnique(testFloatSlice.slice); SliceIsEqual(actual, testFloatSlice.result, false) == false {
		t.Error("SliceUnique() error.")
	}
	testBoolSlice := struct {
		slice, result []bool
	}{
		[]bool{true, false, false, false, true, false},
		[]bool{true, false},
	}
	if actual := SliceUnique(testBoolSlice.slice); SliceIsEqual(actual, testBoolSlice.result, false) == false {
		t.Error("SliceUnique() error.")
	}
	testIntSlice := struct {
		slice, result []int
	}{
		[]int{1, 3, 5, 3, 2, 6, 7, 9, 0, 4, 6, 8, 3, 2, 1, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	if actual := SliceUnique(testIntSlice.slice); SliceIsEqual(actual, testIntSlice.result, false) == false {
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

func TestIndexOfSlice(t *testing.T) {
	TestIntSlice := struct {
		slice          []int
		target, result int
		ok             bool
	}{
		[]int{1, 3, 5, 100},
		3,
		1,
		true,
	}
	if actualResult, actualOk := IndexOfSlice(TestIntSlice.slice, TestIntSlice.target); actualResult != TestIntSlice.result || actualOk != TestIntSlice.ok {
		t.Error("IndexOfSlice() error.")
	}
	TestIntSlice = struct {
		slice          []int
		target, result int
		ok             bool
	}{
		[]int{1, 3, 5, 100},
		100,
		3,
		true,
	}
	if actualResult, actualOk := IndexOfSlice(TestIntSlice.slice, TestIntSlice.target); actualResult != TestIntSlice.result || actualOk != TestIntSlice.ok {
		t.Error("IndexOfSlice() error.")
	}
	TestStringSlice := struct {
		slice  []string
		target string
		result int
		ok     bool
	}{
		[]string{"a", "b", "c", "d", "100"},
		"c",
		2,
		true,
	}
	if actualResult, actualOk := IndexOfSlice(TestStringSlice.slice, TestStringSlice.target); actualResult != TestStringSlice.result || actualOk != TestStringSlice.ok {
		t.Error("IndexOfSlice() error.")
	}
	TestStringSlice = struct {
		slice  []string
		target string
		result int
		ok     bool
	}{
		[]string{"a", "b", "c", "d", "100"},
		"100",
		4,
		true,
	}
	if actualResult, actualOk := IndexOfSlice(TestStringSlice.slice, TestStringSlice.target); actualResult != TestStringSlice.result || actualOk != TestStringSlice.ok {
		t.Error("IndexOfSlice() error.")
	}
}

func TestFirstAndLastValueOfSlice(t *testing.T) {
	TestSlice := struct {
		slice       []string
		first, last string
	}{
		[]string{"a", "b", "c"},
		"a",
		"c",
	}
	if actualFirst, actualLast := FirstAndLastValueOfSlice(TestSlice.slice); actualFirst != TestSlice.first || actualLast != TestSlice.last {
		t.Error("FirstAndLastValueOfSlice() error.")
	}
	TestSlice = struct {
		slice       []string
		first, last string
	}{
		[]string{"a"},
		"a",
		"a",
	}
	if actualFirst, actualLast := FirstAndLastValueOfSlice(TestSlice.slice); actualFirst != TestSlice.first || actualLast != TestSlice.last {
		t.Error("FirstAndLastValueOfSlice() error.")
	}
}

func TestValueOfSlice(t *testing.T) {
	TestSlice := struct {
		slice []string
		index int
		value string
		ok    bool
	}{
		[]string{"a", "b", "c"},
		2,
		"c",
		true,
	}
	if actualValue, actualOk := ValueOfSlice(TestSlice.slice, TestSlice.index); actualValue != TestSlice.value || actualOk != TestSlice.ok {
		t.Error("ValueOfSlice() error.")
	}
	TestSlice = struct {
		slice []string
		index int
		value string
		ok    bool
	}{
		[]string{"a", "b", "c"},
		0,
		"a",
		true,
	}
	if actualValue, actualOk := ValueOfSlice(TestSlice.slice, TestSlice.index); actualValue != TestSlice.value || actualOk != TestSlice.ok {
		t.Error("ValueOfSlice() error.")
	}
}

func TestSliceDiff(t *testing.T) {
	TestSlice := []struct {
		s1, s2, result []int
	}{
		{
			[]int{1, 3, 5, 7, 9},
			[]int{2, 4, 6, 8, 0},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{1, 0, 5, 6, 7, 8, 9, 3, 4, 5},
			[]int{1, 2, 5, 6, 8, 8, 7, 3, 0, 5},
			[]int{9, 4, 2},
		},
	}
	for _, v := range TestSlice {
		if actual := SliceDiff(v.s1, v.s2); SliceIsEqual(actual, v.result, false) == false {
			t.Error("SliceDiff() error.")
		}
	}
}

func TestSliceIsEqual(t *testing.T) {
	TestIntSlice := []struct {
		s1, s2 []int
		result bool
	}{
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0},
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0},
			true,
		},
		{
			[]int{0, 1, 3, 5, 7, 9, 2, 4, 6, 8, 0},
			[]int{0, 1, 3, 5, 7, 9, 2, 4, 6, 8, 0},
			true,
		},
		{
			[]int{9, 2, 4, 6, 8, 0},
			[]int{9, 2, 4, 6, 8, 0},
			true,
		},
		{
			[]int{6, 8, 1},
			[]int{9, 2, 4, 6, 8, 0},
			false,
		},
		{
			[]int{6, 8, 0, 9, 2, 3},
			[]int{9, 2, 4, 6, 8, 0},
			false,
		},
	}
	for _, v := range TestIntSlice {
		if actual := SliceIsEqual(v.s1, v.s2); actual != v.result {
			t.Error("SliceIsEqual() error.")
		}
		if actual := SliceIsEqual(v.s1, v.s2, false); actual != v.result {
			t.Error("SliceIsEqual() error.")
		}
	}
	if actual := SliceIsEqual([]int{1, 3, 5, 7, 9}, []int{3, 9, 1, 7, 5}); actual == true {
		t.Error("SliceIsEqual() error.")
	}
	if actual := SliceIsEqual([]int{1, 3, 5, 7, 9}, []int{3, 9, 1, 7, 5}, false); actual == false {
		t.Error("SliceIsEqual() error.")
	}
	if actual := SliceIsEqual([]string{"a", "b", "c", "d"}, []string{"d", "b", "c", "a"}); actual == true {
		t.Error("SliceIsEqual() error.")
	}
	if actual := SliceIsEqual([]string{"a", "b", "c", "d"}, []string{"d", "b", "c", "a"}, false); actual == false {
		t.Error("SliceIsEqual() error.")
	}
}

func TestSliceIntersect(t *testing.T) {
	TestSlice := []struct {
		s1, s2, result []int
	}{
		{
			[]int{1, 3, 5, 7, 9},
			[]int{2, 4, 6, 8, 0},
			[]int{},
		},
		{
			[]int{1, 0, 5, 6, 7, 8, 9, 3, 4, 5},
			[]int{1, 2, 5, 6, 8, 8, 7, 3, 0, 5},
			[]int{0, 1, 3, 5, 6, 7, 8},
		},
		{
			[]int{1, 3, 5},
			[]int{6, 1},
			[]int{1},
		},
	}
	for _, v := range TestSlice {
		if actual := SliceIntersect(v.s1, v.s2); SliceIsEqual(actual, v.result, false) == false {
			t.Error("SliceIntersect() error.")
		}
	}
}

func TestSliceMerge(t *testing.T) {
	if actual := SliceMerge([]int{1, 2, 3}); SliceIsEqual(actual, []int{1, 2, 3}, false) == false {
		t.Error("SliceMerge() error.")
	}
	if actual := SliceMerge([]int{1, 2, 3}, []int{4, 5, 6}); SliceIsEqual(actual, []int{1, 2, 3, 4, 5, 6}, false) == false {
		t.Error("SliceMerge() error.")
	}
	if actual := SliceMerge([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}); SliceIsEqual(actual, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, false) == false {
		t.Error("SliceMerge() error.")
	}
	if actual := SliceMerge([]int{1, 2, 3, 4}, []int{4, 5, 6, 7}, []int{7, 8, 9, 1}); SliceIsEqual(actual, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 4, 7}, false) == false {
		t.Error("SliceMerge() error.")
	}
}

func TestSliceShuffle(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	SliceShuffle(&slice)
	if SliceIsEqual(slice, []int{1, 2, 3, 4}, false) == false {
		t.Error("SliceShuffle() error.")
	}
}

func TestSliceReverse(t *testing.T) {
	slice := []string{"e", "d", "c", "b", "a"}
	newSlice := SliceReverse(slice)
	if len(slice) != len(newSlice) {
		t.Error("SliceReverse() error.")
	}
	if newSlice[0] != "a" {
		t.Error("SliceReverse() error.")
	}
	if newSlice[1] != "b" {
		t.Error("SliceReverse() error.")
	}
	if newSlice[2] != "c" {
		t.Error("SliceReverse() error.")
	}
	if newSlice[3] != "d" {
		t.Error("SliceReverse() error.")
	}
	if newSlice[4] != "e" {
		t.Error("SliceReverse() error.")
	}
}
