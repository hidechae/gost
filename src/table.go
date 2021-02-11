package src

import (
	"fmt"
	"strings"
	"time"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type Table struct {
	TableCatalog   string     `gorm:"column:TABLE_CATALOG"`
	TableSchema    string     `gorm:"column:TABLE_SCHEMA"`
	TableName      string     `gorm:"column:TABLE_NAME"`
	TableType      string     `gorm:"column:TABLE_TYPE"`
	Engine         *string    `gorm:"column:ENGINE"`
	Version        *uint64    `gorm:"column:VERSION"`
	RowFormat      *string    `gorm:"column:ROW_FORMAT"`
	TableRows      *uint64    `gorm:"column:TABLE_ROWS"`
	AvgRowLength   *uint64    `gorm:"column:AVG_ROW_LENGTH"`
	DataLength     *uint64    `gorm:"column:DATA_LENGTH"`
	MaxDataLength  *uint64    `gorm:"column:MAX_DATA_LENGTH"`
	IndexLength    *uint64    `gorm:"column:INDEX_LENGTH"`
	DataFree       *uint64    `gorm:"column:DATA_FREE"`
	AutoIncrement  *uint64    `gorm:"column:AUTO_INCREMENT"`
	CreateTime     *time.Time `gorm:"column:CREATE_TIME"`
	UpdateTime     *time.Time `gorm:"column:UPDATE_TIME"`
	CheckTime      *time.Time `gorm:"column:CHECK_TIME"`
	TableCollation *string    `gorm:"column:TABLE_COLLATION"`
	Checksum       *uint64    `gorm:"column:CHECKSUM"`
	CreateOptions  *string    `gorm:"column:CREATE_OPTIONS"`
	TableComment   string     `gorm:"column:TABLE_COMMENT"`

	// Association
	Columns []Column `gorm:"foreignKey:TableName;references:TableName"`
}

func (t *Table) GetSingularTableName() string {
	plu := pluralize.NewClient()
	return plu.Singular(t.TableName)
}

func (t *Table) GetDefinition() (string, error) {
	columnDefinitions := make([]string, 0, len(t.Columns))
	for _, c := range t.Columns {
		r, err := c.GetDefinition()
		if err != nil {
			return "", err
		}
		columnDefinitions = append(columnDefinitions, r)
	}
	return fmt.Sprintf(
		"type %s struct {\n\t%s\n}",
		strcase.ToCamel(t.GetSingularTableName()),
		strings.Join(columnDefinitions, "\n\t"),
	), nil
}
