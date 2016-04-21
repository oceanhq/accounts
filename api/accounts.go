package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oceanhq/accounts/uuid"
)

func AccountGetByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	responseId := uuid.Generate()
	w.Header().Add("X-Response-ID", fmt.Sprint(responseId))

	fmt.Fprintf(w, "Hello, "+username)
}

func AccountGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountId"]

	responseId := uuid.Generate()
	w.Header().Add("X-Response-ID", fmt.Sprint(responseId))

	fmt.Fprintf(w, "Hello, "+accountId)
}
