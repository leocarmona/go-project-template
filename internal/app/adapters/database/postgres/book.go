package postgres

import (
	"context"
	"database/sql"
	"github.com/leocarmona/go-project-template/internal/app/domain/book"
	"github.com/leocarmona/go-project-template/internal/infra/database"
)

type BookRepository struct {
	dbs *database.Databases
}

func NewBookRepository(dbs *database.Databases) *BookRepository {
	return &BookRepository{
		dbs: dbs,
	}
}

func (r *BookRepository) Create(ctx context.Context, book *book.Book) error {
	return r.dbs.Write.Connection().QueryRowContext(ctx, createBookSql, book.Name).Scan(&book.Id)
}

func (r *BookRepository) ReadById(ctx context.Context, id int64) (*book.Book, error) {
	var model book.Book
	err := r.dbs.Read.Connection().QueryRowContext(ctx, readBookByIdSql, id).
		Scan(&model.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	model.Id = id
	return &model, nil
}

func (r *BookRepository) UpdateById(ctx context.Context, book *book.Book) (bool, error) {
	result, err := r.dbs.Write.Connection().ExecContext(ctx, updateBookByIdSql, book.Name, book.Id)
	if err != nil {
		return false, nil
	}

	rows, err := result.RowsAffected()
	return rows > 0, err
}

func (r *BookRepository) DeleteById(ctx context.Context, id int64) (bool, error) {
	result, err := r.dbs.Write.Connection().ExecContext(ctx, deleteBookByIdSql, id)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	return rows > 0, err
}

func (r *BookRepository) List(ctx context.Context) ([]*book.Book, error) {
	rows, err := r.dbs.Read.Connection().QueryContext(ctx, listBookSql)
	if err != nil {
		return make([]*book.Book, 0), err
	}

	defer rows.Close()
	var books []*book.Book

	for rows.Next() {
		var b book.Book
		if err := rows.Scan(&b.Id, &b.Name); err != nil {
			return books, nil
		}

		books = append(books, &b)
	}

	if rows.Err() != nil {
		return books, err
	}

	return books, nil
}

const (
	createBookSql     = "insert into book (name) values ($1) RETURNING id"
	readBookByIdSql   = "select name from book where id = $1"
	updateBookByIdSql = "update book set name = $1 where id = $2"
	deleteBookByIdSql = "delete from book where id = $1"
	listBookSql       = "select id, name from book"
)
