package ports

import (
	"context"
	"go-code-snippets/internal/core/domain/dto"
)

type HttpPort interface {
	CreateBook(ctx context.Context, createBookdto dto.CreateBookDto) (interface{}, error)
	GetAllBooks(ctx context.Context) (interface{}, error)
	GetBookById(ctx context.Context, bookId string) (interface{}, error)
	UpdateBook(ctx context.Context, bookId string, updateBookDto dto.UpdateBookDto) (interface{}, error)
	DeleteBookById(ctx context.Context, bookId string) (interface{}, error)
}
