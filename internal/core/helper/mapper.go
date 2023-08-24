package helper

import (
	"go-code-snippets/internal/core/domain/dto"
	"go-code-snippets/internal/core/domain/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCreateBookDtoToBook(createBookdto dto.CreateBookDto) entity.Book {
	return entity.Book{
		Id:        primitive.NewObjectID(),
		Name:      createBookdto.Name,
		Author:    createBookdto.Author,
		CreatedOn: time.Now().Format(time.RFC3339),
		UpdatedOn: time.Now().Format(time.RFC3339),
	}
}
