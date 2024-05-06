package controllers

import (
    "encoding/json"
    "net/http"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
	"github.com/matheusmiguelsa/go-bookstore-mysql-gorilla/pkg/utils"
	"github.com/matheusmiguelsa/go-bookstore-mysql-gorilla/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err!= nil {
        fmt.Println("Error while parsing")
    }
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    bookId := vars["id"]
    ID, err := strconv.ParseInt(bookId,0,0)
    if err!= nil {
        fmt.Println("Error while parsing")
    }
    book := models.DeleteBook(ID)
    res, _ := json.Marshal(book)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
    bookId := vars["id"]
    ID, err := strconv.ParseInt(bookId,0,0)
    if err!= nil {
        fmt.Println("Error while parsing")
    }
    bookDetails, _ := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
	if updateBook.Publication != "" {
        bookDetails.Publication = updateBook.Publication
    }
	db.save(&bookDetails)
    res, _ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}