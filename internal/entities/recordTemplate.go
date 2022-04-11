package entities

type MedicalRecordTemplate struct {
	RecordType string `gorm:"column:record_type"`
	Template   string `gorm:"column:template"`
	// Creater    string `gorm:"creater"`
	// Department string `gorm:"department"`
	// CreateTime string `gorm:"create_time"`
	// UsageType  int    `gorm:"usage_type"`
	// Status     string `gorm:"status"`
	// RecordId   string `gorm:"record_id"`
	// Info       string `gorm:"info"`
}
