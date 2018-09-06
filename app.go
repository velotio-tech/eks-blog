package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "velotio.com/books-app/dao"
	. "velotio.com/books-app/models"
)

var dao = BooksDAO{}

// GET list of books
func AllBooksEndPoint(w http.ResponseWriter, r *http.Request) {
	books, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

// POST a new book
func CreateBookEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.ID = bson.NewObjectId()
	if err := dao.Insert(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, book)
}

// DELETE an existing book
func DeleteBookEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {

	dao.Server = "mongodb-service.default"
	dao.Database = "books_db"
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", AllBooksEndPoint).Methods("GET")
	r.HandleFunc("/books", CreateBookEndPoint).Methods("POST")
	r.HandleFunc("/books", DeleteBookEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
