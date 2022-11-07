package orm

import "github.com/daimayun/golang/conv"

// Handel ORM连接配置参数处理助手
func (data Config) handler() Config {
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
	if data.ConnMaxIdleTime == 0 {
		data.ConnMaxIdleTime = DefaultConnMaxIdleTime
	}
	return data
}

// DSN data source name
func dsnHandler(data Config) string {
	dsn := data.UserName + ":" + data.Password + "@" + data.NetWork + "(" + data.Host + ":" + conv.IntToString(data.Port) + ")/" + data.Database + "?charset=" + data.Charset + "&loc=" + data.Loc
	if !data.NotUseParseTime {
		dsn += "&parseTime=" + ParseTimeTrueVal
	}
	return dsn
}
