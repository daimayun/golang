package orm

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Time time.Time

// MarshalJSON Json后的数据处理
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).In(time.Local).Format("2006-01-02 15:04:05") + `"`), nil
}

// Value 存入数据库
func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

type List []any

// Value 存入数据库
func (l List) Value() (driver.Value, error) {
	return json.Marshal(l)
}

// Scan 从数据库读出
func (l *List) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &l)
}
