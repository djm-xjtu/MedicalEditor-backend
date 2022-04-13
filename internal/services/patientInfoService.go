package services

import (
	"editor-backend/internal/database"
	"editor-backend/internal/entities"
	"errors"
	"log"
)

var RecordTypeList = []string{"门诊初诊病历", "门诊复诊病历"}

type PatientMeta struct {
	Id       int
	title    string
	IdNumber string
}

type TreePatientInfo struct {
	Title       string            `json:"title"`
	Expand      bool              `json:"expand"`
	Contextmenu bool              `json:"contextmenu"`
	Children    []TreePatientInfo `json:"children"`
	Id         string              `json:"id"`
	Name        string            `json:"name"`
	Department  string            `json:"department"`
	IdNumber    string            `json:"idNumber"`
}

func GetPatientInfos() ([]TreePatientInfo, error) {
	var patientInfos []entities.PatientInfo
	if err := database.DB.Find(&patientInfos).Error; err != nil {
		return nil, err
	}

	tempTreeMap := make(map[string][]entities.PatientInfo)
	for _, patientInfo := range patientInfos {
		if _, ok := tempTreeMap[patientInfo.Department]; !ok {
			tempTreeMap[patientInfo.Department] = make([]entities.PatientInfo, 0)
		}

		tempTreeMap[patientInfo.Department] = append(tempTreeMap[patientInfo.Department], entities.PatientInfo{
			PatientId:  patientInfo.PatientId,
			PatientName:       patientInfo.PatientName,
			IdNumber:   patientInfo.IdNumber,
			Department: patientInfo.Department,
		})
	}

	treePatientInfos := make([]TreePatientInfo, 0)
	for department, patientInfoList := range tempTreeMap {
		departmentItem := TreePatientInfo{
			Title:       department,
			Expand:      false,
			Contextmenu: false,
			Children:    make([]TreePatientInfo, 0),
		}

		for _, patientInfo := range patientInfoList {
			patientItem := TreePatientInfo{
				Title:       patientInfo.PatientName,
				IdNumber:    patientInfo.IdNumber,
				Expand:      false,
				Contextmenu: false,
				Children:    make([]TreePatientInfo, 0),
			}

			for _, recordType := range RecordTypeList {
				recordItem := TreePatientInfo{
					Title:       recordType,
					Id:          patientInfo.PatientId,
					Name:        patientInfo.PatientName,
					IdNumber:    patientInfo.IdNumber,
					Department:  department,
					Expand:      false,
					Contextmenu: true,
				}
				patientItem.Children = append(patientItem.Children, recordItem)
			}
			departmentItem.Children = append(departmentItem.Children, patientItem)
		}
		treePatientInfos = append(treePatientInfos, departmentItem)
	}

	return treePatientInfos, nil
}

func GetPatientInfoByPatientId(patientId string) ([]TreePatientInfo, error) {
	var patientInfo entities.PatientInfo
	log.Printf("patientId %s", patientId)
	log.Println(11)
	if err := database.DB.Where("patient_id = ?", patientId).Find(&patientInfo).Error; err != nil {
		log.Println(11)
		return nil, err
	}

	if patientId != patientInfo.PatientId {
		log.Printf("+%v",patientInfo)
		return nil, errors.New("no record")
	}
	log.Printf("12")
	patientItem := TreePatientInfo{
		Title:       patientInfo.PatientName,
		IdNumber:    patientInfo.IdNumber,
		Expand:      false,
		Contextmenu: false,
		Children:    make([]TreePatientInfo, 0),
	}

	for _, recordType := range RecordTypeList {
		recordItem := TreePatientInfo{
			Title:       recordType,
			Id:          patientInfo.PatientId,
			IdNumber:    patientInfo.IdNumber,
			Department:  patientInfo.Department,
			Name:        patientInfo.PatientName,
			Expand:      false,
			Contextmenu: true,
		}
		patientItem.Children = append(patientItem.Children, recordItem)
	}

	return []TreePatientInfo{patientItem}, nil
}
