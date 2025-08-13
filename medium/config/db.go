package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SetUpDb() (*sql.DB, error) {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("error loading env")
		return nil, err
	}
	cfg := mysql.NewConfig()

	cfg.User = "root"
	cfg.DBName = "medium"
	cfg.Addr = "127.0.0.1"
	res, ok := os.LookupEnv("DB_PASSWORD")

	if !ok {
		fmt.Println("error getting db password from env")
	}

	cfg.Passwd = res

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
