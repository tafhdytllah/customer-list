package service

import (
	"errors"

	"github.com/tafhdytllah/customer-list/dto"
	"github.com/tafhdytllah/customer-list/entity"
	"github.com/tafhdytllah/customer-list/repository"
	"gorm.io/gorm"
)

type CustomerService interface {
	FindCustomerById(ID uint) (entity.Customer, error)

	Validation(request *dto.CustomerRequest) error

	CreateCustomer(request *dto.CustomerRequest) error
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
