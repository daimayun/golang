package orm

import (
	"errors"
	"strings"
)

// 创建表数据
type createTableData struct {
	Model       interface{}
	Options     *string
	Rename      *string
	AddColumn   *[]string
	DropColumn  *[]string
	AlterColumn *[]string
	CreateIndex *[]string
	DropIndex   *[]string
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

	// 是否添加字段
	if data.AddColumn != nil {
		for _, v := range *data.AddColumn {
			if !(Db.Migrator().HasColumn(data.Model, v)) {
				err = Db.Migrator().AddColumn(data.Model, v)
				if err != nil {
					return errors.New("Add Column error [" + err.Error() + "]")
				}
			}
		}
	}

	// 是否修改字段
	if data.AlterColumn != nil {
		for _, v := range *data.AlterColumn {
			if Db.Migrator().HasColumn(data.Model, v) {
				err = Db.Migrator().AlterColumn(data.Model, v)
				if err != nil {
					return errors.New("Alter Column error [" + err.Error() + "]")
				}
			}
		}
	}

	// 是否添加普通索引
	if data.CreateIndex != nil {
		for _, v := range *data.CreateIndex {
			if !(Db.Migrator().HasIndex(data.Model, v)) {
				err = Db.Migrator().CreateIndex(data.Model, v)
				if err != nil {
					return errors.New("Create index error [" + err.Error() + "]")
				}
			}
		}
	}

	// 是否删除字段
	if data.DropColumn != nil {
		for _, v := range *data.DropColumn {
			if Db.Migrator().HasColumn(data.Model, v) {
				err = Db.Migrator().DropColumn(data.Model, v)
				if err != nil {
					return errors.New("Drop column error [" + err.Error() + "]")
				}
			}
		}
	}

	// 是否删除索引
	if data.DropIndex != nil {
		for _, v := range *data.DropIndex {
			if Db.Migrator().HasIndex(data.Model, v) {
				err = Db.Migrator().DropIndex(data.Model, v)
				if err != nil {
					return errors.New("Drop index error [" + err.Error() + "]")
				}
			}
		}
	}

	return
}
