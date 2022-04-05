package services

import (
	"editor-backend/database"
	"editor-backend/entities"
	"log"
)

func GetDepartments() ([]entities.Department, error) {
	var departments []entities.Department
	log.Println(&departments)
	if err := database.DB.Find(&departments).Error; err != nil {
		log.Println("2")
		return nil, err
	}

	log.Printf("departments: %v", departments)

	return departments, nil
}
