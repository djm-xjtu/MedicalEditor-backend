package services

import (
	"editor-backend/database"
	"editor-backend/entities"
	"log"
)

func UpdateOrInsertMedicalRecord(patientId int, recordType, record string) error {
	medicalRecord := entities.MedicalRecord{
		PatientId:  patientId,
		RecordType: recordType,
		Record:     record,
	}

	db := database.DB
	log.Println("record")
	log.Println(medicalRecord)
	if err := db.Save(&medicalRecord).Error; err != nil {
		log.Println(5)
		return err
	}
	log.Println(6)
	return nil
}

func GetMedicalRecord(patientId int, recordType string) (string, error) {
	db := database.DB
	medicalRecord := entities.MedicalRecord{
		PatientId:  patientId,
		RecordType: recordType,
	}

	if err := db.Where(&medicalRecord).First(&medicalRecord).Error; err != nil {
		return "", err
	}

	return medicalRecord.Record, nil
}
