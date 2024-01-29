package routers

import (
	"Book_api/pkg/handlers"

	"github.com/gorilla/mux"
)

func InitializeRoute() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/createbook", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/getbooks", handlers.GetAllBook).Methods("GET")
	r.HandleFunc("/getbook/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/updatebook/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/deletebook/{id}", handlers.DeleteBook).Methods("DELETE")
	return r

}
