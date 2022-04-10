package services

import (
	"editor-backend/internal/database"
	"editor-backend/internal/entities"
	"encoding/xml"
	"fmt"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Xml struct {
	Controls []Control `xml:"controls"`
	Html     string    `xml:"html"`
}

type Control struct {
	Id    string `xml:"id"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
}

func InsertMedicalRecordTemplate(recordType, template string) error {
	medicalRecordTemplate := entities.MedicalRecordTemplate{
		RecordType: recordType,
		Template:   template,
	}

	var templateXml Xml
	xml.Unmarshal([]byte(template), &templateXml)
	fmt.Printf("xml:\n%+v\n", templateXml)

	enEscapeUrl, _ := url.QueryUnescape(templateXml.Html)
	fmt.Println("解码:", enEscapeUrl)

	db := database.DB
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "record_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"template"}),
	}).Create(&medicalRecordTemplate).Error; err != nil {
		return err
	}

	return nil
}

func GetMedicalRecordTemplate(recordType string) (string, error) {
	medicalRecordTemplate := entities.MedicalRecordTemplate{
		RecordType: recordType,
	}

	db := database.DB
	if err := db.Where(&medicalRecordTemplate).First(&medicalRecordTemplate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("[recordType: %s] template not found\n", recordType)
		}

		return "", err
	}

	return medicalRecordTemplate.Template, nil
}
