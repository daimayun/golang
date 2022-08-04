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

// Orm ORM操作数据
type Orm struct {
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
func (o Orm) Handel() Orm {
	if o.Host == "" {
		o.Host = DefaultHost
	}
	if o.Port == 0 {
		o.Port = DefaultPort
	}
	if o.NetWork == "" {
		o.NetWork = DefaultNetWork
	}
	if o.UserName == "" {
		o.UserName = DefaultUserName
	}
	if o.Password == "" {
		o.Password = DefaultPassword
	}
	if o.Charset == "" {
		o.Charset = DefaultCharset
	}
	if o.Loc == "" {
		o.Loc = DefaultLoc
	}
	if o.MaxIdleConnects == 0 {
		o.MaxIdleConnects = DefaultMaxIdleConnects
	}
	if o.MaxOpenConnects == 0 {
		o.MaxOpenConnects = DefaultMaxOpenConnects
	}
	if o.ConnectMaxLifetime == 0 {
		o.ConnectMaxLifetime = DefaultConnectMaxLifetime
	}
	return o
}

// NewOrm 实例化ORM
func NewOrm(data Orm) (*gorm.DB, error) {
	var err error
	if data.Database == "" {
		err = errors.New("database not empty")
		return Db, err
	}
	data = data.Handel()
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