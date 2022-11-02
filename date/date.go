package date

import (
	"errors"
	"time"
)

// FebruaryIsLeapYear 二月是否是闰年
func FebruaryIsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

// MonthDays 该月有多少天
func MonthDays(ts ...time.Time) int8 {
	t := TimeNow()
	if len(ts) > 0 {
		t = ts[0]
	}
	month := int8(t.Month())
	type dayType map[int8]struct{}
	day31 := dayType{
		1:  {},
		3:  {},
		5:  {},
		7:  {},
		8:  {},
		10: {},
		12: {},
	}
	if _, ok := day31[month]; ok {
		return 31
	}
	day30 := dayType{
		4:  {},
		6:  {},
		9:  {},
		11: {},
	}
	if _, ok := day30[month]; ok {
		return 30
	}
	if FebruaryIsLeapYear(t.Year()) {
		return 29
	}
	return 28
}

// SubMonths 两个时间相隔多少个月[t1减t2][优先使用]
func SubMonths(t1, t2 time.Time) (months int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	chaY := y1 - y2
	chaM := 12 - m2 + m1
	months = chaM + ((chaY - 1) * 12)
	return
}

// subMonths 两个时间相隔多少个月[t1减t2]
func subMonths(t1, t2 time.Time) (months int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()
	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	months = yearInterval*12 + monthInterval
	return
}

// SubDays 两个时间相隔多少天[t1减t2]
func SubDays(t1, t2 time.Time) (days int) {
	swap := false
	if t1.Unix() < t2.Unix() {
		t_ := t1
		t1 = t2
		t2 = t_
		swap = true
	}
	days = int(t1.Sub(t2).Hours() / 24)
	// 计算在被24整除外的时间是否存在跨自然日
	if int(t1.Sub(t2).Milliseconds())%86400000 > int(86400000-t2.Unix()%86400000) {
		days += 1
	}
	if swap {
		days = -days
	}
	return
}

// GenerateBetweenDays 根据开始日期和结束日期生成该时间段内所有日期[两个日期内所有天数]
func GenerateBetweenDays(startTime, endTime time.Time) (days []time.Time) {
	if endTime.Before(startTime) {
		startTime, endTime = endTime, startTime
	}
	startDay := startTime.Format(TimeLayoutYMD)
	endDay := endTime.Format(TimeLayoutYMD)
	days = append(days, startTime)
	if startDay == endDay {
		return
	}
	for {
		startTime = startTime.AddDate(0, 0, 1)
		if startTime.Format(TimeLayoutYMD) == endDay {
			days = append(days, endTime)
			break
		}
		days = append(days, startTime)
	}
	return
}

// GenerateBetweenMonths 根据开始日期和结束日期生成该时间段内所有月份
func GenerateBetweenMonths(startTime, endTime time.Time) (months []time.Time) {
	if endTime.Before(startTime) {
		startTime, endTime = endTime, startTime
	}
	startMonth := startTime.Format(TimeLayoutYM)
	endMonth := endTime.Format(TimeLayoutYM)
	months = append(months, startTime)
	if startMonth == endMonth {
		return
	}
	for {
		startTime = startTime.AddDate(0, 1, 0)
		if startTime.Format(TimeLayoutYM) == endMonth {
			months = append(months, endTime)
			break
		}
		months = append(months, startTime)
	}
	return
}

// GenerateBetweenDates 根据开始日期和结束日期生成该时间段内指定的所有日期
func GenerateBetweenDates(dateType DateType, startDate, endDate string, layouts ...string) (dates []time.Time, err error) {
	layout := TimeLayoutYMD
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	layoutLength := len(layout)
	if len(startDate) != layoutLength || len(endDate) != layoutLength {
		err = errors.New("日期格式有误")
		return
	}
	var startTime, endTime time.Time
	startTime, err = time.Parse(layout, startDate)
	if err != nil {
		return
	}
	endTime, err = time.Parse(layout, endDate)
	if err != nil {
		return
	}
	switch dateType {
	case DateTypeDay:
		dates = GenerateBetweenDays(startTime, endTime)
		return
	case DateTypeMonth:
		dates = GenerateBetweenMonths(startTime, endTime)
		return
	}
	return
}
