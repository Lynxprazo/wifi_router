package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func main() {

	db, err := sql.Open("mysql", "root@tcp(0.0.0.0:3306)/Router_Db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error occurred during database connection:", err)
	}

	log.Println("Successfully connected to the database!")


	registrationQuery := `
	CREATE TABLE IF NOT EXISTS user_reg (
		user_id INT AUTO_INCREMENT,
		Username VARCHAR(255),
		Password VARCHAR(255),
		Phone_Number VARCHAR(255),
		PRIMARY KEY(user_id)
	);`

	_, err = db.Exec(registrationQuery)
	if err != nil {
		log.Fatal("Failed to create table user_reg:", err)
	}

	log.Println("Table user_reg created successfully!")


	loginQuery := `
	CREATE TABLE IF NOT EXISTS user_log (
		user_id2 INT,
		username VARCHAR(255),
		Password VARCHAR(255),
		FOREIGN KEY (user_id2) REFERENCES user_reg(user_id)
	);`

	_, err = db.Exec(loginQuery)
	if err != nil {
		log.Fatal("Failed to create table user_log:", err)
	}

	log.Println("Table user_log created successfully!")
}