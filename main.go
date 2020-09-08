package main

import (
	"github.com/alifudin-a/go-echo-api/database"
	"github.com/alifudin-a/go-echo-api/routers"
	_ "github.com/lib/pq"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file!")
	// }

	database.DBConn()
	e := routers.Init()
	e.Logger.Fatal(e.Start(":5000"))
}
