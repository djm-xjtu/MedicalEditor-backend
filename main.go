package main

import (
	"editor-backend/internal/database"
	"editor-backend/internal/routers"
	"log"
)


func main() {
	err := database.InitMssqlDB()
	if err != nil {
		log.Panic(err)
	}

	log.Println("DB connected success")
	
	r := routers.InitRouter()
	r.Run(":8000")
}

