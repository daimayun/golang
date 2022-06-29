package function

import (
	"testing"
)

func TestMapKeys(t *testing.T) {
	testMap := struct {
		m    map[string]string
		keys []string
	}{
		map[string]string{"a": "one", "b": "two", "c": "three"},
		[]string{"a", "b", "c"},
	}
	if actual := MapKeys(testMap.m); SliceIsEqual(actual, testMap.keys, false) == false {
		t.Error("MapKeys() error.")
	}
	testIntMap := struct {
		m    map[string]int
		keys []string
	}{
		map[string]int{"aa": 1, "bb": 2, "cc": 3},
		[]string{"aa", "bb", "cc"},
	}
	if actual := MapKeys(testIntMap.m); SliceIsEqual(actual, testIntMap.keys, false) == false {
		t.Error("MapKeys() error.")
	}
}

func TestMapValues(t *testing.T) {
	testMap := struct {
		m      map[string]string
		values []string
	}{
		map[string]string{"a": "one", "b": "two", "c": "three"},
		[]string{"one", "two", "three"},
	}
	if actual := MapValues(testMap.m); SliceIsEqual(actual, testMap.values, false) == false {
		t.Error("MapValues() error.")
	}
	testIntMap := struct {
		m      map[string]int
		values []int
	}{
		map[string]int{"aa": 1, "bb": 2, "cc": 3},
		[]int{1, 2, 3},
	}
	if actual := MapValues(testIntMap.m); SliceIsEqual(actual, testIntMap.values, false) == false {
		t.Error("MapValues() error.")
	}
}

func TestMapKeyExists(t *testing.T) {
	testMap := struct {
		m          map[string]string
		key, value string
		ok         bool
	}{
		map[string]string{"a": "one", "b": "two", "c": "three"},
		"a",
		"one",
		true,
	}
	if actualValue, actualOk := MapKeyExists(testMap.m, testMap.key); actualValue != testMap.value || actualOk != testMap.ok {
		t.Error("MapKeyExists() error.")
	}
	testIntMap := struct {
		m     map[string]int
		key   string
		value int
		ok    bool
	}{
		map[string]int{"aa": 1, "bb": 2, "cc": 3},
		"cc",
		3,
		true,
	}
	if actualValue, actualOk := MapKeyExists(testIntMap.m, testIntMap.key); actualValue != testIntMap.value || actualOk != testIntMap.ok {
		t.Error("MapKeyExists() error.")
	}
}

func TestMapValueExists(t *testing.T) {
	testMap := struct {
		m          map[string]string
		key, value string
		ok         bool
	}{
		map[string]string{"a": "one", "b": "two", "c": "three"},
		"a",
		"one",
		true,
	}
	if actualKey, actualOk := MapValueExists(testMap.m, testMap.value); actualKey != testMap.key || actualOk != testMap.ok {
		t.Error("MapValueExists() error.")
	}
	testIntMap := struct {
		m     map[string]int
		key   string
		value int
		ok    bool
	}{
		map[string]int{"aa": 1, "bb": 2, "cc": 3},
		"cc",
		3,
		true,
	}
	if actualKey, actualOk := MapValueExists(testIntMap.m, testIntMap.value); actualKey != testIntMap.key || actualOk != testIntMap.ok {
		t.Error("MapValueExists() error.")
	}
}

func TestMapMerge(t *testing.T) {
	testMap := struct {
		m1, m2, result map[string]string
	}{
		map[string]string{"a": "one", "c": "three", "b": "2"},
		map[string]string{"b": "two", "d": "four", "e": "five"},
		map[string]string{"a": "one", "b": "two", "c": "three", "d": "four", "e": "five"},
	}
	if actual := MapMerge(testMap.m1, testMap.m2); MapIsEqual(actual, testMap.result) == false {
		t.Error("MapMerge() error.")
	}
	testMaps := struct {
		m1, m2, m3, m4, result map[string]string
	}{
		map[string]string{"a": "one", "c": "three", "b": "two", "g": "seven"},
		map[string]string{"b": "two", "d": "four", "h": "eight", "e": "five"},
		map[string]string{"j": "ten", "b": "two", "f": "six", "d": "four", "e": "five"},
		map[string]string{"b": "2", "d": "four", "e": "5", "i": "nine"},
		map[string]string{"a": "one", "b": "2", "c": "three", "d": "four", "e": "5", "f": "six", "g": "seven", "h": "eight", "i": "nine", "j": "ten"},
	}
	if actual := MapMerge(testMaps.m1, testMaps.m2, testMaps.m3, testMaps.m4); MapIsEqual(actual, testMaps.result) == false {
		t.Error("MapMerge() error.")
	}
}

func TestMapSum(t *testing.T) {
	testMap := struct {
		m   map[string]int
		sum int
	}{
		map[string]int{"one": 1, "two": 3, "three": 5, "four": 10},
		19,
	}
	if actual := MapSum(testMap.m); actual != testMap.sum {
		t.Error("MapSum() error.")
	}
	testFloatMap := struct {
		m   map[string]float64
		sum float64
	}{
		map[string]float64{"one": 1.8, "two": 3, "three": 5.1, "four": 10},
		19.9,
	}
	if actual := MapSum(testFloatMap.m); actual != testFloatMap.sum {
		t.Error("MapSum() error.")
	}
}

func TestMapMergeRecursive(t *testing.T) {
	testMap := struct {
		m1, m2 map[string]string
		result map[string][]string
	}{
		map[string]string{"a": "one", "c": "three", "b": "2"},
		map[string]string{"b": "two", "d": "four", "e": "five"},
		map[string][]string{"a": {"one"}, "b": {"two", "2"}, "c": {"three"}, "d": {"four"}, "e": {"five"}},
	}
	actual := MapMergeRecursive(testMap.m1, testMap.m2)
	if len(actual) != len(testMap.result) {
		t.Error("MapMergeRecursive() error.")
	}
	for k, v := range actual {
		if SliceIsEqual(v, testMap.result[k], false) == false {
			t.Error("MapMergeRecursive() error.")
		}
	}
	testMaps := struct {
		m1, m2, m3, m4 map[string]string
		result         map[string][]string
	}{
		map[string]string{"a": "one", "c": "three", "b": "two", "g": "seven"},
		map[string]string{"b": "two", "d": "four", "h": "eight", "e": "five"},
		map[string]string{"j": "ten", "b": "二", "f": "six", "d": "four", "e": "five"},
		map[string]string{"b": "2", "d": "four", "e": "5", "i": "nine"},
		map[string][]string{"a": {"one"}, "b": {"2", "two", "二"}, "c": {"three"}, "d": {"four"}, "e": {"5", "five"}, "f": {"six"}, "g": {"seven"}, "h": {"eight"}, "i": {"nine"}, "j": {"ten"}},
	}
	actual = MapMergeRecursive(testMaps.m1, testMaps.m2, testMaps.m3, testMaps.m4)
	if len(actual) != len(testMaps.result) {
		t.Error("MapMergeRecursive() error.")
	}
	for k, v := range actual {
		if SliceIsEqual(v, testMaps.result[k], false) == false {
			t.Error("MapMergeRecursive() error.")
		}
	}
}

func TestMapToSlice(t *testing.T) {
	actualKeys, actualValues := MapToSlice(map[string]string{"a": "one", "b": "two", "three": "3", "4": "four"})
	if SliceIsEqual(actualKeys, []string{"a", "b", "three", "4"}, false) == false {
		t.Error("MapToSlice() error.")
	}
	if SliceIsEqual(actualValues, []string{"one", "two", "3", "four"}, false) == false {
		t.Error("MapToSlice() error.")
	}
}
