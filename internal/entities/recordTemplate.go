package entities

type MedicalRecordTemplate struct {
	RecordType string `gorm:"column:record_type"`
	Template   string `gorm:"column:template"`
}
