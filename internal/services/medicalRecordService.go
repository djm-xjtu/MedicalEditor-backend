package services

import (
	"editor-backend/internal/database"
	"editor-backend/internal/entities"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func UpdateMedicalRecord(record string, recordNo int) error {
	db := database.DB
	log.Printf("record_no: %d", recordNo)
	if err := db.Model(&entities.MedicalRecord{}).Where("record_no = ?", recordNo).Update("record", record).Error; err != nil {
		return err
	}

	return nil
}

func GetMedicalRecord(patientCdno string, templateName string) (string, bool, error) {
	db := database.DB
	medicalRecord := entities.MedicalRecord{
		PatientCdno: patientCdno,
		RecordType:  templateName,
	}

	if err := db.Where(&medicalRecord).First(&medicalRecord).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("[patientId: %s TemplateName: %s] record not found, try to find template\n", patientCdno, templateName)
			medicalRecordTemplate := entities.MedicalRecordTemplate{
				TemplateName: templateName,
			}

			if err := db.Where(&medicalRecordTemplate).First(&medicalRecordTemplate).Error; err != nil {
				return "", false, err
			}

			return medicalRecordTemplate.Template, false, nil
		}

		return "", false, err
	}

	return medicalRecord.Record, true, nil
}

func GetMedicalRecords(patientCdno string) ([]entities.MedicalRecord, error) {
	db := database.DB
	var medicalRecords []entities.MedicalRecord

	if err := db.Where("patient_cdno = ?", patientCdno).Find(&medicalRecords).Error; err != nil {
		return nil, err
	}

	return medicalRecords, nil
}

func InsertMedicalRecord(patientCdno, recordType, record string) error {
	medicalRecord := entities.MedicalRecord{
		PatientCdno: patientCdno,
		RecordType:  recordType,
		Record:      record,
	}

	err := entities.AddMedicalRecord(medicalRecord)
	if err != nil {
		return err
	}

	return nil
}
