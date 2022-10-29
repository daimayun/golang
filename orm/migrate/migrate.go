package migrate

import "github.com/daimayun/golang/orm"

// AutoMigrate AutoMigrate
func AutoMigrate(dst ...interface{}) error {
	return orm.Db.Migrator().AutoMigrate(dst...)
}
