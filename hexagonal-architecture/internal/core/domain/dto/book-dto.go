package dto

type CreateBookDto struct {
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
}

type UpdateBookDto struct {
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
}
