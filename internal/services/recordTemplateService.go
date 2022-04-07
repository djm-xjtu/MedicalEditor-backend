package services

import (
	"editor-backend/internal/database"
	"editor-backend/internal/entities"

	"gorm.io/gorm/clause"
)

func InsertMedicalRecordTemplate(recordType, template string) error {
	medicalRecordTemplate := entities.MedicalRecordTemplate {
		RecordType: recordType,
		Template: template,
	}
	
	db := database.DB
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "record_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"template"}),
	}).Create(&medicalRecordTemplate).Error; err != nil {
		return err
	}

	return nil
}