package migrate

import (
	"github.com/daimayun/golang/orm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// CurrentDatabases 当前数据库名
func CurrentDatabases() string {
	return orm.Db.Migrator().CurrentDatabase()
}

// FullDataTypeOf Full Data Type Of
func FullDataTypeOf(field *schema.Field) clause.Expr {
	return orm.Db.Migrator().FullDataTypeOf(field)
}
