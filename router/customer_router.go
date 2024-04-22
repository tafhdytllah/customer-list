package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tafhdytllah/customer-list/config"
	"github.com/tafhdytllah/customer-list/handler"
	"github.com/tafhdytllah/customer-list/repository"
	"github.com/tafhdytllah/customer-list/service"
)

func CustomerRouter(api *mux.Router) {
	repository := repository.NewCustomerRepository(config.DB)
	service := service.NewCustomerService(repository)
	handler := handler.NewCustomerHandler(service)

	r := api.PathPrefix("/customers").Subrouter()

	// Get Customer
	r.HandleFunc("/{customer_id}", handler.FindCustomerById).Methods(http.MethodGet)
	// Create Customer
	r.HandleFunc("/", handler.CreateCustomer).Methods(http.MethodPost)
	// Update Customer
	r.HandleFunc("/{customer_id}", handler.UpdateCustomerById).Methods(http.MethodPut)
	// Delete Family Member
	r.HandleFunc("/{customer_id}/{family_id}", handler.DeleteFamilyById).Methods(http.MethodDelete)
}
