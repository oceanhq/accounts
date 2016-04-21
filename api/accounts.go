package api

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func AccountGetByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	fmt.Fprintf(w, "Hello, "+username)
}

func AccountGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]

	fmt.Fprintf(w, "Hello, "+accountId)
}
