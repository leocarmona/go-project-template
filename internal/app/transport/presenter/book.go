package presenter

import (
	"github.com/leocarmona/go-project-template/internal/app/domain/book"
	"github.com/leocarmona/go-project-template/internal/app/transport/outbound"
)

func CreateBook(book *book.Book) *outbound.CreateBookResponse {
	return &outbound.CreateBookResponse{
		Id:   book.Id,
		Name: book.Name,
	}
}

func ReadBookById(book *book.Book) *outbound.ReadBookByIdResponse {
	return &outbound.ReadBookByIdResponse{
		Id:   book.Id,
		Name: book.Name,
	}
}

func UpdateBookById(updated bool) *outbound.UpdateBookByIdResponse {
	return &outbound.UpdateBookByIdResponse{
		Updated: updated,
	}
}

func DeleteBookById(deleted bool) *outbound.DeleteBookByIdResponse {
	return &outbound.DeleteBookByIdResponse{
		Deleted: deleted,
	}
}

func ListBooks(books []*book.Book) []*outbound.ListBookResponse {
	var response = make([]*outbound.ListBookResponse, len(books))
	for index, b := range books {
		response[index] = &outbound.ListBookResponse{
			Id:   b.Id,
			Name: b.Name,
		}
	}
	return response
}
