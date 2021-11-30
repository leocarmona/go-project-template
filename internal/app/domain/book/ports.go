package book

import "context"

type Repository interface {
	Create(ctx context.Context, book *Book) error
	ReadById(ctx context.Context, id int64) (*Book, error)
	UpdateById(ctx context.Context, book *Book) (bool, error)
	DeleteById(ctx context.Context, id int64) (bool, error)
	List(ctx context.Context) ([]*Book, error)
}
