package function

import "testing"

func TestMapKeys(t *testing.T) {
	testMap := struct {
		m    map[string]string
		keys []string
	}{
		map[string]string{"a": "one", "b": "two", "c": "three"},
		[]string{"a", "b", "c"},
	}
	if actual := MapKeys(testMap.m); len(actual) != len(testMap.keys) {
		t.Error("MapKeys() error.")
	}
	testIntMap := struct {
		m    map[string]int
		keys []string
	}{
		map[string]int{"aa": 1, "bb": 2, "cc": 3},
		[]string{"aa", "bb", "cc"},
	}
	if actual := MapKeys(testIntMap.m); len(actual) != len(testIntMap.keys) {
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
	if actual := MapValues(testMap.m); len(actual) != len(testMap.values) {
		t.Error("MapValues() error.")
	}
	testIntMap := struct {
		m      map[string]int
		values []int
	}{
		map[string]int{"aa": 1, "bb": 2, "cc": 3},
		[]int{1, 2, 3},
	}
	if actual := MapValues(testIntMap.m); len(actual) != len(testIntMap.values) {
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
