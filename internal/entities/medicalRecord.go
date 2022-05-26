package entities

import (
	"editor-backend/internal/database"
	"log"
)

type MedicalRecord struct {
	Mzghxh      string `gorm:"column:mzghxh;primaryKey"`
	PatientCdno string `gorm:"column:patient_cdno"`
	RecordType  string `gorm:"column:record_type"`
	Record      string `gorm:"column:record"`
	Xm          string `gorm:"column:xm"`
	Xb          string `gorm:"column:xb"`
	Cssj        string `gorm:"column:cssj"`
	Jzks        string `gorm:"column:jzks"`
	Tel         string `gorm:"column:tel"`
	UpdateBy    string `gorm:"column:update_by"`
	UpdateTime  string `gorm:"column:update_time"`
	ChangeLog   string `gorm:"column:change_log"`
	RecordXml   string `gorm:"column:record_xml"`
}

func AddMedicalRecord(record MedicalRecord) error {
	db := database.DB
	if err := db.Create(&record).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
