package service

import (
	"errors"

	"github.com/tafhdytllah/customer-list/entity"
	"github.com/tafhdytllah/customer-list/repository"
	"gorm.io/gorm"
)

type CustomerService interface {
	FindCustomerById(ID uint) (entity.Customer, error)
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
