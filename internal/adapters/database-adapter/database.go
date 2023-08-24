package databaseadapter

import (
	"context"
	"go-code-snippets/internal/core/domain/dto"
	"go-code-snippets/internal/core/domain/entity"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type dbAdapter struct {
	collection *mongo.Collection
}

func NewDBAdapter(collection *mongo.Collection) *dbAdapter {
	return &dbAdapter{
		collection: collection,
	}
}

func (d *dbAdapter) CreateBook(ctx context.Context, book entity.Book) (interface{}, error) {
	_, err := d.collection.InsertOne(ctx, book)
	if err != nil {
		log.Println("error inserting book:", err)
		return nil, err
	}

	return book.Id, nil
}

func (d *dbAdapter) GetBookById(ctx context.Context, bookId string) (interface{}, error) {
	book := entity.Book{}
	id, _ := primitive.ObjectIDFromHex(bookId)
	if err := d.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book); err != nil {
		log.Println("error getting or decoding book:", err)
		return nil, err
	}

	return book, nil
}

func (d *dbAdapter) GetAllBooks(ctx context.Context) (interface{}, error) {
	books := []entity.Book{}

	cur, err := d.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("error getting all books:", err)
		return nil, err
	}

	if err := cur.All(ctx, &books); err != nil {
		log.Println("error decoding all books:", err)
		return nil, err
	}

	return books, nil
}

func (d *dbAdapter) UpdateBook(ctx context.Context, bookId string, updateBookDto dto.UpdateBookDto) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	_, err := d.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"name":       updateBookDto.Name,
			"author":     updateBookDto.Author,
			"updated_on": time.Now().Format(time.RFC3339),
		},
	})

	if err != nil {
		log.Println("error updating book:", err)
		return err, nil
	}

	return bookId, nil
}

func (d *dbAdapter) DeleteBookById(ctx context.Context, bookId string) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	_, err := d.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println("error updating book:", err)
		return err, nil
	}

	return bookId, nil
}
