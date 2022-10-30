package orm

import (
	"fmt"
	"github.com/daimayun/golang/conv"
	"github.com/daimayun/golang/function"
	"reflect"
	"strings"
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
		data := createTableData{Model: model}
		// 获取Option
		var options, sign string
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
		if options != "" {
			data.Options = &options
		}
		// 是否修改表名
		tableRename := getTableRename(val)
		if tableRename != "" {
			data.Rename = &tableRename
		}
		var indexSlice []string
		// 是否添加字段
		tableAddColumn := getTableAddColumn(val)
		if len(tableAddColumn) > 0 {
			data.AddColumn = &tableAddColumn
			for _, v := range tableAddColumn { // 处理普通索引
				indexName := getFieldIndexName(function.GetValueByInterface(model), v)
				if indexName != "" {
					indexSlice = append(indexSlice, indexName)
				}
			}
		}
		// 是否修改字段
		tableAlterColumn := getTableAlterColumn(val)
		if len(tableAlterColumn) > 0 {
			data.AlterColumn = &tableAlterColumn
			for _, v := range tableAlterColumn { // 处理普通索引
				indexName := getFieldIndexName(function.GetValueByInterface(model), v)
				if indexName != "" {
					indexSlice = append(indexSlice, indexName)
				}
			}
		}
		tableCreateIndex := getTableCreateIndex(val)
		if len(tableCreateIndex) > 0 || len(indexSlice) > 0 {
			indexSlice = append(indexSlice, tableCreateIndex...)
		}
		// 是否添加索引
		if len(indexSlice) > 0 {
			data.CreateIndex = &indexSlice
		}
		// 是否删除表字段
		tableDropColumn := getTableDropColumn(val)
		if len(tableDropColumn) > 0 {
			data.DropColumn = &tableDropColumn
		}
		// 是否删除索引
		tableDropIndex := getTableDropIndex(val)
		if len(tableDropIndex) > 0 {
			data.DropIndex = &tableDropIndex
		}
		err = createTable(data)
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

// getOrmSliceString 获取表设置字符串切片相关类型
func getOrmSliceString(val reflect.Value, key string) []string {
	if fun := val.MethodByName(key); fun.IsValid() {
		values := fun.Call([]reflect.Value{})
		if len(values) > 0 && values[0].Kind() == reflect.Slice {
			return values[0].Interface().([]string)
		}
	}
	return []string{}
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

// getTableAddColumn Get model struct table add column
func getTableAddColumn(val reflect.Value) []string {
	return getOrmSliceString(val, "TableAddColumn")
}

// getTableDropColumn Get model struct table drop column
func getTableDropColumn(val reflect.Value) []string {
	return getOrmSliceString(val, "TableDropColumn")
}

// getTableAlterColumn Get model struct table alter column
func getTableAlterColumn(val reflect.Value) []string {
	return getOrmSliceString(val, "TableAlterColumn")
}

// getTableCreateIndex Get model struct table create index
func getTableCreateIndex(val reflect.Value) []string {
	return getOrmSliceString(val, "TableCreateIndex")
}

// getTableDropIndex Get model struct table drop index
func getTableDropIndex(val reflect.Value) []string {
	return getOrmSliceString(val, "TableDropIndex")
}

// getFieldIndexName 获取字段索引名称
func getFieldIndexName(model interface{}, field string) string {
	label := "idx_"
	str, ok := function.GetStructTag(model, field, "gorm")
	if ok {
		var hasIndex bool
		var column, index string
		arr := strings.Split(str, ";")
		for _, v := range arr {
			s := strings.Split(v, ":")
			sLength := len(s)
			if sLength > 0 {
				if s[0] == "column" {
					if sLength > 1 {
						column = s[1]
					}
				}
				if s[0] == "index" {
					hasIndex = true
					if sLength == 1 && column != "" {
						break
					}
					if sLength == 2 {
						index = s[1]
						break
					}
				}
			}
		}
		if hasIndex {
			if index == "" && column != "" {
				index = column
			}
			if index != "" {
				if strings.Index(index, label) != 0 {
					index = label + index
				}
				return index
			}
		}
	}
	return ""
}
