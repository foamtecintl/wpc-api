package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var databaseURL string

//InitDB is setting url db
func InitDB(database string) {
	databaseURL = database
	createDatabase()
}

func getConnection() *sql.DB {
	dbConnect, err := sql.Open("mysql", databaseURL)
	if err != nil {
		log.Fatalf("can not connect database : %v", err)
	}
	return dbConnect
}

func createDatabase() {
	db := getConnection()
	defer db.Close()
	sqlCreateTable := `
	CREATE TABLE IF NOT EXISTS APP_USER (
		AU_ID INTEGER PRIMARY KEY AUTO_INCREMENT,
		AU_CREATE_DATE datetime NOT NULL,
		AU_EMPLOYEE_ID varchar(255) DEFAULT NULL,
		AU_FIRST_NAME varchar(255) DEFAULT NULL,
		AU_LAST_NAME varchar(255) DEFAULT NULL,
		AU_STATUS varchar(255) NOT NULL,
		AU_USERNAME varchar(255) NOT NULL UNIQUE,
		AU_PASSWORD varchar(255) NOT NULL,
		AU_EMAIL varchar(255) DEFAULT NULL,
		AU_DEPARTMENT varchar(255) DEFAULT NULL,
		AU_PHONE varchar(255) DEFAULT NULL,
		AU_LEVEL varchar(255) NOT NULL
	);
	`
	db.Exec(sqlCreateTable)
	sqlCreateTable = `
	CREATE TABLE IF NOT EXISTS USER_LOG (
		LOG_ID INTEGER PRIMARY KEY AUTO_INCREMENT,
		LOG_CREATE_DATE datetime NOT NULL,
		LOG_STATUS varchar(255) NOT NULL,
		LOG_IP varchar(255) NOT NULL,
		LOG_AU INTEGER NOT NULL
	);
	`
	db.Exec(sqlCreateTable)
}
