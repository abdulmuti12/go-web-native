package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDb() {
	// db, err := sql.Open("mysql", "root:@/go_product") //connecting to db
	db, err := sql.Open("mysql", "root:@/go_product?parseTime=true")

	// db, err := sql.Open("mysql", "root:@/?parseTime=true")
	// db, err := sqlx.Connect("mysql", "myuser:mypass@tcp(127.0.0.1:3306)/mydb?parseTime=true")

	if err != nil {
		log.Println("Error connecting to the database:", err)
		log.Println("Please check your database credentials.")
		return
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging the database:", err)
		log.Println("Please check your database credentials.")
		return
	}

	DB = db

	log.Println("Database connected")
}
