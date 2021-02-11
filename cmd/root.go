package cmd

import (
	"fmt"

	"github.com/hidechae/gost/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gost -u root --host 127.0.0.1 -P 3306 -d test_db -t suffix_%",
		Short: "Create golang struct definitions from MySQL table schema.",
		Run: func(cmd *cobra.Command, args []string) {
			c := newMySQLConfig()
			h, err := src.NewGormHandler(c, false)
			if err != nil {
				cobra.CheckErr(err)
			}

			db := viper.GetString("database")
			table := viper.GetString("table")
			var tables []src.Table
			h.Preload("Columns", fmt.Sprintf("table_schema = '%s'", db)).
				Where(fmt.Sprintf("table_schema = '%s' and table_name like '%s'", db, table)).
				Find(&tables)

			for _, table := range tables {
				r, err := table.GetDefinition()
				if err != nil {
					cobra.CheckErr(err)
				}
				fmt.Println(r + "\n")
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("user", "u", "root", "User name")
	rootCmd.PersistentFlags().StringP("host", "", "127.0.0.1", "Host address")
	rootCmd.PersistentFlags().StringP("port", "P", "3306", "Port")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Password")
	rootCmd.PersistentFlags().StringP("database", "d", "", "Database")
	rootCmd.PersistentFlags().StringP("encoding", "", "utf8mb4", "Encoding")
	rootCmd.PersistentFlags().StringP("table", "t", "", "table name")

	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("database", rootCmd.PersistentFlags().Lookup("database"))
	viper.BindPFlag("encoding", rootCmd.PersistentFlags().Lookup("encoding"))
	viper.BindPFlag("table", rootCmd.PersistentFlags().Lookup("table"))
}

func newMySQLConfig() src.MySQLConfig {
	return src.MySQLConfig{
		User:     viper.GetString("user"),
		Host:     viper.GetString("host"),
		Port:     viper.GetInt("port"),
		Password: viper.GetString("password"),
		Database: "information_schema",
		Encoding: viper.GetString("encoding"),
	}
}
