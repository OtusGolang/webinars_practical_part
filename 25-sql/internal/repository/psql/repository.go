package psql

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/OtusGolang/webinars_practical_part/25-sql/internal/repository"
)

var _ repository.BaseRepo = (*Repo)(nil)

type Repo struct {
	db *sql.DB
}

func (r *Repo) Connect(ctx context.Context, dsn string) (err error) {
	r.db, err = sql.Open("pgx", dsn)
	if err != nil {
		return
	}

	return r.db.PingContext(ctx)
}

func (r *Repo) Close() error {
	return r.db.Close()
}

func (r *Repo) GetBooks(ctx context.Context) ([]repository.Book, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT title, created_at, updated_at, meta FROM books
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []repository.Book

	for rows.Next() {
		var b repository.Book

		var updatedAt sql.NullTime

		if err := rows.Scan(
			&b.Title,
			&b.CreatedAt,
			&updatedAt,
			&b.Meta,
		); err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			b.UpdatedAt = updatedAt.Time
		}
		books = append(books, b)
	}
	return books, rows.Err()
}
