package orm

// IOptionTableComment 表备注
type IOptionTableComment interface {
	TableComment() string
}

// IOptionTableEngine 表引擎
type IOptionTableEngine interface {
	TableEngine() string
}

// IOptionTableCharset 表字符集
type IOptionTableCharset interface {
	TableCharset() string
}

// IOptionTableCollate 表排序规则
type IOptionTableCollate interface {
	TableCollate() string
}

// IOptionTableAutoIncrement 表主键自增长起始值
type IOptionTableAutoIncrement interface {
	TableAutoIncrement() int64
}

// IOptionTableRename 重命名表
type IOptionTableRename interface {
	TableRename() string
}
