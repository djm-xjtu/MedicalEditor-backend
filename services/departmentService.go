package services

import (
	"editor-backend/database"
	"editor-backend/entities"
)

func GetDepartments() ([]entities.Department, error) {
	var departments []entities.Department
	if err := database.DB.Find(&departments).Error; err != nil {
		return nil, err
	}

	return departments, nil
}
