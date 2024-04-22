package service

import (
	"errors"
	"strings"

	"github.com/tafhdytllah/customer-list/dto"
	"github.com/tafhdytllah/customer-list/entity"
	"github.com/tafhdytllah/customer-list/repository"
	"gorm.io/gorm"
)

type CustomerService interface {
	FindCustomerById(ID uint) (entity.Customer, error)

	Validation(request *dto.CustomerRequest) error

	CreateCustomer(request *dto.CustomerRequest) error

	UpdateCustomerById(customerID uint, request *dto.CustomerRequest) (*entity.Customer, *[]entity.FamilyList, error)
}

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(r repository.CustomerRepository) *customerService {
	return &customerService{r}
}

// FIND CUSTOMER BY ID
func (s *customerService) FindCustomerById(ID uint) (entity.Customer, error) {
	customer, err := s.repository.FindCustomerById(ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Customer{}, errors.New("article not found")
		}

		return entity.Customer{}, errors.New(err.Error())
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
		return errors.New("create customer failed")
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
		return errors.New("create family failed")
	}

	return nil
}

// UPDATE CUSTOMER BY ID
func (s *customerService) UpdateCustomerById(customerID uint, request *dto.CustomerRequest) (*entity.Customer, *[]entity.FamilyList, error) {

	customer, err := s.repository.FindCustomerById(customerID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entity.Customer{}, &[]entity.FamilyList{}, errors.New("customer not found")
		}

		return &entity.Customer{}, &[]entity.FamilyList{}, errors.New(err.Error())
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

	updatedCustomer, err1 := s.repository.UpdateCustomer(&customer)
	if err1 != nil {
		return &entity.Customer{}, &[]entity.FamilyList{}, errors.New("update customer failed")
	}

	err2 := s.repository.DeleteFamilyByCustomerId(customerID)
	if err2 != nil {
		return &entity.Customer{}, &[]entity.FamilyList{}, errors.New("family not found")
	}

	updatedFamily, err3 := s.repository.UpdateFamily(&familyList)
	if err3 != nil {
		return &entity.Customer{}, &[]entity.FamilyList{}, errors.New("update family failed")
	}

	return updatedCustomer, updatedFamily, nil
}

// REQUEST BODY VALIDATION
func (s *customerService) Validation(request *dto.CustomerRequest) error {
	if request == nil {
		err1 := errors.New("request body is required")
		return err1
	}
	if request.Name == "" {
		err2 := errors.New("nama is required")
		return err2
	}
	if request.Dob == "" {
		err3 := errors.New("tanggal lahir is required")
		return err3
	}
	if request.NationalityID == 0 {
		err4 := errors.New("kewarganegaraan is required")
		return err4
	}
	if request.Phone == "" {
		err5 := errors.New("telepon is required")
		return err5
	}
	if request.Email == "" {
		err6 := errors.New("email is required")
		return err6
	}
	return nil
}
