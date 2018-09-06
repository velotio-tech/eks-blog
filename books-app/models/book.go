package models

import "gopkg.in/mgo.v2/bson"

// Represents a Book, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Book struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Author string        `bson:"author" json:"author"`
	Price  float64       `bson:"price" json:"price"`
}
