package migrate

import (
	"github.com/daimayun/golang/orm"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// AddColumn 添加字段
func AddColumn(dst interface{}, field string) error {
	return orm.Db.Migrator().AddColumn(dst, field)
}

// DropColumn 删除字段
func DropColumn(dst interface{}, field string) error {
	return orm.Db.Migrator().DropColumn(dst, field)
}

// AlterColumn 修改字段
func AlterColumn(dst interface{}, field string) error {
	return orm.Db.Migrator().AlterColumn(dst, field)
}

func Column(dst interface{}, field *schema.Field, columnType gorm.ColumnType) error {
	return orm.Db.Migrator().MigrateColumn(dst, field, columnType)
}

// HasColumn 字段是否存在
func HasColumn(dst interface{}, field string) bool {
	return orm.Db.Migrator().HasColumn(dst, field)
}

// RenameColumn 重命名字段名
func RenameColumn(dst interface{}, oldName, field string) error {
	return orm.Db.Migrator().RenameColumn(dst, oldName, field)
}

// ColumnTypes 字段类型
func ColumnTypes(dst interface{}) ([]gorm.ColumnType, error) {
	return orm.Db.Migrator().ColumnTypes(dst)
}
