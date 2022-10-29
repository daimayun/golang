package orm

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
				return
			}
			tableIsExist = false
		}
	}
	// 表存在则直接返回
	if tableIsExist {
		return
	}

	// 开始创建表
	db := Db
	if data.Options != nil {
		db = db.Set("gorm:table_options", *data.Options)
	}
	err = db.AutoMigrate(data.Model)
	if err != nil {
		return
	}
	return
}
