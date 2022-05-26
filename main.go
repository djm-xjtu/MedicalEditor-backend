package main

import (
	"editor-backend/internal/database"
	"editor-backend/internal/routers"
	"log"
)

// @title 接口文档
// @version 1.0
// @description ziannchen-emr-api
func main() {
	err := database.InitMssqlDB()
	if err != nil {
		log.Panic(err)
	}

	log.Println("DB connected succes")
	
	r := routers.InitRouter()
	r.Run(":8000")
}

