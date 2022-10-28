package orm

import (
	"database/sql"
	"errors"
	"github.com/daimayun/golang/conv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// ORM配置默认值
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

var Db *gorm.DB

var notAutoCreateTable, forceResetTable bool

// Config ORM配置数据
type Config struct {
	Host               string
	Port               int
	NetWork            string
	UserName           string
	Password           string
	Database           string
	Charset            string
	Loc                string
	NotUseParseTime    bool
	TablePrefix        string
	MaxIdleConnects    int
	MaxOpenConnects    int
	ConnectMaxLifetime time.Duration
	GormConfig         *gorm.Config
	MysqlConfig        mysql.Config
	NotAutoCreateTable bool // 不自动创建表[默认自动创建表]
	ForceResetTable    bool // 强制重置表[默认不强制重置表]
}

// NewOrm 实例化ORM
func NewOrm(data Config) (*gorm.DB, error) {
	var err error

	// 数据库不能为空
	if data.Database == "" {
		err = errors.New("database not empty")
		return Db, err
	}

	// ORM连接配置参数处理
	data = data.handler()

	// DSN data source name
	dsn := data.UserName + ":" + data.Password + "@" + data.NetWork + "(" + data.Host + ":" + conv.IntToString(data.Port) + ")/" + data.Database + "?charset=" + data.Charset + "&loc=" + data.Loc
	if !data.NotUseParseTime {
		dsn += "&parseTime=" + ParseTimeTrueVal
	}
	if data.MysqlConfig.DSN == "" {
		data.MysqlConfig.DSN = dsn
	}

	// GORM OPEN
	Db, err = gorm.Open(mysql.New(data.MysqlConfig), data.GormConfig)
	if err != nil {
		return Db, err
	}

	// 是否自动创建表
	notAutoCreateTable = data.NotAutoCreateTable
	// 是否强制重置表
	forceResetTable = data.ForceResetTable

	// 使用连接池
	var sqlDB *sql.DB
	sqlDB, err = Db.DB()
	if err != nil {
		return Db, err
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(data.MaxIdleConnects)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(data.MaxOpenConnects)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(data.ConnectMaxLifetime)

	return Db, nil
}
