package mysql

import (
	"fmt"

	"github.com/kenji-yamane/mgr8/domain"
)

type deparser struct{}

func (d *deparser) CreateTable(table *domain.Table) string {
	// TODO: how to mount this string?
	return ""
}

func (d *deparser) DropTable(tableName string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
}

func (d *deparser) ColumnDatatype(columnName string, column *domain.Column) string {
	columnDatatype := column.Datatype
	if size, ok := column.Parameters["size"]; ok {
		columnDatatype = fmt.Sprintf("%s(%d)", column.Datatype, size)
	}
	return fmt.Sprintf("%s %v", columnName, columnDatatype)
}

func (d *deparser) ColumnDefinition(columnName string, column *domain.Column) string {
	columnDefinition := d.ColumnDatatype(columnName, column)
	if column.IsNotNull {
		columnDefinition = fmt.Sprintf("%s NOT NULL", columnDefinition)
	}
	return columnDefinition
}

func (d *deparser) AddColumn(tableName, columnName string, column *domain.Column) string {
	columnDefinition := d.ColumnDefinition(columnName, column)
	return fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", tableName, columnDefinition)
}

func (d *deparser) DropColumn(tableName, columnName string) string {
	return fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", tableName, columnName)
}
func (d *deparser) MakeColumnNotNull(tableName, columnName string, column *domain.Column) string {
	newColumn := &domain.Column{
		Datatype:   column.Datatype,
		Parameters: column.Parameters,
		IsNotNull:  true,
	}
	columnDatatype := d.ColumnDatatype(columnName, newColumn)
	return fmt.Sprintf("ALTER TABLE %s MODIFY %s NOT NULL", tableName, columnDatatype)
}

func (d *deparser) MakeColumnNullable(tableName, columnName string, column *domain.Column) string {
	newColumn := &domain.Column{
		Datatype:   column.Datatype,
		Parameters: column.Parameters,
		IsNotNull:  false,
	}
	columnDatatype := d.ColumnDatatype(columnName, newColumn)
	return fmt.Sprintf("ALTER TABLE %s MODIFY %s NULL", tableName, columnDatatype)
}
