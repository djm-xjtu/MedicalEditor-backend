package main

import (
	"editor-backend/database"
	"editor-backend/routers"
	"log"

	_ "github.com/go-sql-driver/mysql"
)



func main() {
	err := database.InitDB()
	if err != nil {
		log.Panic(err)
	}
	log.Println("db ok")
	defer database.Close()

	r := routers.InitRouter()
	r.Run(":8000")
}

