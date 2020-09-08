package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// DB global database env
var DB *sqlx.DB

// DBConn : open database connection
func DBConn() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	dbInfo := "user=" + fmt.Sprint(os.Getenv("DB_USER")) + " "
	dbInfo += "password=" + fmt.Sprint(os.Getenv("DB_PASS")) + " "
	dbInfo += "host=" + fmt.Sprint(os.Getenv("DB_HOST")) + " "
	dbInfo += "port=" + fmt.Sprint(os.Getenv("DB_PORT")) + " "
	dbInfo += "dbname=" + fmt.Sprint(os.Getenv("DB_NAME")) + " "
	dbInfo += "sslmode=disable"

	log.Println("Connecting to Database . . .")
	log.Println(dbInfo)

	db, err := sqlx.Open("postgres", dbInfo)
	if err != nil {
		log.Println("Failed to connect database", err)
	}

	log.Println("Database Successfully Connected!")

	return db
}
