package entity

type Customer struct {
	ID            uint         `gorm:"primaryKey;column:cst_id"`
	NationalityID uint         `gorm:"column:nationality_id"`
	CstName       string       `gorm:"column:cst_name"`
	CstDob        string       `gorm:"column:cst_dob"`
	CstPhone      string       `gorm:"column:cst_phonenum"`
	CstEmail      string       `gorm:"column:cst_email"`
	FamilyList    []FamilyList `json:"FamilyList"`
	Nationality   Nationality  `gorm:"foreignKey:NationalityID" json:"Nationality"`
}
