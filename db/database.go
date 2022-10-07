package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//Save connection
const url = "mysql_user:An0thrS3crt@tcp(localhost:3306)/mysql_db"

var db *sql.DB

//Connect to the DB
func Connect() {
	connection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conexión Exitosa")
	db = connection
}

//Disconnect to the DB
func Disconnect() {
	db.Close()
}

// Health : Show error if the connection fail
func Health() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//ExistTable : Check if the table is already created
func ExistTable(tableName string) bool {
	sqltable := fmt.Sprintf("SHOW TABLES LIKE '%s' ", tableName)
	rows, err := Query(sqltable)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

//CreateTable : Create a new table
func CreateTable(schema string, name string) {
	if !ExistTable(name) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//TruncateTable : Clean table records
func TruncateTable(tableName string) {
	sqlTable := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sqlTable)
}

//Exec : Execute an action inside a table
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

//Query : Perform a query within a table
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}
