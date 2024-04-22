package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tafhdytllah/customer-list/config"
	"github.com/tafhdytllah/customer-list/router"
)

func main() {

	config.InitConfig()
	config.InitDB()

	r := mux.NewRouter()

	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	router.CustomerRouter(apiV1)

	address := config.C.ADDRESS
	port := config.C.PORT

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", address, port), r))

}
