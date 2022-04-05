package entities

type Department struct {
	Id     int    `gorm:"column:id"`
	First  string `gorm:"column:first"`
	Second string `gorm:"column:second"`
	Name   string `gorm:"column:name"`
}

func (Department) TableName() string {
	return "departments"
}
