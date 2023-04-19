package demo

import (
	"encoding/base64"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	Class  uint
	Grade  uint
	Sex    bool
	School string
}

type User struct {
	ID          int64     `json:"id" gorm:"column:id;primaryKey;not_null;autoIncrement;comment:用户ID"`
	Name        string    `json:"name" gorm:"column:name;type:varchar(50);comment:用户名"`
	IdNum       string    `json:"id_num" gorm:"column:id_num;unique;comment:身份证号码"`
	Money       uint64    `json:"money" gorm:"column:money;unsigned;default:0;comment:用户余额"`
	Ages        string    `json:"ages" gorm:"column:ages;type:char(3);default:1;comment:年龄"`
	Description string    `json:"description" gorm:"column:description;index:idx_description;type:varchar(255);comment:用户描述"`
	Info        UserInfo  `json:"info" gorm:"column:info;type:json;serializer:json;comment:用户信息"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:datetime;comment:创建日期"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:datetime;comment:更新日期"`
}

func (*User) TableName() string {
	return "user"
}

func (*User) TableComment() string {
	return "用户表"
}

//func (User) TableCharset() string {
//	return "utf8mb4"
//}
//
//func (User) TableCollate() string {
//	return "utf8mb4_general_ci"
//}
//
//func (User) TableEngine() string {
//	return "InnoDB"
//}

//func (User) TableAutoIncrement() int64 {
//	return 10086
//}

//func (User) TableRename() string {
//	return "users->user_info"
//}

//func (User) TableAddColumn() []string {
//	return []string{"Ages", "Description"}
//}

//func (User) TableAlterColumn() []string {
//	return []string{"Description", "Ages"}
//}

//func (User) TableDropColumn() []string {
//	return []string{"Description", "Ages"}
//}

//func (User) TableDropIndex() []string {
//	return []string{"idx_description", "idx_ages"}
//}

//func (User) TableCreateIndex() []string {
//	return []string{"idx_money"}
//}

// Base64Encode base64_encode
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode base64_decode
func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("------------------ before create -------------------")
	fmt.Println("%+v", u)
	u.IdNum = Base64Encode(u.IdNum)
	return
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("------------------ after find -------------------")
	fmt.Println("%+v", u)
	u.IdNum = Base64Decode(u.IdNum)
	return
}

//func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
//	fmt.Println("------------------ before update -------------------")
//	fmt.Println("%+v", u)
//	return
//}
//
//func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
//	fmt.Println("------------------ after update -------------------")
//	fmt.Println("user_id:", u.ID)
//	fmt.Println("%+v", u)
//	if u.IdNum != "" {
//		u.IdNum = Base64Encode(u.IdNum)
//	}
//	return
//}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("------------------ before save/update -------------------")
	fmt.Printf("%+v", u)
	if tx.Statement.Changed("IdNum") {
		tx.Statement.SetColumn("id_num", Base64Encode(u.IdNum))
	}
	return
}
