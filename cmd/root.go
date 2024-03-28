package cmd

import (
	"errors"
	"fmt"

	"github.com/hidechae/gost/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gost -u root --host 127.0.0.1 -P 3306 -d test_db -t suffix_%",
		Short: "Generate golang struct definitions from MySQL table schema.",
		Run: func(cmd *cobra.Command, args []string) {
			flags := getFlags()
			if err := flags.validate(); err != nil {
				fmt.Println(err.Error())
				return
			}
			c := newMySQLConfig(flags)
			db := flags.Database
			tableHint := flags.Table

			tables, err := src.Parse(c, db, tableHint)
			if err != nil {
				cobra.CheckErr(err)
			}

			r := src.GetStructDefinitionString(tables)
			fmt.Println(r)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("user", "u", "root", "User name")
	rootCmd.PersistentFlags().StringP("host", "", "127.0.0.1", "Host address")
	rootCmd.PersistentFlags().IntP("port", "P", 3306, "Port")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Password")
	rootCmd.PersistentFlags().StringP("database", "d", "", "Database")
	rootCmd.PersistentFlags().StringP("encoding", "", "utf8mb4", "Encoding")
	rootCmd.PersistentFlags().StringP("table", "t", "%", "table name")

	_ = viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	_ = viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	_ = viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	_ = viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	_ = viper.BindPFlag("database", rootCmd.PersistentFlags().Lookup("database"))
	_ = viper.BindPFlag("encoding", rootCmd.PersistentFlags().Lookup("encoding"))
	_ = viper.BindPFlag("table", rootCmd.PersistentFlags().Lookup("table"))
}

type Flags struct {
	User     string
	Host     string
	Port     int
	Password string
	Database string
	Encoding string
	Table    string
}

func (f *Flags) validate() error {
	if f.Database == "" {
		return errors.New("database is required")
	}
	return nil
}

func getFlags() Flags {
	return Flags{
		User:     viper.GetString("user"),
		Host:     viper.GetString("host"),
		Port:     viper.GetInt("port"),
		Password: viper.GetString("password"),
		Database: viper.GetString("database"),
		Encoding: viper.GetString("encoding"),
		Table:    viper.GetString("table"),
	}
}

func newMySQLConfig(flags Flags) src.MySQLConfig {
	return src.MySQLConfig{
		User:     flags.User,
		Host:     flags.Host,
		Port:     flags.Port,
		Password: flags.Password,
		Database: "information_schema",
		Encoding: flags.Encoding,
	}
}
