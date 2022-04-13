package services

import (
	"editor-backend/internal/entities"
	"fmt"
	"log"
)

var OldRecordList = []string{"历史病历1", "历史病历2", "历史病历3"}
var RecordTemplateList = []string{"模板病历1", "模板病历2", "模板病历3"}

type PatientInfo struct {
	Mzghxh string
	Xm     string
	Xb     string
	Cssj   string
	Nl     string
	Kb     string
	Cdno   string
	Sfzhm  string
}

type TreeData struct {
	Title       string     `json:"title"`
	Expand      bool       `json:"expand"`
	Contextmenu bool       `json:"contextmenu"`
	Children    []TreeData `json:"children"`
	Xml         string     `json:"xml"`
	RecordType  string     `json:"recordType"`
	RecordNo    int        `json:"recordNo"`
	TemplateNo  string     `json:"templateNo"`
	IsRecord    bool       `json:"isRecord"`
}

type Data struct {
	PatientInfo PatientInfo
	MenuData    []TreeData
}

func GetMenuInfo() ([]TreeData, error) {
	oldRecordItem := TreeData{
		Title:       "历史病历",
		Expand:      false,
		Contextmenu: false,
	}

	for _, str := range OldRecordList {
		oldRecordItem.Children = append(oldRecordItem.Children, TreeData{
			Title:       str,
			Expand:      false,
			Contextmenu: false,
		})
	}

	recordTemplateItem := TreeData{
		Title:       "模板病历",
		Expand:      false,
		Contextmenu: false,
	}

	for _, str := range RecordTemplateList {
		recordTemplateItem.Children = append(recordTemplateItem.Children, TreeData{
			Title:       str,
			Expand:      false,
			Contextmenu: false,
		})
	}
	treeDatas := []TreeData{oldRecordItem, recordTemplateItem}
	return treeDatas, nil
}

func GetData(mzghxh string) (Data, error) {
	m, err := entities.GetMzghDfsy(mzghxh)
	if err != nil {
		return Data{}, err
	}

	records, err := GetMedicalRecords(m.Cdno)
	if err != nil {
		return Data{}, err
	}

	templates, err := GetMedicalRecordTemplates(m.Df, m.Kb)
	if err != nil {
		return Data{}, err
	}

	patientInfo := PatientInfo{
		Mzghxh: m.Mzghxh,
		Xm:     m.Xm,
		Xb:     m.Xb,
		Cssj:   m.Cssj,
		Nl:     m.Nl,
		Kb:     m.Kb,
		Cdno:   m.Cdno,
		Sfzhm:  m.Sfzhm,
	}

	historyRecordItem := TreeData{
		Title:       "历史病历",
		Expand:      false,
		Contextmenu: false,
	}

	for _, r := range records {
		historyRecordItem.Children = append(historyRecordItem.Children, TreeData{
			Title:    fmt.Sprintf("%s - %d", r.RecordType, r.RecordNo),
			Xml:      r.Record,
			RecordNo: r.RecordNo,
			IsRecord: true,
			RecordType: r.RecordType,
		})
		log.Printf("XmlNo: %s", string(rune(r.RecordNo)))
	}

	recordTemplateItem := TreeData{
		Title:       "病历模板",
		Expand:      false,
		Contextmenu: false,
	}

	for _, t := range templates {
		recordTemplateItem.Children = append(recordTemplateItem.Children, TreeData{
			Title:      t.TemplateName,
			Xml:        t.Template,
			TemplateNo: t.TemplateNo,
			IsRecord:   false,
			RecordType: t.TemplateName,
		})
	}

	return Data{
		PatientInfo: patientInfo,
		MenuData:    []TreeData{historyRecordItem, recordTemplateItem},
	}, nil
}
