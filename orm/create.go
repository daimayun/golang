package orm

// 创建表
func createTable(model interface{}, options ...string) (err error) {
	if notAutoCreateTable {
		return
	}
	var tableIsExist bool
	if Db.Migrator().HasTable(model) {
		tableIsExist = true
		if forceResetTable {
			err = Db.Migrator().DropTable(model)
			if err != nil {
				return
			}
			tableIsExist = false
		}
	}
	if tableIsExist {
		return
	}
	if len(options) > 0 {
		if options[0] != "" {
			return Db.Set("gorm:table_options", options[0]).AutoMigrate(model)
		}
	}
	return Db.AutoMigrate(model)
}
