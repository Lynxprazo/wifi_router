package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" 
)
var DB *sql.DB
func InitDB() {
     var err error
	DB, err = sql.Open("mysql", "root@tcp(0.0.0.0:3306)/Router_Db?parseTime=true")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
    // configuration of new connection pool
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime( 5 *time.Minute)

	// Ping the database to verify the connection
	if err := DB.Ping(); err != nil {
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

	_, err = DB.Exec(registrationQuery)
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

	_, err = DB.Exec(loginQuery)
	if err != nil {
		log.Fatal("Failed to create table user_log:", err)
	}

	log.Println("Table user_log created successfully!")
}