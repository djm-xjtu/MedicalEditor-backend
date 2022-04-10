package entities

type PatientInfo struct {
	PatientId  int    `gorm:"column:patient_id"`
	Department string `gorm:"column:department"`
	Name       string `gorm:"column:name"`
	IdNumber   string `gorm:"column:id_number"`
}
