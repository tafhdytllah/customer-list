package repository

import "github.com/tafhdytllah/customer-list/entity"

type CustomerRepository interface {
	FindCustomerById(ID uint) (entity.Customer, error)
}
