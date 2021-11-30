package book

import (
	"context"
)

type Service struct {
	repository Repository
}

func (s *Service) Create(ctx context.Context, book *Book) error {
	return s.repository.Create(ctx, book)
}

func (s *Service) ReadById(ctx context.Context, id int64) (*Book, error) {
	return s.repository.ReadById(ctx, id)
}

func (s *Service) UpdateById(ctx context.Context, book *Book) (bool, error) {
	return s.repository.UpdateById(ctx, book)
}

func (s *Service) DeleteById(ctx context.Context, id int64) (bool, error) {
	return s.repository.DeleteById(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]*Book, error) {
	return s.repository.List(ctx)
}
