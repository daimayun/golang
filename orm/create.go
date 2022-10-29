package orm

import (
	"errors"
	"strings"
)

// 创建表数据
type createTableData struct {
	Model   interface{}
	Options *string
	Rename  *string
}

// 创建表
func createTable(data createTableData) (err error) {
	// 判断是否自动创建表
	if notAutoCreateTable {
		return
	}

	// 判断表是否存在
	var tableIsExist bool
	if Db.Migrator().HasTable(data.Model) {
		tableIsExist = true
		// 如果强制重置表
		if forceResetTable {
			// 删除表
			err = Db.Migrator().DropTable(data.Model)
			if err != nil {
				return errors.New("Drop table error [" + err.Error() + "]")
			}
			tableIsExist = false
		}
	}

	// 表不存在则创建
	if !tableIsExist {
		// 开始创建表
		db := Db
		if data.Options != nil {
			db = db.Set("gorm:table_options", *data.Options)
		}
		err = db.AutoMigrate(data.Model)
		if err != nil {
			return errors.New("Create table error [" + err.Error() + "]")
		}
	}

	// 是否需要重命名表名
	if data.Rename != nil {
		arr := strings.Split(*data.Rename, "->")
		if len(arr) == 2 {
			oldTableName := arr[0]
			newTableName := arr[1]
			if Db.Migrator().HasTable(oldTableName) && !(Db.Migrator().HasTable(newTableName)) {
				err = Db.Migrator().RenameTable(oldTableName, newTableName)
				if err != nil {
					return errors.New("Rename table error [" + err.Error() + "]")
				}
			}
		}
	}

	return
}
