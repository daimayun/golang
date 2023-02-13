package orm

import "time"

type BaseModel struct {
	ID        int64     `json:"id" gorm:"column:id;primaryKey;not_null;autoIncrement;comment:ID"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;comment:创建日期"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;comment:更新日期"`
}
