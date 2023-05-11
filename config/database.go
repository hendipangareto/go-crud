package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBCOnnection() (*sql.DB, error) {
	dbDriver 		:= "mysql"
	dbUser 			:= "root"
	dbPassword 		:= ""
	dbName 			:= "go_crud"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	return db, err
}