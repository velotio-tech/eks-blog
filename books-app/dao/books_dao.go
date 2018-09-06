package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	. "velotio.com/books-app/models"
)

type BooksDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "books"
)

// Establish a connection to database
func (m *BooksDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of books
func (m *BooksDAO) FindAll() ([]Book, error) {
	var books []Book
	err := db.C(COLLECTION).Find(bson.M{}).All(&books)
	return books, err
}

// Insert a book into database
func (m *BooksDAO) Insert(book Book) error {
	err := db.C(COLLECTION).Insert(&book)
	return err
}

// Delete an existing book
func (m *BooksDAO) Delete(book Book) error {
	err := db.C(COLLECTION).Remove(&book)
	return err
}
