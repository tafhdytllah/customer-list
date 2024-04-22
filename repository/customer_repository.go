package repository

import (
	"github.com/tafhdytllah/customer-list/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindCustomerById(ID uint) (entity.Customer, error)

	CreateCustomer(customer *entity.Customer) (uint, error)

	CreateFamilyList(familyList *[]entity.FamilyList) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *customerRepository {
	return &customerRepository{db}
}

// FIND CUSTOMER BY ID
func (r *customerRepository) FindCustomerById(ID uint) (entity.Customer, error) {
	var customer entity.Customer

	err := r.db.Preload("FamilyList").Preload("Nationality").Where("customer.cst_id = ?", ID).First(&customer).Error

	return customer, err
}

// CREATE CUSTOMER
func (r *customerRepository) CreateCustomer(customer *entity.Customer) (uint, error) {
	err := r.db.Create(&customer).Error

	return customer.ID, err
}

// CREATE FAMILY LIST
func (r *customerRepository) CreateFamilyList(familyList *[]entity.FamilyList) error {
	err := r.db.Create(&familyList).Error

	return err
}
