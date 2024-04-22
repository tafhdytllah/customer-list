package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tafhdytllah/customer-list/dto"
	"github.com/tafhdytllah/customer-list/entity"
	errorhandler "github.com/tafhdytllah/customer-list/errorHandler"
	"github.com/tafhdytllah/customer-list/service"
)

type CustomerHandler interface {
	FindCustomerById(res http.ResponseWriter, req *http.Request)

	CreateCustomer(res http.ResponseWriter, req *http.Request)
}

type customerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(s service.CustomerService) *customerHandler {
	return &customerHandler{s}
}

// FIND CUSTOMER BY ID
func (h *customerHandler) FindCustomerById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	vars := mux.Vars(req)

	ID, err := strconv.ParseInt(vars["customer_id"], 10, 32)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorhandler.ServiceError{
			Message: err.Error(),
		})
		return
	}

	customer, err1 := h.service.FindCustomerById(uint(ID))
	if err1 != nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(errorhandler.ServiceError{
			Message: err1.Error(),
		})
		return
	}

	result := convertToCustomerResponse(customer, customer.FamilyList)

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

// CREATE CUSTOMER
func (h *customerHandler) CreateCustomer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var customerRequest dto.CustomerRequest

	err := json.NewDecoder(req.Body).Decode(&customerRequest)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorhandler.ServiceError{
			Message: "Error unmarshalling the request",
		})
		return
	}

	err1 := h.service.Validation(&customerRequest)
	if err1 != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errorhandler.ServiceError{
			Message: err1.Error(),
		})
		return
	}

	err2 := h.service.CreateCustomer(&customerRequest)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errorhandler.ServiceError{
			Message: err2.Error(),
		})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(customerRequest)
}

// CONVERT TO CUSTOMER RESPONSE
func convertToCustomerResponse(customer entity.Customer, fam []entity.FamilyList) dto.CustomerResponse {
	familyList := make([]dto.FamilyList, len(fam))
	for i, family := range fam {
		familyList[i] = dto.FamilyList{
			Relation: family.Relation,
			Name:     family.Name,
			Dob:      family.Dob,
		}
	}

	return dto.CustomerResponse{
		FamilyList:      familyList,
		Name:            customer.CstName,
		Dob:             customer.CstDob,
		Phone:           customer.CstPhone,
		NationalityName: customer.Nationality.Name,
		Email:           customer.CstEmail,
	}
}
