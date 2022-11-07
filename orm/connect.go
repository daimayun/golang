package orm

import (
	"database/sql"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
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
	DefaultConnMaxIdleTime    = time.Hour

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
	ConnMaxIdleTime    time.Duration
	GormConfig         *gorm.Config
	MysqlConfig        mysql.Config
	NotAutoCreateTable bool // 不自动创建表[默认自动创建表]
	ForceResetTable    bool // 强制重置表[默认不强制重置表]
}

// ResolverConfig ORM读写分离配置数据
type ResolverConfig struct {
	Sources  []Config
	Replicas []Config
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
	if data.MysqlConfig.DSN == "" {
		data.MysqlConfig.DSN = dsnHandler(data)
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
	// 设置连接使用最大时间
	sqlDB.SetConnMaxIdleTime(data.ConnMaxIdleTime)

	return Db, nil
}

// NewResolverOrm 实例化读写分离ORM
func NewResolverOrm(data ResolverConfig) (*gorm.DB, error) {
	if len(data.Replicas) == 0 {
		if len(data.Sources) == 0 {
			return NewOrm(Config{})
		}
		return NewOrm(data.Sources[0])
	}

	var err error
	var sourcesConfigs, replicasConfigs []Config

	for _, v := range data.Sources {
		if v.Database == "" {
			err = errors.New("source database not empty")
			return Db, err
		}
		// 连接配置参数处理
		sourcesConfigs = append(sourcesConfigs, v.handler())
	}

	for _, v := range data.Replicas {
		if v.Database == "" {
			err = errors.New("replica database not empty")
			return Db, err
		}
		// 连接配置参数处理
		replicasConfigs = append(replicasConfigs, v.handler())
	}

	for k, v := range sourcesConfigs {
		if v.MysqlConfig.DSN == "" {
			sourcesConfigs[k].MysqlConfig.DSN = dsnHandler(v)
		}
	}

	for k, v := range replicasConfigs {
		if v.MysqlConfig.DSN == "" {
			replicasConfigs[k].MysqlConfig.DSN = dsnHandler(v)
		}
	}

	// GORM MASTER OPEN
	Db, err = gorm.Open(mysql.New(sourcesConfigs[0].MysqlConfig), sourcesConfigs[0].GormConfig)
	if err != nil {
		return Db, err
	}

	// Sources
	var sourcesDialector, replicasDialector []gorm.Dialector
	for _, v := range sourcesConfigs {
		sourcesDialector = append(sourcesDialector, mysql.New(v.MysqlConfig))
	}
	for _, v := range replicasConfigs {
		replicasDialector = append(replicasDialector, mysql.New(v.MysqlConfig))
	}

	// 是否自动创建表
	notAutoCreateTable = sourcesConfigs[0].NotAutoCreateTable
	// 是否强制重置表
	forceResetTable = sourcesConfigs[0].ForceResetTable

	// DBResolver
	dbResolverConfig := dbresolver.Config{
		Sources:  sourcesDialector,
		Replicas: replicasDialector,
		Policy:   dbresolver.RandomPolicy{},
	}
	err = Db.Use(dbresolver.Register(dbResolverConfig).
		SetConnMaxIdleTime(sourcesConfigs[0].ConnMaxIdleTime).
		SetConnMaxLifetime(sourcesConfigs[0].ConnectMaxLifetime).
		SetMaxIdleConns(sourcesConfigs[0].MaxIdleConnects).
		SetMaxOpenConns(sourcesConfigs[0].MaxOpenConnects))
	if err != nil {
		return Db, err
	}

	return Db, nil
}
