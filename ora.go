package gorm

import (
	"fmt"
)

type ora struct {
	commonDialect
}

func (ora) SelectFromDummyTable() string {
	return "FROM DUAL"
}

func (ora) Quote(key string) string {
	return fmt.Sprintf("%s", key)
}

func (c ora) HasTable(scope *Scope, tableName string) bool {
	var count int
	c.RawScanInt(scope, &count, "select * from user_tables where table_name = ?", tableName)
	return count > 0
}

func (c ora) HasColumn(scope *Scope, tableName string, columnName string) bool {
	var count int
	c.RawScanInt(scope, &count, "select * from user_tab_cols where table_name = ? and column_name = ?",  tableName, columnName)
	return count > 0
}

func (c ora) HasIndex(scope *Scope, tableName string, indexName string) bool {
	var count int
	c.RawScanInt(scope, &count, "select count(*) from user_indexes where index_name = ?", indexName)
	return count > 0
}

func (ora) RemoveIndex(scope *Scope, indexName string) {
	scope.Err(scope.NewDB().Exec(fmt.Sprintf("drop index %v", indexName)).Error)
}

func (ora) CurrentDatabase(scope *Scope) (name string) {
	scope.Err(scope.NewDB().Raw("select SYS_CONTEXT ('userenv', 'current_schema') as name from dual").Row().Scan(&name))
	return
}
