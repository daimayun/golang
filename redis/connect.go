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

// configOptions REDIS配置数据
type configOptions struct {
	Host           string
	Port           int
	Password       string
	Database       int
	NetWork        string
	MaxIdle        int
	MaxActive      int
	IdleTimeout    time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	ConnectTimeout time.Duration
}

type Option func(*configOptions)

func WithHost(host string) Option {
	return func(options *configOptions) {
		options.Host = host
	}
}

func WithPort(port int) Option {
	return func(options *configOptions) {
		options.Port = port
	}
}

func WithPassword(password string) Option {
	return func(options *configOptions) {
		options.Password = password
	}
}

func WithDatabase(database int) Option {
	return func(options *configOptions) {
		options.Database = database
	}
}

func WithNetWork(netWork string) Option {
	return func(options *configOptions) {
		options.NetWork = netWork
	}
}

func WithMaxIdle(maxIdle int) Option {
	return func(options *configOptions) {
		options.MaxIdle = maxIdle
	}
}

func WithMaxActive(maxActive int) Option {
	return func(options *configOptions) {
		options.MaxActive = maxActive
	}
}

func WithIdleTimeout(idleTimeout time.Duration) Option {
	return func(options *configOptions) {
		options.IdleTimeout = idleTimeout
	}
}

func WithReadTimeout(readTimeout time.Duration) Option {
	return func(options *configOptions) {
		options.ReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(options *configOptions) {
		options.WriteTimeout = writeTimeout
	}
}

func WithConnectTimeout(connectTimeout time.Duration) Option {
	return func(options *configOptions) {
		options.ConnectTimeout = connectTimeout
	}
}

func newOptions(options ...Option) configOptions {
	configOpts := &configOptions{
		Host:           "127.0.0.1",
		Port:           6379,
		Password:       "",
		Database:       0,
		NetWork:        "tcp",
		MaxIdle:        10,
		MaxActive:      0,
		IdleTimeout:    time.Hour,
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		ConnectTimeout: time.Second,
	}
	for _, option := range options {
		option(configOpts)
	}
	return *configOpts
}

// Rds Redis
var Rds *redisServer

// redisLock 锁
var redisLock *redsync.Redsync

// NewRedis 实例化Redis
func NewRedis(options ...Option) {
	data := newOptions(options...)
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
