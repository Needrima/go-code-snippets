package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Author    string             `json:"author" bson:"author"`
	CreatedOn string             `json:"created_on" bson:"created_on"`
	UpdatedOn string             `json:"updated_on" bson:"updated_on"`
}
