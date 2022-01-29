package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

var DBClient *sql.DB

func InitDbConnection(username string, password string, address string, dbport string, dbname string) {

	connString := fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v", username, password, address, dbport, dbname)

	db, err := sql.Open("sqlserver", connString)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("DB connection established")

	DBClient = db
}
