package entity

type FamilyList struct {
	ID         uint   `gorm:"primaryKey;column:fl_id"`
	CustomerID uint   `gorm:"foreignKey;column:cst_id"`
	Relation   string `gorm:"column:fl_relation"`
	Name       string `gorm:"column:fl_name"`
	Dob        string `gorm:"column:fl_dob"`
}
