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

// Set Set命令
func Set(key string, val interface{}) (i interface{}, err error) {
	var str string
	str, err = conv.InterfaceToString(val)
	if err != nil {
		return
	}
	i, err = Cmd("set", key, str)
	return
}

// SetEx SetEx命令
func SetEx(key string, val interface{}, ttl int64) (i interface{}, err error) {
	var str string
	str, err = conv.InterfaceToString(val)
	if err != nil {
		return
	}
	i, err = Cmd("setEx", key, ttl, str)
	return
}

func Expire(key string, ttl int64) (int, error) {
	return rds.Int(Cmd("expire", key, ttl))
}

func Del(key string) (bool, error) {
	return rds.Bool(Cmd("del", key))
}

func Ttl(key string) (int64, error) {
	return rds.Int64(Cmd("ttl", key))
}

func HMSet(key string, val interface{}) (interface{}, error) {
	return Exec("HMSet", rds.Args{}.Add(key).AddFlat(val)...)
}
