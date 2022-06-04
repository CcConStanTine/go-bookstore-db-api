package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CcConStanTine/go-bookstore/pkg/models"
	"github.com/CcConStanTine/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	allBooks := models.GetAllBooks()
	res, _ := json.Marshal(allBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID")
	}
	book, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	newBook = newBook.CreateBook()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID")
	}
	book := models.DeleteBook(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID")
	}
	bookFromDB, db := models.GetBookById(bookId)
	if updateBook.Name != "" {
		bookFromDB.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookFromDB.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookFromDB.Publication = updateBook.Publication
	}
	db.Save(&bookFromDB)
	res, _ := json.Marshal(bookFromDB)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
