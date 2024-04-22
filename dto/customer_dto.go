package dto

type CustomerResponse struct {
	FamilyList      []FamilyList `json:"keluarga"`
	Name            string       `json:"nama"`
	Dob             string       `json:"tanggal_lahir"`
	Phone           string       `json:"telepon"`
	NationalityName string       `json:"kewarganegaraan"`
	Email           string       `json:"email"`
}

type CustomerRequest struct {
	Name          string       `json:"nama"`
	Dob           string       `json:"tanggal_lahir"`
	NationalityID uint         `json:"kewarganegaraan"`
	Phone         string       `json:"telepon"`
	Email         string       `json:"email"`
	FamilyList    []FamilyList `json:"keluarga"`
}

type FamilyList struct {
	Relation string `json:"hubungan"`
	Name     string `json:"nama"`
	Dob      string `json:"tanggal_lahir"`
}
