package controller

import "github.com/gorilla/mux"

func Routes(controller Controller) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books/", controller.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controller.GetBookById).Methods("GET")
	r.HandleFunc("/book/title", controller.SearchTitle).Methods("GET")
	r.HandleFunc("/book/desc", controller.GetBooksByCostDescending).Methods("GET")
	r.HandleFunc("/book/asc", controller.GetBooksByCostAscending).Methods("GET")
	r.HandleFunc("/books/{id}", controller.DeleteById).Methods("DELETE")
	r.HandleFunc("/books/", controller.AddBook).Methods("POST")
	r.HandleFunc("/books/{id}", controller.UpdateBookById).Methods("PUT")
	return r
}
