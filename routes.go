package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oceanhq/accounts/api"
)

func buildRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	}).Methods("GET")
	r.HandleFunc("/accounts", api.AccountGetByUsernameHandler).
		Methods("GET").
		Queries("username", "{username}")
	r.HandleFunc("/accounts/{accountId}", api.AccountGetByIdHandler).
		Methods("GET")

	return r
}
