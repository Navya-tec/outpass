package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB
var err error

func InitDB(){
	DB,err=sql.Open("sqlite","api.db")
	if err!=nil{
		log.Fatalf("Error initialising the Database: %v",err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables(){

	createUserTable:=`
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	hostel TEXT NOT NULL,
	username TEXT NOT NULL,
	password TEXT NOT NULL
	)
	`

	if _,err:=DB.Exec(createUserTable);err!=nil{
		log.Fatalf("Couldn't create user table:%v",err)
	}

	createRequestTable:=`
	CREATE TABLE IF NOT EXISTS requests(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	proceedingTo TEXT NOT NULL,
	date_of_visit DATETIME NOT NULL,
	time_to_leave DATETIME NOT NULL,
	conveyence TEXT NOT NULL,
	status TEXT NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	if _,err:=DB.Exec(createRequestTable);err!=nil{
		log.Fatalf("Couldn't create request table:%v",err)
	}
}