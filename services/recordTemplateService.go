package services

import (
	"editor-backend/database"
	"editor-backend/entities"
)

func InsertMedicalRecordTemplate(recordType, template string) error {
	medicalRecordTemplate := entities.MedicalRecordTemplate {
		RecordType: recordType,
		Template: template,
	}

	if err := database.DB.Create(&medicalRecordTemplate).Error; err != nil {
		return err
	}

	return nil
}