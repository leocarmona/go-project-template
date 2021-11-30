package mapper

import (
	"github.com/leocarmona/go-project-template/internal/app/domain/book"
	"github.com/leocarmona/go-project-template/internal/app/transport/inbound"
)

func BookFromCreateBookRequest(request *inbound.CreateBookRequest) *book.Book {
	return &book.Book{
		Name: request.Name,
	}
}

func BookFromUpdateBookRequest(id int64, request *inbound.UpdateBookRequest) *book.Book {
	return &book.Book{
		Id:   id,
		Name: request.Name,
	}
}
