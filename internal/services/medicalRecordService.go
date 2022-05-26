package services

import (
	"editor-backend/internal/database"
	"editor-backend/internal/entities"
	"fmt"

	"gorm.io/gorm"
)

func UpdateMedicalRecord(record, mzghxh, updateBy, updateTime, changeLog, recordXml string) error {
	db := database.DB
	medicalRecord := &entities.MedicalRecord{
		Mzghxh:     mzghxh,
		Record:     record,
		UpdateBy:   updateBy,
		UpdateTime: updateTime,
		ChangeLog:  changeLog,
		RecordXml:  recordXml,
	}
	if err := db.Model(&medicalRecord).Updates(medicalRecord).Error; err != nil {
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

func InsertMedicalRecord(mzghxh, patientCdno, recordType, record, xm, xb, cssj, jzks, tel, updateBy, updateTime, changeLog, recordXml string) error {
	medicalRecord := entities.MedicalRecord{
		Mzghxh:      mzghxh,
		PatientCdno: patientCdno,
		RecordType:  recordType,
		Record:      record,
		Xm:          xm,
		Xb:          xb,
		Cssj:        cssj,
		Jzks:        jzks,
		Tel:         tel,
		UpdateBy:    updateBy,
		UpdateTime:  updateTime,
		ChangeLog:   changeLog,
		RecordXml:   recordXml,
	}

	err := entities.AddMedicalRecord(medicalRecord)
	if err != nil {
		return err
	}

	return nil
}
