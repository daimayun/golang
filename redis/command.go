package redis

import "github.com/daimayun/golang/conv"

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
