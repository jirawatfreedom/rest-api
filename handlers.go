package handlers

import (
	"github.com/gorilla/mux"
)

// Router route
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	return r
}
