package service

import "github.com/tafhdytllah/customer-list/entity"

type CustomerService interface {
	FindCustomerById(ID uint) (entity.Customer, error)
}
