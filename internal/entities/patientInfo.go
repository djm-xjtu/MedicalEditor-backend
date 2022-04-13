package entities

type PatientInfo struct {
	PatientId   string `gorm:"column:patient_id"`
	Department  string `gorm:"column:department"`
	PatientName string `gorm:"column:patient_name"`
	IdNumber    string `gorm:"column:id_number"`
}
