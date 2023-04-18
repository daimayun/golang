package demo

import (
	"time"

	"github.com/daimayun/golang/orm"
)

type User struct {
	ID        int64      `json:"id" gorm:"column:id;primaryKey;not_null;autoIncrement;comment:ID"`
	Name      string     `json:"name" gorm:"type:varchar(255);comment:用户名"`
	Image     orm.List   `json:"image" gorm:"type:varchar(1000);comment:用户图片列表"`
	CreatedAt orm.MyTime `json:"created_at" gorm:"type:datetime;comment:创建日期"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:datetime;comment:更新日期"`
}

func (*User) TableName() string {
	return "user"
}

func (*User) TableComment() string {
	return "用户表"
}

func test() {
	var u User
	imageList := orm.List{
		"https://www.image.com/1.jpg",
		"https://www.image.com/1.jpg",
		"https://www.image.com/1.jpg",
	}
	u.Image = imageList

	imageList = orm.List{
		map[string]string{"name": "张三"},
		map[string]string{"age": "18"},
	}
	u.Image = imageList
}
