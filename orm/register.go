package orm

import (
	"fmt"
	"github.com/daimayun/golang/conv"
	"reflect"
)

// RegisterModel 注册模型
func RegisterModel(models ...interface{}) (err error) {
	for _, model := range models {
		val := reflect.ValueOf(model)
		typ := reflect.Indirect(val).Type()
		if val.Kind() != reflect.Ptr {
			err = fmt.Errorf("cannot use non-ptr model struct `%s`", getFullName(typ))
			return
		}
		if typ.Kind() == reflect.Ptr {
			err = fmt.Errorf("only allow ptr model struct, it looks you use two reference to the struct `%s`", typ)
			return
		}
		if val.Elem().Kind() == reflect.Slice {
			val = reflect.New(val.Elem().Type().Elem())
		}
		options := ""
		sign := ""
		tableEngine := getTableEngine(val)
		if tableEngine != "" {
			options += sign + "ENGINE=" + tableEngine
			sign = " "
		}
		tableAutoIncrement := getTableAutoIncrement(val)
		if tableAutoIncrement > 0 {
			options += sign + "AUTO_INCREMENT=" + conv.Int64ToString(tableAutoIncrement)
			sign = " "
		}
		tableCharset := getTableCharset(val)
		if tableCharset != "" {
			options += sign + "CHARSET=" + tableCharset
			sign = " "
		}
		tableCollate := getTableCollate(val)
		if tableCollate != "" {
			options += sign + "COLLATE=" + tableCollate
			sign = " "
		}
		tableComment := getTableComment(val)
		if tableComment != "" {
			options += sign + "COMMENT='" + tableComment + "'"
			sign = " "
		}
		err = createTable(model, options)
		if err != nil {
			return
		}
	}
	return
}

// get reflect.Type name with package path.
func getFullName(typ reflect.Type) string {
	return typ.PkgPath() + "." + typ.Name()
}

// getOrmString 获取表设置字符串相关类型
func getOrmString(val reflect.Value, key string) string {
	if fun := val.MethodByName(key); fun.IsValid() {
		values := fun.Call([]reflect.Value{})
		if len(values) > 0 && values[0].Kind() == reflect.String {
			return values[0].String()
		}
	}
	return ""
}

// getOrmInt64 获取表设置整型相关类型
func getOrmInt64(val reflect.Value, key string) int64 {
	if fun := val.MethodByName(key); fun.IsValid() {
		values := fun.Call([]reflect.Value{})
		if len(values) > 0 && values[0].Kind() == reflect.Int64 {
			return values[0].Int()
		}
	}
	return 0
}

// getTableEngine Get model struct table engine
func getTableEngine(val reflect.Value) string {
	return getOrmString(val, "TableEngine")
}

// getTableComment Get model struct table comment
func getTableComment(val reflect.Value) string {
	return getOrmString(val, "TableComment")
}

// getTableCharset Get model struct table charset
func getTableCharset(val reflect.Value) string {
	return getOrmString(val, "TableCharset")
}

// getTableCollate Get model struct table collate
func getTableCollate(val reflect.Value) string {
	return getOrmString(val, "TableCollate")
}

// getTableAutoIncrement Get model struct table auto increment
func getTableAutoIncrement(val reflect.Value) int64 {
	return getOrmInt64(val, "TableAutoIncrement")
}

// getTableName Get model struct table name
func getTableName(val reflect.Value) string {
	return getOrmString(val, "TableName")
}

// getTableRename Get model struct table rename
func getTableRename(val reflect.Value) string {
	return getOrmString(val, "TableRename")
}
