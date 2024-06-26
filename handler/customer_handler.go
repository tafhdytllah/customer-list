package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tafhdytllah/customer-list/dto"
	"github.com/tafhdytllah/customer-list/entity"
	errorhandler "github.com/tafhdytllah/customer-list/errorHandler"
	"github.com/tafhdytllah/customer-list/helper"
	"github.com/tafhdytllah/customer-list/service"
)

type CustomerHandler interface {
	FindCustomerById(res http.ResponseWriter, req *http.Request)

	CreateCustomer(res http.ResponseWriter, req *http.Request)

	UpdateCustomerById(res http.ResponseWriter, req *http.Request)

	DeleteFamilyById(res http.ResponseWriter, req *http.Request)
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
		errorhandler.HandlerError(res, &errorhandler.BadRequestError{
			Message: "customer id is invalid",
		})
		return
	}

	customer, err1 := h.service.FindCustomerById(uint(ID))
	if err1 != nil {
		errorhandler.HandlerError(res, &errorhandler.NotFoundError{
			Message: err1.Error(),
		})
		return
	}

	result := convertToCustomerResponse(customer, &customer.FamilyList)

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

// CREATE CUSTOMER
func (h *customerHandler) CreateCustomer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var customerRequest dto.CustomerRequest

	err := json.NewDecoder(req.Body).Decode(&customerRequest)
	if err != nil {
		errorhandler.HandlerError(res, &errorhandler.BadRequestError{
			Message: "request is not valid",
		})
		return
	}

	err1 := h.service.Validation(&customerRequest)
	if err1 != nil {
		errorhandler.HandlerError(res, &errorhandler.BadRequestError{
			Message: err1.Error(),
		})
		return
	}

	err2 := h.service.CreateCustomer(&customerRequest)
	if err2 != nil {
		errorhandler.HandlerError(res, &errorhandler.InternalServerError{
			Message: err2.Error(),
		})
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(customerRequest)
}

// UPDATE CUSTOMER BY ID
func (h *customerHandler) UpdateCustomerById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var customerRequest dto.CustomerRequest

	err := json.NewDecoder(req.Body).Decode(&customerRequest)
	if err != nil {
		errorhandler.HandlerError(res, &errorhandler.BadRequestError{
			Message: "request is not valid",
		})
		return
	}

	vars := mux.Vars(req)

	customerID, err := strconv.ParseInt(vars["customer_id"], 10, 32)
	if err != nil {
		errorhandler.HandlerError(res, &errorhandler.BadRequestError{
			Message: "customer id is invalid",
		})
		return
	}

	err1 := h.service.Validation(&customerRequest)
	if err1 != nil {
		errorhandler.HandlerError(res, &errorhandler.BadRequestError{
			Message: err1.Error(),
		})
		return
	}

	updatedCustomer, updatedFamily, err2 := h.service.UpdateCustomerById(uint(customerID), &customerRequest)
	if err2 != nil {
		errorhandler.HandlerError(res, &errorhandler.BadRequestError{
			Message: err2.Error(),
		})
		return
	}

	result := convertToCustomerResponse(updatedCustomer, updatedFamily)

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)

}

// DELETE FAMILY BY ID
func (h *customerHandler) DeleteFamilyById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	vars := mux.Vars(req)

	customerID, err := strconv.ParseInt(vars["customer_id"], 10, 32)
	if err != nil {
		errorhandler.HandlerError(res, &errorhandler.BadRequestError{
			Message: "customer id is invalid",
		})
		return
	}

	familyID, err1 := strconv.ParseInt(vars["family_id"], 10, 32)
	if err1 != nil {
		errorhandler.HandlerError(res, &errorhandler.BadRequestError{
			Message: "family id is invalid",
		})
		return
	}

	err2 := h.service.CheckCustomerById(uint(customerID))
	if err2 != nil {
		errorhandler.HandlerError(res, &errorhandler.NotFoundError{
			Message: err2.Error(),
		})
		return
	}

	err3 := h.service.CheckFamilyById(uint(familyID))
	if err3 != nil {
		errorhandler.HandlerError(res, &errorhandler.NotFoundError{
			Message: err3.Error(),
		})
		return
	}

	err4 := h.service.DeleteFamilyByCustomerIdAndFamilyId(uint(customerID), uint(familyID))
	if err4 != nil {
		errorhandler.HandlerError(res, &errorhandler.InternalServerError{
			Message: err4.Error(),
		})
		return
	}

	result := helper.Response(dto.ResponseParams{
		Message: "Success delete family",
	})

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

// CONVERT TO CUSTOMER RESPONSE
func convertToCustomerResponse(customer *entity.Customer, fam *[]entity.FamilyList) dto.CustomerResponse {
	familyList := make([]dto.FamilyList, len(*fam))
	for i, family := range *fam {
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
