package demo

import "github.com/daimayun/golang/orm"

func list() {
	var u []User
	orm.Db.Scopes(orm.Paginate(1, 10)).Order("id desc").Find(&u)
}
