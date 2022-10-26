package times

import (
	"fmt"
	"github.com/daimayun/golang/function"
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

func TestSubDays(t *testing.T) {
	startTime := TimeNow()
	endTime := TimeNow()
	if SubDays(endTime, startTime) != 0 {
		t.Error("SubDays() error.")
	}
	endTime = AfterDayTime(45, startTime)
	if SubDays(endTime, startTime) != 45 {
		t.Error("SubDays() error.")
	}
}

func TestGenerateBetweenDates(t *testing.T) {
	startDate := "2022-08-08"
	endDate := "2022-09-09"
	days := []string{
		"2022-08-08",
		"2022-08-09",
		"2022-08-10",
		"2022-08-11",
		"2022-08-12",
		"2022-08-13",
		"2022-08-14",
		"2022-08-15",
		"2022-08-16",
		"2022-08-17",
		"2022-08-18",
		"2022-08-19",
		"2022-08-20",
		"2022-08-21",
		"2022-08-22",
		"2022-08-23",
		"2022-08-24",
		"2022-08-25",
		"2022-08-26",
		"2022-08-27",
		"2022-08-28",
		"2022-08-29",
		"2022-08-30",
		"2022-08-31",
		"2022-09-01",
		"2022-09-02",
		"2022-09-03",
		"2022-09-04",
		"2022-09-05",
		"2022-09-06",
		"2022-09-07",
		"2022-09-08",
		"2022-09-09",
	}
	timeList, err := GenerateBetweenDates(DateTypeDay, startDate, endDate, TimeLayoutYMD)
	if err != nil {
		t.Error("GenerateBetweenDates() error.")
	}
	var timeDays []string
	for _, v := range timeList {
		timeDays = append(timeDays, v.Format(TimeLayoutYMD))
	}
	fmt.Println(timeDays)
	if function.SliceIsEqual(days, timeDays, false) == false {
		t.Error("GenerateBetweenDates() error.")
	}
	endDate = "2023-09-09"
	months := []string{
		"2022-08",
		"2022-09",
		"2022-10",
		"2022-11",
		"2022-12",
		"2023-01",
		"2023-02",
		"2023-03",
		"2023-04",
		"2023-05",
		"2023-06",
		"2023-07",
		"2023-08",
		"2023-09",
	}
	timeList, err = GenerateBetweenDates(DateTypeMonth, startDate, endDate, TimeLayoutYMD)
	if err != nil {
		t.Error("GenerateBetweenDates() error.")
	}
	var timeMonths []string
	for _, v := range timeList {
		timeMonths = append(timeMonths, v.Format(TimeLayoutYM))
	}
	fmt.Println(timeMonths)
	if function.SliceIsEqual(months, timeMonths, false) == false {
		t.Error("GenerateBetweenDates() error.")
	}
}
