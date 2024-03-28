package src

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Parse(c MySQLConfig, db string, tableHint string) ([]Table, error) {
	h, err := NewGormHandler(c, false)
	if err != nil {
		return nil, err
	}

	var tables []Table
	h.Preload("Columns", func(db *gorm.DB) *gorm.DB {
		return db.Order("ordinal_position")
	}).
		Where(fmt.Sprintf("table_schema = '%s' and table_name like '%s'", db, tableHint)).
		Find(&tables)

	return tables, nil
}

func GetStructDefinitionString(tables []Table) string {
	var s string
	for _, table := range tables {
		r, err := table.GetDefinition()
		if err != nil {
			cobra.CheckErr(err)
		}
		s += fmt.Sprintln(r + "\n")
	}
	return s
}
