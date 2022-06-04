package routes

import (
	"github.com/CcConStanTine/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBookById).Methods("DELETE")
}
