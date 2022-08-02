package orm

// Migrate 执行迁移
func Migrate(model interface{}, options ...string) (err error) {
	if len(options) > 0 {
		if options[0] != "" {
			return Db.Set("gorm:table_options", options[0]).AutoMigrate(model)
		}
	}
	return Db.AutoMigrate(model)
}
