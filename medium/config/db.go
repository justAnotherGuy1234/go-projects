package config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetUpDb() (*sql.DB, error) {
	cfg := mysql.NewConfig()

	cfg.User = "root"
	cfg.DBName = "medium"
	cfg.Addr = "127.0.0.1"
	cfg.Passwd = "hello"

	cfg.Net = "tcp"

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("error connecting to db", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("error pinging db")
		return nil, err
	}
	fmt.Println("connected to db", cfg.DBName)

	return db, nil

}
