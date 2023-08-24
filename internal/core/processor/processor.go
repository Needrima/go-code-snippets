package processor

import (
	"context"
	"go-code-snippets/internal/core/domain/dto"
	"go-code-snippets/internal/core/helper"
	"go-code-snippets/internal/ports"
)

type processor struct {
	databasePort ports.DatabasePort
}

func NewProcessor(databasePort ports.DatabasePort) *processor {
	return &processor{
		databasePort: databasePort,
	}
}

func (p *processor) CreateBook(ctx context.Context, createBookdto dto.CreateBookDto) (interface{}, error) {
	book := helper.MapCreateBookDtoToBook(createBookdto)
	return p.databasePort.CreateBook(ctx, book)
}

func (p *processor) GetBookById(ctx context.Context, bookId string) (interface{}, error) {
	return p.databasePort.GetBookById(ctx, bookId)
}

func (p *processor) UpdateBook(ctx context.Context, bookId string, updateBookDto dto.UpdateBookDto) (interface{}, error) {
	return p.databasePort.UpdateBook(ctx, bookId, updateBookDto)
}

func (p *processor) DeleteBookById(ctx context.Context, bookId string) (interface{}, error) {
	return p.databasePort.DeleteBookById(ctx, bookId)
}

func (p *processor) GetAllBooks(ctx context.Context) (interface{}, error) {
	return p.databasePort.GetAllBooks(ctx)
}
