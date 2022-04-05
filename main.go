package main

import (
	"editor-backend/database"
	"editor-backend/routers"
	"log"
)



func main() {
	err := database.InitDB()
	if err != nil {
		log.Panic(err)
	}

	log.Println("DB connected succes")

	r := routers.InitRouter()
	r.Run(":8000")
}

