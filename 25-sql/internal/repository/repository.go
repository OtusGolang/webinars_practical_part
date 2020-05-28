package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type BooksRepo interface {
	GetBooks(ctx context.Context) ([]Book, error)
}

type BaseRepo interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	BooksRepo
}

type Book struct {
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Meta      BookMeta
}

type BookMeta struct {
	Author string
}

func (b *BookMeta) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	d, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", src)
	}
	return json.Unmarshal(d, b)
}
