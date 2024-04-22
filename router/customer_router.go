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

	r.HandleFunc("/{customer_id}", handler.FindCustomerById).Methods(http.MethodGet)
	r.HandleFunc("/", handler.CreateCustomer).Methods(http.MethodPost)
}
