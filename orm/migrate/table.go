package migrate

import "github.com/daimayun/golang/orm"

// CreateTable 创建表
func CreateTable(dst ...interface{}) error {
	return orm.Db.Migrator().CreateTable(dst...)
}

// DropTable 删除表
func DropTable(dst ...interface{}) error {
	return orm.Db.Migrator().DropTable(dst...)
}

// HasTable 表是否存在
func HasTable(dst interface{}) bool {
	return orm.Db.Migrator().HasTable(dst)
}

// RenameTable 重命名表名
func RenameTable(oldName, newName interface{}) error {
	return orm.Db.Migrator().RenameTable(oldName, newName)
}

// GetTables 获取数据库中的所有表
func GetTables() (tableList []string, err error) {
	return orm.Db.Migrator().GetTables()
}
