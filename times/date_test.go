package times

import (
	"testing"
)

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

func TestMonthDays(t *testing.T) {
	ts, err := StringToTime("2020-02-15", TimeLayoutYMD)
	if err != nil {
		t.Error("StringToTime() error.")
	}
	if MonthDays(ts) != 29 {
		t.Error("MonthDays() error.")
	}
	ts, err = StringToTime("2020-01-15", TimeLayoutYMD)
	if err != nil {
		t.Error("StringToTime() error.")
	}
	if MonthDays(ts) != 31 {
		t.Error("MonthDays() error.")
	}
	ts, err = StringToTime("2020-04-15", TimeLayoutYMD)
	if err != nil {
		t.Error("StringToTime() error.")
	}
	if MonthDays(ts) != 30 {
		t.Error("MonthDays() error.")
	}
}

func TestSubMonths(t *testing.T) {
	startTime := TimeNow()
	endTime := TimeNow()
	if SubMonths(startTime, endTime) != 0 {
		t.Error("SubMonths() error.")
	}
	if subMonths(startTime, endTime) != 0 {
		t.Error("subMonths() error.")
	}
	endTime = AfterMonthTime(5, startTime)
	if SubMonths(endTime, startTime) != 5 {
		t.Error("SubMonths() error.")
	}
	if subMonths(endTime, startTime) != 5 {
		t.Error("subMonths() error.")
	}
}
