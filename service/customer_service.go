package service

import (
	"strings"

	"github.com/tafhdytllah/customer-list/dto"
	"github.com/tafhdytllah/customer-list/entity"
	errorhandler "github.com/tafhdytllah/customer-list/errorHandler"
	"github.com/tafhdytllah/customer-list/repository"
	"gorm.io/gorm"
)

type CustomerService interface {
	FindCustomerById(ID uint) (*entity.Customer, error)

	Validation(request *dto.CustomerRequest) error

	CreateCustomer(request *dto.CustomerRequest) error

	UpdateCustomerById(customerID uint, request *dto.CustomerRequest) (*entity.Customer, *[]entity.FamilyList, error)

	CheckCustomerById(ID uint) error

	CheckFamilyById(ID uint) error

	DeleteFamilyByCustomerIdAndFamilyId(customerID uint, familyID uint) error
}

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(r repository.CustomerRepository) *customerService {
	return &customerService{r}
}

// FIND CUSTOMER BY ID
func (s *customerService) FindCustomerById(ID uint) (*entity.Customer, error) {
	customer, err := s.repository.FindCustomerById(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &errorhandler.NotFoundError{Message: "article not found"}
		}

		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	return customer, nil
}

// CREATE CUSTOMER
func (s *customerService) CreateCustomer(request *dto.CustomerRequest) error {
	customer := entity.Customer{
		NationalityID: request.NationalityID,
		CstName:       request.Name,
		CstDob:        request.Dob,
		CstPhone:      request.Phone,
		CstEmail:      request.Email,
	}

	idCustomer, err := s.repository.CreateCustomer(&customer)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	var familyList []entity.FamilyList

	for _, f := range request.FamilyList {
		familyList = append(familyList, entity.FamilyList{
			CustomerID: idCustomer,
			Relation:   f.Relation,
			Name:       f.Name,
			Dob:        f.Dob,
		})
	}

	if err1 := s.repository.CreateFamilyList(&familyList); err1 != nil {
		return &errorhandler.InternalServerError{Message: err1.Error()}
	}

	return nil
}

// UPDATE CUSTOMER BY ID
func (s *customerService) UpdateCustomerById(customerID uint, request *dto.CustomerRequest) (*entity.Customer, *[]entity.FamilyList, error) {

	customer, err := s.repository.FindCustomerById(customerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, &errorhandler.NotFoundError{Message: "customer not found"}
		}

		return nil, nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	newName := strings.TrimSpace(request.Name)
	newDob := strings.TrimSpace(request.Dob)
	newNationalityId := uint(request.NationalityID)
	newPhone := strings.TrimSpace(request.Phone)
	newEmail := strings.TrimSpace(request.Email)

	customer.CstName = newName
	customer.CstDob = newDob
	customer.NationalityID = newNationalityId
	customer.CstPhone = newPhone
	customer.CstEmail = newEmail

	var familyList []entity.FamilyList
	for _, f := range request.FamilyList {
		familyList = append(familyList, entity.FamilyList{
			CustomerID: customerID,
			Relation:   f.Relation,
			Name:       f.Name,
			Dob:        f.Dob,
		})
	}

	updatedCustomer, err1 := s.repository.UpdateCustomer(customer)
	if err1 != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err1.Error()}
	}

	if err2 := s.repository.DeleteFamilyByCustomerId(customerID); err2 != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err2.Error()}
	}

	updatedFamily, err3 := s.repository.UpdateFamily(&familyList)
	if err3 != nil {
		return nil, nil, &errorhandler.InternalServerError{Message: err3.Error()}
	}

	return updatedCustomer, updatedFamily, nil
}

// CHECK CUSTOMER BY ID
func (s *customerService) CheckCustomerById(ID uint) error {
	if err := s.repository.CheckCustomerById(ID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &errorhandler.NotFoundError{Message: "customer not found"}
		}

		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

// CHECK FAMILY BY ID
func (s *customerService) CheckFamilyById(ID uint) error {
	if err := s.repository.CheckFamilyById(ID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &errorhandler.NotFoundError{Message: "family not found"}
		}

		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

// DELETE FAMILY BY CUSTOMER ID AND FAMILY ID
func (s *customerService) DeleteFamilyByCustomerIdAndFamilyId(customerID uint, familyID uint) error {
	if err := s.repository.DeleteFamilyByCustomerIdAndFamilyId(customerID, familyID); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

// REQUEST BODY VALIDATION
func (s *customerService) Validation(request *dto.CustomerRequest) error {
	if request == nil {
		return &errorhandler.BadRequestError{Message: "request body is required"}
	}
	if request.Name == "" {
		return &errorhandler.BadRequestError{Message: "nama is required"}
	}
	if request.Dob == "" {
		return &errorhandler.BadRequestError{Message: "tanggal lahir is required"}
	}
	if request.NationalityID == 0 {
		return &errorhandler.BadRequestError{Message: "kewarganegaraan is required"}
	}
	if request.Phone == "" {
		return &errorhandler.BadRequestError{Message: "telepon is required"}
	}
	if request.Email == "" {
		return &errorhandler.BadRequestError{Message: "email is required"}
	}
	return nil
}
