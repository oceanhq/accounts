package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/oceanhq/accounts/api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	}).Methods("GET")
	r.HandleFunc("/accounts", api.AccountGetByUsernameHandler).
		Methods("GET").
		Queries("username", "{username}")
	r.HandleFunc("/accounts/{accountId}", api.AccountGetByIdHandler).
		Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	n := negroni.New()
	n.UseHandler(r)
	n.Run(":" + port)
}
