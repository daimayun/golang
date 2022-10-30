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

// IOptionTableAddColumn 表添加字段
type IOptionTableAddColumn interface {
	TableAddColumn() []string
}

// IOptionTableDropColumn 表删除字段
type IOptionTableDropColumn interface {
	TableDropColumn() []string
}

// IOptionTableAlterColumn 表修改字段
type IOptionTableAlterColumn interface {
	TableAlterColumn() []string
}

// IOptionTableCreateIndex 表添加索引
type IOptionTableCreateIndex interface {
	TableCreateIndex() []string
}

// IOptionTableDropIndex 表删除索引
type IOptionTableDropIndex interface {
	TableDropIndex() []string
}
