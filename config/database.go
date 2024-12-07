package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "anakTuhan1."
	dbName := "final-projectcrud"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
