package repository

import (
	"github.com/tafhdytllah/customer-list/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindCustomerById(ID uint) (entity.Customer, error)
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
