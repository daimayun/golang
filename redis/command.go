package redis

import (
	"github.com/daimayun/golang/conv"
	rds "github.com/gomodule/redigo/redis"
)

// Cmd 封装后的redis执行命令
func Cmd(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	params := make([]interface{}, 0)
	params = append(params, key)
	if len(args) > 0 {
		for _, v := range args {
			params = append(params, v)
		}
	}
	return Exec(cmd, params...)
}

// Exec redigo原生执行命令
func Exec(cmd string, params ...interface{}) (interface{}, error) {
	con := Rds.Pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	return con.Do(cmd, params...)
}

// Set set命令
func Set(key string, val interface{}) (i interface{}, err error) {
	var str string
	str, err = conv.InterfaceToString(val)
	if err != nil {
		return
	}
	i, err = Cmd("set", key, str)
	return
}

// Get get命令
func Get(key string) (val string, err error) {
	var exist interface{}
	exist, err = Cmd("get", key)
	if err != nil {
		return
	}
	if exist != nil {
		return rds.String(exist, nil)
	}
	return
}

// SetEx setEx命令
func SetEx(key string, ttl int64, val interface{}) (i interface{}, err error) {
	var str string
	str, err = conv.InterfaceToString(val)
	if err != nil {
		return
	}
	i, err = Cmd("setEx", key, ttl, str)
	return
}

// Expire expire命令
func Expire(key string, ttl int64) (int64, error) {
	return rds.Int64(Cmd("expire", key, ttl))
}

// Del del命令
func Del(key string) (bool, error) {
	return rds.Bool(Cmd("del", key))
}

// Ttl ttl命令
func Ttl(key string) (int64, error) {
	return rds.Int64(Cmd("ttl", key))
}

// HMSet hMSet命令
func HMSet(key string, val interface{}) (interface{}, error) {
	return Exec("HMSet", rds.Args{}.Add(key).AddFlat(val)...)
}

// HGetAll hGetAll命令
func HGetAll(key string) ([]interface{}, error) {
	return rds.Values(Exec("HGetAll", key))
}

// Keys keys命令
func Keys(keys string) ([]string, error) {
	return rds.Strings(Exec("keys", keys))
}
