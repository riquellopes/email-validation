package main

import (
	"github.com/gorilla/mux"

	"github.com/riquellopes/email-validation/api/views"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/validate/", views.EmailHandler).Methods("POST")
}
