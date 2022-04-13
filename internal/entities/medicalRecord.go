package entities

import "editor-backend/internal/database"

type MedicalRecord struct {
	PatientCdno string `gorm:"column:patient_cdno"`
	RecordType  string `gorm:"column:record_type"`
	Record      string `gorm:"column:record"`
	RecordNo    int    `gorm:"column:record_no;primaryKey;autoIncrement:true"`
}

func AddMedicalRecord(record MedicalRecord) error {
	db := database.DB
	if err := db.Create(&record).Error; err != nil {
		return err
	}

	return nil
}
