package services

import (
	"editor-backend/internal/database"
	"editor-backend/internal/entities"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// type Xml struct {
// 	Controls []Control `xml:"controls"`
// 	Html     string    `xml:"html"`
// }

// type Control struct {
// 	Id    string `xml:"id"`
// 	Type  string `xml:"type"`
// 	Value string `xml:"value"`
// }

type RecordTemplate struct {
	TemplateName string
	Template     string
	TemplateNo   string
	Department   string
	UsageType    string
	Creater      string
	CreationTime string
	Luruma       string
	Status       string
	Comment      string
}

func InsertMedicalRecordTemplate(recordTemplate RecordTemplate) error {
	medicalRecordTemplate := entities.MedicalRecordTemplate{
		TemplateName: recordTemplate.TemplateName,
		Template:     recordTemplate.Template,
		TemplateNo:   recordTemplate.TemplateNo,
		Department:   recordTemplate.Department,
		UsageType:    recordTemplate.UsageType,
		Creater:      recordTemplate.Creater,
		CreationTime: recordTemplate.CreationTime,
		Luruma:       recordTemplate.Luruma,
		Status:       recordTemplate.Status,
		Comment:      recordTemplate.Comment,
	}

	// var templateXml Xml
	// xml.Unmarshal([]byte(template), &templateXml)
	// fmt.Printf("xml:\n%+v\n", templateXml)

	// enEscapeUrl, _ := url.QueryUnescape(templateXml.Html)
	// fmt.Println("解码:", enEscapeUrl)

	db := database.DB
	// if err := db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "record_type"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"template"}),
	// }).Create(&medicalRecordTemplate).Error; err != nil {
	// 	return err
	// }
	if err := db.Create(&medicalRecordTemplate).Error; err != nil {
		return err
	}

	return nil
}

func GetMedicalRecordTemplate(templateName string) (string, error) {
	medicalRecordTemplate := entities.MedicalRecordTemplate{
		TemplateName: templateName,
	}

	db := database.DB
	if err := db.Where(&medicalRecordTemplate).First(&medicalRecordTemplate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("[recordType: %s] template not found\n", templateName)
		}

		return "", err
	}

	return medicalRecordTemplate.Template, nil
}

func GetMedicalRecordTemplates(creater, department string) ([]entities.MedicalRecordTemplate, error) {
	db := database.DB
	var templates []entities.MedicalRecordTemplate
	log.Printf("creater: %s, department: %s", creater, department)
	if err := db.Where(entities.MedicalRecordTemplate{
		Creater: creater,
		Status:  "正常",
	}).Or(entities.MedicalRecordTemplate{
		Department: department,
		UsageType:  "科室",
		Status:     "正常",
	}).Find(&templates).Error; err != nil {
		return nil, err
	}

	return templates, nil
}
