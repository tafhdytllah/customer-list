package handler

import "github.com/tafhdytllah/customer-list/entity"

type CustomerHandler interface {
	FindCustomerById(ID uint) (entity.Customer, error)
}
