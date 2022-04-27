package services

import (
	"editor-backend/internal/entities"
	"log"
)

var OldRecordList = []string{"历史病历1", "历史病历2", "历史病历3"}
var RecordTemplateList = []string{"模板病历1", "模板病历2", "模板病历3"}

type PatientInfo struct {
	Mzghxh string `json:"mzghxh"`
	Xm     string `json:"xm"`
	Xb     string `json:"xb"`
	Cssj   string `json:"cssj"`
	Nl     string `json:"nl"`
	Kb     string `json:"kb"`
	Cdno   string `json:"cdno"`
	Sfzhm  string `json:"sfzhm"`
}

type TreeData struct {
	Label        string     `json:"label"`
	Expand       bool       `json:"expand"`
	Contextmenu  bool       `json:"contextmenu"`
	Children     []TreeData `json:"children"`
	Xml          string     `json:"xml"`
	RecordType   string     `json:"recordType"`
	RecordNo     int        `json:"recordNo"`
	Template     string     `json:"template"`
	TemplateNo   string     `json:"templateNo"`
	TemplateName string     `json:"templateName"`
	IsRecord     bool       `json:"isRecord"`
	Record       string     `json:"record"`
}

type RecordTemplateData struct {
	TemplateNo   string `json:"templateNo"`
	TemplateName string `json:"templateName"`
	Template     string `json:"template"`
}

type HistoryRecordData struct {
	RecordNo   int    `json:"recordNo"`
	RecordType string `json:"recordType"`
	Record     string `json:"record"`
}

type Data struct {
	PatientInfo        PatientInfo          `json:"patientInfo"`
	HistoryRecordData  []HistoryRecordData  `json:"historyRecords"`
	RecordTemplateData []RecordTemplateData `json:"availableTemplates"`
}

func GetMenuInfo() ([]TreeData, error) {
	oldRecordItem := TreeData{
		Label:       "历史病历",
		Expand:      false,
		Contextmenu: false,
	}

	for _, str := range OldRecordList {
		oldRecordItem.Children = append(oldRecordItem.Children, TreeData{
			Label:       str,
			Expand:      false,
			Contextmenu: false,
		})
	}

	recordTemplateItem := TreeData{
		Label:       "模板病历",
		Expand:      false,
		Contextmenu: false,
	}

	for _, str := range RecordTemplateList {
		recordTemplateItem.Children = append(recordTemplateItem.Children, TreeData{
			Label:       str,
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

	historyRecordItem := []HistoryRecordData{}

	for _, r := range records {
		historyRecordItem = append(historyRecordItem, HistoryRecordData{
			RecordNo:   r.RecordNo,
			RecordType: r.RecordType,
			Record:     r.Record,
		})
		log.Printf("XmlNo: %s", string(rune(r.RecordNo)))
	}

	recordTemplateItem := []RecordTemplateData{}

	for _, t := range templates {
		recordTemplateItem = append(recordTemplateItem, RecordTemplateData{
			TemplateNo:   t.TemplateNo,
			TemplateName: t.TemplateName,
			Template:     t.Template,
		})
	}

	return Data{
		PatientInfo:        patientInfo,
		HistoryRecordData:  historyRecordItem,
		RecordTemplateData: recordTemplateItem,
	}, nil
}
