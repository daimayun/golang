package times

import "testing"

func TestFebruaryIsLeapYear(t *testing.T) {
	year := 1990
	if FebruaryIsLeapYear(year) {
		t.Error("FebruaryIsLeapYear() error.")
	}
	year = 1996
	if FebruaryIsLeapYear(year) == false {
		t.Error("FebruaryIsLeapYear() error.")
	}
	year = 2000
	if FebruaryIsLeapYear(year) == false {
		t.Error("FebruaryIsLeapYear() error.")
	}
	year = 2008
	if FebruaryIsLeapYear(year) == false {
		t.Error("FebruaryIsLeapYear() error.")
	}
	year = 2010
	if FebruaryIsLeapYear(year) {
		t.Error("FebruaryIsLeapYear() error.")
	}
	year = 2020
	if FebruaryIsLeapYear(year) == false {
		t.Error("FebruaryIsLeapYear() error.")
	}
}
