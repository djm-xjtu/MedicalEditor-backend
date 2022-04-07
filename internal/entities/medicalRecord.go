package entities

type MedicalRecord struct {
	PatientId  int    `gorm:"column:patient_id"`
	RecordType string `gorm:"column:record_type"`
	Record     string `gorm:"column:record"`
}
