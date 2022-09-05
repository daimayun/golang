package times

var (
	TimeLayoutYMDHIS string = "20060102150405"
	TimeLayout       string = "2006-01-02 15:04:05"
	TimeLayoutYMD    string = "2006-01-02"
	TimeLayoutYM     string = "2006-01"
)

// TimeVariable 时间变量[有可能是正负数]
var TimeVariable int64

type DateType string

var (
	// DateTypeDay 天日期类型
	DateTypeDay DateType = "day"
	// DateTypeMonth 月日期类型
	DateTypeMonth DateType = "month"
)
