package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MYSQL_USER     = "MYSQL_USER"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"
	MYSQL_HOST     = "MYSQL_HOST"
	MYSQL_PORT     = "MYSQL_PORT"
	MYSQL_DB       = "MYSQL_DB"
)

func NewMySQLConnection() (*sql.DB, error) {
	mysqlUser := os.Getenv(MYSQL_USER)
	mysqlPassword := os.Getenv(MYSQL_PASSWORD)
	mysqlHost := os.Getenv(MYSQL_HOST)
	mysqlPort := os.Getenv(MYSQL_PORT)
	mysqlDB := os.Getenv(MYSQL_DB)

	// Construct the MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDB)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
