package src

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	User     string
	Host     string
	Port     int
	Password string
	Database string
	Encoding string
}

// GetDSN
// ex) iam:pass@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local
func (c MySQLConfig) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset%s&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.Encoding,
	)
}

func NewGormHandler(c MySQLConfig, debug bool) (*gorm.DB, error) {
	handler, err := gorm.Open(mysql.Open(c.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if debug {
		handler = handler.Debug()
	}
	return handler, nil
}
