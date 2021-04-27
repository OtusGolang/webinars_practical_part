package maindb

import (
	"context"
	"time"

	"github.com/OtusGolang/webinars_practical_part/19-clean-architecture/internal/domain/entities"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type PgEventStorage struct {
	db *sqlx.DB
}

func NewPgEventStorage(dsn string) (*PgEventStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgEventStorage{db: db}, nil
}

func (pges *PgEventStorage) SaveEvent(ctx context.Context, event *entities.Event) error {
	query := `
		INSERT INTO events(id, owner, title, text, start_time, end_time)
		VALUES (:id, :owner, :title, :text, :start_time, :end_time)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         event.ID.String(),
		"owner":      event.Owner,
		"title":      event.Title,
		"text":       event.Text,
		"start_time": event.StartTime,
		"end_time":   event.EndTime,
	})
	return err
}

func (pges *PgEventStorage) GetEventByID(ctx context.Context, id string) (*entities.Event, error) {
	var res entities.Event

	query := `
		SELECT
			*
		FROM events
		WHERE id = $1
	`
	return &res, pges.db.GetContext(ctx, &res, query, id)
}

func (pges *PgEventStorage) GetEventsByOwnerStartDate(ctx context.Context, owner string, startTime time.Time) ([]entities.Event, error) {
	var res []entities.Event

	query := `
		SELECT
			*
		FROM events
		WHERE owner = $1 AND start_time = $2
	`
	return res, pges.db.SelectContext(ctx, &res, query, owner, startTime)
}
