package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewMySQL(config *Config) (*sql.DB, error) {
	conf := &mysql.Config{
		User:                 config.User,
		Passwd:               config.Password,
		Addr:                 config.Host + config.Port,
		Net:                  "tcp",
		DBName:               config.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	return sql.Open("mysql", conf.FormatDSN())
}
