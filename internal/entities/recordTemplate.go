package entities

type MedicalRecordTemplate struct {
	TemplateName string `gorm:"column:template_name"`
	Template     string `gorm:"column:template"`
	Creater      string `gorm:"creater"`
	Department   string `gorm:"department"`
	UsageType    string `gorm:"usage_type"`
	CreationTime string `gorm:"creation_time"`
	Status       string `gorm:"status"`
	TemplateNo   string `gorm:"template_no"`
	Comment      string `gorm:"comment"`
	Luruma       string `gorm:"luruma"`
}
