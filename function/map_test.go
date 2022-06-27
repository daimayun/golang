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
}
