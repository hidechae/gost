package src

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/iancoleman/strcase"
)

type Column struct {
	TableCatalog           string  `gorm:"column:TABLE_CATALOG"`
	TableSchema            string  `gorm:"column:TABLE_SCHEMA"`
	TableName              string  `gorm:"column:TABLE_NAME"`
	ColumnName             string  `gorm:"column:COLUMN_NAME"`
	OrdinalPosition        uint64  `gorm:"column:ORDINAL_POSITION"`
	ColumnDefault          *string `gorm:"column:COLUMN_DEFAULT"`
	IsNullable             string  `gorm:"column:IS_NULLABLE"`
	DataType               string  `gorm:"column:DATA_TYPE"`
	CharacterMaximumLength *uint64 `gorm:"column:CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   *uint64 `gorm:"column:CHARACTER_OCTET_LENGTH"`
	NumericPrecision       *uint64 `gorm:"column:NUMERIC_PRECISION"`
	NumericScale           *uint64 `gorm:"column:NUMERIC_SCALE"`
	DatetimePrecision      *uint64 `gorm:"column:DATETIME_PRECISION"`
	CharacterSetName       *string `gorm:"column:CHARACTER_SET_NAME"`
	CollationName          *string `gorm:"column:COLLATION_NAME"`
	ColumnType             string  `gorm:"column:COLUMN_TYPE"`
	ColumnKey              string  `gorm:"column:COLUMN_KEY"`
	Extra                  string  `gorm:"column:EXTRA"`
	Privileges             string  `gorm:"column:PRIVILEGES"`
	ColumnComment          string  `gorm:"column:COLUMN_COMMENT"`
}

func (c *Column) GetColumnName() string {
	camelName := strcase.ToCamel(c.ColumnName)
	re := regexp.MustCompile("Id$")
	return re.ReplaceAllString(camelName, "ID")
}

func (c *Column) IsUnsigned() bool {
	b, _ := regexp.MatchString(`^.*unsigned$`, c.ColumnType)
	return b
}

func (c *Column) GetDefinition() (string, error) {
	dataType, err := c.GetGolangDataType()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%s %s",
		c.GetColumnName(),
		dataType,
	), nil
}

func (c *Column) GetGolangDataType() (string, error) {
	var t string
	switch c.DataType {
	// integer
	case "tinyint":
		t = "int8"
	case "smallint":
		t = "int16"
	case "mediumint":
		fallthrough
	case "int":
		t = "int"
	case "bigint":
		t = "int64"

	// float
	case "float":
		t = "float32"
	case "double":
		t = "float64"

	// decimal TODO
	case "decimal":
		t = "decimal.Decimal /* TODO */"

	// text
	case "varchar":
		fallthrough
	case "text":
		fallthrough
	case "longtext":
		t = "string"

	// time
	case "date":
		fallthrough
	case "datetime":
		fallthrough
	case "timestamp":
		t = "time.Time"

	default:
		err := errors.New(fmt.Sprintf("unknown data type `%s` for %s:%s", c.DataType, c.TableName, c.ColumnName))
		return "", err
	}

	if c.IsUnsigned() {
		t = "u" + t
	}

	if c.IsNullable == "YES" {
		t = "*" + t
	}
	return t, nil
}
