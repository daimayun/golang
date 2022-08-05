package orm

import (
	"database/sql"
	"errors"
	"github.com/daimayun/golang/conv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	DefaultHost               = "127.0.0.1"
	DefaultPort               = 3306
	DefaultNetWork            = "tcp"
	DefaultUserName           = "root"
	DefaultPassword           = "root"
	DefaultCharset            = "utf8mb4"
	DefaultLoc                = "Local"
	DefaultMaxIdleConnects    = 10
	DefaultMaxOpenConnects    = 100
	DefaultConnectMaxLifetime = time.Hour

	ParseTimeTrueVal = "True"
)

var (
	Db                 *gorm.DB
	notAutoCreateTable bool
	forceResetTable    bool
)

// Data ORM操作数据
type Data struct {
	Host               string        `json:"host"`
	Port               int           `json:"port"`
	NetWork            string        `json:"net_work"`
	UserName           string        `json:"user_name"`
	Password           string        `json:"password"`
	Database           string        `json:"database"`
	Charset            string        `json:"charset"`
	Loc                string        `json:"loc"`
	NotUseParseTime    bool          `json:"not_use_parse_time"`
	TablePrefix        string        `json:"table_prefix"`
	MaxIdleConnects    int           `json:"max_idle_connects"`
	MaxOpenConnects    int           `json:"max_open_connects"`
	ConnectMaxLifetime time.Duration `json:"connect_max_lifetime"`
	NotAutoCreateTable bool          `json:"not_auto_create_table"` // 不自动创建表[默认自动创建表]
	ForceResetTable    bool          `json:"force_reset_table"`     // 强制重置表[默认不强制重置表]
	Config             *gorm.Config  `json:"config"`
}

// Handel ORM数据处理
func (data Data) handel() Data {
	if data.Host == "" {
		data.Host = DefaultHost
	}
	if data.Port == 0 {
		data.Port = DefaultPort
	}
	if data.NetWork == "" {
		data.NetWork = DefaultNetWork
	}
	if data.UserName == "" {
		data.UserName = DefaultUserName
	}
	if data.Password == "" {
		data.Password = DefaultPassword
	}
	if data.Charset == "" {
		data.Charset = DefaultCharset
	}
	if data.Loc == "" {
		data.Loc = DefaultLoc
	}
	if data.MaxIdleConnects == 0 {
		data.MaxIdleConnects = DefaultMaxIdleConnects
	}
	if data.MaxOpenConnects == 0 {
		data.MaxOpenConnects = DefaultMaxOpenConnects
	}
	if data.ConnectMaxLifetime == 0 {
		data.ConnectMaxLifetime = DefaultConnectMaxLifetime
	}
	return data
}

// NewOrm 实例化ORM
func NewOrm(data Data) (*gorm.DB, error) {
	var err error
	if data.Database == "" {
		err = errors.New("database not empty")
		return Db, err
	}
	data = data.handel()
	dsn := data.UserName + ":" + data.Password + "@" + data.NetWork + "(" + data.Host + ":" + conv.IntToString(data.Port) + ")/" + data.Database + "?charset=" + data.Charset + "&loc=" + data.Loc
	if !data.NotUseParseTime {
		dsn += "&parseTime=" + ParseTimeTrueVal
	}
	Db, err = gorm.Open(mysql.Open(dsn), data.Config)
	if err != nil {
		return Db, err
	}
	var sqlDB *sql.DB
	sqlDB, err = Db.DB()
	if err != nil {
		return Db, err
	}
	notAutoCreateTable = data.NotAutoCreateTable
	forceResetTable = data.ForceResetTable
	sqlDB.SetMaxIdleConns(data.MaxIdleConnects)
	sqlDB.SetMaxOpenConns(data.MaxOpenConnects)
	sqlDB.SetConnMaxLifetime(data.ConnectMaxLifetime)
	return Db, nil
}
