package redis

import (
	"github.com/daimayun/golang/conv"
	"github.com/go-redsync/redsync"
	rds "github.com/gomodule/redigo/redis"
	"time"
)

// redisServer Redis Server
type redisServer struct {
	Pool *rds.Pool
}

// Data 连接数据
type Data struct {
	Host           string        `json:"host"`
	Port           int           `json:"port"`
	Password       string        `json:"password"`
	Database       int           `json:"database"`
	NetWork        string        `json:"net_work"`
	MaxIdle        int           `json:"max_idle"`
	MaxActive      int           `json:"max_active"`
	IdleTimeout    time.Duration `json:"idle_timeout"`
	ReadTimeout    time.Duration `json:"read_timeout"`
	WriteTimeout   time.Duration `json:"write_timeout"`
	ConnectTimeout time.Duration `json:"connect_timeout"`
}

const (
	DefaultHost           = "127.0.0.1"
	DefaultPort           = 6379
	DefaultMaxIdle        = 10
	DefaultIdleTimeout    = time.Hour
	DefaultNetWork        = "tcp"
	DefaultReadTimeout    = time.Second
	DefaultWriteTimeout   = time.Second
	DefaultConnectTimeout = time.Second
)

// Rds Redis
var Rds *redisServer

// redisLock 锁
var redisLock *redsync.Redsync

func (data Data) handel() Data {
	if data.Host == "" {
		data.Host = DefaultHost
	}
	if data.Port == 0 {
		data.Port = DefaultPort
	}
	if data.MaxIdle == 0 {
		data.MaxIdle = DefaultMaxIdle
	}
	if data.NetWork == "" {
		data.NetWork = DefaultNetWork
	}
	if data.IdleTimeout == 0 {
		data.IdleTimeout = DefaultIdleTimeout
	}
	if data.ReadTimeout == 0 {
		data.ReadTimeout = DefaultReadTimeout
	}
	if data.WriteTimeout == 0 {
		data.WriteTimeout = DefaultWriteTimeout
	}
	if data.ConnectTimeout == 0 {
		data.ConnectTimeout = DefaultConnectTimeout
	}
	return data
}

// NewRedis 实例化Redis
func NewRedis(data Data) {
	data = data.handel()
	Rds = new(redisServer)
	Rds.Pool = &rds.Pool{
		MaxIdle:     data.MaxIdle,
		MaxActive:   data.MaxActive,
		IdleTimeout: data.IdleTimeout,
		Dial: func() (rds.Conn, error) {
			return rds.Dial(
				data.NetWork,
				data.Host+":"+conv.IntToString(data.Port),
				rds.DialReadTimeout(data.ReadTimeout),
				rds.DialWriteTimeout(data.WriteTimeout),
				rds.DialConnectTimeout(data.ConnectTimeout),
				rds.DialPassword(data.Password),
				rds.DialDatabase(data.Database),
			)
		},
		TestOnBorrow: func(c rds.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	redisLock = redsync.New([]redsync.Pool{Rds.Pool})
}
