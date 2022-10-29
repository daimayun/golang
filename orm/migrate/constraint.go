package migrate

import "github.com/daimayun/golang/orm"

// CreateConstraint Create Constraint
func CreateConstraint(dst interface{}, name string) error {
	return orm.Db.Migrator().CreateConstraint(dst, name)
}

// DropConstraint Drop Constraint
func DropConstraint(dst interface{}, name string) error {
	return orm.Db.Migrator().DropConstraint(dst, name)
}

// HasConstraint Has Constraint
func HasConstraint(dst interface{}, name string) bool {
	return orm.Db.Migrator().HasConstraint(dst, name)
}
