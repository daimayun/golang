package orm

// Migrate 执行迁移
func Migrate(model interface{}, options ...string) (err error) {
	var tableIsExist bool
	if Db.Migrator().HasTable(model) {
		tableIsExist = true
		if autoResetTable {
			err = Db.Migrator().DropTable(model)
			if err != nil {
				return
			}
			tableIsExist = false
		}
	}
	if tableIsExist {
		return nil
	}
	if len(options) > 0 {
		if options[0] != "" {
			return Db.Set("gorm:table_options", options[0]).AutoMigrate(model)
		}
	}
	return Db.AutoMigrate(model)
}
