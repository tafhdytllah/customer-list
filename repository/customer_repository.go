package repository

import (
	"github.com/tafhdytllah/customer-list/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindCustomerById(ID uint) (entity.Customer, error)

	CreateCustomer(customer *entity.Customer) (uint, error)

	CreateFamilyList(familyList *[]entity.FamilyList) error

	UpdateCustomer(customer *entity.Customer) (*entity.Customer, error)

	DeleteFamilyByCustomerId(ID uint) error

	UpdateFamily(family *[]entity.FamilyList) (*[]entity.FamilyList, error)

	CheckCustomerById(ID uint) error

	CheckFamilyById(ID uint) error

	DeleteFamilyByCustomerIdAndFamilyId(customerID uint, familyID uint) error
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

// UPDATE CUSTOMER
func (r *customerRepository) UpdateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	err := r.db.Save(&customer).Error

	return customer, err
}

// DELETE FAMILY BY CUSTOMER ID
func (r *customerRepository) DeleteFamilyByCustomerId(ID uint) error {

	var family entity.FamilyList

	err := r.db.Where("cst_id = ?", ID).Delete(&family).Error

	return err
}

// UPDATE FAMILY
func (r *customerRepository) UpdateFamily(family *[]entity.FamilyList) (*[]entity.FamilyList, error) {
	err := r.db.Save(&family).Error

	return family, err
}

// CHECK CUSTOMER BY ID
func (r *customerRepository) CheckCustomerById(ID uint) error {
	var customer entity.Customer

	err := r.db.First(&customer, ID).Error

	return err
}

// CHECK FAMILY BY ID
func (r *customerRepository) CheckFamilyById(ID uint) error {
	var family entity.FamilyList

	err := r.db.First(&family, ID).Error

	return err
}

// DELETE FAMILY BY CUSTOMER ID AND FAMILY ID
func (r *customerRepository) DeleteFamilyByCustomerIdAndFamilyId(customerID uint, familyID uint) error {
	var family entity.FamilyList

	err := r.db.Where("cst_id = ?", customerID).Where("fl_id = ?", familyID).Delete(&family).Error

	return err
}
