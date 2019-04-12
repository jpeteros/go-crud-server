package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Price  float64       `json:"price" bson:"price"`
	Status bool          `json:"status" bson:"status"`
}

type User struct {
	UserId    float64 `json:"userId" bson:"userId"`
	Avatar    string  `json:"avatar" bson:"avatar"`
	FirstName string  `json:"firstName" bson:"firstName"`
	LastName  string  `json:"lastName" bson:"lastName"`
}
