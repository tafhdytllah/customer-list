package entity

type Nationality struct {
	ID   uint   `gorm:"column:nationality_id"`
	Name string `gorm:"column:nationality_name"`
	Code string `gorm:"column:nationality_code"`
}
