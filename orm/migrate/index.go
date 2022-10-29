package migrate

import "github.com/daimayun/golang/orm"

// CreateIndex 创建索引
func CreateIndex(dst interface{}, name string) error {
	return orm.Db.Migrator().CreateIndex(dst, name)
}

// DropIndex 删除索引
func DropIndex(dst interface{}, name string) error {
	return orm.Db.Migrator().DropIndex(dst, name)
}

// HasIndex 索引是否存在
func HasIndex(dst interface{}, name string) bool {
	return orm.Db.Migrator().HasIndex(dst, name)
}

// RenameIndex 重命名索引
func RenameIndex(dst interface{}, oldName, newName string) error {
	return orm.Db.Migrator().RenameIndex(dst, oldName, newName)
}
